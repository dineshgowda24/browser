package matchers

type Alipay struct {
	p Parser
}

var (
	aliPayName          = "Alipay"
	aliPayVersionRegexp = []string{`AlipayClient(?:HK)?/([\d.]+)`}
	aliPayMatchRegex    = []string{`(?i)Alipay`}
)

func NewAlipay(p Parser) *Alipay {
	return &Alipay{
		p: p,
	}
}

func (a *Alipay) Name() string {
	return aliPayName
}

func (a *Alipay) Version() string {
	return a.p.Version(aliPayVersionRegexp, 1)
}

func (a *Alipay) Match() bool {
	return a.p.Match(aliPayMatchRegex)
}
