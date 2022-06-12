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
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**NameFirm**|String|true|true|false|false|||||Y|NameFirm||
|**NameCentre**|String|true|true|false|false|||||Y|NameCentre||
|**Address1**|String|false|true|false|false|||||Y|Address1||
|**Address2**|String|false|true|false|false|||||Y|Address2||
|**CityTown**|String|false|true|false|false|||||Y|CityTown||
|**County**|String|false|true|false|false|||||Y|County||
|**Postcode**|String|false|true|false|false|||||Y|Postcode||
|**CompID**|String|true|true|false|false|||||Y|CompID||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/counterpartyAddress.go_tmp |
| code | **adaptor** | /adaptor/counterpartyAddress_impl.template |
| code | **dao** | /dao/counterpartyAddress.go_tmp |
| code | **datamodel** | /datamodel/counterpartyAddress.go_tmp |
| code | **menu** | /design/menu/counterpartyAddress.json |
| html | **list** | /html/CounterpartyAddress_List.html |
| html | **view** | /html/CounterpartyAddress_View.html |
| html | **edit** | /html/CounterpartyAddress_Edit.html |
| html | **new** | /html/CounterpartyAddress_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:17:58**
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