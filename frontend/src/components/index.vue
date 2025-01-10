<template>
  <div class="main-container">
    <h1>HEU-Wisedu首页</h1>
    
    <div class="login-container" v-if="!globalStore.isLoggedIn">
      <el-form 
        :model="loginForm" 
        class="login-form"
        @keyup.enter="handleLogin"
      >
        <el-form-item>
          <el-input
            v-model="loginForm.account"
            placeholder="请输入账号"
            prefix-icon="User"
          />
        </el-form-item>
        
        <el-form-item>
          <el-input
            v-model="loginForm.password"
            type="password"
            placeholder="请输入密码"
            prefix-icon="Lock"
            show-password
          />
        </el-form-item>
        
        <el-form-item>
          <div class="captcha-container">
            <el-input
              v-model="loginForm.captcha"
              placeholder="请输入验证码"
              style="width: 60%"
            />
            <div class="captcha-image" @click="refreshCaptcha">
              <img v-if="captchaImage" :src="captchaImage" alt="验证码" />
              <el-button v-else type="primary" link>
                获取验证码
              </el-button>
            </div>
          </div>
        </el-form-item>
        
        <el-form-item>
          <el-button
            type="primary"
            @click="handleLogin"
            :loading="loading"
            style="width: 100%"
          >
            登录选课系统
          </el-button>
        </el-form-item>
      </el-form>
    </div>
    
    <div v-else class="welcome-container">
      <div class="user-info-card">
        <div class="header-row">
          <el-button 
            type="primary" 
            @click="verifyToken"
            :loading="verifyLoading"
            :style="{
              padding: '8px 20px',
              fontSize: '14px',
              height: '32px',
              width: '120px'
            }"
          >
            验证登录状态
          </el-button>
          <h2 class="welcome-text">欢迎回来，{{ globalStore.userName }}</h2>
          <el-button 
            @click="handleLogout" 
            type="danger" 
            :style="{
              padding: '8px 20px',
              fontSize: '14px',
              height: '32px',
              width: '120px'
            }"
          >
            退出登录
          </el-button>
        </div>
        <div class="info-grid">
          <div class="info-item">
            <span class="label">学号：</span>
            <span>{{ globalStore.userInfo?.student?.XH }}</span>
          </div>
          <div class="info-item">
            <span class="label">年级：</span>
            <span>{{ globalStore.userInfo?.student?.NJ }}级</span>
          </div>
          <div class="info-item">
            <span class="label">学院：</span>
            <span>{{ globalStore.userInfo?.student?.YXMC }}</span>
          </div>
          <div class="info-item">
            <span class="label">专业：</span>
            <span>{{ globalStore.userInfo?.student?.ZYMC }}</span>
          </div>
          <div class="info-item">
            <span class="label">班级：</span>
            <span>{{ globalStore.userInfo?.student?.schoolClassName }}</span>
          </div>
          <div class="info-item">
            <span class="label">校区：</span>
            <span>{{ globalStore.userInfo?.student?.campusName }}</span>
          </div>
        </div>
        <div class="token-container">
          <span class="label">Token：</span>
          <div class="token-scroll">
            <span class="token-text">{{ globalStore.Authorization }}</span>
          </div>
          <el-button 
            type="primary" 
            size="small"
            @click="copyToken"
          >
            复制
          </el-button>
        </div>
      </div>
      
      <!-- 选课批次信息 -->
      <div class="elective-batch-container">
        <h3>可选课程批次</h3>
        <div class="batch-info" v-for="batch in globalStore.electiveBatchList" :key="batch.code" :class="{ 'batch-selected': batch.code === globalStore.batchId }" @click="globalStore.setBatchId(batch.code)">
          <div class="batch-header">
            <div class="term-info">
              <span class="term-name">{{ batch.schoolTermName }}</span>
              <span class="batch-code">批次代码：{{ batch.code }}</span>
            </div>
            <div class="batch-tags">
              <el-tag type="info" size="small">{{ batch.typeName }}</el-tag>
              <el-tag :type="batch.canSelect === '1' ? 'success' : 'warning'" size="small">
                {{ batch.canSelect === '1' ? '可选' : '不可选' }}
              </el-tag>
              <el-tag 
                v-if="batch.code === globalStore.batchId" 
                type="danger" 
                size="small"
              >
                已选中该批次
              </el-tag>
            </div>
          </div>
          <div class="batch-time">
            <div class="time-item">
              <span class="time-label">开始时间：</span>
              <span>{{ batch.beginTime }}</span>
            </div>
            <div class="time-item">
              <span class="time-label">结束时间：</span>
              <span>{{ batch.endTime }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useGlobalStore } from '../stores/globalStore'
import { ElMessage } from 'element-plus'
import CryptoJS from 'crypto-js'
import { useRouter } from 'vue-router'

const globalStore = useGlobalStore()
const router = useRouter()
const loading = ref(false)
const verifyLoading = ref(false)
const captchaImage = ref('')
const captchaUuid = ref('')

