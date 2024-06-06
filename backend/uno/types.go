package uno

import (
	"backend/utils"
)

type CardType string

const (
	NumberCard  CardType = "NumberCard"
	SpecialCard CardType = "SpecialCard"
	WildCard    CardType = "WildCard"
)

type SpecialCardType string

const (
	Draw2Card   SpecialCardType = "Draw2Card"
	ReverseCard SpecialCardType = "ReverseCard"
	SkipCard    SpecialCardType = "SkipCard"
)

type WildCardType string

const (
	PlainWildCard WildCardType = "PlainWildCard"
	Draw4WildCard WildCardType = "Draw4WildCard"
)

type CardColor string

const (
	Red    CardColor = "Red"
	Blue   CardColor = "Red"
	Yellow CardColor = "Red"
	Green  CardColor = "Red"
	Black  CardColor = "Black"
)

type Card struct {
	Name            string
	ImageUrl        string
	Type            CardType
	Color           CardColor
	Number          *int
	SpecialCardType *SpecialCardType
	WildCardType    *WildCardType
}

type User struct {
	Name          string
	Token         string
	profilePicUrl string
}

type GameType string

const (
	VanillaGame GameType = "VanillaGame"
)

type GameDirection string

const (
	ClockwiseDirection GameDirection = "ClockwiseDirection"
	AntiClockwiseDirection GameDirection = "AntiClockwise Direction"
)

type GamePlayer struct {
	User User
	Hand []Card
}

type GameState struct {
	GameType GameType
	DrawPile utils.Stack[Card]
	ActivePile utils.Stack[Card]
	Players []GamePlayer
	CurrentPlayerIndex int
	CardColorToPlay CardColor
	Direction GameDirection
}