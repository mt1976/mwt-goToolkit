# **ExternalMessage** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**ExternalMessage** (externalmessage) |
|Endpoint 	    |**/ExternalMessage...** [^1]|
|Endpoint Query |**MessageID**|
|REST API|**/API/ExternalMessage/**|
Glyph|**fas fa-mail-bulk** (text-info)
Friendly Name|**ExternalMessage**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/ExternalMessage/ExternalMessageList) [Exportable]
* **View** (/ExternalMessage/ExternalMessageView)
* **Edit** (/ExternalMessage/ExternalMessageEdit)
* **Save** (/ExternalMessage/ExternalMessageSave)

* **Delete** (/ExternalMessage/ExternalMessageDelete)







##  Provides


* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **externalMessageStore**
SQL Table Key | **MessageID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|
|**MessageID**|String|false|true|false|false|||||Y|messageID||
|**MessageFormat**|String|false|true|false|false|||||Y|messageFormat||
|**MessageDeliveredTo**|String|false|true|false|false|||||Y|messageDeliveredTo||
|**MessageBody**|String|false|true|false|false|||||Y|messageBody||
|**MessageFilename**|String|false|true|false|false|||||Y|messageFilename||
|**MessageLife**|String|false|true|false|false|||||Y|messageLife||
|**MessageDate**|String|false|true|false|false|||||Y|messageDate||
|**MessageTime**|String|false|true|false|false|||||Y|messageTime||
|**MessageTimeoutAction**|String|false|true|false|false|||||Y|messageTimeoutAction||
|**MessageACKNAK**|String|false|true|false|false|||||Y|messageACKNAK||
|**ResponseID**|String|false|true|false|false|||||Y|responseID||
|**ResponseFilename**|String|false|true|false|false|||||Y|responseFilename||
|**ResponseBody**|String|false|true|false|false|||||Y|responseBody||
|**ResponseDate**|String|false|true|false|false|||||Y|responseDate||
|**ResponseTime**|String|false|true|false|false|||||Y|responseTime||
|**ResponseAdditionalInfo**|String|false|true|false|false|||||Y|responseAdditionalInfo||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||
|**MessageTimeout**|String|false|true|false|false|||||Y|messageTimeout||
|**MessageEmitted**|String|false|true|false|false|||||Y|messageEmitted||
|**ResponseRecieved**|String|false|true|false|false|||||Y|responseRecieved||
|**MessageClass**|String|false|true|false|false|||||Y|messageClass||
|**AppID**|String|false|true|false|false|||||Y|appID||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/externalMessage.go_tmp |
| code | **api** | /application/externalMessage_api.go |
| code | **dao** | /dao/externalMessage.go_tmp |
| code | **datamodel** | /datamodel/externalMessage.go_tmp |
| code | **menu** | /design/menu/externalMessage.json |
| html | **list** | /html/ExternalMessage_List.html |
| html | **view** | /html/ExternalMessage_View.html |
| html | **edit** | /html/ExternalMessage_Edit.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:18:07**
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