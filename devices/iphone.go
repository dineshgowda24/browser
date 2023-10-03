package devices

type Iphone struct {
	base
}

var (
	iPhoneName       = "iPhone"
	iPhoneMatchRegex = []string{`iPhone`}
)

func NewIphone(userAgent string) *Iphone {
	return &Iphone{
		base{
			userAgent: userAgent,
		},
	}
}

func (i *Iphone) Name() string {
	return iPhoneName
}

func (i *Iphone) Match() bool {
	return i.match(iPhoneMatchRegex)
}
