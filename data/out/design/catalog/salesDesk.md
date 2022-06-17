# **SalesDesk** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**SalesDesk** (salesdesk) |
|Endpoint 	    |**/SalesDesk...** [^1]|
|Endpoint Query |**Desk**|
Glyph|**fas fa-industry** ()
Friendly Name|**Sales Desk**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}













##  Provides
 * Lookup (Name Name)
 * Reverse Lookup (Name)





##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaSalesDesk**
SQL Table Key | **Name**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**Name**|String|true|true|false|false|||||Y|Name||false|false|false|
|**ReportDealsOver**|String|false|true|false|false|||||Y|ReportDealsOver||false|false|false|
|**ReportDealsOverCCY**|String|false|true|false|false|||||Y|ReportDealsOverCCY||false|false|false|
|**AccountTransferCutOffTime**|Time|false|true|false|false|||||Y|AccountTransferCutOffTime||false|false|false|
|**AccountTransferCutOffTimeTimeZone**|String|false|true|false|false|||||Y|AccountTransferCutOffTimeTimeZone||false|false|false|
|**AccountTransferCutOffTimeCutOffPeriod**|String|false|true|false|false|||||Y|AccountTransferCutOffTimeCutOffPeriod||false|false|false|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/salesDesk_core.go_tmp |
| code | **dao** | /dao/salesDesk_core.go_tmp |
| code | **datamodel** | /datamodel/salesDesk_core.go_tmp |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **16/06/2022** at **13:14:04**
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