// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package m3ter

import (
	"context"
	"net/http"
	"os"
	"slices"

	"github.com/m3ter-com/m3ter-sdk-go/internal/requestconfig"
	"github.com/m3ter-com/m3ter-sdk-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the m3ter API. You should not instantiate this client directly,
// and instead use the [NewClient] method instead.
type Client struct {
	Options []option.RequestOption
	// Endpoint for retrieving a JSON Web Token (JWT) bearer token for a ServiceUser
	// using the Client Credentials Grant flow.
	//
	// A ServiceUser represents the automated process you want to grant access to your
	// Organization - that is, as an API user.
	Authentication *AuthenticationService
	// Endpoints for Account related operations such as creation, update, list and
	// delete. An Account represents one of your end-customer accounts.
	//
	// Accounts do not belong to a Product to allow for cases where an end customer
	// takes more than one of your Products, and the charges for these Products differ.
	//
	// You typically attach a priced Plan or Plan Template to an Account before you can
	// generate bills for the Account:
	//
	//   - If a customer consumes several of your Products, you can attach a priced Plan
	//     or Plan Template to the Account for charging against each Product.
	//   - If an Account is charged solely on the basis of an agreed
	//     Prepayment/Commitment amount but not all of the Prepayment is prepaid, you can
	//     use a customized billing schedule for outstanding fees without having to
	//     attach a Plan to the Account to generate Bills.
	//
	// You can create Child Accounts for end customers who hold multiple Accounts with
	// you. You can then set up billing for the Parent/Child Account usage to have the
	// end-customer billed once for the Parent Account, instead of having separate
	// bills issued for usage against each of their multiple Accounts.
	//
	// **IMPORTANT! - use of PII:** The use of any of your end-customers' Personally
	// Identifiable Information (PII) in m3ter is restricted to a few fields on the
	// **Account** entity. Please ensure that only the `name`, `address`, or
	// `emailAddress` fields contain any end-customer PII data on any Accounts you
	// create. See the
	// [Introduction section](https://www.m3ter.com/docs/api#section/Introduction)
	// above for more details.
	Accounts *AccountService
	// Endpoints for AccountPlan and AccountPlanGroup related operations such as
	// creation, update, list and delete.
	//
	// **AccountPlans** An Account represents one of your end-customer accounts. To
	// create an AccountPlan, you attach a Product Plan to an Account. The AccountPlan
	// then determines the charges incurred at billing by your end customer for
	// consuming the Product the Plan is for:
	//
	//   - **AccountPlan Active/Inactive**. Set start and end dates to define the period
	//     the AccountPlan is active for the Account.
	//   - **AccountPlan per Product**. If an end customer consumes multiple Products,
	//     create separate AccountPlans to charge for each Product.
	//
	// **AccountPlan Constraints:**
	//
	//   - Only one AccountPlan per Product can be active at any one time for an Account.
	//   - If you create a Plan as a custom Plan for a specific Account, you can only use
	//     it to create an AccountPlan for that Account.
	//
	// **AccountPlanGroups** Plan Groups are used when you want to apply a minimum
	// spend amount at billing across several of your Products each of which are priced
	// separately - when you create the Plan Group, you define an overall minimum spend
	// and then add any priced Plans you want to include in the Group. To create an
	// AccounPlanGroup, you can attach a Plan Group to an Account that consumes the
	// separate Products which are priced using the included Plans. At billing, the
	// minimum spend you've defined for the Plan Group is applied:
	//
	//   - **Active AccountPlanGroup**. Set the start and end dates to define the period
	//     for which the Plan Group will be active for the Account.
	//
	// **Plan Group Notes:**
	//
	//   - You can only add _one Plan for the same Product_ to a Plan Group. See the
	//     [Plan Group](https://www.m3ter.com/docs/api#tag/PlanGroup) in this API
	//     Reference for more details on creating Plan Groups.
	//   - You can create a _custom Plan Group_ for an Account, which means the Plan
	//     Group can only be attached to that Account to create an AccountPlanGroup.
	//
	// **AcountPlanGroup - Notes and Constraints:**
	//
	//   - **AccountPlanGroup is type of AccountPlan** When you attach a Plan Group to an
	//     Account, this creates an AccountPlanGroup. However, the m3ter data model _does
	//     not support a separate AccountPlanGroup entity_, and an AccountPlanGroup is a
	//     type of AccountPlan where a `planGroupId` is used instead of a `planId` when
	//     it's created. See the
	//     [Create AccountPlan](https://www.m3ter.com/docs/api#tag/AccountPlan/operation/PostAccountPlan)
	//     call in this section and
	//     [Attaching Plan Groups to an Account](https://www.m3ter.com/docs/guides/end-customer-accounts/attaching-plan-groups-to-an-account)
	//     in our main User Documentation.
	//   - **Multiple AccountPlan Groups:** You can attach more than one Plan Group to an
	//     Account to create multiple AccountPlanGroups, but the rule that _only one
	//     attached Plan per Product can be active at any one time for an Account_ is
	//     preserved:
	//   - Multiple attached Plan Groups on an Account can have overlapping dates only
	//     if none of the Plan Groups contain a Plan belonging to the same Product. If
	//     you try to attach a Plan Group to an Account with Plan Groups already
	//     attached and:
	//   - The new Plan Group contains a Product Plan that also belongs to a Plan
	//     Group already attached to the Account.
	//   - The dates for these "matched Plan" Plan Groups being active for the
	//     Account would overlap.
	//   - Then you'll receive an error and the attachment will be blocked.
	AccountPlans *AccountPlanService
	// Endpoints for listing, creating, updating, retrieving, or deleting Aggregations.
	//
	// An Aggregation links to a Meter and targets a Data Field or Derived Field on the
	// Meter. You define the method of aggregation used to convert the usage data
	// collected by the targeted Meter field into a numerical unit of measurement.
	//
	// You can then use the unit of measurement an Aggregation yields as a metric for
	// pricing Product Plans and apply usage-based pricing to your products and
	// services. You might also want to aggregate raw data measures for other purposes,
	// such as to feed into analytical or business performance tools.
	//
	// **Notes:**
	//
	//   - **Contrast with Compound Aggregations**. Standard or simple Aggregations of
	//     this type, which apply an aggregation method directly to Meter usage data
	//     fields, are contrasted with
	//     [Compound Aggregations](https://www.m3ter.com/docs/api#tag/CompoundAggregation).
	//     A Compound Aggregation typically references one or more simple Aggregations
	//     and applies a calculation to them to derive pricing metrics needed to serve
	//     more complex usage-based pricing scenarios.
	//   - **Segmented Aggregations**. Segmented Aggregations allow you to segment the
	//     usage data collected by a single Meter. This capability is very useful for
	//     implementing some pricing and billing use cases. See
	//     [Segmented Aggregations](https://www.m3ter.com/docs/guides/usage-data-aggregations/segmented-aggregations)
	//     in our main documentation for more details.
	Aggregations *AggregationService
	// Endpoints for creating/retrieving/updating/deleting Balances on Accounts.
	//
	// When you have created a Balance for an Account, you can create a positive or
	// negative Transaction amounts for the Balance. To do this, you must first define
	// Transaction Types for your Organization, and then use one of these Transaction
	// Types when you add a specific Transaction to a Balance - see the
	// [Create TransactionType](https://www.m3ter.com/docs/api#tag/TransactionType/operation/CreateTransactionType)
	// call in the Transaction Type section in this API Reference for more details.
	//
	// Balances are typically used when a customer prepays an amount to add a credit to
	// their Account, which can then be draw-down against charges due for product or
	// service consumption. You can include options to top-up the original Balance.
	//
	// Examples of how Balances for end customer Accounts can be used:
	//
	//   - Onboarding Balance/Free Trials. Offering an onboarding incentive to new
	//     customers as an initial free credit Balance on their Account.
	//
	//   - Balance as initial commitment. Add a Balance amount to a new customer Account.
	//     This acts as an initial commitment, which allows them to use the service and
	//     gain an accurate insight into their usage level.
	//
	//   - Managing Customer Satisfaction. Use Balance as credits that will be applied to
	//     subsequent Bills as compensation for acknowledged service delivery issues.
	//
	// - Facilitating Balance Adjustments:
	//   - Apply negative amounts to immediately write-off outstanding Balances.
	//
	// #### What is the difference between Balances and Commitments/Prepayments?
	//
	// To manage credit amounts for your end-customer Accounts, you can use Balances or
	// Commitments/Prepayments. However, these two kinds of credits for Accounts serve
	// different purposes.
	//
	// Commitments - also referred to as Prepayments - are used for amounts
	// end-customers have agreed to pay for consuming your product or services across a
	// full contract term. A customer might pay the entire or only part of the agreed
	// amount upfront, but **_the commitment or prepayment amount is payable regardless
	// of the actual usage by the customer of your service or product._**
	//
	// In contrast, a Balance - often referred to as a Top-Up or Prepaid draw-down - is
	// used when a customer wants to add a credit amount to their Account at any time
	// during the service period or when you as service provider want to add a credit
	// to a customer Account. This Balance credit can then be drawn-down against for
	// billing the Account for usage, minimum spend, standing charges, or recurring
	// charges due. Balances therefore serve payment use cases in a more flexible way,
	// for example to be used for a "Free Credit" sign-up scheme you offer to encourage
	// sales or to enhance customer satisfaction by adding credit to an Account to
	// compensate for service delivery issues.
	//
	// You can use Commitments/Prepayments and Balances together on Account, and define
	// at Organization or individual Account level the order in which any
	// Balance/Commitment credit on an Account is drawn-down - Balance amounts first or
	// Commitment/Prepayment amounts first.
	Balances *BalanceService
	// Endpoints for billing operations such as creating, updating,
	// listing,downloading, and deleting Bills.
	//
	// Bills are generated for an Account, and are calculated in accordance with the
	// usage-based pricing Plans applied for the Products the Account consumes. These
	// endpoints enable interaction with the billing system, allowing you to obtain
	// billing details and insights into the consumption patterns and charges of your
	// end-customer Accounts.
	Bills *BillService
	// Endpoints for updating and retreiving the Bill Configuration for an
	// Organization. The Organization represents your company as a direct customer of
	// the m3ter service.
	//
	// You can use the **Update BillConfig** endpoint to set a global lock date for
	// **all** Bills - any Bill with a service period end date on or before the set
	// date will be locked and cannot be updated.
	//
	// **Warning: Ensure all Bills are Approved!** If you try to set a global lock date
	// when there remains Bills in a _Pending_ state whose service period end date is
	// on or before the specified lock date, then you'll receive an error.
	BillConfig *BillConfigService
	// Endpoints that manage Commitments _(also known as Prepayments)_ in the context
	// of usage-based pricing and billing. A Commitment represents an agreement where
	// the end-customer has agreed to pay a fixed minimum amount throughout the
	// contract period. **_The commitment amount is payable regardless of the actual
	// usage by the customer of your service or product._**
	//
	// These endpoints enable the creation, updating, retrieval, and deletion of
	// Commitments. Use them to manage your customer's Commitments and ensure optimal
	// revenue recognition:
	//
	//   - Specify which type of charges can draw-down against a Commitment amount on an
	//     Account at billing: usage, minimum spend, standing charges, or recurring
	//     charges.
	//   - Define overage surcharge percentages, which are applied when the usage charges
	//     exceed the agreed Commitment amount within the contract duration.
	//
	// #### What is the difference between Balances and Commitments/Prepayments?
	//
	// To manage credit amounts for your end-customer Accounts, you can use Balances or
	// Commitments/Prepayments. However, these two kinds of credits for Accounts serve
	// different purposes.
	//
	// Commitments/Prepayments are used for amounts end-customers have agreed to pay
	// for consuming your product or services across a full contract term. A customer
	// might pay the entire or only part of the agreed amount upfront, but **_the
	// prepayment amount is payable regardless of the actual usage by the customer of
	// your service or product._**
	//
	// In contrast, a Balance - often referred to as a Top-Up or Prepaid draw-down - is
	// used when a customer wants to add a credit amount to their Account at any time
	// during the service period or when you as service provider want to add a credit
	// to a customer Account. This Balance credit can then be drawn-down against for
	// billing the Account for usage, minimum spend, standing charges, or recurring
	// charges due. Balances therefore serve payment use cases in a more flexible way,
	// for example to be used for a "Free Credit" sign-up scheme you offer to encourage
	// sales or to enhance customer satisfaction by adding credit to an Account to
	// compensate for service delivery issues.
	//
	// You can use Prepayments/Commitments and Balances together on Account, and define
	// at Organization or individual Account level the order in which any
	// Balance/Prepayment credit on an Account is drawn-down - Balance amounts first or
	// Prepayment amounts first.
	//
	// #### Billing for Commitments
	//
	// If not all of an agreed Commitment amount is paid at the start of an
	// end-customer contract period, you can choose one of two options for billing the
	// outstanding fees due on the customer Account:
	//
	// - Select a Product _Plan to bill with_.
	// - Define a _schedule of billing dates_.
	Commitments *CommitmentService
	// Endpoints for creating, retrieving, listing, and cancelling Bill Jobs.
	//
	// Bill Jobs are critical components in billing management, providing asynchronous
	// mechanisms to calculate and handle bills.
	//
	// Bill Jobs give you the flexibiity to run Bills manually for Accounts to suit
	// different billing management purposes. For example, some historical usage data
	// has come in for an Account and you want to run a Bill for a specific date on
	// that Account to check that the Bill is showing correctly for the charges due on
	// the new usage data.
	BillJobs *BillJobService
	// Endpoints for creating/updating/deleting Charges.
	//
	// Create Charges for your end-customer Accounts to create ad-hoc line items for
	// Account billing. Charges are:
	//
	//   - Created for either debit or credit amounts.
	//   - Linked to a Product for accounting purposes.
	//   - Optionally linked to a Contract.
	//   - Given a specific date for billing. When a bill job has run for the specified
	//     Charge bill date, a Charge appears as an Ad-hoc line item on the Bill.
	//   - Assigned a service period.
	//   - Available in any currency defined for your Organization. See
	//     [Creating Charges for Accounts](https://www.m3ter.com/docs/guides/end-customer-accounts/creating-charges-for-accounts)
	//     in our main user documentation for more details.
	//
	// Alternatively, you can create a Charge for a Balance on an end-customer Account
	// to create balance fee line items for Account billing. See
	// [Creating Charges for Balances](https://www.m3ter.com/docs/guides/end-customer-accounts/creating-balances-for-accounts/creating-charges-for-balances)
	// in our main user documentation for more details.
	Charges *ChargeService
	// Endpoints for Compound Aggregation related operations such as creation, update,
	// list and delete.
	//
	// Use Compound Aggregations to create numerical measures from usage data by
	// applying a calculation to one or more simple Aggregations or Custom Fields.
	// These numerical measures can then be used as pricing metrics to price your
	// Product Plans, enabling you to implement a wide range of usage-based pricing use
	// cases.
	//
	// You can create two types of Compound Aggregation:
	//
	// **Global**
	//
	//   - Pricing: Not tied to any specific product and can be used to price Plans
	//     belonging to any Product.
	//   - Calculation: can reference all simple Aggregations - both Global simple
	//     Aggregations and any product-specific simple Aggregations.
	//
	// **Product-specific**
	//
	//   - Pricing: belong to a specific Product and can only be used to price Plans
	//     belonging to the same Product.
	//   - Calculation: can reference any simple Aggregations belonging to the same
	//     Product and any Global simple Aggregations.
	//
	// **IMPORTANT!** If a simple Aggregation referenced by a Compound Aggregation has
	// a **Quantity per unit** defined or a **Rounding** defined, these will not be
	// factored into the value used by the calculation. For example, if the simple
	// Aggregation referenced has a base value of 100 and has **Quantity per unit** set
	// at 10, the Compound Aggregation calculation _will use the base value of 100 not
	// 10_.
	//
	// To better understand and use Compound Aggregations, refer to the example
	// [Compound Aggregation Use Case](https://www.m3ter.com/docs/guides/setting-up-usage-data-meters-and-aggregations/compound-aggregations#example-use-case)
	// in the m3ter documentation.
	CompoundAggregations *CompoundAggregationService
	// Endpoints for Contract related operations such as creation, update, list and
	// delete.
	//
	// Contracts are created for Accounts, which are your end-user customers. Contracts
	// can be used for:
	//
	//   - **Accounts Reporting**. To serve your general accounting operations and
	//     processes, you can report on total Contract values for an Account.
	//   - **Contract Billing**. Various billing entities associated with an Account can
	//     be linked to Contracts on the Account to meet your specific Contract billing
	//     use cases.
	Contracts *ContractService
	// Endpoints for listing, creating, retrieving, updating, or deleting Counters.
	//
	// You can create Counters for your m3ter Organization, which can then be used as
	// pricing metrics to apply a unit-based
	// [CounterPricing](https://www.m3ter.com/docs/api#tag/CounterPricing) to Product
	// Plans or Plan Templates for recurring subscription charges on Accounts.
	//
	// Counters can then be used to post
	// [CounterAdjustments](https://www.m3ter.com/docs/api#tag/CounterAdjustments) on
	// your end-customer Accounts.
	//
	// Accounts are then billed in accordance with the CounterPricing on Plans attached
	// to the Accounts and for the actual Counter quantities Accounts subscribe to. See
	// [Recurring Charges: Counters](https://www.m3ter.com/docs/guides/recurring-charges-counters)
	// in our main user documentation for more details.
	Counters *CounterService
	// Endpoints for listing, creating, updating, retrieving, or deleting
	// CounterAdjustments.
	//
	// If you attach a Plan to an Account which is priced using a Counter to apply
	// unit-based pricing, you can then create CounterAdjustments for the Account using
	// that Counter to ensure the Account is billed according to the number of Counter
	// units the Account subscribes to in a given billing period.
	//
	// See
	// [Understanding and Creating Counter Adjustments for Accounts](https://www.m3ter.com/docs/guides/recurring-charges-counters/creating-counter-adjustments-for-accounts)
	// for more information.
	CounterAdjustments *CounterAdjustmentService
	// Endpoints for listing, creating, updating, retrieving, or deleting
	// CounterPricing.
	//
	// Create the CounterPricing for a Plan/PlanTemplate using a Counter, and define a
	// unit-based pricing structure for charging end customer Accounts put on the Plan.
	//
	// See
	// [Creating Counters and Pricing Plans](https://www.m3ter.com/docs/guides/recurring-charges-counters/creating-counters)
	// for more information.
	CounterPricings *CounterPricingService
	// Endpoints for CreditReason operations such as creation, update, list, and
	// delete.
	//
	// You can create CreditReasons for your Organization, and then use them when
	// creating a credit line item on a bill, or applying a product credit to a bill.
	// CreditReasons provide contextual information as to why a credit was applied.
	CreditReasons *CreditReasonService
	// Endpoints for Currency operations such as creation, update, list, and delete.
	// Currencies are stored for your Organization, and can then be used to specify
	// currencies on various entities such as plan groups and plan templates.
	//
	// **IMPORTANT!** The Currencies you want to use in your Organization must be
	// created first.
	//
	// The currency you select for your Organization determines the billing currency
	// and overrides any currency settings in your pricing Plans. For example, if the
	// Organization currency is set to USD and a pricing Plan used for an Account is
	// set to GBP, the bill for an Account using that Plan is calculated in GBP, and
	// then each bill line item converted to USD amounts.
	//
	// Currency conversion rates are setup in the _OrganizationConfig_. For more
	// details, see
	// [Creating and Managing Currencies](https://www.m3ter.com/docs/guides/organization-and-access-management/viewing-and-editing-organization#creating-and-managing-currencies)
	// in the m3ter Documentation.
	Currencies *CurrencyService
	// Endpoints for retrieving and updating Custom Fields at the Organization level
	// for all entities that support them.
	//
	// Custom Fields in m3ter allow you to store custom data in the form of number or
	// string values against m3ter entities in a way that does not directly affect the
	// normal working operation of the m3ter platform. Having this capability to store
	// data in a free-hand fashion can prove very useful in helping you to meet
	// specific usage-based pricing and other operational business use cases.
	//
	// However, you can exploit the values stored on Custom Fields in a more direct way
	// by referencing them in Derived Field and Compound Aggregation calculations.
	// Given the key role these calculations can play when implementing usage-based
	// pricing schema, any Custom Fields you reference will then affect how the
	// platform behaves. Referencing Custom Field values in your calculations offers a
	// much wider scope of options when it comes to resolving complex usage-based
	// pricing use cases.
	//
	// Custom Fields can be added to the following entities at Organizational level:
	//
	// - Organization
	// - Account
	// - AccountPlan
	// - Aggregation
	// - Compound Aggregation
	// - Meter
	// - Product
	// - Plan
	// - PlanTemplate
	// - Contract
	//
	// These all follow the same pattern - a new _(optional)_ field is available on the
	// entity request and response bodies called "customFields" which is a object in
	// this format:
	//
	// ```
	//
	//	"customFields": {
	//			"exampleCustomField1": 7.1,
	//			"exampleCustomField2": "stringValue"
	//	}
	//
	// ```
	//
	// The value for a Custom Field can be a string or a number.
	//
	// **Using Custom Field values in calculations:**
	//
	//   - You can add Custom Fields at two levels - the Organization level and the
	//     individual entity level.
	//   - The Organizational level field provides a default value and _must be added_ if
	//     you want to also add a Custom Field of the same name at the corresponding
	//     individual entity level. If you reference the Custom Field in a calculation,
	//     the value for the individual entity level field is used. If no field is
	//     defined at individual entity level, then the Organization level field value is
	//     used.
	//
	// **Important: Constraints and Exceptions!**
	//
	// **Custom Fields at Organization Level**. Currently, you cannot create Custom
	// Fields at the Organization-level for the following enitites:
	//
	// - Plan Group
	// - Balance
	// - Balance Transaction Schedule
	// - Balance Charge Schedule
	//
	// Therefore you cannot reference the Custom Fields values created at the
	// individual entity level for these entities in your Derived Field or Compound
	// Aggregation calculations.
	//
	// **Derived Field Calculations**. You can _only reference Custom Fields_ for the
	// following entities:
	//
	// - Organization
	// - Meter
	// - Account
	//
	// However, if you are using Meters belonging to _a specific Product_, that is, not
	// _Global Meters_, you can also reference Custom Fields added to a Product in
	// Derived Field calculations.
	//
	// **Compound Aggregation Calculations - Meter Custom Fields**. The value of the
	// _Organization level Meter Custom Field will always be used_, even if you have
	// defined a corresponding field at the individual Meter level.
	//
	// See
	// [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields)
	// in the m3ter documentation for more information.
	CustomFields *CustomFieldService
	// Endpoints for triggering one-off, ad-hoc Data Exports. You can set up and run
	// ad-hoc Exports to export two kinds of data from your m3ter Organization:
	//
	// - Usage data.
	// - Operational data for entities.
	//
	// **Ad-Hoc Export Destinations** When setting up and running an ad-hoc Export:
	//
	//   - You can define one or more Export Destinations - see the
	//     [ExportDestination](https://www.m3ter.com/docs/api#tag/ExportDestination)
	//     section of this API Reference. When the export runs, the data is sent through
	//     to the sepecified Destination. However, the export file is also made available
	//     for you to download it locally.
	//   - You can set up and run Data Exports without defining a Destination. The data
	//     is not exported but the compiled export file is made available for downloading
	//     locally.
	//   - For details on downloading an export file, see the
	//     [Get Data Export File Download URL](https://www.m3ter.com/docs/api#tag/ExportDestination/operation/GenerateDataExportFileDownloadUrl)
	//     endpoint in this API Reference.
	//
	// **Preview Version!** The Data Export feature is currently available only in
	// Preview release version. See
	// [Feature Release Stages](https://www.m3ter.com/docs/guides/getting-started/feature-release-stages)
	// for Preview release definition. ExportAdHoc endpoints will only be available if
	// Data Export has been enabled for your Organization. For more details see
	// [Data Export(Preview)](https://www.m3ter.com/docs/guides/data-exports) in our
	// main User documentation. If you're interested in previewing the Data Export
	// feature, please get in touch with m3ter Support or your m3ter contact.
	DataExports *DataExportService
	// Endpoints for DebitReason operations such as creation, update, list, and delete.
	//
	// You can create DebitReasons for your Organization, and then use them when
	// creating a debit line item on a bill, or applying a product debit to a bill.
	// DebitReasons provide contextual information as to why a debit was applied.
	DebitReasons *DebitReasonService
	// This section provides Endpoints for operations that allow you to retrieve
	// detailed information about individual Events, list all Events or specific Event
	// Types, and explore dynamic fields available for each Event Type.
	//
	// Events encompass specific instances of state changes within the system, such as
	// the creation of a new Prepayment/Commitment for an Account. Each Event is
	// classified under an Event Type framework, providing context about what kind of
	// change occurred to generate the Event.
	//
	// **Events for Configuration and Billing Entities**
	//
	// Many Event Types cover common configuration and billing objects, where the Event
	// is generated for a state change of one of these objects - for when the
	// configuration or billing object is **created**, **deleted**, or **updated**.
	//
	// For example:
	//
	// - configuration.commitment.created
	// - configuration.commitment.deleted
	// - configuration.commitment.updated
	// - configuration.account.created
	// - configuration.account.deleted
	// - configuration.account.updated
	// - billing.bill.created
	// - billing.bill.deleted
	// - billing.bill.created
	//
	// **Events for Errors or Failures**
	//
	// There are also Event Types for certain kinds of error that can occur:
	//
	// - For an Integration:
	//
	//   - validation
	//   - authentication
	//   - perform
	//   - missing account mapping
	//   - disabled
	//
	// - For a Usage Data Ingest Submission:
	//
	//   - validation failure
	//
	// - For Data Export Jobs:
	//   - data export job failure
	//
	// **Scheduled Events**
	//
	// In addition to system-generated Events that occur when a configuration entity
	// undergoes a state change at creation, update, or deletion of the entity, you can
	// use API calls to create and configure _Scheduled Event Configurations_.
	// Scheduled Events are custom Event types, which you can set up by referencing
	// Date/Time fields on configuration and billing entities. See the
	// [ScheduledEventConfigurations](https://www.m3ter.com/docs/api#tag/ScheduledEventConfigurations)
	// section of this API Reference for more details.
	//
	// **Notifications for Events**
	//
	// You can create Notification rules based on Events and these rules can reference
	// and apply calculations to the Event's fields. This allows you to set up
	// customized alerts to be sent out via webhooks when the Event occurs and any
	// conditions you've built into the Notification rule's calculation are satisfied.
	//
	// See the [Notifications](https://www.m3ter.com/docs/api#tag/Notifications)
	// section for more details.
	//
	// **Other Events**
	//
	// When Events occur, they can cause other Events, such as when a Notification is
	// triggered by the Event it is based on. For these Events there are currently two
	// categories:
	//
	// - Notification
	// - IntegrationEvent
	//
	// Also see
	// [Utilizing Events and Notifications](https://www.m3ter.com/docs/guides/utilizing-events-and-notifications)
	// and
	// [Object Definitions and API Calls](https://www.m3ter.com/docs/guides/utilizing-events-and-notifications/object-definitions-and-api-calls)
	// in the m3ter documentation for more guidance.
	Events *EventService
	// Endpoints for managing External Mapping related operations such as creation,
	// update, list and delete.
	//
	// When you integrate your 3rd-party systems with the m3ter platform, a mapping
	// between entities in the local system _(m3ter)_ and external systems is
	// constructed. This _External Mapping_ is crucial in scenarios where data from
	// external systems is consumed or where data from the local system is to be
	// synchronized with external systems.
	//
	// When you are working to set up your Integrations and want to test or
	// troubleshoot your implementation before going live, you might need to create
	// External Mappings manually and, at a later date, edit or delete them.
	ExternalMappings *ExternalMappingService
	// A suite of endpoints for configuring and managing third party integrations
	// within the m3ter platform. The integration endpoints in this section facilitate
	// various operations such as creating, updating, listing, and deletion of
	// integrations.
	//
	// m3ter integrations enable seamless data synchronization and mapping with
	// external systems required in core business processes. These processes often
	// include sales, pricing, billing and invoicing, and general finance.
	//
	// With m3ter integrations, you can establish robust connections with popular
	// business platforms, enhancing your operational capabilities. For example:
	//
	// - Chargebee
	// - Salesforce
	// - Stripe
	// - Netsuite
	// - Paddle
	// - Xero
	// - QuickBooks
	IntegrationConfigurations *IntegrationConfigurationService
	// Endpoints for creating/updating/deleting Lookup Tables.
	//
	// Lookup Tables enable you to manage dynamic data mappings that your calculations
	// reference. Use them for currency conversion, pricing tiers, discount rates, and
	// similar scenarios where you require values to change operationally but for
	// calculation logic to remain constant.
	//
	// **Beta Version!** The Lookup Table feature is currently available in Beta
	// release version. See
	// [Feature Release Stages](https://www.m3ter.com/docs/guides/getting-started/feature-release-stages)
	// for Beta release definition. Lookup Table endpoints will only be available if
	// Lookup Tables have been enabled for your Organization. For more details see
	// [Lookup Tables (Beta)](https://www.m3ter.com/docs/guides/lookup-tables) in our
	// main User documentation.
	LookupTables *LookupTableService
	// Endpoints for listing, creating, updating, retrieving, or deleting Meters.
	//
	// Use Meters to submit usage data for the consumption of your products and
	// services by end customers. This usage data then becomes the basis for setting up
	// usage-based pricing for your products and services.
	//
	// Examples of usage data collected in Meters:
	//
	// - Number of logins.
	// - Duration of session.
	// - Amount of data downloaded.
	//
	// To collect usage data and ingest it into the platform, you can define two types
	// of fields for Meters:
	//
	//   - `dataFields` Used to collect raw usage data measures - numeric quantitative
	//     data values or non-numeric point data values.
	//   - `derivedFields` Used to derive usage data measures that are the result of
	//     applying a calculation to `dataFields`, `customFields`, or system `Timestamp`
	//     fields.
	//
	// You can also:
	//
	//   - Create `customFields` for a Meter, which allows you to attach custom data to
	//     the Meter as name/value pairs.
	//   - Create Global Meters, which are not tied to a specific Product and allow you
	//     to collect usage data that will form the basis of usage-based pricing across
	//     multiple Products.
	//
	// **IMPORTANT! - use of PII:** The use of any of your end-customers' Personally
	// Identifiable Information (PII) in m3ter is restricted to a few fields on the
	// **Account** entity. Please ensure that any fields you configure for Meters, such
	// as Data Fields or Derived Fields, do not contain any end-customer PII data. See
	// the [Introduction section](https://www.m3ter.com/docs/api#section/Introduction)
	// above for more details.
	//
	// See also:
	//
	// - [Reviewing Meter Options](https://www.m3ter.com/docs/guides/setting-up-usage-data-meters-and-aggregations/reviewing-meter-options).
	Meters *MeterService
	// This section provides endpoints for managing Event Notifications.
	//
	// You can create Notifications based on system Events generated by the platform.
	// When you base a Notification on a specific Event type, you can include a
	// calculation that references the fields available on that Event type to define
	// precise conditions that must be met for the Notification to be triggered when an
	// Event of that type occurs. In this way, you can set up highly customized
	// Notifications that act as timely alerts to inform you about significant
	// occurrences within your Organization. For instance, if you provide a sign-up
	// bonus to new end-customer Accounts, you can set up a Notification to alert you
	// when an end-customer Account has used up a certain percentage of their bonus
	// credit.
	//
	// You can also set up Notifications based on Scheduled Event types you've created
	// for your Organization. See the
	// [ScheduledEventConfigurations](https://www.m3ter.com/docs/api#tag/ScheduledEventConfigurations)
	// section of this API Reference and
	// [Working with Scheduled Events](https://www.m3ter.com/docs/guides/alerts-events-and-notifications/utilizing-events-and-notifications/working-with-scheduled-events)
	// in our user documentation.
	//
	// For more details on Event types and their fields, see the
	// [Events](https://www.m3ter.com/docs/api#tag/Events) section.
	//
	// For detailed guidance on working with Events and Notifications, refer to the
	// [Utilizing Events and Notifications](https://www.m3ter.com/docs/guides/utilizing-events-and-notifications)
	// section of the m3ter user documentation.
	NotificationConfigurations *NotificationConfigurationService
	// Endpoints for retrieving or updating the Organization Config.
	//
	// Organization represents your company as a direct customer of m3ter. Use
	// Organization configuration to define _Organization-wide_ settings. For example:
	//
	// - Timezone.
	// - Currencies and currency conversions.
	// - Billing operations settings, such as:
	//   - Epoch dates to control first billing dates.
	//   - Whether to bill customer accounts in advance/in arrears for standing charge
	//     amounts, minimum spend amounts, and commitment fees.
	//
	// For other aspects of your Organization setup and configuration, see the
	// following sections in this API Reference:
	//
	// - [Custom Fields](https://www.m3ter.com/docs/api#tag/CustomField)
	// - [Currencies](https://www.m3ter.com/docs/api#tag/Currency)
	// - [Credit Reasons](https://www.m3ter.com/docs/api#tag/CreditReason)
	// - [Debit Reason](https://www.m3ter.com/docs/api#tag/DebitReason)
	// - [Transaction Types](https://www.m3ter.com/docs/api#tag/TransactionType)
	//
	// See also:
	//
	// - [Managing your Organization](https://www.m3ter.com/docs/guides/managing-organization-and-users/viewing-and-editing-organization).
	OrganizationConfig *OrganizationConfigService
	// Endpoints for Permission Policy related operations such as creation, update, add
	// and retrieve.
	//
	// Permission Policies can restrict or grant access to specific resources for both
	// Users _(people)_ and Service Users _(automated processes with direct API
	// access)_. This enables you to control precisely what a User can do in your m3ter
	// Organization.
	//
	// For more details, see
	// [Understanding, Creating, and Managing Permission Policies](https://www.m3ter.com/docs/guides/organization-and-access-management/creating-and-managing-permissions#permission-policy-statements---available-actions-and-resources)
	// in our main Documentation.
	PermissionPolicies *PermissionPolicyService
	// Endpoints for listing, creating, updating, retrieving, or deleting Plans.
	//
	// A Plan is based on a PlanTemplate and represents a specific pricing plan for one
	// of your products or services. Each Plan inherits general billing attributes or
	// pricing structure from its parent Plan Template. Some attributes can be
	// overriden for the specific Plan.
	//
	// When you've created the Plan Templates and Plans you need for your Products, you
	// can configure the exact pricing structures for Plans to charge customers that
	// consume one or more of your Products.
	//
	// You can then attach the appropriately priced Plans to customer Accounts to
	// create [Account Plans](https://www.m3ter.com/docs/api#tag/AccountPlan) and
	// enable charges to be calculated correctly for billing against those Accounts.
	//
	// See also:
	//
	// - [Reviewing Options for Plans and Plan Templates](https://www.m3ter.com/docs/guides/working-with-plan-templates-and-plans/reviewing-configuration-options-for-plans-and-plan-templates).
	Plans *PlanService
	// Endpoints for PlanGroup related operations such as creation, update, retrieve,
	// list and delete.
	//
	// PlanGroups are constructs that group multiple plans together. This enables a
	// unified approach to efficiently handle various uses cases across different
	// plans. For example applying a minimum spend amount at billing, across several of
	// your products or features that are each priced separately.
	PlanGroups *PlanGroupService
	// Endpoints for PlanGroupLink related operations such as creation, update, list
	// and delete.
	//
	// PlanGroupLinks are the intersection table between a PlanGroup and its associated
	// Plans. A PlanGroupLink is only created when at least 1 Plan is linked to a
	// PlanGroup.
	PlanGroupLinks *PlanGroupLinkService
	// Endpoints for listing, creating, updating, retrieving, or deleting
	// PlanTemplates.
	//
	// Use PlanTemplates to define default values for Plans. These default values
	// control the billing operations you want applied to your products. PlanTemplates
	// avoid repetition in configuration work - many Plans will share settings for
	// billing operations and differ only in the details of their pricing structures.
	//
	// A PlanTemplate is linked to a Product, and each Plan is a child of a
	// PlanTemplate.
	PlanTemplates *PlanTemplateService
	// Endpoints for listing, creating, updating, retrieving, or deleting Pricing.
	//
	// Create the Pricing for a Plan/PlanTemplate with usage data Aggregations, and
	// define a usage-based pricing structure for charging end customer Accounts put on
	// the Plan.
	//
	// See
	// [Reviewing Pricing Options for Plans and Plan Templates](https://www.m3ter.com/docs/guides/pricing-plans/reviewing-pricing-options-and-pricing-plans)
	// for more information.
	Pricings *PricingService
	// Endpoints for listing, creating, updating, retrieving, or deleting Products.
	//
	// A Product represents the products and services you offer to your end customers.
	// Products act as a container for the Meters, Aggregations, Pricing, and Plans
	// required to implement usage-based and other pricing models for your
	// Organization.
	Products *ProductService
	// Endpoints for ResourceGroup related operations such as creation, update, list
	// and delete.
	//
	// ResourceGroups are used in the context of Permission Policies, which controls
	// what a User who has been given access to your Organization can and cannot do.
	// For example, you might want to create a Permissions Policy that denies Users the
	// ability to retrieve Meters.
	//
	// Resources are defined as m3ter Resource Identifiers _(MRIs)_ in the format:
	//
	// ```
	// service:resource-type/item-type/id
	// ```
	//
	// Where:
	//
	//   - service is a distinct part of the overall m3ter system, and which forms a
	//     natural functional grouping, such as "config" or "billing".
	//
	//   - resource-type is the resource type item accessed - for example: "Plan",
	//     "Meter", "Bill"
	//
	// - item-type is one of:
	//
	//   - "item" - to specify an individual item.
	//
	//   - "group" - to specify a resource group.
	//
	// - id is the resource group id or the resource item id
	//
	// Resources can be assigned to one or more ResourceGroups. For example, a Plan can
	// be assigned to Plan ResourceGroups, a Meter can be assigned to Meter
	// ResourceGroups, and so on. This is useful for cases where you want to create
	// Permission Policies which allow or deny access to a specific subset of
	// resources. For example, grant a user access to only some of the Plans in your
	// Organization.
	//
	// This concept of grouping resources applies to every resource in m3ter, including
	// ResourceGroups themselves. This allows you to nest ResourceGroups to support
	// hierarchies of groups.
	//
	// See
	// [Understanding, Creating, and Managing Permission Policies](https://www.m3ter.com/docs/guides/managing-organization-and-users/creating-and-managing-permissions)
	// in the m3ter documentation for more information.
	//
	// **Note: User Resource Groups** You can create a User Resource Group to group
	// resources of type = `user`. You can then retrieve a list of the User Resource
	// Groups a user belongs to. For more details, see the
	// [Retrieve OrgUser Groups](https://www.m3ter.com/docs/api#tag/OrgUsers/operation/GetOrgUserGroups)
	// call in the OrgUsers section.
	ResourceGroups *ResourceGroupService
	// Endpoints for retrieving and managing scheduled Events' configurations.
	//
	// Scheduled Event Configurations define custom Event types that reference
	// Date/Time fields belonging to configuration and billing entities. They therefore
	// provide you with an extra degree of flexibility over and above system-generated
	// Events for setting up Notifications based on Events.
	//
	// For more details, see the
	// [Working with Scheduled Events](https://www.m3ter.com/docs/guides/alerts-events-and-notifications/utilizing-events-and-notifications/working-with-scheduled-events)
	// in our Documenation.
	ScheduledEventConfigurations *ScheduledEventConfigurationService
	// Endpoints for billing operations such as creating, updating,
	// listing,downloading, and deleting Bills.
	//
	// Bills are generated for an Account, and are calculated in accordance with the
	// usage-based pricing Plans applied for the Products the Account consumes. These
	// endpoints enable interaction with the billing system, allowing you to obtain
	// billing details and insights into the consumption patterns and charges of your
	// end-customer Accounts.
	Statements *StatementService
	// Endpoints for TransactionType operations such as creation, update, list,
	// retrieve, and delete.
	//
	// You can create TransactionTypes for your Organization, which can then be used
	// when creating and updating Balances. Example TransactionTypes: "Balance Amount"
	// or "Add Funds".
	//
	// For details on creating a Transaction amount for a Balance using a
	// TransactionType you've created for your Organization, see the
	// [Create Balance Transaction](https://www.m3ter.com/docs/api#tag/Balances/operation/PostBalanceTransaction)
	// call in the [Balances](https://www.m3ter.com/docs/api#tag/Balances) section of
	// this API Reference.
	TransactionTypes *TransactionTypeService
	Usage            *UsageService
	Users            *UserService
	// A suite of endpoints for configuring and managing third party integrations
	// within the m3ter platform. The integration endpoints in this section facilitate
	// various operations such as creating, updating, listing, and deletion of
	// integrations.
	//
	// m3ter integrations enable seamless data synchronization and mapping with
	// external systems required in core business processes. These processes often
	// include sales, pricing, billing and invoicing, and general finance.
	//
	// With m3ter integrations, you can establish robust connections with popular
	// business platforms, enhancing your operational capabilities. For example:
	//
	// - Chargebee
	// - Salesforce
	// - Stripe
	// - Netsuite
	// - Paddle
	// - Xero
	// - QuickBooks
	Webhooks *WebhookService
}

// DefaultClientOptions read from the environment (M3TER_API_KEY, M3TER_API_SECRET,
// M3TER_API_TOKEN, M3TER_ORG_ID, M3TER_BASE_URL). This should be used to
// initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("M3TER_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
	if o, ok := os.LookupEnv("M3TER_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("M3TER_API_SECRET"); ok {
		defaults = append(defaults, option.WithAPISecret(o))
	}
	if o, ok := os.LookupEnv("M3TER_API_TOKEN"); ok {
		defaults = append(defaults, option.WithToken(o))
	}
	if o, ok := os.LookupEnv("M3TER_ORG_ID"); ok {
		defaults = append(defaults, option.WithOrgID(o))
	}
	return defaults
}

