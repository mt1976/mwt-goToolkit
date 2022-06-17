# **Translation** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Translation** (translation) |
|Endpoint 	    |**/Translation...** [^1]|
|Endpoint Query |**Message**|
Glyph|**fas fa-city** ()
Friendly Name|**Message Translation**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Translation/TranslationList) [Exportable]
* **View** (/Translation/TranslationView)
* **Edit** (/Translation/TranslationEdit)
* **Save** (/Translation/TranslationSave)









##  Provides


* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **translationStore**
SQL Table Key | **Id**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|
|**Id**|String|false|true|false|true|||||N|Id||false|false|false|
|**Class**|String|true|true|false|true|||||Y|Class||true|false|false|
|**Message**|String|true|true|false|true|||||Y|Message||true|false|false|
|**Translation**|String|false|true|false|false|||||Y|Translation||false|false|false|
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|
|**SYSWho**|String|false|true|false|false|||||NH|_who||false|false|true|
|**SYSHost**|String|false|true|false|false|||||NH|_host||false|false|true|
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/translation_core.go_tmp |
| code | **dao** | /dao/translation_core.go_tmp |
| code | **datamodel** | /datamodel/translation_core.go_tmp |
| code | **menu** | /design/menu/translation.json_tmp |
| html/base | **list** | /Translation_List.html |
| html/base | **view** | /Translation_View.html |
| html/base | **edit** | /Translation_Edit.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **16/06/2022** at **13:14:05**
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