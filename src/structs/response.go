package structs

type Response struct {
	Data string `json:"data"`
	Tag  string `json:"tag,omitempty"`
}
