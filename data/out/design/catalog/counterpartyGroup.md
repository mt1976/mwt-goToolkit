# **CounterpartyGroup** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**CounterpartyGroup** (counterpartygroup) |
|Endpoint 	    |**/CounterpartyGroup...** [^1]|
|Endpoint Query |**Group**|
Glyph|**fa fa-users** ()
Friendly Name|**Counterparty Group**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/CounterpartyGroup/CounterpartyGroupList) [Exportable]
* **View** (/CounterpartyGroup/CounterpartyGroupView)
* **Edit** (/CounterpartyGroup/CounterpartyGroupEdit)
* **Save** (/CounterpartyGroup/CounterpartyGroupSave)
* **New** (/CounterpartyGroup/CounterpartyGroupNew)
* **Delete** (/CounterpartyGroup/CounterpartyGroupDelete)







##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaCounterpartyGroup**
SQL Table Key | **Name**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**Name**|String|true|true|false|false|||||Y|Name||
|**CountryCode**|String|false|true|false|false|||||Y|CountryCode||
|**SuperGroup**|String|false|true|false|false|||||Y|SuperGroup||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/counterpartyGroup.go_tmp |
| code | **adaptor** | /adaptor/counterpartyGroup_impl.template |
| code | **dao** | /dao/counterpartyGroup.go_tmp |
| code | **datamodel** | /datamodel/counterpartyGroup.go_tmp |
| code | **menu** | /design/menu/counterpartyGroup.json |
| html | **list** | /html/CounterpartyGroup_List.html |
| html | **view** | /html/CounterpartyGroup_View.html |
| html | **edit** | /html/CounterpartyGroup_Edit.html |
| html | **new** | /html/CounterpartyGroup_New.html |


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