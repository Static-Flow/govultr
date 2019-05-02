package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestApplicationServiceHandler_GetList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/app/list", func(w http.ResponseWriter, r *http.Request) {

		response := `
		{
			"1": {
				"APPID": "1",
				"name": "LEMP",
				"short_name": "lemp",
				"deploy_name": "LEMP on CentOS 6 x64",
				"surcharge": 0
			}
		}
		`

		fmt.Fprint(w, response)
	})

	apps, err := client.Application.GetList(ctx)
	if err != nil {
		t.Errorf("Application.GetList returned error: %v", err)
	}

	expected := []Application{
		{
			AppID:      "1",
			Name:       "LEMP",
			ShortName:  "lemp",
			DeployName: "LEMP on CentOS 6 x64",
			Surcharge:  0,
		},
	}

	if !reflect.DeepEqual(apps, expected) {
		t.Errorf("Application.GetList returned %+v, expected %+v", apps, expected)
	}
}

func TestApplicationServiceHandler_GetListEmpty(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/app/list", func(w http.ResponseWriter, r *http.Request) {

		response := `[]`

		fmt.Fprint(w, response)
	})

	apps, err := client.Application.GetList(ctx)
	if err != nil {
		t.Errorf("Application.GetList returned error: %v", err)
	}

	var expected []Application

	if !reflect.DeepEqual(apps, expected) {
		t.Errorf("Application.GetList returned %+v, expected %+v", apps, expected)
	}
}
