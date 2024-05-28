package response

import "strings"

//response is use for static shape json return
type Response struct {
	Messgae string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Build(msg string, dta interface{}) Response {
	result := Response{
		Messgae: msg,
		Data:    dta,
	}
	return result
}

func Error(msg string, err string) Response {
	errors := strings.Split(err, "\n")
	result := Response{
		Messgae: msg,
		Errors:  errors,
	}
	return result
}
