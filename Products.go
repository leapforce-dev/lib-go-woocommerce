package woocommerce

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
	w_types "github.com/leapforce-libraries/go_woocommerce/types"
)

// Product stores Product from Service
type Product struct {
	Id                *int64                  `json:"id,omitempty"`
	Name              *string                 `json:"name,omitempty"`
	Slug              *string                 `json:"slug,omitempty"`
	Permalink         *string                 `json:"permalink,omitempty"`
	DateCreated       *w_types.DateTimeString `json:"date_created,omitempty"`
	DateCreatedGmt    *w_types.DateTimeString `json:"date_created_gmt,omitempty"`
	DateModified      *w_types.DateTimeString `json:"date_modified,omitempty"`
	DateModifiedGmt   *w_types.DateTimeString `json:"date_modified_gmt,omitempty"`
	Type              *string                 `json:"type,omitempty"`
	Status            *string                 `json:"status,omitempty"`
	Featured          *bool                   `json:"featured,omitempty"`
	CatalogVisibility *string                 `json:"catalog_visibility,omitempty"`
	Description       *string                 `json:"description,omitempty"`
	ShortDescription  *string                 `json:"short_description,omitempty"`
	Sku               *string                 `json:"sku,omitempty"`
	Price             *go_types.Float64String `json:"price,omitempty"`
	RegularPrice      *go_types.Float64String `json:"regular_price,omitempty"`
	SalePrice         *go_types.Float64String `json:"sale_price,omitempty"`
	DateOnSaleFrom    *w_types.DateTimeString `json:"date_on_sale_from,omitempty"`
	DateOnSaleFromGmt *w_types.DateTimeString `json:"date_on_sale_from_gmt,omitempty"`
	DateOnSaleTo      *w_types.DateTimeString `json:"date_on_sale_to,omitempty"`
	DateOnSaleToGmt   *w_types.DateTimeString `json:"date_on_sale_to_gmt,omitempty"`
	PriceHtml         *string                 `json:"price_html,omitempty"`
	OnSale            *bool                   `json:"on_sale,omitempty"`
	Purchasable       *bool                   `json:"purchasable,omitempty"`
	TotalSales        *go_types.Int64String   `json:"total_sales,omitempty"`
	Virtual           *bool                   `json:"virtual,omitempty"`
	Downloadable      *bool                   `json:"downloadable,omitempty"`
	Downloads         *json.RawMessage        `json:"downloads,omitempty"`
	DownloadLimit     *int64                  `json:"download_limit,omitempty"`
	DownloadExpiry    *int64                  `json:"download_expiry,omitempty"`
	ExternalUrl       *string                 `json:"external_url,omitempty"`
	ButtonText        *string                 `json:"button_text,omitempty"`
	TaxStatus         *string                 `json:"tax_status,omitempty"`
	TaxClass          *string                 `json:"tax_class,omitempty"`
	ManageStock       *bool                   `json:"manage_stock,omitempty"`
	StockQuantity     *go_types.Int64String   `json:"stock_quantity,omitempty"`
	StockStatus       *string                 `json:"stock_status,omitempty"`
	Backorders        *string                 `json:"backorders,omitempty"`
	BackordersAllowed *bool                   `json:"backorders_allowed,omitempty"`
	Backordered       *bool                   `json:"backordered,omitempty"`
	SoldIndividually  *bool                   `json:"sold_individually,omitempty"`
	Weight            *go_types.Float64String `json:"weight,omitempty"`
	Dimensions        *ProductDimensions      `json:"dimensions,omitempty"`
	ShippingRequired  *bool                   `json:"shipping_required,omitempty"`
	ShippingTaxable   *bool                   `json:"shipping_taxable,omitempty"`
	ShippingClass     *string                 `json:"shipping_class,omitempty"`
	ShippingClassId   *int64                  `json:"shipping_class_id,omitempty"`
	ReviewsAllowed    *bool                   `json:"reviews_allowed,omitempty"`
	AverageRating     *go_types.Float64String `json:"average_rating,omitempty"`
	RatingCount       *int64                  `json:"rating_count,omitempty"`
	RelatedIds        *[]int64                `json:"related_ids,omitempty"`
	UpsellIds         *[]int64                `json:"upsell_ids,omitempty"`
	CrossSellIds      *[]int64                `json:"cross_sell_ids,omitempty"`
	ParentId          *int64                  `json:"parent_id,omitempty"`
	PurchaseNote      *string                 `json:"purchase_note,omitempty"`
	Categories        *[]ProductCategory      `json:"categories,omitempty"`
	Tags              *[]ProductTag           `json:"tags,omitempty"`
	Images            *[]ProductImage         `json:"images,omitempty"`
	Attributes        *[]ProductAttribute     `json:"attributes,omitempty"`
	DefaultAttributes *[]ProductAttribute     `json:"default_attributes,omitempty"`
	Variations        *[]int64                `json:"variations,omitempty"`
	GroupedProducts   *[]int64                `json:"grouped_products,omitempty"`
	MenuOrder         *int64                  `json:"menu_order,omitempty"`
	MetaData          *[]ProductMetaData      `json:"meta_data,omitempty"`
	Brands            *[]ProductBrand         `json:"brands,omitempty"`
}

