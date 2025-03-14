package fourth_pos

var (
	DimensionTypeOne DimensionType = "One"
	DimensionTypeTwo DimensionType = "Two"
)

var (
	DimensionStateActive   DimensionState = "Active"
	DimensionStateInactive DimensionState = "Inactive"
	DimensionStateAll      DimensionState = "All"
)

type DimensionType string
type DimensionState string

var (
	AccountStateActive AccountState = "Active"
	AccountStateAll    AccountState = "All"
)

type AccountState string

// var (
// 	ContactTypeCustomer ContactTypes = 2
// 	ContactTypeSupplier ContactTypes = 3
// )

// type ContactTypes int

type VatCodesResponse struct {
	ID                 string  `json:"id"`
	VatCode            int     `json:"vatCode"`
	VatCodeDescription string  `json:"vatCodeDescription"`
	VatPercent         float64 `json:"vatPercent"`
	VatType            int     `json:"vatType"`
}

type AccountsResponse struct {
	ID                string   `json:"id"`
	CreatedBy         string   `json:"createdBy"`
	UpdatedBy         string   `json:"updatedBy"`
	CreatedDate       DateTime `json:"createdDate"`
	UpdatedDate       DateTime `json:"updatedDate"`
	Version           int      `json:"version"`
	FinancialYearID   string   `json:"financialYearId"`
	OrganizationID    string   `json:"organizationId"`
	IsTemplateAccount bool     `json:"isTemplateAccount"`
	AccountPlanID     string   `json:"accountPlanId"`
	Number            int      `json:"number"`
	SequenceID        string   `json:"sequenceId"`
	RowVersion        int      `json:"rowVersion"`
	Name              string   `json:"name"`
	IsActive          bool     `json:"isActive"`
	VatCodeID         string   `json:"vatCodeId"`
}

type ContactsResponse struct {
	ID                       string   `json:"id"`
	CreatedBy                string   `json:"createdBy"`
	UpdatedBy                string   `json:"updatedBy"`
	CreatedDate              DateTime `json:"createdDate"`
	UpdatedDate              DateTime `json:"updatedDate"`
	Version                  int      `json:"version"`
	OrganizationID           string   `json:"organizationId"`
	ContactChannels          []any    `json:"contactChannels"`
	Addresses                []any    `json:"addresses"`
	Relations                []any    `json:"relations"`
	SupplierPaymentMethods   []any    `json:"supplierPaymentMethods"`
	CustomerRutRotApplicants []any    `json:"customerRutRotApplicants"`
	SequenceID               string   `json:"sequenceId"`
	RowVersion               int      `json:"rowVersion"`
	ContactTypes             []int    `json:"contactTypes"`
	BusinessType             int      `json:"businessType"`
	CountryCode              string   `json:"countryCode"`
	Country                  string   `json:"country"`
	CompanyName              string   `json:"companyName"`
	Number                   string   `json:"number"`
	GivenName                string   `json:"givenName"`
	SurName                  string   `json:"surName"`
	IdentificationNumber     string   `json:"identificationNumber"`
	Supplier                 Supplier `json:"supplier"`
	Customer                 Customer `json:"customer"`
	VatNumber                string   `json:"vatNumber"`
	Language                 string   `json:"language"`
	IsOrganization           bool     `json:"isOrganization"`
}
type Supplier struct {
	PaymentTerms           int    `json:"paymentTerms"`
	IsActive               bool   `json:"isActive"`
	SelectedCurrencyCode   string `json:"selectedCurrencyCode"`
	IsAutogiroActive       bool   `json:"isAutogiroActive"`
	DebtAccountNumber      any    `json:"debtAccountNumber"`
	IsAutomatedInvoiceFlow bool   `json:"isAutomatedInvoiceFlow"`
}
type EdiAddress struct {
	ID     string `json:"id"`
	Scheme int    `json:"scheme"`
}
type Customer struct {
	PaymentTerms                int        `json:"paymentTerms"`
	IsRutRot                    bool       `json:"isRutRot"`
	RutRotType                  int        `json:"rutRotType"`
	RutRotPropertyDescription   string     `json:"rutRotPropertyDescription"`
	RutRotBrfOrganizationNumber string     `json:"rutRotBrfOrganizationNumber"`
	RutRotApartmentNumber       string     `json:"rutRotApartmentNumber"`
	EInvoiceReceiverEmail       string     `json:"eInvoiceReceiverEmail"`
	InvoiceReference            string     `json:"invoiceReference"`
	DistributionMethod          int        `json:"distributionMethod"`
	IsActive                    bool       `json:"isActive"`
	SelectedCurrencyCode        string     `json:"selectedCurrencyCode"`
	EInvoicePartyID             string     `json:"eInvoicePartyId"`
	GlnNumber                   string     `json:"glnNumber"`
	IsReverseCharge             bool       `json:"isReverseCharge"`
	EdiAddress                  EdiAddress `json:"ediAddress"`
	CustomerDebtAccountNumber   int        `json:"customerDebtAccountNumber"`
	UseDiscount                 bool       `json:"useDiscount"`
	DiscountPercent             float64    `json:"discountPercent"`
	CfarNumber                  string     `json:"cfarNumber"`
}

