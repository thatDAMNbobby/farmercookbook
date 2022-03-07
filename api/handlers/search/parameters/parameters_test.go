package parameters

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	t.Run("no page returns 1", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/thing?q=thing+to+query&page=8&perPage=5", nil)
		actual, err := Get(req)
		require.NoError(t, err)
		expected := &Params{
			Query:   "thing to query",
			Page:    8,
			PerPage: 5,
		}
		require.Equal(t, expected, actual)
	})
}

func TestGetQuery(t *testing.T) {
	t.Run("no query returns empty string", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/thing", nil)

		actual := getQuery(req)
		expected := ""
		require.Equal(t, expected, actual)
	})

	t.Run("query returns query", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/thing?q=thing+to+query", nil)

		actual := getQuery(req)
		expected := "thing to query"
		require.Equal(t, expected, actual)
	})
}

func TestGetPage(t *testing.T) {
	t.Run("valid page returns page", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/thing?page=5", nil)

		expected := 5
		actual := getPage(req)
		require.Equal(t, expected, actual)
	})

	t.Run("no page returns 1", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/thing", nil)

		expected := 1
		actual := getPage(req)
		require.Equal(t, expected, actual)
	})
}

func TestGetPerPage(t *testing.T) {
	t.Run("valid perPage returns perPage", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/thing?perPage=5", nil)

		expected := 5
		actual := getPerPage(req)
		require.Equal(t, expected, actual)
	})

	t.Run("non-numeric perPage returns default", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/thing?perPage=five", nil)

		expected := defaultPerPage
		actual := getPerPage(req)
		require.Equal(t, expected, actual)
	})

	t.Run("negative perPage returns default", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/thing?perPage=-1", nil)

		expected := defaultPerPage
		actual := getPerPage(req)
		require.Equal(t, expected, actual)
	})

	t.Run("zero perPage returns default", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/thing?perPage=0", nil)

		expected := defaultPerPage
		actual := getPerPage(req)
		require.Equal(t, expected, actual)
	})

	t.Run("no perPage returns default", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/thing", nil)

		expected := defaultPerPage
		actual := getPerPage(req)
		require.Equal(t, expected, actual)
	})

	t.Run("perPage beyond max returns default", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/thing?perPage=99", nil)

		expected := defaultPerPage
		actual := getPerPage(req)
		require.Equal(t, expected, actual)
	})
}
