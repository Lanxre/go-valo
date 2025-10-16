package types

type Response[T any] struct {
	Data   T      `json:"data"`
	Status int    `json:"status"`
	Error  string `json:"error,omitempty"`
}
