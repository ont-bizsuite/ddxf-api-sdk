package mp

// SwaggerPublishItemInput ...
// swagger:parameters SwaggerPublishItemInput
type SwaggerPublishItemInput struct {
	// in: body
	Input PublishItemInput
}

// SwaggerUpdateItemInput ...
// swagger:parameters SwaggerUpdateItemInput
type SwaggerUpdateItemInput struct {
	// in: body
	Input UpdateItemInput
}

// SwaggerDeleteItemInput ...
// swagger:parameters SwaggerDeleteItemInput
type SwaggerDeleteItemInput struct {
	// in: body
	Input DeleteItemInput
}

// SwaggerBuyItemInput ...
// swagger:parameters SwaggerBuyItemInput
type SwaggerBuyItemInput struct {
	// in: body
	Input BuyItemInput
}

// SwaggerGetItemInput ...
// swagger:parameters SwaggerGetItemInput
type SwaggerGetItemInput struct {
	// in: body
	Input GetItemInput
}
