package devices

type Xbox360 struct {
	base
}

var (
	xbox360Name       = "Xbox 360"
	xbox360MatchRegex = []string{`(?i)Xbox`}
)

func NewXbox360(userAgent string) *Xbox360 {
	return &Xbox360{
		base{
			userAgent: userAgent,
		},
	}
}

func (x *Xbox360) Name() string {
	return xbox360Name
}

func (x *Xbox360) Match() bool {
	return x.match(xbox360MatchRegex)
}
