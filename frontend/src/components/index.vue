<template>
  <div class="main-container">
    <div class="login-container">
      <!-- 公共账号密码输入区域 -->
      <div class="common-login-section">
        <h2>统一身份认证</h2>
        <el-form class="login-form">
          <el-form-item>
            <label for="account">学号</label>
            <el-input
              id="account"
              v-model="loginForm.account"
              placeholder="请输入账号"
              prefix-icon="User"
            />
          </el-form-item>
          
          <el-form-item>
            <label for="password">密码</label>
            <el-input
              id="password"
              v-model="loginForm.password"
              type="password"
              placeholder="请输入密码"
              prefix-icon="Lock"
              show-password
            />
          </el-form-item>
        </el-form>
      </div>

      <!-- 系统登录卡片区域 -->
      <div class="systems-container">
        <!-- 选课系统登录卡片 -->
        <div class="system-card">
          <div class="card-header">
            <h3>选课系统登录</h3>
            <template v-if="globalStore.isLoggedIn">
              <el-tag type="success" class="status-tag">
                <el-icon class="status-icon"><CircleCheckFilled /></el-icon>
                {{ globalStore.userName }}
              </el-tag>
            </template>
          </div>
          <!-- 未登录状态显示验证码和登录按钮 -->
          <el-form v-if="!globalStore.isLoggedIn" class="system-form">
            <el-form-item>
              <div class="captcha-container">
                <el-input
                  v-model="loginForm.captcha"
                  placeholder="请输入验证码"
                  style="width: 60%"
                />
                <div class="captcha-image" @click="refreshXKCaptcha">
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

          <!-- 已登录状态显示操作按钮 -->
          <div v-else class="button-group">
            <el-button
              type="primary"
              size="default"
              @click="verifyToken"
              :loading="verifyLoading"
            >
              验证登录状态
            </el-button>
            <el-button
              type="danger"
              size="default"
              @click="handleLogout"
            >
              退出登录
            </el-button>
          </div>
        </div>

        <!-- 教务系统登录卡片 -->
        <div class="system-card">
          <div class="card-header">
            <h3>教务系统登录</h3>
            <template v-if="jwglCookies">
              <el-tag type="success" class="status-tag">
                <el-icon class="status-icon"><CircleCheckFilled /></el-icon>
                已登录
              </el-tag>
            </template>
          </div>

          <!-- 未登录状态显示验证码和登录按钮 -->
          <el-form v-if="!jwglCookies" class="system-form">
            <div class="captcha-group">
              <el-form-item>
                <el-input
                  v-model="jwglForm.captcha1"
                  placeholder="验证码1"
                  maxlength="4"
                />
                <img
                  :src="captcha1.image"
                  alt="验证码1"
                  class="captcha-image"
                  @click="refreshJWGLCaptcha(1)"
                />
              </el-form-item>
              <el-form-item>
                <el-input
                  v-model="jwglForm.captcha2"
                  placeholder="验证码2"
                  maxlength="4"
                />
                <img
                  :src="captcha2.image"
                  alt="验证码2"
                  class="captcha-image"
                  @click="refreshJWGLCaptcha(2)"
                />
              </el-form-item>
            </div>
            <el-form-item>
              <el-button
                type="primary"
                :loading="jwglLoading"
                @click="handleJWGLLogin"
                style="width: 100%;"
              >
                登录教务系统
              </el-button>
            </el-form-item>
          </el-form>

          <!-- 已登录状态显示操作按钮 -->
          <div v-else class="button-group">
            <el-button
              type="danger"
              size="default"
              @click="clearJwglLogin"
            >
              清除Cookies
            </el-button>
            <el-button
              type="primary"
              size="default"
              @click="showCookiesDetail"
            >
              查看Cookies
            </el-button>
          </div>
        </div>
      </div>
    </div>
    
    <!-- 登录成功后显示的内容 -->
    <div v-if="globalStore.isLoggedIn" class="welcome-container">
      <div class="user-info-card">
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
          <span class="label">Authorization：</span>
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
              <span class="batch-code">batchId：{{ batch.code }}</span>
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

    <!-- cookies详情弹窗 -->
    <el-dialog
      v-model="cookiesDialogVisible"
      title="教务系统Cookies"
      width="600px"
      :modal-append-to-body="true"
      :append-to-body="true"
    >
      <div class="cookies-detail-content">
        <pre>{{ jwglCookies }}</pre>
      </div>
      <template #footer>
        <div class="dialog-footer">
          <el-button
            type="primary"
            @click="copyCookies"
            style="margin-right: auto;"
          >
            复制
          </el-button>
          <el-button
            type="success"
            @click="cookiesDialogVisible = false"
          >
            确定
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, watch } from 'vue'
import { useGlobalStore } from '../stores/globalStore'
import { ElMessage, ElMessageBox } from 'element-plus'
import { User, Lock } from '@element-plus/icons-vue'
import CryptoJS from 'crypto-js'
import { useRouter, useRoute } from 'vue-router'
import { Check, CircleCheckFilled } from '@element-plus/icons-vue'

