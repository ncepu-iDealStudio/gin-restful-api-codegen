// coding: utf-8
// @Author : lryself
// @Date : 2022/4/13 19:17
// @Software: GoLand

package httpReq

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"tem_go_project/internal/globals/codes"
)

func PostJWTServer(token string) (map[string]string, error) {
	body, err := HttpPostServer(
		fmt.Sprintf("%s/api_1_0/jwt/verifyToken", viper.GetString("remote.UserCenter")),
		map[string]interface{}{
			"Token": token,
		})
	if err != nil {
		return nil, err
	}

	var r map[string]interface{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}
	if r["code"].(string) != codes.OK {
		return nil, errors.New(r["message"].(string))
	}
	data := r["data"].(map[string]string)
	return data, nil
}
