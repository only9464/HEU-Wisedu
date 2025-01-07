package XGKC

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (a *App) GetXGKC(Authorization string, batchId string) (map[string]interface{}, error) {
	// 构造请求数据，严格按照Python代码
	jsonData := map[string]interface{}{
		"teachingClassType": "XGKC",
		"pageNumber":        1,
		"pageSize":          1000,
		"orderBy":           "",
		"campus":            "01",
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

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 解析响应
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
