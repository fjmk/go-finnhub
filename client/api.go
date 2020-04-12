package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/llonchj/go-finnhub"
)

const (
	// APIEndpoint the api url
	APIEndpoint = "https://finnhub.io/api"

	// APIVersion the api version
	APIVersion = "v1"

	// UserAgentFmt the user agent for the client
	UserAgentFmt = "go-finnhub-%v"
)

var (
	// ErrUnauthorized if the token is invalid
	ErrUnauthorized = ErrServer("Invalid API Key")

	// ErrTooManyRequests if the api returns that you have made too many requests
	ErrTooManyRequests = ErrServer("you are over the request limit - you may of not entered a valid token")
)

//ErrServer is a server error
type ErrServer string

//Error implements errors.Error
func (e ErrServer) Error() string {
	return string(e)
}

// API is the data structure for holding the API details
type API struct {
	*http.Client

	Key           string
	ClientVersion string
	UserAgent     string

	Endpoint string
}

// NewAPI returns a new api client
func NewAPI(key string, clientVersion string) *API {
	return &API{
		Client:        http.DefaultClient,
		Key:           key,
		ClientVersion: clientVersion,
		Endpoint:      APIEndpoint,
		UserAgent:     fmt.Sprintf(UserAgentFmt, clientVersion),
	}
}

// Get requests a get from the api
func (a *API) Get(path string, params finnhub.URLParams, response interface{}) error {
	return a.Call(http.MethodGet, path, params, response)
}

// Call calls the api using a supplied method
func (a *API) Call(method string, path string, params finnhub.URLParams, response interface{}) error {
	q := url.Values{}
	if params != nil {
		if _, found := params[finnhub.ParamToken]; !found {
			params[finnhub.ParamToken] = a.Key
		}
	}
	for k, v := range params {
		q.Add(k, v)
	}

	endpoint := fmt.Sprintf("%v/%v/%v", a.Endpoint, APIVersion, path)
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return err
	}
	req.URL.RawQuery = q.Encode()
	resp, err := a.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		ct := strings.Split(resp.Header.Get("Content-type"), ";")[0]
		if ct == "application/json" {
			return json.Unmarshal(body, &response)
		}
		return errors.New(string(body))
	case http.StatusUnauthorized:
		return ErrUnauthorized
	case http.StatusTooManyRequests:
		return ErrTooManyRequests
	default:
		return ErrServer(body)
	}
}
