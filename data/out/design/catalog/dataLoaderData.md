# **DataLoaderData** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**DataLoaderData** (dataloaderdata) |
|Endpoint 	    |**/DataLoaderData...** [^1]|
|Endpoint Query |**Id**|
Glyph|**fas fa-city** ()
Friendly Name|**Data Loader Data**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/DataLoaderData/DataLoaderDataList) [Exportable]
* **View** (/DataLoaderData/DataLoaderDataView)
* **Edit** (/DataLoaderData/DataLoaderDataEdit)
* **Save** (/DataLoaderData/DataLoaderDataSave)









##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **loaderDataStore**
SQL Table Key | **Id**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|
|**Id**|String|true|true|false|false|||||Y|id||
|**Row**|String|false|true|false|false|||||Y|row||
|**Position**|String|false|true|false|false|||||Y|position||
|**Value**|String|false|true|false|false|||||Y|value||
|**Loader**|String|false|true|false|false|OL|DataLoader|Loader|Name|Y|loader||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||
|**SYSWho**|String|false|true|false|false|||||NH|_who||
|**SYSHost**|String|false|true|false|false|||||NH|_host||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||
|**Map**|String|false|true|false|false|||||Y|map||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/dataLoaderData.go_tmp |
| code | **dao** | /dao/dataLoaderData.go_tmp |
| code | **datamodel** | /datamodel/dataLoaderData.go_tmp |
| code | **menu** | /design/menu/dataLoaderData.json |
| html | **list** | /html/DataLoaderData_List.html |
| html | **view** | /html/DataLoaderData_View.html |
| html | **edit** | /html/DataLoaderData_Edit.html |


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