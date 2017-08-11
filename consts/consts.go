package consts

//需要自己添加
const (
	AppID  = "AppID"
	AppKey = "AppKey"
	Host   = "Host"
	PushID = "PushID"
)
const (
	//透传消息by pushId
	PushThroughMessageByPushId = Host + "//garcia/api/server/push/unvarnished/pushByPushId"
	//通知栏消息 by pushId
	PushNotificationMessageByPushId = Host + "/garcia/api/server/push/varnished/pushByPushId"

	PushThroughMessageByAlias      = Host + "/garcia/api/server/push/unvarnished/pushByAlias"
	PushNotificationMessageByAlias = Host + "/garcia/api/server/push/varnished/pushByAlias"

	PushAllMessage     = Host + "/garcia/api/server/push/pushTask/pushToApp"
	PushMessageByTipic = Host + "/garcia/api/server/push/pushTask/pushToTag"

	//推送统计
	PushStatistics = Host + "/garcia/api/server/push/statistics/getTaskStatistics"
)
