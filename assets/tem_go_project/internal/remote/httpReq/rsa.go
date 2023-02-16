// coding: utf-8
// @Author : lryself
// @Date : 2022/1/26 15:19
// @Software: GoLand

package httpReq

import (
	"encoding/json"
	"errors"
	"fmt"
	"tem_go_project/internal/globals/codes"
)

func PostRSADecryptServer(CipherText string) (string, error) {
	serverHost, err := GetServerHostServer(4)
	if err != nil {
		return "", err
	}
	body, err := HttpPostServer(fmt.Sprintf("http://%s/api/rsa/decrypt", serverHost), map[string]interface{}{
		"CipherText": CipherText,
	})
	if err != nil {
		return "", err
	}

	var r map[string]interface{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return "", err
	}
	if r["code"].(string) != codes.OK {
		return "", errors.New(r["message"].(string))
	}
	data := r["data"].(string)
	return data, nil
}
