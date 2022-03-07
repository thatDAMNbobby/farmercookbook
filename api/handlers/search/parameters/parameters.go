package parameters

import (
	"net/http"
	"strconv"
)

func Get(r *http.Request) (Params, error) {

	params := Params{
		Query:   getQuery(r),
		Page:    getPage(r),
		PerPage: getPerPage(r),
	}

	return params, nil
}

func getQuery(r *http.Request) string {
	return r.URL.Query().Get("q")
}

func getPage(r *http.Request) int {
	values := r.URL.Query()
	page, err := strconv.Atoi(values.Get("page"))
	if err != nil || page < 1 {
		return 1
	}

	return page
}

func getPerPage(r *http.Request) int {
	values := r.URL.Query()
	perPage, err := strconv.Atoi(values.Get("perPage"))
	if err != nil || perPage < 1 || perPage > maxPerPage {
		return defaultPerPage
	}

	return perPage
}
