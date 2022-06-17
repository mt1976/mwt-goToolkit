# **CounterpartyImport** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**CounterpartyImport** (counterpartyimport) |
|Endpoint 	    |**/CounterpartyImport...** [^1]|
|Endpoint Query |**ID**|
Glyph|**fas fa-passport** ()
Friendly Name|**Counterparty Import**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/CounterpartyImport/CounterpartyImportList) [Exportable]
* **View** (/CounterpartyImport/CounterpartyImportView)
* **Edit** (/CounterpartyImport/CounterpartyImportEdit)
* **Save** (/CounterpartyImport/CounterpartyImportSave)
* **New** (/CounterpartyImport/CounterpartyImportNew)
* **Delete** (/CounterpartyImport/CounterpartyImportDelete)







##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaCounterpartyImportID**
SQL Table Key | **CompID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**KeyImportID**|String|true|true|false|false|||||Y|KeyImportID||false|false|false|
|**Firm**|String|false|true|false|false|||||Y|Firm||false|false|false|
|**Centre**|String|false|true|false|false|||||Y|Centre||false|false|false|
|**FirmName**|String|true|true|false|false|||||Y|FirmName||false|false|false|
|**CentreName**|String|false|true|false|false|||||Y|CentreName||false|false|false|
|**KeyOriginID**|String|true|true|false|false|||||Y|KeyOriginID||false|false|false|
|**FullName**|String|false|true|false|false|||||Y|FullName||false|false|false|
|**CompID**|String|true|true|false|false|||||Y|CompID||false|false|false|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/counterpartyImport_core.go_tmp |
| code | **adaptor** | /adaptor/counterpartyImport_impl.go_template_tmp |
| code | **dao** | /dao/counterpartyImport_core.go_tmp |
| code | **datamodel** | /datamodel/counterpartyImport_core.go_tmp |
| code | **menu** | /design/menu/counterpartyImport.json_tmp |
| html/base | **list** | /CounterpartyImport_List.html |
| html/base | **view** | /CounterpartyImport_View.html |
| html/base | **edit** | /CounterpartyImport_Edit.html |
| html/base | **new** | /CounterpartyImport_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **16/06/2022** at **13:13:59**
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