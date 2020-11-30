package utility

import (
	"encoding/json"
	"fmt"
	"io"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Validator function
func Validator(data io.ReadCloser, request interface{}, fields ...*validation.FieldRules) error {
	fmt.Printf("Just recieved: %v\n", request)
	err := json.NewDecoder(data).Decode(request)

	if err != nil {
		return fmt.Errorf("Invalid JSON body")
	}

	verr := validation.ValidateStruct(request, fields...)

	if verr != nil {
		return verr
	}

	return nil
}