package excel

import (
	"fmt"
	"log"

	"github.com/xuri/excelize"
)

func createExcel() {
	file, err := excelize.OpenFile("sample.xlsx")
	if err != nil {
		log.Fatalln(err)
	}

	cell, err := file.GetCellValue("Sheet1", "B2")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cell)
}

func testFunc() {
	fmt.Println("TEST")
}
