package fourth_pos

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-fourth-pos/utils"
)

func (c *Client) NewSendFilesRequest() SendFilesRequest {
	return SendFilesRequest{
		client:      c,
		queryParams: c.NewSendFilesQueryParams(),
		pathParams:  c.NewSendFilesPathParams(),
		formParams:  c.NewSendFilesFormParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewSendFilesRequestBody(),
	}
}

type SendFilesRequest struct {
	client      *Client
	queryParams *SendFilesQueryParams
	pathParams  *SendFilesPathParams
	formParams  *SendFilesFormParams
	method      string
	headers     http.Header
	requestBody SendFilesRequestBody
}

func (c *Client) NewSendFilesQueryParams() *SendFilesQueryParams {
	return &SendFilesQueryParams{}
}

type SendFilesQueryParams struct{}

type SendFilesFormParams struct {
	ExtfFile FormFile
}

func (p SendFilesFormParams) IsMultiPart() bool {
	return true
}

func (p SendFilesFormParams) Files() map[string]FormFile {
	return map[string]FormFile{
		"file": p.ExtfFile,
	}
}

func (p SendFilesFormParams) Values() url.Values {
	return url.Values{}
}

func (p SendFilesQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *SendFilesRequest) QueryParams() *SendFilesQueryParams {
	return r.queryParams
}

func (r *SendFilesRequest) FormParams() *SendFilesFormParams {
	return r.formParams
}

func (r *SendFilesRequest) FormParamsInterface() Form {
	return r.formParams
}

func (c *Client) NewSendFilesPathParams() *SendFilesPathParams {
	return &SendFilesPathParams{}
}

func (c *Client) NewSendFilesFormParams() *SendFilesFormParams {
	return &SendFilesFormParams{}
}

type SendFilesPathParams struct {
}

func (p *SendFilesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *SendFilesRequest) PathParams() *SendFilesPathParams {
	return r.pathParams
}

func (r *SendFilesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *SendFilesRequest) SetMethod(method string) {
	r.method = method
}

func (r *SendFilesRequest) Method() string {
	return r.method
}

func (s *Client) NewSendFilesRequestBody() SendFilesRequestBody {
	return SendFilesRequestBody{}
}

type SendFilesRequestBody struct{}

func (r *SendFilesRequest) RequestBody() *SendFilesRequestBody {
	return nil
}

func (r *SendFilesRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *SendFilesRequest) SetRequestBody(body SendFilesRequestBody) {
	r.requestBody = body
}

func (r *SendFilesRequest) NewResponseBody() *SendFilesResponseBody {
	return &SendFilesResponseBody{}
}

type SendFilesResponseBody []SendFilesResponse

func (r *SendFilesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api/files", r.PathParams())
	return &u
}

func (r *SendFilesRequest) Do() (SendFilesResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	req.Header.Add("Filename", r.FormParams().ExtfFile.Filename)
	req.Header.Set("Content-Disposition", "Fourth-POS-File")

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, true)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