// NewClient generates a new client with the default option read from the
// environment (M3TER_API_KEY, M3TER_API_SECRET, M3TER_API_TOKEN, M3TER_ORG_ID,
// M3TER_BASE_URL). The option passed in as arguments are applied after these
// default arguments, and all option will be passed down to the services and
// requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	opts = append(DefaultClientOptions(), opts...)

	r = &Client{Options: opts}

	r.Authentication = NewAuthenticationService(opts...)
	r.Accounts = NewAccountService(opts...)
	r.AccountPlans = NewAccountPlanService(opts...)
	r.Aggregations = NewAggregationService(opts...)
	r.Balances = NewBalanceService(opts...)
	r.Bills = NewBillService(opts...)
	r.BillConfig = NewBillConfigService(opts...)
	r.Commitments = NewCommitmentService(opts...)
	r.BillJobs = NewBillJobService(opts...)
	r.Charges = NewChargeService(opts...)
	r.CompoundAggregations = NewCompoundAggregationService(opts...)
	r.Contracts = NewContractService(opts...)
	r.Counters = NewCounterService(opts...)
	r.CounterAdjustments = NewCounterAdjustmentService(opts...)
	r.CounterPricings = NewCounterPricingService(opts...)
	r.CreditReasons = NewCreditReasonService(opts...)
	r.Currencies = NewCurrencyService(opts...)
	r.CustomFields = NewCustomFieldService(opts...)
	r.DataExports = NewDataExportService(opts...)
	r.DebitReasons = NewDebitReasonService(opts...)
	r.Events = NewEventService(opts...)
	r.ExternalMappings = NewExternalMappingService(opts...)
	r.IntegrationConfigurations = NewIntegrationConfigurationService(opts...)
	r.LookupTables = NewLookupTableService(opts...)
	r.Meters = NewMeterService(opts...)
	r.NotificationConfigurations = NewNotificationConfigurationService(opts...)
	r.OrganizationConfig = NewOrganizationConfigService(opts...)
	r.PermissionPolicies = NewPermissionPolicyService(opts...)
	r.Plans = NewPlanService(opts...)
	r.PlanGroups = NewPlanGroupService(opts...)
	r.PlanGroupLinks = NewPlanGroupLinkService(opts...)
	r.PlanTemplates = NewPlanTemplateService(opts...)
	r.Pricings = NewPricingService(opts...)
	r.Products = NewProductService(opts...)
	r.ResourceGroups = NewResourceGroupService(opts...)
	r.ScheduledEventConfigurations = NewScheduledEventConfigurationService(opts...)
	r.Statements = NewStatementService(opts...)
	r.TransactionTypes = NewTransactionTypeService(opts...)
	r.Usage = NewUsageService(opts...)
	r.Users = NewUserService(opts...)
	r.Webhooks = NewWebhookService(opts...)

	return
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// If a byte slice or an [io.Reader] is supplied to params, it will be used as-is
// for the request body.
//
// The params is by default serialized into the body using [encoding/json]. If your
// type implements a MarshalJSON function, it will be used instead to serialize the
// request. If a URLQuery method is implemented, the returned [url.Values] will be
// used as query strings to the url.
//
// If your params struct uses [param.Field], you must provide either [MarshalJSON],
// [URLQuery], and/or [MarshalForm] functions. It is undefined behavior to use a
// struct uses [param.Field] without specifying how it is serialized.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url.
//
// The response body will be deserialized into the res variable, depending on its
// type:
//
//   - A pointer to a [*http.Response] is populated by the raw response.
//   - A pointer to a byte array will be populated with the contents of the request
//     body.
//   - A pointer to any other type uses this library's default JSON decoding, which
//     respects UnmarshalJSON if it is defined on the type.
//   - A nil value will not read the response body.
//
// For even greater flexibility, see [option.WithResponseInto] and
// [option.WithResponseBodyInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	opts = slices.Concat(r.Options, opts)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}
