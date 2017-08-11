package flyme_push

import (
	"github.com/cocotyty/httpclient"
	"encoding/json"
	"fmt"
	"sort"
	"crypto/md5"
	"encoding/hex"
	"github.com/Houjingchao/flyme_push/consts"
)

/**
 * 通过PushId推送透传消息
 */
func SendThroughByPushIds(appId, pushIds, messageJson, appKey string) string {
	sendByPushIds := map[string]string{
		"appId":       appId,
		"pushIds":     pushIds,
		"messageJson": messageJson,
	}
	sign := GenerateSign(sendByPushIds, appKey)
	req, err := httpclient.
	Post(consts.PushThroughMessageByPushId).
		Head("charset", "UTF-8").
		Param("appId", appId).
		Param("pushIds", pushIds). //多个逗号隔开
		Param("sign", sign).
		Param("messageJson", messageJson).
		Send().String()
	if err != nil {
		panic("SendThroughByPushIds 报错了～～")
	}
	fmt.Println(req)
	return req
}

//pushId推送接口（通知栏消息）
func SendNotificationMessageByPushId(appId string, pushIds string, messageJson string, appKey string) string {
	pushNotificationMessageMap := map[string]string{
		"appId":       appId,
		"pushIds":     pushIds,
		"messageJson": messageJson,
	}
	sign := GenerateSign(pushNotificationMessageMap, appKey)
	req, err := httpclient.
	Post(consts.PushNotificationMessageByPushId).
		Head("charset", "UTF-8").
		Param("appId", appId).
		Param("pushIds", pushIds). //多个逗号隔开
		Param("sign", sign).
		Param("messageJson", messageJson).
		Send().String()
	if err != nil {
		panic("SendNotificationMessageByPushId 报错了～～")
	}
	fmt.Println(req)
	return req
}

//别名推送接口（通知栏消息）
func SendNotificationMessageByAlias(appId string, alias string, messageJson string, appKey string) string {
	maps := map[string]string{
		"appId":       appId,
		"alias":       alias,
		"messageJson": messageJson,
	}

	sign := GenerateSign(maps, appKey)
	req, err := httpclient.
	Post(consts.PushNotificationMessageByAlias).
		Head("charset", "UTF-8").
		Param("appId", appId).
		Param("alias", alias). //推送别名，一批最多不能超过1000个多个英文逗号分割（必填）
		Param("sign", sign).
		Param("messageJson", messageJson).
		Send().String()
	if err != nil {
		panic("PushNotificationMessageByAlias 报错了～～")
	}
	fmt.Println(req)
	return req
}
/*******************************************标签推送****************************************/
//全部推送
func SendAllMessage(appId string, pushType string, messageJson string, appKey string) string {
	maps := map[string]string{
		"appId":       appId,
		"pushType":    pushType,
		"messageJson": messageJson,
	}

	sign := GenerateSign(maps, appKey)
	req, err := httpclient.
	Post(consts.PushAllMessage).
		Head("charset", "UTF-8").
		Param("appId", appId).
		Param("pushType", pushType).
		Param("sign", sign).
		Param("messageJson", messageJson).
		Send().String()
	if err != nil {
		panic("SendAllMessage 报错了～～")
	}
	fmt.Println(req)
	return req
}

// 标签推送
func SendMessageByTipic(appId string, pushType string, tagNames string, scope string, messageJson string, appKey string) string {
	maps := map[string]string{
		"appId":       appId,
		"pushType":    pushType,
		"tagNames":    tagNames,
		"scope":       scope,
		"messageJson": messageJson,
	}

	sign := GenerateSign(maps, appKey)
	req, err := httpclient.
	Post(consts.PushMessageByTipic).
		Head("charset", "UTF-8").
		Param("appId", appId).
		Param("pushType", pushType).
		Param("tagNames", tagNames).
		Param("scope", scope).//标签集合（必填）0 并集1 交集
		Param("sign", sign).
		Param("messageJson", messageJson).
		Send().String()
	if err != nil {
		panic("SendAllMessage 报错了～～")
	}
	fmt.Println(req)
	return req
}

/********************************推送统计*******************************/
//PushStatistics get 请求
func SendStatistics(appId string, taskId string , appKey string) string {
	maps := map[string]string{
		"appId":       appId,
		"taskId":    taskId,
	}
	sign := GenerateSign(maps, appKey)
	req, err := httpclient.
	Post(consts.PushStatistics).
		Head("charset", "UTF-8").
		Param("appId", appId).
		Param("taskId", taskId).
		Param("sign", sign).
		Send().String()
	if err != nil {
		panic("PushStatistics 报错了～～")
	}
	fmt.Println(req)
	return req
}
func GenerateSign(params map[string]string, appKey string) string {
	var signStr string
	if params != nil {
		keys := make([]string, len(params))
		i := 0
		for key, _ := range params {
			keys[i] = key
			i++
		}
		sort.Strings(keys)
		for _, k := range keys {
			signStr += k + "=" + params[k]
		}
		signStr += appKey
	}
	return PushParamMD5(signStr)
}

func PushParamMD5(encodeStr string) string {
	hasher := md5.New()
	hasher.Write([]byte(encodeStr))
	return hex.EncodeToString(hasher.Sum(nil))
}

/*json方法*/
func JSON(i interface{}) string {
	outi, err := json.Marshal(i)
	if err != nil {
		panic("json出错了～～")
	}
	return string(outi)
}
