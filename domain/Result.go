package domain

type Result struct {
	Message    string
	StatusCode int
	Data       interface{}
}

func Success() *Result {
	return &Result{
		Message:    "success",
		StatusCode: 200,
	}
}

func Fail() *Result {
	return &Result{Message: "failure", StatusCode: 500}
}
