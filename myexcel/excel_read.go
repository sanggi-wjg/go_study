package myexcel

import (
	"errors"
	"fmt"
	"time"

	"github.com/xuri/excelize"
)

var errExcelOpen = errors.New("can't open file")
var errExcelSave = errors.New("can't save as excel file")

type InvoiceResult struct {
	invoiceNo string
	toWhere   string
	weight    string
	cost      string
}

func NewInvoiceResult(invoiceNo string, toWhere string, weight string, cost string) *InvoiceResult {
	res := InvoiceResult{invoiceNo: invoiceNo, toWhere: toWhere, weight: weight, cost: cost}
	return &res
}

func openfile(filepath string) ([][]string, error) {
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println("OpenFile error:", err)
		return nil, err
	}

	rows, err := f.GetRows("账单明细")
	if err != nil {
		fmt.Println("GetRows error:", err)
		return nil, err
	}
	return rows, nil
}

func readRows(rows [][]string) []InvoiceResult {
	invoiceResults := []InvoiceResult{}

	for _, row := range rows {
		res := InvoiceResult{row[2], row[4], row[5], row[6]}
		invoiceResults = append(invoiceResults, res)
	}

	return invoiceResults
}

func ExcelReadMain() {
	startTime := time.Now()

	rows, _ := openfile("myexcel/data/YTO-2021-10-SHA.xlsx")
	readRows(rows)

	// for _, v := range results {
	// 	if v.invoiceNo == "YTG003192303174" {
	// 		fmt.Print(v)
	// 		break
	// 	}
	// }

	finishTime := time.Since(startTime)
	fmt.Println(finishTime)
}
