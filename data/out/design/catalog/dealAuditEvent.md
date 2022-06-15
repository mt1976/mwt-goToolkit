# **DealAuditEvent** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**DealAuditEvent** (dealauditevent) |
|Endpoint 	    |**/DealAuditEvent...** [^1]|
|Endpoint Query |**ID**|
Glyph|**fas fa-toolbox** ()
Friendly Name|**Deal Audit History**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/DealAuditEvent/DealAuditEventList) [Exportable]
* **View** (/DealAuditEvent/DealAuditEventView)











##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **DealAuditEvent**
SQL Table Key | **InternalID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**DealRefNo**|String|true|true|false|false|||||Y|DealRefNo||
|**EventIndex**|Int|true|true|false|false|||||Y|EventIndex|0|
|**CommonRefNo**|String|false|true|false|false|||||Y|CommonRefNo||
|**Timestamp**|Time|false|true|false|false|||||Y|Timestamp||
|**UTCTimestamp**|String|false|true|false|false|||||Y|UTCTimestamp||
|**EventType**|String|false|true|false|false|||||Y|EventType||
|**Status**|String|false|true|false|false|||||Y|Status||
|**LimitOrderStatus**|String|false|true|false|false|||||Y|LimitOrderStatus||
|**Usr**|String|false|true|false|false|||||Y|Usr||
|**DealingInterface**|String|false|true|false|false|||||Y|DealingInterface||
|**SourceIP**|String|false|true|false|false|||||Y|SourceIP||
|**MessageID**|String|false|true|false|false|||||Y|MessageID||
|**Details**|String|false|true|false|false|||||Y|Details||
|**InternalId**|Int|true|true|false|false|||||Y|InternalId|0|
|**InternalDeleted**|Time|false|true|false|false|||||Y|InternalDeleted||
|**UpdatedTransactionId**|String|false|true|false|false|||||Y|UpdatedTransactionId||
|**UpdatedUserId**|String|false|true|false|false|||||Y|UpdatedUserId||
|**UpdatedDateTime**|Time|false|true|false|false|||||Y|UpdatedDateTime||
|**DeletedTransactionId**|String|false|true|false|false|||||Y|DeletedTransactionId||
|**DeletedUserId**|String|false|true|false|false|||||Y|DeletedUserId||
|**ChangeType**|String|false|true|false|false|||||Y|ChangeType||
|**DealRefNo**|String|true|true|false|false|||||Y|DealRefNo||
|**EventIndex**|Int|true|true|false|false|||||Y|EventIndex|0|
|**CommonRefNo**|String|false|true|false|false|||||Y|CommonRefNo||
|**Timestamp**|Time|false|true|false|false|||||Y|Timestamp||
|**UTCTimestamp**|String|false|true|false|false|||||Y|UTCTimestamp||
|**EventType**|String|false|true|false|false|||||Y|EventType||
|**Status**|String|false|true|false|false|||||Y|Status||
|**LimitOrderStatus**|String|false|true|false|false|||||Y|LimitOrderStatus||
|**Usr**|String|false|true|false|false|||||Y|Usr||
|**DealingInterface**|String|false|true|false|false|||||Y|DealingInterface||
|**SourceIP**|String|false|true|false|false|||||Y|SourceIP||
|**MessageID**|String|false|true|false|false|||||Y|MessageID||
|**Details**|String|false|true|false|false|||||Y|Details||
|**InternalId**|Int|true|true|false|false|||||Y|InternalId|0|
|**InternalDeleted**|Time|false|true|false|false|||||Y|InternalDeleted||
|**UpdatedTransactionId**|String|false|true|false|false|||||Y|UpdatedTransactionId||
|**UpdatedUserId**|String|false|true|false|false|||||Y|UpdatedUserId||
|**UpdatedDateTime**|Time|false|true|false|false|||||Y|UpdatedDateTime||
|**DeletedTransactionId**|String|false|true|false|false|||||Y|DeletedTransactionId||
|**DeletedUserId**|String|false|true|false|false|||||Y|DeletedUserId||
|**ChangeType**|String|false|true|false|false|||||Y|ChangeType||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/dealAuditEvent_core.go_tmp |
| code | **adaptor** | /adaptor/dealAuditEvent_impl.go_tmp |
| code | **dao** | /dao/dealAuditEvent_core.go_tmp |
| code | **datamodel** | /datamodel/dealAuditEvent_core.go_tmp |
| code | **menu** | /design/menu/dealAuditEvent.json_tmp |
| html | **list** | /html/DealAuditEvent_List.html |
| html | **view** | /html/DealAuditEvent_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **16:16:32**
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