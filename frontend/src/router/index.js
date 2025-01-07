import { createRouter, createWebHistory } from 'vue-router'
import TJKCView from '../views/TJKCView.vue'
import FAWKCView from '../views/FAWKCView.vue'
import XGKCView from '../views/XGKCView.vue'
import YXKCYXView from '../views/YXKCYXView.vue'
import only9464View from '../views/only9464View.vue'
import IndexView from '../components/index.vue'
import SettingView from '../components/setting.vue'
const routes = [
  { path: '/', component: IndexView }, // 首页
  { path: '/TJKC', component: TJKCView }, // 培养方案内课程
  { path: '/FAWKC', component: FAWKCView }, // 跨专业选修课
  { path: '/XGKC', component: XGKCView }, // 公选课
  { path: '/YXKCYX', component: YXKCYXView }, // 已选课程
  { path: '/only9464', component: only9464View }, // 抢课助手
  { path: '/setting', component: SettingView }, // 设置页面
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router