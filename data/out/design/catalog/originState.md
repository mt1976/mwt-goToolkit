# **OriginState** - Object Definition
##  Information
| Information  | Value  |
|---|---|
|Object         |**OriginState** (originstate) |
|Endpoint 	    |**/OriginState...** [^1]|
|Endpoint Query |**OriginStateID**|
|REST API|**/API/OriginState/**|
Glyph|**fas fa-cogs** ()
Friendly Name|**Origin State**|
|For Project    |github.com/mt1976/ebEstimates/|

##  Actions {#action-id}
* **List** (/OriginState/OriginStateList) [Exportable]
* **View** (/OriginState/OriginStateView)
* **Edit** (/OriginState/OriginStateEdit)
* **Save** (/OriginState/OriginStateSave)
* **New** (/OriginState/OriginStateNew)








##  Provides
 * Lookup (Code Name)

* Auditing 




##  Data Source 
| Information  | Value  |
|---|---|
SQL Table Name       | **originStateStore**
SQL Table Key | **originStateID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**OriginStateID**|String|false|true|false|true|||||H|originStateID||true|false|false|text||
|**Code**|String|true|true|false|false|||||Y|code||false|false|false|text||
|**Name**|String|true|true|false|false|||||Y|name||false|false|false|text||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|text||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|text||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|text||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|text||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|text||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|text||
|**SYSDeleted**|String|false|true|false|false|||||NH|_deleted||false|false|true|text||
|**SYSDeletedBy**|String|false|true|false|false|||||NH|_deletedBy||false|false|true|text||
|**SYSDeletedHost**|String|false|true|false|false|||||NH|_deletedHost||false|false|true|text||
|**SYSDbVersion**|String|false|true|false|false|||||NH|_dbVersion||false|false|true|text||
|**IsLocked**|String|false|true|false|false|||||Y|isLocked||false|false|false|text||
|**Notify**|String|false|true|false|false|||||Y|notify||false|false|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/originState_core.go_tmp |
| code | **adaptor** | /dao/originState_adaptor.go_template_tmp |
| code | **api** | /application/originState_api.go_tmp |
| code | **dao** | /dao/originState_core.go_tmp |
| code | **datamodel** | /datamodel/originState_core.go_tmp |
| code | **menu** | /design/menu/originState.json_tmp |
| html | **list** | /OriginState_List.html |
| html | **view** | /OriginState_View.html |
| html | **edit** | /OriginState_Edit.html |
| html | **new** | /OriginState_New.html |


## Audit Information
| Information  | Value |
|---|---|
Template Generator Version   | **Einsteinium [r5-23.01.23]**
Date & Time		     | **23/01/2023** at **11:31:03**
Who & Where		     | **matttownsend (Matt Townsend)** on **silicon.local**

---
### Footnotes
[^1]: **Endpoint**
    * The full list of endpoints can be found in the [Actions](#action-id) section
[^2]: **Lookup Fields**
    * LL = A List Lookup. Define list in lits.cfg
    * OL = An Object Lookup. Get a list of values from an Object
    * FL = Fetches 1 value from an object based on the content of the field. 
    * ∀ = This lookup has a filter that can be defined in the Data Object
[^3]: **Inputtable**   
    * H = Hidden Field
    * N = No Input Field
    * Y = Inputtable Field