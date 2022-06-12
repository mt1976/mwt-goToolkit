# **Product** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Product** (product) |
|Endpoint 	    |**/Product...** [^1]|
|Endpoint Query |**Code**|
|REST API|**/API/Product/**|
Glyph|**fas fa-object-group** ()
Friendly Name|**Product Group**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Product/ProductList) [Exportable]
* **View** (/Product/ProductView)
* **Edit** (/Product/ProductEdit)
* **Save** (/Product/ProductSave)
* **New** (/Product/ProductNew)
* **Delete** (/Product/ProductDelete)







##  Provides

 * Reverse Lookup (Name)





##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaProduct**
SQL Table Key | **Code**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**Code**|String|true|true|false|false|||||Y|Code||
|**Name**|String|false|true|false|false|||||Y|Name||
|**Factor**|Float|false|true|false|false|||||Y|Factor|0.00|
|**MaxTermPrecedence**|Bool|false|true|false|false|||||Y|MaxTermPrecedence|True|
|**InternalId**|Int|true|true|false|false|||||Y|InternalId|0|
|**InternalDeleted**|Time|false|true|false|false|||||Y|InternalDeleted||
|**UpdatedTransactionId**|String|false|true|false|false|||||Y|UpdatedTransactionId||
|**UpdatedUserId**|String|false|true|false|false|||||Y|UpdatedUserId||
|**UpdatedDateTime**|Time|false|true|false|false|||||Y|UpdatedDateTime||
|**DeletedTransactionId**|String|false|true|false|false|||||Y|DeletedTransactionId||
|**DeletedUserId**|String|false|true|false|false|||||Y|DeletedUserId||
|**ChangeType**|String|false|true|false|false|||||Y|ChangeType||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/product.go_tmp |
| code | **adaptor** | /adaptor/product_impl.template |
| code | **api** | /application/product_api.go |
| code | **dao** | /dao/product.go_tmp |
| code | **datamodel** | /datamodel/product.go_tmp |
| code | **menu** | /design/menu/product.json |
| html | **list** | /html/Product_List.html |
| html | **view** | /html/Product_View.html |
| html | **edit** | /html/Product_Edit.html |
| html | **new** | /html/Product_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:18:12**
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