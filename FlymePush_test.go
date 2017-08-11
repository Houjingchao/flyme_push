package flyme_push

import (
	"testing"
	"fmt"
	"github.com/Houjingchao/flyme_push/model"
	"github.com/Houjingchao/flyme_push/consts"
)

func TestSendNotificationMessageByPushId(t *testing.T) {
	flymePush := FlymePush{
		AppId:  "appid",
		AppKey: "appkey",
	}
	json := model.NewdNotificationMessage().
		NoticeBarType(2). //android 原生
		NoticeTitle("komiko").
		NoticeContent("测试内容").JSON()
	message := flymePush.SendNotificationMessageByPushId(consts.PushID, json)
	fmt.Println(message)
}

//别名推送接口（通知栏消息）
func TestPushNotificationMessageByAlias(t *testing.T) {
	flymePush := FlymePush{
		AppId:  "appid",
		AppKey: "appkey",
	}
	json := model.NewdNotificationMessage().
		NoticeBarType(2).
		NoticeTitle("komiko-别名推送").
		NoticeContent("测试内容-别名推送").JSON()
	message := flymePush.SendNotificationMessageByAlias("1827541", json)
	fmt.Println(message)
}

//SendAllMessage
func TestSendAllMessage(t *testing.T) {
	flymePush := FlymePush{
		AppId:  "appid",
		AppKey: "appkey",
	}
	json := model.NewdNotificationMessage().
		NoticeBarType(2).
		NoticeTitle("komiko-Sendall").
		NoticeContent("测试内容").JSON()
	message := flymePush.SendAllMessage("0", json) //pushTye=0通知栏消息
	fmt.Println(message)
}

//根据标签推送 SendMessageByTipic  废弃
func TestSendMessageByTipic(t *testing.T) {
	flymePush := FlymePush{
		AppId:  "appid",
		AppKey: "appkey",
	}
	json := model.NewdNotificationMessage().
		NoticeBarType(2).
		NoticeTitle("komiko").
		NoticeContent("测试内容").JSON()
	message := flymePush.SendMessageByTipic("0", "tagNames", "scope", json) //pushTye=0通知栏消息
	fmt.Println(message)
}

//统计结果
//PushStatistics
func TestPushStatistics(t *testing.T) {
	flymePush := FlymePush{
		AppId:  "appid",
		AppKey: "appkey",
	}
	message := flymePush.SendStatistics("116675", )
	fmt.Println(message)
}
