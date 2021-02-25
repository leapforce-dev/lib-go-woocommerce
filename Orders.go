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
	ID          int    `json:"id"`
	DateCreated string `json:"date_created"`
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
			URL:           service.url(path),
			ResponseModel: &_orders,
		}
		fmt.Println(service.url(path))
		_, response, e := service.get(&requestConfig)
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
