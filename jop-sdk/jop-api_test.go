package jop_sdk

import (
	"testing"
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
	NewUnion().QueryOrder(params)
}
