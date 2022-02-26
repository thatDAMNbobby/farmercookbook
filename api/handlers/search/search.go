package search

import (
	"github.com/thatDAMNbobby/farmercookbook/clients/elasticsearch"
	"github.com/thatDAMNbobby/farmercookbook/clients/render"
	searchsvc "github.com/thatDAMNbobby/farmercookbook/services/search"
	"net/http"
)

type Deps struct {
	Search searchsvc.Search
	Render render.Render
}

type impl struct {
	Deps *Deps
}

type Search interface {
	Query(w http.ResponseWriter, r *http.Request)
}

func New(deps *Deps) Search {
	return &impl{Deps: deps}
}

func (impl *impl) Query(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	query := buildQuery(r)

	resp, err := impl.Deps.Search.Search(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	impl.renderResponse(w, http.StatusOK, resp)

}

func (impl *impl) renderResponse(w http.ResponseWriter, status int, v interface{}) {
	impl.Deps.Render.JSON(w, status, v)
}

func buildQuery(r *http.Request) elasticsearch.SearchRequest {
	return elasticsearch.SearchRequest{
		Query: r.FormValue("q"),
	}
}
