package Settings

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"
)

// 下载文件函数
func downloadFile(url, dest string) error {
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// 更新程序
func Update() {
	url := "https://ghproxy.mioe.me/https://raw.githubusercontent.com/only9464/Acrylic/master/build/bin/Acrylic.exe"

	// 获取当前程序路径
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return
	}

	// 临时文件路径（下载的文件加 .new 后缀）
	tempExePath := exePath + ".new"

	// 下载新版本的 A.exe
	err = downloadFile(url, tempExePath)
	if err != nil {
		fmt.Println("Error downloading new version:", err)
		return
	}

	// 稍微等待，确保下载完毕
	time.Sleep(1 * time.Second)

	// 创建批处理文件路径
	batchFilePath := exePath + ".bat"
	batchFile, err := os.Create(batchFilePath)
	if err != nil {
		fmt.Println("Error creating batch file:", err)
		return
	}
	defer batchFile.Close()

	// 批处理文件内容：替换原程序并启动新程序，同时删除自己
	batchContent := fmt.Sprintf(`@echo off
timeout /t 2
move /y "%s" "%s"
start "" "%s"
start "" cmd /c del "%%~f0"
exit`, tempExePath, exePath, exePath)

	_, err = batchFile.WriteString(batchContent)
	if err != nil {
		fmt.Println("Error writing batch file:", err)
		return
	}

	// 让当前程序退出
	fmt.Println("Successfully downloaded new version. Restarting...")

	// 启动批处理文件的一个新进程来替换文件并重启程序
	cmd := exec.Command("cmd", "/C", batchFilePath) // 使用 cmd.exe 启动批处理文件
	err = cmd.Start()
	if err != nil {
		fmt.Println("Error starting batch file:", err)
		return
	}

	// 程序退出
	os.Exit(0)
}
