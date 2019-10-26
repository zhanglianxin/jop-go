package jop_sdk

import (
	"testing"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/zhanglianxin/jop-go/config"
)

var (
	siteId string
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
	})
	conf := config.GetConfig("../config.toml")
	AppKey = conf.Jop.AppKey
	SecretKey = conf.Jop.SecretKey
	siteId = conf.Jop.SiteId
}

func TestMakeRequest(t *testing.T) {
	paramBs, _ := json.Marshal(map[string]interface{}{
		"promotionCodeReq": map[string]interface{}{
			"materialId": "https://item.jd.com/23484023378.html",
			"siteId":     siteId,
		},
	})
	params := map[string]string{
		"method":     "jd.union.open.promotion.common.get",
		"v":          "1.0",
		"param_json": string(paramBs),
	}
	header, body := makeRequest("GET", params)
	t.Logf("%#v, %#v", header.StatusCode(), string(body))
}

func TestMakeRequestNew(t *testing.T) {
	param := NewParam("jd.union.open.promotion.common.get", "1.0")
	param.ParamJson = map[string]interface{}{
		"promotionCodeReq": map[string]interface{}{
			"materialId": "https://item.jd.com/23484023378.html",
			"siteId":     siteId,
		},
	}
	header, body := makeRequestNew("POST", param)
	t.Logf("%#v, %#v", header.StatusCode(), string(body))
}
