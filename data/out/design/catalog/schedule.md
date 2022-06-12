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
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|
|**Id**|String|true|true|false|false|||||Y|id||
|**Name**|String|false|true|false|false|||||Y|name||
|**Description**|String|false|true|false|false|||||Y|description||
|**Schedule**|String|false|true|false|false|||||Y|schedule||
|**Started**|String|false|true|false|false|||||Y|started||
|**Lastrun**|String|false|true|false|false|||||Y|lastrun||
|**Message**|String|false|true|false|false|||||Y|message||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||
|**SYSWho**|String|false|true|false|false|||||NH|_who||
|**SYSHost**|String|false|true|false|false|||||NH|_host||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||
|**Type**|String|false|true|false|false|||||Y|type||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||
|**Human**|String|false|true|false|false|||||Y|human||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/schedule.go_tmp |
| code | **api** | /application/schedule_api.go |
| code | **dao** | /dao/schedule.go_tmp |
| code | **datamodel** | /datamodel/schedule.go_tmp |
| code | **menu** | /design/menu/schedule.json |
| html | **list** | /html/Schedule_List.html |
| html | **view** | /html/Schedule_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:18:12**
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