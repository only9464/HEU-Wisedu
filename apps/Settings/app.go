package Settings

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const version_code = "8.8.7"

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// 供前端直接调用的函数都写在这里

func (s *App) Settings(a, b int) int {
	return settings(a, b)
}

func (s *App) Check_now_is_latest(currentVersion, latestVersion string) bool {
	return check_now_is_latest(currentVersion, latestVersion)
}

func (s *App) Get_latest_version_code() string {
	// fmt.Println("latestVersion", latestVersion)
	return get_latest_version_code()
}
func (s *App) Get_version_code() string {
	return version_code
}

// Get_current_program_path 返回当前程序的完整路径
func (a *App) Get_current_program_path() string {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("无法获取可执行文件路径:", err)
		return ""
	}
	execDir := filepath.Dir(exePath)
	return execDir
	// absPath, err := filepath.Abs(exePath)
	// if err != nil {
	// 	fmt.Println("无法获取绝对路径:", err)
	// 	return ""
	// }
	// return absPath
}

func (a *App) Get_config_file_path() string {
	// 检查当前目录下是否存在config.json文件，如果存在，则返回该文件的完整路径
	// 如果不存在，则返回空字符串
	filePath := filepath.Join(a.Get_current_program_path(), "config.json")
	if _, err := os.Stat(filePath); err == nil {
		return filePath
	}
	return ""
}

func (a *App) Download_config_file() bool {
	// 将 "https://ghproxy.mioe.me/https://raw.githubusercontent.com/only9464/Acrylic/master/build/bin/config.json" 下载到当前目录下并保存成config.json
	url := "https://ghproxy.mioe.me/https://raw.githubusercontent.com/only9464/Acrylic/master/config.json"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error occurred:", err)
		return false
	}
	// 如果状态码为200，则下载成功
	if response.StatusCode == 200 {
		fmt.Println("配置文件下载成功")
	} else {
		fmt.Println("配置文件下载失败")
		return false
	}
	defer response.Body.Close()
	// 将下载到的内容写入到当前目录下的config.json文件中
	filePath := filepath.Join(a.Get_current_program_path(), "config.json")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error occurred:", err)
		return false
	}
	defer file.Close()
	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("Error occurred:", err)
		return false
	}
	return true
}

func (a *App) UpdateAndRestart() {
	Update()
}
