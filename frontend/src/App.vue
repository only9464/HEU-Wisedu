<script setup>
import { ref, onMounted, computed } from 'vue'
import Sidebar from './components/nav/Sidebar.vue'

// 授权相关状态
const authToken = ref('')
const isAuthorized = ref(true)
const showAuthDialog = ref(false)
//const isAuthorized = ref(false)
// const showAuthDialog = ref(false)
const authInput = ref('')
const authError = ref('')
const isVerifying = ref(false)

// Cloudflare Worker URL
const WORKER_URL = 'https://heu.qzz.io'

// 检查本地存储的授权码
const checkLocalAuth = () => {
  const token = localStorage.getItem('auth_token')
  if (token) {
    authToken.value = token
    // 验证本地token是否有效
    verifyToken(token)
  } else {
    showAuthDialog.value = true
  }
}

// 验证授权码
const verifyToken = async (token) => {
  isVerifying.value = true
  authError.value = ''
  
  try {
    const response = await fetch(`${WORKER_URL}/verify`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ authCode: token })
    })
    
    if (response.ok) {
      const data = await response.json()
      if (data.valid) {
        isAuthorized.value = true
        authToken.value = token
        localStorage.setItem('auth_token', token)
        showAuthDialog.value = false
      } else {
        authError.value = '授权码无效或已过期'
        localStorage.removeItem('auth_token')
      }
    } else {
      authError.value = '验证服务暂时不可用'
    }
  } catch (error) {
    authError.value = '网络错误，请检查连接'
    console.error('验证授权码时出错:', error)
  } finally {
    isVerifying.value = false
  }
}

// 提交授权码
const submitAuthCode = () => {
  if (!authInput.value.trim()) {
    authError.value = '请输入授权码'
    return
  }
  verifyToken(authInput.value.trim())
}

// 打开获取授权码页面
const getAuthCode = () => {
  window.runtime.BrowserOpenURL(`${WORKER_URL}/login`, '_blank')
}
const getWatt = () => {
  window.runtime.BrowserOpenURL(`https://steampp.net/`, '_blank')
}

// 组件挂载时检查授权状态
onMounted(() => {
  // checkLocalAuth()
})
</script>

<template>
  <div v-if="isAuthorized" class="app-container">
    <Sidebar />
    <div class="main-content">
      <div class="content">
        <div class="content-wrapper">
          <router-view></router-view>
        </div>
      </div>
    </div>
  </div>

  <!-- 授权对话框 -->
  <div v-if="showAuthDialog" class="auth-dialog-overlay">
    <div class="auth-dialog">
      <h2>授权</h2>
      <p>本软件完全免费，如果你是花钱的，呵呵......为了防止倒卖，我们需要您提供一个授权码（免费的）。请按照以下步骤操作：</p>
      
      <div class="auth-instructions">
        <ol>
          <li>点击下方按钮在浏览器中打开授权页面</li>
          <li>使用GitHub账号登录(没有的话就去注册一个)</li>
          <li>确保已给我们的仓库点Star(为了让你知道免费的从哪里下载)</li>
          <li>复制获得的授权码</li>
          <li>在此处粘贴授权码并验证</li>
        </ol>
      </div>
      
      <button class="get-auth-btn" @click="getAuthCode">
        获取授权码
      </button>
      <button class="get-watt-btn" @click="getWatt">
        Github打不开？点我！
      </button>
      <div class="auth-input-section">
        <input 
          v-model="authInput" 
          type="text" 
          placeholder="将授权码粘贴到这里，然后点“验证”按钮" 
          :disabled="isVerifying"
          @keyup.enter="submitAuthCode"
        >
        <button 
          @click="submitAuthCode" 
          :disabled="isVerifying || !authInput"
          class="verify-btn"
        >
          {{ isVerifying ? '验证中...' : '验证' }}
        </button>
      </div>
      
      <div v-if="authError" class="auth-error">
        {{ authError }}
      </div>
    </div>
  </div>
</template>

<style scoped>
.app-container {
  display: flex;
  width: 100%;
  height: 100vh;
  overflow: hidden;
}

.main-content {
  flex: 1;
  margin-left: 240px; /* 调整左侧间距，200px + 40px */
  min-width: 0;
  padding: 20px 20px 20px 0; /* 上、右、下、左的内边距 */
}

.content {
  height: 100%;
  overflow-y: auto;
  width: 100%;
  padding-right: 20px; /* 右侧内边距 */
}

.content-wrapper {
  max-width: 1200px;
  margin: 0 auto;
  height: 100%;
  display: flex;
  flex-direction: column;
}

/* 授权对话框样式 */
.auth-dialog-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.7);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.auth-dialog {
  background-color: #2d3748;
  border-radius: 12px;
  padding: 30px;
  width: 90%;
  max-width: 500px;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.5);
  color: white;
}

.auth-dialog h2 {
  margin-top: 0;
  color: #63b3ed;
  text-align: center;
}

.auth-instructions {
  margin: 20px 0;
  background-color: rgba(255, 255, 255, 0.05);
  padding: 15px;
  border-radius: 8px;
}

.auth-instructions ol {
  padding-left: 20px;
  margin: 0;
}

.auth-instructions li {
  margin-bottom: 8px;
  line-height: 1.4;
}

.get-auth-btn {
  width: 100%;
  padding: 12px;
  background-color: #0ba543;
  color: white;
  border: none;
  border-radius: 6px;
  font-size: 16px;
  font-weight: bold;
  cursor: pointer;
  transition: background-color 0.2s;
  margin-bottom: 20px;
}
.get-watt-btn {
  width: 100%;
  padding: 12px;
  background-color: #246ac5;
  color: white;
  border: none;
  border-radius: 6px;
  font-size: 16px;
  font-weight: bold;
  cursor: pointer;
  transition: background-color 0.2s;
  margin-bottom: 20px;
}

.get-auth-btn:hover {
  background-color: #31ce31;
}

.auth-input-section {
  display: flex;
  gap: 10px;
  margin-bottom: 15px;
}

.auth-input-section input {
  flex: 1;
  padding: 12px;
  border: 1px solid #4a5568;
  border-radius: 6px;
  background-color: #1a202c;
  color: white;
  font-size: 16px;
}

.auth-input-section input:focus {
  outline: none;
  border-color: #4299e1;
}

.verify-btn {
  padding: 12px 20px;
  background-color: #48bb78;
  color: white;
  border: none;
  border-radius: 6px;
  font-weight: bold;
  cursor: pointer;
  transition: background-color 0.2s;
}

.verify-btn:hover:not(:disabled) {
  background-color: #38a169;
}

.verify-btn:disabled {
  background-color: #4a5568;
  cursor: not-allowed;
}

.auth-error {
  color: #fc8181;
  text-align: center;
  font-size: 14px;
}

/* 自定义滚动条样式 */
.content::-webkit-scrollbar {
  width: 8px;
}

.content::-webkit-scrollbar-track {
  background: transparent;
}

.content::-webkit-scrollbar-thumb {
  background: rgba(255, 255, 255, 0);
  border-radius: 4px;
}

.content::-webkit-scrollbar-thumb:hover {
  background: rgba(255, 255, 255, 0);
}

/* 响应式设计 */
@media (max-width: 768px) {
  .auth-dialog {
    width: 95%;
    padding: 20px;
  }
  
  .auth-input-section {
    flex-direction: column;
  }
}
</style>