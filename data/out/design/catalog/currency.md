# **Currency** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Currency** (currency) |
|Endpoint 	    |**/Currency...** [^1]|
|Endpoint Query |**Code**|
|REST API|**/API/Currency/**|
Glyph|**fas fa-dollar-sign** ()
Friendly Name|**Currency**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Currency/CurrencyList) [Exportable]
* **View** (/Currency/CurrencyView)











##  Provides
 * Lookup (Code Name)
 * Reverse Lookup (Name)





##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaCurrency**
SQL Table Key | **Code**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**Code**|String|true|true|false|false|||||Y|Code||
|**Name**|String|false|true|false|false|||||Y|Name||
|**AmountDp**|Int|false|true|false|false|||||Y|AmountDp|0|
|**Country**|String|false|true|false|false|||||Y|Country||
|**CountryName**|String|false|true|false|false|||||Y|CountryName||
|**IntBase**|String|false|true|false|false|||||Y|IntBase||
|**KeydateBase**|String|false|true|false|false|||||Y|KeydateBase||
|**InterestRateTolerance**|Float|false|true|false|false|||||Y|InterestRateTolerance|0.00|
|**CheckPayTo**|Bool|false|true|false|false|||||Y|CheckPayTo|True|
|**LatinAmericanSettlement**|Bool|false|true|false|false|||||Y|LatinAmericanSettlement|True|
|**DefaultLayOffBookKey**|String|false|true|false|false|||||Y|DefaultLayOffBookKey||
|**CutOffTimeCutOffTime**|Time|false|true|false|false|||||Y|CutOffTimeCutOffTime||
|**CutOffTimeTimeZone**|String|false|true|false|false|||||Y|CutOffTimeTimeZone||
|**CutOffTimeDerivedDataUTCOffset**|String|false|true|false|false|||||Y|CutOffTimeDerivedDataUTCOffset||
|**CutOffTimeDerivedDataHasDaylightSaving**|Bool|false|true|false|false|||||Y|CutOffTimeDerivedDataHasDaylightSaving|True|
|**CutOffTimeDerivedDataDaylightStart**|Time|false|true|false|false|||||Y|CutOffTimeDerivedDataDaylightStart||
|**CutOffTimeDerivedDataDaylightEnd**|Time|false|true|false|false|||||Y|CutOffTimeDerivedDataDaylightEnd||
|**DealerInterventionQuoteTimeout**|Int|false|true|false|false|||||Y|DealerInterventionQuoteTimeout|0|
|**CutOffTimeCutOffPeriod**|String|false|true|false|false|||||Y|CutOffTimeCutOffPeriod||
|**StripRateFutureExchangeCode**|String|false|true|false|false|||||Y|StripRateFutureExchangeCode||
|**StripRateFutureCurrencyContractCurrencyIsoCode**|String|false|true|false|false|||||Y|StripRateFutureCurrencyContractCurrencyIsoCode||
|**StripRateFutureCurrencyContractFutureContractCode**|String|false|true|false|false|||||Y|StripRateFutureCurrencyContractFutureContractCode||
|**OvernightFundingSpreadBid**|Float|false|true|false|false|||||Y|OvernightFundingSpreadBid|0.00|
|**OvernightFundingSpreadOffer**|Float|false|true|false|false|||||Y|OvernightFundingSpreadOffer|0.00|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/currency.go_tmp |
| code | **adaptor** | /adaptor/currency_impl.template |
| code | **api** | /application/currency_api.go |
| code | **dao** | /dao/currency.go_tmp |
| code | **datamodel** | /datamodel/currency.go_tmp |
| code | **menu** | /design/menu/currency.json |
| html | **list** | /html/Currency_List.html |
| html | **view** | /html/Currency_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:18:02**
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