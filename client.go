package fourth_pos

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/textproto"
	"net/url"
	"path"
	"strings"

	"github.com/hashicorp/go-multierror"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-fourth-pos/" + libraryVersion
	mediaType      = "application/json"
	charset        = "utf-8"
)

var (
	BaseURL = url.URL{
		Scheme: "https",
		Host:   "api.fourth.com",
		Path:   "api/eposgateway",
	}
)

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{
		http: httpClient,
	}

	client.SetBaseURL(BaseURL)
	client.SetDebug(true)
	client.SetUserAgent(userAgent)
	client.SetMediaType(mediaType)
	client.SetCharset(charset)

	return client
}

// Client manages communication with InvoiceXpress Client
type Client struct {
	// HTTP client used to communicate with the Client.
	http *http.Client

	debug   bool
	baseURL url.URL

	// PMS Code
	pmsCode string

	// User agent for client
	userAgent string

	mediaType             string
	charset               string
	disallowUnknownFields bool

	// Optional function called after every successful request made to the DO Clients
	onRequestCompleted RequestCompletionCallback
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func (c *Client) Debug() bool {
	return c.debug
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c *Client) BaseURL() url.URL {
	return c.baseURL
}

func (c *Client) SetBaseURL(baseURL url.URL) {
	c.baseURL = baseURL
}

func (c *Client) SetMediaType(mediaType string) {
	c.mediaType = mediaType
}

func (c *Client) MediaType() string {
	return mediaType
}

func (c *Client) SetCharset(charset string) {
	c.charset = charset
}

func (c *Client) Charset() string {
	return charset
}

func (c *Client) SetPMSCode(pmsCode string) {
	c.pmsCode = pmsCode
}

func (c *Client) PMSCode() string {
	return c.pmsCode
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c *Client) UserAgent() string {
	return userAgent
}

func (c *Client) SetDisallowUnknownFields(disallowUnknownFields bool) {
	c.disallowUnknownFields = disallowUnknownFields
}

func (c *Client) GetEndpointURL(relative string, pathParams PathParams) url.URL {
	clientURL := c.BaseURL()
	relativeURL, err := url.Parse(relative)
	if err != nil {
		log.Fatal(err)
	}

	clientURL.Path = path.Join(clientURL.Path, relativeURL.Path)

	query := url.Values{}
	for k, v := range clientURL.Query() {
		query[k] = append(query[k], v...)
	}
	for k, v := range relativeURL.Query() {
		query[k] = append(query[k], v...)
	}
	clientURL.RawQuery = query.Encode()

	tmpl, err := template.New("endpoint_url").Parse(clientURL.Path)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	params := pathParams.Params()
	err = tmpl.Execute(buf, params)
	if err != nil {
		log.Fatal(err)
	}
	clientURL.Path = buf.String()

	return clientURL
}

func (c *Client) NewRequest(ctx context.Context, req Request) (*http.Request, error) {
	var contentType string

	// convert body struct to json
	buf := new(bytes.Buffer)

	if req.RequestBodyInterface() != nil {
		// Request has a method that returns a request body
		if r, ok := req.RequestBodyInterface().(io.Reader); ok {
			// request body is a io.Reader
			_, err := io.Copy(buf, r)
			if err != nil {
				return nil, err
			}
		} else {
			// request body is a struct/slice; marshal to json
			err := json.NewEncoder(buf).Encode(req.RequestBodyInterface())
			if err != nil {
				return nil, err
			}
		}
	} else if i, ok := req.(interface{ FormParamsInterface() Form }); ok {
		if i.FormParamsInterface().IsMultiPart() {
			// @TODO implement this as RequestBodyInterface()
			// Request has a form as body
			var err error
			w := multipart.NewWriter(buf)

			for k, f := range i.FormParamsInterface().Files() {
				var part io.Writer
				if x, ok := f.Content.(io.Closer); ok {
					defer x.Close()
				}

				if part, err = CreateFormFile(w, f.Content, k, f.Filename); err != nil {
					return nil, err
				}

				if _, err = io.Copy(part, f.Content); err != nil {
					return nil, err
				}
			}

			for k := range i.FormParamsInterface().Values() {
				var part io.Writer

				// Add other fields
				if part, err = w.CreateFormField(k); err != nil {
					return nil, err
				}

				fv := strings.NewReader(i.FormParamsInterface().Values().Get(k))
				if _, err = io.Copy(part, fv); err != nil {
					return nil, err
				}
			}

			// Don't forget to close the multipart writer.
			// If you don't close it, your request will be missing the terminating boundary.
			w.Close()

			// Don't forget to set the content type, this will contain the boundary.
			contentType = w.FormDataContentType()

		} else {
			buf.WriteString(i.FormParamsInterface().Values().Encode())
		}
	}

	// create new http request
	r, err := http.NewRequest(req.Method(), req.URL().String(), buf)
	if err != nil {
		return nil, err
	}

	// values := url.Values{}
	// err = utils.AddURLValuesToRequest(values, req, true)
	// if err != nil {
	// 	return nil, err
	// }

	// optionally pass along context
	if ctx != nil {
		r = r.WithContext(ctx)
	}

	// set other headers
	if contentType != "" {
		r.Header.Add("Content-Type", contentType)
	} else {
		r.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", c.MediaType(), c.Charset()))
	}
	r.Header.Add("Accept", c.MediaType())
	r.Header.Add("User-Agent", c.UserAgent())

	return r, nil
}

// Do sends an Client request and returns the Client response. The Client response is json decoded and stored in the value
// pointed to by v, or returned as an error if an Client error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, responseBody interface{}) (*http.Response, error) {
	if c.debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	httpResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, nil
	}

	// try to decode body into interface parameter
	if responseBody == nil {
		return httpResp, nil
	}
	errorResponse := &ErrorResponse{Response: httpResp}
	statusErrResponse := &StatusErrorResponse{Response: httpResp}
	err = c.Unmarshal(httpResp.Body, responseBody, errorResponse, statusErrResponse)
	if err != nil {
		return httpResp, err
	}

	if errorResponse.Error() != "" {
		return httpResp, errorResponse
	}

	if statusErrResponse.Error() != "" {
		return httpResp, statusErrResponse
	}

	return httpResp, nil
}

