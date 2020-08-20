package bad_word

import "strings"

type Model struct {
	Word string `json:"word" `
}

const (
	//FlagReplace = 0 // 替换
	//FlagOnlyCheck = 1 // 只检查 是否符合
	DefaultReplace = "**"
)

func Filter(word string, ls []*Model, replaceSt string) (bool, string) {
	back := word
	for _, v := range ls {
		back = strings.ReplaceAll(back, v.Word, replaceSt)
	}
	println("word:", word)
	println("back:", back)
	if word != back {
		return false, back
	}
	return true, back
}
