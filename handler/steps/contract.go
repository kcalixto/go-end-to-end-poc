package steps

type Steps interface {
	Execute(body string) (string, error)
}