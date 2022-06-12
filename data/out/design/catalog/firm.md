# **Firm** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Firm** (firm) |
|Endpoint 	    |**/Firm...** [^1]|
|Endpoint Query |**FirmName**|
|REST API|**/API/Firm/**|
Glyph|**fas fa-city** ()
Friendly Name|**Firm**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Firm/FirmList) [Exportable]
* **View** (/Firm/FirmView)
* **Edit** (/Firm/FirmEdit)
* **Save** (/Firm/FirmSave)
* **New** (/Firm/FirmNew)
* **Delete** (/Firm/FirmDelete)







##  Provides
 * Lookup (Firm FullName)
 * Reverse Lookup (FullName)





##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaFirm**
SQL Table Key | **FirmName**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**FirmName**|String|true|true|false|false|||||Y|FirmName||
|**FullName**|String|false|true|false|false|||||Y|FullName||
|**Country**|String|false|true|false|false|OL|Country|Country|Name|N|Country||
|**Sector**|String|false|true|false|false|OL|Sector|Sector|Name|Y|Sector||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/firm.go_tmp |
| code | **adaptor** | /adaptor/firm_impl.template |
| code | **api** | /application/firm_api.go |
| code | **dao** | /dao/firm.go_tmp |
| code | **datamodel** | /datamodel/firm.go_tmp |
| code | **menu** | /design/menu/firm.json |
| html | **list** | /html/Firm_List.html |
| html | **view** | /html/Firm_View.html |
| html | **edit** | /html/Firm_Edit.html |
| html | **new** | /html/Firm_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:18:08**
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