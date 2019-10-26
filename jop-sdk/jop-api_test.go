package jop_sdk

import (
	"testing"
	"github.com/zhanglianxin/jop-go/config"
)

func TestUnion_QueryOrder(t *testing.T) {
	params := map[string]interface{}{
		"orderReq": map[string]interface{}{
			"pageNo":   "1",
			"pageSize": "20",
			"type":     "1",
			"time":     "20190924000000",
		},
	}
	body := NewUnion().QueryOrder(params)
	t.Log(string(body))
}

func TestUnion_GetCommonPromotion(t *testing.T) {
	siteId := config.GetConfig("../config.toml").Jop.SiteId
	t.Log(siteId)
	params := map[string]interface{}{
		"promotionCodeReq": map[string]interface{}{
			"materialId": "https://item.jd.com/23484023378.html",
			// "materialId": "https://u.jd.com/z2QZog",
			"siteId":     siteId,
		},
	}
	body := NewUnion().GetCommonPromotion(params)
	t.Log(string(body))
}

func TestUnion_QueryJingfenGoods(t *testing.T) {
	params := map[string]interface{}{
		"goodsReq": map[string]interface{}{
			"eliteId": "1",
		},
	}
	body := NewUnion().QueryJingfenGoods(params)
	t.Log(string(body))
}
