package jop_sdk

// 京东联盟
type Union struct {
}

// 宙斯
type Jos struct {
}

func NewUnion() *Union {
	return &Union{}
}

func NewJos() *Jos {
	return &Jos{}
}

func (u *Union) QueryOrder(params map[string]interface{}) []byte {
	checkRequiredKeys([]string{"orderReq", "pageNo", "pageSize", "type", "time"}, params)
	param := NewParam("jd.union.open.order.query", "1.0")
	param.ParamJson = params
	_, body := makeRequestNew("GET", param)
	return body
}

func (u *Union) QueryJingfenGoods(params map[string]interface{}) []byte {
	checkRequiredKeys([]string{"goodsReq", "eliteId"}, params)
	param := NewParam("jd.union.open.order.query", "1.0")
	param.ParamJson = params
	_, body := makeRequestNew("GET", param)
	return body
}

func (u *Union) GetCommonPromotion(params map[string]interface{}) ([]byte) {
	checkRequiredKeys([]string{"promotionCodeReq", "materialId", "siteId"}, params)
	param := NewParam("jd.union.open.promotion.common.get", "1.0")
	param.ParamJson = params
	_, body := makeRequestNew("POST", param)
	return body
}

// 检查必选参数
func checkRequiredKeys(required []string, params map[string]interface{}) {
	for _, v := range required {
		if found, _ := findNested(params, v); !found {
			panic(v + " is required")
		}
	}
}

// https://eli.thegreenplace.net/2019/go-json-cookbook/
// findNested looks for a key named s in map m. If values in m map to other
// maps, findNested looks into them recursively. Returns true if found, and
// the value found.
func findNested(m map[string]interface{}, s string) (bool, interface{}) {
	// Try to find key s at this level
	for k, v := range m {
		if k == s {
			return true, v
		}
	}
	// Not found on this level, so try to find it nested
	for _, v := range m {
		nm, ok := v.(map[string]interface{})
		if ok {
			found, val := findNested(nm, s)
			if found {
				return found, val
			}
		}
	}
	// Not found recursively
	return false, nil
}
