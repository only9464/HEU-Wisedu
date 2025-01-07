<template>
  <div class="main-container">
    <h1>HEU-Wisedu首页</h1>
    
    <div class="login-container" v-if="!globalStore.isLoggedIn">
      <el-form :model="loginForm" class="login-form">
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
          <el-button
            type="primary"
            @click="handleLogin"
            :loading="loading"
            style="width: 100%"
          >
            登录
          </el-button>
        </el-form-item>
      </el-form>
    </div>
    
    <div v-else class="welcome-container">
      <h2>欢迎回来</h2>
      <p>当前登录账号：{{ globalStore.userAccount }}</p>
      <el-button @click="handleLogout" type="warning">退出登录</el-button>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useGlobalStore } from '../stores/globalStore'
import { ElMessage } from 'element-plus'

const globalStore = useGlobalStore()
const loading = ref(false)

const loginForm = reactive({
  account: '',
  password: ''
})

// 处理登录
const handleLogin = async () => {
  if (!loginForm.account || !loginForm.password) {
    ElMessage.warning('请输入账号和密码')
    return
  }
  
  loading.value = true
  try {
    // 这里可以添加实际的登录验证逻辑
    // await validateCredentials(loginForm.account, loginForm.password)
    
    // 保存账号密码到全局状态
    globalStore.setUserCredentials(loginForm.account, loginForm.password)
    ElMessage.success('登录成功')
  } catch (error) {
    ElMessage.error('登录失败：' + error.message)
  } finally {
    loading.value = false
  }
}

// 处理退出登录
const handleLogout = () => {
  globalStore.setUserCredentials('', '')
  ElMessage.success('已退出登录')
}
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
</style>
