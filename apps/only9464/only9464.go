package only9464

import (
	"fmt"
	"github.com/tidwall/gjson"
	_ "github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// AddClazz 实现选课功能
func (a *App) AddClazz(Authorization string, batchId string, clazzType string, clazzId string, secretVal string, chooseVolunteer string) (string, error) {
	// 构造请求URL
	apiUrl := "https://jwxk.hrbeu.edu.cn/xsxk/elective/hrbeu/add"

	// 构造表单数据
	formData := url.Values{}
	formData.Set("clazzType", clazzType)
	formData.Set("clazzId", clazzId)
	formData.Set("secretVal", secretVal)
	formData.Set("chooseVolunteer", chooseVolunteer)

	// 创建请求
	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(formData.Encode()))
	if err != nil {
		return "", err
	}

	// 设置请求头
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Authorization", Authorization)
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", "https://jwxk.hrbeu.edu.cn")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", fmt.Sprintf("https://jwxk.hrbeu.edu.cn/xsxk/elective/grablessons?batchId=%s", batchId))
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0")
	req.Header.Set("batchId", batchId)
	req.Header.Set("sec-ch-ua", `"Microsoft Edge";v="131", "Chromium";v="131", "Not_A Brand";v="24"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)

	// 设置Cookie
	req.AddCookie(&http.Cookie{
		Name:  "Authorization",
		Value: Authorization,
	})

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 直接返回JSON字符串
	return string(body), nil
}

// VerifyToken 验证Token有效性
func (a *App) VerifyToken(Authorization string, batchId string) (string, error) {
	// 构造请求URL
	apiUrl := "https://jwxk.hrbeu.edu.cn/xsxk/elective/user"

	// 构造表单数据
	formData := url.Values{}
	formData.Set("batchId", batchId)

	// 创建请求
	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(formData.Encode()))
	if err != nil {
		return "", err
	}

	// 设置请求头
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Authorization", Authorization)
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", "https://jwxk.hrbeu.edu.cn")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://jwxk.hrbeu.edu.cn/xsxk/profile/index.html")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0")
	req.Header.Set("sec-ch-ua", `"Microsoft Edge";v="131", "Chromium";v="131", "Not_A Brand";v="24"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)

	// 设置Cookie
	req.AddCookie(&http.Cookie{
		Name:  "Authorization",
		Value: Authorization,
	})

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 直接返回JSON字符串
	return string(body), nil
}

// DelClazz 实现退选功能
func (a *App) DelClazz(Authorization string, batchId string, clazzType string, clazzId string, secretVal string) (string, error) {
	// 构造请求URL
	apiUrl := "https://jwxk.hrbeu.edu.cn/xsxk/elective/clazz/del"

	// 构造表单数据
	formData := url.Values{}
	formData.Set("clazzType", clazzType)
	formData.Set("clazzId", clazzId)
	formData.Set("secretVal", secretVal)

	// 创建请求
	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(formData.Encode()))
	if err != nil {
		return "", err
	}

	// 设置请求头
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Authorization", Authorization)
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", "https://jwxk.hrbeu.edu.cn")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", fmt.Sprintf("https://jwxk.hrbeu.edu.cn/xsxk/elective/grablessons?batchId=%s", batchId))
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0")
	req.Header.Set("batchId", batchId)
	req.Header.Set("sec-ch-ua", `"Microsoft Edge";v="131", "Chromium";v="131", "Not_A Brand";v="24"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)

	// 设置Cookie
	req.AddCookie(&http.Cookie{
		Name:  "Authorization",
		Value: Authorization,
	})

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 直接返回JSON字符串
	return string(body), nil
}

// TaskListAddClazz 批量添加课程
func (a *App) TaskListAddClazz(Authorization string, batchID string, taskListJSON string) ([]string, []error) {
	// 解析 JSON 数组
	result := gjson.Parse(taskListJSON)
	if !result.IsArray() {
		return nil, []error{fmt.Errorf("JSON 数据不是数组")}
	}

	results := make([]string, len(result.Array()))
	errors := make([]error, len(result.Array()))

	for i, task := range result.Array() {
		fmt.Println("TaskListAddClazz", task.Get("clazzType").String())
		fmt.Println("TaskListAddClazz", task.Get("JXBID").String())
		fmt.Println("TaskListAddClazz", task.Get("secretVal").String())
		clazzType := task.Get("clazzType").String()
		clazzId := task.Get("clazzId").String()
		secretVal := task.Get("secretVal").String()
		//chooseVolunteer := task.Get("chooseVolunteer").String()
		chooseVolunteer := "1"

		result, err := a.AddClazz(Authorization, batchID, clazzType, clazzId, secretVal, chooseVolunteer)
		if err != nil {
			errors[i] = err
		} else {
			results[i] = result
		}
	}

	return results, errors
}
