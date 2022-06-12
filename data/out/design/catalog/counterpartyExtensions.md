# **CounterpartyExtensions** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**CounterpartyExtensions** (counterpartyextensions) |
|Endpoint 	    |**/CounterpartyExtensions...** [^1]|
|Endpoint Query |**ID**|
Glyph|**fas fa-passport** ()
Friendly Name|**Counterparty Extension**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/CounterpartyExtensions/CounterpartyExtensionsList) [Exportable]
* **View** (/CounterpartyExtensions/CounterpartyExtensionsView)
* **Edit** (/CounterpartyExtensions/CounterpartyExtensionsEdit)
* **Save** (/CounterpartyExtensions/CounterpartyExtensionsSave)
* **New** (/CounterpartyExtensions/CounterpartyExtensionsNew)
* **Delete** (/CounterpartyExtensions/CounterpartyExtensionsDelete)







##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaCounterpartyExtensions**
SQL Table Key | **CompID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**NameFirm**|String|true|true|false|false|||||Y|NameFirm||
|**NameCentre**|String|true|true|false|false|||||Y|NameCentre||
|**BICCode**|String|false|true|false|false|||||Y|BICCode||
|**ContactIndicator**|Bool|false|true|false|false|||||Y|ContactIndicator|True|
|**CoverTrade**|Bool|false|true|false|false|||||Y|CoverTrade|True|
|**CustomerCategory**|String|false|true|false|false|||||Y|CustomerCategory||
|**FSCSInclusive**|Bool|false|true|false|false|||||Y|FSCSInclusive|True|
|**FeeFactor**|Float|false|true|false|false|||||Y|FeeFactor|0.00|
|**InactiveStatus**|Bool|false|true|false|false|||||Y|InactiveStatus|True|
|**Indemnity**|Bool|false|true|false|false|||||Y|Indemnity|True|
|**KnowYourCustomerStatus**|Bool|false|true|false|false|||||Y|KnowYourCustomerStatus|True|
|**LERLimitCarveOut**|Bool|false|true|false|false|||||Y|LERLimitCarveOut|True|
|**LastAmended**|Time|false|true|false|false|||||Y|LastAmended||
|**LastLogin**|Time|false|true|false|false|||||Y|LastLogin||
|**LossGivenDefault**|Float|false|true|false|false|||||Y|LossGivenDefault|0.00|
|**MIC**|String|false|true|false|false|||||Y|MIC||
|**ProtectedDepositor**|Bool|false|true|false|false|||||Y|ProtectedDepositor|True|
|**RPTCurrency**|String|false|true|false|false|||||Y|RPTCurrency||
|**RateTimeout**|Int|false|true|false|false|||||Y|RateTimeout|0|
|**RateValidation**|String|false|true|false|false|||||Y|RateValidation||
|**Registered**|Bool|false|true|false|false|||||Y|Registered|True|
|**RegulatoryCategory**|String|false|true|false|false|||||Y|RegulatoryCategory||
|**SecuredSettlement**|Bool|false|true|false|false|||||Y|SecuredSettlement|True|
|**SettlementLimitCarveOut**|Bool|false|true|false|false|||||Y|SettlementLimitCarveOut|True|
|**SortCode**|String|false|true|false|false|||||Y|SortCode||
|**Training**|Bool|false|true|false|false|||||Y|Training|True|
|**TrainingCode**|String|false|true|false|false|||||Y|TrainingCode||
|**TrainingReceived**|Bool|false|true|false|false|||||Y|TrainingReceived|True|
|**Unencumbered**|Bool|false|true|false|false|||||Y|Unencumbered|True|
|**LEIExpiryDate**|Time|false|true|false|false|||||Y|LEIExpiryDate||
|**MIFIDReviewDate**|Time|false|true|false|false|||||Y|MIFIDReviewDate||
|**GDPRReviewDate**|Time|false|true|false|false|||||Y|GDPRReviewDate||
|**DelegatedReporting**|String|false|true|false|false|||||Y|DelegatedReporting||
|**BOReconcile**|Bool|false|true|false|false|||||Y|BOReconcile|True|
|**MIFIDReportableDealsAllowed**|Bool|false|true|false|false|||||Y|MIFIDReportableDealsAllowed|True|
|**SignedInvestmentAgreement**|Bool|false|true|false|false|||||Y|SignedInvestmentAgreement|True|
|**AppropriatenessAssessment**|Bool|false|true|false|false|||||Y|AppropriatenessAssessment|True|
|**FinancialCounterparty**|Bool|false|true|false|false|||||Y|FinancialCounterparty|True|
|**Collateralisation**|String|false|true|false|false|||||Y|Collateralisation||
|**PortfolioCode**|String|false|true|false|false|||||Y|PortfolioCode||
|**ReconciliationLetterFrequency**|String|false|true|false|false|||||Y|ReconciliationLetterFrequency||
|**DirectDealing**|Bool|false|true|false|false|||||Y|DirectDealing|True|
|**CompID**|String|true|true|false|false|||||Y|CompID||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/counterpartyExtensions.go_tmp |
| code | **adaptor** | /adaptor/counterpartyExtensions_impl.template |
| code | **dao** | /dao/counterpartyExtensions.go_tmp |
| code | **datamodel** | /datamodel/counterpartyExtensions.go_tmp |
| code | **menu** | /design/menu/counterpartyExtensions.json |
| html | **list** | /html/CounterpartyExtensions_List.html |
| html | **view** | /html/CounterpartyExtensions_View.html |
| html | **edit** | /html/CounterpartyExtensions_Edit.html |
| html | **new** | /html/CounterpartyExtensions_New.html |


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