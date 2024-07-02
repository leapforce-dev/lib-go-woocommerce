package woocommerce

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
	w_types "github.com/leapforce-libraries/go_woocommerce/types"
)

// Order stores Order from Service
type Order struct {
	Id                 int64                   `json:"id"`
	ParentId           int64                   `json:"parent_id"`
	Number             string                  `json:"number"`
	OrderKey           string                  `json:"order_key"`
	CreatedVia         string                  `json:"created_via"`
	Version            string                  `json:"version"`
	Status             string                  `json:"status"`
	Currency           string                  `json:"currency"`
	CurrencySymbol     string                  `json:"currency_symbol"`
	DateCreated        w_types.DateTimeString  `json:"date_created"`
	DateCreatedGmt     w_types.DateTimeString  `json:"date_created_gmt"`
	DateModified       w_types.DateTimeString  `json:"date_modified"`
	DateModifiedGmt    w_types.DateTimeString  `json:"date_modified_gmt"`
	DiscountTotal      go_types.Float64String  `json:"discount_total"`
	DiscountTax        go_types.Float64String  `json:"discount_tax"`
	ShippingTotal      go_types.Float64String  `json:"shipping_total"`
	ShippingTax        go_types.Float64String  `json:"shipping_tax"`
	CartTax            go_types.Float64String  `json:"cart_tax"`
	Total              go_types.Float64String  `json:"total"`
	TotalTax           go_types.Float64String  `json:"total_tax"`
	PricesIncludeTax   bool                    `json:"prices_include_tax"`
	CustomerId         int64                   `json:"customer_id"`
	CustomerIpAddress  string                  `json:"customer_ip_address"`
	CustomerUserAgent  string                  `json:"customer_user_agent"`
	CustomerNote       string                  `json:"customer_note"`
	Billing            OrderBilling            `json:"billing"`
	Shipping           OrderShipping           `json:"shipping"`
	PaymentMethod      string                  `json:"payment_method"`
	PaymentMethodTitle string                  `json:"payment_method_title"`
	TransactionId      string                  `json:"transaction_id"`
	DatePaid           *w_types.DateTimeString `json:"date_paid"`
	DatePaidGmt        *w_types.DateTimeString `json:"date_paid_gmt"`
	DateCompleted      *w_types.DateTimeString `json:"date_completed"`
	DateCompletedGmt   *w_types.DateTimeString `json:"date_completed_gmt"`
	CartHash           string                  `json:"cart_hash"`
	MetaData           []OrderMetaData         `json:"meta_data"`
	LineItems          []OrderLineItem         `json:"line_items"`
	TaxLines           []OrderTaxLine          `json:"tax_lines"`
	ShippingLines      []OrderShippingLine     `json:"shipping_lines"`
	FeeLines           []OrderFeeLine          `json:"fee_lines"`
	CouponLines        []OrderCouponLine       `json:"coupon_lines"`
	Refunds            []OrderRefund           `json:"refunds"`
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
	Id    int64           `json:"id"`
	Key   string          `json:"key"`
	Value json.RawMessage `json:"value"`
}

func (o OrderMetaData) GetValueString() (string, error) {
	var s string

	err := json.Unmarshal(o.Value, &s)
	if err != nil {
		return "", err
	}

	return s, nil
}

func (o OrderMetaData) GetValueMap() (map[string]string, error) {
	m := make(map[string]string)

	err := json.Unmarshal(o.Value, &m)
	if err != nil {
		return m, err
	}

	return m, nil
}

type OrderLineItem struct {
	Id          int64           `json:"id"`
	Name        string          `json:"name"`
	ProductId   int64           `json:"product_id"`
	VariationId int64           `json:"variation_id"`
	Quantity    int64           `json:"quantity"`
	TaxClass    string          `json:"tax_class"`
	Subtotal    string          `json:"subtotal"`
	SubtotalTax string          `json:"subtotal_tax"`
	Total       string          `json:"total"`
	TotalTax    string          `json:"total_tax"`
	Taxes       []OrderTax      `json:"taxes"`
	MetaData    []OrderMetaData `json:"meta_data"`
	Sku         string          `json:"sku"`
	Price       float64         `json:"price"`
}

