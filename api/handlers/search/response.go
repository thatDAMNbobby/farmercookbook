package search

import "github.com/thatDAMNbobby/farmercookbook/models"

type Pagination struct {
	TotalResults int `json:"totalResults"`
	CurrentPage  int `json:"page"`
	TotalPages   int `json:"limit"`
	NextPage     int `json:"nextPage"`
	PreviousPage int `json:"previousPage"`
	PerPage      int `json:"perPage"`
	Begin        int `json:"begin"`
	End          int `json:"end"`
}

type Response struct {
	Pagination Pagination      `json:"pagination"`
	Results    []models.Recipe `json:"results"`
}
