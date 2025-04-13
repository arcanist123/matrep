package engine

import (
	"fmt"
	"strings"
)

type Report_def struct {
	matrix [][]string
}

type csv_coordinate struct {
	row int
	col int
}
type report_page struct {
	tuple      []string
	position   csv_coordinate
	is_present bool
}
type report_rows struct {
	tuples     [][]string
	position   csv_coordinate
	is_present bool
}
type report_cols struct {
	tuples     [][]string
	position   csv_coordinate
	is_present bool
}
type report_result struct {
	cells    [][]string
	position csv_coordinate
}
type resolved_report struct {
	page   report_page
	rows   report_rows
	cols   report_cols
	result report_result
}

func NewReportDef(matrix [][]string) *Report_def {
	return &Report_def{matrix: matrix}
}

func (me Report_def) GetReportDefinition() (result *resolved_report, errs []error) {
	normalisedReport, errs := me.normaliseCellContents(me.matrix)
	if errs != nil {
		return nil, errs
	}
	fmt.Println(normalisedReport)
	return nil, nil
}
func (me Report_def) normaliseCellContents(matrix [][]string) (result [][]string, errs []error) {
	result = matrix
	for row, matrixRow := range matrix {
		for column, matrixCell := range matrixRow {
			//ensure that the string contains  equals sign and it is in a middle of the string
			if strings.Count(matrixCell, "=") != 0 || len(strings.Split(matrixCell, "=")) != 2 {
				errs = append(errs, fmt.Errorf("value in cell R%dC%d is not key=value pair", row, column))
			}
		}
	}
	return
}

func (me Report_def) get_page(matrix [][]string) (result report_page, err error) {

	for rowIndex, row := range matrix {
		if me.isRowEmpty(row) != true {
			report_values, is_resolved := me.resolve_row_to_page(row)
			if is_resolved {
				return report_page{
					tuple:      report_values,
					is_present: true,
					position: csv_coordinate{
						row: rowIndex,
						col: 0,
					},
				}, nil
			} else {
				return report_page{}, nil

			}
		}
	}

	return report_page{}, fmt.Errorf("report is empty")
}

func (me Report_def) isRowEmpty(row []string) bool {
	for _, cell := range row {
		if cell != "" {
			return false
		}
	}
	return true
}
func (me Report_def) resolve_row_to_page(row []string) (result []string, is_resolved bool) {
	keys := map[string]int{}
	for _, cell := range row {
		if cell != "" {
			cell_components := strings.Split(cell, "=")
			_, ok := keys[cell_components[0]]
			if ok {
				//if we have found an
				return []string{}, false
			} else {
				keys[cell_components[0]] = 0
			}

		}
	}
	return row, true
}
