package bots

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

// Name returns an empty string.
func (e *Empty) Name() string {
	return ""
}

// Match returns true if the user agent is empty.
func (e *Empty) Match() bool {
	return e.userAgent == ""
}
