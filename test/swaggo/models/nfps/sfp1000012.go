/**
 * @Author: fangwen
 * @Description:
 * @File: sfp1000007
 * @Version: 1.0.0
 * @Date: 2022/3/18 1:45 AM
 */

package models

import "time"

type FP100012I struct {
	AlertID            string               `json:"alertID"`
	StrategyName       string               `json:"strategyName"`
	AlertType          string               `json:"alertType"`
	AlertEntity        string               `json:"alertEntity"`
	AlertDetails       string               `json:"alertDetails"`
	TimeStart          time.Time            `json:"timeStart"`
	Actions            string               `json:"actions"`
	WorkSpace          string               `json:"workSpace"`
	Env                string               `json:"env"`
	Su                 string               `json:"su"`
	AlertName          string               `json:"alertName"`
	AlertLevel         string               `json:"alertLevel"`
	AlertLabels        string               `json:"alertLabels"`
	Expression         string               `json:"expression"`
	Lable              string               `json:"lable"`
	Type               string               `json:"type"`
	ParticipantConnect []ParticipantConnect `json:"participant"`
}

type SedEmailBody struct {
	Bodies []SedInfo
}

type SedInfo struct {
	AlertID          string    `json:"AlertID"`
	AlertType        string    `json:"AlertType"`
	AlertEntity      string    `json:"AlertEntity"`
	AlertDetails     string    `json:"AlertDetails"`
	TimeStart        time.Time `json:"TimeStart"`
	WorkSpace        string    `json:"WorkSpace"`
	Env              string    `json:"Env"`
	Su               string    `json:"Su"`
	PartClearingCode string    `json:"partClearingCode"`
	AlarmStrategy    string    `json:"alarmStrategy"`
	AlertName        string    `json:"AlertName"`
	AlertLevel       string    `json:"AlertLevel"`
	AlertLabels      string    `json:"AlertLabels"`
}

type ParticipantConnect struct {
	ParticipantClearingCode string `json:"participantClearingCode"`
	ContactName             string `json:"contactName"`
	ContactMethod           string `json:"contactMethod" xorm:"'contact_method'"`
	ContactInformation      string `json:"contactInformation" xorm:"'contact_information'"`
	Priority                string `json:"priority" xorm:"'priority'"`
	Necessary               string `json:"necessary" xorm:"'necessary'"`
	Remark                  string `json:"remark" xorm:"'remark'"`
}

type FP100012O struct {
	AssociatedStates bool `json:"associatedStates"`
}

type McEmailRequest struct {
	AttachList []string `json:"attachList"`
	Body       string   `json:"body"`
	CcList     []string `json:"ccList"`
	MessageId  string   `json:"messageId"`
	SendTime   string   `json:"sendTime"`
	TemplateId string   `json:"templateId"`
	Title      string   `json:"title"`
	ToList     []string `json:"toList"`
}

func (h *FP100012I) GetMcTopic() string {
	return "ACCEPTMSG_SYNC_ITMX"
}

func (h *FP100012I) GetMcMattermostTopic() string {
	return "ACCEPTMSG_SYNC_MM"
}

func (h *FP100012I) GetLineTopic() string {
	return "ACCEPTMSG_SYNC_Line"
}

type McEmailResponse struct {
	ErrorCode string     `json:"errorCode"`
	ErrorMsg  string     `json:"errorMsg"`
	Response  MCResponse `json:"response"`
}

type MattermostRequest struct {
	MessageId         string `json:"messageId"`
	SendTime          string `json:"sendTime"`
	MattermostChannel string `json:"mattermostChannel"`
	TemplateId        string `json:"templateId"`
	Text              string `json:"text"`
}

type LineRequest struct {
	MessageId        string `json:"messageId"`
	SendTime         string `json:"sendTime"`
	LineGroupChannel string `json:"lineGroupChannel"`
	LineUserChannel  string `json:"lineUserChannel"`
	TemplateId       string `json:"templateId"`
	Text             string `json:"text"`
}

type Test struct {
	ADFASDFASD string `json:"ADFASDFASD"`
}

type MCResponse struct {
	Data         bool   `json:"data"`
	DeliveryTime string `json:"deliveryTime"`
	ResponseTime string `json:"responseTime"`
}

func (*FP100012I) GetServiceKey() string {
	return "FP100012"
}
