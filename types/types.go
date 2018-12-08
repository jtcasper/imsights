package types

type Class struct {
	Name      string
	Functions []Function
}

type Function struct {
	Name  string
	Calls []Call
}

type Call struct {
	Name string
}
