# **Owner** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Owner** (owner) |
|Endpoint 	    |**/Owner...** [^1]|
|Endpoint Query |**Owner**|
Glyph|**fas fa-industry** ()
Friendly Name|**Owner**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}













##  Provides
 * Lookup (UserName FullName)
 * Reverse Lookup (FullName)





##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaOwner**
SQL Table Key | **UserName**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**UserName**|String|true|true|false|false|||||Y|UserName||
|**FullName**|String|false|true|false|false|||||Y|FullName||
|**Type**|String|false|true|false|false|||||Y|Type||
|**TradingEntity**|String|false|true|false|false|||||Y|TradingEntity||
|**DefaultEnterBook**|String|false|true|false|false|||||Y|DefaultEnterBook||
|**EmailAddress**|String|false|true|false|false|||||Y|EmailAddress||
|**Enabled**|String|false|true|false|false|||||Y|Enabled||
|**ExternalUserIds**|String|false|true|false|false|||||Y|ExternalUserIds||
|**Language**|String|false|true|false|false|||||Y|Language||
|**LocalCurrency**|String|false|true|false|false|||||Y|LocalCurrency||
|**Role**|String|false|true|false|false|||||Y|Role||
|**TelephoneNumber**|String|false|true|false|false|||||Y|TelephoneNumber||
|**TokenId**|String|false|true|false|false|||||Y|TokenId||
|**Entity**|String|false|true|false|false|||||Y|Entity||
|**UserCode**|String|false|true|false|false|||||Y|UserCode||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/owner.go_tmp |
| code | **dao** | /dao/owner.go_tmp |
| code | **datamodel** | /datamodel/owner.go_tmp |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:18:10**
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