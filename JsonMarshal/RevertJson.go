package JsonMarshal

type StatusDetails struct {
	Name              string        `json:"name,omitempty"`
	Kind              string        `json:"kind,omitempty"`
	Causes            []StatusCause `json:"causes,omitempty"`
	RetryAfterSeconds int           `json:"retryAfterSeconds,omitempty"`
}

type StatusCause struct {
	Type    string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
	Field   string `json:"field,omitempty"`
}

//func NewNotFound(kind, name string) error {
//	return &StatusError{api.Status{
//		Status: api.StatusFailure,
//		Code:   http.StatusNotFound,
//		Reason: api.StatusReasonNotFound,
//		Details: &api.StatusDetails{
//			Kind: kind,
//			Name: name,
//		},
//		Message: fmt.Sprintf("%s %q not found", kind, name),
//	}}
//}
