# **CounterpartyName** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**CounterpartyName** (counterpartyname) |
|Endpoint 	    |**/CounterpartyName...** [^1]|
|Endpoint Query |**ID**|
Glyph|**fas fa-passport** ()
Friendly Name|**Counterparty Name**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/CounterpartyName/CounterpartyNameList) [Exportable]
* **View** (/CounterpartyName/CounterpartyNameView)











##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaCounterpartyNameLookup**
SQL Table Key | **CompID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**NameFirm**|String|true|true|false|false|||||Y|NameFirm||
|**NameCentre**|String|true|true|false|false|||||Y|NameCentre||
|**FullName**|String|false|true|false|false|||||Y|FullName||
|**CompID**|String|true|true|false|false|||||Y|CompID||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/counterpartyName.go_tmp |
| code | **dao** | /dao/counterpartyName.go_tmp |
| code | **datamodel** | /datamodel/counterpartyName.go_tmp |
| code | **menu** | /design/menu/counterpartyName.json |
| html | **list** | /html/CounterpartyName_List.html |
| html | **view** | /html/CounterpartyName_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:18:01**
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