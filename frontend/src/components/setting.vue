<template>
  <div class="main-container">
    <h1>设置界面</h1>
    <div class="settings-options">

      <!-- 当前程序所在位置 -->
      <div class="setting-option">
        <span class="option-label">当前程序所在位置：</span>
        <div class="option-control">
          <span>{{ currentProgramPath }}</span>
          <el-button type="primary" size="small" color="#1e8fa3" @click="OpenFileExplorer(currentProgramPath)">打开文件夹</el-button>
        </div>
      </div>

      <!-- 选项-1 -->
      <div class="setting-option">
        <span class="option-label">当前版本号：</span>
        <div class="option-control">
          <span>{{ currentVersion }}</span>
          <div v-if="!is_latest_version">
            <span>(发现新版本：{{ latestVersion }}
                          <el-button type="primary" size="small" color="#1e8fa3" @click="UpdateAndRestart">下载并更新重启</el-button>
              )</span>
          </div>
        </div>
      </div>

      <!-- 选项0 - 更新源下拉选择及完整下载链接显示 -->
      <div class="setting-option">
        <span class="option-label">更新源：</span>
        <div class="option-control">
          <span
            class="download-link"
            @click="openUrlLink(currentDownloadLink)"
            v-if="currentDownloadLink"
          >
            {{ currentDownloadLink }}
          </span>
          <el-select
            v-model="selectedSource"
            @change="updateDownloadLink"
            placeholder="请选择更新源"
            class="adaptive-select"
          >
            <el-option
              v-for="source in sources"
              :key="source.value"
              :label="source.label"
              :value="source.value"
            ></el-option>
          </el-select>
        </div>
      </div>

      <!-- 选项1 - 切换深浅模式 -->
      <div class="setting-option">
        <span class="option-label">切换深浅模式</span>
        <div class="option-control">
          <el-switch
            v-model="isDarkMode"
            active-text="深色模式"
            inactive-text="浅色模式"
            inline-prompt
            style="--el-switch-on-color: #000000; --el-switch-off-color: #1e8fa3"
            @change="toggleTheme"
          ></el-switch>
        </div>
      </div>

      <!-- 配置文件 -->
      <div class="setting-option">
        <span class="option-label">配置文件：</span>
        <div class="option-control">
          <span v-if="!configFilePath">配置文件缺失</span>
          <span v-else @click="openFile(configFilePath)">{{ configFilePath }}</span>
          <el-button v-if="!configFilePath" type="primary" size="small" color="#1e8fa3" @click="DownloadConfigFile">立即下载</el-button>
        </div>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
// 如果需要按需引入 Element Plus 组件，请取消注释以下行
// import { ElSelect, ElOption } from 'element-plus'

// 定义响应式变量
const latestVersion = ref('')
const currentVersion = ref('')
const is_latest_version = ref(true)
const currentProgramPath = ref('')

// 定义可选的更新源
const sources = ref([
  { label: 'Github', value: 'github', downloadLink: 'https://github.com/only9464/Acrylic' },
  { label: 'Gitlab', value: 'gitlab', downloadLink: 'https://gitlab.com/only9464/Acrylic' },
  { label: 'Gitee', value: 'gitee', downloadLink: 'https://gitee.com/only9464/Acrylic' },
  { label: 'Atom', value: 'atomgit', downloadLink: 'https://atomgit.com/only9464/Acrylic' },
  // 可以根据需要添加更多源
])

// 定义选中的更新源，默认为Github
const selectedSource = ref('github')

// 定义当前的下载链接
const currentDownloadLink = ref('')

// 定义配置文件路径
const configFilePath = ref('')

// 定义深色模式状态
const isDarkMode = ref(false)

// 更新下载链接的方法
const updateDownloadLink = () => {
  const source = sources.value.find(src => src.value === selectedSource.value)
  currentDownloadLink.value = source ? source.downloadLink : ''
  // console.log(`当前选择的下载链接: ${currentDownloadLink.value}`)
}

