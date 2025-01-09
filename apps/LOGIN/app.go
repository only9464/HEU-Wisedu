package login

import (
	"context"
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

func (a *App) QueryAllGrade(username, password, ocrAPIURL string) (map[string]interface{}, error) {
	return QueryAllGrade(username, password, ocrAPIURL)
}

// 添加保存文件的方法
func (a *App) SaveGradeToFile(data map[string]interface{}) error {
	return SaveGradeToFile(data)
}
