package Settings

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

// 将版本号字符串 (x.x.x) 转换为数字数组
func parse_version(version string) []int {
	// 去除可能存在的前后空格
	version = strings.TrimSpace(version)

	// 使用正则验证版本号格式 (x.x.x)
	re := regexp.MustCompile(`^\d+\.\d+\.\d+$`)
	if !re.MatchString(version) {
		return nil
	}

	// 分割版本号字符串
	parts := strings.Split(version, ".")
	var versionInts []int

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err == nil {
			versionInts = append(versionInts, num)
		} else {
			return nil
		}
	}

	return versionInts
}

// 以下都是供本包内的函数进行调用的
func settings(a, b int) int {
	fmt.Println("settings")
	return a / b
}

// 获取远程的最新版本号
func get_latest_version_code() string {
	url := "https://ghproxy.mioe.me/https://raw.githubusercontent.com/only9464/Acrylic/master/Version.code"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error occurred:", err)
		return ""
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return ""
		}
		return string(body)
	}
	return ""
}

// 检查当前版本是否是最新版本
func check_now_is_latest(currentVersion, latestVersion string) bool {
	// 解析本地版本号和远程版本号
	parsedCurrentVersion := parse_version(currentVersion)
	parsedLatestVersion := parse_version(latestVersion)

	if parsedCurrentVersion == nil || parsedLatestVersion == nil {
		fmt.Println("版本号格式错误")
		return false
	}

	// 补充缺失的版本号部分（如果少于3个部分，补充0）
	for len(parsedCurrentVersion) < 3 {
		parsedCurrentVersion = append(parsedCurrentVersion, 0)
	}
	for len(parsedLatestVersion) < 3 {
		parsedLatestVersion = append(parsedLatestVersion, 0)
	}

	// 比较两个版本号
	for i := 0; i < 3; i++ {
		if parsedCurrentVersion[i] < parsedLatestVersion[i] {
			// 当前版本号小于远程版本号，返回 false
			return false
		} else if parsedCurrentVersion[i] > parsedLatestVersion[i] {
			// 当前版本号大于远程版本号，返回 true
			return true
		}
	}

	// 版本号相等
	return true
}
