package myexcel

import (
	"errors"

	"github.com/xuri/excelize"
)

var errExcelWrite = errors.New("can't set value on excel sheet")
var errExcelSave = errors.New("can't save as excel file")

func CreateExcel() error {
	file := excelize.NewFile()

	sheetIndex := file.NewSheet("Sheet1")
	file.SetActiveSheet(sheetIndex)

	err := file.SetCellValue("Sheet1", "B2", "Hello World")
	if err != nil {
		return errExcelWrite
	}

	err = file.SaveAs("Sample.xlsx")
	if err != nil {
		return errExcelSave
	}

	return nil
}