func (c *Client) Unmarshal(r io.Reader, vv ...interface{}) error {
	if len(vv) == 0 {
		return nil
	}

	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	errs := []error{}
	for _, v := range vv {
		r := bytes.NewReader(b)
		dec := json.NewDecoder(r)
		if c.disallowUnknownFields {
			dec.DisallowUnknownFields()
		}

		err := dec.Decode(v)
		if err != nil && err != io.EOF {
			errs = append(errs, err)
		}

	}

	if len(errs) == len(vv) {
		// Everything errored
		msgs := make([]string, len(errs))
		for i, e := range errs {
			msgs[i] = fmt.Sprint(e)
		}
		return errors.New(strings.Join(msgs, ", "))
	}

	return nil
}

// CheckResponse checks the Client response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. Client error responses are expected to have either no response
// body, or a json response body that maps to ErrorResponse. Any other response
// body will be silently ignored.
func CheckResponse(r *http.Response) error {
	errorResponse := &ErrorResponse{Response: r}

	if c := r.StatusCode; (c >= 200 && c <= 299) || c == 400 {
		return nil
	}

	err := checkContentType(r)
	if err != nil {
		return errors.New(r.Status)
	}

	// read data and copy it back
	data, err := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewReader(data))
	if err != nil {
		return errorResponse
	}

	if len(data) == 0 {
		return nil
	}

	// convert json to struct
	err = json.Unmarshal(data, errorResponse)
	if err != nil {
		return err
	}

	return errorResponse
}

type StatusErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response
}

func (r *StatusErrorResponse) Error() string {
	if r.Response.StatusCode != 0 && (r.Response.StatusCode < 200 || r.Response.StatusCode > 299) {
		return fmt.Sprintf(r.Response.Status)
	}

	return ""
}

// {
//   "errors": {
//       "sequenceId": [
//           "Error converting value \"\" to type 'System.Guid'. Path 'sequenceId', line 1, position 338."
//       ]
//   },
//   "type": "https://tools.ietf.org/html/rfc9110#section-15.5.1",
//   "title": "One or more validation errors occurred.",
//   "status": 400,
//   "traceId": "00-3c171ea3c263877259335dc4a4a1b970-ad0f4ea6ae7a2ebe-00"
// }

type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	ValidationErrors []ValidationError
	Errors           Errors `json:"errors"`
	Message          string `json:"message"`
	Type             string `json:"type"`
	Title            string `json:"title"`
	Status           int    `json:"status"`
	TraceID          string `json:"traceId"`
}

type Errors struct {
	SequenceID []string `json:"sequenceId"`
}

func (r *ErrorResponse) UnmarshalJSON(data []byte) error {
	var validationErrors []ValidationError
	if err := json.Unmarshal(data, &validationErrors); err == nil {
		r.ValidationErrors = validationErrors
		return nil
	}

	type Alias ErrorResponse
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}

	return json.Unmarshal(data, &aux)
}

func (r *ErrorResponse) Error() string {
	var result *multierror.Error

	// Collect validation errors
	for _, v := range r.ValidationErrors {
		if v.ErrorCode != nil && v.Message != "" {
			result = multierror.Append(result, fmt.Errorf("ErrorCode: %d, Message: %s", *v.ErrorCode, v.Message))
		}
	}

	// Add sequence ID errors
	for _, seq := range r.Errors.SequenceID {
		result = multierror.Append(result, fmt.Errorf("%s", seq))
	}

	// Add title error
	if r.Title != "" {
		result = multierror.Append(result, fmt.Errorf("%d: %s", r.Status, r.Title))
	}

	// Add message error
	if r.Message != "" {
		result = multierror.Append(result, errors.New(r.Message))
	}

	if result == nil {
		return ""
	}

	return result.Error()
}

// [
//
//	{
//	  "errorCode": 1001004,
//	  "errorScopeId": "9ab8f26a-e31d-4b36-b3f3-26c5c5340abd",
//	  "message": "Invalid Voucher type, valid Values (Unspecified,SupplierInvoice,AutoVat,SupplierInvoicePayment,CustomerInvoice,CustomerInvoicePayment",
//	  "sequentialId": "00000000-0000-0000-0000-000000000000"
//	}
//
// ]
type ValidationError struct {
	ErrorCode    *int   `json:"errorCode"`
	ErrorScopeID string `json:"errorScopeId"`
	Message      string `json:"message"`
	Reference    string `json:"reference"`
	SequentialID string `json:"sequentialId"`
	SourceHint   string `json:"sourceHint"`
}

func checkContentType(response *http.Response) error {
	header := response.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]
	if contentType != mediaType {
		return fmt.Errorf("expected content-type \"%s\", got \"%s\"", mediaType, contentType)
	}

	return nil
}

func CreateFormFile(w *multipart.Writer, data io.Reader, fieldname, filename string) (io.Writer, error) {
	var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

	escapeQuotes := func(s string) string {
		return quoteEscaper.Replace(s)
	}

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes(fieldname), escapeQuotes(filename)))

	// contentType, err := GetFileContentType(data)
	// if err != nil {
	// 	return nil, err
	// }
	h.Set("Content-Type", "text/csv")
	return w.CreatePart(h)
}

func GetFileContentType(file io.Reader) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
