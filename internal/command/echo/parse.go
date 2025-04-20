package echo

type Echo interface {
	Parse(buff []byte) (string, error)
}
