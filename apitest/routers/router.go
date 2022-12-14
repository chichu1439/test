// @APIVersion 1.0.0
// @Title PaymentSwitch Api
// @Description PaymentSwitch solution
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://10.20.7.91:30002/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"apitest/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("",
		//beego.NSNamespace("/v1/bnpl/order-and-repayment-plan", beego.NSInclude(&controllers.BP000001Controller{})),
		//beego.NSNamespace("/v1/bnpl/repayment-plan/trial-calculation", beego.NSInclude(&controllers.BP000002Controller{})),
		//beego.NSNamespace("/v1/bnpl/payment/confirm", beego.NSInclude(&controllers.BP000003Controller{})),
		//beego.NSNamespace("/v1/bnpl/loan/list", beego.NSInclude(&controllers.BP000004Controller{})),
		//beego.NSNamespace("/v1/bnpl/loan/detail", beego.NSInclude(&controllers.BP000005Controller{})),
		//beego.NSNamespace("/v1/bnpl/repayment/plan", beego.NSInclude(&controllers.BP000006Controller{})),
		//beego.NSNamespace("/v1/bnpl/repayment/calculate", beego.NSInclude(&controllers.BP000007Controller{})),
		//beego.NSNamespace("/v1/bnpl/repayment/confirm", beego.NSInclude(&controllers.BP000008Controller{})),
		//beego.NSNamespace("/v1/bnpl/order", beego.NSInclude(&controllers.BP000031Controller{})),
		//beego.NSNamespace("/v1/bnpl/customer", beego.NSInclude(&controllers.BP000032Controller{})),
		//beego.NSNamespace("/v1/bnpl/loan-product", beego.NSInclude(&controllers.BP000034Controller{})),
		//beego.NSNamespace("/v1/bnpl/repayment/method", beego.NSInclude(&controllers.BP000035Controller{})),

		//beego.NSNamespace("/v1/MaaiPrepaid/document/query", beego.NSInclude(&controllers.SMP0000001Controller{})),
		//beego.NSNamespace("/v1/MaaiPrepaid/appVerion/query", beego.NSInclude(&controllers.SMP0000002Controller{})),
		//beego.NSNamespace("/v1/MaaiPrepaid/tc/query", beego.NSInclude(&controllers.SMP0000004Controller{})),
		//beego.NSNamespace("/v1/MaaiPrepaid/otp/send", beego.NSInclude(&controllers.SMP0000005Controller{})),
		//beego.NSNamespace("/v1/MaaiPrepaid/otp/verify", beego.NSInclude(&controllers.SMP0000006Controller{})),
		//beego.NSNamespace("/v1/MaaiPrepaid/otp/resend", beego.NSInclude(&controllers.SMP0000007Controller{})),
		//beego.NSNamespace("/v1/MaaiPrepaid/phoneState/check", beego.NSInclude(&controllers.SMP0000008Controller{})),
		//beego.NSNamespace("/v1/MaaiPrepaid/agreement/sign", beego.NSInclude(&controllers.SMP0000009Controller{})),
		//beego.NSNamespace("/v1/MaaiPrepaid/citizenIDpicture/record", beego.NSInclude(&controllers.SMP0000010Controller{})),
		//beego.NSNamespace("/v1/MaaiPrepaid/citizenID/check", beego.NSInclude(&controllers.SMP0000011Controller{})),
		//beego.NSNamespace("/v1/MaaiPrepaid/selfileImage/send", beego.NSInclude(&controllers.SMP0000012Controller{})),
		////beego.NSNamespace("/v1/MaaiPrepaid/publicKey/query", beego.NSInclude(&controllers.SMP0000013Controller{})),
		//beego.NSNamespace("/v1/MaaiPrepaid/register/Usepin", beego.NSInclude(&controllers.SMP0000014Controller{})),
		//beego.NSNamespace("/v1/MaaiPrepaid/walletAccont/query", beego.NSInclude(&controllers.SMP0000015Controller{})),
		////beego.NSNamespace("/v1/MaaiPrepaid/allNotification/query", beego.NSInclude(&controllers.SMP0000016Controller{})),
		//beego.NSNamespace("/v1/MaaiPrepaid/verify/pin", beego.NSInclude(&controllers.SMP0000017Controller{})),
		//beego.NSNamespace("/v1/MaaiPrepaid/reset/pin", beego.NSInclude(&controllers.SMP0000018Controller{})),
		//beego.NSNamespace("/v1/MaaiPrepaid/login/bycustomerNo", beego.NSInclude(&controllers.SMP0000019Controller{})),
		////beego.NSNamespace("/v1/MaaiPrepaid/Notification/send", beego.NSInclude(&controllers.SMP0000021Controller{})),
		////beego.NSNamespace("/v1/MaaiPrepaid/Limit/query", beego.NSInclude(&controllers.SMP0000022Controller{})),
		////beego.NSNamespace("/v1/MaaiPrepaid/Limit/modify", beego.NSInclude(&controllers.SMP0000023Controller{})),
		////beego.NSNamespace("/v1/MaaiPrepaid/personalInfo/query", beego.NSInclude(&controllers.SMP0000024Controller{})),
		////beego.NSNamespace("/v1/MaaiPrepaid/transactionHistory/query", beego.NSInclude(&controllers.SMP0000027Controller{})),
		////beego.NSNamespace("/v1/MaaiPrepaid/notification/query", beego.NSInclude(&controllers.SMP0000028Controller{})),
		//beego.NSNamespace("/v1/MaaiPrepaid/citizenInformation/query", beego.NSInclude(&controllers.SMP0000029Controller{})),
		//beego.NSNamespace("/v1/MaaiPrepaid/money/transfer", beego.NSInclude(&controllers.SMP0000030Controller{})),
		//beego.NSNamespace("/v1/MaaiPrepaid/myqr/query", beego.NSInclude(&controllers.SMP0000031Controller{})),
		//

		//beego.NSNamespace("/v1/icredit/repayment_amount_trial", beego.NSInclude(&controllers.AL000003Controller{})),
		//beego.NSNamespace("/v1/icredit_bnpl/new_customer", beego.NSInclude(&controllers.CU000061Controller{})),
		//beego.NSNamespace("/v1/icredit_bnpl/new_loan", beego.NSInclude(&controllers.IC000069Controller{})),
		//beego.NSNamespace("/v1/icredit_bnpl/repayment", beego.NSInclude(&controllers.IC000070Controller{})),
		//beego.NSNamespace("/v1/icredit_bnpl/trial_repayment_plan", beego.NSInclude(&controllers.IC000071Controller{})),
		//beego.NSNamespace("/v1/icredit_bnpl/loan_list", beego.NSInclude(&controllers.IC000072Controller{})),
		//beego.NSNamespace("/v1/icredit_bnpl/loan_detail", beego.NSInclude(&controllers.IC000073Controller{})),

		//beego.NSNamespace("/v1/icredit/writeoff/apply", beego.NSInclude(&controllers.AL000012Controller{})),
		//beego.NSNamespace("/v1/icredit/writeoff/approve", beego.NSInclude(&controllers.AL000013Controller{})),
		//beego.NSNamespace("/v1/icredit/query-write-off-register-list", beego.NSInclude(&controllers.AL000014Controller{})),
		//beego.NSNamespace("/v1/icredit/query-write-off-flow-record", beego.NSInclude(&controllers.AL000015Controller{})),
		//beego.NSNamespace("/v1/icredit/writeoff/cancel", beego.NSInclude(&controllers.AL000018Controller{})),
		//beego.NSNamespace("/v1/icredit/query-loan-account-period-info", beego.NSInclude(&controllers.AL000054Controller{})),
		//beego.NSNamespace("/v1/icredit/query-write-off-status", beego.NSInclude(&controllers.AL000055Controller{})),
		//beego.NSNamespace("/v1/2c2p/payment-link-inquiry", beego.NSInclude(&controllers.AL000060Controller{})),
		//beego.NSNamespace("/v1/2c2p/payment-info-encrypt", beego.NSInclude(&controllers.AL000061Controller{})),
		//beego.NSNamespace("/v1/2c2p/payment-status-inquiry", beego.NSInclude(&controllers.AL000063Controller{})),
		//beego.NSNamespace("/v1/2c2p/payment", beego.NSInclude(&controllers.AL000065Controller{})),
		//beego.NSNamespace("/v1/common/file-upload", beego.NSInclude(&controllers.CM000006Controller{})),
		//beego.NSNamespace("/v1/common/file-download", beego.NSInclude(&controllers.CM000007Controller{})),
		//beego.NSNamespace("/v1/common/get-option-info", beego.NSInclude(&controllers.CM000021Controller{})),
		//beego.NSNamespace("/v1/common/random-code-generate", beego.NSInclude(&controllers.CM000040Controller{})),
		//beego.NSNamespace("/v1/common/random-code-delete", beego.NSInclude(&controllers.CM000041Controller{})),
		//beego.NSNamespace("/v1/common/key-maintain", beego.NSInclude(&controllers.CM000044Controller{})),
		//beego.NSNamespace("/v1/icredit/collection/summarystatus", beego.NSInclude(&controllers.CM000052Controller{})),
		//beego.NSNamespace("/v1/customer/check-newcustomer", beego.NSInclude(&controllers.CU000001Controller{})),
		//beego.NSNamespace("/v1/customer/create-new-customer", beego.NSInclude(&controllers.CU000002Controller{})),
		//beego.NSNamespace("/v1/customer/query-customer-information", beego.NSInclude(&controllers.CU000003Controller{})),
		//beego.NSNamespace("/v1/customer/modify-customer", beego.NSInclude(&controllers.CU000004Controller{})),
		//beego.NSNamespace("/v1/customer/create/customer/contract", beego.NSInclude(&controllers.CU000005Controller{})),
		//beego.NSNamespace("/v1/customer/query/customer/status", beego.NSInclude(&controllers.CU000006Controller{})),
		//beego.NSNamespace("/v1/customer/contract/count", beego.NSInclude(&controllers.CU000007Controller{})),
		//beego.NSNamespace("/v1/customer/query/customer/contract/info", beego.NSInclude(&controllers.CU000008Controller{})),
		//beego.NSNamespace("/v1/customer/update/customer/contract/info", beego.NSInclude(&controllers.CU000009Controller{})),
		//beego.NSNamespace("/v1/customer/update/customer/status", beego.NSInclude(&controllers.CU000010Controller{})),
		//beego.NSNamespace("/v1/customer/customer-count", beego.NSInclude(&controllers.CU000011Controller{})),
		//beego.NSNamespace("/v1/customer/customer-list", beego.NSInclude(&controllers.CU000012Controller{})),
		//beego.NSNamespace("/v1/icredit/repayment_application", beego.NSInclude(&controllers.IC000001Controller{})),
		//beego.NSNamespace("/v1/icredit/repayment_approve", beego.NSInclude(&controllers.IC000002Controller{})),
		//beego.NSNamespace("/v1/icredit/loan-application", beego.NSInclude(&controllers.IC000006Controller{})),
		//beego.NSNamespace("/v1/icredit/loan-approve", beego.NSInclude(&controllers.IC000007Controller{})),
		//beego.NSNamespace("/v1/icredit/customer-detail-inquiry", beego.NSInclude(&controllers.IC000008Controller{})),
		//beego.NSNamespace("/v1/icredit/receive-customer", beego.NSInclude(&controllers.IC000009Controller{})),
		//beego.NSNamespace("/v1/icredit/approve-new-customer", beego.NSInclude(&controllers.IC000010Controller{})),
		//beego.NSNamespace("/v1/icredit/approve-modify-customer", beego.NSInclude(&controllers.IC000011Controller{})),
		//beego.NSNamespace("/v1/icredit/customer-pdf-email", beego.NSInclude(&controllers.IC000012Controller{})),
		//beego.NSNamespace("/v1/icredit/query_balance_statistics", beego.NSInclude(&controllers.IC000013Controller{})),
		//beego.NSNamespace("/v1/customer/query_customer_statistics", beego.NSInclude(&controllers.IC000014Controller{})),
		//beego.NSNamespace("/v1/icredit/query_risk_classification", beego.NSInclude(&controllers.IC000015Controller{})),
		//beego.NSNamespace("/v1/icredit/modify-customer-application", beego.NSInclude(&controllers.IC000017Controller{})),
		//beego.NSNamespace("/v1/icredit/repayment_plan_trial", beego.NSInclude(&controllers.IC000018Controller{})),
		//beego.NSNamespace("/v1/icredit/query-loan-list", beego.NSInclude(&controllers.IC000022Controller{})),
		//beego.NSNamespace("/v1/icredit/query-loan-detail", beego.NSInclude(&controllers.IC000023Controller{})),
		//beego.NSNamespace("/v1/icredit/loan_transaction", beego.NSInclude(&controllers.IC000024Controller{})),
		//beego.NSNamespace("/v1/icredit/manual-disbursement/application", beego.NSInclude(&controllers.IC000025Controller{})),
		//beego.NSNamespace("/v1/icredit/manual-disbursement/approve", beego.NSInclude(&controllers.IC000026Controller{})),
		//beego.NSNamespace("/v1/icredit/repayment-plan/query", beego.NSInclude(&controllers.IC000027Controller{})),
		//beego.NSNamespace("/v1/icredit/modify-loan-detail", beego.NSInclude(&controllers.IC000028Controller{})),
		//beego.NSNamespace("/v1/icredit/approve-modify-loan-detail", beego.NSInclude(&controllers.IC000029Controller{})),
		//beego.NSNamespace("/v1/icredit/loanpdf", beego.NSInclude(&controllers.IC000030Controller{})),
		//beego.NSNamespace("/v1/icredit/customer-history/query", beego.NSInclude(&controllers.IC000031Controller{})),
		//beego.NSNamespace("/v1/icredit/query-loan-change-history", beego.NSInclude(&controllers.IC000032Controller{})),
		//beego.NSNamespace("/v1/icredit/loan_transaction_detail", beego.NSInclude(&controllers.IC000034Controller{})),
		//beego.NSNamespace("/v1/icredit/loan_transaction_delete", beego.NSInclude(&controllers.IC000035Controller{})),
		//beego.NSNamespace("/v1/icredit/collection/newtask", beego.NSInclude(&controllers.IC000036Controller{})),
		//beego.NSNamespace("/v1/icredit/collection/tasklist", beego.NSInclude(&controllers.IC000037Controller{})),
		//beego.NSNamespace("/v1/icredit/collection/taskdetail", beego.NSInclude(&controllers.IC000038Controller{})),
		//beego.NSNamespace("/v1/icredit/collection/updtask", beego.NSInclude(&controllers.IC000039Controller{})),
		//beego.NSNamespace("/v1/icredit/collection/summaryamt", beego.NSInclude(&controllers.IC000040Controller{})),
		//beego.NSNamespace("/icredit/v1/collection/summarystatus", beego.NSInclude(&controllers.IC000041Controller{})),
		//beego.NSNamespace("/v1/icredit/collection/newcontact", beego.NSInclude(&controllers.IC000042Controller{})),
		//beego.NSNamespace("/v1/icredit/collection/contactlist", beego.NSInclude(&controllers.IC000043Controller{})),
		//beego.NSNamespace("/v1/icredit/collection/contactdetail", beego.NSInclude(&controllers.IC000044Controller{})),
		//beego.NSNamespace("/v1/icredit/new-loan-daily-monthly-report", beego.NSInclude(&controllers.IC000049Controller{})),
		//beego.NSNamespace("/v1/icredit/closed-loan-monthly-report", beego.NSInclude(&controllers.IC000050Controller{})),
		//beego.NSNamespace("/v1/icredit/non-performing-loan-monthly-report", beego.NSInclude(&controllers.IC000052Controller{})),
		//beego.NSNamespace("/v1/icredit/report-file-list-query", beego.NSInclude(&controllers.IC000053Controller{})),
		//beego.NSNamespace("/v1/icredit/loan-extension/modify", beego.NSInclude(&controllers.IC000054Controller{})),
		//beego.NSNamespace("/v1/icredit/query-to-be-write-off-account-list", beego.NSInclude(&controllers.IC000055Controller{})),
		//beego.NSNamespace("/v1/icredit/quiry-customer-quota", beego.NSInclude(&controllers.IC000059Controller{})),
		//beego.NSNamespace("/v1/icredit/customerproductzcreditlist/query", beego.NSInclude(&controllers.IC000060Controller{})),
		//beego.NSNamespace("/v1/icredit/customerproductzcreditdetail/query", beego.NSInclude(&controllers.IC000061Controller{})),
		//beego.NSNamespace("/v1/icredit/customerproductcreditline/application", beego.NSInclude(&controllers.IC000062Controller{})),
		//beego.NSNamespace("/v1/banking-unc/template/create", beego.NSInclude(&controllers.NT000001Controller{})),
		//beego.NSNamespace("/v1/banking-unc/template/query-list", beego.NSInclude(&controllers.NT000002Controller{})),
		//beego.NSNamespace("/v1/banking-unc/template/query-detail", beego.NSInclude(&controllers.NT000003Controller{})),
		//beego.NSNamespace("/v1/banking-unc/template/update", beego.NSInclude(&controllers.NT000004Controller{})),
		//beego.NSNamespace("/v1/banking-unc/template/delete", beego.NSInclude(&controllers.NT000005Controller{})),
		//beego.NSNamespace("/v1/banking-unc/strategy/create", beego.NSInclude(&controllers.NT000006Controller{})),
		//beego.NSNamespace("/v1/banking-unc/strategy/query-list", beego.NSInclude(&controllers.NT000007Controller{})),
		//beego.NSNamespace("/v1/banking-unc/strategy/query-detail", beego.NSInclude(&controllers.NT000008Controller{})),
		//beego.NSNamespace("/v1/banking-unc/strategy/update", beego.NSInclude(&controllers.NT000009Controller{})),
		//beego.NSNamespace("/v1/banking-unc/strategy/delete", beego.NSInclude(&controllers.NT000010Controller{})),
		//beego.NSNamespace("/v1/banking-unc/target/add", beego.NSInclude(&controllers.NT000011Controller{})),
		//beego.NSNamespace("/v1/banking-unc/target/update", beego.NSInclude(&controllers.NT000012Controller{})),
		//beego.NSNamespace("/v1/banking-unc/message/send", beego.NSInclude(&controllers.NT000013Controller{})),
		//beego.NSNamespace("/v1/product/create-new-category", beego.NSInclude(&controllers.PD000001Controller{})),
		//beego.NSNamespace("/v1/product/query-category-tree", beego.NSInclude(&controllers.PD000002Controller{})),
		//beego.NSNamespace("/v1/product/modify-category", beego.NSInclude(&controllers.PD000003Controller{})),
		//beego.NSNamespace("/v1/product/delete-category", beego.NSInclude(&controllers.PD000004Controller{})),
		//beego.NSNamespace("/v1/product/query-category-product", beego.NSInclude(&controllers.PD000005Controller{})),
		//beego.NSNamespace("/v1/product/query-product-list", beego.NSInclude(&controllers.PD000006Controller{})),
		//beego.NSNamespace("/v1/product/modify-current-deposit-product", beego.NSInclude(&controllers.PD000009Controller{})),
		//beego.NSNamespace("/v1/product/delete-current-deposit-product", beego.NSInclude(&controllers.PD000010Controller{})),
		//beego.NSNamespace("/v1/product/create-micro-loan-product", beego.NSInclude(&controllers.PD000011Controller{})),
		//beego.NSNamespace("/v1/product/query-micro-loan-product", beego.NSInclude(&controllers.PD000012Controller{})),
		//beego.NSNamespace("/v1/product/modify-micro-loan-product", beego.NSInclude(&controllers.PD000013Controller{})),
		//beego.NSNamespace("/v1/product/delete-micro-loan-product", beego.NSInclude(&controllers.PD000014Controller{})),
		//beego.NSNamespace("/v1/product/query-interest-strategy", beego.NSInclude(&controllers.PD000015Controller{})),
		//beego.NSNamespace("/v1/product/modify-interest-strategy", beego.NSInclude(&controllers.PD000016Controller{})),
		//beego.NSNamespace("/v1/product/query-fee-strategy", beego.NSInclude(&controllers.PD000017Controller{})),
		//beego.NSNamespace("/v1/product/modify-fee-strategy", beego.NSInclude(&controllers.PD000018Controller{})),
		//beego.NSNamespace("/v1/product/create-product-item", beego.NSInclude(&controllers.PD000019Controller{})),
		//beego.NSNamespace("/v1/product/modify-product-item", beego.NSInclude(&controllers.PD000020Controller{})),
		//beego.NSNamespace("/v1/product/delete-product-item", beego.NSInclude(&controllers.PD000021Controller{})),
		//beego.NSNamespace("/v1/product/query-loan-interest-strategy", beego.NSInclude(&controllers.PD000044Controller{})),
		//beego.NSNamespace("/v1/product/modify-loan-interest-strategy", beego.NSInclude(&controllers.PD000045Controller{})),
		//beego.NSNamespace("/v1/product/query-loan-fee-strategy", beego.NSInclude(&controllers.PD000046Controller{})),
		//beego.NSNamespace("/v1/product/modify-loan-fee-strategy", beego.NSInclude(&controllers.PD000047Controller{})),
		//beego.NSNamespace("/v1/product/query-current-deposit-full", beego.NSInclude(&controllers.PD000052Controller{})),
		//beego.NSNamespace("/v1/product/query-micro-loan-full", beego.NSInclude(&controllers.PD000053Controller{})),
		//beego.NSNamespace("/v1/product/modify-category-product", beego.NSInclude(&controllers.PD000054Controller{})),
		//beego.NSNamespace("/v1/product/delete-loan-interest-strategy", beego.NSInclude(&controllers.PD000057Controller{})),
		//beego.NSNamespace("/v1/product/delete-loan-fee-strategy", beego.NSInclude(&controllers.PD000058Controller{})),
		//beego.NSNamespace("/v1/product/create-micro-loan-simple", beego.NSInclude(&controllers.PD000068Controller{})),
		//beego.NSNamespace("/v1/product/query-product-service-pool", beego.NSInclude(&controllers.PD000069Controller{})),
		//beego.NSNamespace("/v1/product/modify-micro-loan-product-simple", beego.NSInclude(&controllers.PD000070Controller{})),
		//beego.NSNamespace("/v1/product/product-notification-strategy-list/query", beego.NSInclude(&controllers.PD000078Controller{})),
		//beego.NSNamespace("/v1/product/product-fee-his/query", beego.NSInclude(&controllers.PD000080Controller{})),
		//beego.NSNamespace("/v1/product/product-limit-his/query", beego.NSInclude(&controllers.PD000082Controller{})),
		//beego.NSNamespace("/v1/product/product-fee-info/query", beego.NSInclude(&controllers.PD000084Controller{})),
		//beego.NSNamespace("/v1/product/product-interest-his/query", beego.NSInclude(&controllers.PD000087Controller{})),
		//beego.NSNamespace("/v1/product/limit-control", beego.NSInclude(&controllers.PD000089Controller{})),
		//beego.NSNamespace("/v1/product/product-examine", beego.NSInclude(&controllers.PD000096Controller{})),
		//beego.NSNamespace("/v1/product/change/lastrecord", beego.NSInclude(&controllers.PD000100Controller{})),
		//beego.NSNamespace("/v1/pricing/fee/create-fee-item", beego.NSInclude(&controllers.PF000001Controller{})),
		//beego.NSNamespace("/v1/pricing/fee/query-fee-item", beego.NSInclude(&controllers.PF000002Controller{})),
		//beego.NSNamespace("/v1/pricing/fee/modify-fee-item", beego.NSInclude(&controllers.PF000003Controller{})),
		//beego.NSNamespace("/v1/pricing/fee/delete-fee-item", beego.NSInclude(&controllers.PF000004Controller{})),
		//beego.NSNamespace("/v1/pricing/fee/create-strategy", beego.NSInclude(&controllers.PF000005Controller{})),
		//beego.NSNamespace("/v1/pricing/fee/query-strategy-list", beego.NSInclude(&controllers.PF000006Controller{})),
		//beego.NSNamespace("/v1/pricing/fee/modify-strategy", beego.NSInclude(&controllers.PF000008Controller{})),
		//beego.NSNamespace("/v1/pricing/fee/delete-strategy", beego.NSInclude(&controllers.PF000009Controller{})),
		//beego.NSNamespace("/v1/pricing/fee/calculate-fee", beego.NSInclude(&controllers.PF000010Controller{})),
		//beego.NSNamespace("/pr/PI000001", beego.NSInclude(&controllers.PI000001Controller{})),
		//beego.NSNamespace("/pr/PI000002", beego.NSInclude(&controllers.PI000002Controller{})),
		//beego.NSNamespace("/pr/PI000003", beego.NSInclude(&controllers.PI000003Controller{})),
		//beego.NSNamespace("/pr/PI000004", beego.NSInclude(&controllers.PI000004Controller{})),
		//beego.NSNamespace("/pr/PI000005", beego.NSInclude(&controllers.PI000005Controller{})),
		//beego.NSNamespace("/pr/PI000006", beego.NSInclude(&controllers.PI000006Controller{})),
		//beego.NSNamespace("/pr/PI000007", beego.NSInclude(&controllers.PI000007Controller{})),
		//beego.NSNamespace("/pr/PI000008", beego.NSInclude(&controllers.PI000008Controller{})),
		//beego.NSNamespace("/pr/PI000009", beego.NSInclude(&controllers.PI000009Controller{})),
		//beego.NSNamespace("/pr/PI000010", beego.NSInclude(&controllers.PI000010Controller{})),
		//beego.NSNamespace("/pr/PR000001", beego.NSInclude(&controllers.PR000001Controller{})),
		//beego.NSNamespace("/pr/PR000002", beego.NSInclude(&controllers.PR000002Controller{})),
		//beego.NSNamespace("/pr/PR000003", beego.NSInclude(&controllers.PR000003Controller{})),
		//beego.NSNamespace("/pr/PR000004", beego.NSInclude(&controllers.PR000004Controller{})),
		//beego.NSNamespace("/pr/PR000005", beego.NSInclude(&controllers.PR000005Controller{})),
		//beego.NSNamespace("/pr/PR000009", beego.NSInclude(&controllers.PR000009Controller{})),
		//beego.NSNamespace("/pr/PR000010", beego.NSInclude(&controllers.PR000010Controller{})),
		//beego.NSNamespace("/pr/PR000012", beego.NSInclude(&controllers.PR000012Controller{})),
		//beego.NSNamespace("/pr/PR000013", beego.NSInclude(&controllers.PR000013Controller{})),

		beego.NSNamespace("/v1/nfps/participant-list/inquiry", beego.NSInclude(&controllers.FP000001Controller{})),
		beego.NSNamespace("/v1/nfps/participant-detail/inquiry", beego.NSInclude(&controllers.FP000002Controller{})),
		beego.NSNamespace("/v1/nfps/participant/amendment", beego.NSInclude(&controllers.FP000003Controller{})),
		beego.NSNamespace("/v1/nfps/participant/suspension", beego.NSInclude(&controllers.FP000004Controller{})),
		beego.NSNamespace("/v1/nfps/participant/unsuspension", beego.NSInclude(&controllers.FP000005Controller{})),
		beego.NSNamespace("/v1/nfps/participant/creation", beego.NSInclude(&controllers.FP000006Controller{})),
		beego.NSNamespace("/v1/nfps/participant-relation-list/inquiry", beego.NSInclude(&controllers.FP000007Controller{})),
		beego.NSNamespace("/v1/nfps/participant-agency-relation-detail/inquiry", beego.NSInclude(&controllers.FP000008Controller{})),
		beego.NSNamespace("/v1/nfps/participant-agency-relation/creation", beego.NSInclude(&controllers.FP000009Controller{})),
		beego.NSNamespace("/v1/nfps/other-participant-detail/inquiry", beego.NSInclude(&controllers.FP000010Controller{})),
		beego.NSNamespace("/v1/nfps/rtgs-participant/inquiry", beego.NSInclude(&controllers.FP000011Controller{})),
		beego.NSNamespace("/v1/nfps/settlement-account/creation", beego.NSInclude(&controllers.FP000012Controller{})),
		beego.NSNamespace("/v1/nfps/random-code/creation", beego.NSInclude(&controllers.FP000013Controller{})),
		beego.NSNamespace("/v1/nfps/random-code/verification", beego.NSInclude(&controllers.FP000014Controller{})),
		beego.NSNamespace("/v1/nfps/settlement-account/inquiry", beego.NSInclude(&controllers.FP000015Controller{})),
		beego.NSNamespace("/v1/nfps/settlement-account-balance/adjustment", beego.NSInclude(&controllers.FP000016Controller{})),
		beego.NSNamespace("/v1/nfps/participant-agency-relation/deletion", beego.NSInclude(&controllers.FP000017Controller{})),
		beego.NSNamespace("/v1/nfps/settlement-account-detail/inquiry", beego.NSInclude(&controllers.FP000018Controller{})),
		beego.NSNamespace("/v1/nfps/participant-balance/debit-credit", beego.NSInclude(&controllers.FP000019Controller{})),
		beego.NSNamespace("/v1/nfps/alert-balance/setting", beego.NSInclude(&controllers.FP000020Controller{})),
		beego.NSNamespace("/v1/nfps/transaction/limit/add", beego.NSInclude(&controllers.FP000021Controller{})),
		beego.NSNamespace("/v1/nfps/transaction/limit/update", beego.NSInclude(&controllers.FP000024Controller{})),
		beego.NSNamespace("/v1/nfps/transaction/limit/query/list", beego.NSInclude(&controllers.FP000025Controller{})),
		beego.NSNamespace("/v1/nfps/transaction/limit/query/history", beego.NSInclude(&controllers.FP000026Controller{})),
		beego.NSNamespace("/v1/nfps/prod/fee/insert", beego.NSInclude(&controllers.FP000027Controller{})),
		beego.NSNamespace("/v1/nfps/balance/sweeping/update", beego.NSInclude(&controllers.FP000029Controller{})),
		beego.NSNamespace("/v1/nfps/report/file/query", beego.NSInclude(&controllers.FP000030Controller{})),
		beego.NSNamespace("/v1/nfps/banch/daily/and/monthly/report", beego.NSInclude(&controllers.FP000031Controller{})),
		beego.NSNamespace("/v1/nfps/download/file", beego.NSInclude(&controllers.FP000035Controller{})),
		beego.NSNamespace("/v1/nfps/balance/alter/noti", beego.NSInclude(&controllers.FP000036Controller{})),
		beego.NSNamespace("/v1/nfps/prod/fee/history/query/list", beego.NSInclude(&controllers.FP000039Controller{})),
		beego.NSNamespace("/v1/nfps/prod/fee/update", beego.NSInclude(&controllers.FP000040Controller{})),
		beego.NSNamespace("/v1/nfps/prod/fee/query", beego.NSInclude(&controllers.FP000041Controller{})),
		beego.NSNamespace("/v1/nfps/prod/fee/delete", beego.NSInclude(&controllers.FP000042Controller{})),
		beego.NSNamespace("/v1/nfps/report/log/query", beego.NSInclude(&controllers.FP000043Controller{})),
		beego.NSNamespace("/v1/nfps/alert/participant/list", beego.NSInclude(&controllers.FP000044Controller{})),
		beego.NSNamespace("/v1/nfps/alert/query", beego.NSInclude(&controllers.FP000045Controller{})),
		beego.NSNamespace("/v1/nfps/alert/participant/add", beego.NSInclude(&controllers.FP000046Controller{})),
		beego.NSNamespace("/v1/nfps/alert/participant/update", beego.NSInclude(&controllers.FP000047Controller{})),
		beego.NSNamespace("/v1/nfps/alert/participant/delete", beego.NSInclude(&controllers.FP000048Controller{})),
		beego.NSNamespace("/v1/nfps/alert/trigger", beego.NSInclude(&controllers.FP000049Controller{})),
		beego.NSNamespace("/v1/nfps/test/scenario/list", beego.NSInclude(&controllers.FP000050Controller{})),
		beego.NSNamespace("/v1/nfps/test/scenario/detail", beego.NSInclude(&controllers.FP000051Controller{})),
		beego.NSNamespace("/v1/nfps/test/step/detail", beego.NSInclude(&controllers.FP000052Controller{})),
		beego.NSNamespace("/v1/nfps/test/api/query", beego.NSInclude(&controllers.FP000053Controller{})),
		beego.NSNamespace("/v1/nfps/test/scenario/update", beego.NSInclude(&controllers.FP000055Controller{})),
		beego.NSNamespace("/v1/nfps/test/scenario/delete", beego.NSInclude(&controllers.FP000056Controller{})),
		beego.NSNamespace("/v1/nfps/test/scenario/create", beego.NSInclude(&controllers.FP0000572Controller{})),
		beego.NSNamespace("/v1/nfps/test/scenario/step/creation", beego.NSInclude(&controllers.FP000057Controller{})),
		beego.NSNamespace("/v1/nfps/test/scenario/step/update", beego.NSInclude(&controllers.FP000058Controller{})),
		beego.NSNamespace("/v1/nfps/test/scenario/step/delete", beego.NSInclude(&controllers.FP000059Controller{})),
		beego.NSNamespace("/v1/nfps/test/scenario/fields/values/creation", beego.NSInclude(&controllers.FP000060Controller{})),
		beego.NSNamespace("/v1/nfps/test/scenario/fields/values/update", beego.NSInclude(&controllers.FP000061Controller{})),
		beego.NSNamespace("/v1/nfps/test/scenario/fields/values/delete", beego.NSInclude(&controllers.FP000062Controller{})),
		beego.NSNamespace("/v1/nfps/test/api/schema", beego.NSInclude(&controllers.FP000063Controller{})),
		beego.NSNamespace("/v1/nfps/test/aig/serviceID", beego.NSInclude(&controllers.FP000064Controller{})),
		beego.NSNamespace("/v1/nfps/test/host/list", beego.NSInclude(&controllers.FP000065Controller{})),
		beego.NSNamespace("/v1/nfps/test/host/start", beego.NSInclude(&controllers.FP000066Controller{})),
		beego.NSNamespace("/v1/nfps/test/host/query", beego.NSInclude(&controllers.FP000067Controller{})),
		beego.NSNamespace("/v1/nfps/test/scenario/enum/query", beego.NSInclude(&controllers.FP000068Controller{})),
		beego.NSNamespace("/v1/nfps/partclearing/query", beego.NSInclude(&controllers.FP000070Controller{})),
		beego.NSNamespace("/v1/nfps/simulator/tcp/sending", beego.NSInclude(&controllers.FP000071Controller{})),
		beego.NSNamespace("/v1/nfps/simulator/tcp/receive", beego.NSInclude(&controllers.FP000072Controller{})),
		beego.NSNamespace("/v2/b2b/nfps/alert/inquiry", beego.NSInclude(&controllers.FP000077Controller{})),
		beego.NSNamespace("/v2/b2b/nfps/alert-balance/setting", beego.NSInclude(&controllers.FP000078Controller{})),
		beego.NSNamespace("/v1/nfps/test/scenario/simulato3/set/message", beego.NSInclude(&controllers.FP000097Controller{})),
		beego.NSNamespace("/v1/nfps/test/scenario/query/redis/message", beego.NSInclude(&controllers.FP000098Controller{})),
		beego.NSNamespace("/v1/nfps/test/scenario/simulato1/get/messages/info", beego.NSInclude(&controllers.FP000099Controller{})),
		beego.NSNamespace("/v1/nfps/test/scenario/simulato1/aog/wait", beego.NSInclude(&controllers.FP000100Controller{})),
		beego.NSNamespace("/v1/nfps/report/sample/query", beego.NSInclude(&controllers.FP000102Controller{})),
		beego.NSNamespace("/v1/nfps/netting/param/enquiry", beego.NSInclude(&controllers.FP000103Controller{})),
		beego.NSNamespace("/v1/nfps/netting/param/update", beego.NSInclude(&controllers.FP000104Controller{})),
		beego.NSNamespace("/v1/nfps/balance/adjustment/apply", beego.NSInclude(&controllers.FP000105Controller{})),
		beego.NSNamespace("/v1/nfps/balance/adjustment/update", beego.NSInclude(&controllers.FP000106Controller{})),
		beego.NSNamespace("/v1/nfps/balance/adjustment/delete", beego.NSInclude(&controllers.FP000107Controller{})),
		beego.NSNamespace("/v1/nfps/balance/adjustment/enquiry", beego.NSInclude(&controllers.FP000108Controller{})),
		beego.NSNamespace("/v1/nfps/balance/adjustment/confirm", beego.NSInclude(&controllers.FP000109Controller{})),
		beego.NSNamespace("/v1/nfps/balance/adjustment/checker/enquiry", beego.NSInclude(&controllers.FP000110Controller{})),
		beego.NSNamespace("/v1/nfps/netting/result/enquiry", beego.NSInclude(&controllers.FP000111Controller{})),
		beego.NSNamespace("/v3/nfps/payment/enum/list", beego.NSInclude(&controllers.FP000112Controller{})),
		beego.NSNamespace("/v1/nfps/payment/creat/list", beego.NSInclude(&controllers.FP000113Controller{})),
		beego.NSNamespace("/v1/nfps/balance/adjustment/history/enquiry", beego.NSInclude(&controllers.FP000114Controller{})),
		beego.NSNamespace("/v1/nfps/credit-transfer", beego.NSInclude(&controllers.FP010001Controller{})),
		beego.NSNamespace("/v1/nfps/payment-return", beego.NSInclude(&controllers.FP010004Controller{})),
		beego.NSNamespace("/v1/nfps/payment-summary", beego.NSInclude(&controllers.FP010005Controller{})),
		beego.NSNamespace("/v1/nfps/credit-transfer/detail", beego.NSInclude(&controllers.FP010006Controller{})),
		beego.NSNamespace("/v1/nfps/message-list/inquiry", beego.NSInclude(&controllers.FP010007Controller{})),
		beego.NSNamespace("/v1/nfps/message-detail/inquiry", beego.NSInclude(&controllers.FP010008Controller{})),
		beego.NSNamespace("/v1/nfps/message/redeliver", beego.NSInclude(&controllers.FP010009Controller{})),
		beego.NSNamespace("/v1/nfps/direct-debit/detail", beego.NSInclude(&controllers.FP010010Controller{})),
		beego.NSNamespace("/v1/nfps/payment-return/detail", beego.NSInclude(&controllers.FP010011Controller{})),
		beego.NSNamespace("/v1/nfps/direct-debit", beego.NSInclude(&controllers.FP010012Controller{})),
		beego.NSNamespace("/v1/nfps/balance/sweeping/enquiry", beego.NSInclude(&controllers.FP010013Controller{})),
		beego.NSNamespace("/v1/nfps/payment/detail/enquiry", beego.NSInclude(&controllers.FP010014Controller{})),
		beego.NSNamespace("/v1/nfps/iso20022/lookup", beego.NSInclude(&controllers.FP010015Controller{})),
		beego.NSNamespace("/v1/nfps/8583/lookup-and-account-verify", beego.NSInclude(&controllers.FP010016Controller{})),
		beego.NSNamespace("/v1/nfps/8583/credit-transfer", beego.NSInclude(&controllers.FP010017Controller{})),
		beego.NSNamespace("/v1/nfps/request/to/pay", beego.NSInclude(&controllers.FP010018Controller{})),
		beego.NSNamespace("/v1/nfps/8583/request-to-pay", beego.NSInclude(&controllers.FP010019Controller{})),
		beego.NSNamespace("/v1/nfps/8583/echo/check", beego.NSInclude(&controllers.FP010020Controller{})),
		beego.NSNamespace("/v1/nfps/20022/echo/check", beego.NSInclude(&controllers.FP010021Controller{})),
		beego.NSNamespace("/v1/nfps/20022/sign/on", beego.NSInclude(&controllers.FP010022Controller{})),
		beego.NSNamespace("/v1/nfps/8583/sign/on", beego.NSInclude(&controllers.FP010023Controller{})),
		beego.NSNamespace("/v1/nfps/20022/sign/off", beego.NSInclude(&controllers.FP010024Controller{})),
		beego.NSNamespace("/v1/nfps/8583/sign/off", beego.NSInclude(&controllers.FP010025Controller{})),
		beego.NSNamespace("/v1/nfps/8583/credit/notification", beego.NSInclude(&controllers.FP010026Controller{})),
		beego.NSNamespace("/v1/nfps/8583/bill-payment", beego.NSInclude(&controllers.FP010027Controller{})),
		beego.NSNamespace("/v3/nfps/iso20022/lookup", beego.NSInclude(&controllers.FP013015Controller{})),
		beego.NSNamespace("/v3/nfps/request/to/pay", beego.NSInclude(&controllers.FP013018Controller{})),
		beego.NSNamespace("/v3/nfps/adjust/mockinfo", beego.NSInclude(&controllers.FP013028Controller{})),
		beego.NSNamespace("/v1/nfps/addressing/register", beego.NSInclude(&controllers.FP020001Controller{})),
		beego.NSNamespace("/v1/nfps/addressing/cancel", beego.NSInclude(&controllers.FP020002Controller{})),
		beego.NSNamespace("/v1/nfps/addressing/amendment", beego.NSInclude(&controllers.FP020003Controller{})),
		beego.NSNamespace("/v1/nfps/addressing/inquiry", beego.NSInclude(&controllers.FP020004Controller{})),
		beego.NSNamespace("/v1/nfps/import/execl/add", beego.NSInclude(&controllers.FP020007Controller{})),
		beego.NSNamespace("/v1/nfps/import/execl/list", beego.NSInclude(&controllers.FP020009Controller{})),
		beego.NSNamespace("/v1/nfps/import/execl/download", beego.NSInclude(&controllers.FP020010Controller{})),
	)
	beego.AddNamespace(ns)
}
