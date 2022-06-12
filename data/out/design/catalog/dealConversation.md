# **DealConversation** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**DealConversation** (dealconversation) |
|Endpoint 	    |**/DealConversation...** [^1]|
|Endpoint Query |**ID**|
Glyph|**fas fa-exchange-alt** ()
Friendly Name|**Deal Conversation**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/DealConversation/DealConversationList) [Exportable]
* **View** (/DealConversation/DealConversationView)











##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaDealConversationLog**
SQL Table Key | **MessageLogReference**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**SienaReference**|String|false|true|false|false|||||Y|SienaReference||
|**Status**|String|false|true|false|false|||||Y|Status||
|**MessageType**|String|false|true|false|false|||||Y|MessageType||
|**ContractNumber**|String|false|true|false|false|||||Y|ContractNumber||
|**AckReference**|String|false|true|false|false|||||Y|AckReference||
|**NewTX**|Bool|false|true|false|false|||||Y|NewTX|True|
|**LegNo**|Int|false|true|false|false|||||Y|LegNo|0|
|**Summary**|String|false|true|false|false|||||Y|Summary||
|**BusinessDate**|Time|false|true|false|false|||||Y|BusinessDate||
|**TXNo**|Int|false|true|false|false|||||Y|TXNo|0|
|**ExternalSystem**|String|false|true|false|false|||||Y|ExternalSystem||
|**MessageLogReference**|String|true|true|false|false|||||Y|MessageLogReference||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/dealConversation.go_tmp |
| code | **dao** | /dao/dealConversation.go_tmp |
| code | **datamodel** | /datamodel/dealConversation.go_tmp |
| code | **menu** | /design/menu/dealConversation.json |
| html | **list** | /html/DealConversation_List.html |
| html | **view** | /html/DealConversation_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:18:05**
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