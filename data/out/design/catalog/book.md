# **Book** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Book** (book) |
|Endpoint 	    |**/Book...** [^1]|
|Endpoint Query |**Book**|
|REST API|**/API/Book/**|
Glyph|**fas fa-book** ()
Friendly Name|**Book**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Book/BookList) [Exportable]
* **View** (/Book/BookView)
* **Edit** (/Book/BookEdit)
* **Save** (/Book/BookSave)

* **Delete** (/Book/BookDelete)







##  Provides

 * Reverse Lookup (FullName)





##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaBook**
SQL Table Key | **BookName**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**BookName**|String|true|true|false|false|||||Y|BookName||
|**FullName**|String|false|true|false|false|||||Y|FullName||
|**PLManage**|String|false|true|false|false|||||Y|PLManage||
|**PLTransfer**|String|false|true|false|false|||||Y|PLTransfer||
|**DerivePL**|Bool|false|true|false|false|||||Y|DerivePL|True|
|**CostOfCarry**|Bool|false|true|false|false|||||Y|CostOfCarry|True|
|**CostOfFunding**|Bool|false|true|false|false|||||Y|CostOfFunding|True|
|**LotAllocationMethod**|String|false|true|false|false|||||Y|LotAllocationMethod||
|**InternalId**|Int|true|true|false|false|||||Y|InternalId|0|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/book.go_tmp |
| code | **adaptor** | /adaptor/book_impl.template |
| code | **api** | /application/book_api.go |
| code | **dao** | /dao/book.go_tmp |
| code | **datamodel** | /datamodel/book.go_tmp |
| code | **menu** | /design/menu/book.json |
| html | **list** | /html/Book_List.html |
| html | **view** | /html/Book_View.html |
| html | **edit** | /html/Book_Edit.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:17:54**
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