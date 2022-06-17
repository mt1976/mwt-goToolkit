# **CounterpartyAddress** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**CounterpartyAddress** (counterpartyaddress) |
|Endpoint 	    |**/CounterpartyAddress...** [^1]|
|Endpoint Query |**ID**|
Glyph|**fas fa-passport** ()
Friendly Name|**Counterparty Address**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/CounterpartyAddress/CounterpartyAddressList) [Exportable]
* **View** (/CounterpartyAddress/CounterpartyAddressView)
* **Edit** (/CounterpartyAddress/CounterpartyAddressEdit)
* **Save** (/CounterpartyAddress/CounterpartyAddressSave)
* **New** (/CounterpartyAddress/CounterpartyAddressNew)
* **Delete** (/CounterpartyAddress/CounterpartyAddressDelete)







##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaCounterpartyAddress**
SQL Table Key | **CompID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**NameFirm**|String|true|true|false|false|||||Y|NameFirm||false|false|false|
|**NameCentre**|String|true|true|false|false|||||Y|NameCentre||false|false|false|
|**Address1**|String|false|true|false|false|||||Y|Address1||false|false|false|
|**Address2**|String|false|true|false|false|||||Y|Address2||false|false|false|
|**CityTown**|String|false|true|false|false|||||Y|CityTown||false|false|false|
|**County**|String|false|true|false|false|||||Y|County||false|false|false|
|**Postcode**|String|false|true|false|false|||||Y|Postcode||false|false|false|
|**CompID**|String|true|true|false|false|||||Y|CompID||false|false|false|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/counterpartyAddress_core.go_tmp |
| code | **adaptor** | /adaptor/counterpartyAddress_impl.go_template_tmp |
| code | **dao** | /dao/counterpartyAddress_core.go_tmp |
| code | **datamodel** | /datamodel/counterpartyAddress_core.go_tmp |
| code | **menu** | /design/menu/counterpartyAddress.json_tmp |
| html/base | **list** | /CounterpartyAddress_List.html |
| html/base | **view** | /CounterpartyAddress_View.html |
| html/base | **edit** | /CounterpartyAddress_Edit.html |
| html/base | **new** | /CounterpartyAddress_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **16/06/2022** at **13:13:58**
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