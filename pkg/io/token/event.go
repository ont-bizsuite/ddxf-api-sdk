package token

// UseTokenEvent ...
type UseTokenEvent struct {
	Acct string
	TokenId string
	N int
}

// UseTokenEvent ...
type DeleteTokenEvent struct {
	Acct string
	TokenId string
}
// UseTokenEvent ...
type UseTokenByAgentEvent struct {
	Acct string
	TokenId string
	N int
}