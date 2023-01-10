package woocommerce

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	w_types "github.com/leapforce-libraries/go_woocommerce/types"
	"net/http"
	"net/url"
)

type ProductBrand struct {
	Id          int64              `json:"id,omitempty"`
	Name        string             `json:"name,omitempty"`
	Slug        string             `json:"slug,omitempty"`
	Parent      int64              `json:"parent,omitempty"`
	Description string             `json:"description,omitempty"`
	Image       *ProductBrandImage `json:"image,omitempty"`
	MenuOrder   int64              `json:"menu_order,omitempty"`
	Count       int64              `json:"count,omitempty"`
}

type ProductBrandImage struct {
	Id              int64                  `json:"id"`
	DateCreated     w_types.DateTimeString `json:"date_created"`
	DateCreatedGmt  w_types.DateTimeString `json:"date_created_gmt"`
	DateModified    w_types.DateTimeString `json:"date_modified"`
	DateModifiedGmt w_types.DateTimeString `json:"date_modified_gmt"`
	Src             string                 `json:"src"`
	Name            string                 `json:"name"`
	Alt             string                 `json:"alt"`
}

// GetProductBrands returns all productBrands
//
func (service *Service) GetProductBrands() (*[]ProductBrand, *errortools.Error) {
	var page int64 = 1
	var perPage int64 = 100

	values := url.Values{}
	values.Set("per_page", fmt.Sprintf("%v", perPage))
	endpoint := "products/brands"

	var productBrands []ProductBrand

	for {
		values.Set("page", fmt.Sprintf("%v", page))

		var productBrands_ []ProductBrand

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("%s?%s", endpoint, values.Encode())),
			ResponseModel: &productBrands_,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		if len(productBrands_) == 0 {
			break
		}

		productBrands = append(productBrands, productBrands_...)
		page++
	}

	return &productBrands, nil
}

// CreateProductBrand creates a productBrand
//
func (service *Service) CreateProductBrand(productBrand *ProductBrand) (*ProductBrand, *errortools.Error) {
	if productBrand == nil {
		return nil, errortools.ErrorMessage("ProductBrand is a nil pointer")
	}

	createdProductBrand := ProductBrand{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           service.url("products/brands"),
		BodyModel:     productBrand,
		ResponseModel: &createdProductBrand,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &createdProductBrand, nil
}

// DeleteProductBrand deletes a productBrand
//
func (service *Service) DeleteProductBrand(id int64) *errortools.Error {
	requestConfig := go_http.RequestConfig{
		Method: http.MethodDelete,
		Url:    service.url(fmt.Sprintf("products/brands/%v?force=true", id)),
	}

	_, _, e := service.httpRequest(&requestConfig)
	return e
}
