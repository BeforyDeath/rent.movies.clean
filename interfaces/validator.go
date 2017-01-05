package interfaces

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func NewValidator(params map[string]interface{}, obj interface{}) error {
	v := reflect.ValueOf(obj).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		vField := v.Field(i)
		tField := t.Field(i)

		tag := tField.Tag.Get("validate")
		if tag == "" {
			continue
		}

		if tag == "required" {
			res, err := find(tField.Name, vField.Kind(), params)
			if err != nil {
				return fmt.Errorf("%v (required)", err)
			}
			err = write(vField, res)
			if err != nil {
				return fmt.Errorf("%v: %v (required)", tField.Name, err)
			}
		}
		if tag == "neglect" {
			res, err := find(tField.Name, vField.Kind(), params)
			if err == nil {
				err := write(vField, res)
				if err != nil && !strings.HasPrefix(err.Error(), "Invalid type field") {
					return fmt.Errorf("%v: %v", tField.Name, err)
				}
			}
		}
	}
	return nil
}

func find(nField string, tField reflect.Kind, r map[string]interface{}) (interface{}, error) {
	name := strings.ToLower(nField)
	if nObjField, ok := r[name]; ok {

		tObjField := reflect.ValueOf(nObjField).Kind()
		if _, ok := nObjField.(json.Number); ok && tField == reflect.Int64 {
			tField = reflect.Int64
			tObjField = tField
		}

		if tObjField == tField {
			return nObjField, nil
		}

		return nil, fmt.Errorf("Invalid type field '%v', expected %v", name, tField)
	}
	return nil, fmt.Errorf("Field '%v' not found", name)
}

func write(f reflect.Value, r interface{}) error {
	switch f.Kind() {
	case reflect.String:
		f.SetString(r.(string))
	case reflect.Int64:
		n, err := r.(json.Number).Int64()
		if err != nil {
			return fmt.Errorf("Invalid type field '%s', expected int", r.(json.Number))
		}
		f.SetInt(n)
	case reflect.Bool:
		f.SetBool(r.(bool))
	default:
		return fmt.Errorf("Unsupported type '%s'", f.Kind())
	}
	return nil
}
