package matchers

type Edge struct {
	InternetExplorer
}

var (
	edgeName         = "Microsoft Edge"
	edgeMatchRegex   = []string{`((?:Edge|Edg|EdgiOS|EdgA)/[\d.]+|Trident/8)`}
	edgeVersionRegex = []string{`(?:Edge|Edg|EdgiOS|EdgA)/([\d.]+)`}
)

func NewEdge(userAgent string) *Edge {
	return &Edge{
		InternetExplorer: InternetExplorer{
			base: base{
				userAgent: userAgent,
			},
		},
	}
}

func (e *Edge) Name() string {
	return edgeName
}

func (e *Edge) Version() string {
	v := e.version(edgeVersionRegex, 1)
	if v != "0.0" {
		return v
	}

	return e.InternetExplorer.Version()
}

func (e *Edge) Match() bool {
	return e.match(edgeMatchRegex)
}
