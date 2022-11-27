# **Inbox** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Inbox** (inbox) |
|Endpoint 	    |**/Inbox...** [^1]|
|Endpoint Query |**Message**|
|REST API|**/API/Inbox/**|
Glyph|**fas fa-inbox** ()
Friendly Name|**Inbox**|
|For Project    |github.com/mt1976/ebEstimates/|

##  Actions {#action-id}
* **List** (/Inbox/InboxList) [Exportable]
* **View** (/Inbox/InboxView)



* **Delete** (/Inbox/InboxDelete)







##  Provides


* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **inboxMessages**
SQL Table Key | **MailId**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|text||
|**SYSWho**|String|false|true|false|false|||||NH|_who||false|false|true|text||
|**SYSHost**|String|false|true|false|false|||||NH|_host||false|false|true|text||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|text||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|text||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|text||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|text||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|text||
|**MailId**|String|false|true|false|false|||||Y|MailId||false|false|false|text||
|**MailTo**|String|false|true|false|false|||||Y|MailTo||false|false|false|text||
|**MailFrom**|String|false|true|false|false|||||Y|MailFrom||false|false|false|text||
|**MailSource**|String|false|true|false|false|||||Y|MailSource||false|false|false|text||
|**MailSent**|String|false|true|false|false|||||Y|MailSent||false|false|false|text||
|**MailUnread**|String|false|true|false|false|LL|tf|||Y|MailUnread||false|false|false|text||
|**MailSubject**|String|false|true|false|false|||||Y|MailSubject||false|false|false|text||
|**MailContent**|String|false|true|false|false|||||Y|MailContent||false|false|false|text||
|**MailAcknowledged**|String|false|true|false|false|LL|tf|||Y|MailAcknowledged||false|false|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/inbox_core.go_tmp |
| code | **api** | /application/inbox_api.go_tmp |
| code | **dao** | /dao/inbox_core.go_tmp |
| code | **datamodel** | /datamodel/inbox_core.go_tmp |
| code | **menu** | /design/menu/inbox.json_tmp |
| html | **list** | /Inbox_List.html |
| html | **view** | /Inbox_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **24/06/2022** at **09:39:16**
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