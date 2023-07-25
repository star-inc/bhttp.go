// bhttp.go
// (c) 2023 Star Inc.

package bHttp

type StatusCode int

func (s StatusCode) IsEqual(code StatusCode) bool {
	return s == code
}
