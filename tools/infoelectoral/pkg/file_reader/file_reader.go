package file_reader

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type Column struct {
	name       string
	position   int
	length     int
	columnType string
}

type FileReader[T any] struct {
	file     fs.File
	lineSize int
	columns  []Column
}

func (fr FileReader[T]) Read() (T, error) {
	var structuredData T
	var data = make([]byte, fr.lineSize+1)

	var err error
	_, err = fr.file.Read(data)
	if err != nil && err != io.EOF {
		return structuredData, err
	}

	var merr error
	structuredData, merr = unMarshaling[T](data, fr.columns)
	if merr != nil {
		log.Printf("Error reading data: %s", err)
	}

	return structuredData, err
}

func NewFileReader[T any](file fs.File) (FileReader[T], error) {
	var structType T

	columns, err := extractColumns[T](structType)
	if err != nil {
		return FileReader[T]{}, err
	}

	return FileReader[T]{
		file:     file,
		columns:  columns,
		lineSize: calculateLineLength(columns),
	}, nil
}

func unMarshaling[T any](data []byte, columns []Column) (T, error) {
	var structuredData T

	for _, column := range columns {
		field := reflect.ValueOf(&structuredData).Elem().FieldByName(column.name)
		rawValue := data[column.position : column.position+column.length]

		switch column.columnType {
		case "int":
			number, err := strconv.Atoi(string(rawValue))
			if err != nil {
				return structuredData, err
			}
			field.SetInt(int64(number))
		case "bool":
			field.SetBool(string(rawValue) == "1")
		case "string":
			field.SetString(strings.TrimSpace(isoToUtf8(rawValue)))
		default:
			errorMessage := fmt.Sprintf("could not parse type %s", column.columnType)
			return structuredData, errors.New(errorMessage)
		}
	}

	return structuredData, nil
}

func isoToUtf8(bytes []byte) string {
	buf := make([]rune, len(bytes))
	for i, b := range bytes {
		buf[i] = rune(b)
	}
	return string(buf)
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

func calculateLineLength(columns []Column) int {
	var totalLength int

	for _, column := range columns {
		totalLength += column.length
	}

	return totalLength
}
