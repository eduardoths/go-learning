package structs

type Response struct {
	Data interface{} `json:"data"`
	Tag  string `json:"tag,omitempty"`
}
