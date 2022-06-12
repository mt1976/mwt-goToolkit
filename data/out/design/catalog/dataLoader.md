# **DataLoader** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**DataLoader** (dataloader) |
|Endpoint 	    |**/DataLoader...** [^1]|
|Endpoint Query |**Id**|
Glyph|**fas fa-city** ()
Friendly Name|**Data Loader**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/DataLoader/DataLoaderList) 
* **View** (/DataLoader/DataLoaderView)
* **Edit** (/DataLoader/DataLoaderEdit)
* **Save** (/DataLoader/DataLoaderSave)









##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **loaderStore**
SQL Table Key | **Id**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|
|**Id**|String|true|true|false|false|||||Y|id||
|**Name**|String|false|true|false|false|||||Y|name||
|**Description**|String|false|true|false|false|||||Y|description||
|**Filename**|String|false|true|false|false|||||Y|filename||
|**Lastrun**|String|false|true|false|false|||||Y|lastrun||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||
|**SYSWho**|String|false|true|false|false|||||NH|_who||
|**SYSHost**|String|false|true|false|false|||||NH|_host||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||
|**Type**|String|false|true|false|false|||||Y|type||
|**Instance**|String|false|true|false|false|||||Y|instance||
|**Extension**|String|false|true|false|false|||||Y|extension||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/dataLoader.go_tmp |
| code | **dao** | /dao/dataLoader.go_tmp |
| code | **datamodel** | /datamodel/dataLoader.go_tmp |
| code | **menu** | /design/menu/dataLoader.json |
| html | **list** | /html/DataLoader_List.html |
| html | **view** | /html/DataLoader_View.html |
| html | **edit** | /html/DataLoader_Edit.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:18:04**
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