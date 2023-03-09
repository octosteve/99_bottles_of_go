package bottles

import (
	"strings"

	"github.com/octosteve/99_bottles_of_go/bottle_verse"
)

type Bottles struct {
}

func NewBottles() Bottles {
	return Bottles{}
}

func (b Bottles) Verses(start, finish int) (verses string) {
	for i := start; i >= finish; i-- {
		verse := bottle_verse.NewBottleVerse(i).Lyrics()
		verses += verse + "\n"
	}
	verses = strings.Trim(verses, "\n")
	return 
}

func (b Bottles) Song() string {
	return b.Verses(99, 0)
}
