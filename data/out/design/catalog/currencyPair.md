# **CurrencyPair** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**CurrencyPair** (currencypair) |
|Endpoint 	    |**/CurrencyPair...** [^1]|
|Endpoint Query |**Code**|
|REST API|**/API/CurrencyPair/**|
Glyph|**fas fa-coins** ()
Friendly Name|**Currency Pair**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/CurrencyPair/CurrencyPairList) [Exportable]
* **View** (/CurrencyPair/CurrencyPairView)
* **Edit** (/CurrencyPair/CurrencyPairEdit)
* **Save** (/CurrencyPair/CurrencyPairSave)
* **New** (/CurrencyPair/CurrencyPairNew)
* **Delete** (/CurrencyPair/CurrencyPairDelete)







##  Provides
 * Lookup (Code Code)






##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaCurrencyPair**
SQL Table Key | **Code**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**CodeMajorCurrencyIsoCode**|String|true|true|false|true|OL|Currency|CodeMajorCurrencyIsoCode|Name|Y|CodeMajorCurrencyIsoCode||true|false|false|
|**CodeMinorCurrencyIsoCode**|String|true|true|false|true|OL|Currency|CodeMinorCurrencyIsoCode|Name|Y|CodeMinorCurrencyIsoCode||true|false|false|
|**ReciprocalActive**|Bool|false|true|false|false|LL|tf|||Y|ReciprocalActive|True|false|false|false|
|**Code**|String|true|true|false|true|||||N|Code||true|true|false|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/currencyPair_core.go_tmp |
| code | **adaptor** | /adaptor/currencyPair_impl.go_template_tmp |
| code | **api** | /application/currencyPair_api.go_tmp |
| code | **dao** | /dao/currencyPair_core.go_tmp |
| code | **datamodel** | /datamodel/currencyPair_core.go_tmp |
| code | **menu** | /design/menu/currencyPair.json_tmp |
| html/base | **list** | /CurrencyPair_List.html |
| html/base | **view** | /CurrencyPair_View.html |
| html/base | **edit** | /CurrencyPair_Edit.html |
| html/base | **new** | /CurrencyPair_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **16/06/2022** at **13:14:00**
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