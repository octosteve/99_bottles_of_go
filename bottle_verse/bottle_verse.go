package bottle_verse

import "fmt"

type BottleVerse struct {
	Number int
}

func NewBottleVerse(number int) BottleVerse {
	return BottleVerse{number}
}

func (verse BottleVerse) Lyrics() string {
	switch verse.Number {
	case 0:
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"
	case 1:
		return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n"
	case 2:
		return "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n"
	default:
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", verse.Number, verse.Number, verse.Number-1)
	}
}
