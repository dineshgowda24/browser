package matchers

type Alipay struct {
	base
}

var (
	aliPayName          = "Alipay"
	aliPayVersionRegexp = []string{`AlipayClient(?:HK)?/([\d.]+)`}
	aliPayMatchRegex    = []string{`(?i)Alipay`}
)

func NewAlipay(userAgent string) *Alipay {
	return &Alipay{
		base: newBase(userAgent),
	}
}

func (a *Alipay) Name() string {
	return aliPayName
}

func (a *Alipay) Version() string {
	return a.version(aliPayVersionRegexp, 1)
}

func (a *Alipay) Match() bool {
	return a.match(aliPayMatchRegex)
}
