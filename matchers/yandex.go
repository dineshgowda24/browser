package matchers

type Yandex struct {
	p Parser
}

var (
	yandexName          = "Yandex"
	yandexVersionRegexp = []string{`YaBrowser/([\d.]+)`}
	yandexMatchRegexp   = []string{`YaBrowser`}
)

func NewYandex(p Parser) *Yandex {
	return &Yandex{
		p: p,
	}
}

func (y *Yandex) Name() string {
	return yandexName
}

func (y *Yandex) Version() string {
	return y.p.Version(yandexVersionRegexp, 1)
}

func (y *Yandex) Match() bool {
	return y.p.Match(yandexMatchRegexp)
}
