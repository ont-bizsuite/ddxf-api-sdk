package sdk

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ontio/ontology-crypto/keypair"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ont-bizsuite/ddxf-api-sdk/pkg/forward"
	io "github.com/ont-bizsuite/ddxf-api-sdk/pkg/io/meta"
)

// MetaSdk ...
type MetaSdk struct {
	addr string
}

// NewMetaSdk ...
func NewMetaSdk(addr string) *MetaSdk {
	return &MetaSdk{
		addr: addr,
	}
}

// CreateMeta ...
func (sdk *MetaSdk) CreateMeta(ontID string, ontIDAcc *ontology_go_sdk.Account, input io.CreateMetaInput) (metaID string, err error) {

	code, body, err := sdk.request(ontID, ontIDAcc, "/ddxf/meta/create", input)
	if code != http.StatusOK {
		err = fmt.Errorf("code:%d body:%s", code, body)
		return
	}

	var output io.CreateMetaOutput
	err = json.Unmarshal(body, &output)
	if err != nil {
		return
	}

	err = output.Error()
	if err != nil {
		return
	}

	metaID = output.MetaID
	return
}

// UpdateMeta ...
func (sdk *MetaSdk) UpdateMeta(ontID string, ontIDAcc *ontology_go_sdk.Account, input io.UpdateMetaInput) (err error) {
	code, body, err := sdk.request(ontID, ontIDAcc, "/ddxf/meta/update", input)
	if code != http.StatusOK {
		err = fmt.Errorf("code:%d body:%s", code, body)
		return
	}

	var output io.UpdateMetaOutput
	err = json.Unmarshal(body, &output)
	if err != nil {
		return
	}

	err = output.Error()
	if err != nil {
		return
	}
	return
}

// DeleteMeta ...
func (sdk *MetaSdk) DeleteMeta(ontID string, ontIDAcc *ontology_go_sdk.Account, input io.DeleteMetaInput) (err error) {
	code, body, err := sdk.request(ontID, ontIDAcc, "/ddxf/meta/delete", input)
	if code != http.StatusOK {
		err = fmt.Errorf("code:%d body:%s", code, body)
		return
	}

	var output io.DeleteMetaOutput
	err = json.Unmarshal(body, &output)
	if err != nil {
		return
	}

	err = output.Error()
	if err != nil {
		return
	}
	return
}

// GetMeta ...
func (sdk *MetaSdk) GetMeta(input io.GetMetaInput) (metaResult map[string]interface{}, err error) {
	code, body, err := sdk.request("", nil, "/ddxf/meta/get", input)
	if code != http.StatusOK {
		err = fmt.Errorf("code:%d body:%s", code, body)
		return
	}

	var output io.GetMetaOutput
	err = json.Unmarshal(body, &output)
	if err != nil {
		return
	}

	err = output.Error()
	if err != nil {
		return
	}

	metaResult = output.Meta
	return
}

func (sdk *MetaSdk) request(ontID string, ontIDAcc *ontology_go_sdk.Account, uri string, input interface{}) (code int, body []byte, err error) {
	bs, err := json.Marshal(input)
	if err != nil {
		return
	}

	var header map[string]string
	if ontID != "" {
		pk := keypair.SerializePublicKey(ontIDAcc.GetPublicKey())
		var sig []byte
		sig, err = ontIDAcc.Sign(bs)
		if err != nil {
			return
		}

		header = map[string]string{
			"DDXF_ONTID": ontID,
			"DDXF_PK":    hex.EncodeToString(pk),
			"DDXF_SIGN":  hex.EncodeToString(sig),
		}
	}

	code, _, body, err = forward.PostJSONRequest(sdk.addr+uri, bs, header)
	return
}
