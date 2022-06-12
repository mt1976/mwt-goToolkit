# **Payee** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Payee** (payee) |
|Endpoint 	    |**/Payee...** [^1]|
|Endpoint Query |**ID**|
Glyph|**fas fa-wallet** ("")
Friendly Name|**Payee**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Payee/PayeeList) [Exportable]
* **View** (/Payee/PayeeView)











##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaCounterpartyPayee**
SQL Table Key | **KeyCounterpartyFirm**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**SourceTable**|String|true|true|false|false|||||Y|SourceTable||
|**KeyCounterpartyFirm**|String|true|true|false|false|||||Y|KeyCounterpartyFirm||
|**KeyCounterpartyCentre**|String|true|true|false|false|||||Y|KeyCounterpartyCentre||
|**KeyCurrency**|String|false|true|false|false|||||Y|KeyCurrency||
|**KeyName**|String|true|true|false|false|||||Y|KeyName||
|**KeyNumber**|String|true|true|false|false|||||Y|KeyNumber||
|**KeyDirection**|String|true|true|false|false|||||Y|KeyDirection||
|**KeyType**|String|true|true|false|false|||||Y|KeyType||
|**FullName**|String|false|true|false|false|||||Y|FullName||
|**Address**|String|false|true|false|false|||||Y|Address||
|**PhoneNo**|String|false|true|false|false|||||Y|PhoneNo||
|**Country**|String|false|true|false|false|OL|Country|Country|Name|N|Country||
|**Bic**|String|false|true|false|false|||||Y|Bic||
|**Iban**|String|false|true|false|false|||||Y|Iban||
|**AccountNo**|String|false|true|false|false|||||Y|AccountNo||
|**FedWireNo**|String|false|true|false|false|||||Y|FedWireNo||
|**SortCode**|String|false|true|false|false|||||Y|SortCode||
|**BankName**|String|false|true|false|false|||||Y|BankName||
|**BankPinCode**|String|false|true|false|false|||||Y|BankPinCode||
|**BankAddress**|String|false|true|false|false|||||Y|BankAddress||
|**Reason**|String|false|true|false|true|||||N|Reason||
|**BankSettlementAcct**|Bool|false|true|false|false|||||Y|BankSettlementAcct|True|
|**UpdatedUserId**|String|false|true|false|false|||||Y|UpdatedUserId||
|**Status**|String|false|false|true|false|||||N|||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/payee.go_tmp |
| code | **adaptor** | /adaptor/payee_impl.template |
| code | **dao** | /dao/payee.go_tmp |
| code | **datamodel** | /datamodel/payee.go_tmp |
| code | **menu** | /design/menu/payee.json |
| html | **list** | /html/Payee_List.html |
| html | **view** | /html/Payee_View.html |


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