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
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**KeyImportID**|String|true|true|false|false|||||Y|KeyImportID||
|**Firm**|String|false|true|false|false|||||Y|Firm||
|**Centre**|String|false|true|false|false|||||Y|Centre||
|**FirmName**|String|true|true|false|false|||||Y|FirmName||
|**CentreName**|String|false|true|false|false|||||Y|CentreName||
|**KeyOriginID**|String|true|true|false|false|||||Y|KeyOriginID||
|**FullName**|String|false|true|false|false|||||Y|FullName||
|**CompID**|String|true|true|false|false|||||Y|CompID||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/counterpartyImport.go_tmp |
| code | **adaptor** | /adaptor/counterpartyImport_impl.template |
| code | **dao** | /dao/counterpartyImport.go_tmp |
| code | **datamodel** | /datamodel/counterpartyImport.go_tmp |
| code | **menu** | /design/menu/counterpartyImport.json |
| html | **list** | /html/CounterpartyImport_List.html |
| html | **view** | /html/CounterpartyImport_View.html |
| html | **edit** | /html/CounterpartyImport_Edit.html |
| html | **new** | /html/CounterpartyImport_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:18:00**
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