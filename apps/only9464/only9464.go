package only9464

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// AddClazz 实现选课功能
func (a *App) AddClazz(Authorization string, batchId string, clazzType string, clazzId string, secretVal string, chooseVolunteer string) (map[string]interface{}, error) {
	// 构造请求数据
	data := map[string]interface{}{
		"clazzType":       clazzType,
		"clazzId":         clazzId,
		"secretVal":       secretVal,
		"chooseVolunteer": chooseVolunteer,
	}
	fmt.Println(data)
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// 创建请求
	req, err := http.NewRequest("POST", "https://jwxk.hrbeu.edu.cn/xsxk/elective/hrbeu/add", bytes.NewBuffer(jsonBytes))
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
	fmt.Println(result)
	return result, nil
}
