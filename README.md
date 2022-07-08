# sap-api-integrations-activity-reads
sap-api-integrations-activity-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 活動データを取得するマイクロサービスです。    
sap-api-integrations-activity-reads には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-activity-reads は、オンプレミス版である（＝クラウド版ではない）SAPC4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/activity/overview


## 動作環境  
sap-api-integrations-activity-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）　　

## クラウド環境での利用
sap-api-integrations-activity-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-activity-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/activity/overview
* APIサービス名(=baseURL): c4codataapi

## 本レポジトリ に 含まれる API名
sap-api-integrations-activity-reads には、次の API をコールするためのリソースが含まれています。  

* AppointmentCollection（活動 - アポイントメント）※アポイントメントの詳細データを取得するために、ToAppointmentInvolvedParties、ToAppointmentOtherParties、と合わせて利用されます。
* ToAppointmentInvolvedParties（活動 - アポイントメント関与先 ※To）
* ToAppointmentOtherParties（活動 - アポイントメントその他関与先 ※To）
* EmailCollection(活動　- Eメール)※Eメールの詳細データを取得するために、ToEmailRecipients、ToEmailSender、と合わせて利用されます。
* ToEmailRecipients(活動 - Eメール受信者　※To)
* ToEmailSender(活動 - Eメール送信者 ※To)
* TasksCollection(活動　- タスク)

## API への 値入力条件 の 初期値
sap-api-integrations-activity-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.AppointmentCollection.ID（ID）
* inoutSDC.EmailCollection.Subject（主題）
* inoutSDC.TaskCollection.Category（カテゴリ）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"AppointmentCollection" が指定されています。

```
  "api_schema": "AppointmentCollection",
  "accepter": ["AppointmentCollection"],
  "contract_code": "2862",
  "deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
  "api_schema": "AppointmentCollection",
  "accepter": ["All"],
  "contract_code": "2862",
  "deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetActivity(iD, subject, category string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "AppointmentCollection":
			func() {
				c.AppointmentCollection(iD)
				wg.Done()
			}()
		case "EmailCollection":
			func() {
				c.EmailCollection(subject)
				wg.Done()
			}()
		case "TaskCollection":
			func() {
				c.TaskCollection(category)
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
以下の sample.json の例は、SAP 活動 の アポイントメントデータ が取得された結果の JSON の例です。  
以下の項目のうち、"ObjectID" ～ "AppointmentOtherParties" は、/SAP_API_Output_Formatter/type.go 内 の Type AppointmentCollection {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integtrations-activity-reads/SAP_API_Caller/caller.go#L58",
	"function": "sap-api-integrations-activity-reads/SAP_API_Caller.(*SAPAPICaller).AppointmentCollection",
	"level": "INFO",
	"message": [
		{
			"ObjectID": "00163E05A2E01EE4ADBBB58D491DDB6E",
			"DocumentType": "0001",
			"DocumentTypeText": "Appointment",
			"ID": "2862",
			"Subject": "Roadshow 2019",
			"LifeCycleStatusCode": "1",
			"LifeCycleStatusCodeText": "Open",
			"StartDate": "2019-11-13T06:00:00+09:00",
			"StartDatetimeZoneCode": "UTC",
			"StartDatetimeZoneCodeText": "UTC+0",
			"EndDate": "2019-11-14T07:00:00+09:00",
			"EnddatetimeZoneCode": "UTC",
			"EnddatetimeZoneCodeText": "UTC+0",
			"OrganizerUUID": "00163E03-A070-1EE2-88BA-39174566D0B3",
			"OrganizerPartyID": "8000000006",
			"Owner": "00163E03-A070-1EE2-88BA-3BF3813A50C6",
			"PrimaryContact": "00163E0C-6CDB-1ED6-84D4-217462CEE486",
			"OwnerPartyID": "8000000020",
			"FullDayIndicator": false,
			"Category": "0001",
			"CategoryText": "Customer Visit",
			"Account": "00163E0C-6CDB-1EE5-8DE9-4B48F9E4BBBA",
			"MainAccountPartyID": "1001710",
			"MainContactPartyID": "1004188",
			"Location": "ALPHA Center / 9009 North Avenue / MINNEAPOLIS MN 55427 / US",
			"Priority": "3",
			"PriorityText": "Normal",
			"SalesTerritoryUUID": "",
			"SalesTerritoryID": "",
			"SalesTerritoryName": "",
			"SalesOrganisation": "US1100",
			"DistributionChannel": "01",
			"DistributionChannelText": "Direct sales",
			"Division": "00",
			"DivisionText": "Cross Division",
			"VisitTourPlanUUID": "",
			"InformationSensitivityCode": "1",
			"InformationSensitivityCodeText": "Normal",
			"AccountName": "ALPHA Center",
			"AdditionalLocationName": "",
			"ChangedBy": "Kate Jacob",
			"CompletionPercent": "0.00",
			"CreatedBy": "Diane Lacey",
			"CreatedOn": "2015-02-16T22:18:54+09:00",
			"DataOriginTypeCode": "1",
			"DataOriginTypeCodeText": "Manual Data Entry",
			"GroupwareItemID": "",
			"OrganizerEmailURI": "ussalesman@ondemand.com",
			"OrganizerName": "Kate Jacob",
			"OrganizerPartyName": "Kate Jacob",
			"OwnerName": "Diane Lacey",
			"Phone": "+1 847-778-8971",
			"PrimaryContactName": "Lou Daly",
			"ReportedDateTime": "2019-11-13T06:00:00+09:00",
			"UUID": "00163E05-A2E0-1EE4-ADBB-B58D491DDB6E",
			"AvailabilityCode": "",
			"AvailabilityCodeText": "",
			"LastChangeDateTime": "2019-11-12T05:54:02+09:00",
			"EntityLastChangedOn": "2019-11-12T05:54:02+09:00",
			"ETag": "2019-11-18T15:16:31+09:00",
			"VisitTypeCode": "",
			"VisitTypeCodeText": "",
			"OccurrenceDate": "",
			"RecurrenceActGroupwareID": "",
			"RecurringBusinessActivityNodeID": "",
			"CollaborationChannelCode": "",
			"CollaborationChannelCodeText": "",
			"CollaborationURL": "",
			"Draft": false,
			"CreatedBySSI": "",
			"CreatedBySSIText": "",
			"dataloader1_KUT": "",
			"AppointmentInvolvedParties": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/AppointmentCollection('00163E05A2E01EE4ADBBB58D491DDB6E')/AppointmentInvolvedParties",
			"AppointmentOtherParties": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/AppointmentCollection('00163E05A2E01EE4ADBBB58D491DDB6E')/AppointmentOtherParties"
		}
	],
	"time": "2022-07-08T15:51:58+09:00"
}

```