type DimensionsResponse struct {
	ID              string   `json:"id"`
	Version         int      `json:"version"`
	CreatedBy       string   `json:"createdBy"`
	UpdatedBy       string   `json:"updatedBy"`
	CreatedDate     DateTime `json:"createdDate"`
	UpdatedDate     DateTime `json:"updatedDate"`
	OrganizationID  string   `json:"organizationId"`
	DimensionInfoID string   `json:"dimensionInfoId"`
	SequenceID      string   `json:"sequenceId"`
	RowVersion      int      `json:"rowVersion"`
	Number          string   `json:"number"`
	Type            int      `json:"type"`
	Name            string   `json:"name"`
	IsActive        bool     `json:"isActive"`
	StartDate       DateTime `json:"startDate"`
	EndDate         DateTime `json:"endDate"`
}

type VatSettingsResponse struct {
	ID              string   `json:"id"`
	CreatedBy       string   `json:"createdBy"`
	CreatedDate     DateTime `json:"createdDate"`
	UpdatedBy       string   `json:"updatedBy"`
	UpdatedDate     DateTime `json:"updatedDate"`
	Version         int      `json:"version"`
	VatValue        float64  `json:"vatValue"`
	OrganizationID  string   `json:"organizationId"`
	FinancialYearID string   `json:"financialYearId"`
	SequenceID      string   `json:"sequenceId"`
	RowVersion      int      `json:"rowVersion"`
	AccountNumber   int      `json:"accountNumber"`
	VatID           string   `json:"vatId"`
	VatType         int      `json:"vatType"`
}

type FinancialYearsResponse struct {
	ID             string   `json:"id"`
	Version        int      `json:"version"`
	CreatedBy      string   `json:"createdBy"`
	UpdatedBy      string   `json:"updatedBy"`
	CreatedDate    string   `json:"createdDate"`
	UpdatedDate    string   `json:"updatedDate"`
	OrganizationID string   `json:"organizationId"`
	SequenceID     string   `json:"sequenceId"`
	RowVersion     int      `json:"rowVersion"`
	StartDate      DateTime `json:"startDate"`
	EndDate        DateTime `json:"endDate"`
}

type VoucherDraftsResponse struct {
	VoucherID string `json:"voucherId"`
}

