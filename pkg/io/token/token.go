package token

import (
	"github.com/ont-bizsuite/ddxf-api-sdk/pkg/io"
	"github.com/ontio/ontology/common"
)

// VerifyTokenInput ...
// AddressSign
type VerifyTokenInput struct {
	// owner account
	Address common.Address
	// token id, generated by /ddxf/mp/buy_item or /ddxf/dtoken/generate
	TokenID string
	// # of token
	N int
}

// VerifyTokenOutput ...
// swagger:model VerifyTokenOutput
type VerifyTokenOutput struct {
	io.BaseResp
	// generated by /ddxf/dtoken/create_template
	TokenTemplateID string
	// token template for this token
	TokenTemplate TokenTemplate
	// whether the owner has so many tokens
	OK bool
}

// UseTokenInput ...
// AddressSign
type UseTokenInput struct {
	// owner account
	Address common.Address
	// token id, generated by /ddxf/mp/buy_item or /ddxf/dtoken/generate
	TokenID string
	// # of token
	N int
}

// UseTokenOutput ...
// swagger:model UseTokenOutput
type UseTokenOutput struct {
	io.BaseResp
	// send with ontology sdk
	Tx string
}

// TransferTokenInput ...
// AddressSign
type TransferTokenInput struct {
	// transfer from account
	From common.Address
	// transfer to account
	To common.Address
	// token id
	TokenID string
	// amount of token
	N int
}

// TransferTokenOutput ...
// swagger:model TransferTokenOutput
type TransferTokenOutput struct {
	io.BaseResp
	// send with ontology sdk
	Tx string
}

// SetTokenAgentInput ...
type SetTokenAgentInput struct {
	// owner account
	Owner common.Address
	// token id
	TokenID string
	// amount for agent, len(N) == len(Agents)
	N []int
	// agents for the token
	Agents []common.Address
}

// SetTokenAgentOutput ...
// swagger:model SetTokenAgentOutput
type SetTokenAgentOutput struct {
	io.BaseResp
	// send with ontology sdk
	Tx string
}

// AddTokenAgentInput ...
type AddTokenAgentInput struct {
	// owner account
	Owner common.Address
	// token id
	TokenID string
	// additional amount for agent, len(N) == len(Agents)
	N []int
	// agents for the token
	Agents []common.Address
}

// AddTokenAgentOutput ...
// swagger:model AddTokenAgentOutput
type AddTokenAgentOutput struct {
	io.BaseResp
	// send with ontology sdk
	Tx string
}

// RemoveTokenAgentInput ...
type RemoveTokenAgentInput struct {
	// owner account
	Owner common.Address
	// token id
	TokenID string
	// agents for the token
	Agents []common.Address
}

// RemoveTokenAgentOutput ...
// swagger:model RemoveTokenAgentOutput
type RemoveTokenAgentOutput struct {
	io.BaseResp
	// send with ontology sdk
	Tx string
}

// ClearTokenAgentInput ...
type ClearTokenAgentInput struct {
	// owner account
	Owner common.Address
	// token id
	TokenID string
	// agents to clear
	Agents []common.Address
}

// ClearTokenAgentOutput ...
// swagger:model ClearTokenAgentOutput
type ClearTokenAgentOutput struct {
	io.BaseResp
	// send with ontology sdk
	Tx string
}

// GetTokenInput ...
type GetTokenInput struct {
	// owner account
	Owner common.Address
	// token id
	TokenID string
}

// GetTokenOutput ...
// swagger:model GetTokenOutput
type GetTokenOutput struct {
	io.BaseResp
	// token template id for this token, generated by /ddxf/dtoken/create_template
	TokenTemplateID string
	// token template for this token
	TokenTemplate TokenTemplate
	// # of token owned
	N int
}

// TokenTemplate ...
type TokenTemplate struct {
	// data id is generated by /ddxf/data_meta/create
	DataID string
	// token meta hash for token meta
	TokenMetaHash string
	// endpoint for meta service
	Endpoint string
}

