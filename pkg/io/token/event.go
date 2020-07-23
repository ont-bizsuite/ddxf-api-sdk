package token

import "github.com/ont-bizsuite/ddxf-sdk/market_place_contract"

// UseTokenEvent ...
type UseTokenEvent struct {
	Acct    string
	TokenId string
	N       int
}

// UseTokenEvent ...
type DeleteTokenEvent struct {
	Acct    string
	TokenId string
}

// UseTokenEvent ...
type UseTokenByAgentEvent struct {
	Acct    string
	TokenId string
	N       int
}

// CreateTokenTemplateEvent ...
type CreateTokenTemplateEvent struct {
	Creator         string
	TT              *market_place_contract.TokenTemplate
	TokenTemplateID string
}

type GenerateTokenEvent struct {
	Account         string
	TokenTemplateId string
	N               int
	TokenId         string
}
