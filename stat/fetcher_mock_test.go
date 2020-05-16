package stat

import "errors"

type FetcherMock struct {
	Err        error
	Res        string
	PassedData string
}

func (f FetcherMock) GetBodyAsString(url string) (string, error) {
	f.PassedData = url
	return f.Res, f.Err
}

func (f FetcherMock) ReturnsSuccess(data string) {
	f.Res = data
}

func (f FetcherMock) ReturnsError() {
	f.Err = errors.New("some err")
}
