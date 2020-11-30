package utility

// ResponseError struct
type ResponseError struct {
	Code         string `json:"code"`
	Deescription string `json:"description"`
}

// Responses struct
type Responses struct {
	Status      string          `json:"status"`
	Description string          `json:"description"`
	Data        interface{}     `json:"data,omitempty"`
	Errors      []ResponseError `json:"error"`
}

// ErrorResponses schema
func (r *Responses) ErrorResponses(description string, data []ResponseError) {
	r.Status = "001"
	r.Description = description
	r.Errors = data
}

// SuccessResponse schema
func (r *Responses) SuccessResponse(status string, description string, data interface{}) {
	r.Status = status
	r.Description = description
	r.Data = data
}
