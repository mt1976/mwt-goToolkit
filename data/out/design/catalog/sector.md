# **Sector** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Sector** (sector) |
|Endpoint 	    |**/Sector...** [^1]|
|Endpoint Query |**Sector**|
|REST API|**/API/Sector/**|
Glyph|**fas fa-industry** ()
Friendly Name|**Sector**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Sector/SectorList) [Exportable]
* **View** (/Sector/SectorView)
* **Edit** (/Sector/SectorEdit)
* **Save** (/Sector/SectorSave)
* **New** (/Sector/SectorNew)
* **Delete** (/Sector/SectorDelete)







##  Provides
 * Lookup (Code Name)
 * Reverse Lookup (Name)





##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaSector**
SQL Table Key | **Code**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**Code**|String|true|true|false|false|||||Y|Code||
|**Name**|String|false|true|false|false|||||Y|Name||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/sector.go_tmp |
| code | **adaptor** | /adaptor/sector_impl.template |
| code | **api** | /application/sector_api.go |
| code | **dao** | /dao/sector.go_tmp |
| code | **datamodel** | /datamodel/sector.go_tmp |
| code | **menu** | /design/menu/sector.json |
| html | **list** | /html/Sector_List.html |
| html | **view** | /html/Sector_View.html |
| html | **edit** | /html/Sector_Edit.html |
| html | **new** | /html/Sector_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:18:13**
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