type CustomerInvoicesResponse struct {
	ID                        string                      `json:"id"`
	Version                   int                         `json:"version"`
	OrganizationID            string                      `json:"organizationId"`
	CreatedBy                 string                      `json:"createdBy"`
	CreatedDate               DateTime                    `json:"createdDate"`
	UpdatedBy                 string                      `json:"updatedBy"`
	UpdatedDate               DateTime                    `json:"updatedDate"`
	RutRot                    RutRot                      `json:"rutRot"`
	InvoiceLines              []InvoiceLinesResp          `json:"invoiceLines"`
	GrossAmount               float64                     `json:"grossAmount"`
	IsClosed                  bool                        `json:"isClosed"`
	IsBilled                  bool                        `json:"isBilled"`
	NetAmount                 float64                     `json:"netAmount"`
	VatAmount                 float64                     `json:"vatAmount"`
	AmountLeftToPay           float64                     `json:"amountLeftToPay"`
	AmountPaid                float64                     `json:"amountPaid"`
	CustomerInvoiceState      string                      `json:"customerInvoiceState"`
	VatDistributions          []VatDistributions          `json:"vatDistributions"`
	IsExpired                 bool                        `json:"isExpired"`
	IsCanceled                bool                        `json:"isCanceled"`
	CustomerInvoiceReferences []CustomerInvoiceReferences `json:"customerInvoiceReferences"`
	AmountToPay               float64                     `json:"amountToPay"`
	ReminderCount             int                         `json:"reminderCount"`
	IsCustomerPaid            bool                        `json:"isCustomerPaid"`
	InvoiceHeader             InvoiceHeaderResp           `json:"invoiceHeader"`
	PaymentChannels           []PaymentChannelsResp       `json:"paymentChannels"`
	Attachments               []Attachments               `json:"attachments"`
	CustomerInvoiceEvents     []CustomerInvoiceEvents     `json:"customerInvoiceEvents"`
	RoundingAmount            float64                     `json:"roundingAmount"`
	SequenceID                string                      `json:"sequenceId"`
	AccountingGrossAmount     float64                     `json:"accountingGrossAmount"`
	AccountingRoundingAmount  float64                     `json:"accountingRoundingAmount"`
	AccountingVatAmount       float64                     `json:"accountingVatAmount"`
	AccountingNetAmount       float64                     `json:"accountingNetAmount"`
	CurrencyInfo              CurrencyInfoResp            `json:"currencyInfo"`
	Dimensions                []DimensionsResp            `json:"dimensions"`
	RowVersion                int                         `json:"rowVersion"`
	ReminderFeeAmount         float64                     `json:"reminderFeeAmount"`
	Discount                  int                         `json:"discount"`
	NetAmountWithoutDiscount  int                         `json:"netAmountWithoutDiscount"`
	Organization              OrganizationResp            `json:"organization"`
	IsReverseCharge           bool                        `json:"isReverseCharge"`
	CustomerDebtAccountNumber int                         `json:"customerDebtAccountNumber"`
	ProjectNumber             string                      `json:"projectNumber"`
	AccountingMethod          string                      `json:"accountingMethod"`
	RoundingAccountNumber     int                         `json:"roundingAccountNumber"`
	ReminderFeeText           string                      `json:"reminderFeeText"`
}
type RutRotApplicantsResp struct {
	ID                   string              `json:"id"`
	CustomerInvoiceID    string              `json:"customerInvoiceId"`
	CustomerInvoice      CustomerInvoiceResp `json:"customerInvoice"`
	OrganizationID       string              `json:"organizationId"`
	Version              int                 `json:"version"`
	CreatedBy            string              `json:"createdBy"`
	CreatedDate          DateTime            `json:"createdDate"`
	UpdatedBy            string              `json:"updatedBy"`
	UpdatedDate          DateTime            `json:"updatedDate"`
	SequenceID           string              `json:"sequenceId"`
	RowVersion           int                 `json:"rowVersion"`
	IdentificationNumber string              `json:"identificationNumber"`
	Name                 string              `json:"name"`
	Amount               int                 `json:"amount"`
	ApplicationNumber    int                 `json:"applicationNumber"`
	UtcExportedDate      DateTime            `json:"utcExportedDate"`
}
type RutRotWorkTypeLinesResp struct {
	Version           int    `json:"version"`
	OrganizationID    string `json:"organizationId"`
	MaterialCost      int    `json:"materialCost"`
	WorkHours         int    `json:"workHours"`
	Name              string `json:"name"`
	CustomerInvoiceID string `json:"customerInvoiceId"`
}
type AddressResp struct {
	Street          string `json:"street"`
	StreetExtraLine string `json:"streetExtraLine"`
	ZipCode         string `json:"zipCode"`
	City            string `json:"city"`
	Country         string `json:"country"`
}
type RutRotResp struct {
	RutRotApplicants      []RutRotApplicantsResp    `json:"rutRotApplicants"`
	RutRotWorkTypeLines   []RutRotWorkTypeLinesResp `json:"rutRotWorkTypeLines"`
	AmountPaid            float64                   `json:"amountPaid"`
	IsRutRotPaid          bool                      `json:"isRutRotPaid"`
	ReminderText          string                    `json:"reminderText"`
	IsRutRot              bool                      `json:"isRutRot"`
	DebtAccountNumber     int                       `json:"debtAccountNumber"`
	Type                  string                    `json:"type"`
	ReductionAmount       float64                   `json:"reductionAmount"`
	PropertyDescription   string                    `json:"propertyDescription"`
	BrfOrganizationNumber string                    `json:"brfOrganizationNumber"`
	ApartmentNumber       string                    `json:"apartmentNumber"`
	InformationText       string                    `json:"informationText"`
	Address               AddressResp               `json:"address"`
	OtherCosts            int                       `json:"otherCosts"`
	TotalWorkCost         int                       `json:"totalWorkCost"`
}

