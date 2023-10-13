package browser

import (
	"embed"
	"log"
	"reflect"
	"strings"

	"github.com/dineshgowda24/browser/bots"
	"gopkg.in/yaml.v2"
)

// Embed the assets/internal directory
//
//go:embed assets/internal
var fs embed.FS

// BotMatcher is the interface for bot matchers
type BotMatcher interface {
	Matcher
}

var genericBotName = "Generic Bot"

// Bot is a struct that contains information about the user agent's bot
type Bot struct {
	userAgent  string     // user agent string
	exceptions []string   // list of user agents that are exceptions
	matcher    BotMatcher // bot matcher detected
	registered bool       // indicates if the bot matcher has registered
}

// NewBot returns a new Bot
// It will return an error if the bot exceptions file cannot be read
// It will return an error if the bot exceptions file cannot be unmarshalled
func NewBot(userAgent string) (*Bot, error) {
	be, err := fs.ReadFile("assets/internal/bot_exceptions.yml")
	if err != nil {
		return nil, err
	}

	e := make([]string, 0)
	if err := yaml.Unmarshal(be, &e); err != nil {
		return nil, err
	}

	bot := &Bot{
		userAgent:  userAgent,
		exceptions: e,
	}
	bot.register()

	return bot, nil
}

// register registers the bot matcher
func (b *Bot) register() {
	if b.exception() {
		return
	}

	matchers := []Matcher{
		bots.NewEmpty(b.userAgent),
		bots.NewKnown(b.userAgent, getKnownBots()),
		bots.NewKeyword(b.userAgent),
	}

	for _, m := range matchers {
		if m.Match() {
			b.matcher = m
			break
		}
	}

	b.registered = true
}

// getKnownBots returns the known bots
// It will panic if the known bots file cannot be read
// TODO: return an error
func getKnownBots() map[string]string {
	kb, err := fs.ReadFile("assets/internal/known_bots.yml")
	if err != nil {
		log.Panicln(err)
	}

	b := make(map[string]string)
	if err := yaml.Unmarshal(kb, &b); err != nil {
		log.Panicln(err)
	}

	return b
}

// getMatcher returns the bot matcher detected from the user agent string
// It will register the bot matcher if it has not been registered
func (b *Bot) getMatcher() Matcher {
	if !b.registered {
		b.register()
	}

	return b.matcher
}

// exception returns true if the user agent is an exception
func (b *Bot) exception() bool {
	for _, e := range b.exceptions {
		if strings.Contains(b.userAgent, e) {
			return true
		}
	}

	return false
}

// IsBot returns true if the user agent is a bot
func (b *Bot) IsBot() bool {
	return b.getMatcher() != nil
}

// Name returns the name of the bot
func (b *Bot) Name() string {
	if b.matcher == nil {
		return ""
	}

	return b.matcher.Name()
}

// Why returns the reason why the user agent is a bot
// It will return an empty string if the user agent is not a bot
// It will return the type of the bot matcher detected
func (b *Bot) Why() string {
	if b.matcher == nil {
		return ""
	}

	return reflect.TypeOf(b.matcher).String()
}
