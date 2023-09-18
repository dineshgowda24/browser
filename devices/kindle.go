package devices

type Kindle struct {
	base
}

var (
	kindleName       = "Kindle"
	kindleMatchRegex = []string{`Kindle`}
)

func NewKindle(userAgent string) *Kindle {
	return &Kindle{
		base{
			userAgent: userAgent,
		},
	}
}

func (k *Kindle) Name() string {
	return kindleName
}

func (k *Kindle) Match() bool {
	return k.match(kindleMatchRegex)
}
