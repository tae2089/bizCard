package domain

type Result struct {
	Message    string      `form:"message" json:"message"`
	StatusCode int         `form:"statusCode" json:"statusCode"`
	Data       interface{} `form:"data" json:"data"`
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
