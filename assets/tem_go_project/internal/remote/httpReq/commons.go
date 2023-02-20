// coding: utf-8
// @Author : lryself
// @Date : 2022/1/18 10:57
// @Software: GoLand

package httpReq

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"math/rand"
	"net/http"
	"tem_go_project/internal/globals/codes"
)

func interfaceToMap(m interface{}) map[string]interface{} {
	return m.(map[string]interface{})
}

func HttpGetServer(reqUrl string) ([]byte, error) {
	resp, err := http.Get(reqUrl)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(string(body))
	}
	return body, nil
}

func HttpPostServer(reqUrl string, requestBody map[string]interface{}) ([]byte, error) {
	marshal, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(reqUrl, "application/json", bytes.NewReader(marshal))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(string(body))
	}
	return body, nil
}

func GetServerHostServer(serverID int) (string, error) {
	body, err := HttpGetServer(fmt.Sprintf("%s/api/server/list?ServerID=%d",
		viper.GetString("remote.RegistryCenterURL"), serverID))
	if err != nil {
		return "", err
	}
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}
	if result["code"].(string) != codes.OK {
		return "", errors.New(result["message"].(string))
	}
	data := interfaceToMap(result["data"])
	var hosts []string
	for _, host := range data["Host"].([]interface{}) {
		hosts = append(hosts, host.(string))
	}
	return hosts[rand.Intn(len(hosts))], nil
}
