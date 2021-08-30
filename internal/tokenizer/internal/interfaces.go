package internal

type WordTokenizer interface{
	TokenizeWord(word string) []string
}