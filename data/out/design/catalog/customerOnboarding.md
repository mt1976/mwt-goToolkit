# **CustomerOnboarding** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**CustomerOnboarding** (customeronboarding) |
|Endpoint 	    |**/CustomerOnboarding...** [^1]|
|Endpoint Query |**ID**|
|REST API|**/API/CustomerOnboarding/**|
Glyph|**fas fa-tasks** (text-secondary)
Friendly Name|**Customer Onboarding Tool**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/CustomerOnboarding/CustomerOnboardingList) [Exportable]
* **View** (/CustomerOnboarding/CustomerOnboardingView)
* **Edit** (/CustomerOnboarding/CustomerOnboardingEdit)
* **Save** (/CustomerOnboarding/CustomerOnboardingSave)
* **New** (/CustomerOnboarding/CustomerOnboardingNew)
* **Delete** (/CustomerOnboarding/CustomerOnboardingDelete)







##  Provides



* Monitoring 



##  Data Source 
|   |   |
|---|---|

Fetch|<ul><li>**Implement in Adaptor**</li><li> func CustomerOnboarding_GetList_impl() (int, []dm.CustomerOnboarding, error) {return 0, nil, nil}</li><li>func CustomerOnboarding_GetByID_impl(id string) (int, dm.CustomerOnboarding, error) {return 0, dm.CustomerOnboarding{}, nil}</li></ul>
Store|<ul><li>**Implement in Adaptor**</li><li>func CustomerOnboarding_NewID_impl(rec dm.CustomerOnboarding) (string) { return rec.ID } </li><li>func CustomerOnboarding_Delete_impl(id string) (error) {return nil}</li><li>func CustomerOnboarding_Update_impl(id string,rec dm.CustomerOnboarding, usr string) (error) {return nil}</li></ul>

##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**ID**|String|true|true|false|false|||||Y|ID||false|false|false|
|**CustomerName**|String|true|true|false|false|||||Y|CustomerName||false|false|false|
|**CustomerAddress**|String|false|true|false|false|||||Y|CustomerAddress||false|false|false|
|**CustomerBOLID**|String|false|true|false|false|||||Y|CustomerBOLID||false|false|false|
|**CustomerFirmName**|String|false|true|false|false|||||Y|CustomerFirmName||false|false|false|
|**CustomerType**|String|false|true|false|false|||||Y|CustomerType||false|false|false|
|**CustomerRDC**|String|false|true|false|false|||||Y|CustomerRDC||false|false|false|
|**CustomerSortCode**|String|false|true|false|false|||||Y|CustomerSortCode||false|false|false|
|**CustomerGMClientNo**|String|false|true|false|false|||||Y|CustomerGMClientNo||false|false|false|
|**CustomerDefaultBook**|String|false|true|false|false|||||Y|CustomerDefaultBook||false|false|false|
|**CustomerRegion**|String|false|true|false|false|||||Y|CustomerRegion||false|false|false|
|**CustomerCategory**|String|false|true|false|false|||||Y|CustomerCategory||false|false|false|
|**CustomerTelephoneNo**|String|false|true|false|false|||||Y|CustomerTelephoneNo||false|false|false|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/customerOnboarding_core.go_tmp |
| code | **adaptor** | /adaptor/customerOnboarding_impl.go_template_tmp |
| code | **api** | /application/customerOnboarding_api.go_tmp |
| code | **dao** | /dao/customerOnboarding_core.go_tmp |
| code | **datamodel** | /datamodel/customerOnboarding_core.go_tmp |
| code | **menu** | /design/menu/customerOnboarding.json_tmp |
| html/base | **list** | /CustomerOnboarding_List.html |
| html/base | **view** | /CustomerOnboarding_View.html |
| html/base | **edit** | /CustomerOnboarding_Edit.html |
| html/base | **new** | /CustomerOnboarding_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **16/06/2022** at **13:13:55**
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