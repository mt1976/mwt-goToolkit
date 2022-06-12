# **Test1** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Test1** (test1) |
|Endpoint 	    |**/Test1...** [^1]|
|Endpoint Query |**ID**|
|REST API|**/API/Test1/**|
Glyph|**fas fa-cubes** (text-info)
Friendly Name|**Test1Template**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Test1/Test1List) [Exportable]
* **View** (/Test1/Test1View)
* **Edit** (/Test1/Test1Edit)
* **Save** (/Test1/Test1Save)
* **New** (/Test1/Test1New)
* **Delete** (/Test1/Test1Delete)







##  Provides
 * Lookup (ID Descr)
 * Reverse Lookup (name_fisherRLU)

* Monitoring 



##  Data Source 
|   |   |
|---|---|

Fetch|<ul><li>**Implement in Adaptor**</li><li> func Test1_GetList_impl() (int, []dm.Test1, error) {return 0, nil, nil}</li><li>func Test1_GetByID_impl(id string) (int, dm.Test1, error) {return 0, dm.Test1{}, nil}</li></ul>
Store|<ul><li>**Implement in Adaptor**</li><li>func Test1_NewID_impl(rec dm.Test1) (string) { return rec.ID } </li><li>func Test1_Delete_impl(id string) (error) {return nil}</li><li>func Test1_Update_impl(id string,rec dm.Test1, usr string) (error) {return nil}</li></ul>

##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**ID**|String|true|true|false|false|||||Y|ID||
|**Endpoint**|String|true|true|false|false|||||Y|Endpoint||
|**Descr**|String|false|true|false|true|||||N|Descr||
|**Query**|String|true|true|false|false|||||Y|Query||
|**Source**|String|true|true|false|false|||||Y|Source||
|**Firm**|String|false|true|false|false|OL|Firm|Firmx|FullName|N|Firm||
|**YN**|String|false|true|false|false|LL||yn||Y|YN||
|**User**|String|false|true|false|false|LL||users||Y|User||
|**Cheese**|String|false|false|true|false|||||Y|||
|**Onion**|String|true|false|true|false|||||Y|||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/test1.go_tmp |
| code | **adaptor** | /adaptor/test1_impl.template |
| code | **api** | /application/test1_api.go |
| code | **dao** | /dao/test1.go_tmp |
| code | **datamodel** | /datamodel/test1.go_tmp |
| code | **job** | /jobs/test1_core.go |
| code | **menu** | /design/menu/test1.json |
| html | **list** | /html/Test1_List.html |
| html | **view** | /html/Test1_View.html |
| html | **edit** | /html/Test1_Edit.html |
| html | **new** | /html/Test1_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:18:15**
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