package engine

import (
	"fmt"
	"strings"

	"github.com/arcanist123/matrep/config"
)

type MatrixHandler struct {
	matrix [][]string
	config config.Config
}

func NewMatrixHandler(matrix [][]string, config config.Config) *MatrixHandler {
	return &MatrixHandler{matrix: matrix, config: config}
}

func (me MatrixHandler) GetReportWithData() (matrix [][]string, errors []error) {
	normalisedMatrix, errors := me.normaliseCellContents(me.matrix)

	return normalisedMatrix, errors
}

func (me MatrixHandler) normaliseCellContents(matrix [][]string) (result [][]string, err []error) {
	result = matrix
	for row, matrixRow := range matrix {
		for column, matrixCell := range matrixRow {
			if strings.Count(matrixCell, "=") != 0 {
				err = append(err, fmt.Errorf("value in cell R%dC%d is not key=value pair", row, column))
			}
		}
	}
	return
}

func (me MatrixHandler) getPage(matrix [][]string) ([]string, int, error) {
	var finalPosition int
	// for position, element := range matrix {
	// 	finalPosition = position

	// }

	return []string{}, finalPosition, nil
}
