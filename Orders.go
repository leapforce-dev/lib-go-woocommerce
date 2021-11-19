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

// Order stores Order from Service
//
type Order struct {
	ID                 int                     `json:"id"`
	ParentID           int                     `json:"parent_id"`
	Number             string                  `json:"number"`
	OrderKey           string                  `json:"order_key"`
	CreatedVia         string                  `json:"created_via"`
	Version            string                  `json:"version"`
	Status             string                  `json:"status"`
	Currency           string                  `json:"currency"`
	CurrencySymbol     string                  `json:"currency_symbol"`
	DateCreated        w_types.DateTimeString  `json:"date_created"`
	DateCreatedGMT     w_types.DateTimeString  `json:"date_created_gmt"`
	DateModified       w_types.DateTimeString  `json:"date_modified"`
	DateModifiedGMT    w_types.DateTimeString  `json:"date_modified_gmt"`
	DiscountTotal      go_types.Float64String  `json:"discount_total"`
	DiscountTax        go_types.Float64String  `json:"discount_tax"`
	ShippingTotal      go_types.Float64String  `json:"shipping_total"`
	ShippingTax        go_types.Float64String  `json:"shipping_tax"`
	CartTax            go_types.Float64String  `json:"cart_tax"`
	Total              go_types.Float64String  `json:"total"`
	TotalTax           go_types.Float64String  `json:"total_tax"`
	PricesIncludeTax   bool                    `json:"prices_include_tax"`
	CustomerID         int                     `json:"customer_id"`
	CustomerIPAddress  string                  `json:"customer_ip_address"`
	CustomerUserAgent  string                  `json:"customer_user_agent"`
	CustomerNote       string                  `json:"customer_note"`
	Billing            *OrderBilling           `json:"billing"`
	Shipping           *OrderShipping          `json:"shipping"`
	PaymentMethod      string                  `json:"payment_method"`
	PaymentMethodTitle string                  `json:"payment_method_title"`
	TransactionID      string                  `json:"transaction_id"`
	DatePaid           *w_types.DateTimeString `json:"date_paid"`
	DatePaidGMT        *w_types.DateTimeString `json:"date_paid_gmt"`
	DateCompleted      *w_types.DateTimeString `json:"date_completed"`
	DateCompletedGMT   *w_types.DateTimeString `json:"date_completed_gmt"`
	CartHash           string                  `json:"cart_hash"`
	MetaData           []*OrderMetaData        `json:"meta_data"`
	LineItems          []*OrderLineItem        `json:"line_items"`
	TaxLines           []*OrderTaxLine         `json:"tax_lines"`
	ShippingLines      []*OrderShippingLine    `json:"shipping_lines"`
	FeeLines           []*OrderFeeLine         `json:"fee_lines"`
	CouponLines        []*OrderCouponLine      `json:"coupon_lines"`
	Refunds            []*OrderRefund          `json:"refunds"`
	SetPaid            bool                    `json:"set_paid"`
}

type OrderBilling struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Company   string `json:"company"`
	Address1  string `json:"address_1"`
	Address2  string `json:"address_2"`
	City      string `json:"city"`
	State     string `json:"state"`
	Postcode  string `json:"postcode"`
	Country   string `json:"country"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type OrderShipping struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Company   string `json:"company"`
	Address1  string `json:"address_1"`
	Address2  string `json:"address_2"`
	City      string `json:"city"`
	State     string `json:"state"`
	Postcode  string `json:"postcode"`
	Country   string `json:"country"`
}

type OrderMetaData struct {
	ID    int    `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type OrderMetaDataJSON struct {
	ID    int             `json:"id"`
	Key   string          `json:"key"`
	Value json.RawMessage `json:"value"`
}

func (p *OrderMetaData) UnmarshalJSON(data []byte) error {
	var res OrderMetaDataJSON

	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	p.ID = res.ID
	p.Key = res.Key
	p.Value = strings.Trim(string(res.Value), `"`)

	return nil
}

type OrderLineItem struct {
	ID          int              `json:"id"`
	Name        string           `json:"name"`
	ProductID   int              `json:"product_id"`
	VariationID int              `json:"variation_id"`
	Quantity    int              `json:"quantity"`
	TaxClass    string           `json:"tax_class"`
	Subtotal    string           `json:"subtotal"`
	SubtotalTax string           `json:"subtotal_tax"`
	Total       string           `json:"total"`
	TotalTax    string           `json:"total_tax"`
	Taxes       []*OrderTax      `json:"taxes"`
	MetaData    []*OrderMetaData `json:"meta_data"`
	SKU         string           `json:"sku"`
	Price       float64          `json:"price"`
}

type OrderTax struct {
	ID               int              `json:"id"`
	RateCode         string           `json:"rate_code"`
	RateID           string           `json:"rate_id"`
	Label            string           `json:"label"`
	Compound         bool             `json:"compound"`
	TaxTotal         string           `json:"tax_total"`
	ShippingTaxTotal string           `json:"shipping_tax_total"`
	MetaData         []*OrderMetaData `json:"meta_data"`
}

type OrderTaxLine struct {
	ID               int              `json:"id"`
	RateCode         string           `json:"rate_code"`
	RateID           string           `json:"rate_id"`
	Label            string           `json:"label"`
	Compound         bool             `json:"compound"`
	TaxTotal         string           `json:"tax_total"`
	ShippingTaxTotal string           `json:"shipping_tax_total"`
	MetaData         []*OrderMetaData `json:"meta_data"`
}

