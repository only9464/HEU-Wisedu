package Settings

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

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
