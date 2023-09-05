package file_reader

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"reflect"
	"sort"
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
	var data []byte
	var b = make([]byte, 1)

	var err error
	for {
		_, err = fr.file.Read(b)
		if err != nil && err != io.EOF {
			return structuredData, err
		}
		if b[0] == 10 || err == io.EOF {
			break
		}
		data = append(data, b[0])
	}

	if len(data) == 0 {
		return structuredData, io.EOF
	}

	var merr error
	structuredData, merr = unMarshaling[T](data, fr.columns)
	if merr != nil {
		fmt.Printf("Error reading data: %s\n", merr)
	}

	return structuredData, nil
}

func (fr FileReader[T]) Close() {
	fr.file.Close()
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
			rawValueString := strings.TrimSpace(string(rawValue))
			number, err := strconv.Atoi(rawValueString)
			if err != nil {
				return structuredData, fmt.Errorf("error converting %s to int in column %s: %w", rawValue, column.name, err)
			}
			field.SetInt(int64(number))
		case "bool":
			field.SetBool(string(rawValue) == "1")
		case "string":
			field.SetString(strings.TrimSpace(isoToUtf8(rawValue)))
		default:
			return structuredData, fmt.Errorf("could not parse type %s in column %s", rawValue, column.name)
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

	sortColumns(columns)

	return columns, nil
}

func sortColumns(columns []Column) {
	sort.Slice(columns, func(i int, j int) bool {
		return columns[i].position < columns[j].position
	})
}

func calculateLineLength(columns []Column) int {
	var totalLength int

	for _, column := range columns {
		totalLength += column.length
	}

	return totalLength
}
