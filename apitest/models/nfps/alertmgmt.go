package models



type PushAlert struct {
	/*
	   the tablename for alert
	   if the alert is about rdb cluster,then it is RDBCluster
	   if the alert is about rdb node,then it is RDBNode
	*/
	CIName string `json:"ciName"`

	/*
	   the field of table for alert
	   if the alert is about rdb cluster,then it is RDBCluster's Cluster_ID
	   if the alert is about rdb node,then it is RDBNode's IP
	*/
	CIFieldName string `json:"ciFieldName"`

	/*
	   the field value of table for alert。
	   if the alert is about rdb cluster,then it is RDBCluster's Cluster_ID="clusterXXX",target="clusterXXX"。
	   if the alert is about rdb node,then it is RDBNode's IP="10.10.10.10",target="10.10.10.10"。
	*/
	AlertTarget string `json:"target"`
	Org         string                 `json:"org"`
	Az string `json:"az"`
	WorkSpace string `json:"workspace"`
	Env string `json:"env"`
	Su string `json:"su"`
	/*
	   First letter is upper。its values are as follow，
	   if your alert level type is not in，please contact nanili
	   {
	      Host
	      Rdb
	      Kubernetes
	      Cache
	      Rados
	      Service
	      Mesh
	   }
	*/
	AlertType   string                 `json:"type"`
	/*
	  A short, one-word description of the object of the alarm。
	  such as cpuUsage，memUsage，tps，qps
	*/
	AlertName   string                 `json:"name"`
	/*
	     First letter is upper。its values are as follow，
	   if your alert level value is not in，please contact nanili
	     {
	        Warning
	        Minor
	        Major
	        Critical
	        Fatal
	     }
	*/
	AlertLevel  string                 `json:"level"`
	/*
	   First letter is upper。its values are as follow，
	   if your alert receive type value is not in，please contact nanili
	   {
	      Email
	      Mu
	   }
	*/
	AlertReceiveType string `json:"receivetype"`
	/*
	   First letter is upper。its values are as follow，
	   if your type value is not in，please contact nanili
	   {
	      Type=Email：receiver mail，more than one，split by “;”
	      Type=Mu:show in mu page
	   }
	*/
	AlertReceives string `json:"receives"`
	/*
	   other info of alert
	*/
	AlertLabels map[string]interface{} `json:"labels"`
	/*
	  alert detail info
	*/
	AlertMsg    string                 `json:"msg"`
}