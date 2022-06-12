# **Systems** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Systems** (systems) |
|Endpoint 	    |**/Systems...** [^1]|
|Endpoint Query |**ID**|
Glyph|**fas fa-plug** (text-info)
Friendly Name|**Connected System**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Systems/SystemsList) [Exportable]
* **View** (/Systems/SystemsView)
* **Edit** (/Systems/SystemsEdit)
* **Save** (/Systems/SystemsSave)
* **New** (/Systems/SystemsNew)
* **Delete** (/Systems/SystemsDelete)







##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **systemStore**
SQL Table Key | **Id**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|
|**Id**|String|true|true|false|false|||||Y|Id||
|**Name**|String|false|true|false|false|||||Y|Name||
|**Staticin**|String|false|true|false|false|||||Y|Staticin||
|**Staticout**|String|false|true|false|false|||||Y|Staticout||
|**Txnin**|String|false|true|false|false|||||Y|Txnin||
|**Txnout**|String|false|true|false|false|||||Y|Txnout||
|**Fundscheckin**|String|false|true|false|false|||||Y|Fundscheckin||
|**Fundscheckout**|String|false|true|false|false|||||Y|Fundscheckout||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||
|**SYSWho**|String|false|true|false|false|||||NH|_who||
|**SYSHost**|String|false|true|false|false|||||NH|_host||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||
|**SWIFTin**|String|false|true|false|false|||||Y|SWIFTin||
|**SWIFTout**|String|false|true|false|false|||||Y|SWIFTout||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/systems.go_tmp |
| code | **dao** | /dao/systems.go_tmp |
| code | **datamodel** | /datamodel/systems.go_tmp |
| code | **menu** | /design/menu/systems.json |
| html | **list** | /html/Systems_List.html |
| html | **view** | /html/Systems_View.html |
| html | **edit** | /html/Systems_Edit.html |
| html | **new** | /html/Systems_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:18:15**
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