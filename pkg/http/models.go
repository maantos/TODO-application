package http

type ContextBodyKey struct{}

type Response struct {
	Code int         `json:"-"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}