// AES-ECB加密函数
const encryptPassword = (plainText) => {
  // 内置密钥（16 字节）
  const key = CryptoJS.enc.Utf8.parse('MWMqg2tPcDkxcm11')
  
  // 加密
  const encrypted = CryptoJS.AES.encrypt(plainText, key, {
    mode: CryptoJS.mode.ECB,
    padding: CryptoJS.pad.Pkcs7
  })
  
  // 返回Base64编码的密文
  return encrypted.toString()
}

const loginForm = reactive({
  account: localStorage.getItem('userAccount') || '',
  password: localStorage.getItem('userPassword') || '',
  captcha: '',
})

// 获取验证码
const refreshCaptcha = async () => {
  try {
    const response = await fetch('https://jwxk.hrbeu.edu.cn/xsxk/auth/captcha', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    const data = await response.json()
    if (data?.data?.captcha && data?.data?.uuid) {
      captchaImage.value = data.data.captcha
      captchaUuid.value = data.data.uuid
      loginForm.captcha = ''
    } else {
      throw new Error('验证码获取失败')
    }
  } catch (error) {
    console.error('获取验证码失败:', error)
    ElMessage.error('获取验证码失败，请重试')
  }
}

// 发送登录请求
const sendLoginRequest = async (loginData) => {
  const formData = new FormData()
  formData.append('loginname', loginData.username)
  formData.append('password', loginData.password)
  formData.append('captcha', loginData.captcha)
  formData.append('uuid', loginData.uuid)

  const response = await fetch('https://jwxk.hrbeu.edu.cn/xsxk/auth/hrbeu/login', {
    method: 'POST',
    body: formData,
    headers: {
      'Accept': 'application/json'
    }
  })

  if (!response.ok) {
    throw new Error(`登录请求失败: ${response.status}`)
  }

  const data = await response.json()
  if (data.code !== 200) {
    throw new Error(`登录失败 [${data.code}]: ${data.msg || '未知错误'}`)
  }

  return data
}

// 处理登录
const handleLogin = async () => {
  if (!loginForm.account || !loginForm.password || !loginForm.captcha) {
    ElMessage.warning('请输入账号、密码和验证码')
    return
  }
  
  // 在登录请求前就保存账号密码
  globalStore.setUserCredentials(loginForm.account, loginForm.password)
  
  loading.value = true
  try {
    // 加密密码
    const encryptedPassword = encryptPassword(loginForm.password)
    
    const loginData = {
      username: loginForm.account,
      password: encryptedPassword,  // 发送加密后的密码
      captcha: loginForm.captcha,
      uuid: captchaUuid.value
    }
    
    const loginResult = await sendLoginRequest(loginData)

    // 验证登录状态
    if (loginResult.code === 200 && loginResult.msg === '登录成功') {
      // 保存用户信息到全局状态
      globalStore.setUserInfo(loginResult.data)
      // 初始化选中批次
      globalStore.initSelectedBatch()
      
      ElMessage.success(`欢迎回来，${loginResult.data.student.XM}`)
    } else {
      throw new Error(`登录失败 [${loginResult.code}]: ${loginResult.msg || '未知错误'}`)
    }
  } catch (error) {
    ElMessage.error(error.message)
    refreshCaptcha()
  } finally {
    loading.value = false
  }
}

// 处理退出登录
const handleLogout = () => {
  // 不再清除账号密码，只清除用户信息
  globalStore.clearUserInfo()
  captchaImage.value = ''
  loginForm.captcha = ''
  // 刷新验证码
  refreshCaptcha()
  ElMessage.success('已退出登录')
}

// 验证Token的方法
const verifyToken = async () => {
  verifyLoading.value = true
  try {
    ElMessage.info('正在验证Token...')
    if (!globalStore.Authorization || !globalStore.batchId) {
      throw new Error('未找到Authorization信息或批次ID')
    }
    
    const result = await window.go.only9464.App.VerifyToken(
      globalStore.Authorization,
      globalStore.batchId
    )
    
    const data = JSON.parse(result)
    
    if (data.code === 200 && data.msg === 'success') {
      ElMessage.success('登录状态有效')
    } else {
      throw new Error('登录状态已失效')
    }
  } catch (error) {
    ElMessage.error(error.message || '验证失败')
    // 如果验证失败，清除用户信息并刷新验证码
    globalStore.clearUserInfo()
    refreshCaptcha()
  } finally {
    verifyLoading.value = false
  }
}

// 复制Token功能
const copyToken = () => {
  navigator.clipboard.writeText(globalStore.Authorization)
    .then(() => {
      ElMessage.success('Token已复制到剪贴板')
    })
    .catch(() => {
      ElMessage.error('复制失败')
    })
}

// 组件加载时获取验证码
onMounted(() => {
  refreshCaptcha()
})
</script>

<style scoped>
.main-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;
}

