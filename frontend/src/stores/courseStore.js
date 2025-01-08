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
    }
  }
}) 