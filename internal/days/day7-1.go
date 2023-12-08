package day

import (
	"sort"
)

type Card struct {
	Rank string
}

type Hand []Card

func DaySevenPartOne(cardsBids map[string]int) int {
	cardsRank := make(map[string]int)

	for cards := range cardsBids {
		var h Hand
		for i := range cards {
			h = append(h, Card{string(cards[i])})
		}
		score := scoreHand1(h)

		cardsRank[cards] = score
	}

	keys := make([]string, 0, len(cardsRank))

	for key := range cardsRank {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		if cardsRank[keys[i]] == cardsRank[keys[j]] {
			for k := range keys[i] {
				card_i := string(keys[i][k])
				card_j := string(keys[j][k])

				if rankValue1(card_i) > rankValue1(card_j) {
					return false
				}
				if rankValue1(card_j) > rankValue1(card_i) {
					return true
				}
			}
		}
		return cardsRank[keys[i]] < cardsRank[keys[j]]
	})

	var part1 int
	for i := range keys {
		part1 += cardsBids[keys[i]] * (i + 1)
	}

	return part1
}

func scoreHand1(hand Hand) int {
	sort.Slice(hand, func(i, j int) bool {
		return rankValue1(hand[i].Rank) < rankValue1(hand[j].Rank)
	})

	// Check for specific hand types
	if isFiveOfAKind1(hand) {
		return 7
	} else if isFourOfAKind1(hand) {
		return 6
	} else if isFullHouse1(hand) {
		return 5
	} else if isThreeOfAKind1(hand) {
		return 4
	} else if isTwoPair1(hand) {
		return 3
	} else if isOnePair1(hand) {
		return 2
	}

	return 1
}

func rankValue1(rank string) int {
	rankValues := map[string]int{
		"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
		"T": 10, "J": 11, "Q": 12, "K": 13, "A": 14,
	}

	return rankValues[rank]
}

func isFiveOfAKind1(hand Hand) bool {
	countMap := handMap(hand)

	for _, count := range countMap {
		if count == 5 {
			return true
		}
	}

	return false
}

func isFourOfAKind1(hand Hand) bool {
	countMap := handMap(hand)

	for _, count := range countMap {
		if count == 4 {
			return true
		}
	}

	return false
}

func isFullHouse1(hand Hand) bool {
	countMap := handMap(hand)

	hasThreeOfAKind := false
	hasPair := false

	for _, count := range countMap {
		if count == 3 {
			hasThreeOfAKind = true
		} else if count == 2 {
			hasPair = true
		}
	}

	return hasThreeOfAKind && hasPair
}

func isThreeOfAKind1(hand Hand) bool {
	countMap := handMap(hand)

	for _, count := range countMap {
		if count == 3 {
			return true
		}
	}

	return false
}

func isTwoPair1(hand Hand) bool {
	countMap := handMap(hand)

	pairCount := 0
	for _, count := range countMap {
		if count == 2 {
			pairCount++
		}
	}

	return pairCount == 2
}

func isOnePair1(hand Hand) bool {
	countMap := handMap(hand)

	for _, count := range countMap {
		if count == 2 {
			return true
		}
	}

	return false
}
