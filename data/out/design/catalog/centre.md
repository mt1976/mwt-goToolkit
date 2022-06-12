# **Centre** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Centre** (centre) |
|Endpoint 	    |**/Centre...** [^1]|
|Endpoint Query |**Code**|
|REST API|**/API/Centre/**|
Glyph|**fas fa-map-marker-alt** ()
Friendly Name|**Centre**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Centre/CentreList) [Exportable]
* **View** (/Centre/CentreView)
* **Edit** (/Centre/CentreEdit)
* **Save** (/Centre/CentreSave)
* **New** (/Centre/CentreNew)
* **Delete** (/Centre/CentreDelete)







##  Provides
 * Lookup (Code Name)
 * Reverse Lookup (Name)





##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaCentre**
SQL Table Key | **Code**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**Code**|String|true|true|false|false|||||Y|Code||
|**Name**|String|false|true|false|false|||||Y|Name||
|**Country**|String|false|true|false|false|OL|Country|Country|Name|N|Country||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/centre.go_tmp |
| code | **adaptor** | /adaptor/centre_impl.template |
| code | **api** | /application/centre_api.go |
| code | **dao** | /dao/centre.go_tmp |
| code | **datamodel** | /datamodel/centre.go_tmp |
| code | **menu** | /design/menu/centre.json |
| html | **list** | /html/Centre_List.html |
| html | **view** | /html/Centre_View.html |
| html | **edit** | /html/Centre_Edit.html |
| html | **new** | /html/Centre_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:17:57**
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