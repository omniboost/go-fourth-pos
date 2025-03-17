package fourth_pos

var FOURTH_POS_GATEWAY_CSV_HEADERS = []string{
	"transactionid",
	"unitid",
	"sitelocationcode",
	"tradingdate",
	"time",
	"timefact",
	"terminalcode",
	"terminaldesc",
	"recordactivitycode",
	"receiptcode",
	"checkcode",
	"tablecode",
	"revenuecentrecode",
	"revenuecentredesc",
	"transactiontypecode",
	"salesitemid",
	"salesitemPLU",
	"salesitemGUID",
	"salesitemdesc",
	"tendertypecode",
	"tendertypedesc",
	"deductioncode",
	"deductiondesc",
	"covers",
	"qty",
	"currency",
	"listprice",
	"tax",
	"pricepaid",
	"deduction",
	"tenderamount",
	"costpricetheo",
	"listpriceconv",
	"taxconv",
	"pricepaidconv",
	"deductionconv",
	"tenderamountconv",
	"costpricetheoconv",
	"ordertypedesc",
	"menuband",
	"majorgroupdesc",
	"familygroupdesc",
	"subgroupdesc",
	"tabowner",
	"tabownerdesc",
	"originaltabowner",
	"originaltabownerdesc",
	"oldtablecode",
	"prevtransactioncode",
	"authorisedby",
	"textfield",
	"guestdesc",
	"guestcode",
	"timesenttoprep",
	"bumptime",
	"universaltimeslotid",
	"timeslotdesc",
	"transactionstartend",
	"IsDeleted",
	"customfield1",
	"customfield2",
	"customfield3",
	"customfact1",
	"customfact2",
	"datefact",
}

type TransactionDataset struct {
	SiteLocationCode string
	tabs             []TransactionDatasetTab
	rows             []TransactionDatasetRow
}

type TransactionDatasetTab struct {
	SiteLocationCode     string
	TransactionID        string
	TradingDate          DateTime
	Time                 DateTime
	ReceiptCode          string
	CheckCode            string
	TableCode            string
	RevenueCentreCode    string
	RevenueCentreDesc    string
	Currency             string
	TabOwner             string
	TabOwnerDesc         string
	OriginalTabOwner     string
	OriginalTabOwnerDesc string
	TerminalCode         string
	TerminalDesc         string
}

type TransactionDatasetRow struct {
	TransactionID        string
	CheckCode            string
	SiteLocationCode     string
	TransactionTypeCode  string
	TradingDate          DateTime
	Time                 DateTime
	RecordActivityCode   int
	UnitID               string
	Timefact             string
	TerminalCode         string
	TerminalDesc         string
	ReceiptCode          string
	TableCode            string
	RevenueCentreCode    string
	RevenueCentreDesc    string
	SalesItemID          string
	SalesItemPLU         string
	SalesItemGUID        string
	SalesItemDesc        string
	TenderTypeCode       string
	TenderTypeDesc       string
	DeductionCode        string
	DeductionDesc        string
	Covers               int
	Qty                  float64
	Currency             string
	ListPrice            float64
	Tax                  float64
	PricePaid            float64
	Deduction            float64
	TenderAmount         float64
	CostPriceTheo        float64
	ListPriceConv        float64
	TaxConv              float64
	PricePaidConv        float64
	DeductionConv        float64
	TenderAmountConv     float64
	CostPriceTheoConv    float64
	OrderTypeDesc        string
	MenuBand             string
	MajorGroupDesc       string
	FamilyGroupDesc      string
	SubGroupDesc         string
	TabOwner             string
	TabOwnerDesc         string
	OriginalTabOwner     string
	OriginalTabOwnerDesc string
	OldTableCode         string
	PrevTransactionCode  string
	AuthorisedBy         string
	TextField            string
	GuestDesc            string
	TimeSentToPrep       string
	BumpTime             string
	UniversalTimesLotID  string
	TimeslotDesc         string
	TransactionStartEnd  string
	IsDeleted            string
	CustomField1         string
	CustomField2         string
	CustomField3         string
	CustomFact1          string
	CustomFact2          string
	DateFact             string
}

func NewTransactionDataset(siteLocationCode string) *TransactionDataset {
	return &TransactionDataset{
		SiteLocationCode: siteLocationCode,
		tabs:             []TransactionDatasetTab{},
		rows:             []TransactionDatasetRow{},
	}
}

type SendFilesResponse struct {
	Name             string `json:"Name"`
	ContainerName    string `json:"ContainerName"`
	Size             int    `json:"Size"`
	ContentType      string `json:"ContentType"`
	Location         string `json:"Location"`
	ModifiedDateTime string `json:"ModifiedDateTime"`
	DataLoadBatchID  string `json:"DataLoadBatchId"`
}