.login-container {
  width: 100%;
  max-width: 400px;
  margin-top: 40px;
}

.login-form {
  background: linear-gradient(
    135deg,
    rgba(255, 255, 255, 0.1),
    rgba(255, 255, 255, 0.05)
  );
  backdrop-filter: blur(10px);
  padding: 30px;
  border-radius: 12px;
  box-shadow: 0 8px 32px 0 rgba(0, 0, 0, 0.1);
}

.welcome-container {
  text-align: center;
  margin-top: 40px;
}

.welcome-container h2 {
  margin-bottom: 20px;
}

.welcome-container p {
  margin-bottom: 20px;
  color: #666;
}

@media (prefers-color-scheme: dark) {
  .login-form {
    background: linear-gradient(
      135deg,
      rgba(255, 255, 255, 0.1),
      rgba(255, 255, 255, 0.05)
    );
  }
  
  .welcome-container p {
    color: #999;
  }
}

.captcha-container {
  display: flex;
  align-items: center;
  gap: 10px;
}

.captcha-image {
  flex: 0 0 40%;
  height: 32px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 4px;
  overflow: hidden;
}

.captcha-image img {
  height: 100%;
  width: 100%;
  object-fit: contain;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .captcha-image {
    background: rgba(255, 255, 255, 0.05);
  }
}

.elective-batch-container {
  margin: 20px 0;
  width: 100%;
  max-width: 1000px;
  padding: 0;
}

.batch-info {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border-radius: 8px;
  padding: 15px 20px;
  margin-bottom: 15px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.batch-info:hover {
  background: rgba(255, 255, 255, 0.15);
}

.batch-selected {
  background: rgba(30, 143, 163, 0.15);
  border: 1px solid rgba(30, 143, 163, 0.3);
}

.batch-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}

.term-info {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 4px;
}

.term-name {
  font-size: 1.1em;
  font-weight: 500;
}

.batch-code {
  font-size: 0.9em;
  color: #666;
  font-family: monospace;
}

.batch-tags {
  display: flex;
  gap: 8px;
}

.batch-time {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.user-info-card {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border-radius: 12px;
  padding: 0;
  margin-bottom: 20px;
  padding-top: 15px;
  padding-bottom: 15px;
  width: 100%;
  max-width: 1000px;
}

.info-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 15px;
  padding: 0 20px;
  margin-bottom: 25px;
}

.info-item {
  text-align: left;
  padding: 8px;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 6px;
}

.label {
  color: #666;
  margin-right: 8px;
}

.time-range {
  display: flex;
  flex-direction: column;
  gap: 4px;
  font-size: 0.9em;
}

.time-item {
  display: flex;
  align-items: center;
  margin: 0;
}

.time-label {
  color: #666;
  margin-right: 8px;
  min-width: 45px;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .label {
    color: #999;
  }
  
  .user-info-card {
    background: rgba(255, 255, 255, 0.05);
  }
  
  .info-item {
    background: rgba(255, 255, 255, 0.02);
  }
  
  .time-label {
    color: #999;
  }
}

.token-container {
  display: flex;
  align-items: center;
  padding: 8px;
  margin-left: 20px;
  margin-right: 20px;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 6px;
  gap: 8px;
}

.token-scroll {
  flex: 1;
  overflow-x: auto;
  white-space: nowrap;
  padding: 4px;
  margin-right: 8px;
}

.token-text {
  font-family: monospace;
}

/* 自定义滚动条样式 */
.token-scroll::-webkit-scrollbar {
  height: 4px;
}

.token-scroll::-webkit-scrollbar-track {
  background: rgba(0, 0, 0, 0.1);
  border-radius: 2px;
}

.token-scroll::-webkit-scrollbar-thumb {
  background: rgba(0, 0, 0, 0.2);
  border-radius: 2px;
}

.token-scroll::-webkit-scrollbar-thumb:hover {
  background: rgba(0, 0, 0, 0.3);
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .token-scroll::-webkit-scrollbar-track {
    background: rgba(255, 255, 255, 0.05);
  }
  
  .token-scroll::-webkit-scrollbar-thumb {
    background: rgba(255, 255, 255, 0.1);
  }
  
  .token-scroll::-webkit-scrollbar-thumb:hover {
    background: rgba(255, 255, 255, 0.15);
  }
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .batch-selected {
    background: rgba(30, 143, 163, 0.25);
    border: 1px solid rgba(30, 143, 163, 0.4);
  }
}

.header-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 25px;
  padding: 0 20px;
  height: 50px;
}

.welcome-text {
  margin: 0;
  flex: 1;
  text-align: center;
  font-size: 20px;
  color: #333;
  font-weight: 500;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .welcome-text {
    color: #fff;
  }
}

/* 统一按钮容器样式 */
.header-row .el-button {
  min-width: 100px;
  flex-shrink: 0;
}

.elective-batch-container h3 {
  padding: 0 20px;
  margin-bottom: 15px;
  text-align: left;
}
</style>
