package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "king":
		return 10
	case "queen":
		return 10
	case "jack":
		return 10
	case "ten":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	card1Value := ParseCard(card1)
	card2Value := ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)
	sumCard := card1Value + card2Value
	switch {
	case card1 == card2 && card1Value == 11:
		return "P"
	case sumCard == 21:
		if dealerCardValue < 10 {
			return "W"
		}
		return "S"
	case sumCard >= 17 && sumCard <= 20:
		return "S"
	case sumCard >= 12 && sumCard <= 16:
		if dealerCardValue < 7 {
			return "S"
		}
		return "H"
	default:
		return "H"
	}

}
