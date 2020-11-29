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


// func Validator(data io.ReadCloser, request interface{}, schema validation.MapRule) (interface{}, error) {
// 	fmt.Printf("Just recieved: %v\n", request)
// 	err := json.NewDecoder(data).Decode(request)

// 	if err != nil {
// 		return nil, fmt.Errorf("Invalid JSON body")
// 	}

// 	verr := validation.Validate(request, schema)

// 	if verr != nil {
// 		return nil, verr
// 	}

// 	return request, nil
// }
