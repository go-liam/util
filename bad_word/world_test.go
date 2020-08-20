package bad_word

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestFilter1(t *testing.T) {
	word := "xxx非常好11-ddd棒极了22ddd-值得拥有333 -ee棒极了22-棒极了22- 非常好11 - 非常好11-值得拥有33-"
	ls := make([]*Model,0)
	ls= append(ls,&Model{Word:"非常好"})
	ls= append(ls,&Model{Word:"棒极了"})
	ls= append(ls,&Model{Word:"值得拥有"})
	f,got := Filter(word,ls,DefaultReplace)
	log.Printf("num=%+v\n",f)
	log.Printf("got=%+v\n",got)
	assert.Equal(t, false, f)
}

func TestFilter2(t *testing.T) {
	word := "xxx非常好11-ddd棒极了22ddd-值得拥有333 -ee棒极了22-棒极了22- 非常好11 - 非常好11-值得拥有33-"
	ls := make([]*Model,0)
	ls= append(ls,&Model{Word:"非常好A"})
	ls= append(ls,&Model{Word:"棒极了A"})
	ls= append(ls,&Model{Word:"值得拥有A"})
	f,got := Filter(word,ls,DefaultReplace)
	log.Printf("num=%+v\n",f)
	log.Printf("got=%+v\n",got)
	assert.Equal(t, true, f)
}