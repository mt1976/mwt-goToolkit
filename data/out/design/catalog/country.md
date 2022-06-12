# **Country** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Country** (country) |
|Endpoint 	    |**/Country...** [^1]|
|Endpoint Query |**Code**|
|REST API|**/API/Country/**|
Glyph|**fas fa-globe-americas** ()
Friendly Name|**Country**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Country/CountryList) [Exportable]
* **View** (/Country/CountryView)
* **Edit** (/Country/CountryEdit)
* **Save** (/Country/CountrySave)
* **New** (/Country/CountryNew)








##  Provides
 * Lookup (Code Name)
 * Reverse Lookup (Name)





##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaCountry**
SQL Table Key | **Code**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**Code**|String|true|true|false|false|||||Y|Code||
|**Name**|String|false|true|false|false|||||Y|Name||
|**ShortCode**|String|false|true|false|false|||||Y|ShortCode||
|**EU_EEA**|Bool|false|true|false|false|||||Y|EU_EEA|True|
|**HolidaysWeekend**|String|false|true|false|false|||||Y|HolidaysWeekend||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/country.go_tmp |
| code | **adaptor** | /adaptor/country_impl.template |
| code | **api** | /application/country_api.go |
| code | **dao** | /dao/country.go_tmp |
| code | **datamodel** | /datamodel/country.go_tmp |
| code | **menu** | /design/menu/country.json |
| html | **list** | /html/Country_List.html |
| html | **view** | /html/Country_View.html |
| html | **edit** | /html/Country_Edit.html |
| html | **new** | /html/Country_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:18:01**
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