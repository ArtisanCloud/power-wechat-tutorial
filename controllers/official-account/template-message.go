package official_account

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/templateMessage/request"
	"github.com/gin-gonic/gin"
	"power-wechat-tutorial/services"
)

// TemplateMessageSetIndustry 修改账号所属行业
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html#0
func TemplateMessageSetIndustry(c *gin.Context) {
	industryId1 := c.DefaultQuery("industryId1", "1")
	industryId2 := c.DefaultQuery("industryId2", "2")

	rs, err := services.OfficialAccountApp.TemplateMessage.SetIndustry(c.Request.Context(), industryId1, industryId2, nil)
	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// TemplateMessageGetIndustry 获取支持的行业列表
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html#1
func TemplateMessageGetIndustry(c *gin.Context) {
	rs, err := services.OfficialAccountApp.TemplateMessage.GetIndustry(c.Request.Context())
	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// TemplateMessageAddTemplate 添加模板
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html#2
func TemplateMessageAddTemplate(c *gin.Context) {
	shortID := c.DefaultQuery("shortID", "TM00015")
	rs, err := services.OfficialAccountApp.TemplateMessage.AddTemplate(c.Request.Context(), shortID, []string{})
	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// GetPrivateTemplates 获取所有模板列表
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html#3
func GetPrivateTemplates(c *gin.Context) {
	rs, err := services.OfficialAccountApp.TemplateMessage.GetPrivateTemplates(c.Request.Context())
	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// DelPrivateTemplate 删除模板
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html#4
func DelPrivateTemplate(c *gin.Context) {
	templateID := c.DefaultQuery("templateID", "123")
	rs, err := services.OfficialAccountApp.TemplateMessage.DeletePrivateTemplate(c.Request.Context(), templateID)
	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// TemplateMessageSend 发送模板消息
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html#5
func TemplateMessageSend(c *gin.Context) {
	toUser := c.DefaultQuery("toUser", "o1PjF1K4W5fjQYqC0Fm9qX5YqXqk")
	templateID := c.DefaultQuery("templateID", "MTvUCMmZfl-Dv66C5fVWdf4zPJkYSaRbnrtk6DXIfTQ")
	rs, err := services.OfficialAccountApp.TemplateMessage.Send(c.Request.Context(), &request.RequestTemlateMessage{
		ToUser:     toUser,
		TemplateID: templateID,
		URL:        "https://www.artisan-cloud.com/",
		Data: &power.HashMap{
			"first": &power.HashMap{
				"value": "恭喜你购买成功！",
				"color": "#173177",
			},
			"DateTime": &power.HashMap{
				"value": "2022-3-5 16:22",
				"color": "#173177",
			},
			"PayAmount": &power.HashMap{
				"value": "59.8元",
				"color": "#173177",
			},
			"Location": &power.HashMap{
				"value": "上海市长宁区",
				"color": "#173177",
			},
			"remark": &power.HashMap{
				"value": "欢迎再次购买！",
				"color": "#173177",
			},
		},
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 发送一次性订阅消息
// https://developers.weixin.qq.com/doc/offiaccount/Subscription_Messages/api.html#send
func SendSubscribe(c *gin.Context) {
	toUser := c.DefaultQuery("toUser", "o1PjF1K4W5fjQYqC0Fm9qX5YqXqk")
	templateID := c.DefaultQuery("templateID", "MTvUCMmZfl-Dv66C5fVWdf4zPJkYSaRbnrtk6DXIfTQ")
	rs, err := services.OfficialAccountApp.TemplateMessage.SendSubscription(
		c.Request.Context(),
		&request.RequestTemlateMessageSubscribe{
			ToUser:     toUser,
			TemplateID: templateID,
			URL:        "https://www.artisan-cloud.com/",
			Data: &power.HashMap{
				"first": &power.HashMap{
					"value": "恭喜你购买成功！",
					"color": "#173177",
				},
				"keyword1": &power.HashMap{
					"value": "巧克力",
					"color": "#173177",
				},
				"keyword2": &power.HashMap{
					"value": "39.8元",
					"color": "#173177",
				},
				"keyword3": &power.HashMap{
					"value": "2014年9月22日",
					"color": "#173177",
				},
				"remark": &power.HashMap{
					"value": "欢迎再次购买！",
					"color": "#173177",
				},
			},
		})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}
