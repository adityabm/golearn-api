package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func JsonResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{message, code, status}

	response := Response{meta, data}

	return response
}

func FormatError(err error) []string {
	var errors []string

	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, err.Error())
	}

	return errors
}