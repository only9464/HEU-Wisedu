import { defineStore } from 'pinia'
import { useCourseStore } from './courseStore'

export const useGlobalStore = defineStore('global', {
  state: () => ({
    // 从localStorage读取初始值，如果没有则默认为false
    isDarkMode: localStorage.getItem('isDarkMode') === 'true' || false,
    // 添加账号密码字段，从localStorage读取保存的值
    userAccount: localStorage.getItem('userAccount') || '',
    userPassword: localStorage.getItem('userPassword') || '',
    // 添加用户信息字段
    userInfo: localStorage.getItem('userInfo') ? JSON.parse(localStorage.getItem('userInfo')) : null,
    // 当前选中的批次代码
    batchId: localStorage.getItem('batchId') || '',
    Authorization: localStorage.getItem('Authorization') || '',
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
    // 添加设置用户信息的方法
    setUserInfo(info) {
      this.userInfo = info
      this.Authorization = info.token
      localStorage.setItem('userInfo', JSON.stringify(info))
      localStorage.setItem('Authorization', info.token)
    },
    // 清除用户信息
    clearUserInfo() {
      this.userInfo = null
      this.batchId = ''
      this.Authorization = ''
      localStorage.removeItem('userInfo')
      localStorage.removeItem('batchId')
      localStorage.removeItem('Authorization')
      // 清除课程数据
      const courseStore = useCourseStore()
      courseStore.clearAllData()
    },
    // 设置选中的批次代码
    setBatchId(code) {
      this.batchId = code
      localStorage.setItem('batchId', code)
    },
    // 初始化选中批次（登录后自动选择第一个批次）
    initSelectedBatch() {
      const batches = this.userInfo?.student?.electiveBatchList || []
      if (batches.length > 0) {
        this.setBatchId(batches[0].code)
      }
    }
  },
  
  getters: {
    // 计算属性
    currentTheme: (state) => state.isDarkMode ? 'dark' : 'light',
    // 判断是否已登录
    isLoggedIn: (state) => !!(state.userInfo),
    // 获取用户姓名
    userName: (state) => state.userInfo?.student?.XM || '',
    // 正确获取选课批次列表，并只提取需要的字段
    electiveBatchList: (state) => {
      const list = state.userInfo?.student?.electiveBatchList || []
      return list.map(batch => ({
        code: batch.code,
        name: batch.name,
        beginTime: batch.beginTime,
        endTime: batch.endTime,
        typeName: batch.typeName,
        schoolTermName: batch.schoolTermName,
        canSelect: batch.canSelect
      }))
    },
    // 其他getter...
  }
}) 