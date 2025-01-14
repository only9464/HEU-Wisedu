import { defineStore } from 'pinia'

export const useFAWKCStore = defineStore('fawkc', {
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