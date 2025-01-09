<template>
  <div class="allthepast-container">
    <h1>所有学期已修课程</h1>
    
    <!-- 登录表单 -->
    <div class="login-form">
      <el-input
        v-model="loginForm.username"
        placeholder="请输入账号"
        class="input-item"
      />
      <el-input
        v-model="loginForm.password"
        type="password"
        placeholder="请输入密码"
        show-password
        class="input-item"
      />
      <el-input
        v-model="loginForm.ocrAPIURL"
        placeholder="请输入验证码识别API地址"
        class="input-item"
      />
      <el-button 
        type="primary" 
        @click="handleLogin"
        :loading="loading"
      >
        登录教务系统
      </el-button>
    </div>

    <!-- 登录结果展示 -->
    <div v-if="loginResult" class="result-container">
      <h3>登录结果：</h3>
      <div class="result-content">
        <pre>{{ JSON.stringify(loginResult, null, 2) }}</pre>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import { useGlobalStore } from '../../stores/globalStore'

const loading = ref(false)
const loginResult = ref(null)
const globalStore = useGlobalStore()

const loginForm = reactive({
  username: globalStore.userAccount,
  password: globalStore.userPassword,
  ocrAPIURL: globalStore.ocrAPIURL
})

// 处理登录
const handleLogin = async () => {
  if (!loginForm.username || !loginForm.password || !loginForm.ocrAPIURL) {
    ElMessage.warning('请填写完整的登录信息')
    return
  }

  loading.value = true
  try {
    const result = await window.go.login.App.LoginToJwgl(
      loginForm.username,
      loginForm.password,
      loginForm.ocrAPIURL
    )
    loginResult.value = result
    ElMessage.success('登录成功')
  } catch (error) {
    ElMessage.error('登录失败：' + error.message)
    loginResult.value = null
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.allthepast-container {
  padding: 20px;
  max-width: 800px;
  margin: 0 auto;
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 16px;
  margin: 20px 0;
  padding: 20px;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border-radius: 12px;
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.input-item {
  width: 100%;
}

.result-container {
  margin-top: 20px;
  padding: 20px;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border-radius: 12px;
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.result-content {
  background: rgba(0, 0, 0, 0.05);
  padding: 16px;
  border-radius: 8px;
  overflow-x: auto;
}

pre {
  margin: 0;
  white-space: pre-wrap;
  word-wrap: break-word;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .login-form,
  .result-container {
    background: rgba(255, 255, 255, 0.05);
  }

  .result-content {
    background: rgba(255, 255, 255, 0.05);
  }
}
</style>