package login

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

// LoginToJwgl 登录教务管理系统
// 参数：统一身份认证的账号、密码、验证码识别API的URL
// 返回值：教务管理系统首页的cookies
func LoginToJwgl(username, password, ocrAPIURL string) (map[string]*http.Cookie, error) {
	// 第一次登录获取 astraeus_session
	_, lastAstraeusSession, err := loginAndFollowRedirects(username, password, "", ocrAPIURL)
	if err != nil {
		return nil, fmt.Errorf("第一次登录失败: %v", err)
	}

	// 使用获取到的 astraeus_session 进行第二次登录
	cookies2, _, err := loginAndFollowRedirects(username, password, lastAstraeusSession, ocrAPIURL)
	if err != nil {
		return nil, fmt.Errorf("第二次登录失败: %v", err)
	}

	return cookies2, nil
}

// 定义一个结构体来存储登录参数
type LoginParams struct {
	Lt        string
	Source    string
	Pid       string
	Execution string
	Cookies   map[string]*http.Cookie
}

// extractFormParams 提取登录表单中的关键参数
func extractFormParams(htmlContent string) map[string]string {
	params := make(map[string]string)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		return params
	}

	doc.Find("input.for-form").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("name")
		value, _ := s.Attr("value")
		if name == "lt" || name == "source" || name == "pid" || name == "execution" {
			params[name] = value
		}
	})
	return params
}

// getLoginParams 获取登录所需的所有参数
func getLoginParams(astraeusSession string) (*LoginParams, error) {
	session := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	initialURL := "https://jwgl-443.wvpn.hrbeu.edu.cn/jwapp/sys/emaphome/portal/index.do"
	params := url.Values{
		"origin": {initialURL},
		"reason": {"site jwgl-443.wvpn.hrbeu.edu.cn not found"},
	}

	// 获取第一个响应
	resp, err := session.Get(initialURL + "?" + params.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// fmt.Printf("\n获取登录参数第1次请求: %s\n", initialURL)
	// fmt.Printf("响应状态码: %d\n", resp.StatusCode)

	if resp.StatusCode == http.StatusFound {
		redirectURL := resp.Header.Get("Location")
		// fmt.Printf("下一个重定向: %s\n", redirectURL)

		jar, _ := cookiejar.New(nil)
		client := &http.Client{Jar: jar}

		// 如果提供了astraeus_session，添加到cookies
		if astraeusSession != "" {
			u, _ := url.Parse(redirectURL)
			client.Jar.SetCookies(u, []*http.Cookie{{
				Name:  "_astraeus_session",
				Value: astraeusSession,
			}})
		}

		resp, err = client.Get(redirectURL)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			formParams := extractFormParams(string(body))
			cookies := make(map[string]*http.Cookie)
			for _, cookie := range client.Jar.Cookies(resp.Request.URL) {
				cookies[cookie.Name] = cookie
			}

			return &LoginParams{
				Lt:        formParams["lt"],
				Source:    formParams["source"],
				Pid:       formParams["pid"],
				Execution: formParams["execution"],
				Cookies:   cookies,
			}, nil
		}
	}
	return nil, fmt.Errorf("无法获取登录参数")
}

