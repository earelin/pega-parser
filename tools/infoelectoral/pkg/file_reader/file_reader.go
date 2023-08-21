package file_reader

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type Column struct {
	name       string
	position   int
	length     int
	columnType string
}

type FileReader[T any] struct {
	file     *os.File
	lineSize int
	columns  []Column
}

func (fr FileReader[T]) Read() T {
	var line T
	return line
}

func NewFileReader[T any](file *os.File) (FileReader[T], error) {
	var structType T

	columns, err := extractColumns[T](structType)
	if err != nil {
		return FileReader[T]{}, err
	}

	return FileReader[T]{
		file:    file,
		columns: columns,
	}, nil
}

func extractColumns[T any](structType T) ([]Column, error) {
	var columns []Column

	t := reflect.TypeOf(structType)
	if t.Kind().String() != "struct" {
		errorMessage := fmt.Sprintf("not valid type %s", t.Kind().String())
		return columns, errors.New(errorMessage)
	}

	e := reflect.ValueOf(&structType).Elem()
	for i := 0; i < e.NumField(); i++ {
		f := e.Type().Field(i)

		tags := f.Tag
		position, err := strconv.Atoi(tags.Get("position"))
		if err != nil {
			return []Column{}, err
		}

		length, err := strconv.Atoi(tags.Get("length"))
		if err != nil {
			return []Column{}, err
		}

		columns = append(columns, Column{
			columnType: f.Type.Name(),
			name:       f.Name,
			position:   position,
			length:     length,
		})
	}

	return columns, nil
}

func calculateColumnsTotalLength(columns []Column) int {
	var totalLength int

	for _, column := range columns {
		totalLength += column.length
	}

	return totalLength
}
