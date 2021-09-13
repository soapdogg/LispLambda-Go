package internal

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o fakes/fake_word_tokenizer.go . WordTokenizer
type WordTokenizer interface{
	TokenizeWord(word string) []string
}