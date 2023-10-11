package browser

// BotMatcher is the interface for bot matchers
type BotMatcher interface {
	Matcher
}

var genericBotName = "Generic Bot"

// Bot is a struct that contains information about the user agent's bot
type Bot struct {
	userAgent  string     // user agent string
	matcher    BotMatcher // bot matcher detected
}

