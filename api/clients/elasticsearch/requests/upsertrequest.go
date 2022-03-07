package requests

type UpsertRequest struct {
	Index string
	Type  string
	Id    *string
	Body  interface{}
}
