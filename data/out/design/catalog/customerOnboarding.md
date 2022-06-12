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







##  Data Source 
|   |   |
|---|---|

Fetch|<ul><li>**Implement in Adaptor**</li><li> func CustomerOnboarding_GetList_impl() (int, []dm.CustomerOnboarding, error) {return 0, nil, nil}</li><li>func CustomerOnboarding_GetByID_impl(id string) (int, dm.CustomerOnboarding, error) {return 0, dm.CustomerOnboarding{}, nil}</li></ul>
Store|<ul><li>**Implement in Adaptor**</li><li>func CustomerOnboarding_NewID_impl(rec dm.CustomerOnboarding) (string) { return rec.ID } </li><li>func CustomerOnboarding_Delete_impl(id string) (error) {return nil}</li><li>func CustomerOnboarding_Update_impl(id string,rec dm.CustomerOnboarding, usr string) (error) {return nil}</li></ul>

##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**ID**|String|true|true|false|false|||||Y|ID||
|**CustomerName**|String|true|true|false|false|||||Y|CustomerName||
|**CustomerAddress**|String|false|true|false|false|||||Y|CustomerAddress||
|**CustomerBOLID**|String|false|true|false|false|||||Y|CustomerBOLID||
|**CustomerFirmName**|String|false|true|false|false|||||Y|CustomerFirmName||
|**CustomerType**|String|false|true|false|false|||||Y|CustomerType||
|**CustomerRDC**|String|false|true|false|false|||||Y|CustomerRDC||
|**CustomerSortCode**|String|false|true|false|false|||||Y|CustomerSortCode||
|**CustomerGMClientNo**|String|false|true|false|false|||||Y|CustomerGMClientNo||
|**CustomerDefaultBook**|String|false|true|false|false|||||Y|CustomerDefaultBook||
|**CustomerRegion**|String|false|true|false|false|||||Y|CustomerRegion||
|**CustomerCategory**|String|false|true|false|false|||||Y|CustomerCategory||
|**CustomerTelephoneNo**|String|false|true|false|false|||||Y|CustomerTelephoneNo||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/customerOnboarding.go_tmp |
| code | **adaptor** | /adaptor/customerOnboarding_impl.template |
| code | **api** | /application/customerOnboarding_api.go |
| code | **dao** | /dao/customerOnboarding.go_tmp |
| code | **datamodel** | /datamodel/customerOnboarding.go_tmp |
| code | **menu** | /design/menu/customerOnboarding.json |
| html | **list** | /html/CustomerOnboarding_List.html |
| html | **view** | /html/CustomerOnboarding_View.html |
| html | **edit** | /html/CustomerOnboarding_Edit.html |
| html | **new** | /html/CustomerOnboarding_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:17:51**
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