type OrganizationResp struct {
	Name                 string                 `json:"name"`
	OrganizationNumber   string                 `json:"organizationNumber"`
	VatRegNumber         string                 `json:"vatRegNumber"`
	Address              AddressResp            `json:"address"`
	ContactChannels      []string               `json:"contactChannels"`
	TypedContactChannels []TypedContactChannels `json:"typedContactChannels"`
	IsFTaxApproved       bool                   `json:"isFTaxApproved"`
	SubjectToVAT         bool                   `json:"subjectToVAT"`
	Headquarters         string                 `json:"headquarters"`
}
type CustomerInvoiceResp struct {
	Organization              OrganizationResp `json:"organization"`
	IsReverseCharge           bool             `json:"isReverseCharge"`
	CustomerDebtAccountNumber int              `json:"customerDebtAccountNumber"`
	ProjectNumber             string           `json:"projectNumber"`
	AccountingMethod          string           `json:"accountingMethod"`
	RoundingAccountNumber     int              `json:"roundingAccountNumber"`
	ReminderFeeText           string           `json:"reminderFeeText"`
}
type InvoiceLinesResp struct {
	UnitPrice                float64      `json:"unitPrice"`
	UnitType                 string       `json:"unitType"`
	NetAmount                float64      `json:"netAmount"`
	NetAmountWithoutDiscount float64      `json:"netAmountWithoutDiscount"`
	VatAmount                float64      `json:"vatAmount"`
	GrossAmount              float64      `json:"grossAmount"`
	Quantity                 float64      `json:"quantity"`
	Vat                      Vat          `json:"vat"`
	SalesAccountNumber       int          `json:"salesAccountNumber"`
	RutRotWorkTypeName       string       `json:"rutRotWorkTypeName"`
	ArticleType              string       `json:"articleType"`
	StockAccountNumber       int          `json:"stockAccountNumber"`
	CostAccountNumber        int          `json:"costAccountNumber"`
	CostAmount               float64      `json:"costAmount"`
	Dimensions               []Dimensions `json:"dimensions"`
	ProjectNumber            string       `json:"projectNumber"`
	AccountingCostAmount     float64      `json:"accountingCostAmount"`
	AccountingGrossAmount    float64      `json:"accountingGrossAmount"`
	AccountingNetAmount      float64      `json:"accountingNetAmount"`
	AccountingVatAmount      float64      `json:"accountingVatAmount"`
	DiscountPercent          float64      `json:"discountPercent"`
	Discount                 float64      `json:"discount"`
	Name                     string       `json:"name"`
	ID                       string       `json:"id"`
	IsClosed                 bool         `json:"isClosed"`
	OrganizationID           string       `json:"organizationId"`
	CustomerInvoice          any          `json:"customerInvoice"`
	Type                     string       `json:"type"`
	LineIndex                int          `json:"lineIndex"`
}
type VatDistributions struct {
	OrganizationID string  `json:"organizationId"`
	Percent        float64 `json:"percent"`
	Amount         float64 `json:"amount"`
}
type CustomerInvoiceReferences struct {
	CustomerInvoiceID           string `json:"customerInvoiceId"`
	ReferencedCustomerInvoiceID string `json:"referencedCustomerInvoiceId"`
	Type                        string `json:"type"`
}
type ShippingAddressResp struct {
	Street          string `json:"street"`
	StreetExtraLine string `json:"streetExtraLine"`
	ZipCode         string `json:"zipCode"`
	City            string `json:"city"`
	Country         string `json:"country"`
}
type BillingAddressResp struct {
	Street          string `json:"street"`
	StreetExtraLine string `json:"streetExtraLine"`
	ZipCode         string `json:"zipCode"`
	City            string `json:"city"`
	Country         string `json:"country"`
}
type CustomerResp struct {
	Name                 string `json:"name"`
	Number               string `json:"number"`
	IdentificationNumber string `json:"identificationNumber"`
	VatNumber            string `json:"vatNumber"`
	CountryCode          string `json:"countryCode"`
}
type EdiAddressResp struct {
	ID     string `json:"id"`
	Scheme int    `json:"scheme"`
}
type InvoiceHeaderResp struct {
	ReminderText                string              `json:"reminderText"`
	InvoiceDate                 DateTime            `json:"invoiceDate"`
	InvoiceNumber               int                 `json:"invoiceNumber"`
	InvoiceExpirationDate       DateTime            `json:"invoiceExpirationDate"`
	ShippingAddress             ShippingAddressResp `json:"shippingAddress"`
	BillingAddress              BillingAddressResp  `json:"billingAddress"`
	Customer                    Customer            `json:"customer"`
	EdiAddress                  EdiAddressResp      `json:"ediAddress"`
	Comment                     string              `json:"comment"`
	OurReference                string              `json:"ourReference"`
	YourReference               string              `json:"yourReference"`
	DistributionMethod          string              `json:"distributionMethod"`
	DistributionEmail           string              `json:"distributionEmail"`
	DistributionEInvoicePartyID string              `json:"distributionEInvoicePartyId"`
	Language                    string              `json:"language"`
	CustomerOrderNumber         string              `json:"customerOrderNumber"`
	CustomerProjectNumber       string              `json:"customerProjectNumber"`
	Ocr                         string              `json:"ocr"`
	SupplierOrderNumber         string              `json:"supplierOrderNumber"`
	InvoiceReference            string              `json:"invoiceReference"`
	IsAutogiro                  bool                `json:"isAutogiro"`
}
type PaymentChannelsResp struct {
	OrganizationID string `json:"organizationId"`
	Type           string `json:"type"`
	Value          string `json:"value"`
	IsPrimary      bool   `json:"isPrimary"`
}
type Attachments struct {
	ID                string `json:"id"`
	CustomerInvoiceID string `json:"customerInvoiceId"`
	FileName          string `json:"fileName"`
	MimeType          string `json:"mimeType"`
	AttachmentType    string `json:"attachmentType"`
	FileSizeKb        int    `json:"fileSizeKb"`
	BlobName          string `json:"blobName"`
}
type DimensionsResp struct {
	DimensionNumber string `json:"dimensionNumber"`
	DimensionLevel  int    `json:"dimensionLevel"`
}
type CustomerInvoiceEventRows struct {
	AccountNumber  int              `json:"accountNumber"`
	Amount         int              `json:"amount"`
	ID             string           `json:"id"`
	Version        int              `json:"version"`
	CreatedBy      string           `json:"createdBy"`
	CreatedDate    DateTime         `json:"createdDate"`
	UpdatedBy      string           `json:"updatedBy"`
	UpdatedDate    DateTime         `json:"updatedDate"`
	OrganizationID string           `json:"organizationId"`
	ProjectNumber  string           `json:"projectNumber"`
	Dimensions     []DimensionsResp `json:"dimensions"`
}
type CustomerInvoiceEvents struct {
	CustomerInvoiceEventRows  []CustomerInvoiceEventRows `json:"customerInvoiceEventRows"`
	ID                        string                     `json:"id"`
	Version                   int                        `json:"version"`
	UpdatedDate               DateTime                   `json:"updatedDate"`
	UpdatedBy                 string                     `json:"updatedBy"`
	CreatedBy                 string                     `json:"createdBy"`
	CreatedDate               DateTime                   `json:"createdDate"`
	OrganizationID            string                     `json:"organizationId"`
	CustomerInvoiceID         string                     `json:"customerInvoiceId"`
	CustomerInvoiceEventType  string                     `json:"customerInvoiceEventType"`
	Date                      DateTime                   `json:"date"`
	CurrencyID                string                     `json:"currencyId"`
	ExchangeRate              int                        `json:"exchangeRate"`
	PaymentAccountNumber      int                        `json:"paymentAccountNumber"`
	Amount                    int                        `json:"amount"`
	InvoiceNumber             int                        `json:"invoiceNumber"`
	NotificationID            string                     `json:"notificationId"`
	NotificationStatusType    string                     `json:"notificationStatusType"`
	NotificationReceiverEmail string                     `json:"notificationReceiverEmail"`
	ReminderFeeAmount         int                        `json:"reminderFeeAmount"`
	VoucherID                 string                     `json:"voucherId"`
	VoucherSerie              string                     `json:"voucherSerie"`
	VoucherNumber             int                        `json:"voucherNumber"`
	SequenceID                string                     `json:"sequenceId"`
}
type CurrencyResp struct {
	Decimals int    `json:"decimals"`
	Code     string `json:"code"`
}
type AccountingCurrencyResp struct {
	Decimals int    `json:"decimals"`
	Code     string `json:"code"`
}
type CurrencyInfoResp struct {
	CurrencyProfitsAccountNumber int                    `json:"currencyProfitsAccountNumber"`
	CurrencyLossesAccountNumber  int                    `json:"currencyLossesAccountNumber"`
	ExchangeRate                 float64                `json:"exchangeRate"`
	RoundingFactor               float64                `json:"roundingFactor"`
	Currency                     CurrencyResp           `json:"currency"`
	AccountingCurrency           AccountingCurrencyResp `json:"accountingCurrency"`
}

