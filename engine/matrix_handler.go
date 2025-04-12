package engine

import (
	"github.com/arcanist123/matrep/config"
)

type MatrixHandler struct {
	matrix [][]string
	config config.Config
}

func NewMatrixHandler(matrix [][]string, config config.Config) *MatrixHandler {
	return &MatrixHandler{matrix: matrix, config: config}
}

func (me MatrixHandler) GetReportWithData() (matrix [][]string, err error) {
	return me.matrix, nil
}
