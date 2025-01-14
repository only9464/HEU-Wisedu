package only9464

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// AddClazz 实现选课功能
func (a *App) AddClazz(Authorization string, batchId string, clazzType string, clazzId string, secretVal string) (string, error) {
	// 构造请求URL
	apiUrl := "https://jwxk.hrbeu.edu.cn/xsxk/elective/hrbeu/add"

	// 构造表单数据
	formData := url.Values{}
	formData.Set("clazzType", clazzType)
	formData.Set("clazzId", clazzId)
	formData.Set("secretVal", secretVal)
	// formData.Set("chooseVolunteer", chooseVolunteer)

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

// BatchAddClazz 批量选(qiang)课
func (a *App) BatchAddClazz(Authorization string, batchId string, courses []map[string]string, interval float64, mode string) error {
	fmt.Printf("开始批量选(qiang)课，共 %d 门课程，模式：%s\n", len(courses), mode)

	if mode == "crazy" {
		return a.CrazyBatchAddClazz(Authorization, batchId, courses, interval)
	}
	return a.NormalBatchAddClazz(Authorization, batchId, courses, interval)
}

// NormalBatchAddClazz 老实人模式选(qiang)课（原通道模式）
func (a *App) NormalBatchAddClazz(Authorization string, batchId string, courses []map[string]string, interval float64) error {
	// 如果没有指定间隔，使用默认值0.1秒
	if interval == 0 {
		interval = 0.1
	}

	// 创建课程通道
	courseChan := make(chan map[string]string, len(courses))
	// 创建间隔时间通道
	intervalChan := make(chan float64, 1)
	intervalChan <- interval
	// 创建停止信号通道
	stopChan := make(chan struct{})

	// 监听停止事件
	runtime.EventsOn(a.ctx, "only9464:stop", func(optionalData ...interface{}) {
		close(stopChan)
	})

	// 监听间隔时间更新事件
	runtime.EventsOn(a.ctx, "only9464:interval", func(optionalData ...interface{}) {
		if len(optionalData) > 0 {
			if newInterval, ok := optionalData[0].(float64); ok {
				// 清空旧的间隔时间
				select {
				case <-intervalChan:
				default:
				}
				// 写入新的间隔时间
				intervalChan <- newInterval
			}
		}
	})

	// 将所有课程放入通道
	for _, course := range courses {
		courseChan <- course
	}

	// 记录成功的课程数
	successCount := 0

	// 单独的goroutine处理选(qiang)课
	go func() {
		defer close(courseChan)
		defer close(intervalChan)
		defer runtime.EventsOff(a.ctx, "only9464:interval")
		defer runtime.EventsOff(a.ctx, "only9464:stop")
		defer func() {
			// 发送任务完成事件
			runtime.EventsEmit(a.ctx, "only9464:complete")
		}()

		for course := range courseChan {
			// 检查是否收到停止信号
			select {
			case <-stopChan:
				return
			default:
			}

			result := map[string]interface{}{
				"clazzId":    course["clazzId"],
				"courseName": course["courseName"],
				"teacher":    course["teacher"],
				"message":    "排队中",
				"code":       0,
			}

			// 发送状态
			runtime.EventsEmit(a.ctx, "only9464:result", result)

			response, err := a.AddClazz(Authorization, batchId, course["clazzType"], course["clazzId"], course["secretVal"])
			if err != nil {
				result["message"] = "请求失败"
				result["code"] = 500
				runtime.EventsEmit(a.ctx, "only9464:result", result)
				// 检查是否收到停止信号
				select {
				case <-stopChan:
					return
				default:
					// 失败后重新加入队列
					courseChan <- course
				}
				// 获取当前间隔时间
				currentInterval := <-intervalChan
				intervalChan <- currentInterval
				time.Sleep(time.Duration(currentInterval * float64(time.Second)))
				continue
			}

			// 解析JSON响应
			var respData map[string]interface{}
			if err := json.Unmarshal([]byte(response), &respData); err != nil {
				result["message"] = "解析失败"
				result["code"] = 500
				runtime.EventsEmit(a.ctx, "only9464:result", result)
				// 检查是否收到停止信号
				select {
				case <-stopChan:
					return
				default:
					// 失败后重新加入队列
					courseChan <- course
				}
				// 获取当前间隔时间
				currentInterval := <-intervalChan
				intervalChan <- currentInterval
				time.Sleep(time.Duration(currentInterval * float64(time.Second)))
				continue
			}

			// 获取响应码和消息
			code := int(respData["code"].(float64))
			result["code"] = code

			// 根据响应码设置消息
			if code == 200 {
				result["message"] = "选(qiang)课成功"
				runtime.EventsEmit(a.ctx, "only9464:result", result)
				successCount++
			} else {
				if code == 403 {
					result["message"] = "请求过快"
				} else if code == 500 {
					result["message"] = "课容量已满"
				} else {
					result["message"] = respData["msg"].(string)
				}
				runtime.EventsEmit(a.ctx, "only9464:result", result)
				// 检查是否收到停止信号
				select {
				case <-stopChan:
					return
				default:
					// 不是200就重新加入队列
					courseChan <- course
				}
			}

			// 获取当前间隔时间
			currentInterval := <-intervalChan
			intervalChan <- currentInterval
			time.Sleep(time.Duration(currentInterval * float64(time.Second)))

			// 如果所有课程都成功了，就退出
			if successCount == len(courses) {
				return
			}
		}
	}()

	return nil
}

// CrazyBatchAddClazz 狂暴模式选(qiang)课（多线程模式）
func (a *App) CrazyBatchAddClazz(Authorization string, batchId string, courses []map[string]string, interval float64) error {
	// 如果没有指定间隔，使用默认值0.1秒
	if interval == 0 {
		interval = 0.1
	}

	// 创建停止信号通道
	stopChan := make(chan struct{})

	// 监听停止事件
	runtime.EventsOn(a.ctx, "only9464:stop", func(optionalData ...interface{}) {
		close(stopChan)
	})

	// 创建等待组
	var wg sync.WaitGroup
	successMap := sync.Map{}

	// 为每个课程创建一个goroutine
	for _, course := range courses {
		wg.Add(1)
		go func(course map[string]string) {
			defer wg.Done()

			for {
				// 检查是否收到停止信号
				select {
				case <-stopChan:
					return
				default:
				}

				// 检查是否已经选(qiang)课成功
				if _, ok := successMap.Load(course["clazzId"]); ok {
					return
				}

				result := map[string]interface{}{
					"clazzId":    course["clazzId"],
					"courseName": course["courseName"],
					"teacher":    course["teacher"],
					"message":    "选(qiang)课中",
					"code":       0,
				}

				// 发送状态
				runtime.EventsEmit(a.ctx, "only9464:result", result)

				response, err := a.AddClazz(Authorization, batchId, course["clazzType"], course["clazzId"], course["secretVal"])
				if err != nil {
					result["message"] = "请求失败"
					result["code"] = 500
					runtime.EventsEmit(a.ctx, "only9464:result", result)
					time.Sleep(time.Duration(interval * float64(time.Second)))
					continue
				}

				// 解析JSON响应
				var respData map[string]interface{}
				if err := json.Unmarshal([]byte(response), &respData); err != nil {
					result["message"] = "解析失败"
					result["code"] = 500
					runtime.EventsEmit(a.ctx, "only9464:result", result)
					time.Sleep(time.Duration(interval * float64(time.Second)))
					continue
				}

				// 获取响应码和消息
				code := int(respData["code"].(float64))
				result["code"] = code

				// 根据响应码设置消息
				if code == 200 {
					result["message"] = "选(qiang)课成功"
					runtime.EventsEmit(a.ctx, "only9464:result", result)
					successMap.Store(course["clazzId"], true)
					return
				} else {
					if code == 403 {
						result["message"] = "请求过快"
					} else if code == 500 {
						result["message"] = "课容量已满"
					} else {
						result["message"] = respData["msg"].(string)
					}
					runtime.EventsEmit(a.ctx, "only9464:result", result)
				}

				time.Sleep(time.Duration(interval * float64(time.Second)))
			}
		}(course)
	}

	// 启动一个goroutine等待所有任务完成
	go func() {
		wg.Wait()
		runtime.EventsOff(a.ctx, "only9464:stop")
		runtime.EventsEmit(a.ctx, "only9464:complete")
	}()

	return nil
}