// Requests

type VoucherDraftsRequest struct {
	VoucherDraft     VoucherDraft       `json:"voucherDraft"`
	VoucherDraftRows []VoucherDraftRows `json:"voucherDraftRows"`
	SequenceID       string             `json:"sequenceId"`
}

type VoucherDraft struct {
	VoucherDate  DateTime `json:"voucherDate"`
	VoucherType  string   `json:"voucherType"`
	Text         string   `json:"text"`
	CurrencyCode string   `json:"currencyCode"`
	Series       string   `json:"series"`
}

type VoucherDraftRows struct {
	Amount        float64      `json:"amount"`
	NumberOf      int          `json:"numberOf"`
	ProjectNumber string       `json:"projectNumber"`
	Specification string       `json:"specification"`
	Text          string       `json:"text"`
	AccountNumber int          `json:"accountNumber"`
	Dimensions    []Dimensions `json:"dimensions"`
}

type CustomerInvoicesRequest struct {
	Dimensions                []Dimensions      `json:"dimensions"`
	SequenceID                string            `json:"sequenceId"`
	InvoiceHeader             InvoiceHeader     `json:"invoiceHeader"`
	CurrencyInfo              CurrencyInfo      `json:"currencyInfo"`
	InvoiceLines              []InvoiceLines    `json:"invoiceLines"`
	RutRot                    RutRot            `json:"rutRot"`
	PaymentChannels           []PaymentChannels `json:"paymentChannels"`
	Organization              Organization      `json:"organization"`
	IsReverseCharge           bool              `json:"isReverseCharge"`
	CustomerDebtAccountNumber int               `json:"customerDebtAccountNumber"`
	ProjectNumber             string            `json:"projectNumber"`
	AccountingMethod          int               `json:"accountingMethod"`
	RoundingAccountNumber     int               `json:"roundingAccountNumber"`
}

