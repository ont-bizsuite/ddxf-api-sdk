package datameta

import "github.com/zhiqiangxu/ddxf-sdk/pkg/io"

// CreateDataMetaInput ...
// IDSign
type CreateDataMetaInput struct {
	// controller ontids for this data meta
	OntIDs []string
	// data fingerprint
	Fingerprint string
	// data meta id
	DataMetaHash string
	// endpoint for data meta hash
	Endpoint string
}

// ContainsOntID ...
func (input *CreateDataMetaInput) ContainsOntID(ontID string) bool {
	exists := false
	for _, oid := range input.OntIDs {
		if oid == ontID {
			exists = true
			break
		}
	}
	return exists
}

// DataID, MetaID, ...

// CreateDataMetaOutput ...
// swagger:model CreateDataMetaOutput
type CreateDataMetaOutput struct {
	io.BaseResp
	// uuid for this data meta
	DataID string
	// send to chain with ontology sdk
	Tx string
}

// UpdateDataMetaInput ...
type UpdateDataMetaInput struct {
	// uuid for this data meta
	DataID string
	// data fingerprint
	Fingerprint string
	// data meta hash
	DataMetaHash string
	// endpoint for data meta hash
	Endpoint string
}

// UpdateDataMetaOutput ...
// swagger:model UpdateDataMetaOutput
type UpdateDataMetaOutput struct {
	io.BaseResp
	// send to chain with ontology sdk
	Tx string
}

// RemoveDataMetaInput ...
type RemoveDataMetaInput struct {
	// uuid for this data meta
	DataID string
}

// RemoveDataMetaOutput ...
// swagger:model RemoveDataMetaOutput
type RemoveDataMetaOutput struct {
	io.BaseResp
	// send to chain with ontology sdk
	Tx string
}

// GetDataMetaInput ...
type GetDataMetaInput struct {
	// uuid for this data meta
	DataID string
}

// GetDataMetaOutput ...
// swagger:model GetDataMetaOutput
type GetDataMetaOutput struct {
	io.BaseResp
	// data fingerprint
	Fingerprint string
	// data meta stored at endpoint
	DataMeta map[string]interface{}
}
