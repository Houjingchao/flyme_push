package flyme_push

import (
	"github.com/cocotyty/httpclient"
	"encoding/json"
	"sort"
	"crypto/md5"
	"encoding/hex"
	"github.com/Houjingchao/flyme_push/consts"
)

type FlymePush struct {
	AppId  string
	AppKey string
}

/**
 * 通过PushId推送透传消息
 */
func (f FlymePush) SendThroughByPushIds(pushIds, messageJson string) error {
	sendByPushIds := map[string]string{
		"appId":       f.AppId,
		"pushIds":     pushIds,
		"messageJson": messageJson,
	}
	sign := GenerateSign(sendByPushIds, f.AppKey)
	_, err := httpclient.
	Post(consts.PushThroughMessageByPushId).
		Head("charset", "UTF-8").
		Param("appId", f.AppId).
		Param("pushIds", pushIds). //多个逗号隔开
		Param("sign", sign).
		Param("messageJson", messageJson).
		Send().String()
	if err != nil {
		return err
	}
	return nil
}

//pushId推送接口（通知栏消息）
func (f FlymePush) SendNotificationMessageByPushId(pushIds string, messageJson string) error {
	pushNotificationMessageMap := map[string]string{
		"appId":       f.AppId,
		"pushIds":     pushIds,
		"messageJson": messageJson,
	}
	sign := GenerateSign(pushNotificationMessageMap, f.AppKey)
	_, err := httpclient.
	Post(consts.PushNotificationMessageByPushId).
		Head("charset", "UTF-8").
		Param("appId", f.AppId).
		Param("pushIds", pushIds). //多个逗号隔开
		Param("sign", sign).
		Param("messageJson", messageJson).
		Send().String()
	if err != nil {
		return err
	}

	return nil
}

//别名推送接口（通知栏消息）
func (f FlymePush) SendNotificationMessageByAlias(alias string, messageJson string) error {
	maps := map[string]string{
		"appId":       f.AppId,
		"alias":       alias,
		"messageJson": messageJson,
	}

	sign := GenerateSign(maps, f.AppKey)
	_, err := httpclient.
	Post(consts.PushNotificationMessageByAlias).
		Head("charset", "UTF-8").
		Param("appId", f.AppId).
		Param("alias", alias). //推送别名，一批最多不能超过1000个多个英文逗号分割（必填）
		Param("sign", sign).
		Param("messageJson", messageJson).
		Send().String()
	if err != nil {
		return err
	}
	return nil
}

/*******************************************标签推送****************************************/
//全部推送
func (f FlymePush) SendAllMessage(pushType string, messageJson string) error {
	maps := map[string]string{
		"appId":       f.AppId,
		"pushType":    pushType,
		"messageJson": messageJson,
	}

	sign := GenerateSign(maps, f.AppKey)
	_, err := httpclient.
	Post(consts.PushAllMessage).
		Head("charset", "UTF-8").
		Param("appId", f.AppId).
		Param("pushType", pushType).
		Param("sign", sign).
		Param("messageJson", messageJson).
		Send().String()
	if err != nil {
		return err
	}
	return nil
}

// 标签推送
func (f FlymePush) SendMessageByTipic(pushType string, tagNames string, scope string, messageJson string) error {
	maps := map[string]string{
		"appId":       f.AppId,
		"pushType":    pushType,
		"tagNames":    tagNames,
		"scope":       scope,
		"messageJson": messageJson,
	}

	sign := GenerateSign(maps, f.AppKey)
	_, err := httpclient.
	Post(consts.PushMessageByTipic).
		Head("charset", "UTF-8").
		Param("appId", f.AppId).
		Param("pushType", pushType).
		Param("tagNames", tagNames).
		Param("scope", scope). //标签集合（必填）0 并集1 交集
		Param("sign", sign).
		Param("messageJson", messageJson).
		Send().String()
	if err != nil {
		return err
	}
	return nil
}

/********************************推送统计*******************************/
//PushStatistics get 请求
func (f FlymePush) SendStatistics( taskId string) error {
	maps := map[string]string{
		"appId":   f.AppId,
		"taskId": taskId,
	}
	sign := GenerateSign(maps, f.AppKey)
	_, err := httpclient.
	Post(consts.PushStatistics).
		Head("charset", "UTF-8").
		Param("appId",  f.AppId).
		Param("taskId", taskId).
		Param("sign", sign).
		Send().String()
	if err != nil {
		return err
	}
	return nil
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
