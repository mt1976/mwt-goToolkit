# **Message** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Message** (message) |
|Endpoint 	    |**/Message...** [^1]|
|Endpoint Query |**Message**|
Glyph|**fas fa-info-circle** (text-info)
Friendly Name|**System Message**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Message/MessageList) [Exportable]
* **View** (/Message/MessageView)
* **Edit** (/Message/MessageEdit)
* **Save** (/Message/MessageSave)









##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **messageStore**
SQL Table Key | **id**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|
|**Id**|String|false|true|false|false|||||Y|Id||
|**Message**|String|false|true|false|false|||||Y|Message||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||
|**SYSWho**|String|false|true|false|false|||||NH|_who||
|**SYSHost**|String|false|true|false|false|||||NH|_host||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/message.go_tmp |
| code | **dao** | /dao/message.go_tmp |
| code | **datamodel** | /datamodel/message.go_tmp |
| code | **menu** | /design/menu/message.json |
| html | **list** | /html/Message_List.html |
| html | **view** | /html/Message_View.html |
| html | **edit** | /html/Message_Edit.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:18:10**
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