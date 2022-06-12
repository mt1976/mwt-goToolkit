# **DealType** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**DealType** (dealtype) |
|Endpoint 	    |**/DealType...** [^1]|
|Endpoint Query |**DealTypeKey**|
Glyph|**fas fa-toolbox** ()
Friendly Name|**Deal Type**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/DealType/DealTypeList) [Exportable]
* **View** (/DealType/DealTypeView)
* **Edit** (/DealType/DealTypeEdit)
* **Save** (/DealType/DealTypeSave)
* **New** (/DealType/DealTypeNew)
* **Delete** (/DealType/DealTypeDelete)







##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaDealType**
SQL Table Key | **DealTypeKey**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**DealTypeKey**|String|true|true|false|false|||||Y|DealTypeKey||
|**DealTypeShortName**|String|false|true|false|false|||||Y|DealTypeShortName||
|**HostKey**|String|false|true|false|false|||||Y|HostKey||
|**IsActive**|Bool|false|true|false|false|||||Y|IsActive|True|
|**Interbook**|Bool|false|true|false|false|||||Y|Interbook|True|
|**BackOfficeLink**|Bool|false|true|false|false|||||Y|BackOfficeLink|True|
|**HasTicket**|Bool|false|true|false|false|||||Y|HasTicket|True|
|**CurrencyOverride**|Bool|false|true|false|false|||||Y|CurrencyOverride|True|
|**CurrencyHolderCurrency**|String|false|true|false|false|||||Y|CurrencyHolderCurrency||
|**AllBooks**|Bool|false|true|false|false|||||Y|AllBooks|True|
|**FundamentalDealTypeKey**|String|false|true|false|false|||||Y|FundamentalDealTypeKey||
|**RelatedDealType**|String|false|true|false|false|||||Y|RelatedDealType||
|**BookName**|String|false|true|false|false|||||Y|BookName||
|**ExportMethod**|String|false|true|false|false|||||Y|ExportMethod||
|**DefaultUserLayoffBooks**|Bool|false|true|false|false|||||Y|DefaultUserLayoffBooks|True|
|**RFQ**|Bool|false|true|false|false|||||Y|RFQ|True|
|**OBS**|Bool|false|true|false|false|||||Y|OBS|True|
|**KID**|Bool|false|true|false|false|||||Y|KID|True|
|**InternalId**|Int|true|true|false|false|||||Y|InternalId|0|
|**InternalDeleted**|Time|false|true|false|false|||||Y|InternalDeleted||
|**UpdatedTransactionId**|String|false|true|false|false|||||Y|UpdatedTransactionId||
|**UpdatedUserId**|String|false|true|false|false|||||Y|UpdatedUserId||
|**UpdatedDateTime**|Time|false|true|false|false|||||Y|UpdatedDateTime||
|**DeletedTransactionId**|String|false|true|false|false|||||Y|DeletedTransactionId||
|**DeletedUserId**|String|false|true|false|false|||||Y|DeletedUserId||
|**ChangeType**|String|false|true|false|false|||||Y|ChangeType||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/dealType.go_tmp |
| code | **adaptor** | /adaptor/dealType_impl.template |
| code | **dao** | /dao/dealType.go_tmp |
| code | **datamodel** | /datamodel/dealType.go_tmp |
| code | **menu** | /design/menu/dealType.json |
| html | **list** | /html/DealType_List.html |
| html | **view** | /html/DealType_View.html |
| html | **edit** | /html/DealType_Edit.html |
| html | **new** | /html/DealType_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:18:06**
Who & Where		     | **matttownsend (Matt Townsend)** on **silicon.local**

### Footnotes
[^1]: **Endpoint**
    * The full list of endpoints can be found in the [Actions](#action-id) section
[^2]: **Lookup Fields**
    * LL = A List Lookup. Define list in lits.cfg
    * OL = An Object Lookup. Get a list of values from an Object
    * FL = Fetches 1 value from an object based on the content of the field. 
[^3]: **Inputtable**   
    * H = Hidden Field
    * N = No Input Field
    * Y = Inputtable Field