package woocommerce

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
	"net/url"
)

// ProductAttributeDef stores ProductAttributeDef from Service
//
type ProductAttributeDef struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Type        string `json:"type"`
	OrderBy     string `json:"order_by"`
	HasArchives bool   `json:"has_archives"`
}

type GetProductAttributeDefsContext string

const (
	GetProductAttributeDefsContextView GetProductAttributeDefsContext = "view"
	GetProductAttributeDefsContextEdit GetProductAttributeDefsContext = "edit"
)

type GetProductAttributeDefsConfig struct {
	Context *GetProductAttributeDefsContext
}

// GetProductAttributeDefs returns all productAttributeDefs
//
func (service *Service) GetProductAttributeDefs(config *GetProductAttributeDefsConfig) (*[]ProductAttributeDef, *errortools.Error) {
	values := url.Values{}
	endpoint := "products/attributes"

	if config != nil {
		if config.Context != nil {
			values.Set("context", string(*config.Context))
		}
	}

	productAttributeDefs := []ProductAttributeDef{}

	path := fmt.Sprintf("%s?%s", endpoint, values.Encode())

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(path),
		ResponseModel: &productAttributeDefs,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &productAttributeDefs, nil
}

// UpdateProductAttributeDef updates all productAttributeDefs
//
func (service *Service) UpdateProductAttributeDef(productAttributeDef *ProductAttributeDef) (*ProductAttributeDef, *errortools.Error) {
	if productAttributeDef == nil {
		return nil, errortools.ErrorMessage("ProductAttributeDef is a nil pointer")
	}

	updatedProductAttributeDef := ProductAttributeDef{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPut,
		Url:           service.url(fmt.Sprintf("products/attributes/%v", productAttributeDef.Id)),
		BodyModel:     productAttributeDef,
		ResponseModel: &updatedProductAttributeDef,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &updatedProductAttributeDef, nil
}
