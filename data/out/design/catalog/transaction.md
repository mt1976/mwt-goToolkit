# **Transaction** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Transaction** (transaction) |
|Endpoint 	    |**/Transaction...** [^1]|
|Endpoint Query |**Ref**|
Glyph|**far fa-handshake** ()
Friendly Name|**Transaction**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Transaction/TransactionList) [Exportable]
* **View** (/Transaction/TransactionView)











##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaDealList**
SQL Table Key | **SienaReference**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**SienaReference**|String|true|true|false|false|||||Y|SienaReference||
|**Status**|String|false|true|false|false|||||Y|Status||
|**ValueDate**|Time|false|true|false|false|||||Y|ValueDate||
|**MaturityDate**|Time|false|true|false|false|||||Y|MaturityDate||
|**ContractNumber**|String|false|true|false|false|||||Y|ContractNumber||
|**ExternalReference**|String|false|true|false|false|||||Y|ExternalReference||
|**Book**|String|false|true|false|false|||||Y|Book||
|**MandatedUser**|String|false|true|false|false|||||Y|MandatedUser||
|**Portfolio**|String|false|true|false|false|||||Y|Portfolio||
|**AgreementId**|Int|false|true|false|false|||||Y|AgreementId|0|
|**BackOfficeRefNo**|String|false|true|false|false|||||Y|BackOfficeRefNo||
|**ISIN**|String|false|true|false|false|||||Y|ISIN||
|**UTI**|String|false|true|false|false|||||Y|UTI||
|**BookName**|String|false|true|false|false|||||Y|BookName||
|**Centre**|String|false|true|false|false|||||Y|Centre||
|**Firm**|String|false|true|false|false|||||Y|Firm||
|**DealTypeShortName**|String|false|true|false|false|||||Y|DealTypeShortName||
|**FullDealType**|String|false|true|false|false|||||Y|FullDealType||
|**TradeDate**|Time|false|true|false|false|||||Y|TradeDate||
|**DealtCcy**|String|false|true|false|false|||||Y|DealtCcy||
|**DealtAmount**|Float|false|true|false|false|||||Y|DealtAmount|0.00|
|**AgainstAmount**|Float|false|true|false|false|||||Y|AgainstAmount|0.00|
|**AgainstCcy**|String|false|true|false|false|||||Y|AgainstCcy||
|**AllInRate**|Float|false|true|false|false|||||Y|AllInRate|0.00|
|**MktRate**|Float|false|true|false|false|||||Y|MktRate|0.00|
|**SettleCcy**|String|false|true|false|false|||||Y|SettleCcy||
|**Direction**|String|false|true|false|false|||||Y|Direction||
|**NpvRate**|Float|false|true|false|false|||||Y|NpvRate|0.00|
|**OriginUser**|String|false|true|false|false|||||Y|OriginUser||
|**PayInstruction**|String|false|true|false|false|||||Y|PayInstruction||
|**ReceiptInstruction**|String|false|true|false|false|||||Y|ReceiptInstruction||
|**NIName**|String|false|true|false|false|||||Y|NIName||
|**CCYPair**|String|false|true|false|false|||||Y|CCYPair||
|**Instrument**|String|false|true|false|false|||||Y|Instrument||
|**PortfolioName**|String|false|true|false|false|||||Y|PortfolioName||
|**RVDate**|Time|false|true|false|false|||||Y|RVDate||
|**RVMTM**|Float|false|true|false|false|||||Y|RVMTM|0.00|
|**CounterBook**|String|false|true|false|false|||||Y|CounterBook||
|**CounterBookName**|String|false|true|false|false|||||Y|CounterBookName||
|**Party**|String|false|true|false|false|||||Y|Party||
|**PartyName**|String|false|true|false|false|||||Y|PartyName||
|**NameCentre**|String|false|true|false|false|||||Y|NameCentre||
|**NameFirm**|String|false|true|false|false|||||Y|NameFirm||
|**CustomerExternalView**|String|false|true|false|false|||||Y|CustomerExternalView||
|**CustomerSienaView**|String|false|true|false|false|||||Y|CustomerSienaView||
|**CompID**|String|false|true|false|false|||||Y|CompID||
|**SienaDealer**|String|false|true|false|false|||||Y|SienaDealer||
|**DealOwner**|String|false|true|false|false|||||Y|DealOwner||
|**DealOwnerMnemonic**|String|false|true|false|false|||||Y|DealOwnerMnemonic||
|**EditedByUser**|String|false|true|false|false|||||Y|EditedByUser||
|**UTCOriginTime**|String|false|true|false|false|||||Y|UTCOriginTime||
|**UTCUpdateTime**|String|false|true|false|false|||||Y|UTCUpdateTime||
|**MarginTrading**|Bool|false|true|false|false|||||Y|MarginTrading|True|
|**SwapPoints**|Float|false|true|false|false|||||Y|SwapPoints|0.00|
|**SpotDate**|Time|false|true|false|false|||||Y|SpotDate||
|**SpotRate**|Float|false|true|false|false|||||Y|SpotRate|0.00|
|**MktSpotRate**|Float|false|true|false|false|||||Y|MktSpotRate|0.00|
|**SpotSalesMargin**|Float|false|true|false|false|||||Y|SpotSalesMargin|0.00|
|**SpotChlMargin**|Float|false|true|false|false|||||Y|SpotChlMargin|0.00|
|**RerouteCcy**|String|false|true|false|false|||||Y|RerouteCcy||
|**CustomerPayInstruction**|String|false|true|false|false|||||Y|CustomerPayInstruction||
|**CustomerReceiptInstruction**|String|false|true|false|false|||||Y|CustomerReceiptInstruction||
|**BackOfficeNotes**|String|false|true|false|false|||||Y|BackOfficeNotes||
|**CustomerStatementNotes**|String|false|true|false|false|||||Y|customerStatementNotes||
|**NotesMargin**|Float|false|true|false|false|||||Y|NotesMargin|0.00|
|**RequestedBy**|String|false|true|false|false|||||Y|RequestedBy||
|**EditReason**|String|false|true|false|false|||||Y|EditReason||
|**EditOtherReason**|String|false|true|false|false|||||Y|EditOtherReason||
|**NICleanPrice**|Float|false|true|false|false|||||Y|NICleanPrice|0.00|
|**NIDirtyPrice**|Float|false|true|false|false|||||Y|NIDirtyPrice|0.00|
|**NIYield**|Float|false|true|false|false|||||Y|NIYield|0.00|
|**NIClearingSystem**|String|false|true|false|false|||||Y|NIClearingSystem||
|**Acceptor**|String|false|true|false|false|||||Y|Acceptor||
|**NIDiscount**|Float|false|true|false|false|||||Y|NIDiscount|0.00|
|**FastPay**|Bool|false|true|false|false|||||Y|FastPay|True|
|**PaymentFee**|Float|false|true|false|false|||||Y|PaymentFee|0.00|
|**PaymentFeePolicy**|String|false|true|false|false|||||Y|PaymentFeePolicy||
|**PaymentReason**|String|false|true|false|false|||||Y|PaymentReason||
|**PaymentDate**|Time|false|true|false|false|||||Y|PaymentDate||
|**SettlementDate**|Time|false|true|false|false|||||Y|SettlementDate||
|**FixingDate**|Time|false|true|false|false|||||Y|FixingDate||
|**VenueUTI**|String|false|true|false|false|||||Y|VenueUTI||
|**EditVersion**|Int|false|true|false|false|||||Y|EditVersion|0|
|**BrokeragePercentage**|Float|false|true|false|false|||||Y|BrokeragePercentage|0.00|
|**BrokerageAmount**|Float|false|true|false|false|||||Y|BrokerageAmount|0.00|
|**BrokerageCurrency**|String|false|true|false|false|||||Y|BrokerageCurrency||
|**BrokerageDate**|Time|false|true|false|false|||||Y|BrokerageDate||
|**AccountName**|String|false|true|false|false|||||Y|AccountName||
|**AccountNumber**|String|false|true|false|false|||||Y|AccountNumber||
|**CashBalance**|Float|false|true|false|false|||||Y|CashBalance|0.00|
|**DebitFrequency**|String|false|true|false|false|||||Y|DebitFrequency||
|**CreditFrequency**|String|false|true|false|false|||||Y|CreditFrequency||
|**ManuallyQuoted**|String|false|true|false|false|||||Y|ManuallyQuoted||
|**LedgerBalance**|Float|false|true|false|false|||||Y|LedgerBalance|0.00|
|**SettAmtOutstanding**|Float|false|true|false|false|||||Y|SettAmtOutstanding|0.00|
|**FeePercentage**|Float|false|true|false|false|||||Y|FeePercentage|0.00|
|**FeeAmount**|Float|false|true|false|false|||||Y|FeeAmount|0.00|
|**Venue**|String|false|true|false|false|||||Y|Venue||
|**EURAmount**|Float|false|true|false|false|||||Y|EURAmount|0.00|
|**EUROtherAmount**|Float|false|true|false|false|||||Y|EUROtherAmount|0.00|
|**LEI**|String|false|true|false|false|||||Y|LEI||
|**Equity**|String|false|true|false|false|||||Y|Equity||
|**Shares**|Int|false|true|false|false|||||Y|Shares|0|
|**QuoteExpiryDate**|Time|false|true|false|false|||||Y|QuoteExpiryDate||
|**Commodity**|String|false|true|false|false|||||Y|Commodity||
|**PaymentSystemSienaView**|String|false|true|false|false|||||Y|PaymentSystemSienaView||
|**PaymentSystemExternalView**|String|false|true|false|false|||||Y|PaymentSystemExternalView||
|**SalesProfit**|Float|false|true|false|false|||||Y|SalesProfit|0.00|
|**RejectReason**|String|false|true|false|false|||||Y|RejectReason||
|**PaymentError**|String|false|true|false|false|||||Y|PaymentError||
|**RepoPrincipal**|Float|false|true|false|false|||||Y|RepoPrincipal|0.00|
|**FixingFrequency**|String|false|true|false|false|||||Y|FixingFrequency||
|**Dealt**|String|false|false|true|false|||||N|||
|**Against**|String|false|false|true|false|||||N|||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/transaction.go_tmp |
| code | **adaptor** | /adaptor/transaction_impl.template |
| code | **dao** | /dao/transaction.go_tmp |
| code | **datamodel** | /datamodel/transaction.go_tmp |
| code | **menu** | /design/menu/transaction.json |
| html | **list** | /html/Transaction_List.html |
| html | **view** | /html/Transaction_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:17:52**
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