//Version: v1
package models

import "github.com/shopspring/decimal"

type BP000001Request struct {
	CustomerId string `json:"customerId" description:"Customer ID" validate:"required,max=20"`
	OrderId    string `json:"orderId" description:"Order ID"  validate:"required,max=20"`
}

type BP000001Response struct {
	OrderInfo         OrderInfo           `json:"orderInfo" description:"Order Information"`
	RepaymentPlanList []RepaymentPlanInfo `json:"repaymentPlanList" description:"Repayment Plan Information"`
	PaymentMethodList []PaymentMethodInfo `json:"paymentMethodList" description:"Payment Method Information"`
}

type OrderInfo struct {
	OrderId          string          `json:"orderId" description:"Order ID" validate:"max=20" `
	ProductStoreId   string          `json:"productStoreId" description:"Product Store ID" validate:"max=20"`
	ProductStoreName string          `json:"productStoreName" description:"Product Store Name" validate:"max=100"`
	TotalOrderAmount decimal.Decimal `json:"totalOrderAmount" description:"Total Order Amount"`
	ProductList      []ProductInfo   `json:"productList" description:"Product List"`
}

type ProductInfo struct {
	ProductId       string          `json:"productId" description:"Product ID" validate:"max=20"`
	ProductName     string          `json:"productName" description:"Product Name" validate:"max=100"`
	ProductCategory string          `json:"productCategory" description:"Product Category" validate:"max=100"`
	ProductPrice    decimal.Decimal `json:"productPrice" description:"Product Price"`
	ProductQuantity decimal.Decimal `json:"productQuantity" description:"Product Quantity"`
}

type RepaymentPlanInfo struct {
	PeriodNum                  int             `json:"periodNum" description:"Period Number"`
	InterestCalculateStartDate string          `json:"interestCalculateStartDate" description:"Interest Calcuation Start Date" validate:"datetime=2006-01-02"`
	InterestCalculateEndDate   string          `json:"interestCalculateEndDate" description:"Interest Calculation End Date" validate:"datetime=2006-01-02"`
	PlanRepayDate              string          `json:"planRepayDate" description:"Plan Repayment Date" validate:"datetime=2006-01-02"`
	PlanRepayTotalAmount       decimal.Decimal `json:"planRepayTotalAmount" description:"Total Plan Repayment Amount"`
	PlanRepayPrinciple         decimal.Decimal `json:"planRepayPrinciple" description:"Plan Repayment Principle"`
	PlanRepayInterest          decimal.Decimal `json:"planRepayInterest" description:"Plan Repayment Interest"`
	MaintainPrinciple          decimal.Decimal `json:"maintainPrinciple" description:"Maintain Principle"`
	DaysOfInterestCalculate    int             `json:"daysOfInterestCalculate" description:"Days Of Interest Calculation"`
}

type PaymentMethodInfo struct {
	BindAcctNo      string `json:"bindAcctNo" description:"Bind Account Number" validate:"max=32"`
	BindAcctType    string `json:"bindAcctType" description:"Bind Account Type" validate:"max=2"`
	BindAcctOweBank string `json:"bindAcctOweBank" description:"Bind Account Owe Bank" validate:"max=3"`
}

func (*BP000001Request) GetServiceKey() string {
	return "BP000001"
}
