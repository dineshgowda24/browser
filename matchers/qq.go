package matchers

type QQ struct {
	base
}

var (
	qqName          = "QQ Browser"
	qqVersionRegexp = []string{`(?i)(?:MQQBrowser|QQBrowserLite|QQBrowser|QQ)/([\d.]+)`}
	qqMatchRegexp   = []string{`QQ/|QQBrowser`}
)

func NewQQ(userAgent string) *QQ {
	return &QQ{
		base: newBase(userAgent),
	}
}

func (q *QQ) Name() string {
	return qqName
}

func (q *QQ) Version() string {
	return q.version(qqVersionRegexp, 1)
}

func (q *QQ) Match() bool {
	return q.match(qqMatchRegexp)
}
