# **MarketRates** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**MarketRates** (marketrates) |
|Endpoint 	    |**/MarketRates...** [^1]|
|Endpoint Query |**ID**|
|REST API|**/API/MarketRates/**|
Glyph|**fas fa-file-invoice-dollar** (text-info)
Friendly Name|**Rates Store Content**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/MarketRates/MarketRatesList) [Exportable]
* **View** (/MarketRates/MarketRatesView)











##  Provides


* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **rateDataStore**
SQL Table Key | **Id**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|
|**Id**|String|true|true|false|false|||||Y|id||
|**Bid**|String|false|true|false|false|||||Y|bid||
|**Mid**|String|false|true|false|false|||||Y|mid||
|**Offer**|String|false|true|false|false|||||Y|offer||
|**Market**|String|false|true|false|false|||||Y|market||
|**Tenor**|String|false|true|false|false|||||Y|tenor||
|**Series**|String|false|true|false|false|||||Y|series||
|**Name**|String|false|true|false|false|||||Y|name||
|**Source**|String|false|true|false|false|||||Y|source||
|**Destination**|String|false|true|false|false|||||Y|destination||
|**Class**|String|false|true|false|false|||||Y|class||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||
|**SYSWho**|String|false|true|false|false|||||NH|_who||
|**SYSHost**|String|false|true|false|false|||||NH|_host||
|**Date**|String|false|true|false|false|||||Y|date||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/marketRates.go_tmp |
| code | **api** | /application/marketRates_api.go |
| code | **dao** | /dao/marketRates.go_tmp |
| code | **datamodel** | /datamodel/marketRates.go_tmp |
| code | **menu** | /design/menu/marketRates.json |
| html | **list** | /html/MarketRates_List.html |
| html | **view** | /html/MarketRates_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:18:09**
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