package engine

type RequestData struct {
	matrix [][]string
}

func NewRequestData(matrix [][]string) RequestData {
	return RequestData{matrix: matrix}
}

func (this RequestData) getReportDef() int {
	return 0
}
