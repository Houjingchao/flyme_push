package flyme_push

import (
	"testing"
	"fmt"
	"github.com/Houjingchao/flyme_push/model"
	"github.com/Houjingchao/flyme_push/consts"
)

func TestSendNotificationMessageByPushId(t *testing.T) {
	json := model.NewdNotificationMessage().
		NoticeBarType(2). //android 原生
		NoticeTitle("komiko").
		NoticeContent("测试内容").JSON()
	message := SendNotificationMessageByPushId(consts.AppID, consts.PushID, json, consts.AppKey)
	fmt.Println(message)
}

//别名推送接口（通知栏消息）
func TestPushNotificationMessageByAlias(t *testing.T) {
	json := model.NewdNotificationMessage().
		NoticeBarType(2).
		NoticeTitle("komiko-别名推送").
		NoticeContent("测试内容-别名推送").JSON()
	message := SendNotificationMessageByAlias(consts.AppID, "1827541", json, consts.AppKey)
	fmt.Println(message)
}

//SendAllMessage
func TestSendAllMessage(t *testing.T) {
	json := model.NewdNotificationMessage().
		NoticeBarType(2).
		NoticeTitle("komiko-Sendall").
		NoticeContent("测试内容").JSON()
	message := SendAllMessage(consts.AppID, "0", json, consts.AppKey) //pushTye=0通知栏消息
	fmt.Println(message)
}

//根据标签推送 SendMessageByTipic  废弃
func TestSendMessageByTipic(t *testing.T) {
	json := model.NewdNotificationMessage().
		NoticeBarType(2).
		NoticeTitle("komiko").
		NoticeContent("测试内容").JSON()
	message := SendMessageByTipic(consts.AppID, "0", "tagNames", "scope", json, consts.AppKey) //pushTye=0通知栏消息
	fmt.Println(message)
}

//统计结果
//PushStatistics
func TestPushStatistics(t *testing.T) {
	message := SendStatistics(consts.AppID, "116675", consts.AppKey)
	fmt.Println(message)
}
