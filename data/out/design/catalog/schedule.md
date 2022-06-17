# **Schedule** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Schedule** (schedule) |
|Endpoint 	    |**/Schedule...** [^1]|
|Endpoint Query |**Schedule**|
|REST API|**/API/Schedule/**|
Glyph|**far fa-clock** (text-info)
Friendly Name|**Scheduler**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Schedule/ScheduleList) [Exportable]
* **View** (/Schedule/ScheduleView)

* **Save** (/Schedule/ScheduleSave)









##  Provides


* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **scheduleStore**
SQL Table Key | **id**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|
|**Id**|String|true|true|false|false|||||Y|id||false|false|false|
|**Name**|String|false|true|false|false|||||Y|name||false|false|false|
|**Description**|String|false|true|false|false|||||Y|description||false|false|false|
|**Schedule**|String|false|true|false|false|||||Y|schedule||false|false|false|
|**Started**|String|false|true|false|false|||||Y|started||false|false|false|
|**Lastrun**|String|false|true|false|false|||||Y|lastrun||false|false|false|
|**Message**|String|false|true|false|false|||||Y|message||false|false|false|
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|
|**SYSWho**|String|false|true|false|false|||||NH|_who||false|false|true|
|**SYSHost**|String|false|true|false|false|||||NH|_host||false|false|true|
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|
|**Type**|String|false|true|false|false|||||Y|type||false|false|false|
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|
|**Human**|String|false|true|false|false|||||Y|human||false|false|false|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/schedule_core.go_tmp |
| code | **api** | /application/schedule_api.go_tmp |
| code | **dao** | /dao/schedule_core.go_tmp |
| code | **datamodel** | /datamodel/schedule_core.go_tmp |
| code | **menu** | /design/menu/schedule.json_tmp |
| html/base | **list** | /Schedule_List.html |
| html/base | **view** | /Schedule_View.html |


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