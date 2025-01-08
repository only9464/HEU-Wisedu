import { defineStore } from 'pinia'

export const useCourseStore = defineStore('course', {
  state: () => ({
    tjkcData: null,
    fawkcData: null,
    xgkcData: null,
    // 记录是否已经加载过数据
    tjkcLoaded: false,
    fawkcLoaded: false,
    xgkcLoaded: false,
    // 从localStorage读取任务队列，如果没有则为空数组
    taskQueue: JSON.parse(localStorage.getItem('taskQueue') || '[]')
  }),
  
  actions: {
    setTJKCData(data) {
      this.tjkcData = data
      this.tjkcLoaded = true
    },
    setFAWKCData(data) {
      this.fawkcData = data
      this.fawkcLoaded = true
    },
    setXGKCData(data) {
      this.xgkcData = data
      this.xgkcLoaded = true
    },
    // 清除所有数据（比如退出登录时调用）
    clearAllData() {
      this.tjkcData = null
      this.fawkcData = null
      this.xgkcData = null
      this.tjkcLoaded = false
      this.fawkcLoaded = false
      this.xgkcLoaded = false
      // 清除任务队列时同时清除localStorage中的数据
      this.taskQueue = []
      localStorage.removeItem('taskQueue')
    },
    // 添加任务到队列
    addToTaskQueue(task) {
      // 检查是否已存在
      const exists = this.taskQueue.some(t => t.JXBID === task.JXBID)
      if (!exists) {
        // 确保task包含所有必需字段
        const newTask = {
          JXBID: task.JXBID,
          KCM: task.KCM,
          SKJS: task.SKJS,
          XF: task.XF,
          clazzType: task.clazzType,
          secretVal: task.secretVal,
          KCXZ: task.KCXZ
        }
        this.taskQueue.push(newTask)
        // 保存到localStorage
        localStorage.setItem('taskQueue', JSON.stringify(this.taskQueue))
      }
    },
    // 从队列中移除任务
    removeFromTaskQueue(jxbid) {
      const index = this.taskQueue.findIndex(task => task.JXBID === jxbid)
      if (index !== -1) {
        this.taskQueue.splice(index, 1)
        // 更新localStorage
        localStorage.setItem('taskQueue', JSON.stringify(this.taskQueue))
      }
    },
    // 获取任务队列
    getTaskQueue() {
      return this.taskQueue
    },
    // 检查课程是否在队列中
    isInTaskQueue(jxbid) {
      return this.taskQueue.some(task => task.JXBID === jxbid)
    },
    // 添加清空任务队列的方法
    clearTaskQueue() {
      this.taskQueue = []
      localStorage.removeItem('taskQueue')
    }
  }
}) 