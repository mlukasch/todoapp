package validators

import (
	"errors"
	"strings"
)

type Field struct {
	value      string
	validators []FieldValidator
}

func NewField(value string) *Field {
	return &Field{value, nil}
}

type FieldValidator interface {
	Validate(string) (bool, error)
}

func (this *Field) Add(validator FieldValidator) *Field {
	this.validators = append(this.validators, validator)
	return this
}

func (this *Field) Execute() (bool, []error) {
	result := make([]error, 0)
	for _, validator := range this.validators {
		ok, error := validator.Validate(this.value)
		if !ok {
			result = append(result, error)
		}
	}
	ok := len(result) == 0
	return ok, result
}

// InputValidator-Interface:
// A Function-Interface, that also fulfills the Validator-Object Interface
type FieldValidatorFunc func(string) (bool, error)

func (v FieldValidatorFunc) Validate(input string) (bool, error) {
	return v(input)
}

// NotEmptyValidator is: InputValidator-Function + Validator-Object:
var NotEmptyValidator FieldValidatorFunc = func(input string) (bool, error) {
	if len(strings.Trim(input, " ")) == 0 {
		return false, errors.New("Input is empty")
	}
	return true, nil
}

// Form Validator:
type FormValidator struct {
	fields map[string]*Field
}

func NewForm() *FormValidator {
	fields := make(map[string]*Field)
	return &FormValidator{fields}
}

func (f *FormValidator) AddField(fieldName, fieldValue string) *Field {
	field := NewField(fieldValue)
	f.fields[fieldName] = field
	return field
}

func (f *FormValidator) Execute() (bool, map[string][]error) {
	result := make(map[string][]error)
	for fieldName, field := range f.fields {
		_, errors := field.Execute()
		result[fieldName] = errors
	}
	for _, errors := range result {
		if len(errors) > 0 {
			return false, result
		}
	}
	return true, result
}
