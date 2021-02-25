package woocommerce

import (
	"fmt"
	"net/url"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

// Order stores Order from Service
//
type Order struct {
	ID int `json:"id"`
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
	path := "orders"

	if filter != nil {
		if filter.Context != nil {
			values.Add("context", string(*filter.Context))
		}
		if filter.Page != nil {
			values.Add("page", fmt.Sprintf("%v", *filter.Page))
		}
		if filter.PerPage != nil {
			values.Add("per_page", fmt.Sprintf("%v", *filter.PerPage))
		}
		if filter.Search != nil {
			values.Add("search", string(*filter.Search))
		}
		if filter.After != nil {
			values.Add("after", filter.After.Format(DateFormat))
		}
		if filter.Before != nil {
			values.Add("before", filter.Before.Format(DateFormat))
		}
		if filter.Exclude != nil {
			values.Add("exclude", UIntArrayToString(*filter.Exclude))
		}
		if filter.Include != nil {
			values.Add("include", UIntArrayToString(*filter.Include))
		}
		if filter.Offset != nil {
			values.Add("offset", fmt.Sprintf("%v", *filter.Offset))
		}
		if filter.Order != nil {
			values.Add("order", string(*filter.Order))
		}
		if filter.OrderBy != nil {
			values.Add("orderby", string(*filter.OrderBy))
		}
		if filter.Parent != nil {
			values.Add("parent", UIntArrayToString(*filter.Parent))
		}
		if filter.ParentExclude != nil {
			values.Add("parent_exclude", UIntArrayToString(*filter.ParentExclude))
		}
		if filter.Status != nil {
			values.Add("status", string(*filter.Status))
		}
		if filter.Customer != nil {
			values.Add("customer", fmt.Sprintf("%v", *filter.Customer))
		}
		if filter.Product != nil {
			values.Add("product", fmt.Sprintf("%v", *filter.Product))
		}
		if filter.DecimalPositions != nil {
			values.Add("dp", fmt.Sprintf("%v", *filter.DecimalPositions))
		}
	}

	if len(values) > 0 {
		path = fmt.Sprintf("%s?%s", path, values.Encode())
	}

	orders := []Order{}

	requestConfig := go_http.RequestConfig{
		URL:           service.url(path),
		ResponseModel: &orders,
	}
	fmt.Println(service.url(path))
	_, _, e := service.get(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &orders, nil
}
