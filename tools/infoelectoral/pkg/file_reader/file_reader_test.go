package file_reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadFile(t *testing.T) {

}

func Test_extractColumns(t *testing.T) {
	var tStruct testStruct

	columns, err := extractColumns(tStruct)
	assert.Nil(t, err, "Should not return error on valid columns")
	assert.ElementsMatch(t, columnsContent, columns)
}

func Test_calculateColumnsTotalLength(t *testing.T) {
	var columns []Column

	assert.Equal(t, 0, calculateColumnsTotalLength(columns))

	columns = []Column{
		{
			length: 2,
		},
		{
			length: 2,
		},
		{
			length: 4,
		},
	}
	assert.Equal(t, 8, calculateColumnsTotalLength(columns))
}

type testStruct struct {
	year    int    `position:"0" length:"4"`
	month   int    `position:"4" length:"2"`
	day     int    `position:"6" length:"2"`
	weekDay string `position:"8" length:"8"`
	happy   bool   `position:"16" length:"1"`
}

var columnsContent = []Column{
	{
		name:       "year",
		position:   0,
		length:     4,
		columnType: "int",
	},
	{
		name:       "month",
		position:   4,
		length:     2,
		columnType: "int",
	},
	{
		name:       "day",
		position:   6,
		length:     2,
		columnType: "int",
	},
	{
		name:       "weekDay",
		position:   8,
		length:     8,
		columnType: "string",
	},
	{
		name:       "happy",
		position:   16,
		length:     1,
		columnType: "bool",
	},
}
