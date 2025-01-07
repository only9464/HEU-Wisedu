import { defineStore } from 'pinia'

export const useGlobalStore = defineStore('global', {
  state: () => ({
    // 从localStorage读取初始值，如果没有则默认为false
    isDarkMode: localStorage.getItem('isDarkMode') === 'true' || false,
    userName: '',
    // 添加账号密码字段，从localStorage读取保存的值
    userAccount: localStorage.getItem('userAccount') || '',
    userPassword: localStorage.getItem('userPassword') || '',
    // 其他全局变量...
  }),
  
  actions: {
    // 修改暗黑模式
    setDarkMode(value) {
      this.isDarkMode = value
      // 将值保存到localStorage
      localStorage.setItem('isDarkMode', value)
    },
    // 设置用户名
    setUserName(name) {
      this.userName = name
    },
    // 设置账号密码
    setUserCredentials(account, password) {
      this.userAccount = account
      this.userPassword = password
      // 保存到localStorage
      localStorage.setItem('userAccount', account)
      localStorage.setItem('userPassword', password)
    },
  },
  
  getters: {
    // 计算属性
    currentTheme: (state) => state.isDarkMode ? 'dark' : 'light',
    // 判断是否已登录
    isLoggedIn: (state) => !!(state.userAccount && state.userPassword),
    // 其他getter...
  }
}) 