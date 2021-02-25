package woocommerce

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

const (
	APIPath    string = "wp-json/wc/v2"
	DateFormat string = "2006-01-02T15:04:05Z"
)

// type
//
type Service struct {
	domain      string
	token       string
	httpService *go_http.Service
}

type ServiceConfig struct {
	Domain         string
	ConsumerKey    string
	ConsumerSecret string
}

func NewService(config ServiceConfig) (*Service, *errortools.Error) {
	if config.Domain == "" {
		return nil, errortools.ErrorMessage("Domain not provided")
	}

	if config.ConsumerKey == "" {
		return nil, errortools.ErrorMessage("ConsumerKey not provided")
	}

	if config.ConsumerSecret == "" {
		return nil, errortools.ErrorMessage("ConsumerSecret not provided")
	}

	httpServiceConfig := go_http.ServiceConfig{}

	return &Service{
		domain:      config.Domain,
		token:       base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", config.ConsumerKey, config.ConsumerSecret))),
		httpService: go_http.NewService(httpServiceConfig),
	}, nil
}

func (service *Service) httpRequest(httpMethod string, requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	// add authentication header
	header := http.Header{}
	header.Set("Authorization", fmt.Sprintf("Basic %s", service.token))
	(*requestConfig).NonDefaultHeaders = &header

	// add error model
	errorResponse := ErrorResponse{}
	(*requestConfig).ErrorModel = &errorResponse

	request, response, e := service.httpService.HTTPRequest(httpMethod, requestConfig)
	if errorResponse.Message != "" {
		e.SetMessage(errorResponse.Message)
	}

	return request, response, e
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s/%s", service.domain, APIPath, path)
}

func (service *Service) get(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodGet, requestConfig)
}

func (service *Service) post(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodPost, requestConfig)
}

func (service *Service) put(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodPut, requestConfig)
}

func (service *Service) delete(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodDelete, requestConfig)
}

func UIntArrayToString(unints []uint) string {
	ids := []string{}
	for _, include := range unints {
		ids = append(ids, fmt.Sprintf("%v", include))
	}

	return strings.Join(ids, ",")
}
