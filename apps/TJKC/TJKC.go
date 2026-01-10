package TJKC

import (
	"HEU-Wisedu/apps/only9464"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (a *App) GetTJKC(Authorization string, batchId string) (map[string]interface{}, error) {
	// 构造请求数据
	jsonData := map[string]interface{}{
		"teachingClassType": "TJKC",
		"pageNumber":        1,
		"pageSize":          1000,
		"orderBy":           "",
		"campus":            "01",
		"SFCT":              "0",
	}

	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}

	// 创建请求
	req, err := http.NewRequest("POST", "https://jwxk.hrbeu.edu.cn/xsxk/elective/clazz/list", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", Authorization)
	req.Header.Set("batchId", batchId)
	req.AddCookie(&http.Cookie{
		Name:  "JSESSIONID",
		Value: only9464.GOLBALJSESSIONID,
	})
	req.AddCookie(&http.Cookie{
		Name:  "route",
		Value: only9464.GOLBALROUTE,
	})
	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	for _, c := range resp.Cookies() {
		if c.Name == "JSESSIONID" {
			fmt.Println("[GetTJKC]更新之前的全局JSESSIONID:", only9464.GOLBALJSESSIONID)
			fmt.Println("[GetTJKC]从响应中获取到JSESSIONID:", c.Value)
			only9464.GOLBALJSESSIONID = c.Value
			fmt.Println("[GetTJKC]更新后的全局JSESSIONID:", only9464.GOLBALJSESSIONID)
			break
		}
	}
	// 解析响应
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
