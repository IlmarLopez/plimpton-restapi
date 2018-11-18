package apis

import (
	"strconv"

	routing "github.com/go-ozzo/ozzo-routing"
	"myhmp.org/plimpton-restapi/util"
)

const (
	// DefaultPageSize determines the item size per page
	DefaultPageSize int = 100
	// MaxPageSize size maximo per page
	MaxPageSize int = 1000
	// MaxUploadSize Tamaño máximo de carga
	MaxUploadSize int = 2 * 1024 // 2mb
)

func getPaginatedListFromRequest(c *routing.Context, count int) *util.PaginatedList {
	page := parseInt(c.Query("page"), 1)
	perPage := parseInt(c.Query("per_page"), DefaultPageSize)
	if perPage <= 0 {
		perPage = DefaultPageSize
	}
	if perPage > MaxPageSize {
		perPage = MaxPageSize
	}
	return util.NewPaginatedList(page, perPage, count)
}

func parseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil {
		return result
	}
	return defaultValue
}
