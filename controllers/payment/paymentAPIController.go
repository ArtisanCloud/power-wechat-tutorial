package payment

import (
  "github.com/ArtisanCloud/power-wechat/src/kernel/models"
  "github.com/ArtisanCloud/power-wechat/src/kernel/power"
  "github.com/ArtisanCloud/power-wechat/src/payment/notify/request"
  "github.com/gin-gonic/gin"
  "log"
  "net/http"
  "power-wechat-tutorial/services"
)

func APIMakeOrder(c *gin.Context) {

  // 下单
  response, err := services.PaymentApp.Order.JSAPITransaction(&power.HashMap{
    "amount": &power.HashMap{
      "total":    1,
      "currency": "CNY",
    },
    "attach":       "自定义数据说明",
    "description":  "Image形象店-深圳腾大-QQ公仔",
    "mchid":        "1611854986",
    "notify_url":   "https://pay.wangchaoyi.com/wx/notify",
    "out_trade_no": "5519778939773395659222298000", // 这里是商户订单号，不能重复提交给微信
    "payer": &power.HashMap{
      "openid": "oAuaP0TRUMwP169nQfg7XCEAw3HQ", // 用户的openid， 记得也是动态的。
    },
  })

  if err != nil {

    log.Printf("error: %s", err)
    c.JSON(400, response)
    return
  }

  payConf, err := services.PaymentApp.JSSDK.BridgeConfig(response.PrepayID, true)
  if err != nil {
    panic(err)
  }

  c.JSON(200, payConf)
}

func APIQueryOrder(c *gin.Context) {

  traceNo := c.Query("traceNo")

  rs, err := services.PaymentApp.Order.QueryByOutTradeNumber(traceNo)
  if err != nil {
    panic(err)
  }
  c.JSON(http.StatusOK, rs)

}

func APICloseOrder(c *gin.Context) {
  traceNo := c.Query("traceNo")
  log.Printf("traceNo: %s", traceNo)

  rs, err := services.PaymentApp.Order.Close(traceNo)
  if err != nil {
    log.Println("出错了： ", err)
    c.String(400, err.Error())
    return
  }
  c.JSON(http.StatusOK, rs)

}

func CallbackWXNotify(c *gin.Context) {
  //rs, err := PaymentApp.Order.QueryByOutTradeNumber("商户系统的内部订单号 [out_trade_no]")
  //rs, err := PaymentApp.Order.QueryByTransactionId("微信支付订单号 [transaction_id]")
  //_, err := services.PaymentApp.HandlePaidNotify(c.Request, func(message *power.HashMap, content *power.HashMap, fail string) interface{} {
  //if content == nil || (*content)["out_trade_no"] == nil {
  //  return fail("no content notify")
  //}
  //return true
  //})

  res, err := services.PaymentApp.HandlePaidNotify(
    c.Request,
    func(message *request.RequestNotify, transaction *models.Transaction, fail func(message string)) interface{} {

      // 看下支付通知事件状态
      // 这里可能是微信支付失败的通知，所以可能需要在数据库做一些记录，然后告诉微信我处理完成了。
      if message.EventType != "TRANSACTION.SUCCESS" {
        return true
      }

      if transaction.OutTradeNo != "" {
        // 这里对照自有数据库里面的订单做查询以及支付状态改变
        log.Printf("订单号：%s 支付成功", transaction.OutTradeNo)
      } else {
        // 因为微信这个回调不存在订单号，所以可以告诉微信我还没处理成功，等会它会重新发起通知
        // 如果不需要，直接返回true即可
        fail("payment fail")
        return nil
      }
      return true
    },
  )

  // 这里可能是因为不是微信官方调用的，无法正常解析出transaction和message，所以直接抛错。
  if err != nil {
    panic(err)
  }

  // 这里根据之前返回的是true或者fail，框架这边自动会帮你回复微信
  err = res.Send(c.Writer)

  if err != nil {
    panic(err)
  }

}
