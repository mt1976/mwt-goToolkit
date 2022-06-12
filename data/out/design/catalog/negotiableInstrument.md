# **NegotiableInstrument** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**NegotiableInstrument** (negotiableinstrument) |
|Endpoint 	    |**/NegotiableInstrument...** [^1]|
|Endpoint Query |**Id**|
|REST API|**/API/NegotiableInstrument/**|
Glyph|**fas fa-file-signature** (text-info)
Friendly Name|**Negotiable Instrument**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/NegotiableInstrument/NegotiableInstrumentList) [Exportable]
* **View** (/NegotiableInstrument/NegotiableInstrumentView)

* **Save** (/NegotiableInstrument/NegotiableInstrumentSave)









##  Provides


* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **lseGiltsDataStore**
SQL Table Key | **Id**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|
|**Id**|String|false|true|false|false|||||Y|id||
|**LongName**|String|false|true|false|false|||||Y|longName||
|**Isin**|String|false|true|false|false|||||Y|isin||
|**Tidm**|String|false|true|false|false|||||Y|tidm||
|**Sedol**|String|false|true|false|false|||||Y|sedol||
|**IssueDate**|String|false|true|false|false|||||Y|issueDate||
|**MaturityDate**|String|false|true|false|false|||||Y|maturityDate||
|**CouponValue**|String|false|true|false|false|||||Y|couponValue||
|**CouponType**|String|false|true|false|false|||||Y|couponType||
|**Segment**|String|false|true|false|false|||||Y|segment||
|**Sector**|String|false|true|false|false|||||Y|sector||
|**CodeConventionCalculateAccrual**|String|false|true|false|false|||||Y|codeConventionCalculateAccrual||
|**MinimumDenomination**|String|false|true|false|false|||||Y|minimumDenomination||
|**DenominationCurrency**|String|false|true|false|false|||||Y|denominationCurrency||
|**TradingCurrency**|String|false|true|false|false|||||Y|tradingCurrency||
|**Type**|String|false|true|false|false|||||Y|type||
|**FlatYield**|String|false|true|false|false|||||Y|flatYield||
|**PaymentCouponDate**|String|false|true|false|false|||||Y|paymentCouponDate||
|**PeriodOfCoupon**|String|false|true|false|false|||||Y|periodOfCoupon||
|**ExCouponDate**|String|false|true|false|false|||||Y|exCouponDate||
|**DateOfIndexInflation**|String|false|true|false|false|||||Y|dateOfIndexInflation||
|**UnitOfQuotation**|String|false|true|false|false|||||Y|unitOfQuotation||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||
|**SYSWho**|String|false|true|false|false|||||NH|_who||
|**SYSHost**|String|false|true|false|false|||||NH|_host||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||
|**Issuer**|String|false|true|false|false|||||Y|issuer||
|**IssueAmount**|String|false|true|false|false|||||Y|issueAmount||
|**RunningYield**|String|false|true|false|false|||||Y|runningYield||
|**LEI**|String|false|true|false|false|||||Y|LEI||
|**CUSIP**|String|false|true|false|false|||||Y|CUSIP||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/negotiableInstrument.go_tmp |
| code | **api** | /application/negotiableInstrument_api.go |
| code | **dao** | /dao/negotiableInstrument.go_tmp |
| code | **datamodel** | /datamodel/negotiableInstrument.go_tmp |
| code | **menu** | /design/menu/negotiableInstrument.json |
| html | **list** | /html/NegotiableInstrument_List.html |
| html | **view** | /html/NegotiableInstrument_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:18:08**
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