package account_service

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/accountService/customer/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

// 读取消息
// https://work.weixin.qq.com/api/doc/90000/90135/94670
func APIAccountServiceCustomerBatchGet(c *gin.Context) {
	externalUserIDList := []string{c.DefaultQuery("externalUserIDList", "matrix-x")}

	res, err := services.WeComApp.AccountServiceCustomer.BatchGet(c.Request.Context(), externalUserIDList, 1)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 获取配置的专员与客户群
// https://work.weixin.qq.com/api/doc/90000/90135/94674
func APIAccountServiceCustomerGetUpgradeServiceConfig(c *gin.Context) {
	res, err := services.WeComApp.AccountServiceCustomer.GetUpgradeServiceConfig(c.Request.Context())

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 为客户升级为专员或客户群服务
// https://work.weixin.qq.com/api/doc/90000/90135/94674
func APIAccountServiceCustomerUpgradeService(c *gin.Context) {
	options := &request.RequestUpgradeService{
		OpenKFID:       c.DefaultQuery("openKFID", "kfxxxxxxxxxxxxxx"),
		ExternalUserID: c.DefaultQuery("externalUserID", "kfxxxxxxxxxxxxxx"),
		Type:           2, // 表示是升级到专员服务还是客户群服务。1:专员服务。2:客户群服务
		Member: &request.RequestUpgradeServiceMember{
			UserID:  c.DefaultQuery("member", "matrix-x"),
			Wording: "你好，我是你的专属服务专员zhangsan",
		}, // 推荐的服务专员，type等于1时有效
		GroupChat: &request.RequestUpgradeServiceGroupChat{
			ChatID:  "wraaaaaaaaaaaaaaaa",
			Wording: "欢迎加入你的专属服务群",
		}, // 推荐的客户群，type等于2时有效
	}
	res, err := services.WeComApp.AccountServiceCustomer.UpgradeService(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 为客户取消推荐
// https://work.weixin.qq.com/api/doc/90000/90135/94674
func APIAccountServiceCustomerCancelUpgradeService(c *gin.Context) {
	openKFID := c.DefaultQuery("openKFID", "kfxxxxxxxxxxxxxx")
	externalUserID := c.DefaultQuery("externalUserID", "kfxxxxxxxxxxxxxx")
	res, err := services.WeComApp.AccountServiceCustomer.CancelUpgradeService(c.Request.Context(), openKFID, externalUserID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}
