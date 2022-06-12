# **Portfolio** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Portfolio** (portfolio) |
|Endpoint 	    |**/Portfolio...** [^1]|
|Endpoint Query |**Code**|
|REST API|**/API/Portfolio/**|
Glyph|**fas fa-plug** ()
Friendly Name|**Bank Porfolio**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Portfolio/PortfolioList) [Exportable]
* **View** (/Portfolio/PortfolioView)
* **Edit** (/Portfolio/PortfolioEdit)
* **Save** (/Portfolio/PortfolioSave)
* **New** (/Portfolio/PortfolioNew)
* **Delete** (/Portfolio/PortfolioDelete)







##  Provides

 * Reverse Lookup (Description1)





##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaPortfolio**
SQL Table Key | **Code**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**Code**|String|true|true|false|false|||||Y|Code||
|**Description1**|String|false|true|false|false|||||Y|Description1||
|**Description2**|String|false|true|false|false|||||Y|Description2||
|**IsDefault**|Bool|false|true|false|false|||||Y|isDefault|True|
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
| code | **application** | /application/portfolio.go_tmp |
| code | **adaptor** | /adaptor/portfolio_impl.template |
| code | **api** | /application/portfolio_api.go |
| code | **dao** | /dao/portfolio.go_tmp |
| code | **datamodel** | /datamodel/portfolio.go_tmp |
| code | **menu** | /design/menu/portfolio.json |
| html | **list** | /html/Portfolio_List.html |
| html | **view** | /html/Portfolio_View.html |
| html | **edit** | /html/Portfolio_Edit.html |
| html | **new** | /html/Portfolio_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:18:11**
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