# **Session** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Session** (session) |
|Endpoint 	    |**/Session...** [^1]|
|Endpoint Query |**SessionID**|
Glyph|**fas fa-history** (text-info)
Friendly Name|**Session**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Session/SessionList) 
* **View** (/Session/SessionView)

* **Save** (/Session/SessionSave)









##  Provides


* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sessionStore**
SQL Table Key | **Id**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|
|**Apptoken**|String|true|true|false|false|||||Y|Apptoken||
|**Createdate**|String|false|true|false|false|||||Y|Createdate||
|**Createtime**|String|false|true|false|false|||||Y|Createtime||
|**Uniqueid**|String|false|true|false|false|||||Y|Uniqueid||
|**Sessiontoken**|String|false|true|false|false|||||Y|Sessiontoken||
|**Username**|String|false|true|false|false|||||Y|Username||
|**Password**|String|false|true|false|false|||||Y|Password||
|**Userip**|String|false|true|false|false|||||Y|Userip||
|**Userhost**|String|false|true|false|false|||||Y|Userhost||
|**Appip**|String|false|true|false|false|||||Y|Appip||
|**Apphost**|String|false|true|false|false|||||Y|Apphost||
|**Issued**|String|false|true|false|false|||||Y|Issued||
|**Expiry**|String|false|true|false|false|||||Y|Expiry||
|**Expiryraw**|String|false|true|false|false|||||Y|Expiryraw||
|**Brand**|String|false|true|false|false|||||Y|Brand||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||
|**SYSWho**|String|false|true|false|false|||||NH|_who||
|**SYSHost**|String|false|true|false|false|||||NH|_host||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||
|**Id**|String|false|true|false|false|||||Y|Id||
|**Expires**|Time|false|true|false|false|||||Y|Expires||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||
|**SessionRole**|String|false|true|false|false|||||Y|SessionRole||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/session.go_tmp |
| code | **dao** | /dao/session.go_tmp |
| code | **datamodel** | /datamodel/session.go_tmp |
| code | **menu** | /design/menu/session.json |
| html | **list** | /html/Session_List.html |
| html | **view** | /html/Session_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:18:14**
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