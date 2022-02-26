package render

import (
	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
	"net/http"
)

type Deps struct {
	Render *render.Render
}

type Render interface {
	JSON(w http.ResponseWriter, status int, v interface{})
}

type impl struct {
	deps *Deps
}

func New(deps *Deps) Render {
	return &impl{deps: deps}
}

func (impl *impl) JSON(w http.ResponseWriter, status int, v interface{}) {
	err := impl.deps.Render.JSON(w, status, v)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
