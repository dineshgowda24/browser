package devices

type BlackberryPlaybook struct {
	base
}

var (
	bbPlaybookName       = "BlackBerry Playbook"
	bbPlaybookMatchRegex = []string{`PlayBook.*?RIM Tablet`}
)

func NewBlackberryPlaybook(userAgent string) *BlackberryPlaybook {
	return &BlackberryPlaybook{
		base{
			userAgent: userAgent,
		},
	}
}

func (b *BlackberryPlaybook) Name() string {
	return bbPlaybookName
}

func (b *BlackberryPlaybook) Match() bool {
	return b.match(bbPlaybookMatchRegex)
}
