// Butterfly HTTP - The easiest HTTP client for GoLang from Butterfly.
// Copyright(c) 2022 Star Inc. All Rights Reserved.
// The software licensed under Mozilla Public License Version 2.0

package bHttp

type StatusCode int

func (s StatusCode) IsEqual(code StatusCode) bool {
	return s == code
}
