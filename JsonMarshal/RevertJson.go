package JsonMarshal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type StatusDetails struct {
	Name              string        `json:"name,omitempty"`
	Kind              string        `json:"kind,omitempty"`
	Causes            []StatusCause `json:"causes,omitempty"`
	RetryAfterSeconds int           `json:"retryAfterSeconds,omitempty"`
}

type StatusCause struct {
	Typea   string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
	Field   string `json:"field,omitempty"`
}

type Causes struct {
	Reasons []StatusCause
}

func (a *StatusCause) String() string {
	return "{Type:" + a.Typea + ",Message:" + a.Message + ",Field:" + a.Field + "}"
}

func Main() {
	filePath := "/home/milliant/miplace/src/the-way-to-go/JsonMarshal/sample.json"
	file, err := os.Open(filePath)
	if err != nil {
		panic(err.Error())
	}
	content, errRead := ioutil.ReadAll(file)
	if errRead != nil {
		panic(errRead.Error())
	}
	fmt.Println(string(content))
	var causes Causes
	errJson := json.Unmarshal(content, &causes)
	if errJson != nil {
		panic(errJson.Error())
	} else {
		for _, value := range causes.Reasons {
			fmt.Printf(value.String())
		}
	}

	jsonInstance := &StatusDetails{
		Kind: "kind",
		Name: "name",
		Causes: []StatusCause{{
			Typea: "duplicate",
			Field: "",
		}, {
			Typea: "duplicate",
			Field: "hah",
		}},
	}
	j, err := json.Marshal(jsonInstance)
	if err != nil {
		fmt.Print("json convert fail!")
	} else {
		fmt.Print(string(j)) //j 是[]byte类型，可以直接转换成string
	}

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
