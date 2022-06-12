# **DealingInterface** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**DealingInterface** (dealinginterface) |
|Endpoint 	    |**/DealingInterface...** [^1]|
|Endpoint Query |**Name**|
Glyph|**fas fa-share-alt** ()
Friendly Name|**Dealing Interface**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/DealingInterface/DealingInterfaceList) [Exportable]
* **View** (/DealingInterface/DealingInterfaceView)
* **Edit** (/DealingInterface/DealingInterfaceEdit)
* **Save** (/DealingInterface/DealingInterfaceSave)
* **New** (/DealingInterface/DealingInterfaceNew)
* **Delete** (/DealingInterface/DealingInterfaceDelete)







##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaDealingInterface**
SQL Table Key | **Name**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**Name**|String|true|true|false|false|||||Y|Name||
|**AcceptReducedAmount**|Bool|false|true|false|false|||||Y|AcceptReducedAmount|True|
|**QuoteAsIndicative**|Bool|false|true|false|false|||||Y|QuoteAsIndicative|True|
|**RateTimeOut**|Int|false|true|false|false|||||Y|RateTimeOut|0|
|**PropagationDelay**|Int|false|true|false|false|||||Y|PropagationDelay|0|
|**CheckLiquidity**|Bool|false|true|false|false|||||Y|CheckLiquidity|True|
|**ChangeQuoteDirection**|Bool|false|true|false|false|||||Y|ChangeQuoteDirection|True|
|**GenerateRejectedDeals**|Bool|false|true|false|false|||||Y|GenerateRejectedDeals|True|
|**SpotUpdatesForForwardQuotes**|Bool|false|true|false|false|||||Y|SpotUpdatesForForwardQuotes|True|
|**SettlementInstructionStyle**|String|false|true|false|false|||||Y|SettlementInstructionStyle||
|**CanRetractQuotes**|Bool|false|true|false|false|||||Y|CanRetractQuotes|True|
|**CancelESPifNotPriced**|Bool|false|true|false|false|||||Y|CancelESPifNotPriced|True|
|**CancelRFQSifNotPriced**|Bool|false|true|false|false|||||Y|CancelRFQSifNotPriced|True|
|**CancelonDealingSuspended**|Bool|false|true|false|false|||||Y|CancelonDealingSuspended|True|
|**CreditCheckedatDI**|Bool|false|true|false|false|||||Y|CreditCheckedatDI|True|
|**DuplicateCheckonExternalRef**|Bool|false|true|false|false|||||Y|DuplicateCheckonExternalRef|True|
|**LimitCheckQuote**|Bool|false|true|false|false|||||Y|LimitCheckQuote|True|
|**LimitCheckonRFQDealSubmission**|Bool|false|true|false|false|||||Y|LimitCheckonRFQDealSubmission|True|
|**ListenonLimits**|Bool|false|true|false|false|||||Y|ListenonLimits|True|
|**MarginStyle**|String|false|true|false|false|||||Y|MarginStyle||
|**UseRerouteDefinitionOnly**|Bool|false|true|false|false|||||Y|UseRerouteDefinitionOnly|True|
|**BypassConfirmation**|Bool|false|true|false|false|||||Y|BypassConfirmation|True|
|**DIOnAcceptance**|Bool|false|true|false|false|||||Y|DIOnAcceptance|True|
|**IgnoreESPAmountRules**|Bool|false|true|false|false|||||Y|IgnoreESPAmountRules|True|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/dealingInterface.go_tmp |
| code | **adaptor** | /adaptor/dealingInterface_impl.template |
| code | **dao** | /dao/dealingInterface.go_tmp |
| code | **datamodel** | /datamodel/dealingInterface.go_tmp |
| code | **menu** | /design/menu/dealingInterface.json |
| html | **list** | /html/DealingInterface_List.html |
| html | **view** | /html/DealingInterface_View.html |
| html | **edit** | /html/DealingInterface_Edit.html |
| html | **new** | /html/DealingInterface_New.html |


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