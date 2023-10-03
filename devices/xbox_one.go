package devices

type XboxOne struct {
	base
}

var (
	xboxOneName       = "Xbox One"
	xboxOneMatchRegex = []string{`(?i)Xbox One`}
)

func NewXboxOne(userAgent string) *XboxOne {
	return &XboxOne{
		base{
			userAgent: userAgent,
		},
	}
}

func (x *XboxOne) Name() string {
	return xboxOneName
}

func (x *XboxOne) Match() bool {
	return x.match(xboxOneMatchRegex)
}
