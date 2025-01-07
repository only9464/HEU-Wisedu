import { createRouter, createWebHistory } from 'vue-router'
import FirstView from '../views/FirstView.vue'
import SecondView from '../views/SecondView.vue'
import IndexView from '../components/index.vue'
import SettingView from '../components/setting.vue'
const routes = [
  { path: '/', component: IndexView }, // 首页
  { path: '/First', component: FirstView }, // 第一个页面
  { path: '/second', component: SecondView }, // 第二个页面
  { path: '/setting', component: SettingView }, // 设置页面
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router