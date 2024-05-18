package matchers

type QQ struct {
	p Parser
}

var (
	qqName          = "QQ Browser"
	qqVersionRegexp = []string{`(?i)(?:MQQBrowser|QQBrowserLite|QQBrowser|QQ)/([\d.]+)`}
	qqMatchRegexp   = []string{`QQ/|QQBrowser`}
)

func NewQQ(p Parser) *QQ {
	return &QQ{
		p: p,
	}
}

func (q *QQ) Name() string {
	return qqName
}

func (q *QQ) Version() string {
	return q.p.Version(qqVersionRegexp, 1)
}

func (q *QQ) Match() bool {
	return q.p.Match(qqMatchRegexp)
}
