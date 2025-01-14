import { defineStore } from 'pinia'

export const useTJKCStore = defineStore('tjkc', {
  state: () => ({
    hideSelected: false
  }),
  
  actions: {
    setHideSelected(value) {
      this.hideSelected = value
    },
    resetFilters() {
      this.hideSelected = false
    }
  }
}) 