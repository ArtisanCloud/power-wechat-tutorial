package message

import (
  "encoding/xml"
  "github.com/ArtisanCloud/PowerLibs/fmt"
  "github.com/ArtisanCloud/PowerWeChat/src/kernel"
  "github.com/ArtisanCloud/PowerWeChat/src/kernel/contract"
  models2 "github.com/ArtisanCloud/PowerWeChat/src/kernel/models"
  "github.com/ArtisanCloud/PowerWeChat/src/work/server/handlers/models"
  "github.com/gin-gonic/gin"
  "io/ioutil"
  "net/http"
  "power-wechat-tutorial/services"
)

func TestBuffer(c *gin.Context) {

  textXML := "<xml><ToUserName><![CDATA[ww454dfb9d6f6d432a]]></ToUserName><FromUserName><![CDATA[WangChaoYi]]></FromUserName><CreateTime>1634401052</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[thioutrrr]]></Content><MsgId>7019699067840561924</MsgId><AgentID>1000008</AgentID></xml>"
  var md interface{}
  md2 := models.MessageText{}
  err := xml.Unmarshal([]byte(textXML), &md2)
  md = md2
  fmt.Dump(md, err)
}

// 回调配置
// https://work.weixin.qq.com/api/doc/90000/90135/90930
func CallbackVerify(c *gin.Context) {
  rs, err := services.WeComApp.Server.Serve(c.Request)
  if err != nil {
    panic(err)
  }

  text, _ := ioutil.ReadAll(rs.Body)
  c.String(http.StatusOK, string(text))

}

// 回调配置
// https://work.weixin.qq.com/api/doc/90000/90135/90930
func CallbackNotify(c *gin.Context) {

  rs, err := services.WeComApp.Server.Notify(c.Request, func(event contract.EventInterface) interface{} {
    fmt.Dump("event", event)
    //return  "handle callback"

    switch event.GetMsgType() {
    case models2.CALLBACK_MSG_TYPE_TEXT:
      msg := models.MessageText{}
      err := event.ReadMessage(&msg)
      if err != nil {
        println(err.Error())
        return "error"
      }
      fmt.Dump(msg)
    }

    return kernel.SUCCESS_EMPTY_RESPONSE

  })
  if err != nil {
    panic(err)
  }

  err = rs.Send(c.Writer)

  if err != nil {
    panic(err)
  }

}
