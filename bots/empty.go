package bots

var genericBot = "Generic Bot"

// Empty is a bot that matches when the user agent is empty.
// An empty user agent is treated as a bot.
type Empty struct {
	userAgent string
}

func NewEmpty(userAgent string) *Empty {
	return &Empty{
		userAgent: userAgent,
	}
}

// Name returns genericBot.
func (e *Empty) Name() string {
	return genericBot
}

// Match returns true if the user agent is empty.
func (e *Empty) Match() bool {
	return e.userAgent == ""
}