type ProductDimensions struct {
	Length go_types.Float64String `json:"length"`
	Width  go_types.Float64String `json:"width"`
	Height go_types.Float64String `json:"height"`
}

type ProductCategory struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type ProductTag struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type ProductImage struct {
	Id              int64                  `json:"id"`
	DateCreated     w_types.DateTimeString `json:"date_created"`
	DateCreatedGmt  w_types.DateTimeString `json:"date_created_gmt"`
	DateModified    w_types.DateTimeString `json:"date_modified"`
	DateModifiedGmt w_types.DateTimeString `json:"date_modified_gmt"`
	Src             string                 `json:"src"`
	Name            string                 `json:"name"`
	Alt             string                 `json:"alt"`
}

type ProductAttribute struct {
	Id        int64    `json:"id"`
	Name      string   `json:"name"`
	Position  int64    `json:"position"`
	Visible   bool     `json:"visible"`
	Variation bool     `json:"variation"`
	Options   []string `json:"options"`
}

type ProductMetaData struct {
	Id    int64       `json:"id"`
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type ProductMetaDataJSON struct {
	Id    int64           `json:"id"`
	Key   string          `json:"key"`
	Value json.RawMessage `json:"value"`
}

type BatchUpdateProductsInput struct {
	Update []Product
}

func (p *ProductMetaData) UnmarshalJSON(data []byte) error {
	var res ProductMetaDataJSON

	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	p.Id = res.Id
	p.Key = res.Key
	p.Value = strings.Trim(string(res.Value), `,omitempty"`)

	return nil
}

type GetProductsContext string

const (
	GetProductsContextView GetProductsContext = "view"
	GetProductsContextEdit GetProductsContext = "edit"
)

type GetProductsOrder string

const (
	GetProductsOrderAsc  GetProductsOrder = "asc"
	GetProductsOrderDesc GetProductsOrder = "desc"
)

type GetProductsOrderBy string

const (
	GetProductsOrderByDate    GetProductsOrderBy = "date"
	GetProductsOrderById      GetProductsOrderBy = "id"
	GetProductsOrderByInclude GetProductsOrderBy = "include"
	GetProductsOrderByTitle   GetProductsOrderBy = "title"
	GetProductsOrderBySlug    GetProductsOrderBy = "slug"
)

type GetProductsStatus string

const (
	GetProductsStatusAny     GetProductsStatus = "any"
	GetProductsStatusDraft   GetProductsStatus = "draft"
	GetProductsStatusPending GetProductsStatus = "pending"
	GetProductsStatusPrivate GetProductsStatus = "private"
	GetProductsStatusPublish GetProductsStatus = "publish"
)

type GetProductsType string

const (
	GetProductsTypeSimple   GetProductsType = "simple"
	GetProductsTypeGrouped  GetProductsType = "grouped"
	GetProductsTypeExternal GetProductsType = "external"
	GetProductsTypeVariable GetProductsType = "variable"
)

type GetProductsTaxClass string

const (
	GetProductsTaxClassStandard    GetProductsTaxClass = "standard"
	GetProductsTaxClassReducedRate GetProductsTaxClass = "reduced-rate"
	GetProductsTaxClassZeroRate    GetProductsTaxClass = "zero-rate"
)

type GetProductsStockStatus string

const (
	GetProductsStockStatusInStock     GetProductsStockStatus = "instock"
	GetProductsStockStatusOutOfStock  GetProductsStockStatus = "outofstock"
	GetProductsStockStatusOnBackorder GetProductsStockStatus = "onbackorder"
)

type GetProductsConfig struct {
	Context       *GetProductsContext
	Page          *uint // nil = all pages
	PerPage       *uint
	Search        *string
	After         *time.Time
	Before        *time.Time
	Exclude       *[]uint
	Include       *[]uint
	Offset        *uint
	Order         *GetProductsOrder
	OrderBy       *GetProductsOrderBy
	Parent        *[]uint
	ParentExclude *[]uint
	Slug          *string
	Status        *GetProductsStatus
	Type          *GetProductsType
	Sku           *string
	Featured      *bool
	Category      *string
	Tag           *string
	ShippingClass *string
	Attribute     *string
	AttributeTerm *string
	TaxClass      *GetProductsTaxClass
	OnSale        *bool
	MinPrice      *int64
	MaxPrice      *int64
	StockStatus   *GetProductsStockStatus
}

// GetProducts returns all products
func (service *Service) GetProducts(config *GetProductsConfig) (*[]Product, *errortools.Error) {
	values := url.Values{}
	endpoint := "products"

	if config != nil {
		if config.Context != nil {
			values.Set("context", string(*config.Context))
		}
		if config.PerPage != nil {
			values.Set("per_page", fmt.Sprintf("%v", *config.PerPage))
		}
		if config.Search != nil {
			values.Set("search", string(*config.Search))
		}
		if config.After != nil {
			values.Set("after", config.After.Format(DateFormat))
		}
		if config.Before != nil {
			values.Set("before", config.Before.Format(DateFormat))
		}
		if config.Exclude != nil {
			values.Set("exclude", UIntArrayToString(*config.Exclude))
		}
		if config.Include != nil {
			values.Set("include", UIntArrayToString(*config.Include))
		}
		if config.Offset != nil {
			values.Set("offset", fmt.Sprintf("%v", *config.Offset))
		}
		if config.Order != nil {
			values.Set("order", string(*config.Order))
		}
		if config.OrderBy != nil {
			values.Set("orderby", string(*config.OrderBy))
		}
		if config.Parent != nil {
			values.Set("parent", UIntArrayToString(*config.Parent))
		}
		if config.ParentExclude != nil {
			values.Set("parent_exclude", UIntArrayToString(*config.ParentExclude))
		}
		if config.Slug != nil {
			values.Set("slug", *config.Slug)
		}
		if config.Status != nil {
			values.Set("status", string(*config.Status))
		}
		if config.Type != nil {
			values.Set("type", string(*config.Type))
		}
		if config.Sku != nil {
			values.Set("sku", *config.Sku)
		}
		if config.Featured != nil {
			values.Set("featured", fmt.Sprintf("%v", *config.Featured))
		}
		if config.Category != nil {
			values.Set("category", *config.Category)
		}
		if config.Tag != nil {
			values.Set("tag", *config.Tag)
		}
		if config.ShippingClass != nil {
			values.Set("shipping_class", *config.ShippingClass)
		}
		if config.Attribute != nil {
			values.Set("attribute", *config.Attribute)
		}
		if config.AttributeTerm != nil {
			values.Set("attribute_term", *config.AttributeTerm)
		}
		if config.TaxClass != nil {
			values.Set("tax_class", string(*config.TaxClass))
		}
		if config.OnSale != nil {
			values.Set("on_sale", fmt.Sprintf("%v", *config.OnSale))
		}
		if config.MinPrice != nil {
			values.Set("min_price", fmt.Sprintf("%v", *config.MinPrice))
		}
		if config.MaxPrice != nil {
			values.Set("max_price", fmt.Sprintf("%v", *config.MaxPrice))
		}
		if config.StockStatus != nil {
			values.Set("stock_status", string(*config.StockStatus))
		}
	}

	page := 1
	maxPage := page
	if config != nil {
		if config.Page != nil {
			page = int(*config.Page)
		}
	}

	var products []Product

	for page <= maxPage {
		values.Set("per_page", fmt.Sprintf("%v", 100))
		values.Set("page", fmt.Sprintf("%v", page))

		var products_ []Product

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("%s?%s", endpoint, values.Encode())),
			ResponseModel: &products_,
		}

		_, response, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		products = append(products, products_...)

		if config != nil {
			if config.Page != nil {
				break
			}
		}

		maxPage, e = TotalPages(response)
		if e != nil {
			return nil, e
		}

		page++
	}

	return &products, nil
}

