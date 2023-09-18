package devices

type IpodTouch struct {
	base
}

var (
	ipodTouchName       = "iPod Touch"
	ipodTouchMatchRegex = []string{`(?i)iPod`}
)

func NewIpodTouch(userAgent string) *IpodTouch {
	return &IpodTouch{
		base{
			userAgent: userAgent,
		},
	}
}

func (i *IpodTouch) Name() string {
	return ipodTouchName
}

func (i *IpodTouch) Match() bool {
	return i.match(ipodTouchMatchRegex)
}
