package myexcel

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/xuri/excelize"
)

var errExcelWrite = errors.New("can't set value on excel sheet")
var errExcelSave = errors.New("can't save as excel file")

const SHEET_NUMBER = 10
const SAMPLE_NUMBER = 50000

func CreateExcelSample() error {
	file := excelize.NewFile()

	for i := 1; i <= SHEET_NUMBER; i++ {
		CreateSheet(file, "Sheet"+strconv.Itoa(i))
	}

	sampleDatas := sampleData()
	sheetMap := file.GetSheetMap()

	for _, sheetName := range sheetMap {
		channel := make(chan error)
		channelResult := make([]error, SAMPLE_NUMBER)

		for index, sample := range sampleDatas {
			go WriteCellString(file, sheetName, index, sample, channel)
		}

		for i := 0; i < SAMPLE_NUMBER; i++ {
			channelResult = append(channelResult, <-channel)
		}
	}

	err := SaveExcelFile(file, "Sample.xlsx")
	if err != nil {
		return err
	}

	return nil
}

func CreateSheet(f *excelize.File, sheetName string) int {
	defer fmt.Println("create sheet " + sheetName)
	index := f.NewSheet(sheetName)
	return index
}

func WriteCellString(f *excelize.File, sheetName string, col int, value string, c chan error) {
	// fmt.Println(sheetName, col)

	err := f.SetCellValue(sheetName, "B"+strconv.Itoa(col+1), value)
	if err != nil {
		c <- errExcelWrite
	}

	err = f.SetCellValue(sheetName, "C"+strconv.Itoa(col+1), value)
	if err != nil {
		c <- errExcelWrite
	}

	err = f.SetCellValue(sheetName, "D"+strconv.Itoa(col+1), value)
	if err != nil {
		c <- errExcelWrite
	}

	c <- nil
}

func SaveExcelFile(f *excelize.File, filename string) error {
	err := f.SaveAs(filename)
	if err != nil {
		return errExcelSave
	}

	return nil
}

func sampleData() []string {
	s := make([]string, SAMPLE_NUMBER)
	for i := 0; i < SAMPLE_NUMBER; i++ {
		s[i] = strconv.Itoa(i + 1)
	}
	return s
}