type ShippingAddress struct {
	Street          string `json:"street"`
	StreetExtraLine string `json:"streetExtraLine"`
	ZipCode         string `json:"zipCode"`
	City            string `json:"city"`
	Country         string `json:"country"`
}
type BillingAddress struct {
	Street          string `json:"street"`
	StreetExtraLine string `json:"streetExtraLine"`
	ZipCode         string `json:"zipCode"`
	City            string `json:"city"`
	Country         string `json:"country"`
}
type CustomerForInvoice struct {
	Name                 string `json:"name"`
	Number               string `json:"number"`
	IdentificationNumber string `json:"identificationNumber"`
	VatNumber            string `json:"vatNumber"`
	CountryCode          string `json:"countryCode"`
}
type InvoiceHeader struct {
	InvoiceDate           DateTime           `json:"invoiceDate"`
	InvoiceExpirationDate DateTime           `json:"invoiceExpirationDate"`
	ShippingAddress       ShippingAddress    `json:"shippingAddress"`
	BillingAddress        BillingAddress     `json:"billingAddress"`
	Customer              CustomerForInvoice `json:"customer"`
	Comment               string             `json:"comment"`
	OurReference          string             `json:"ourReference"`
	YourReference         string             `json:"yourReference"`
	DistributionMethod    int                `json:"distributionMethod"`
	DistributionEmail     string             `json:"distributionEmail"`
	Language              string             `json:"language"`
	CustomerOrderNumber   string             `json:"customerOrderNumber"`
	CustomerProjectNumber string             `json:"customerProjectNumber"`
	SupplierOrderNumber   string             `json:"supplierOrderNumber"`
	InvoiceReference      string             `json:"invoiceReference"`
}
type Currency struct {
	Decimals int    `json:"decimals"`
	Code     string `json:"code"`
}
type AccountingCurrency struct {
	Decimals int    `json:"decimals"`
	Code     string `json:"code"`
}
type CurrencyInfo struct {
	CurrencyProfitsAccountNumber int                `json:"currencyProfitsAccountNumber"`
	CurrencyLossesAccountNumber  int                `json:"currencyLossesAccountNumber"`
	ExchangeRate                 int                `json:"exchangeRate"`
	RoundingFactor               int                `json:"roundingFactor"`
	Currency                     Currency           `json:"currency"`
	AccountingCurrency           AccountingCurrency `json:"accountingCurrency"`
}
type Vat struct {
	Percent       float64 `json:"percent"`
	AccountNumber int     `json:"accountNumber"`
}

