package myexcel

import (
	"log"

	"github.com/xuri/excelize"
)

func CreateExcel() {
	file := excelize.NewFile()

	sheetIndex := file.NewSheet("Sheet1")
	file.SetActiveSheet(sheetIndex)

	err := file.SetCellValue("Sheet1", "B2", "Hello World")
	if err != nil {
		log.Fatalln(err)
	}

	err = file.SaveAs("Sample.xlsx")
	if err != nil {
		log.Fatalln(err)
	}
}
