# **CounterpartyCreditRating** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**CounterpartyCreditRating** (counterpartycreditrating) |
|Endpoint 	    |**/CounterpartyCreditRating...** [^1]|
|Endpoint Query |**ID**|
Glyph|**fas fa-passport** ()
Friendly Name|**Counterparty Credit Rating**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/CounterpartyCreditRating/CounterpartyCreditRatingList) [Exportable]
* **View** (/CounterpartyCreditRating/CounterpartyCreditRatingView)
* **Edit** (/CounterpartyCreditRating/CounterpartyCreditRatingEdit)
* **Save** (/CounterpartyCreditRating/CounterpartyCreditRatingSave)
* **New** (/CounterpartyCreditRating/CounterpartyCreditRatingNew)
* **Delete** (/CounterpartyCreditRating/CounterpartyCreditRatingDelete)







##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaCounterpartyCreditRating**
SQL Table Key | **CompID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**NameFirm**|String|true|true|false|false|||||Y|NameFirm||
|**NameCentre**|String|true|true|false|false|||||Y|NameCentre||
|**CreditRatingUsage**|String|true|true|false|false|||||Y|CreditRatingUsage||
|**CreditRatingAgency**|String|false|true|false|false|||||Y|CreditRatingAgency||
|**CreditRatingName**|String|false|true|false|false|||||Y|CreditRatingName||
|**CompID**|String|true|true|false|false|||||Y|CompID||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/counterpartyCreditRating.go_tmp |
| code | **adaptor** | /adaptor/counterpartyCreditRating_impl.template |
| code | **dao** | /dao/counterpartyCreditRating.go_tmp |
| code | **datamodel** | /datamodel/counterpartyCreditRating.go_tmp |
| code | **menu** | /design/menu/counterpartyCreditRating.json |
| html | **list** | /html/CounterpartyCreditRating_List.html |
| html | **view** | /html/CounterpartyCreditRating_View.html |
| html | **edit** | /html/CounterpartyCreditRating_Edit.html |
| html | **new** | /html/CounterpartyCreditRating_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:17:59**
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