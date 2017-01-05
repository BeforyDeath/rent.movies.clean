package interfaces

import (
	"encoding/json"
	"testing"
)

type valStructure01 struct {
	Param1 string `validate:"required"`
	Param2 bool   `validate:"required"`
	Param3 string
	Param4 int64 `validate:"neglect"`
}

type valStructure02 struct {
	Param1 int `validate:"neglect"`
	Param2 int `validate:"required"`
}

type valStructure03 struct {
	Param1 int64 `validate:"neglect"`
	Param2 int64 `validate:"required"`
}

func TestNewValidator(t *testing.T) {
	var jsonParams = []map[string]interface{}{
		{"param1": 1},
		{"param1": "res01"},
		{"param1": "res01", "param2": "res02"},
		{"param1": "res01", "param2": false},
		{"param1": "res01", "param2": true, "param4": json.Number("15")},
		{"param1": "res01", "param2": true, "param3": "res02"},

		{"param1": 1},
		{"param2": 1},

		{"param2": json.Number("15.2")},
		{"param1": json.Number("15.2"), "param2": json.Number("15")},
	}

	val01 := new(valStructure01)
	err := NewValidator(jsonParams[0], val01)
	if err.Error() != "Invalid type field 'param1', expected string (required)" {
		t.Error(err)
	}

	val01 = new(valStructure01)
	err = NewValidator(jsonParams[1], val01)
	if err.Error() != "Field 'param2' not found (required)" {
		t.Error(err)
	}

	val01 = new(valStructure01)
	err = NewValidator(jsonParams[2], val01)
	if err.Error() != "Invalid type field 'param2', expected bool (required)" {
		t.Error(err)
	}

	val01 = new(valStructure01)
	err = NewValidator(jsonParams[3], val01)
	if val01.Param2 {
		t.Error("Expected false, got", val01.Param2, err)
	}

	val01 = new(valStructure01)
	err = NewValidator(jsonParams[4], val01)
	if val01.Param4 != 15 {
		t.Error("Expected 15, got", val01.Param4, err)
	}

	val01 = new(valStructure01)
	err = NewValidator(jsonParams[4], val01)
	if err != nil {
		t.Error(err)
	}
	if val01.Param1 != "res01" || !val01.Param2 || val01.Param3 != "" || val01.Param4 != 15 {
		t.Error("Incorrect expected result")
	}

	val01 = new(valStructure01)
	err = NewValidator(jsonParams[5], val01)
	if val01.Param3 != "" {
		t.Error(err)
	}

	val02 := new(valStructure02)
	err = NewValidator(jsonParams[6], val02)
	if err.Error() != "Param1: Unsupported type 'int'" {
		t.Error(err)
	}

	val02 = new(valStructure02)
	err = NewValidator(jsonParams[7], val02)
	if err.Error() != "Param2: Unsupported type 'int' (required)" {
		t.Error(err)
	}

	val03 := new(valStructure03)
	err = NewValidator(jsonParams[8], val03)
	if err.Error() != "Param2: Invalid type field '15.2', expected int (required)" {
		t.Error(err)
	}
	val03 = new(valStructure03)
	err = NewValidator(jsonParams[9], val03)
	if err != nil {
		t.Error(err)
	}
}
