package domain

type Response struct {
	Error interface{} `json:",omitempty"`
	Body  interface{} `json:",omitempty"`
}