// getCaptcha 获取并识别验证码
func getCaptcha(ocrAPIURL string) (string, string, error) {
	// 第一步：获取验证码图片和token
	resp, err := http.Get("https://cas-443.wvpn.hrbeu.edu.cn/sso/apis/v2/open/captcha?imageWidth=100&captchaSize=4")
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	var result struct {
		Img   string `json:"img"`
		Token string `json:"token"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", "", err
	}

	// 第二步：识别验证码
	data := url.Values{}
	data.Set("image", result.Img)
	data.Set("probability", "false")
	data.Set("png_fix", "false")

	ocrResp, err := http.PostForm(ocrAPIURL, data)
	if err != nil {
		return "", "", err
	}
	defer ocrResp.Body.Close()

	body, err := io.ReadAll(ocrResp.Body)
	if err != nil {
		return "", "", err
	}

	code := gjson.Get(string(body), "code").String()
	message := gjson.Get(string(body), "message").String()
	captcha := gjson.Get(string(body), "data").String()

	if code == "200" && message == "Success" {
		return captcha, result.Token, nil
	}
	return "", "", fmt.Errorf("验证码识别失败: %s", message)
}

// login 执行登录请求
func login(INGRESSCOOKIE, JSESSIONID, username, password, captcha, token, lt, source, execution string) (*http.Response, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	data := url.Values{}
	data.Set("username", username)
	data.Set("password", password)
	data.Set("captcha", captcha)
	data.Set("token", token)
	data.Set("_eventId", "submit")
	data.Set("lt", lt)
	data.Set("source", source)
	data.Set("execution", execution)

	req, err := http.NewRequest("POST", "https://cas-443.wvpn.hrbeu.edu.cn/cas/login", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 设置cookies
	req.AddCookie(&http.Cookie{Name: "INGRESSCOOKIE", Value: INGRESSCOOKIE})
	req.AddCookie(&http.Cookie{Name: "JSESSIONID", Value: JSESSIONID})
	req.AddCookie(&http.Cookie{Name: "X_CAPTCHA", Value: "true"})

	return client.Do(req)
}

// loginAndFollowRedirects 执行登录流程并跟踪所有重定向
func loginAndFollowRedirects(username, password, astraeusSession string, ocrAPIURL string) (map[string]*http.Cookie, string, error) {
	// 获取登录参数
	loginParams, err := getLoginParams(astraeusSession)
	if err != nil {
		return nil, "", err
	}

	// 获取验证码和token
	captcha, token, err := getCaptcha(ocrAPIURL)
	if err != nil {
		return nil, "", err
	}

	// 创建session来处理重定向
	jar, _ := cookiejar.New(nil)
	session := &http.Client{
		Jar: jar,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	// 执行登录请求
	cookies := loginParams.Cookies
	resp, err := login(
		cookies["INGRESSCOOKIE"].Value,
		cookies["JSESSIONID"].Value,
		username,
		password,
		captcha,
		token,
		loginParams.Lt,
		loginParams.Source,
		loginParams.Execution,
	)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	// fmt.Println("\n第1次请求：https://cas-443.wvpn.hrbeu.edu.cn/cas/login")
	// fmt.Printf("响应状态码: %d\n", resp.StatusCode)

	// 更新session的cookies
	for _, cookie := range resp.Cookies() {
		u, _ := url.Parse("https://cas-443.wvpn.hrbeu.edu.cn")
		session.Jar.SetCookies(u, []*http.Cookie{cookie})
	}

	// 处理重定向
	currentURL := resp.Header.Get("Location")
	redirectCount := 2
	var lastAstraeusSession string

	// 创建一个map来存储所有的cookies
	allCookies := make(map[string]*http.Cookie)

	// 保存第一次请求的cookies
	for _, cookie := range resp.Cookies() {
		allCookies[cookie.Name] = cookie
	}

	for resp.StatusCode >= 300 && resp.StatusCode < 400 {
		// fmt.Printf("\n第%d次请求: %s\n", redirectCount, currentURL)

		if astraeusSession != "" && redirectCount == 2 {
			u, _ := url.Parse(currentURL)
			session.Jar.SetCookies(u, []*http.Cookie{{
				Name:  "_astraeus_session",
				Value: astraeusSession,
			}})
		}

		resp, err = session.Get(currentURL)
		if err != nil {
			return nil, "", err
		}
		defer resp.Body.Close()

		// fmt.Printf("响应状态码: %d\n", resp.StatusCode)

		// 保存每次重定向的cookies
		for _, cookie := range resp.Cookies() {
			allCookies[cookie.Name] = cookie
			if cookie.Name == "_astraeus_session" {
				lastAstraeusSession = cookie.Value
			}
		}

		if location := resp.Header.Get("Location"); location != "" {
			currentURL = location
			// fmt.Printf("下一个重定向: %s\n", currentURL)
		} else {
			// 最后一次请求，获取所有域名的cookies
			domains := []string{
				"https://cas-443.wvpn.hrbeu.edu.cn",
				"https://jwgl-443.wvpn.hrbeu.edu.cn",
				"https://wvpn.hrbeu.edu.cn",
			}

			for _, domain := range domains {
				u, _ := url.Parse(domain)
				for _, cookie := range session.Jar.Cookies(u) {
					allCookies[cookie.Name] = cookie
				}
			}
			break
		}

		redirectCount++
	}

	// fmt.Println("\n最终Cookies:")
	// fmt.Printf("%+v\n", allCookies)

	// fmt.Println("\n最终响应内容:")
	// if len(body) > 500 {
	//     fmt.Println(string(body[:500]))
	// } else {
	//     fmt.Println(string(body))
	// }

	return allCookies, lastAstraeusSession, nil
}
