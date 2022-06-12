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
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|
|**Id**|String|false|true|false|false|||||Y|Id||
|**Class**|String|false|true|false|false|||||Y|Class||
|**Message**|String|false|true|false|false|||||Y|Message||
|**Translation**|String|false|true|false|false|||||Y|Translation||
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
| code | **application** | /application/translation.go_tmp |
| code | **dao** | /dao/translation.go_tmp |
| code | **datamodel** | /datamodel/translation.go_tmp |
| code | **menu** | /design/menu/translation.json |
| html | **list** | /html/Translation_List.html |
| html | **view** | /html/Translation_View.html |
| html | **edit** | /html/Translation_Edit.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:18:16**
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