const globalStore = useGlobalStore()
const router = useRouter()
const route = useRoute()
const loading = ref(false)
const verifyLoading = ref(false)
const captchaImage = ref('')
const captchaUuid = ref('')

// 监听路由变化
watch(() => route.path, (newPath, oldPath) => {
  // 如果是从选课助手页面跳转回来的
  if (oldPath === '/only9464' && newPath === '/') {
    // 自动执行验证
    verifyToken()
  }
}, { immediate: true })

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

// 修改教务系统登录相关的响应式变量
const jwglForm = reactive({
  account: loginForm.account, // 使用选课系统的账号
  password: loginForm.password, // 使用选课系统的密码
  captcha1: '',
  captcha2: '',
  captchaImg1: '',
  captchaImg2: '',
  token1: '',
  token2: ''
})

// 添加watch来同步两个表单的账号密码
watch(() => loginForm.account, (newVal) => {
  jwglForm.account = newVal
})

watch(() => loginForm.password, (newVal) => {
  jwglForm.password = newVal
})

watch(() => jwglForm.account, (newVal) => {
  loginForm.account = newVal
})

watch(() => jwglForm.password, (newVal) => {
  loginForm.password = newVal
})

// 选课系统的验证码刷新函数
const refreshXKCaptcha = async () => {
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

// 教务系统的验证码刷新函数
const refreshJWGLCaptcha = async (index) => {
  try {
    const result = await window.go.login.App.GetCaptcha()
    if (index === 1) {
      captcha1.image = 'data:image/png;base64,' + result.img
      captcha1.token = result.token
      jwglForm.captcha1 = ''
    } else {
      captcha2.image = 'data:image/png;base64,' + result.img
      captcha2.token = result.token
      jwglForm.captcha2 = ''
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
      // 保存用户信息到全局状态（这个方法会同时设置 token）
      globalStore.setUserInfo(loginResult.data)
      // 初始化选中批次
      globalStore.initSelectedBatch()
      
      ElMessage.success(`欢迎回来，${loginResult.data.student.XM}`)
    } else {
      throw new Error(`登录失败 [${loginResult.code}]: ${loginResult.msg || '未知错误'}`)
    }
  } catch (error) {
    ElMessage.error(error.message)
    refreshXKCaptcha()
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
  refreshXKCaptcha()
  ElMessage.success('已退出登录')
}

// 验证Token的方法
const verifyToken = async () => {
  verifyLoading.value = true
  try {
    ElMessage.info('正在验证Authorization...')
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
    refreshXKCaptcha()
  } finally {
    verifyLoading.value = false
  }
}

// 复制Token功能
const copyToken = () => {
  navigator.clipboard.writeText(globalStore.Authorization)
    .then(() => {
      ElMessage.success('Authorization已复制到剪贴板')
    })
    .catch(() => {
      ElMessage.error('复制失败')
    })
}

// 添加教务系统登录相关的响应式变量
const jwglLoading = ref(false)
const jwglCookies = ref(null)

// 处理教务系统登录
const handleJWGLLogin = async () => {
  if (!jwglForm.account || !jwglForm.password || !jwglForm.captcha1 || !jwglForm.captcha2) {
    ElMessage.warning('请填写完整的登录信息和验证码')
    return
  }

  jwglLoading.value = true
  try {
    const result = await window.go.login.App.LoginToJwgl(
      jwglForm.account,
      jwglForm.password,
      jwglForm.captcha1,
      captcha1.token,
      jwglForm.captcha2,
      captcha2.token
    )
    
    // 将 cookies 对象转换为简单的 name=value 格式
    const cookieStrings = Object.entries(result).map(([name, cookie]) => {
      return `${cookie.Name}=${cookie.Value}`
    })
    
    jwglCookies.value = cookieStrings.join('; ')
    
    // 保存cookies到localStorage，使用gradeCookies作为key以保持一致性
    localStorage.setItem('gradeCookies', JSON.stringify(result))
    
    ElMessage.success('教务系统登录成功，已保存登录状态')
  } catch (error) {
    ElMessage.error(error.message)
    // 刷新验证码
    refreshJWGLCaptcha(1)
    refreshJWGLCaptcha(2)
  } finally {
    jwglLoading.value = false
  }
}

// 复制cookies
const copyCookies = () => {
  if (jwglCookies.value) {
    navigator.clipboard.writeText(jwglCookies.value)
      .then(() => {
        ElMessage.success('Cookies已复制到剪贴板')
      })
      .catch(() => {
        ElMessage.error('复制失败')
      })
  }
}

// 添加清除教务系统登录状态的方法
const clearJwglLogin = () => {
  ElMessageBox.confirm(
    '确定要清除教务系统的登录状态吗？',
    '警告',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    }
  ).then(() => {
    jwglCookies.value = null
    localStorage.removeItem('gradeCookies')
    ElMessage.success('已清除教务系统登录状态')
    // 重新获取验证码
    refreshJWGLCaptcha(1)
    refreshJWGLCaptcha(2)
  }).catch(() => {})
}

// 组件加载时获取验证码和账号密码
onMounted(() => {
  // 从localStorage获取保存的账号密码
  const savedAccount = localStorage.getItem('userAccount')
  const savedPassword = localStorage.getItem('userPassword')
  
  if (savedAccount) {
    loginForm.account = savedAccount
    jwglForm.account = savedAccount
  }
  
  if (savedPassword) {
    loginForm.password = savedPassword
    jwglForm.password = savedPassword
  }

  // 从localStorage获取保存的教务系统cookies
  const savedCookies = localStorage.getItem('gradeCookies')
  if (savedCookies) {
    const cookiesObj = JSON.parse(savedCookies)
    // 将cookies对象转换为字符串格式
    const cookieStrings = Object.entries(cookiesObj).map(([name, cookie]) => {
      return `${cookie.Name}=${cookie.Value}`
    })
    jwglCookies.value = cookieStrings.join('; ')
  }

  refreshXKCaptcha() // 获取选课系统验证码
  // 只有在没有保存的cookies时才获取验证码
  if (!savedCookies) {
    refreshJWGLCaptcha(1) // 获取教务系统第一个验证码
    refreshJWGLCaptcha(2) // 获取教务系统第二个验证码
  }
})

const cookiesDialogVisible = ref(false)

// 添加显示cookies详情的方法
const showCookiesDetail = () => {
  cookiesDialogVisible.value = true
}

// 验证码相关的响应式变量
const captcha1 = reactive({
  image: '',
  token: ''
})

const captcha2 = reactive({
  image: '',
  token: ''
})
</script>

<style scoped>
.main-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;
  gap: 40px; /* 增加登录区域和欢迎区域之间的间距 */
}

