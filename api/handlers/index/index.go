package index

import (
	"encoding/json"
	"github.com/thatDAMNbobby/farmercookbook/clients/elasticsearch/requests"
	"github.com/thatDAMNbobby/farmercookbook/clients/render"
	"github.com/thatDAMNbobby/farmercookbook/services/index"
	"io/ioutil"
	"net/http"
)

type Deps struct {
	IndexService index.Index
	Render       render.Render
}

type impl struct {
	deps *Deps
}

type Index interface {
	Delete(w http.ResponseWriter, r *http.Request)
	Upsert(w http.ResponseWriter, r *http.Request)
}

func New(deps *Deps) Index {
	return &impl{deps}
}

func (impl *impl) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req, err := buildDeleteRequest(r)
	if err != nil {
		impl.respondBadRequest(w, err)
		return
	}

	err = impl.deps.IndexService.Delete(ctx, req)

}

func (impl *impl) Upsert(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req, err := buildUpsertRequest(r)
	if err != nil {
		impl.respondBadRequest(w, err)
	}

	err = impl.deps.IndexService.Upsert(ctx, req)
	if err != nil {
		impl.respondBadRequest(w, err)
	}

}

func buildDeleteRequest(r *http.Request) (*requests.DeleteRequest, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var req requests.DeleteRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		return nil, err
	}

	return &req, nil
}

func buildUpsertRequest(r *http.Request) (*requests.UpsertRequest, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var jsonBody interface{}
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		return nil, err
	}

	return &requests.UpsertRequest{
		Index: "recipes",
		Type:  "recipe",
		Body:  jsonBody,
	}, nil
}

func (impl *impl) respondBadRequest(w http.ResponseWriter, err error) {
	impl.deps.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
}

func (impl *impl) renderResponse(w http.ResponseWriter, status int, v interface{}) {
	impl.deps.Render.JSON(w, status, v)
	return
}
