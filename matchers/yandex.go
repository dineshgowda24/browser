package matchers

type Yandex struct {
	base
}

var (
	yandexName          = "Yandex"
	yandexVersionRegexp = []string{`YaBrowser/([\d.]+)`}
	yandexMatchRegexp   = []string{`YaBrowser`}
)

func NewYandex(userAgent string) *Yandex {
	return &Yandex{
		base: newBase(userAgent),
	}
}

func (y *Yandex) Name() string {
	return yandexName
}

func (y *Yandex) Version() string {
	return y.version(yandexVersionRegexp, 1)
}

func (y *Yandex) Match() bool {
	return y.match(yandexMatchRegexp)
}
