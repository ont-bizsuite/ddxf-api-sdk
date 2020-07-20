package mp

const (
	// PublishItemURI ...
	PublishItemURI = "/ddxf/mp/publish_item"
	// swagger:route POST /ddxf/mp/publish_item SwaggerPublishItemInput
	// publish item
	//     Responses:
	//       200: body:PublishItemOutput

	// UpdateItemURI ...
	UpdateItemURI = "/ddxf/mp/update_item"
	// swagger:route POST /ddxf/mp/update_item SwaggerUpdateItemOutput
	// update item
	//     Responses:
	//       200: body:UpdateItemOutput

	// DeleteItemURI ...
	DeleteItemURI = "/ddxf/mp/delete_item"
	// swagger:route POST /ddxf/mp/delete_item SwaggerDeleteItemInput
	// delete item
	//     Responses:
	//       200: body:DeleteItemOutput

	// BuyItemURI ...
	BuyItemURI = "/ddxf/mp/buy_item"
	// swagger:route POST /ddxf/mp/buy SwaggerBuyItemInput
	// use token
	//     Responses:
	//       200: body:BuyItemOutput

	// GetItemURI ...
	GetItemURI = "/ddxf/mp/get_item"
	// swagger:route POST /ddxf/mp/get_item  SwaggerGetItemInput
	// get item
	//     Responses:
	//       200: body:GetItemOutput

)
