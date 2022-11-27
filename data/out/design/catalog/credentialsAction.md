# **CredentialsAction** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**CredentialsAction** (credentialsaction) |
|Endpoint 	    |**/CredentialsAction...** [^1]|
|Endpoint Query |**ID**|
Glyph|**fas fa-user-cog** ()
Friendly Name|**Credentials Update**|
|For Project    |github.com/mt1976/ebEstimates/|

##  Actions {#action-id}

* **View** (/CredentialsAction/CredentialsActionView)
* **Edit** (/CredentialsAction/CredentialsActionEdit)
* **Save** (/CredentialsAction/CredentialsActionSave)
* **New** (/CredentialsAction/CredentialsActionNew)








##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **inboxMessages**
SQL Table Key | **MailId**
Fetch|<ul><li>**Implement in Adaptor**</li><li> func CredentialsAction_GetList_impl() (int, []dm.CredentialsAction, error) {return 0, nil, nil}</li><li>func CredentialsAction_GetByID_impl(id string) (int, dm.CredentialsAction, error) {return 0, dm.CredentialsAction{}, nil}</li></ul>
Store|<ul><li>**Implement in Adaptor**</li><li>func CredentialsAction_NewID_impl(rec dm.CredentialsAction) (string) { return rec.ID } </li><li>func CredentialsAction_Delete_impl(id string) (error) {return nil}</li><li>func CredentialsAction_Update_impl(id string,rec dm.CredentialsAction, usr string) (error) {return nil}</li></ul>

##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**ID**|String|true|true|false|false|||||Y|ID||false|false|false|text||
|**User**|String|true|true|false|false|OL|Credentials|Credentials_Username|Credentials_Id|Y|User||false|false|false|text||
|**OldPassword**|String|true|true|false|true|||||Y|OldPassword||false|false|false|||
|**NewPassword**|String|true|true|false|true|||||Y|NewPassword||false|false|false|||
|**ConfirmPassword**|String|true|true|false|true|||||Y|ConfirmPassword||false|false|false|||
|**Action**|String|true|true|false|false|LL|credentialStates|||Y|Action||false|false|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/credentialsAction_core.go_tmp |
| code | **adaptor** | /adaptor/credentialsAction_impl.go_template_tmp |
| code | **dao** | /dao/credentialsAction_core.go_tmp |
| code | **datamodel** | /datamodel/credentialsAction_core.go_tmp |
| code | **menu** | /design/menu/credentialsAction.json_tmp |
| html | **view** | /CredentialsAction_View.html |
| html | **edit** | /CredentialsAction_Edit.html |
| html | **new** | /CredentialsAction_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **24/06/2022** at **09:57:47**
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