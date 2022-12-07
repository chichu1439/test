package util

import (
	"apitest/models"
	"encoding/json"
)

//func SedAlert(topicId string,body,target string)  {
//	alertBody := buildAlertInfo(body,target)
//	meshRequest := mesh.NewMeshRequest(alertBody)
//	meshRequest.WithOptions(
//		mesh.WithTopicTypeAlert(),
//		mesh.WithORG("*"),
//		mesh.WithWorkspace("*"),
//		mesh.WithEnvironment("*"),
//		//mesh.WithSU("*"),
//		//mesh.WithNodeID("*"),
//		//mesh.WithInstanceID("*"),
//		mesh.WithEventID(topicId),
//		mesh.WithMaxRetryTimes(1),
//		mesh.WithTimeout(3*time.Second),
//		mesh.WithCodec(codec.BuildCustomCodec(&text.Encoder{}, &text.Decoder{})),
//	)
//	DefaultClient := mesh.NewMeshClient()
//	err := DefaultClient.AsyncCall(context.Background(), meshRequest)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//}

func BuildAlertInfo(body string, target string) []byte {
	alertStruct := models.PushAlert{
		AlertType:   "Service",
		AlertTarget: target,
		AlertName:   "log",
		AlertLevel:  "Warning",
		AlertLabels: map[string]interface{}{},
		AlertMsg:    body,
	}
	request := []models.PushAlert{alertStruct}
	bdata, _ := json.Marshal(request)
	return bdata
}