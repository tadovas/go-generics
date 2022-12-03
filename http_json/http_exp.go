package http_json

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Validator interface {
	Validate() error
}

type Response interface {
	ErrorCode() int
}

type HttpResponse struct {
	Code int
}

func (hr HttpResponse) ErrorCode() int {
	return hr.Code
}

type SomeStruct struct {
	SomeData    string
	SomeCounter int
}

func (ss SomeStruct) Validate() error {
	return errors.New("validation failed")
}

func CallMe[I Validator](input I) error {
	if err := input.Validate(); err != nil {
		return fmt.Errorf("validation: %v", err)
	}
	return nil
}

type JsonHandler[I Validator, R Response] interface {
	HandleJson(I) R
}

type JsonHandlerFunc func(input SomeStruct) HttpResponse

func (f JsonHandlerFunc) HandleJson(input SomeStruct) HttpResponse {
	return f(input)
}

func HandleJson[I Validator, R Response](h JsonHandler[I, R]) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler hit!")
		input := new(I)
		output := h.HandleJson(*input)
		jsonOutput, err := json.Marshal(output)
		if err != nil {
			fmt.Println("Json output error", err)
		}
		if _, err := writer.Write(jsonOutput); err != nil {
			fmt.Println("Writer error")
		}
		writer.WriteHeader(output.ErrorCode())
	}
}
