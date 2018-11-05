package mock

type Crypto struct{}

func (h *Crypto) Generate(s string) (string, error) {
	return s, nil
}

func (h *Crypto) Compare(hash string, s string) error {
	return nil
}
