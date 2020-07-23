package mp

import "github.com/ont-bizsuite/ddxf-sdk/market_place_contract"

// PublishEvent ...
type PublishEvent struct {
	ItemId string
	DDO    *market_place_contract.ResourceDDO
	Item   *market_place_contract.DTokenItem
}

// PublishEvent ...
type UpdateEvent struct {
	ItemId string
	DDO    *market_place_contract.ResourceDDO
	Item   *market_place_contract.DTokenItem
}

// PublishEvent ...
type DeleteEvent struct {
	ItemId string
}

// PublishEvent ...
type BuyDTokenEvent struct {
	ItemId string
	N      int
	Buyer  string
	Payer  string
}
