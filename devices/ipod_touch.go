package devices

type IpodTouch struct {
	p Parser
}

var (
	ipodTouchName       = "iPod Touch"
	ipodTouchMatchRegex = []string{`(?i)iPod`}
)

func NewIpodTouch(p Parser) *IpodTouch {
	return &IpodTouch{
		p: p,
	}
}

func (i *IpodTouch) Name() string {
	return ipodTouchName
}

func (i *IpodTouch) Match() bool {
	return i.p.Match(ipodTouchMatchRegex)
}
