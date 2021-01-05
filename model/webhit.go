package model

// WebHit stores a web request
type WebHit struct {
	// @TODO: Build this out!
	UserAgent    string `json:"useragent"`
	IP           string `json:"ip"`
	Port         string `json:"port"`
	ForwardedFor string `json:"forwardedfor"`
}
