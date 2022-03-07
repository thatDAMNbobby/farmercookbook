package search

import (
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/thatDAMNbobby/farmercookbook/clients/elasticsearch/requests"
	"github.com/thatDAMNbobby/farmercookbook/clients/render"
	"github.com/thatDAMNbobby/farmercookbook/handlers/search/parameters"
	"github.com/thatDAMNbobby/farmercookbook/models"
	searchsvc "github.com/thatDAMNbobby/farmercookbook/services/search"
	"github.com/thatDAMNbobby/farmercookbook/utils"
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

	params, err := parameters.Get(r)
	if err != nil {
		impl.Deps.Render.JSON(w, http.StatusBadRequest, err)
		return
	}

	query := buildQuery(params)
	esResp, err := impl.Deps.Search.Search(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := buildResponse(esResp, params)

	impl.renderResponse(w, http.StatusOK, resp)

}

func buildResponse(esResponse *elastic.SearchResult, params parameters.Params) *Response {
	var resp Response

	totalResults := int(esResponse.Hits.TotalHits.Value)
	page := params.Page
	perPage := params.PerPage
	totalPages := totalResults / perPage

	var previousPage int
	if (page - 1) > 0 {
		previousPage = page - 1
	}

	var nextPage int
	if (page + 1) <= totalPages {
		nextPage = page + 1
	}

	begin := (page-1)*perPage + 1
	end := utils.Min(page*perPage, totalResults)

	resp.Pagination = Pagination{
		TotalResults: totalResults,
		CurrentPage:  page,
		TotalPages:   totalPages,
		NextPage:     nextPage,
		PreviousPage: previousPage,
		PerPage:      perPage,
		Begin:        begin,
		End:          end,
	}

	for _, hit := range esResponse.Hits.Hits {

		var recipe models.Recipe
		sourceJSON, err := hit.Source.MarshalJSON()
		if err != nil {
			fmt.Println(err)
		}
		if err := json.Unmarshal(sourceJSON, &recipe); err != nil {
			fmt.Println(err)
		}
		resp.Results = append(resp.Results, recipe)
	}

	return &resp
}

func (impl *impl) renderResponse(w http.ResponseWriter, status int, v interface{}) {
	impl.Deps.Render.JSON(w, status, v)
}

func buildQuery(params parameters.Params) *requests.SearchRequest {

	query := elastic.NewBoolQuery().
		Should(elastic.NewMatchQuery("name", params.Query).Boost(2)).
		Should(elastic.NewMatchQuery("description", params.Query).Boost(2)).
		Should(elastic.NewTermsQuery("author", params.Query).Boost(1))

	return &requests.SearchRequest{
		Type:    "recipe",
		Source:  []string{"*"},
		Query:   query,
		From:    (params.Page - 1) * params.PerPage,
		Size:    params.PerPage,
		Indices: []string{"recipes"},
	}
}
