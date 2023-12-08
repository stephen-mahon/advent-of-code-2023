package day

import (
	"sort"
)

func DaySevenPartTwo(cardsBids map[string]int) int {

	cardsRank := make(map[string]int)

	for cards := range cardsBids {
		var h Hand
		for i := range cards {
			h = append(h, Card{string(cards[i])})
		}
		score := scoreHand(h)

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

				if rankValue(card_i) > rankValue(card_j) {
					return false
				}
				if rankValue(card_j) > rankValue(card_i) {
					return true
				}
			}
		}
		return cardsRank[keys[i]] < cardsRank[keys[j]]
	})

	var part2 int
	for i := range keys {
		//fmt.Println(keys[i], cardsBids[keys[i]])
		part2 += cardsBids[keys[i]] * (i + 1)
	}

	return part2
}

func scoreHand(hand Hand) int {
	// Order the hand for lowest rank (J) to highest (A)
	sort.Slice(hand, func(i, j int) bool {
		return rankValue(hand[i].Rank) < rankValue(hand[j].Rank)
	})

	// Check for specific hand types
	if isFiveOfAKind(hand) {
		return 7
	} else if isFourOfAKind(hand) {
		return 6
	} else if isFullHouse(hand) {
		return 5
	} else if isThreeOfAKind(hand) {
		return 4
	} else if isTwoPair(hand) {
		return 3
	} else if isOnePair(hand) {
		return 2
	}
	// 1 is for high card
	return 1
}

func rankValue(rank string) int {
	// Map the Card Value to its Rank. Returns the rank
	rankValues := map[string]int{
		"J": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
		"T": 10, "Q": 11, "K": 12, "A": 13,
	}

	return rankValues[rank]
}

func handMap(hand Hand) map[string]int {
	countMap := make(map[string]int)
	for _, card := range hand {
		countMap[card.Rank]++
	}

	return countMap
}

func isFiveOfAKind(hand Hand) bool {
	// I used the handMap function extensively for counting the occurence of each card.
	countMap := handMap(hand)

	for _, count := range countMap {
		if count == 5 || (count+countMap["J"] == 5) {
			return true
		}
	}

	return false
}

func isFourOfAKind(hand Hand) bool {
	countMap := handMap(hand)

	for k, count := range countMap {
		if k == "J" {
			continue
		}
		if count == 4 || (count+countMap["J"] == 4) {
			return true
		}
	}

	return false
}

func isFullHouse(hand Hand) bool {
	countMap := handMap(hand)

	hasThreeOfAKind := false
	hasPair := false

	if countMap["J"] == 1 && isTwoPair(hand[1:]) {
		// A quick test to remove any edge cases. Two pairs and a joker is a full house
		return true
	}

	for k, count := range countMap {
		// Here I do not compare Jokers for three of a kind or two of a kind.
		if k == "J" {
			continue
		}
		if count == 3 || (count+countMap["J"] == 3) {
			hasThreeOfAKind = true
		} else if count == 2 || (count+countMap["J"] == 2) {
			hasPair = true
		}
	}

	return hasThreeOfAKind && hasPair
}

func isThreeOfAKind(hand Hand) bool {
	countMap := handMap(hand)

	for k, count := range countMap {
		if k == "J" {
			continue
		}
		if count == 3 || (count+countMap["J"] == 3) {
			return true
		}
	}

	return false
}

func isTwoPair(hand Hand) bool {
	countMap := handMap(hand)

	pairCount := 0
	for k, count := range countMap {
		if k == "J" {
			continue
		}
		if count == 2 || (count+countMap["J"] == 2) {
			pairCount++
		}
	}
	// Only return true if the number of pairs is 2.
	return pairCount == 2
}

func isOnePair(hand Hand) bool {
	countMap := handMap(hand)

	for _, count := range countMap {
		if count == 2 || (count+countMap["J"] == 2) {
			return true
		}
	}

	return false
}
