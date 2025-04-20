package server

type Request struct {
	Command string
	Args    []string
}

type Response struct {
	Body   string
	IsNull bool
}
