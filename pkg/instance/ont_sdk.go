package instance

import (
	ddxf_sdk "github.com/ont-bizsuite/ddxf-sdk"
	"github.com/zhiqiangxu/ddxf-api/pkg/config"
	"sync"
)

var (
	ontSdkOnce  sync.Once
	ddxfSdkOnce sync.Once
	ddxfSdk     *ddxf_sdk.DdxfSdk
)

// OntSdk is singleton for misc.OntSdk
func DDXFSdk() *ddxf_sdk.DdxfSdk {
	ddxfSdkOnce.Do(func() {
		if config.Load().Prod {
			ddxfSdk = ddxf_sdk.NewDdxfSdk("http://dappnode1.ont.io:20336")
		} else {
			ddxfSdk = ddxf_sdk.NewDdxfSdk("http://polaris2.ont.io:20336")
		}
	})
	ddxfSdk.SetPayer(config.Payer)
	return ddxfSdk
}
