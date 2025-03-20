package fourth_pos

type TransactionDatasetRows []TransactionDatasetRow

type TransactionDatasetRow struct {
	TransactionID        string   `csv:"transactionid"`
	UnitID               string   `csv:"unitid"`
	SiteLocationCode     string   `csv:"sitelocationcode"`
	TradingDate          DateTime `csv:"tradingdate"`
	Time                 DateTime `csv:"time"`
	Timefact             string   `csv:"timefact"`
	TerminalCode         string   `csv:"terminalcode"`
	TerminalDesc         string   `csv:"terminaldesc"`
	RecordActivityCode   int      `csv:"recordactivitycode,omitempty"`
	ReceiptCode          string   `csv:"receiptcode"`
	CheckCode            string   `csv:"checkcode"`
	TableCode            string   `csv:"tablecode"`
	RevenueCentreCode    string   `csv:"revenuecentrecode"`
	RevenueCentreDesc    string   `csv:"revenuecentredesc"`
	TransactionTypeCode  string   `csv:"transactiontypecode"`
	SalesItemID          string   `csv:"salesitemid"`
	SalesItemPLU         string   `csv:"salesitemPLU"`
	SalesItemGUID        string   `csv:"salesitemGUID"`
	SalesItemDesc        string   `csv:"salesitemdesc"`
	TenderTypeCode       string   `csv:"tendertypecode"`
	TenderTypeDesc       string   `csv:"tendertypedesc"`
	DeductionCode        string   `csv:"deductioncode"`
	DeductionDesc        string   `csv:"deductiondesc"`
	Covers               int      `csv:"covers"`
	Qty                  float64  `csv:"qty"`
	Currency             string   `csv:"currency"`
	ListPrice            float64  `csv:"listprice"`
	Tax                  float64  `csv:"tax"`
	PricePaid            float64  `csv:"pricepaid"`
	Deduction            float64  `csv:"deduction"`
	TenderAmount         float64  `csv:"tenderamount"`
	CostPriceTheo        float64  `csv:"costpricetheo"`
	ListPriceConv        float64  `csv:"listpriceconv"`
	TaxConv              float64  `csv:"taxconv"`
	PricePaidConv        float64  `csv:"pricepaidconv"`
	DeductionConv        float64  `csv:"deductionconv"`
	TenderAmountConv     float64  `csv:"tenderamountconv"`
	CostPriceTheoConv    float64  `csv:"costpricetheoconv"`
	OrderTypeDesc        string   `csv:"ordertypedesc"`
	MenuBand             string   `csv:"menuband"`
	MajorGroupDesc       string   `csv:"majorgroupdesc"`
	FamilyGroupDesc      string   `csv:"familygroupdesc"`
	SubGroupDesc         string   `csv:"subgroupdesc"`
	TabOwner             string   `csv:"tabowner"`
	TabOwnerDesc         string   `csv:"tabownerdesc"`
	OriginalTabOwner     string   `csv:"originaltabowner"`
	OriginalTabOwnerDesc string   `csv:"originaltabownerdesc"`
	OldTableCode         string   `csv:"oldtablecode"`
	PrevTransactionCode  string   `csv:"prevtransactioncode"`
	AuthorisedBy         string   `csv:"authorisedby"`
	TextField            string   `csv:"textfield"`
	GuestDesc            string   `csv:"guestdesc"`
	GuestCode            string   `csv:"guestcode"`
	TimeSentToPrep       string   `csv:"timesenttoprep"`
	BumpTime             string   `csv:"bumptime"`
	UniversalTimesLotID  string   `csv:"universaltimeslotid"`
	TimeslotDesc         string   `csv:"timeslotdesc"`
	TransactionStartEnd  string   `csv:"transactionstartend"`
	IsDeleted            string   `csv:"isdeleted"`
	CustomField1         string   `csv:"customfield1"`
	CustomField2         string   `csv:"customfield2"`
	CustomField3         string   `csv:"customfield3"`
	CustomFact1          string   `csv:"customfact1"`
	CustomFact2          string   `csv:"customfact2"`
	DateFact             string   `csv:"datefact"`
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
