package fourth_pos_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gocarina/gocsv"
	fourth_pos "github.com/omniboost/go-fourth-pos"
)

func TestTransactionDatasetRowsCSVMarshalling(t *testing.T) {

	tradingDate, _ := time.Parse("2006-01-02 15:04:05", "2025-03-17 12:30:00")
	dateTime := fourth_pos.DateTime{Time: tradingDate}
	rows := fourth_pos.TransactionDatasetRows{
		{
			TransactionID:        "1234_130500",
			TradingDate:          dateTime,
			Time:                 dateTime,
			ReceiptCode:          "",
			CheckCode:            "28581",
			TableCode:            "20",
			RevenueCentreCode:    "2",
			RevenueCentreDesc:    "Restaurant",
			Currency:             "USD",
			TabOwner:             "138992",
			TabOwnerDesc:         "KAREN",
			OriginalTabOwner:     "138992",
			OriginalTabOwnerDesc: "KAREN",
			TerminalCode:         "",
			TerminalDesc:         "TILL 3",
		},
		{
			CheckCode:            "1234_135700_SALES_ITEM_28581_2",
			SiteLocationCode:     "1234",
			TransactionTypeCode:  "SALES_ITEM",
			TradingDate:          dateTime,
			Time:                 dateTime,
			RecordActivityCode:   2,
			UnitID:               "",
			Timefact:             "",
			TerminalCode:         "",
			TerminalDesc:         "TILL 3",
			ReceiptCode:          "",
			TableCode:            "20",
			RevenueCentreCode:    "2",
			RevenueCentreDesc:    "Restaurant",
			Currency:             "USD",
			TabOwner:             "138992",
			TabOwnerDesc:         "KAREN",
			OriginalTabOwner:     "138992",
			OriginalTabOwnerDesc: "KAREN",
			SalesItemDesc:        "Margherita",
			Qty:                  2,
			ListPrice:            20,
			PricePaid:            20,
			Tax:                  3.33,
			SalesItemPLU:         "24789",
		},
	}

	csv, err := gocsv.MarshalString(rows)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(csv)
}