type Dimensions struct {
	DimensionNumber string `json:"dimensionNumber"`
	DimensionLevel  int    `json:"dimensionLevel"`
	AccountType     int    `json:"accountType,omitempty"`
}
type InvoiceLines struct {
	UnitPrice          int          `json:"unitPrice"`
	UnitType           string       `json:"unitType"`
	Quantity           int          `json:"quantity"`
	Vat                Vat          `json:"vat"`
	SalesAccountNumber int          `json:"salesAccountNumber"`
	RutRotWorkTypeName string       `json:"rutRotWorkTypeName"`
	ArticleType        string       `json:"articleType"`
	StockAccountNumber int          `json:"stockAccountNumber"`
	CostAccountNumber  int          `json:"costAccountNumber"`
	CostAmount         int          `json:"costAmount"`
	Dimensions         []Dimensions `json:"dimensions"`
	ProjectNumber      string       `json:"projectNumber"`
	DiscountPercent    int          `json:"discountPercent"`
	Name               string       `json:"name"`
	Type               string       `json:"type"`
}
type RutRotApplicants struct {
	SequenceID           string   `json:"sequenceId"`
	IdentificationNumber string   `json:"identificationNumber"`
	Name                 string   `json:"name"`
	Amount               int      `json:"amount"`
	ApplicationNumber    int      `json:"applicationNumber"`
	UtcExportedDate      DateTime `json:"utcExportedDate"`
}
type RutRotWorkTypeLines struct {
	MaterialCost      int    `json:"materialCost"`
	WorkHours         int    `json:"workHours"`
	Name              string `json:"name"`
	CustomerInvoiceID string `json:"customerInvoiceId"`
}
type Address struct {
	Street          string `json:"street"`
	StreetExtraLine string `json:"streetExtraLine"`
	ZipCode         string `json:"zipCode"`
	City            string `json:"city"`
	Country         string `json:"country"`
}
type RutRot struct {
	RutRotApplicants      []RutRotApplicants    `json:"rutRotApplicants"`
	RutRotWorkTypeLines   []RutRotWorkTypeLines `json:"rutRotWorkTypeLines"`
	IsRutRot              bool                  `json:"isRutRot"`
	DebtAccountNumber     int                   `json:"debtAccountNumber"`
	Type                  int                   `json:"type"`
	ReductionAmount       int                   `json:"reductionAmount"`
	PropertyDescription   string                `json:"propertyDescription"`
	BrfOrganizationNumber string                `json:"brfOrganizationNumber"`
	ApartmentNumber       string                `json:"apartmentNumber"`
	InformationText       string                `json:"informationText"`
	Address               Address               `json:"address"`
	OtherCosts            int                   `json:"otherCosts"`
	TotalWorkCost         int                   `json:"totalWorkCost"`
}
type PaymentChannels struct {
	Type      string `json:"type"`
	Value     string `json:"value"`
	IsPrimary bool   `json:"isPrimary"`
}
type TypedContactChannels struct {
	Value              string `json:"value"`
	ContactChannelType int    `json:"contactChannelType"`
}
type Organization struct {
	Name                 string                 `json:"name"`
	OrganizationNumber   string                 `json:"organizationNumber"`
	VatRegNumber         string                 `json:"vatRegNumber"`
	Address              Address                `json:"address"`
	ContactChannels      []string               `json:"contactChannels"`
	TypedContactChannels []TypedContactChannels `json:"typedContactChannels"`
	IsFTaxApproved       bool                   `json:"isFTaxApproved"`
	SubjectToVAT         bool                   `json:"subjectToVAT"`
	Headquarters         string                 `json:"headquarters"`
}

type ContactsPost struct {
	ContactTypes         []int        `json:"contactTypes"`
	CountryCode          string       `json:"countryCode"`
	Country              string       `json:"country"`
	CompanyName          string       `json:"companyName"`
	Number               string       `json:"number"`
	GivenName            string       `json:"givenName"`
	SurName              string       `json:"surName"`
	IdentificationNumber string       `json:"identificationNumber"`
	Customer             CustomerPost `json:"customer"`
	VatNumber            string       `json:"vatNumber"`
	Language             string       `json:"language"`
}
type CustomerPost struct {
	PaymentTerms          string `json:"paymentTerms"`
	IsRutRot              bool   `json:"isRutRot"`
	EInvoiceReceiverEmail string `json:"eInvoiceReceiverEmail"`
	InvoiceReference      string `json:"invoiceReference"`
	IsActive              bool   `json:"isActive"`
	SelectedCurrencyCode  string `json:"selectedCurrencyCode"`
	IsReverseCharge       bool   `json:"isReverseCharge"`
}
