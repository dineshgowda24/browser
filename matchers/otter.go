
package matchers

type Otter struct {
	base
}

var (
	otterName          = "Otter"
	otterVersionRegexp = []string{`Otter/([\d.]+)`}
	otterMatchRegexp   = []string{`Otter`}
)

func NewOtter(userAgent string) *Otter {
	return &Otter{
		base: newBase(userAgent),
	}
}

func (o *Otter) Name() string {
	return otterName
}

func (o *Otter) Version() string {
	return o.version(otterVersionRegexp, 1)
}

func (o *Otter) Match() bool {
	return o.match(otterMatchRegexp)
}
