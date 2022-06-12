# **Counterparty** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Counterparty** (counterparty) |
|Endpoint 	    |**/Counterparty...** [^1]|
|Endpoint Query |**ID**|
|REST API|**/API/Counterparty/**|
Glyph|**fas fa-user** ()
Friendly Name|**Counterparty**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Counterparty/CounterpartyList) [Exportable]
* **View** (/Counterparty/CounterpartyView)
* **Edit** (/Counterparty/CounterpartyEdit)
* **Save** (/Counterparty/CounterpartySave)

* **Delete** (/Counterparty/CounterpartyDelete)







##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaCounterparty**
SQL Table Key | **CompID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**NameCentre**|String|true|true|false|false|||||Y|NameCentre||
|**NameFirm**|String|true|true|false|false|||||Y|NameFirm||
|**FullName**|String|false|true|false|false|||||Y|FullName||
|**TelephoneNumber**|String|false|true|false|false|||||Y|TelephoneNumber||
|**EmailAddress**|String|false|true|false|false|||||Y|EmailAddress||
|**CustomerType**|String|false|true|false|false|||||Y|CustomerType||
|**AccountOfficer**|String|false|true|false|false|||||Y|AccountOfficer||
|**CountryCode**|String|false|true|false|false|||||Y|CountryCode||
|**SectorCode**|String|false|true|false|false|||||Y|SectorCode||
|**CpartyGroupName**|String|false|true|false|false|||||Y|CpartyGroupName||
|**Notes**|String|false|true|false|false|||||Y|Notes||
|**Owner**|String|false|true|false|false|||||Y|Owner||
|**Authorised**|Bool|false|true|false|false|||||Y|Authorised|True|
|**NameFirmName**|String|false|true|false|false|||||Y|NameFirmName||
|**NameCentreName**|String|false|true|false|false|||||Y|NameCentreName||
|**CountryCodeName**|String|false|true|false|false|||||Y|CountryCodeName||
|**SectorCodeName**|String|false|true|false|false|||||Y|SectorCodeName||
|**CompID**|String|true|true|false|false|||||Y|CompID||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/counterparty.go_tmp |
| code | **adaptor** | /adaptor/counterparty_impl.template |
| code | **api** | /application/counterparty_api.go |
| code | **dao** | /dao/counterparty.go_tmp |
| code | **datamodel** | /datamodel/counterparty.go_tmp |
| code | **menu** | /design/menu/counterparty.json |
| html | **list** | /html/Counterparty_List.html |
| html | **view** | /html/Counterparty_View.html |
| html | **edit** | /html/Counterparty_Edit.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:17:58**
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