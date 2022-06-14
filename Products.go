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
//
type Product struct {
	Id                int                     `json:"id"`
	Name              string                  `json:"name"`
	Slug              string                  `json:"slug"`
	Permalink         string                  `json:"permalink"`
	DateCreated       w_types.DateTimeString  `json:"date_created"`
	DateCreatedGmt    w_types.DateTimeString  `json:"date_created_gmt"`
	DateModified      w_types.DateTimeString  `json:"date_modified"`
	DateModifiedGmt   w_types.DateTimeString  `json:"date_modified_gmt"`
	Type              string                  `json:"type"`
	Status            string                  `json:"status"`
	Featured          bool                    `json:"featured"`
	CatalogVisibility string                  `json:"catalog_visibility"`
	Description       string                  `json:"description"`
	ShortDescription  string                  `json:"short_description"`
	Sku               string                  `json:"sku"`
	Price             go_types.Float64String  `json:"price"`
	RegularPrice      go_types.Float64String  `json:"regular_price"`
	SalePrice         go_types.Float64String  `json:"sale_price"`
	DateOnSaleFrom    *w_types.DateTimeString `json:"date_on_sale_from"`
	DateOnSaleFromGmt *w_types.DateTimeString `json:"date_on_sale_from_gmt"`
	DateOnSaleTo      *w_types.DateTimeString `json:"date_on_sale_to"`
	DateOnSaleToGmt   *w_types.DateTimeString `json:"date_on_sale_to_gmt"`
	PriceHtml         string                  `json:"price_html"`
	OnSale            bool                    `json:"on_sale"`
	Purchasable       bool                    `json:"purchasable"`
	TotalSales        int                     `json:"total_sales"`
	Virtual           bool                    `json:"virtual"`
	Downloadable      bool                    `json:"downloadable"`
	Downloads         json.RawMessage         `json:"downloads"`
	DownloadLimit     int                     `json:"download_limit"`
	DownloadExpiry    int                     `json:"download_expiry"`
	ExternalUrl       string                  `json:"external_url"`
	ButtonText        string                  `json:"button_text"`
	TaxStatus         string                  `json:"tax_status"`
	TaxClass          string                  `json:"tax_class"`
	ManageStock       bool                    `json:"manage_stock"`
	StockQuantity     *int                    `json:"stock_quantity"`
	StockStatus       string                  `json:"stock_status"`
	Backorders        string                  `json:"backorders"`
	BackordersAllowed bool                    `json:"backorders_allowed"`
	Backordered       bool                    `json:"backordered"`
	SoldIndividually  bool                    `json:"sold_individually"`
	Weight            go_types.Float64String  `json:"weight"`
	Dimensions        ProductDimensions       `json:"dimensions"`
	ShippingRequired  bool                    `json:"shipping_required"`
	ShippingTaxable   bool                    `json:"shipping_taxable"`
	ShippingClass     string                  `json:"shipping_class"`
	ShippingClassId   int                     `json:"shipping_class_id"`
	ReviewsAllowed    bool                    `json:"reviews_allowed"`
	AverageRating     go_types.Float64String  `json:"average_rating"`
	RatingCount       int                     `json:"rating_count"`
	RelatedIds        []int                   `json:"related_ids"`
	UpsellIds         []int                   `json:"upsell_ids"`
	CrossSellIds      []int                   `json:"cross_sell_ids"`
	ParentId          int                     `json:"parent_id"`
	PurchaseNote      string                  `json:"purchase_note"`
	Categories        []ProductCategory       `json:"categories"`
	Tags              []string                `json:"tags"`
	Images            []ProductImage          `json:"images"`
	Attributes        []ProductAttribute      `json:"attributes"`
	DefaultAttributes []ProductAttribute      `json:"default_attributes"`
	Variations        []int                   `json:"variations"`
	GroupedProducts   []int                   `json:"grouped_products"`
	MenuOrder         int                     `json:"menu_order"`
	MetaData          []ProductMetaData       `json:"meta_data"`
}

type ProductDimensions struct {
	Length go_types.Float64String `json:"length"`
	Width  go_types.Float64String `json:"width"`
	Height go_types.Float64String `json:"height"`
}

type ProductCategory struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type ProductImage struct {
	Id              int                    `json:"id"`
	DateCreated     w_types.DateTimeString `json:"date_created"`
	DateCreatedGmt  w_types.DateTimeString `json:"date_created_gmt"`
	DateModified    w_types.DateTimeString `json:"date_modified"`
	DateModifiedGmt w_types.DateTimeString `json:"date_modified_gmt"`
	Src             string                 `json:"src"`
	Name            string                 `json:"name"`
	Alt             string                 `json:"alt"`
}

type ProductAttribute struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Position  int      `json:"position"`
	Visible   bool     `json:"visible"`
	Variation bool     `json:"variation"`
	Options   []string `json:"options"`
}

type ProductMetaData struct {
	Id    int    `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ProductMetaDataJSON struct {
	Id    int             `json:"id"`
	Key   string          `json:"key"`
	Value json.RawMessage `json:"value"`
}

func (p *ProductMetaData) UnmarshalJSON(data []byte) error {
	var res ProductMetaDataJSON

	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	p.Id = res.Id
	p.Key = res.Key
	p.Value = strings.Trim(string(res.Value), `"`)

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
	MinPrice      *int
	MaxPrice      *int
	StockStatus   *GetProductsStockStatus
}

// GetProducts returns all products
//
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
	if config.Page != nil {
		page = int(*config.Page)
	}

	products := []Product{}

	for page <= maxPage {
		values.Set("page", fmt.Sprintf("%v", page))

		path := fmt.Sprintf("%s?%s", endpoint, values.Encode())

		_products := []Product{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(path),
			ResponseModel: &_products,
		}

		_, response, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		products = append(products, _products...)

		if config.Page == nil {
			maxPage, e = TotalPages(response)
			if e != nil {
				return nil, e
			}
		}

		page++
	}

	return &products, nil
}

// UpdateProduct updates all products
//
func (service *Service) UpdateProduct(product *Product) (*Product, *errortools.Error) {
	if product == nil {
		return nil, errortools.ErrorMessage("Product is a nil pointer")
	}

	updatedProduct := Product{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPut,
		Url:           service.url(fmt.Sprintf("products/%v", product.Id)),
		BodyModel:     product,
		ResponseModel: &updatedProduct,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &updatedProduct, nil
}
