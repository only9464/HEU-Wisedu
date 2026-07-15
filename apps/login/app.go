package login

import (
	"context"
	"net/http"
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

func (a *App) QueryAllGrade(cookies map[string]*http.Cookie) (map[string]interface{}, error) {
	return QueryAllGrade(cookies)
}

// 添加保存文件的方法
func (a *App) SaveGradeToFile(data map[string]interface{}) error {
	return SaveGradeToFile(data)
}

func (a *App) LoginToJwgl(username, password, captcha1, token1, captcha2, token2 string) (map[string]*http.Cookie, error) {
	return LoginToJwgl(username, password, captcha1, token1, captcha2, token2)
}

// GetCaptcha 获取验证码
func (a *App) GetCaptcha() (map[string]string, error) {
	img, token, err := GetCaptcha()
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"img":   img,
		"token": token,
	}, nil
}
