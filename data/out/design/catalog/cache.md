# **Cache** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Cache** (cache) |
|Endpoint 	    |**/Cache...** [^1]|
|Endpoint Query |**ID**|
Glyph|**fas fa-database** (text-info)
Friendly Name|**Cache Content**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Cache/CacheList) [Exportable]
* **View** (/Cache/CacheView)
* **Edit** (/Cache/CacheEdit)
* **Save** (/Cache/CacheSave)









##  Provides


* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **cacheStore**
SQL Table Key | **Id**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|
|**Id**|String|true|true|false|false|||||Y|id||
|**Object**|String|false|true|false|false|||||Y|object||
|**Field**|String|false|true|false|false|||||Y|field||
|**Value**|String|false|true|false|false|||||Y|value||
|**Expiry**|String|false|true|false|false|||||Y|expiry||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||
|**SYSWho**|String|false|true|false|false|||||NH|_who||
|**SYSHost**|String|false|true|false|false|||||NH|_host||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||
|**Source**|String|false|true|false|false|||||Y|source||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/cache.go_tmp |
| code | **dao** | /dao/cache.go_tmp |
| code | **datamodel** | /datamodel/cache.go_tmp |
| code | **menu** | /design/menu/cache.json |
| html | **list** | /html/Cache_List.html |
| html | **view** | /html/Cache_View.html |
| html | **edit** | /html/Cache_Edit.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:17:55**
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