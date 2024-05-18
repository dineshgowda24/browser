package matchers

type Edge struct {
	p  Parser
	ie *InternetExplorer
}

var (
	edgeName         = "Microsoft Edge"
	edgeMatchRegex   = []string{`((?:Edge|Edg|EdgiOS|EdgA)/[\d.]+|Trident/8)`}
	edgeVersionRegex = []string{`(?:Edge|Edg|EdgiOS|EdgA)/([\d.]+)`}
)

func NewEdge(p Parser) *Edge {
	return &Edge{
		p:  p,
		ie: NewInternetExplorer(p),
	}
}

func (e *Edge) Name() string {
	return edgeName
}

func (e *Edge) Version() string {
	v := e.p.Version(edgeVersionRegex, 1)
	if v != "0.0" {
		return v
	}

	return e.ie.Version()
}

func (e *Edge) Match() bool {
	return e.p.Match(edgeMatchRegex)
}