type OrderTax struct {
	Id               int64           `json:"id"`
	RateCode         string          `json:"rate_code"`
	RateId           string          `json:"rate_id"`
	Label            string          `json:"label"`
	Compound         bool            `json:"compound"`
	TaxTotal         string          `json:"tax_total"`
	ShippingTaxTotal string          `json:"shipping_tax_total"`
	MetaData         []OrderMetaData `json:"meta_data"`
}

type OrderTaxLine struct {
	Id               int64           `json:"id"`
	RateCode         string          `json:"rate_code"`
	RateId           string          `json:"rate_id"`
	Label            string          `json:"label"`
	Compound         bool            `json:"compound"`
	TaxTotal         string          `json:"tax_total"`
	ShippingTaxTotal string          `json:"shipping_tax_total"`
	MetaData         []OrderMetaData `json:"meta_data"`
}

type OrderShippingLine struct {
	Id          int64           `json:"id"`
	MethodTitle string          `json:"method_title"`
	MethodId    string          `json:"method_id"`
	Total       string          `json:"total"`
	TotalTax    string          `json:"total_tax"`
	Taxes       []OrderTax      `json:"taxes"`
	MetaData    []OrderMetaData `json:"meta_data"`
}

type OrderFeeLine struct {
	Id        int64           `json:"id"`
	Name      string          `json:"name"`
	TaxClass  string          `json:"tax_class"`
	TaxStatus string          `json:"tax_status"`
	Total     string          `json:"total"`
	TotalTax  string          `json:"total_tax"`
	Taxes     []OrderTax      `json:"taxes"`
	MetaData  []OrderMetaData `json:"meta_data"`
}

type OrderCouponLine struct {
	Id          int64           `json:"id"`
	Code        string          `json:"code"`
	Discount    string          `json:"discount"`
	DiscountTax string          `json:"discount_tax"`
	MetaData    []OrderMetaData `json:"meta_data"`
}

type OrderRefund struct {
	Id     int64  `json:"id"`
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
	GetOrdersOrderById      GetOrdersOrderBy = "id"
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
	ModifiedAfter    *time.Time
	ModifiedBefore   *time.Time
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
func (service *Service) GetOrders(config *GetOrdersConfig) (*[]Order, *errortools.Error) {
	values := url.Values{}
	endpoint := "orders"

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
		if config.ModifiedAfter != nil {
			values.Set("modified_after", config.ModifiedAfter.Format(DateFormat))
		}
		if config.ModifiedBefore != nil {
			values.Set("modified_before", config.ModifiedBefore.Format(DateFormat))
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
		if config.Status != nil {
			values.Set("status", string(*config.Status))
		}
		if config.Customer != nil {
			values.Set("customer", fmt.Sprintf("%v", *config.Customer))
		}
		if config.Product != nil {
			values.Set("product", fmt.Sprintf("%v", *config.Product))
		}
		if config.DecimalPositions != nil {
			values.Set("dp", fmt.Sprintf("%v", *config.DecimalPositions))
		}
	}

	page := 1
	maxPage := page
	if config.Page != nil {
		page = int(*config.Page)
	}

	var orders []Order

	for page <= maxPage {
		values.Set("page", fmt.Sprintf("%v", page))

		path := fmt.Sprintf("%s?%s", endpoint, values.Encode())

		var _orders []Order

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(path),
			ResponseModel: &_orders,
		}

		_, response, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		orders = append(orders, _orders...)

		if config.Page == nil {
			maxPage, e = TotalPages(response)
			if e != nil {
				return nil, e
			}
		}

		page++
	}

	return &orders, nil
}

// UpdateOrder updates all orders
func (service *Service) UpdateOrder(order *Order) (*Order, *errortools.Error) {
	if order == nil {
		return nil, errortools.ErrorMessage("Order is a nil pointer")
	}

	updatedOrder := Order{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPut,
		Url:           service.url(fmt.Sprintf("orders/%v", order.Id)),
		BodyModel:     order,
		ResponseModel: &updatedOrder,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &updatedOrder, nil
}
