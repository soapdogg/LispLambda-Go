package tokenizer

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o fakes/fake_tokenizer.go . Tokenizer
type Tokenizer interface {
	Tokenize(input string) []string
}