# DDXF SDK使用文档

## Overview

`DDXF`由4个子模块组成：`meta`、`data meta`、`marketplace`、`token`，分别对应4个子`SDK`，下面分别进行介绍。


------------------------------------------------

### meta sdk

#### 初始化

```golang
NewMetaSdk(addr string)
```
入参`addr`是meta server地址。

创建完实例，然后就可以对meta进行常规的CRUD操作，分别是`CreateMeta`、`UpdateMeta`、`DeleteMeta`、`GetMeta`，下面分别介绍。

#### 创建meta

```golang

func (sdk *MetaSdk) CreateMeta(ontIDAcc *ontology_go_sdk.Account, input io.CreateMetaInput) (metaID string, err error)
```

其中：
1. `ontIDAcc`可以通过[`ontology-go-sdk`](https://github.com/ontio/ontology-go-sdk)的`NewAccountFromPrivateKey`方法创建。
2. `input`是meta相关详情。

#### 更新meta

```golang
func (sdk *MetaSdk) UpdateMeta(ontIDAcc *ontology_go_sdk.Account, input io.UpdateMetaInput) (err error) 
```


#### 删除meta

```golang
func (sdk *MetaSdk) DeleteMeta(ontIDAcc *ontology_go_sdk.Account, input io.DeleteMetaInput) (err error)
```

#### 获取meta

```golang
func (sdk *MetaSdk) GetMeta(input io.GetMetaInput) (metaResult map[string]interface{}, err error)
```
------------------------------------------------

### data meta sdk

#### 初始化

```golang
func NewDataMetaSdk(ddxfAPIAddr, ontologyApiAddr string) *DataMetaSdk
```
入参`ddxfAPIAddr`是data meta server地址，`ontologyApiAddr`是ontology节点地址。

#### 创建data meta

```golang
func (m *DataMetaSdk) CreateDataMeta(ontIdAcc *ontology_go_sdk.Account, input io.CreateDataMetaInput) (out io.CreateDataMetaOutput, err error)
```

#### 更新data meta

```golang
func (m *DataMetaSdk) UpdateDataMeta(ontIdAcc *ontology_go_sdk.Account, input io.UpdateDataMetaInput) (out io.UpdateDataMetaOutput, err error)
```

#### 删除data meta

```golang
func (m *DataMetaSdk) RemoveDataMeta(ontIdAcc *ontology_go_sdk.Account, input io.RemoveDataMetaInput) (out io.RemoveDataMetaOutput, err error)
```

------------------------------------------------

### marketplace sdk

#### 初始化

```golang
func NewMarketplaceSdk(ddxfAPIAddr, ontologyApiAddr, mpContract string) *Marketplace
```


#### 发布商品

```golang
func (mp *Marketplace) PublishItem(ontIdAcc *ontology_go_sdk.Account, input io.PublishItemInput, seller, mpAcc *ontology_go_sdk.Account) (out io.PublishItemOutput, err error)
```

#### 更新商品

```golang
func (mp *Marketplace) UpdateItem(ontIdAcc *ontology_go_sdk.Account, input io.UpdateItemInput, seller *ontology_go_sdk.Account) (out io.UpdateItemOutput, err error)
```

#### 删除商品

```golang
func (mp *Marketplace) DeleteItem(ontIdAcc *ontology_go_sdk.Account, input io.DeleteItemInput, acc *ontology_go_sdk.Account) (out io.DeleteItemOutput, err error)
```

#### 购买商品

```golang
func (mp *Marketplace) BuyItem(ontIdAcc *ontology_go_sdk.Account, input io.BuyItemInput, buyer *ontology_go_sdk.Account) (out io.BuyItemOutput, err error)
```

#### 查询商品

```golang
func (mp *Marketplace) GetItem(input io.GetItemInput) (*io.GetItemOutput, error)
```

------------------------------------------------

### token sdk

#### 初始化

```golang
func NewTokenSdk(ddxfAPIAddr, ontologyApiAddr, tokenContract string) *TokenSdk
```

#### 验证token

```golang
func (ts *TokenSdk) VerifyToken(input io.VerifyTokenInput) (err error) 
```

#### 更新token

```golang
func (ts *TokenSdk) UseToken(ontIdAcc *ontology_go_sdk.Account, input io.UseTokenInput) (out io.UseTokenOutput, err error)
```

#### 转移token
```golang
func (ts *TokenSdk) TransferToken(ontIdAcc *ontology_go_sdk.Account, input io.TransferTokenInput) (out io.TransferTokenOutput, err error)
```

#### 生成token

```golang
func (ts *TokenSdk) GenerateToken(ontIdAcc *ontology_go_sdk.Account, input io.GenerateTokenInput) (tokenId string, err error)
```

#### 重置token代理

```golang
func (ts *TokenSdk) SetTokenAgent(ontIdAcc *ontology_go_sdk.Account, input io.SetTokenAgentInput)
```

#### 添加token代理

```golang
func (ts *TokenSdk) AddTokenAgent(ontIdAcc *ontology_go_sdk.Account, input io.AddTokenAgentInput)
```

#### 删除token代理

```golang
func (ts *TokenSdk) RemoveTokenAgent(ontIdAcc *ontology_go_sdk.Account, input io.AddTokenAgentInput) (err error)
```

#### 查询token

```golang
func (ts *TokenSdk) GetToken(ontIdAcc *ontology_go_sdk.Account, input io.GetTokenInput) (out io.GetTokenOutput, err error)
```


#### 创建token template

```golang
func (ts *TokenSdk) CreateTokenTemplate(ontIdAcc *ontology_go_sdk.Account, input io.CreateTokenTemplateInput) (tokenTemplateId string, err error)
```

#### 更新token template

```golang
func (ts *TokenSdk) UpdateTokenTemplate(ontIdAcc *ontology_go_sdk.Account, input io.UpdateTokenTemplateInput) (err error)
```


#### 删除token template

```golang
func (ts *TokenSdk) RemoveTokenTemplate(ontIdAcc *ontology_go_sdk.Account, input io.RemoveTokenTemplateAuthInput) (err error)
```

#### 查询token template

```golang
func (ts *TokenSdk) GetTokenTemplate(ontIdAcc *ontology_go_sdk.Account, input io.GetTokenTemplateInput) (out io.GetTokenTemplateOutput, err error)
```

#### 重置token template授权

```golang
func (ts *TokenSdk) SetTokenTemplateAuth(ontIdAcc *ontology_go_sdk.Account, input io.SetTokenTemplateAuthInput) (err error)
```

#### 添加token template授权

```golang
func (ts *TokenSdk) AddTokenTemplateAuth(ontIdAcc *ontology_go_sdk.Account, input io.AddTokenTemplateAuthInput) (err error)
```

#### 删除token template授权

```golang
func (ts *TokenSdk) RemoveTokenTemplateAuth(ontIdAcc *ontology_go_sdk.Account, input io.RemoveTokenTemplateAuthInput) (err error)
```

#### 清除token template授权

```golang
func (ts *TokenSdk) ClearTokenTemplateAuth(ontIdAcc *ontology_go_sdk.Account, input io.ClearTokenTemplateAuthInput) (err error)
```
