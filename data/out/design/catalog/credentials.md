# **Credentials** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Credentials** (credentials) |
|Endpoint 	    |**/Credentials...** [^1]|
|Endpoint Query |**Id**|
Glyph|**fas fa-key** (text-danger)
Friendly Name|**Credential**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Credentials/CredentialsList) [Exportable]
* **View** (/Credentials/CredentialsView)
* **Edit** (/Credentials/CredentialsEdit)
* **Save** (/Credentials/CredentialsSave)
* **New** (/Credentials/CredentialsNew)
* **Delete** (/Credentials/CredentialsDelete)







##  Provides


* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **credentialsStore**
SQL Table Key | **id**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|
|**Id**|String|true|true|false|false|||||Y|Id||
|**Username**|String|true|true|false|false|||||Y|Username||
|**Password**|String|false|true|false|false|||||Y|Password||
|**Firstname**|String|false|true|false|false|||||Y|Firstname||
|**Lastname**|String|false|true|false|false|||||Y|Lastname||
|**Knownas**|String|false|true|false|false|||||Y|Knownas||
|**Email**|String|false|true|false|false|||||Y|Email||
|**Issued**|String|false|true|false|false|||||Y|Issued||
|**Expiry**|String|false|true|false|false|||||Y|Expiry||
|**RoleType**|String|false|true|false|false|||||Y|RoleType||
|**Brand**|String|false|true|false|false|||||Y|Brand||
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
| code | **application** | /application/credentials.go_tmp |
| code | **dao** | /dao/credentials.go_tmp |
| code | **datamodel** | /datamodel/credentials.go_tmp |
| code | **menu** | /design/menu/credentials.json |
| html | **list** | /html/Credentials_List.html |
| html | **view** | /html/Credentials_View.html |
| html | **edit** | /html/Credentials_Edit.html |
| html | **new** | /html/Credentials_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:18:02**
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