package utils

type ApiMessage struct {
  Message string 	`json:"message"`
}

type ApiError struct {
  Message string 	`json:"message"`
  Code int 		`json:"code"`
  Field string 	`json:"field"`
}

type ApiErrorsWithMessage struct {
  ApiMessage
  Errors []ApiError 	`json:"errors"`
}
