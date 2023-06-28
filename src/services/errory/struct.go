package errory

type ErrorX struct {
	StatusCode int
	Message    string
	ClientIP   string
	Method     string
	URL        string
	ErrorRoot  error
}

func (s *ErrorX) Error() string {
	if s.Message == `` {
		return `Unknown`
	}

	return s.Message
}