.login-container {
  width: 100%;
  max-width: 800px;
}

.login-container h2 {
  margin-bottom: 20px;
  text-align: center;
  font-size: 1.5em;
  color: var(--el-text-color-primary);
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
  width: 100%;
  max-width: 1000px;
  margin-top: 0; /* 移除原有的顶部边距 */
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

.captcha-group {
  display: flex;
  flex-direction: column;
  gap: 15px;
  margin-bottom: 20px;
}

.captcha-group .el-form-item {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 10px;
}

.captcha-group .el-input {
  flex: 1;
}

.captcha-image {
  width: 100px;
  height: 32px;
  cursor: pointer;
  border-radius: 4px;
  transition: all 0.3s;
}

.captcha-image:hover {
  opacity: 0.8;
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

.jwgl-login-container {
  width: 100%;
  max-width: 400px;
  margin-top: 40px;
}

.jwgl-login-container h2 {
  margin-bottom: 20px;
  text-align: center;
}

.cookies-display {
  margin-top: 20px;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 8px;
  padding: 16px;
}

.cookies-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.cookies-header h3 {
  margin: 0;
}

.cookies-content {
  background: rgba(0, 0, 0, 0.05);
  border-radius: 4px;
  padding: 12px;
  max-height: 300px;
  overflow: auto;
}

.cookies-content pre {
  margin: 0;
  white-space: pre-wrap;
  word-break: break-all;
  font-family: monospace;
  line-height: 1.5;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .cookies-content {
    background: rgba(0, 0, 0, 0.2);
  }
}

.common-login-section {
  margin-bottom: 30px;
}

.common-login-section h2 {
  margin-bottom: 20px;
  text-align: center;
  font-size: 1.5em;
  color: var(--el-text-color-primary);
}

.systems-container {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  gap: 20px;
  margin-top: 20px;
}

.system-card {
  background: rgba(255, 255, 255, 0.2);
  border-radius: 16px;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.1);
  padding: 20px;
}

.system-card h3 {
  font-size: 1.4em;
  color: var(--el-text-color-primary);
  margin-bottom: 15px;
}

.system-form {
  margin-top: 10px;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .system-card {
    background: linear-gradient(
      135deg,
      rgba(255, 255, 255, 0.08),
      rgba(255, 255, 255, 0.03)
    );
  }
  
  .login-form {
    background: linear-gradient(
      135deg,
      rgba(255, 255, 255, 0.08),
      rgba(255, 255, 255, 0.03)
    );
  }
}

.login-success-container {
  padding: 15px;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 8px;
}

.status-section {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 15px;
}

.status-tag {
  display: flex;
  align-items: center;
  gap: 5px;
  padding: 8px 12px;
}

.status-icon {
  font-size: 16px;
}

.cookies-section {
  background: rgba(0, 0, 0, 0.05);
  border-radius: 6px;
  padding: 12px;
}

.cookies-header {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 10px;
  margin-bottom: 20px;
}

.cookies-header h4 {
  margin: 0;
  font-size: 14px;
  color: var(--el-text-color-regular);
}

.cookies-content {
  background: rgba(0, 0, 0, 0.03);
  border-radius: 4px;
  padding: 8px;
  max-height: 100px;
  overflow: auto;
}

.cookies-content pre {
  margin: 0;
  font-size: 12px;
  line-height: 1.4;
  white-space: pre-wrap;
  word-break: break-all;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .login-success-container {
    background: rgba(255, 255, 255, 0.08);
  }
  
  .cookies-section {
    background: rgba(0, 0, 0, 0.2);
  }
  
  .cookies-content {
    background: rgba(0, 0, 0, 0.3);
  }
}

.action-buttons {
  display: flex;
  gap: 8px;
}

.cookies-detail-content {
  background: var(--el-fill-color-light);
  border-radius: 8px;
  padding: 16px;
  max-height: 400px;
  overflow-y: auto;
}

.cookies-detail-content pre {
  margin: 0;
  white-space: pre-wrap;
  word-break: break-all;
  font-family: monospace;
  color: var(--el-text-color-regular);
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  padding-top: 20px;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .cookies-detail-content {
    background: var(--el-fill-color-dark);
  }
}

.card-header {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 10px;
  margin-bottom: 20px;
}

.card-header h3 {
  margin: 0;
  font-size: 1.3em;
  color: var(--el-text-color-primary);
}

.status-tag {
  display: flex;
  align-items: center;
  gap: 5px;
  padding: 8px 12px;
}

.status-icon {
  font-size: 16px;
}

.button-group {
  display: flex;
  gap: 12px;
  justify-content: center;
  margin-top: 15px;
}

.login-form .el-form-item {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 20px;
}

.login-form label {
  min-width: 40px;
  text-align: left;
  color: var(--el-text-color-secondary);
}

.login-form .el-input {
  flex: 1;
}

.el-button {
  border-radius: 8px;
}
</style>
