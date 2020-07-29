package token

import (
	"github.com/ont-bizsuite/ddxf-sdk/market_place_contract"
	"github.com/klaytn/klaytn/common"
)

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

type UpdateTokenTemplateEvent struct {
	TT              *market_place_contract.TokenTemplate
	TokenTemplateID string
}

type RemoveTokenTemplateEvent struct {
	TokenTemplateID string
}

type AuthorizeTokenTemplateEvent struct {
	TokenTemplateID string
	Addrs []common.Address
}

type GenerateTokenEvent struct {
	Account         string
	TokenTemplateId string
	N               int
	TokenId         string
}
