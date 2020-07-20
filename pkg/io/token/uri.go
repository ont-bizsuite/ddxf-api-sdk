package token

const (
	// VerifyTokenURI ...
	VerifyTokenURI = "/ddxf/dtoken/verify"
	// swagger:route POST /ddxf/dtoken/verify SwaggerVerifyTokenInput
	// verify token
	//     Responses:
	//       200: body:VerifyTokenOutput

	// UseTokenURI ...
	UseTokenURI = "/ddxf/dtoken/use"
	// swagger:route POST /ddxf/dtoken/use SwaggerUseTokenInput
	// use token
	//     Responses:
	//       200: body:UseTokenOutput

	// TransferTokenURI ...
	TransferTokenURI = "/ddxf/dtoken/transfer"
	// swagger:route POST /ddxf/dtoken/transfer  SwaggerTransferTokenInput
	// transfer token
	//     Responses:
	//       200: body:TransferTokenOutput

	// SetTokenAgentURI ...
	SetTokenAgentURI = "/ddxf/dtoken/set_agent"
	// swagger:route POST /ddxf/dtoken/set_agent  SwaggerSetTokenAgentInput
	// set token agent
	//     Responses:
	//       200: body:SetTokenAgentOutput

	// AddTokenAgentURI ...
	AddTokenAgentURI = "/ddxf/dtoken/add_agent"
	// swagger:route POST /ddxf/dtoken/add_agent  SwaggerAddTokenAgentInput
	// add token agent
	//     Responses:
	//       200: body:AddTokenAgentOutput

	// RemoveTokenAgentURI ...
	RemoveTokenAgentURI = "/ddxf/dtoken/remove_agent"
	// swagger:route POST /ddxf/dtoken/remove_agent  SwaggerRemoveTokenAgentInput
	// remove token agent
	//     Responses:
	//       200: body:RemoveTokenAgentOutput

	// ClearTokenAgentURI ...
	ClearTokenAgentURI = "/ddxf/dtoken/clear_agent"
	// swagger:route POST /ddxf/dtoken/clear_agent  SwaggerClearTokenAgentInput
	// clear token agent
	//     Responses:
	//       200: body:ClearTokenAgentOutput

	// GetTokenURI ...
	GetTokenURI = "/ddxf/dtoken/get"
	// swagger:route POST /ddxf/dtoken/get  SwaggerGetTokenInput
	// get token
	//     Responses:
	//       200: body:GetTokenOutput

	// CreateTokenTemplateURI ...
	CreateTokenTemplateURI = "/ddxf/dtoken/create_template"
	// swagger:route POST /ddxf/dtoken/create_template  SwaggerCreateTokenTemplateInput
	// create token template
	//     Responses:
	//       200: body:CreateTokenTemplateOutput

	// UpdateTokenTemplateURI ...
	UpdateTokenTemplateURI = "/ddxf/dtoken/update_template"
	// swagger:route POST /ddxf/dtoken/update_template  SwaggerUpdateTokenTemplateInput
	// update token template
	//     Responses:
	//       200: body:UpdateTokenTemplateOutput

	// RemoveTokenTemplateURI ...
	RemoveTokenTemplateURI = "/ddxf/dtoken/remove_template"
	// swagger:route POST /ddxf/dtoken/remove_template  SwaggerRemoveTokenTemplateInput
	// remove token template
	//     Responses:
	//       200: body:RemoveTokenTemplateOutput

	// GetTokenTemplateURI ...
	GetTokenTemplateURI = "/ddxf/dtoken/get_template"
	// swagger:route POST /ddxf/dtoken/get_template  SwaggerGetTokenTemplateInput
	// get token template
	//     Responses:
	//       200: body:GetTokenTemplateOutput

	// SetTokenTemplateAuthURI ...
	SetTokenTemplateAuthURI = "/ddxf/dtoken/set_template_auth"
	// swagger:route POST /ddxf/dtoken/set_template_auth  SwaggerSetTokenTemplateAuthInput
	// set token template auth
	//     Responses:
	//       200: body:SetTokenTemplateAuthOutput

	// AddTokenTemplateAuthURI ...
	AddTokenTemplateAuthURI = "/ddxf/dtoken/add_template_auth"
	// swagger:route POST /ddxf/dtoken/add_template_auth  SwaggerAddTokenTemplateAuthInput
	// add token template auth
	//     Responses:
	//       200: body:AddTokenTemplateAuthOutput

	// RemoveTokenTemplateAuthURI ...
	RemoveTokenTemplateAuthURI = "/ddxf/dtoken/remove_template_auth"
	// swagger:route POST /ddxf/dtoken/remove_template_auth  SwaggerRemoveTokenTemplateAuthInput
	// remove token template auth
	//     Responses:
	//       200: body:RemoveTokenTemplateAuthOutput

	// ClearTokenTemplateAuthURI ...
	ClearTokenTemplateAuthURI = "/ddxf/dtoken/clear_template_auth"
	// swagger:route POST /ddxf/dtoken/clear_template_auth  SwaggerClearTokenTemplateAuthInput
	// clear token template auth
	//     Responses:
	//       200: body:ClearTokenTemplateAuthOutput

	// GenerateTokenURI ...
	GenerateTokenURI = "/ddxf/dtoken/generate"
	// swagger:route POST /ddxf/dtoken/generate  SwaggerGenerateTokenInput
	// generate token
	//     Responses:
	//       200: body:GenerateTokenOutput
)
