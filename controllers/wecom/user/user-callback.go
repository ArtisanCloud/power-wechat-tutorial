package user

import (
	"github.com/ArtisanCloud/PowerLibs/v3/fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/http/helper"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/server/handlers/models"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"power-wechat-tutorial/services"
)

// 回调配置
// https://work.weixin.qq.com/api/doc/90000/90135/90930
func CallbackVerify(c *gin.Context) {
	rs, err := services.WeComContactApp.Server.VerifyURL(c.Request)
	if err != nil {
		panic(err)
	}

	text, _ := ioutil.ReadAll(rs.Body)
	c.String(http.StatusOK, string(text))

	// 选择2
	//err = helper.HttpResponseSend(rs, c.Writer)

}

// 回调配置
// https://work.weixin.qq.com/api/doc/90000/90135/90930
func CallbackNotify(c *gin.Context) {

	rs, err := services.WeComContactApp.Server.Notify(c.Request, func(event contract.EventInterface) interface{} {
		fmt.Dump("event", event)

		// 这里需要获取到事件类型，然后把对应的结构体传递进去进一步解析
		// 所有包含的结构体请参考： https://github.com/ArtisanCloud/PowerWeChat/tree/master/src/work/server/handlers/models
		if event.GetEvent() == models.CALLBACK_EVENT_CHANGE_CONTACT && event.GetChangeType() == models.CALLBACK_EVENT_CHANGE_TYPE_CREATE_PARTY {
			//msg := models.EventPartyCreate{}
			msg := models.EventKFMsgOrEvent{}
			err := event.ReadMessage(&msg)
			if err != nil {
				println(err.Error())
				return "error"
			}
			fmt.Dump(msg)
		}

		// 假设员工给应用发送消息，这里可以直接回复消息文本，
		// return  "I'm recv..."

		// 这里回复success告诉微信我收到了，后续需要回复用户信息可以主动调发消息接口
		return kernel.SUCCESS_EMPTY_RESPONSE
	})
	if err != nil {
		panic(err)
	}

	// 选择1： 直接把gin context writer传入，会自动回复。
	err = helper.HttpResponseSend(rs, c.Writer)
	if err != nil {
		panic(err)
	}

	// 选择2： 或者是把内容读取出来，回复
	//text, _ := ioutil.ReadAll(rs.Body)
	//c.String(http.StatusOK, string(text))

}
