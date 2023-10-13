package matchers

type Weibo struct {
	base
}

var (
	weiboName          = "Weibo"
	weiboVersionRegexp = []string{`(?i)(?:__weibo__)([\d.]+)`}
	weiboMatchRegexp   = []string{`__weibo__`}
)

func NewWeibo(userAgent string) *Weibo {
	return &Weibo{
		base: newBase(userAgent),
	}
}

func (w *Weibo) Name() string {
	return weiboName
}

func (w *Weibo) Version() string {
	return w.version(weiboVersionRegexp, 1)
}

func (w *Weibo) Match() bool {
	return w.match(weiboMatchRegexp)
}