// 监视 selectedSource 的变化，自动更新下载链接
watch(selectedSource, () => {
  updateDownloadLink()
})

// 定义执行加载时的异步函数
const executeOnLoad = async () => {
  try {
    console.log('设置页面已加载，执行初始化函数')
    // 获取当前程序所在位置
    const current_program_path = await window.go.Settings.App.Get_current_program_path()
    currentProgramPath.value = current_program_path
    // 获取配置文件路径
    configFilePath.value = await window.go.Settings.App.Get_config_file_path()
    // 异步获取当前版本号
    const current_version_code = await window.go.Settings.App.Get_version_code()
    currentVersion.value = current_version_code

    // 异步获取最新的版本号，并设置到 latestVersion
    const latest_version_code = await window.go.Settings.App.Get_latest_version_code()
    latestVersion.value = latest_version_code

    // 检查当前版本是否为最新版本
    const now_is_latest_version = await window.go.Settings.App.Check_now_is_latest(current_version_code, latest_version_code)
    is_latest_version.value = now_is_latest_version  // 正确: 使用 .value 更新

    // 设置初始主题模式
    setInitialTheme()
    
    // 其他初始化逻辑
  } catch (error) {
    console.error('执行初始化函数时出错:', error)
  }
}

// 设置初始主题模式
const setInitialTheme = () => {
  // 根据系统或其他逻辑设置初始主题
  // 这里假设默认是浅色模式
  isDarkMode.value = false
}

// 监视 isDarkMode 的变化，切换主题
const toggleTheme = (value) => {
  if (value) {
    window.runtime.WindowSetDarkTheme()
  } else {
    window.runtime.WindowSetLightTheme()
  }
}

onMounted(() => {
  executeOnLoad()
  updateDownloadLink()
})

// 其他响应式数据和方法
const downloadPath = ref('')
const userDescription = ref('')

const toggleNotifications = (event) => {
  const isEnabled = event.target.checked
  console.log(`通知已${isEnabled ? '开启' : '关闭'}`)
  // 添加处理逻辑
}

const handleDownloadPath = () => {
  console.log(`下载路径: ${downloadPath.value}`)
  // 添加处理逻辑
}

const handleUserDescription = () => {
  console.log(`用户描述: ${userDescription.value}`)
  // 添加处理逻辑
}

// 定义打开下载链接的方法
const openUrlLink = (url) => {
  console.log('尝试打开下载链接:', url)
  if (window.runtime && typeof window.runtime.BrowserOpenURL === 'function') {
    window.runtime.BrowserOpenURL(url)
  } else {
    console.error('BrowserOpenURL 方法未定义或不可用')
    // 备用方案，例如直接在新标签页打开链接
    window.open(url, '_blank')
  }
}
const openFile = (filePath) => {
  console.log('尝试打开配置文件:', filePath)
  if (window.runtime && typeof window.runtime.BrowserOpenURL === 'function') {
    window.runtime.BrowserOpenURL("file:///"+filePath)
  } else {
    console.error('BrowserOpenURL 方法未定义或不可用')
    // 备用方案，例如直接在新标签页打开链接
    window.open("file:///"+filePath, '_blank')
  }
}
const OpenFileExplorer = (filePath) => {
  console.log('尝试下载配置文件:', filePath)
  if (window.runtime && typeof window.runtime.BrowserOpenURL === 'function') {
    window.runtime.BrowserOpenURL(filePath)
  } else {
    console.error('BrowserOpenURL 方法未定义或不可用')
    // 备用方案，例如直接在新标签页打开链接
    window.open(filePath, '_blank')
  }
}
const DownloadConfigFile = async () => {
  const result = await window.go.Settings.App.Download_config_file()
  if (result) {
    try {
      // 等待配置文件路径的 Promise 完成
      const path = await window.go.Settings.App.Get_config_file_path()
      configFilePath.value = path
      console.log(`配置文件下载成功，路径: ${configFilePath.value}`)
    } catch (error) {
      console.error('获取配置文件路径时出错:', error)
    }
  } else {
    console.error('下载配置文件失败')
  }
}
const UpdateAndRestart = () => {
  window.go.Settings.App.UpdateAndRestart()
}
</script>

