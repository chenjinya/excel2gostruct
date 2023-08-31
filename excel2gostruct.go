package excel2gostruct

import (
	"encoding/json"
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

func GetColumnName(n int) string {
	columnName := ""
	for n > 0 {
		n--
		columnName = string('A'+n%26) + columnName
		n /= 26
	}
	return columnName
}

func ConvertExcelToGo(filepath string, targets interface{}, sheetIndex int) (err error) {
	if sheetIndex == 0 {
		sheetIndex = 1
	}
	if filepath == "" {
		return errors.New("filepath is empty")
	}
	if targets == nil {
		return errors.New("targets is nil")
	}
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		return err
	}
	var rowIndexToValue = make(map[string]string)
	var datasets []map[string]string
	sheetName := f.GetSheetName(sheetIndex)
	rows := f.GetRows(sheetName)
	for i, row := range rows {
		rawData := make(map[string]string)
		for j := range row {
			colName := GetColumnName(j + 1)
			axis := colName + strconv.Itoa(i+1)
			if i == 0 {
				name := f.GetCellValue(sheetName, axis)
				rowIndexToValue[colName] = name
			} else {
				rawData[rowIndexToValue[colName]] = f.GetCellValue(sheetName, axis)
			}
		}
		if i > 0 {
			datasets = append(datasets, rawData)
		}

	}
	ms, err := json.Marshal(datasets)
	if err != nil {
		return err
	}
	err = json.Unmarshal(ms, &targets)
	if err != nil {
		return err
	}
	return nil
}
