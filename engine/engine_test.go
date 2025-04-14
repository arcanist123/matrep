package engine

import (
	"reflect"
	"testing"
)

func TestReportDef_GetReportDefinition(t *testing.T) {
	tests := []struct {
		name         string
		matrix       [][]string
		wantResult   *resolved_report
		wantErrCount int
	}{
		{
			name:         "empty matrix",
			matrix:       [][]string{},
			wantResult:   nil,
			wantErrCount: 0,
		},
		{
			name: "valid matrix",
			matrix: [][]string{
				{"key1=value1", "key2=value2"},
				{"key3=value3", "key4=value4"},
			},
			wantResult:   nil,
			wantErrCount: 0,
		},
		{
			name: "invalid matrix - invalid format",
			matrix: [][]string{
				{"key1=value1", "key2value2"},
				{"key3=value3", "key4=value4"},
			},
			wantResult:   nil,
			wantErrCount: 1,
		},
		{
			name: "invalid matrix - two equals",
			matrix: [][]string{
				{"key1=value1", "key2=value=2"},
				{"key3=value3", "key4=value4"},
			},
			wantResult:   nil,
			wantErrCount: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := NewReportDef(tt.matrix)
			gotResult, gotErrs := me.GetReportDefinition()
			if tt.wantErrCount != len(gotErrs) {
				t.Errorf("GetReportDefinition() error count = %v, wantErr count %v", len(gotErrs), tt.wantErrCount)
				if gotErrs != nil {
					for _, err := range gotErrs {
						t.Errorf("%s", err.Error())
					}
				}
				return
			}

			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("GetReportDefinition() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestReportDef_normaliseCellContents(t *testing.T) {
	tests := []struct {
		name         string
		matrix       [][]string
		wantResult   [][]string
		wantErrCount int
	}{
		{
			name:         "empty matrix",
			matrix:       [][]string{},
			wantResult:   [][]string{},
			wantErrCount: 0,
		},
		{
			name: "valid matrix",
			matrix: [][]string{
				{"key1=value1", "key2=value2"},
				{"key3=value3", "key4=value4"},
			},
			wantResult: [][]string{
				{"key1=value1", "key2=value2"},
				{"key3=value3", "key4=value4"},
			},
			wantErrCount: 0,
		},
		{
			name: "invalid matrix - invalid format",
			matrix: [][]string{
				{"key1=value1", "key2value2"},
				{"key3=value3", "key4=value4"},
			},
			wantResult: [][]string{
				{"key1=value1", "key2value2"},
				{"key3=value3", "key4=value4"},
			},
			wantErrCount: 1,
		},
		{
			name: "invalid matrix - two equals",
			matrix: [][]string{
				{"key1=value1", "key2=value=2"},
				{"key3=value3", "key4=value4"},
			},
			wantResult: [][]string{
				{"key1=value1", "key2=value=2"},
				{"key3=value3", "key4=value4"},
			},
			wantErrCount: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := NewReportDef(tt.matrix)
			gotResult, gotErrs := me.normaliseCellContents(tt.matrix)
			if tt.wantErrCount != len(gotErrs) {
				t.Errorf("normaliseCellContents() error count = %v, wantErr count %v", len(gotErrs), tt.wantErrCount)
				if gotErrs != nil {
					for _, err := range gotErrs {
						t.Errorf("%s", err.Error())
					}
				}
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("normaliseCellContents() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
func TestReportDef_get_page(t *testing.T) {
	tests := []struct {
		name       string
		matrix     [][]string
		wantResult report_page
		wantErr    bool
		wantErrMsg string
	}{
		{
			name:       "Empty Report",
			matrix:     [][]string{{}, {}, {}},
			wantResult: report_page{},
			wantErr:    true,
			wantErrMsg: "report is empty",
		},
		{
			name: "Single Empty Row",
			matrix: [][]string{
				{"key1=val1", "key2=val2"},
				{},
				{"key3=val3", "key4=val4"},
			},
			wantResult: report_page{
				tuple:      []string{},
				is_present: false,
				position:   csv_coordinate{row: 1, col: 0},
			},
			wantErr: false,
		},
		{
			name: "Multiple Empty Rows",
			matrix: [][]string{
				{"key1=val1", "key2=val2"},
				{},
				{},
			},
			wantResult: report_page{
				tuple:      []string{},
				is_present: false,
				position:   csv_coordinate{row: 1, col: 0},
			},
			wantErr: false,
		},
		{
			name: "Page Row",
			matrix: [][]string{
				{"key1=val1", "key2=val2"},
				{"A=1", "B=2"},
				{"key3=val3", "key4=val4"},
			},
			wantResult: report_page{
				tuple:      []string{"A=1", "B=2"},
				is_present: true,
				position:   csv_coordinate{row: 1, col: 0},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := NewReportDef(tt.matrix)
			gotResult, gotErr := me.get_page(tt.matrix)

			if (gotErr != nil) != tt.wantErr {
				t.Errorf("get_page() error = %v, wantErr %v", gotErr, tt.wantErr)
				return
			}
			if tt.wantErr {
				if gotErr.Error() != tt.wantErrMsg {
					t.Errorf("get_page() got error message '%v', want '%v'", gotErr.Error(), tt.wantErrMsg)
				}
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("get_page() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestReportDef_isRowEmpty(t *testing.T) {
	tests := []struct {
		name string
		row  []string
		want bool
	}{
		{
			name: "Empty row",
			row:  []string{"", "", ""},
			want: true,
		},
		{
			name: "Non-empty row",
			row:  []string{"a", "b", "c"},
			want: false,
		},
		{
			name: "Mixed row",
			row:  []string{"", "b", ""},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := NewReportDef([][]string{})
			if got := me.isRowEmpty(tt.row); got != tt.want {
				t.Errorf("isRowEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReportDef_resolve_row_to_page(t *testing.T) {
	tests := []struct {
		name           string
		row            []string
		wantResult     []string
		wantIsResolved bool
	}{
		{
			name:           "Empty row",
			row:            []string{"", "", ""},
			wantResult:     []string{"", "", ""},
			wantIsResolved: true,
		},
		{
			name:           "Single Key Value Pairs",
			row:            []string{"a=1", "b=2", "c=3"},
			wantResult:     []string{"a=1", "b=2", "c=3"},
			wantIsResolved: true,
		},
		{
			name:           "Duplicate Key",
			row:            []string{"a=1", "b=2", "a=3"},
			wantResult:     []string{},
			wantIsResolved: false,
		},
		{
			name:           "Mixed row",
			row:            []string{"a=1", "", "a=3"},
			wantResult:     []string{},
			wantIsResolved: false,
		},
		{
			name:           "Mixed row with key value pairs and Empty",
			row:            []string{"a=1", "", "c=3"},
			wantResult:     []string{"a=1", "", "c=3"},
			wantIsResolved: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := NewReportDef([][]string{})
			gotResult, gotIsResolved := me.resolve_row_to_page(tt.row)
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("resolve_row_to_page() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if gotIsResolved != tt.wantIsResolved {
				t.Errorf("resolve_row_to_page() gotIsResolved = %v, want %v", gotIsResolved, tt.wantIsResolved)
			}
		})
	}
}
