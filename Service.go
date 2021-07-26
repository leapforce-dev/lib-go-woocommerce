package woocommerce

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

const (
	apiName          string = "WooCommerce"
	apiPath          string = "wp-json/wc/v2"
	dateFormat       string = "2006-01-02T15:04:05"
	totalPagesHeader string = "X-WP-TotalPages"
)

// type
//
type Service struct {
	host        string
	token       string
	httpService *go_http.Service
}

type ServiceConfig struct {
	Host           string
	ConsumerKey    string
	ConsumerSecret string
}

func NewService(config *ServiceConfig) (*Service, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	if config.Host == "" {
		return nil, errortools.ErrorMessage("Host not provided")
	}

	if config.ConsumerKey == "" {
		return nil, errortools.ErrorMessage("ConsumerKey not provided")
	}

	if config.ConsumerSecret == "" {
		return nil, errortools.ErrorMessage("ConsumerSecret not provided")
	}

	httpService, e := go_http.NewService(&go_http.ServiceConfig{})
	if e != nil {
		return nil, e
	}

	return &Service{
		host:        config.Host,
		token:       base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", config.ConsumerKey, config.ConsumerSecret))),
		httpService: httpService,
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
	return fmt.Sprintf("%s/%s/%s", service.host, apiPath, path)
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

func (service *Service) APIName() string {
	return apiName
}

func (service *Service) APIKey() string {
	return service.token
}

func (service *Service) APICallCount() int64 {
	return service.httpService.RequestCount()
}

func (service *Service) APIReset() {
	service.httpService.ResetRequestCount()
}

func UIntArrayToString(unints []uint) string {
	ids := []string{}
	for _, include := range unints {
		ids = append(ids, fmt.Sprintf("%v", include))
	}

	return strings.Join(ids, ",")
}

func TotalPages(response *http.Response) (int, *errortools.Error) {
	if response == nil {
		return 0, nil
	}

	totalPages, err := strconv.Atoi(response.Header.Get(totalPagesHeader))
	if err != nil {
		return 0, errortools.ErrorMessage(fmt.Sprintf("Error while retrieving %s header (%s)", totalPagesHeader, err.Error()))
	}

	return totalPages, nil
}