type OrderShippingLine struct {
	ID          int              `json:"id"`
	MethodTitle string           `json:"method_title"`
	MethodID    string           `json:"method_id"`
	Total       string           `json:"total"`
	TotalTax    string           `json:"total_tax"`
	Taxes       []*OrderTax      `json:"taxes"`
	MetaData    []*OrderMetaData `json:"meta_data"`
}

type OrderFeeLine struct {
	ID        int              `json:"id"`
	Name      string           `json:"name"`
	TaxClass  string           `json:"tax_class"`
	TaxStatus string           `json:"tax_status"`
	Total     string           `json:"total"`
	TotalTax  string           `json:"total_tax"`
	Taxes     []*OrderTax      `json:"taxes"`
	MetaData  []*OrderMetaData `json:"meta_data"`
}

type OrderCouponLine struct {
	ID          int              `json:"id"`
	Code        string           `json:"code"`
	Discount    string           `json:"discount"`
	DiscountTax string           `json:"discount_tax"`
	MetaData    []*OrderMetaData `json:"meta_data"`
}

type OrderRefund struct {
	ID     int    `json:"id"`
	Reason string `json:"reason"`
	Total  string `json:"total"`
}

type GetOrdersContext string

const (
	GetOrdersContextView GetOrdersContext = "view"
	GetOrdersContextEdit GetOrdersContext = "edit"
)

type GetOrdersOrder string

const (
	GetOrdersOrderAsc  GetOrdersOrder = "asc"
	GetOrdersOrderDesc GetOrdersOrder = "desc"
)

type GetOrdersOrderBy string

const (
	GetOrdersOrderByDate    GetOrdersOrderBy = "date"
	GetOrdersOrderByID      GetOrdersOrderBy = "id"
	GetOrdersOrderByInclude GetOrdersOrderBy = "include"
	GetOrdersOrderByTitle   GetOrdersOrderBy = "title"
	GetOrdersOrderBySlug    GetOrdersOrderBy = "slug"
)

type GetOrdersStatus string

const (
	GetOrdersStatusAny        GetOrdersStatus = "any"
	GetOrdersStatusPending    GetOrdersStatus = "pending"
	GetOrdersStatusProcessing GetOrdersStatus = "processing"
	GetOrdersStatusOnHold     GetOrdersStatus = "on-hold"
	GetOrdersStatusCompleted  GetOrdersStatus = "completed"
	GetOrdersStatusCancelled  GetOrdersStatus = "cancelled"
	GetOrdersStatusRefunded   GetOrdersStatus = "refunded"
	GetOrdersStatusFailed     GetOrdersStatus = "failed"
	GetOrdersStatusTrash      GetOrdersStatus = "trash"
)

type GetOrdersConfig struct {
	Context          *GetOrdersContext
	Page             *uint // nil = all pages
	PerPage          *uint
	Search           *string
	After            *time.Time
	Before           *time.Time
	Exclude          *[]uint
	Include          *[]uint
	Offset           *uint
	Order            *GetOrdersOrder
	OrderBy          *GetOrdersOrderBy
	Parent           *[]uint
	ParentExclude    *[]uint
	Status           *GetOrdersStatus
	Customer         *uint
	Product          *uint
	DecimalPositions *uint
}

// GetOrders returns all orders
//
func (service *Service) GetOrders(filter *GetOrdersConfig) (*[]Order, *errortools.Error) {
	values := url.Values{}
	endpoint := "orders"

	if filter != nil {
		if filter.Context != nil {
			values.Set("context", string(*filter.Context))
		}
		if filter.PerPage != nil {
			values.Set("per_page", fmt.Sprintf("%v", *filter.PerPage))
		}
		if filter.Search != nil {
			values.Set("search", string(*filter.Search))
		}
		if filter.After != nil {
			values.Set("after", filter.After.Format(DateFormat))
		}
		if filter.Before != nil {
			values.Set("before", filter.Before.Format(DateFormat))
		}
		if filter.Exclude != nil {
			values.Set("exclude", UIntArrayToString(*filter.Exclude))
		}
		if filter.Include != nil {
			values.Set("include", UIntArrayToString(*filter.Include))
		}
		if filter.Offset != nil {
			values.Set("offset", fmt.Sprintf("%v", *filter.Offset))
		}
		if filter.Order != nil {
			values.Set("order", string(*filter.Order))
		}
		if filter.OrderBy != nil {
			values.Set("orderby", string(*filter.OrderBy))
		}
		if filter.Parent != nil {
			values.Set("parent", UIntArrayToString(*filter.Parent))
		}
		if filter.ParentExclude != nil {
			values.Set("parent_exclude", UIntArrayToString(*filter.ParentExclude))
		}
		if filter.Status != nil {
			values.Set("status", string(*filter.Status))
		}
		if filter.Customer != nil {
			values.Set("customer", fmt.Sprintf("%v", *filter.Customer))
		}
		if filter.Product != nil {
			values.Set("product", fmt.Sprintf("%v", *filter.Product))
		}
		if filter.DecimalPositions != nil {
			values.Set("dp", fmt.Sprintf("%v", *filter.DecimalPositions))
		}
	}

	page := 1
	maxPage := page
	if filter.Page != nil {
		page = int(*filter.Page)
	}

	orders := []Order{}

	for page <= maxPage {
		values.Set("page", fmt.Sprintf("%v", page))

		path := fmt.Sprintf("%s?%s", endpoint, values.Encode())

		_orders := []Order{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			URL:           service.url(path),
			ResponseModel: &_orders,
		}

		_, response, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		orders = append(orders, _orders...)

		if filter.Page == nil {
			maxPage, e = TotalPages(response)
			if e != nil {
				return nil, e
			}
		}

		page++
	}

	return &orders, nil
}
