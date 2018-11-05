package root

type Crypto interface {
	Generate(s string) (string, error)
	Compare(hash string, s string) error
}
