package fourth_pos

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TransactionDatasetTabParams struct {
	TransactionID        string
	Date                 DateTime
	TimeVal              DateTime
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

func (td *TransactionDataset) CreateTab(params TransactionDatasetTabParams) TransactionDatasetTab {
	tab := TransactionDatasetTab{
		SiteLocationCode:     td.SiteLocationCode,
		TransactionID:        params.TransactionID,
		TradingDate:          params.Date,
		Time:                 params.TimeVal,
		ReceiptCode:          params.ReceiptCode,
		CheckCode:            params.CheckCode,
		TableCode:            params.TableCode,
		RevenueCentreCode:    params.RevenueCentreCode,
		RevenueCentreDesc:    params.RevenueCentreDesc,
		Currency:             params.Currency,
		TabOwner:             params.TabOwner,
		TabOwnerDesc:         params.TabOwnerDesc,
		OriginalTabOwner:     params.OriginalTabOwner,
		OriginalTabOwnerDesc: params.OriginalTabOwnerDesc,
		TerminalCode:         params.TerminalCode,
		TerminalDesc:         params.TerminalDesc,
	}

	td.tabs = append(td.tabs, tab)
	return tab
}

func (td *TransactionDataset) AddRow(row TransactionDatasetRow) {
	td.rows = append(td.rows, row)
}

func (td *TransactionDataset) Output() (string, error) {
	var b strings.Builder
	writer := csv.NewWriter(&b)

	// Write headers
	if err := writer.Write(FOURTH_POS_GATEWAY_CSV_HEADERS); err != nil {
		return "", fmt.Errorf("failed to write headers: %w", err)
	}

	// Write tabs
	if err := td.writeTabs(writer); err != nil {
		return "", fmt.Errorf("failed to write tabs: %w", err)
	}

	// Write rows
	if err := td.writeRows(writer); err != nil {
		return "", fmt.Errorf("failed to write rows: %w", err)
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return "", fmt.Errorf("csv writer error: %w", err)
	}

	return b.String(), nil
}

func (td *TransactionDataset) Write(path string) error {
	output, err := td.Output()
	if err != nil {
		return fmt.Errorf("failed to generate output: %w", err)
	}

	return os.WriteFile(path, []byte(output), 0644)
}

func (td *TransactionDataset) writeTabs(writer *csv.Writer) error {
	for _, tab := range td.tabs {
		record := []string{
			tab.TransactionID,
			"", // UnitID placeholder
			tab.SiteLocationCode,
			tab.TradingDate.Format("2006-01-02"),
			tab.Time.Format("15:04:05"),
			"", // Timefact placeholder
			tab.TerminalCode,
			tab.TerminalDesc,
			"", // RecordActivityCode placeholder
			tab.ReceiptCode,
			tab.CheckCode,
			tab.TableCode,
			tab.RevenueCentreCode,
			tab.RevenueCentreDesc,
			tab.Currency,
			tab.TabOwner,
			tab.TabOwnerDesc,
			tab.OriginalTabOwner,
			tab.OriginalTabOwnerDesc,
		}

		if err := writer.Write(record); err != nil {
			return fmt.Errorf("failed to write tab record: %w", err)
		}
	}
	return nil
}

func (td *TransactionDataset) writeRows(writer *csv.Writer) error {
	for _, row := range td.rows {
		if err := td.writeRow(writer, row); err != nil {
			return fmt.Errorf("failed to write row: %w", err)
		}
	}
	return nil
}

func (td *TransactionDataset) writeRow(writer *csv.Writer, row TransactionDatasetRow) error {
	transactionID := fmt.Sprintf("%s_%s_%s_%s_%s_%d",
		row.CheckCode,
		row.SiteLocationCode,
		row.TransactionTypeCode,
		row.TradingDate.Format("020106"),
		row.TradingDate.Format("150405"),
		row.RecordActivityCode,
	)

	// Build the CSV record with all fields according to header order
	// Assuming FOURTH_POS_GATEWAY_CSV_HEADERS defines the correct order
	record := []string{
		transactionID,
		row.UnitID,
		row.SiteLocationCode,
		row.TradingDate.Format("2006-01-02"),
		row.Time.Format("15:04:05"),
		row.Timefact,
		row.TerminalCode,
		row.TerminalDesc,
		fmt.Sprintf("%d", row.RecordActivityCode),
		row.ReceiptCode,
		row.CheckCode,
		row.TableCode,
		row.RevenueCentreCode,
		row.RevenueCentreDesc,
		row.TransactionTypeCode,
		row.SalesItemID,
		row.SalesItemPLU,
		row.SalesItemGUID,
		row.SalesItemDesc,
		row.TenderTypeCode,
		row.TenderTypeDesc,
		row.DeductionCode,
		row.DeductionDesc,
		strconv.Itoa(row.Covers),
		strconv.FormatFloat(row.Qty, 'f', -1, 64),
		row.Currency,
		strconv.FormatFloat(row.ListPrice, 'f', -1, 64),
		strconv.FormatFloat(row.Tax, 'f', -1, 64),
		strconv.FormatFloat(row.PricePaid, 'f', -1, 64),
		strconv.FormatFloat(row.Deduction, 'f', -1, 64),
		strconv.FormatFloat(row.TenderAmount, 'f', -1, 64),
		strconv.FormatFloat(row.CostPriceTheo, 'f', -1, 64),
		strconv.FormatFloat(row.ListPriceConv, 'f', -1, 64),
		strconv.FormatFloat(row.TaxConv, 'f', -1, 64),
		strconv.FormatFloat(row.PricePaidConv, 'f', -1, 64),
		strconv.FormatFloat(row.DeductionConv, 'f', -1, 64),
		strconv.FormatFloat(row.TenderAmountConv, 'f', -1, 64),
		strconv.FormatFloat(row.CostPriceTheoConv, 'f', -1, 64),
		row.OrderTypeDesc,
		row.MenuBand,
		row.MajorGroupDesc,
		row.FamilyGroupDesc,
		row.SubGroupDesc,
		row.TabOwner,
		row.TabOwnerDesc,
		row.OriginalTabOwner,
		row.OriginalTabOwnerDesc,
		row.OldTableCode,
		row.PrevTransactionCode,
		row.AuthorisedBy,
		row.TextField,
		row.GuestDesc,
		row.TimeSentToPrep,
		row.BumpTime,
		row.UniversalTimesLotID,
		row.TimeSentToPrep,
		row.BumpTime,
		row.TransactionStartEnd,
		row.IsDeleted,
		row.CustomField1,
		row.CustomField2,
		row.CustomField3,
		row.CustomFact1,
		row.CustomFact2,
		row.DateFact,
	}

	if err := writer.Write(record); err != nil {
		return fmt.Errorf("failed to write row record: %w", err)
	}
	return nil
}