<style scoped>
.main-container {
  padding: 0px;
}

.settings-options {
  display: flex;
  flex-direction: column;
  gap: 16px;
  margin-top: 20px;
}

.setting-option {
  display: flex;
  justify-content: space-between;
  align-items: center;
  color: #000000;
  padding: 10px 15px;
  border-radius: 8px;
  font-weight: 600;
  letter-spacing: 0.5px;
  text-shadow: 0 1px 1px rgba(255, 255, 255, 0.349);
  background: linear-gradient(
    90deg,
    rgba(173, 216, 230, 0.4),
    rgba(255, 182, 193, 0.6)
  );
  backdrop-filter: blur(20px) saturate(125%);
  -webkit-backdrop-filter: blur(20px) saturate(125%);
  border: 1px solid rgba(255, 255, 255, 0.9);
  box-shadow: 0 8px 32px 0 rgba(0, 0, 0, 0.15);
  cursor: default;
  transition: all 0.3s ease;
  position: relative;
  z-index: 1;
}

.setting-option:hover {
  background: linear-gradient(
    90deg,
    rgba(173, 216, 230, 0.45),
    rgba(255, 182, 193, 0.65)
  );
  color: #000000;
  text-shadow: 0 1px 2px rgba(255, 255, 255, 0.1);
}

.setting-option::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 8px;
  opacity: 0;
  transition: opacity 0.3s ease;
  pointer-events: none;
  z-index: 0;
}

.setting-option:hover::before {
  opacity: 1;
}

.option-label {
  flex: 0 0 auto;
  margin-right: auto;
  text-align: left;
  display: flex;
  align-items: center;
}

.option-control {
  flex: 0 0 auto;
  display: flex;
  align-items: center;
  gap: 8px;
}

/* 更新光标样式 */
.download-link {
  font-size: 0.9em;
  color: #000000;
  word-break: break-all;
  white-space: normal;
  z-index: 2;
  cursor: pointer; /* 添加指针光标 */
}

/* 自适应宽度的下拉菜单 */
.adaptive-select {
  min-width: 120px; /* 根据需要调整 */
  max-width: 250px; /* 根据需要调整 */
  width: auto; /* 自适应宽度 */
}

/* 调整 el-select 内部输入框样式 */
.adaptive-select .el-input__inner {
  border: none;
  background: transparent;
}

/* Hover 状态 */
.adaptive-select:hover .el-input__inner {
  background: rgba(255, 255, 255, 0.1);
}

.adaptive-select .el-input__inner:focus {
  background: rgba(255, 255, 255, 0.1);
  outline: none;
}

.setting-option input[type="text"],
.setting-option textarea {
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 4px;
  background: rgba(255, 255, 255, 0.6);
  transition: background 0.3s ease;
}
.setting-option input[type="text"]:focus,
.setting-option textarea:focus {
  background: rgba(255, 255, 255, 0.1);
  outline: none;
}

@media (prefers-color-scheme: dark) {
  .setting-option {
    color: rgba(255, 255, 255, 0);
    text-shadow: 0 2px 4px rgba(0, 0, 0, 0.5);
    background: linear-gradient(
      90deg,
      rgba(255, 255, 255, 0.4),
      rgba(255, 255, 255, 0.65)
    );
  }


  .setting-option input[type="text"],
  .setting-option textarea {
    background: rgba(255, 255, 255, 0.5);
    color: #000;
  }

  .setting-option input[type="text"]:focus,
  .setting-option textarea:focus {
    background: rgba(255, 255, 255, 0.7);
  }

  .download-link {
    color: #000000;
  }

  .adaptive-select {
    background: rgba(255, 255, 255, 0.5);
    color: #000;
  }

  .adaptive-select:hover .el-input__inner {
    background: rgba(255, 255, 255, 0.3);
  }

  .adaptive-select .el-input__inner:focus {
    background: rgba(255, 255, 255, 0.3);
  }
}
</style>