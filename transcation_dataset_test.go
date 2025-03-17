package fourth_pos_test

import (
	"log"
	"os"
	"testing"
	"time"

	fourth_pos "github.com/omniboost/go-fourth-pos"
)

func TestTransactionDatasetExport(t *testing.T) {
	dataset := fourth_pos.NewTransactionDataset("TEST01")

	log.Println(dataset)

	tradingDate, _ := time.Parse("2006-01-02 15:04:05", "2025-03-17 12:30:00")
	dateTime := fourth_pos.DateTime{Time: tradingDate}

	dataset.CreateTab(fourth_pos.TransactionDatasetTabParams{
		TransactionID:        "1234_130500",
		Date:                 dateTime,
		TimeVal:              dateTime,
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
	})

	row := fourth_pos.TransactionDatasetRow{
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
	}
	dataset.AddRow(row)

	testFilePath := "./test_transaction_export2.csv"

	err := dataset.Write(testFilePath)
	if err != nil {
		log.Fatalf("Failed to write CSV file: %v", err)
	}

	content, err := os.ReadFile(testFilePath)
	if err != nil {
		log.Fatalf("Failed to read CSV file: %v", err)
	}

	if len(content) == 0 {
		log.Fatal("CSV file is empty")
	}

	log.Printf("CSV file successfully exported to: %s", testFilePath)
	log.Printf("File size: %d bytes", len(content))
}
