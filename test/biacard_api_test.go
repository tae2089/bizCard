package test

import (
	"bizCard/router"
	"github.com/gavv/httpexpect"
	"net/http"
	"testing"
)

func TestRegisterBizCard(t *testing.T) {
	handler := router.SetupRouter()
	// Create httpexpect instance
	e := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(handler),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})
	data := map[string]interface{}{
		"name":        "taebin",
		"email":       "tae2089",
		"phoneNumber": "010-xxxx-xxxx",
		"age":         25,
	}

	e.POST("/register").
		WithHeader("Content-Type", "application/json").
		WithJSON(data).Expect().
		JSON().
		Object().
		ContainsKey("name").
		ValueEqual("name", "taebin")
}
