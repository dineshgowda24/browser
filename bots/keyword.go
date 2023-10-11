package bots

import "regexp"

// Keyword is a bot that matches when the user agent contains a keyword.
type Keyword struct {
	userAgent string
}

var (
	keywordMatchRegex = `(?i)(?:crawl|fetch|search|monitoring|spider|bot)`
)

// NewKeyword returns a new Keyword bot.
func NewKeyword(userAgent string) *Keyword {
	return &Keyword{
		userAgent: userAgent,
	}
}

// Name returns the keyword that matched.
func (k *Keyword) Name() string {
	return regexp.MustCompile(keywordMatchRegex).FindString(k.userAgent)
}

// Match returns true if the user agent contains a keyword.
func (k *Keyword) Match() bool {
	return regexp.MustCompile(keywordMatchRegex).MatchString(k.userAgent)
}
