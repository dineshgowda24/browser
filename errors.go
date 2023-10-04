package browser

import "fmt"

var (
	ErrNoUserAgent           = fmt.Errorf("no user agent provided")
	ErrUserAgentSizeExceeded = fmt.Errorf("user agent size exceeds limit of %d", userAgentSizeLimit)
)
