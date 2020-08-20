package bad_word

import "strings"

type Model struct {
	Word string `json:"word" `
}

const (
	DefaultReplace = "**"
)

func Filter(word string, ls []*Model, replaceSt string) (bool, string) {
	back := word
	for _, v := range ls {
		back = strings.ReplaceAll(back, v.Word, replaceSt)
	}
	if word != back {
		return false, back
	}
	return true, back
}
