import { defineStore } from 'pinia'

export const useXGKCStore = defineStore('xgkc', {
  state: () => ({
    selectedType: '全部',
    searchKeyword: '',
    onlyNotFull: false,
    onlyOnline: false,
    hideSelected: false
  }),
  
  actions: {
    setSelectedType(type) {
      this.selectedType = type
    },
    setSearchKeyword(keyword) {
      this.searchKeyword = keyword
    },
    setOnlyNotFull(value) {
      this.onlyNotFull = value
    },
    setOnlyOnline(value) {
      this.onlyOnline = value
    },
    setHideSelected(value) {
      this.hideSelected = value
    },
    resetFilters() {
      this.selectedType = '全部'
      this.searchKeyword = ''
      this.onlyNotFull = false
      this.onlyOnline = false
      this.hideSelected = false
    }
  }
}) 