// CreateTokenTemplateInput ...
// Note: update token_template_id when tx is published
type CreateTokenTemplateInput struct {
	// owner account
	Address common.Address
	// template info
	Template TokenTemplate
}

// CreateTokenTemplateOutput ...
// swagger:model CreateTokenTemplateOutput
type CreateTokenTemplateOutput struct {
	io.BaseResp
	// send with ontology sdk
	Tx string
}

// UpdateTokenTemplateInput ...
type UpdateTokenTemplateInput struct {
	// monitored above
	TokenTemplateID string
	// template info
	Template TokenTemplate
}

// UpdateTokenTemplateOutput ...
// swagger:model UpdateTokenTemplateOutput
type UpdateTokenTemplateOutput struct {
	io.BaseResp
	// send with ontology sdk
	Tx string
}

// RemoveTokenTemplateInput ...
type RemoveTokenTemplateInput struct {
	// monitored above
	TokenTemplateID string
}

// RemoveTokenTemplateOutput ...
// swagger:model RemoveTokenTemplateOutput
type RemoveTokenTemplateOutput struct {
	io.BaseResp
	// send with ontology sdk
	Tx string
}

// GetTokenTemplateInput ...
type GetTokenTemplateInput struct {
	// template id, generated by /ddxf/dtoken/create_template
	TokenTemplateID string
}

// GetTokenTemplateOutput ...
// swagger:model GetTokenTemplateOutput
type GetTokenTemplateOutput struct {
	io.BaseResp
	// authenticated accounts
	Auths []common.Address
	// template info
	Template TokenTemplate
}

// SetTokenTemplateAuthInput ...
type SetTokenTemplateAuthInput struct {
	// template id, generated by /ddxf/dtoken/create_template
	TokenTemplateID string
	// owner account
	Owner common.Address
	// authenticated accounts
	Auths []common.Address
}

// SetTokenTemplateAuthOutput ...
// swagger:model SetTokenTemplateAuthOutput
type SetTokenTemplateAuthOutput struct {
	io.BaseResp
	// send with ontology sdk
	Tx string
}

// AddTokenTemplateAuthInput ...
type AddTokenTemplateAuthInput struct {
	// template id, generated by /ddxf/dtoken/create_template
	TokenTemplateID string
	// owner account
	Owner common.Address
	// authenticated accounts
	Auths []common.Address
}

// AddTokenTemplateAuthOutput ...
// swagger:model AddTokenTemplateAuthOutput
type AddTokenTemplateAuthOutput struct {
	io.BaseResp
	// send with ontology sdk
	Tx string
}

// RemoveTokenTemplateAuthInput ...
type RemoveTokenTemplateAuthInput struct {
	// template id, generated by /ddxf/dtoken/create_template
	TokenTemplateID string
	// owner account
	Owner common.Address
	// authenticated accounts
	Auths []common.Address
}

// RemoveTokenTemplateAuthOutput ...
// swagger:model RemoveTokenTemplateAuthOutput
type RemoveTokenTemplateAuthOutput struct {
	io.BaseResp
	// send with ontology sdk
	Tx string
}

// ClearTokenTemplateAuthInput ...
type ClearTokenTemplateAuthInput struct {
	// template id, generated by /ddxf/dtoken/create_template
	TokenTemplateID string
	// owner account
	Owner common.Address
}

// ClearTokenTemplateAuthOutput ...
// swagger:model ClearTokenTemplateAuthOutput
type ClearTokenTemplateAuthOutput struct {
	io.BaseResp
	// send with ontology sdk
	Tx string
}

// GenerateTokenInput ...
type GenerateTokenInput struct {
	// includes owner
	Auth common.Address
	// template id, generated by /ddxf/dtoken/create_template
	TokenTemplateID string
	// to account
	To common.Address
	// amount of token
	N int
}


// GenerateTokenOutput ...
// swagger:model GenerateTokenOutput
type GenerateTokenOutput struct {
	io.BaseResp
	// send with ontology sdk
	Tx string
}
