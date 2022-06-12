# **DealTypeFundamental** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**DealTypeFundamental** (dealtypefundamental) |
|Endpoint 	    |**/DealTypeFundamental...** [^1]|
|Endpoint Query |**DealTypeKey**|
Glyph|**fas fa-toolbox** ()
Friendly Name|**Deal Type Fundamental**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/DealTypeFundamental/DealTypeFundamentalList) [Exportable]
* **View** (/DealTypeFundamental/DealTypeFundamentalView)
* **Edit** (/DealTypeFundamental/DealTypeFundamentalEdit)
* **Save** (/DealTypeFundamental/DealTypeFundamentalSave)
* **New** (/DealTypeFundamental/DealTypeFundamentalNew)
* **Delete** (/DealTypeFundamental/DealTypeFundamentalDelete)







##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaDealTypeFundamentals**
SQL Table Key | **DealTypeKey**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**DealTypeKey**|String|true|true|false|false|||||Y|DealTypeKey||
|**Amendment**|Bool|false|true|false|false|||||Y|Amendment|True|
|**DefaultFrequency**|Int|false|true|false|false|||||Y|DefaultFrequency|0|
|**Directions**|String|false|true|false|false|||||Y|Directions||
|**SettledTermDealType**|String|false|true|false|false|||||Y|SettledTermDealType||
|**Optn**|Bool|false|true|false|false|||||Y|Optn|True|
|**AllowPledge**|Bool|false|true|false|false|||||Y|AllowPledge|True|
|**Takeup**|Bool|false|true|false|false|||||Y|Takeup|True|
|**MismatchDealType**|String|false|true|false|false|||||Y|MismatchDealType||
|**SettledHedgeTermDealType**|String|false|true|false|false|||||Y|SettledHedgeTermDealType||
|**SettlementCode**|String|false|true|false|false|||||Y|SettlementCode||
|**TermSubType**|String|false|true|false|false|||||Y|TermSubType||
|**FundingDealType**|String|false|true|false|false|||||Y|FundingDealType||
|**TransferType**|String|false|true|false|false|||||Y|TransferType||
|**TermDealType**|String|false|true|false|false|||||Y|TermDealType||
|**NegotiableInstrumentType**|String|false|true|false|false|||||Y|NegotiableInstrumentType||
|**Mismatch**|Bool|false|true|false|false|||||Y|Mismatch|True|
|**ComplexTransferSubType**|String|false|true|false|false|||||Y|ComplexTransferSubType||
|**LayOffDealType**|String|false|true|false|false|||||Y|LayOffDealType||
|**NIAccount**|Int|false|true|false|false|||||Y|NIAccount|0|
|**SimpleMMsubtype**|Int|false|true|false|false|||||Y|SimpleMMsubtype|0|
|**SwapDealType**|String|false|true|false|false|||||Y|SwapDealType||
|**Positions**|String|false|true|false|false|||||Y|Positions||
|**OptionOutright**|String|false|true|false|false|||||Y|OptionOutright||
|**SettledHedgeSpotDealType**|String|false|true|false|false|||||Y|SettledHedgeSpotDealType||
|**StraightThroughInterestMethod**|Bool|false|true|false|false|||||Y|StraightThroughInterestMethod|True|
|**SubType**|String|false|true|false|false|||||Y|SubType||
|**Rollover**|Bool|false|true|false|false|||||Y|Rollover|True|
|**DefaultIssuer**|String|false|true|false|false|||||Y|DefaultIssuer||
|**DefaultStartDate**|Int|false|true|false|false|||||Y|DefaultStartDate|0|
|**Fee**|String|false|true|false|false|||||Y|Fee||
|**NDF**|Bool|false|true|false|false|||||Y|NDF|True|
|**FXFX**|Bool|false|true|false|false|||||Y|FXFX|True|
|**ONIA**|Bool|false|true|false|false|||||Y|ONIA|True|
|**MarginSubType**|Int|false|true|false|false|||||Y|MarginSubType|0|
|**TransferDealType**|String|false|true|false|false|||||Y|TransferDealType||
|**IsFX**|Bool|false|true|false|false|||||Y|IsFX|True|
|**Ordr**|String|false|true|false|false|||||Y|Ordr||
|**OptionStyle**|String|false|true|false|false|||||Y|OptionStyle||
|**SpotDealType**|String|false|true|false|false|||||Y|SpotDealType||
|**CanIssue**|Bool|false|true|false|false|||||Y|CanIssue|True|
|**CanShort**|Bool|false|true|false|false|||||Y|CanShort|True|
|**FXMarginTradingType**|Int|false|true|false|false|||||Y|FXMarginTradingType|0|
|**Internal**|Bool|false|true|false|false|||||Y|Internal|True|
|**TicketBasename**|String|false|true|false|false|||||Y|TicketBasename||
|**InterestRateFutureType**|String|false|true|false|false|||||Y|InterestRateFutureType||
|**TradingLimitProductCode**|String|false|true|false|false|||||Y|TradingLimitProductCode||
|**Forward**|Bool|false|true|false|false|||||Y|Forward|True|
|**MaturityNotificationPeriod**|String|false|true|false|false|||||Y|MaturityNotificationPeriod||
|**NotificationEvents**|String|false|true|false|false|||||Y|NotificationEvents||
|**SwapSubType**|String|false|true|false|false|||||Y|SwapSubType||
|**ProductCode**|String|false|true|false|false|||||Y|ProductCode||
|**Funding**|Bool|false|true|false|false|||||Y|Funding|True|
|**AllocationPricing**|String|false|true|false|false|||||Y|AllocationPricing||
|**CancelPeriod**|String|false|true|false|false|||||Y|CancelPeriod||
|**MMMarginTradingType**|Int|false|true|false|false|||||Y|MMMarginTradingType|0|
|**OptionSpot**|String|false|true|false|false|||||Y|OptionSpot||
|**Transfer**|Bool|false|true|false|false|||||Y|Transfer|True|
|**NotificationPeriod**|String|false|true|false|false|||||Y|NotificationPeriod||
|**Paymentdateshift**|Int|false|true|false|false|||||Y|Paymentdateshift|0|
|**CloseOut**|Bool|false|true|false|false|||||Y|CloseOut|True|
|**FXOptionPricing**|String|false|true|false|false|||||Y|FXOptionPricing||
|**SettledHedgeOutrightDealType**|String|false|true|false|false|||||Y|SettledHedgeOutrightDealType||
|**RepoBond**|String|false|true|false|false|||||Y|RepoBond||
|**RepoTerm**|String|false|true|false|false|||||Y|RepoTerm||
|**RepoType**|Int|false|true|false|false|||||Y|RepoType|0|
|**DateRule**|String|false|true|false|false|||||Y|DateRule||
|**CorpTransferDealType**|String|false|true|false|false|||||Y|CorpTransferDealType||
|**GenerateFXImage**|Bool|false|true|false|false|||||Y|GenerateFXImage|True|
|**Variant**|String|false|true|false|false|||||Y|Variant||
|**HedgeTermDealType**|String|false|true|false|false|||||Y|HedgeTermDealType||
|**PricingModel**|String|false|true|false|false|||||Y|PricingModel||
|**HedgeOutrightDealType**|String|false|true|false|false|||||Y|HedgeOutrightDealType||
|**Fixing**|Bool|false|true|false|false|||||Y|Fixing|True|
|**Payment**|Bool|false|true|false|false|||||Y|Payment|True|
|**MT**|Bool|false|true|false|false|||||Y|MT|True|
|**SettlementInstructionStyle**|String|false|true|false|false|||||Y|SettlementInstructionStyle||
|**QuoteHistoryRequired**|Bool|false|true|false|false|||||Y|QuoteHistoryRequired|True|
|**Brokerage**|Bool|false|true|false|false|||||Y|Brokerage|True|
|**ExposureDisabled**|Bool|false|true|false|false|||||Y|ExposureDisabled|True|
|**CreditLine**|String|false|true|false|false|||||Y|CreditLine||
|**Encumbered**|Bool|false|true|false|false|||||Y|Encumbered|True|
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
| code | **application** | /application/dealTypeFundamental.go_tmp |
| code | **adaptor** | /adaptor/dealTypeFundamental_impl.template |
| code | **dao** | /dao/dealTypeFundamental.go_tmp |
| code | **datamodel** | /datamodel/dealTypeFundamental.go_tmp |
| code | **menu** | /design/menu/dealTypeFundamental.json |
| html | **list** | /html/DealTypeFundamental_List.html |
| html | **view** | /html/DealTypeFundamental_View.html |
| html | **edit** | /html/DealTypeFundamental_Edit.html |
| html | **new** | /html/DealTypeFundamental_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **15:18:07**
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