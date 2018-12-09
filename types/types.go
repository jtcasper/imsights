package types

type Class struct {
	Name      string     `json:"name"`
	Functions []Function `json:"functions"`
}

type Function struct {
	Name  string `json:"name"`
	Calls []Call `json:"calls"`
}

type Call struct {
	Name string `json:"name"`
}
