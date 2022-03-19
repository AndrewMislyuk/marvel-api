package marvel

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type webLogger struct {
	logger io.Writer
	next   http.RoundTripper
}

func (w webLogger) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(w.logger, "[%s] %s %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL)
	return w.next.RoundTrip(r)
}
