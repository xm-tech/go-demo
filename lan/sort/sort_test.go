package sort

import (
	"log"
	"sort"
	"testing"
)

type Card struct {
	Score int
}

func (c *Card) ScoreVal() int {
	return c.Score
}

func TestSortTable(t *testing.T) {
	cards := []*Card{}
	cards = append(cards, &Card{Score: 5})
	cards = append(cards, &Card{Score: 7})
	cards = append(cards, &Card{Score: 6})

	log.Printf("cards:%v,%v,%v", cards[0], cards[1], cards[2])
	sort.SliceStable(cards, func(i, j int) bool {
		// assert i < j , 从小到大牌
		return cards[i].Score < cards[j].Score
	})

	log.Printf("cards:%v,%v,%v", cards[0], cards[1], cards[2])

	val := []int{2, 1, 3}
	sort.SliceStable(val, func(i, j int) bool {
		return val[i] < val[j]
	})

	log.Printf("val:%v", val)

	sort.SliceStable(val, func(i, j int) bool {
		return val[i] > val[j]
	})
	log.Printf("val:%v", val)

}
