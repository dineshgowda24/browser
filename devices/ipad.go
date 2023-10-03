package devices

type Ipad struct {
	base
}

var (
	iPadName       = "iPad"
	iPadMatchRegex = []string{`iPad`}
)

func NewIpad(userAgent string) *Ipad {
	return &Ipad{
		base{
			userAgent: userAgent,
		},
	}
}

func (i *Ipad) Name() string {
	return iPadName
}

func (i *Ipad) Match() bool {
	return i.match(iPadMatchRegex)
}