// UpdateProduct updates a specific product
func (service *Service) UpdateProduct(product *Product) (*Product, *errortools.Error) {
	if product == nil {
		return nil, errortools.ErrorMessage("Product is a nil pointer")
	}
	if product.Id == nil {
		return nil, errortools.ErrorMessage("ProductId is a nil pointer")
	}

	updatedProduct := Product{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPut,
		Url:           service.url(fmt.Sprintf("products/%v", *product.Id)),
		BodyModel:     product,
		ResponseModel: &updatedProduct,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &updatedProduct, nil
}

// BatchUpdateProducts updates multiple products
func (service *Service) BatchUpdateProducts(products []Product) *errortools.Error {
	if products == nil {
		return errortools.ErrorMessage("Products is a nil pointer")
	}

	if len(products) == 0 {
		return nil
	}

	if len(products) > 100 {
		return errortools.ErrorMessage("Maximum 100 products can be updated at once")
	}

	batchUpdateProductsInput := BatchUpdateProductsInput{
		Update: products,
	}

	updatedProducts := []Product{}
	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPut,
		Url:           service.url("products/batch"),
		BodyModel:     batchUpdateProductsInput,
		ResponseModel: &updatedProducts,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return e
	}
	return nil
}

// UpdateProductBrands updates the brands of a specific product
func (service *Service) UpdateProductBrands(productId int64, brands []int64) (*Product, *errortools.Error) {
	updatedProduct := Product{}

	requestConfig := go_http.RequestConfig{
		Method: http.MethodPut,
		Url:    service.url(fmt.Sprintf("products/%v", productId)),
		BodyModel: struct {
			Brands []int64 `json:"brands"`
		}{brands},
		ResponseModel: &updatedProduct,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &updatedProduct, nil
}

// CreateProduct creates a product
func (service *Service) CreateProduct(product *Product) (*Product, *errortools.Error) {
	if product == nil {
		return nil, errortools.ErrorMessage("Product is a nil pointer")
	}

	createdProduct := Product{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           service.url("products"),
		BodyModel:     product,
		ResponseModel: &createdProduct,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &createdProduct, nil
}

// DeleteProduct deletes a product
func (service *Service) DeleteProduct(productId int64, force bool) *errortools.Error {
	var values = url.Values{}
	values.Set("force", fmt.Sprintf("%v", force))

	requestConfig := go_http.RequestConfig{
		Method: http.MethodDelete,
		Url:    service.url(fmt.Sprintf("products/%v?%s", productId, values.Encode())),
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return e
	}

	return nil
}
