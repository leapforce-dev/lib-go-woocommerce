package woocommerce

import (
	"encoding/json"
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
	w_types "github.com/leapforce-libraries/go_woocommerce/types"
	"net/http"
	"net/url"
)

// ProductVariation stores ProductVariation from Service
type ProductVariation struct {
	Id                *int64                       `json:"id,omitempty"`
	DateCreated       *w_types.DateTimeString      `json:"date_created,omitempty"`
	DateCreatedGmt    *w_types.DateTimeString      `json:"date_created_gmt,omitempty"`
	DateModified      *w_types.DateTimeString      `json:"date_modified,omitempty"`
	DateModifiedGmt   *w_types.DateTimeString      `json:"date_modified_gmt,omitempty"`
	Description       *string                      `json:"description,omitempty"`
	Permalink         *string                      `json:"permalink,omitempty"`
	Sku               *string                      `json:"sku,omitempty"`
	Price             *go_types.Float64String      `json:"price,omitempty"`
	RegularPrice      *go_types.Float64String      `json:"regular_price,omitempty"`
	SalePrice         *go_types.Float64String      `json:"sale_price,omitempty"`
	DateOnSaleFrom    *w_types.DateTimeString      `json:"date_on_sale_from,omitempty"`
	DateOnSaleFromGmt *w_types.DateTimeString      `json:"date_on_sale_from_gmt,omitempty"`
	DateOnSaleTo      *w_types.DateTimeString      `json:"date_on_sale_to,omitempty"`
	DateOnSaleToGmt   *w_types.DateTimeString      `json:"date_on_sale_to_gmt,omitempty"`
	OnSale            *bool                        `json:"on_sale,omitempty"`
	Visible           *bool                        `json:"visible,omitempty"`
	Purchasable       *bool                        `json:"purchasable,omitempty"`
	Virtual           *bool                        `json:"virtual,omitempty"`
	Downloadable      *bool                        `json:"downloadable,omitempty"`
	Downloads         *json.RawMessage             `json:"downloads,omitempty"`
	DownloadLimit     *int64                       `json:"download_limit,omitempty"`
	DownloadExpiry    *int64                       `json:"download_expiry,omitempty"`
	TaxStatus         *string                      `json:"tax_status,omitempty"`
	TaxClass          *string                      `json:"tax_class,omitempty"`
	ManageStock       json.RawMessage              `json:"manage_stock,omitempty"`
	StockQuantity     *go_types.Int64String        `json:"stock_quantity,omitempty"`
	InStock           *bool                        `json:"in_stock,omitempty"`
	Backorders        *string                      `json:"backorders,omitempty"`
	BackordersAllowed *bool                        `json:"backorders_allowed,omitempty"`
	Backordered       *bool                        `json:"backordered,omitempty"`
	Weight            *go_types.Float64String      `json:"weight,omitempty"`
	Dimensions        *ProductDimensions           `json:"dimensions,omitempty"`
	ShippingClass     *string                      `json:"shipping_class,omitempty"`
	ShippingClassId   *int64                       `json:"shipping_class_id,omitempty"`
	Image             *ProductImage                `json:"image,omitempty"`
	Attributes        *[]ProductVariationAttribute `json:"attributes,omitempty"`
	MenuOrder         *int64                       `json:"menu_order,omitempty"`
	MetaData          *[]ProductMetaData           `json:"meta_data,omitempty"`
}

type ProductVariationAttribute struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Option string `json:"option"`
}

// GetProductVariations returns all productVariations
func (service *Service) GetProductVariations(productId int64) (*[]ProductVariation, *errortools.Error) {
	values := url.Values{}

	page := 1

	var productVariations []ProductVariation

	for {
		values.Set("per_page", fmt.Sprintf("%v", 100))
		values.Set("page", fmt.Sprintf("%v", page))

		var productVariations_ []ProductVariation

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("products/%v/variations?%s", productId, values.Encode())),
			ResponseModel: &productVariations_,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		if len(productVariations_) == 0 {
			break
		}

		productVariations = append(productVariations, productVariations_...)

		page++
	}

	return &productVariations, nil
}
