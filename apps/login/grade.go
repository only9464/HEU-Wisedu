package login

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// QueryAllGrade 查询所有学期成绩
// 使用已有cookies进行查询
func QueryAllGrade(cookies map[string]*http.Cookie) (map[string]interface{}, error) {
	if cookies == nil {
		return nil, fmt.Errorf("cookies不能为空")
	}

	// 创建HTTP客户端
	client := &http.Client{}

	// 获取新的_WEU
	req, err := http.NewRequest("GET", "https://jwgl-443.wvpn.hrbeu.edu.cn/jwapp/sys/emaphome/appShow.do", nil)
	if err != nil {
		return nil, err
	}

	// 添加查询参数
	q := req.URL.Query()
	q.Add("id", "d71f7b57b4f348368f06c3e9a2a0988f")
	req.URL.RawQuery = q.Encode()

	// 添加必要的cookies
	for _, cookie := range [...]string{"GS_SESSIONID", "_WEU", "_webvpn_key"} {
		if c, ok := cookies[cookie]; ok {
			req.AddCookie(c)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 更新_WEU cookie（如果有）
	if newCookies := resp.Header.Values("Set-Cookie"); len(newCookies) > 0 {
		for _, cookie := range newCookies {
			if strings.Contains(cookie, "_WEU=") {
				parts := strings.Split(cookie, ";")
				weu := strings.TrimPrefix(parts[0], "_WEU=")
				if c, ok := cookies["_WEU"]; ok {
					c.Value = weu
				}
			}
		}
	}

	// 准备查询成绩的请求
	data := url.Values{}
	data.Set("querySetting", `[{"name":"SFYX","caption":"是否有效","linkOpt":"AND","builderList":"cbl_m_List","builder":"m_value_equal","value":"1","value_display":"是"},{"name":"SHOWMAXCJ","caption":"显示最高成绩","linkOpt":"AND","builderList":"cbl_m_List","builder":"m_value_equal","value":"0","value_display":"否"}]`)
	data.Set("*order", "-XNXQDM,-KCH,-KXH")
	data.Set("pageSize", "100")
	data.Set("pageNumber", "1")

	req, err = http.NewRequest("POST", "https://jwgl-443.wvpn.hrbeu.edu.cn/jwapp/sys/cjcx/modules/cjcx/xscjcx.do", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Referer", "https://jwgl-443.wvpn.hrbeu.edu.cn/jwapp/sys/cjcx/*default/index.do?THEME=purple&EMAP_LANG=zh")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Origin", "https://jwgl-443.wvpn.hrbeu.edu.cn")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	// 添加cookies
	for _, cookie := range [...]string{"GS_SESSIONID", "_WEU", "_webvpn_key"} {
		if c, ok := cookies[cookie]; ok {
			req.AddCookie(c)
		}
	}

	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	// 解析 JSON 到 map
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("解析成绩数据失败: %v", err)
	}

	// 将结果和cookies一起返回
	return map[string]interface{}{
		"data":    result,
		"cookies": cookies,
	}, nil
}

// SaveGradeToFile 将成绩数据保存到本地文件
func SaveGradeToFile(data map[string]interface{}) error {
	// 将数据格式化为JSON字符串
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("格式化JSON失败: %v", err)
	}

	// 获取当前可执行文件的路径
	execPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("获取程序路径失败: %v", err)
	}

	// 获取可执行文件所在的目录
	execDir := filepath.Dir(execPath)

	// 构建文件路径
	filePath := filepath.Join(execDir, "grade_data.json")

	// 写入文件
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("写入文件失败: %v", err)
	}

	return nil
}
