# sap-api-integrations-master-recipe-reads
sap-api-integrations-master-recipe-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で マスタレシピデータを取得するマイクロサービスです。    
sap-api-integrations-master-recipe-reads には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-master-recipe-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。    
https://api.sap.com/api/OP_API_MASTER_RECIPE_0001/overview

## 動作環境  
sap-api-integrations-master-recipe-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）　　

## クラウド環境での利用
sap-api-integrations-master-recipe-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-master-recipe-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_MASTER_RECIPE_0001/overview  
* APIサービス名(=baseURL): API_MASTER_RECIPE

## 本レポジトリ に 含まれる API名
sap-api-integrations-master-recipe-reads には、次の API をコールするためのリソースが含まれています。  

* MasterRecipeHeader（マスタレシピ - ヘッダ）※マスタレシピの詳細データを取得するために、ToMaterialAssignment、ToOperation、ToPhase、ToPhaseRelationship、と合わせて利用されます。
* ToMaterialAssignment（マスタレシピ - 品目割当 ※To）
* ToOperation（マスタレシピ - 作業 ※To）
* ToPhase（マスタレシピ - フェーズ ※To）
* ToPhaseRelationship（マスタレシピ - フェーズ関係 ※To）

## API への 値入力条件 の 初期値
sap-api-integrations-master-recipe-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.MasterRecipe.MasterRecipeGroup（マスタレシピグループ）
* inoutSDC.MasterRecipe.MasterRecipe（マスタレシピ）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。

```
	"api_schema": "/MasterRecipeHeader",
	"accepter": ["Header"],
	"master_recipe_code": "1",
	"plant": "",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "/MasterRecipeHeader",
	"accepter": ["All"],
	"master_recipe_code": "1",
	"plant": "",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetMasterRecipe(masterRecipeGroup, masterRecipe string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(masterRecipeGroup, masterRecipe)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP マスタレシピ の ヘッダデータ が取得された結果の JSON の例です。  
以下の項目のうち、"MasterRecipeGroup" ～ "to_Operation" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-master-recipe-reads/SAP_API_Caller/caller.go#L53",
	"function": "sap-api-integrations-master-recipe-reads/SAP_API_Caller.(*SAPAPICaller).Header",
	"level": "INFO",
	"message": [
		{
			"MasterRecipeGroup": "41010007",
			"MasterRecipe": "1",
			"MasterRecipeInternalVersion": "1",
			"IsMarkedForDeletion": false,
			"BillOfOperationsDesc": "SG24 - Ink (Process)",
			"LongTextLanguageCode": "",
			"PlainLongText": "",
			"Plant": "1010",
			"BillOfOperationsStatus": "4",
			"BillOfOperationsUsage": "1",
			"ResponsiblePlannerGroup": "1",
			"BillOfOperationsProfile": "YPI1",
			"MinimumLotSizeQuantity": "0.000",
			"MaximumLotSizeQuantity": "99999999.000",
			"BillOfOperationsUnit": "CCM",
			"CreationDate": "2016-06-24T09:00:00+09:00",
			"LastChangeDate": "2016-06-24T09:00:00+09:00",
			"ValidityStartDate": "2007-01-01T09:00:00+09:00",
			"ValidityEndDate": "9999-12-31T09:00:00+09:00",
			"ChangeNumber": "",
			"ChangedDateTime": "",
			"to_MatlAssgmt": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_MASTER_RECIPE/MasterRecipeHeader(MasterRecipeGroup='41010007',MasterRecipe='1',MasterRecipeInternalVersion='1')/to_MatlAssgmt",
			"to_Operation": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_MASTER_RECIPE/MasterRecipeHeader(MasterRecipeGroup='41010007',MasterRecipe='1',MasterRecipeInternalVersion='1')/to_Operation"
		}
	],
	"time": "2022-01-28T16:10:19+09:00"
}
```
