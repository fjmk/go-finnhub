package client_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/llonchj/go-finnhub"
	"github.com/llonchj/go-finnhub/client"
)

func NewTestAPI() *client.API {
	testSrv := http.NewServeMux()
	testSrv.HandleFunc("/v1/valid", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `""`)
	})
	testSrv.HandleFunc("/v1/invalid", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "text/plain")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid API Key")
	})
	testSrv.HandleFunc("/v1/throttle", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "text/plain")
		w.WriteHeader(http.StatusTooManyRequests)
	})

	srv := httptest.NewServer(testSrv)
	return &client.API{
		Client:   http.DefaultClient,
		Endpoint: srv.URL,
		Key:      "token",
	}
}

func TestAPI_Call(t *testing.T) {
	// t.Parallel()

	for name, tt := range map[string]struct {
		Method string
		Path   string
		Params finnhub.URLParams

		Want interface{}
		Err  error
	}{
		"valid": {
			Path: "valid",
			Want: "",
		},
		"invalid token": {
			Path: "invalid",
			Err:  client.ErrUnauthorized,
		},
		"429 - too many requests": {
			Path: "throttle",
			Err:  client.ErrTooManyRequests,
		},
	} {
		t.Run(name, func(t *testing.T) {
			a := NewTestAPI()

			method := "GET"
			if tt.Method != "" {
				method = tt.Method
			}

			params := finnhub.URLParams{}
			if tt.Params != nil {
				params = tt.Params
			}

			var response interface{}
			err := a.Call(method, tt.Path, params, &response)

			if err != nil {
				if tt.Err == nil || err.Error() != tt.Err.Error() {
					t.Errorf("%v, wantErr %v", err, tt.Err)
				}
			} else if tt.Err != nil {
				t.Errorf("%v, wantErr %v", err, tt.Err)
			}
			if !reflect.DeepEqual(response, tt.Want) {
				t.Errorf("mismatch: got = %v, want %v", response, tt.Want)
			}
		})
	}
}
