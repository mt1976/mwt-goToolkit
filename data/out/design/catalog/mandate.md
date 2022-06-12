# **Mandate** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Mandate** (mandate) |
|Endpoint 	    |**/Mandate...** [^1]|
|Endpoint Query |**Mandate**|
|REST API|**/API/Mandate/**|
Glyph|**fas fa-city** ()
Friendly Name|**Mandate**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Mandate/MandateList) [Exportable]
* **View** (/Mandate/MandateView)
* **Edit** (/Mandate/MandateEdit)
* **Save** (/Mandate/MandateSave)
* **New** (/Mandate/MandateNew)








##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaMandatedUser**
SQL Table Key | **CompID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**MandatedUserKeyCounterpartyFirm**|String|true|true|false|false|||||Y|MandatedUserKeyCounterpartyFirm||
|**MandatedUserKeyCounterpartyCentre**|String|true|true|false|false|||||Y|MandatedUserKeyCounterpartyCentre||
|**MandatedUserKeyUserName**|String|true|true|false|false|||||Y|MandatedUserKeyUserName||
|**TelephoneNumber**|String|false|true|false|false|||||Y|TelephoneNumber||
|**EmailAddress**|String|false|true|false|false|||||Y|EmailAddress||
|**Active**|Bool|false|true|false|false|||||Y|Active|True|
|**FirstName**|String|false|true|false|false|||||Y|FirstName||
|**Surname**|String|false|true|false|false|||||Y|Surname||
|**DateOfBirth**|Time|true|true|false|false|||||Y|DateOfBirth||
|**Postcode**|String|false|true|false|false|||||Y|Postcode||
|**NationalIDNo**|String|false|true|false|false|||||Y|NationalIDNo||
|**PassportNo**|String|false|true|false|false|||||Y|PassportNo||
|**Country**|String|false|true|false|false|OL|Country|Country|Name|N|Country||
|**CountryName**|String|false|true|false|false|||||Y|CountryName||
|**FirmName**|String|false|true|false|false|||||Y|FirmName||
|**CentreName**|String|false|true|false|false|||||Y|CentreName||
|**Notify**|Bool|false|true|false|false|||||Y|Notify|True|
|**SystemUser**|String|false|true|false|false|||||Y|SystemUser||
|**CompID**|String|true|true|false|false|||||Y|CompID||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/mandate.go_tmp |
| code | **adaptor** | /adaptor/mandate_impl.template |
| code | **api** | /application/mandate_api.go |
| code | **dao** | /dao/mandate.go_tmp |
| code | **datamodel** | /datamodel/mandate.go_tmp |
| code | **menu** | /design/menu/mandate.json |
| html | **list** | /html/Mandate_List.html |
| html | **view** | /html/Mandate_View.html |
| html | **edit** | /html/Mandate_Edit.html |
| html | **new** | /html/Mandate_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:18:09**
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