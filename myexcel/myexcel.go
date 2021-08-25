package myexcel

import (
	"errors"
)

var errExcelWrite = errors.New("can't set value on excel sheet")
var errExcelSave = errors.New("can't save as excel file")

const SHEET_NUMBER = 10
const SAMPLE_NUMBER = 50000

func CreateExcelSample() {

}
