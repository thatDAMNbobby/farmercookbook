package parameters

const (
	defaultPerPage = 10
	maxPerPage     = 50
)

type Params struct {
	Query   string `json:"query"`
	Page    int    `json:"page"`
	PerPage int    `json:"perPage"`
}
