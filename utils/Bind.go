package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"reflect"
	"strconv"
)

func BindMultipartForm(form *multipart.Form, target interface{}) error {
	v := reflect.ValueOf(target)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("target must be a pointer to a struct")
	}
	v = v.Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// Obtener el nombre del campo del tag 'form'
		formKey := fieldType.Tag.Get("form")
		if formKey == "" || formKey == "-" {
			continue
		}

		// Obtener el valor del formulario
		formValues, exists := form.Value[formKey]
		fileValues, fileExists := form.File[formKey]

		if !exists && !fileExists {
			continue
		}

		// Manejar archivos y slices de archivos
		if field.Type() == reflect.TypeOf((*multipart.FileHeader)(nil)) {
			if fileExists && len(fileValues) > 0 {
				field.Set(reflect.ValueOf(fileValues[0]))
			}
			continue
		}

		if field.Type() == reflect.TypeOf(([]*multipart.FileHeader)(nil)).Elem() {
			if fileExists {
				field.Set(reflect.ValueOf(fileValues))
			}
			continue
		}

		// Si no hay valores, continuar
		if len(formValues) == 0 {
			continue
		}

		value := formValues[0]

		// Parsear según el tipo de campo
		switch field.Kind() {
		case reflect.String:
			field.SetString(value)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intVal, err := strconv.ParseInt(value, 10, 64)
			if err == nil {
				field.SetInt(intVal)
			}
		case reflect.Float32, reflect.Float64:
			floatVal, err := strconv.ParseFloat(value, 64)
			if err == nil {
				field.SetFloat(floatVal)
			}
		case reflect.Bool:
			boolVal, err := strconv.ParseBool(value)
			if err == nil {
				field.SetBool(boolVal)
			}
		case reflect.Slice:
			// Manejar slices de tipos básicos
			if field.Type().Elem().Kind() == reflect.String {
				var strSlice []string
				if err := json.Unmarshal([]byte(value), &strSlice); err == nil {
					field.Set(reflect.ValueOf(strSlice))
				}
			}
		case reflect.Struct:
			// Manejar structs parseables desde JSON
			if field.CanAddr() {
				if err := json.Unmarshal([]byte(value), field.Addr().Interface()); err != nil {
					log.Printf("Error parsing struct: %v", err)
				}
			}
		}
	}

	return nil
}
