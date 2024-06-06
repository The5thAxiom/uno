package uno

func (gs GameState) isValidCardToPlay(cardToPlay Card) bool {
	topCard := gs.ActivePile.Top()

	if cardToPlay.Type == WildCard {
		return true
	}

	if topCard.Type == WildCard {
		return cardToPlay.Color == gs.CardColorToPlay
	}

	if cardToPlay.Type == NumberCard && topCard.Type == NumberCard {
		return cardToPlay.Color == topCard.Color || cardToPlay.Number == topCard.Number
	}

	// if either is a special card
	return topCard.Color == cardToPlay.Color
}

func (gd *GameDirection) reverse() {
	if *gd == ClockwiseDirection {
		*gd = AntiClockwiseDirection
	} else {
		*gd = ClockwiseDirection
	}
}

func (gs *GameState) incrementTurn() {
	gs.CurrentPlayerIndex = (gs.CurrentPlayerIndex + 1) % len(gs.Players)
}

func (gs GameState) handleSpecialCardPlayed(cardPlayed Card) {
	if cardPlayed.Type == SpecialCard {
		if *cardPlayed.SpecialCardType == SkipCard {
			gs.incrementTurn()
		} else if *cardPlayed.SpecialCardType == ReverseCard {
			gs.Direction.reverse()
		} else { // draw 2
			// TODO: create draw 2 handler
		}
	}
}