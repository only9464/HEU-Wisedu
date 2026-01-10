# HEU-Wisedu

HEU-Wisedu是一款基于[Wails](https://wails.io/zh-Hans/docs/introduction) `框架`、[哈尔滨工程大学教务系统](http://www.hrbeu.edu.cn/) `API`、[Acrylic](https://github.com/only9464/Acrylic) `模版`开发的跨平台桌面应用程序，主要功能为查看课程信息、选课等。

[![Downloads](https://img.shields.io/github/downloads/only9464/HEU-Wisedu/total.svg?label=HEU-Wisedu软件总下载量&color=blue)](https://github.com/only9464/HEU-Wisedu/releases)
[![Release](https://img.shields.io/github/v/release/only9464/HEU-Wisedu?label=最新版本&color=blue)](https://github.com/only9464/HEU-Wisedu/releases/latest)
![Visitors](https://visitor-badge.laobi.icu/badge?page_id=github.only9464.HEU-Wisedu)
[![Stars](https://img.shields.io/github/stars/only9464/HEU-Wisedu?label=Star数量&color=yellow)](https://github.com/only9464/HEU-Wisedu/stargazers)
[![Forks](https://img.shields.io/github/forks/only9464/HEU-Wisedu?label=Forks数量&color=blue)](https://github.com/only9464/HEU-Wisedu/network/members)

> [!CAUTION]
> 【项目介绍】本软件完全免费且部分代码开源。仅供学习交流使用，切勿用于商业用途，一经发现，将拉黑、并依规向校方举报！！！
>
> 【温馨提示】**本软件仅在Gihtub发布且免费，任何其他渠道均为假冒，各位擦亮双眼，谨防上当受骗，否则后果自负**
>
> 【开源说明】2026年1月10日，发现有司马玩意儿恶意使用该免费开源软件牟取利益，这违背了我给大家提供相同便利的初衷，所以从此只逐步开放部分核心源代码
>
> 【严正声明】**该软件新增“一机一号”机制，严防倒卖，如仍发现存在任何使用该免费软件牟利的行为，欢迎各位向我发送邮件举报，事后我在向校方举报的同时，根据本项目开源协议依法起诉**
>
> 【免责声明】任何试图突破“一机一号”机制(包括但不限于对软件进行逆向工程)的行为造成系统损坏的，**后果自负**！！！
>

> [!TIP]
>
> Github地址：[https://github.com/only9464/HEU-Wisedu](https://github.com/only9464/HEU-Wisedu)
>
> Gitee地址：[https://gitee.com/only9464/HEU-Wisedu](https://gitee.com/only9464/HEU-Wisedu)
>
> 欢迎各位同志提交[Pull Requests](https://github.com/only9464/HEU-Wisedu/pulls)，共同完善开放的源代码。[![Pull Requests](https://img.shields.io/github/issues-pr/only9464/HEU-Wisedu?label=PRs&color=blue)](https://github.com/only9464/HEU-Wisedu/pulls)
>
> **BUG反馈**或者**功能建议**欢迎提交[Issues](https://github.com/only9464/HEU-Wisedu/issues)或者发邮件：[sky9464@qq.com](mailto:sky9464@qq.com)[![Issues](https://img.shields.io/github/issues/only9464/HEU-Wisedu?label=Issues&color=blue)](https://github.com/only9464/HEU-Wisedu/issues)
>
> 若软件对你有帮助，还请多多支持🙏🙏🙏！！！Ciallo～(∠・ω< )⌒★~~ [支持作者](#四支持作者)
>
> <details>
> <summary>支持作者</summary>
>
> ![支持作者](image/README/image.png)
>
> </details>
>

> [!IMPORTANT]
> [![Contributors](https://img.shields.io/github/contributors/only9464/HEU-Wisedu?label=当前贡献者数量&color=blue)](https://github.com/only9464/HEU-Wisedu/graphs/contributors)如果你和我一样愿意无偿通过自己的技术为大家提供无差别的便利，认为自己有能力且愿意为该项目开发做出贡献，欢迎致信[sky9464@qq.com](mailto:sky9464@qq.com)，请告知来意与你的项目经历，最好带上你的Github用户名。审核通过将开放全部代码，共同为该项目添砖加瓦。
>

# 一、优势

- [X] **免费**(希望各位尊重他人劳动成果，共同打击倒卖行为🙏🙏🙏)
- [X] **直接**调用[选课系统](https://jwxk.hrbeu.edu.cn/xsxk/auth/cas)API进行课程相关操作，省去繁杂加载，**高峰期选(qiang)课快人一步**（
- [X] ~~支持跨平台，支持Windows、MacOS、Linux~~（现在只提供windows可执行程序）
- [X] 根据已修学分(查成绩)，选课更方便、快捷、具有目的性
- [X] 采用Golang的通道技术多个课程依次选(qiang)课、多线程多个课程同时选(qiang)课，自己设置间隔时间（操作简单，一看就会）
- [X] 界面简约美观(乐)，支持明暗双主题（暗色太拉胯了）（没调好，艹）

[更新日志](#八更新日志)

# 二、功能

- [X] 登录 `教务选课系统`、`教务管理系统(统一身份认证)`
- [X] 查看 `培养方案内课程`
- [X] 查看 `跨专业选修课`
- [X] 查看 `公选课`
- [X] 查看 `本批次已选课程`
- [X] 查看 `所有学期已选课程` 及其 `成绩` (说白了，只能查到出成绩的课)(**若查询失败，请先评教**)
- [X] 查看 `已修学分`
- [X] 退选
- [X] 自动选(qiang)课

[相关功能截图](#七功能截图)

# 三、下载

~~强烈建议各位学会自己下载源码编译成可执行程序，当然，如果你想偷懒：~~

【2026年1月10日22:00】从此以后**不再**提供**完整的源代码**，仅提供**部分源代码**和**windows可执行程序**

[![Downloads](https://img.shields.io/github/downloads/only9464/HEU-Wisedu/total.svg?label=软件总下载次数&color=blue)](https://github.com/only9464/HEU-Wisedu/releases)

- [X] Windows[下载](https://gh-proxy.com/https://github.com/only9464/HEU-Wisedu/releases/latest/download/HEU-Wisedu.exe)
- [ ] MacOS  [下载](https://gh-proxy.com/https://github.com/only9464/HEU-Wisedu/releases/latest/download/HEU-Wisedu.exe)
- [ ] Linux  [下载](https://gh-proxy.com/https://github.com/only9464/HEU-Wisedu/releases/latest/download/HEU-Wisedu.exe)

~~如果想自己下载源码编译，可以参考下面的**二次开发[打包](#3打包)** 部分~~

# 四、支持作者

感谢以下同志的赞助！！！(仅展示2026年1月1日以来的记录)

|        赞助时间        |   ID   | 赞助额度 |         留言         |
| :--------------------: | :----: | :------: | :------------------: |
| 2026年1月14日 09:24:22 | Lil.🐑 |  ￥6.66  |       吊哦       |
| 2026年1月11日 12:24:17 | Qceeen |  ￥5.00  |       感谢学长       |
| 2026年1月7日 21:01:15 |  ***  | ￥10.00 | 哟西，你滴大大滴良民 |

更新不易，还请多多支持🙏🙏🙏！！！Ciallo～(∠・ω< )⌒★

![sponsor](image/README/image.png)


# 五、二次开发

## 1.安装依赖

**以下依赖按照顺序逐个安装即可：**

- [![Go](https://img.shields.io/github/go-mod/go-version/only9464/HEU-Wisedu?logo=go&label=Golang&color=00ADD8)](https://go.dev/)
- [![Wails](https://img.shields.io/github/v/release/wailsapp/wails?label=Wails&color=red&logo=wails)](https://wails.io)
- [![Node](https://img.shields.io/badge/Node.js-v20.12.2-green?logo=node.js)](https://nodejs.org/)
- [![npm](https://img.shields.io/badge/npm-v9.0.0-red?logo=npm)](https://www.npmjs.com/)
- [![Vue](https://img.shields.io/badge/vue-v3.5.13-green?logo=vue)](https://vuejs.org/)

### Windows

暂无

### Mac

暂无

### Linux

- libgtk-3-dev
- libwebkit2gtk-4.0-dev
- libglib2.0-dev

所需执行命令(仅在ubuntu-20.04.6-amd64测试通过，其余自测)：

```bash
sudo apt update
sudo apt-get install libgtk-3-dev libwebkit2gtk-4.0-dev libglib2.0-dev
export PKG_CONFIG_PATH=/usr/lib/x86_64-linux-gnu/pkgconfig:$PKG_CONFIG_PATH
```

> [!NOTE]
>
> 新版 `Linux`安装 `libwebkit2gtk-4.0-dev`编译应用时需要增加 `-tags webkit2_40`
> PS: 目前最新版的依赖不知道是什么版本，你可以先执行 `wails dev`或者 `wails build`查看报错信息来知道都需要安装什么

## 2.调试运行

**在项目的根目录下执行：**

```bash
wails dev
```

## 3.打包

**在项目的根目录下执行：**

```bash
wails build
```

执行完之后会在你项目根目录的 `build/bin`文件夹中看到可执行文件

`更多信息请参考：`[Acrylic 二次开发](https://github.com/only9464/Acrylic#%E4%BA%8C%E4%BA%8C%E6%AC%A1%E5%BC%80%E5%8F%91)

# 六、Star History

[![Star History Chart](https://api.star-history.com/svg?repos=only9464/HEU-Wisedu&type=Date)](https://star-history.com/#only9464/HEU-Wisedu&Date)

# 七、功能截图

![home](image/README/home.png)
![xuanke](image/README/xuanke.png)
![qiangke](image/README/qiangke.png)

# 八、更新日志

- VersionCode：2【2026年1月12日12:20:00】
  - 删除物理端口检测，修复部分同学点击即闪退的问题

- VersionCode：1【2026年1月11日12:00:00】
  - 删除旧版更新方式
  - 删除选课系统自动保存认证信息功能
  - 调整界面信息

- VersionCode：0【2026年1月10日21:56:19】
  - 更改选课系统登录方式，适配学校选课系统，改为统一身份认证
  - 删除选(qiang)课模式：狂暴模式
  - 增加一号一码绑定，采取手段打击代抢课牟利行为

- VersionCode：-1【2025年1月14日15:16:17】
  - 修复系统默认暗色导致表格字体显示不清楚的Bug
  - 新增选(qiang)课模式：老实人模式、狂暴模式
  - 更改默认选(qiang)课间隔时间：0.1s --> 0.275s
  - 优化登录界面
