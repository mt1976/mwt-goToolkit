# **AccountTransaction** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**AccountTransaction** (accounttransaction) |
|Endpoint 	    |**/AccountTransaction...** [^1]|
|Endpoint Query |**AccountNo**|
Glyph|**fas fa-landmark** ()
Friendly Name|**Account Transaction**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/AccountTransaction/AccountTransactionList) [Exportable]
* **View** (/AccountTransaction/AccountTransactionView)











##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaAccountTransactions**
SQL Table Key | **SienaReference**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**SienaReference**|String|true|true|false|false|||||Y|SienaReference||
|**LegNo**|Int|true|true|false|false|||||Y|LegNo|0|
|**MMLegNo**|Int|true|true|false|false|||||Y|MMLegNo|0|
|**Narrative**|String|false|true|false|false|||||Y|Narrative||
|**Amount**|Float|false|true|false|false|||||Y|Amount|0.00|
|**StartInterestDate**|Time|false|true|false|false|||||Y|StartInterestDate||
|**EndInterestDate**|Time|false|true|false|false|||||Y|EndInterestDate||
|**Amortisation**|Float|false|true|false|false|||||Y|Amortisation|0.00|
|**InterestAmount**|Float|false|true|false|false|||||Y|InterestAmount|0.00|
|**InterestAction**|String|false|true|false|false|||||Y|InterestAction||
|**FixingDate**|Time|false|true|false|false|||||Y|FixingDate||
|**InterestCalculationDate**|Time|false|true|false|false|||||Y|InterestCalculationDate||
|**AmendmentAmount**|Float|false|true|false|false|||||Y|AmendmentAmount|0.00|
|**DealtCcy**|String|false|true|false|false|||||Y|DealtCcy||
|**AmountDp**|Int|false|true|false|false|||||Y|AmountDp|0|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/accountTransaction.go_tmp |
| code | **adaptor** | /adaptor/accountTransaction_impl.template |
| code | **dao** | /dao/accountTransaction.go_tmp |
| code | **datamodel** | /datamodel/accountTransaction.go_tmp |
| code | **menu** | /design/menu/accountTransaction.json |
| html | **list** | /html/AccountTransaction_List.html |
| html | **view** | /html/AccountTransaction_View.html |


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