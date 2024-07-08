package types

type ValidationError struct {
	Errors []Error `json:"errors"`
}

type Error struct {
	NameSpace string `json:"namespace"`
	Field     string `json:"field"`
	Tag       string `json:"tag"`
	Kind      string `json:"kind"`
	Type      string `json:"type"`
	Value     any    `json:"value"`
}
