# **Account** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Account** (account) |
|Endpoint 	    |**/Account...** [^1]|
|Endpoint Query |**AccountNo**|
|REST API|**/API/Account/**|
Glyph|**fas fa-landmark** ()
Friendly Name|**Account**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Account/AccountList) [Exportable]
* **View** (/Account/AccountView)











##  Provides
 * Lookup (SienaReference AccountName)
 * Reverse Lookup (CashBalance)





##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaAccount**
SQL Table Key | **SienaReference**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**SienaReference**|String|true|true|false|false|||||Y|SienaReference||
|**CustomerSienaView**|String|false|true|false|false|||||Y|CustomerSienaView||
|**SienaCommonRef**|String|true|true|false|false|||||Y|SienaCommonRef||
|**Status**|String|false|true|false|false|||||Y|Status||
|**StartDate**|Time|false|true|false|false|||||Y|StartDate||
|**MaturityDate**|Time|false|true|false|false|||||Y|MaturityDate||
|**ContractNumber**|String|false|true|false|false|||||Y|ContractNumber||
|**ExternalReference**|String|false|true|false|false|||||Y|ExternalReference||
|**CCY**|String|false|true|false|false|OL|Currency|CCY|Name|Y|CCY||
|**Book**|String|false|true|false|false|OL|Book|Book|FullName|N|Book||
|**MandatedUser**|String|false|true|false|false|||||Y|MandatedUser||
|**BackOfficeNotes**|String|false|true|false|false|||||Y|BackOfficeNotes||
|**CashBalance**|Float|false|true|false|false|||||Y|CashBalance|0.00|
|**AccountNumber**|String|false|true|false|false|||||Y|AccountNumber||
|**AccountName**|String|false|true|false|false|||||Y|AccountName||
|**LedgerBalance**|Float|false|true|false|false|||||Y|LedgerBalance|0.00|
|**Portfolio**|String|false|true|false|false|OL|Portfolio|Portfolio|Description1|N|Portfolio||
|**AgreementId**|Int|false|true|false|false|||||Y|AgreementId|0|
|**BackOfficeRefNo**|String|false|true|false|false|||||Y|BackOfficeRefNo||
|**ISIN**|String|false|true|false|false|||||Y|ISIN||
|**UTI**|String|false|true|false|false|||||Y|UTI||
|**CCYName**|String|false|true|false|false|||||Y|CCYName||
|**BookName**|String|false|true|false|false|||||Y|BookName||
|**PortfolioName**|String|false|true|false|false|||||Y|PortfolioName||
|**Centre**|String|false|true|false|false|OL|Centre|Centre|Name|N|Centre||
|**DealTypeKey**|String|false|true|false|false|||||Y|DealTypeKey||
|**DealTypeShortName**|String|false|true|false|false|||||Y|DealTypeShortName||
|**InternalId**|Int|false|true|false|false|||||Y|InternalId|0|
|**InternalDeleted**|Time|false|true|false|false|||||Y|InternalDeleted||
|**UpdatedTransactionId**|String|false|true|false|false|||||Y|UpdatedTransactionId||
|**UpdatedUserId**|String|false|true|false|false|||||Y|UpdatedUserId||
|**UpdatedDateTime**|Time|false|true|false|false|||||Y|UpdatedDateTime||
|**DeletedTransactionId**|String|false|true|false|false|||||Y|DeletedTransactionId||
|**DeletedUserId**|String|false|true|false|false|||||Y|DeletedUserId||
|**ChangeType**|String|false|true|false|false|||||Y|ChangeType||
|**CCYDp**|Int|false|true|false|false|||||Y|CCYDp|0|
|**CompID**|String|false|true|false|false|||||Y|CompID||
|**Firm**|String|false|true|false|false|OL|Firm|Firm|FullName|N|Firm||
|**DealType**|String|false|true|false|false|||||Y|DealType||
|**FullDealType**|String|false|true|false|false|||||Y|FullDealType||
|**DealingInterface**|String|false|true|false|false|||||Y|DealingInterface||
|**DealtAmount**|Float|false|true|false|false|||||Y|DealtAmount|0.00|
|**ParentContractNumber**|String|false|true|false|false|||||Y|ParentContractNumber||
|**InterestFrequency**|String|false|true|false|false|||||Y|InterestFrequency||
|**InterestAction**|String|false|true|false|false|||||Y|InterestAction||
|**InterestStrategy**|String|false|true|false|false|||||Y|InterestStrategy||
|**InterestBasis**|String|false|true|false|false|||||Y|InterestBasis||
|**SienaDealer**|String|false|true|false|false|||||Y|SienaDealer||
|**DealOwner**|String|false|true|false|false|||||Y|DealOwner||
|**OriginUser**|String|false|true|false|false|||||Y|OriginUser||
|**EditedByUser**|String|false|true|false|false|||||Y|EditedByUser||
|**DealOwnerMnemonic**|String|false|true|false|false|||||Y|DealOwnerMnemonic||
|**UTCOriginTime**|String|false|true|false|false|||||Y|UTCOriginTime||
|**UTCUpdateTime**|String|false|true|false|false|||||Y|UTCUpdateTime||
|**CustomerStatementNotes**|String|false|true|false|false|||||Y|customerStatementNotes||
|**NotesMargin**|Float|false|true|false|false|||||Y|NotesMargin|0.00|
|**RequestedBy**|String|false|true|false|false|||||Y|RequestedBy||
|**EditReason**|String|false|true|false|false|||||Y|EditReason||
|**EditOtherReason**|String|false|true|false|false|||||Y|EditOtherReason||
|**NoticeDays**|Int|false|true|false|false|||||Y|NoticeDays|0|
|**DebitFrequency**|String|false|true|false|false|||||Y|DebitFrequency||
|**CreditFrequency**|String|false|true|false|false|||||Y|CreditFrequency||
|**EURAmount**|Float|false|true|false|false|||||Y|EURAmount|0.00|
|**EUROtherAmount**|Float|false|true|false|false|||||Y|EUROtherAmount|0.00|
|**PaymentSystemSienaView**|String|false|true|false|false|||||Y|PaymentSystemSienaView||
|**PaymentSystemExternalView**|String|false|true|false|false|||||Y|PaymentSystemExternalView||
|**DealtCA**|String|false|false|true|false|||||N|||
|**AgainstCA**|String|false|false|true|false|||||N|||
|**LedgerCA**|String|false|false|true|false|||||N|||
|**CashBalanceCA**|String|false|false|true|false|||||N|||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/account.go_tmp |
| code | **adaptor** | /adaptor/account_impl.template |
| code | **api** | /application/account_api.go |
| code | **dao** | /dao/account.go_tmp |
| code | **datamodel** | /datamodel/account.go_tmp |
| code | **menu** | /design/menu/account.json |
| html | **list** | /html/Account_List.html |
| html | **view** | /html/Account_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:17:53**
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