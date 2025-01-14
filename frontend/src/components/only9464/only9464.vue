<template>
  <div class="only9464-container">
    <!-- 任务队列表格 -->
    <div class="task-table-container">
      <div class="table-header">
        <div class="title-section">
          <h2>任务队列 ({{ taskList.length }})</h2>
          <el-button 
            type="primary" 
            :icon="Refresh"
            circle
            @click="handleRefresh"
            :loading="loading"
          />
        </div>
        <div class="header-buttons">
          <el-button 
            type="danger"
            @click="handleClearAll"
          >
            清空队列
          </el-button>
          <div class="mode-select">
            <span class="mode-label">选(qiang)课模式：</span>
            <el-select
              v-model="selectMode"
              style="width: 120px;"
            >
              <el-option
                v-for="item in modeOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </div>
          <div class="interval-input">
            <span class="interval-label">间隔时间(s)：</span>
            <el-input
              v-model="requestInterval"
              type="number"
              style="width: 120px;"
              placeholder="请输入间隔"
              @input="handleIntervalChange"
            />
          </div>
          <el-button 
            :type="isTaskRunning ? 'warning' : 'primary'"
            @click="handleStartTask"
          >
            {{ isTaskRunning ? '停止选(qiang)课' : '开始选(qiang)课任务' }}
          </el-button>
        </div>
      </div>
      
      <el-table 
        v-loading="loading"
        :data="taskList" 
        style="width: 100%"
        :stripe="true"
        class="acrylic-effect"
        @sort-change="handleSortChange"
      >
        <el-table-column 
          type="index" 
          label="序号" 
          width="60" 
          align="center"
          header-align="center"
        />
        <el-table-column 
          prop="KCM" 
          label="课程名称" 
          min-width="160" 
          align="center"
          header-align="center"
          show-overflow-tooltip
          sortable="custom"
        />
        <el-table-column 
          prop="SKJS" 
          label="授课教师" 
          min-width="120" 
          align="center"
          header-align="center"
          sortable="custom"
        >
          <template #default="scope">
            <div class="multi-line-cell">{{ scope.row.SKJS }}</div>
          </template>
        </el-table-column>
        <el-table-column 
          prop="XF" 
          label="学分" 
          width="80" 
          align="center"
          header-align="center"
          sortable="custom"
        />
        <el-table-column 
          prop="KCXZ" 
          label="课程性质" 
          min-width="120" 
          align="center"
          header-align="center"
          sortable="custom"
        >
          <template #default="scope">
            <div class="multi-line-cell">{{ scope.row.KCXZ || '-' }}</div>
          </template>
        </el-table-column>
        <el-table-column 
          prop="clazzType" 
          label="课程类型" 
          width="120" 
          align="center"
          header-align="center"
          sortable="custom"
        >
          <template #default="scope">
            <el-tag :type="getTagType(scope.row.clazzType)">
              {{ getClazzTypeLabel(scope.row.clazzType) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column 
          prop="status" 
          label="选(qiang)课状态" 
          min-width="120"
          align="center"
          header-align="center"
        >
          <template #default="scope">
            <div class="status-info">
              <el-tag 
                :type="getStatusType(scope.row.status)" 
                effect="plain"
                style="white-space: normal; height: auto; padding: 2px 6px; line-height: 1.5;"
              >
                {{ scope.row.message || scope.row.status || '未开始' }}
              </el-tag>
              <span v-if="scope.row.interval" class="interval-text">
                ({{ scope.row.interval.toFixed(1) }}s)
              </span>
            </div>
          </template>
        </el-table-column>
        <el-table-column 
          label="操作" 
          width="120"
          align="center"
          header-align="center"
          fixed="right"
        >
          <template #default="scope">
            <el-button 
              type="danger"
              size="small"
              @click="handleRemoveTask(scope.row)"
            >
              移除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Refresh, InfoFilled } from '@element-plus/icons-vue'
import { useCourseStore } from '../../stores/courseStore'
import { useGlobalStore } from '../../stores/globalStore'
import { useRouter } from 'vue-router'

const router = useRouter()
const loading = ref(false)
const taskList = ref([])
const courseStore = useCourseStore()
const globalStore = useGlobalStore()
const requestInterval = ref('0.275')

// 修改选课模式相关的数据，使用 store
const selectMode = computed({
  get: () => courseStore.getSelectMode(),
  set: (value) => courseStore.setSelectMode(value)
})
const modeOptions = [
  { value: 'normal', label: '老实人模式' },
  { value: 'crazy', label: '狂暴模式' }
]

// 使用计算属性获取任务运行状态
const isTaskRunning = computed(() => courseStore.getTaskRunning())

// 修改排序相关的变量，移除默认值
const sortOrder = ref({
  prop: null,
  order: null
})

// 添加排序处理函数
const handleSortChange = ({ prop, order }) => {
  if (!prop) {
    taskList.value = courseStore.getTaskQueue()
    return
  }

  const list = [...courseStore.getTaskQueue()]
  list.sort((a, b) => {
    let compareResult = 0
    // 处理不同类型的排序
    switch (prop) {
      case 'XF':  // 学分按数字排序
        compareResult = parseFloat(a[prop]) - parseFloat(b[prop])
        break
      case 'clazzType':  // 课程类型按照预定义顺序排序
        const typeOrder = { 'TJKC': 1, 'FAWKC': 2, 'XGKC': 3 }
        compareResult = (typeOrder[a[prop]] || 0) - (typeOrder[b[prop]] || 0)
        break
      default:  // 其他字段按字符串排序
        compareResult = String(a[prop] || '').localeCompare(String(b[prop] || ''), 'zh-CN')
    }
    // 根据排序方向调整结果
    return order === 'ascending' ? compareResult : -compareResult
  })
  taskList.value = list
}

// 获取课程类型标签
const getClazzTypeLabel = (type) => {
  const labels = {
    'XGKC': '公选课',
    'TJKC': '培养方案内课程',
    'FAWKC': '跨专业选修课'
  }
  return labels[type] || type
}

// 获取标签类型
const getTagType = (type) => {
  const types = {
    'XGKC': 'success',
    'TJKC': 'primary',
    'FAWKC': 'warning'
  }
  return types[type] || 'info'
}

// 获取状态标签类型
const getStatusType = (status) => {
  const types = {
    '未开始': 'info',
    '等待中': 'warning',
    '选(qiang)课成功': 'success',
    '课容量已满': 'danger',
    '请求过快': 'warning'
  }
  return types[status] || 'info'
}

// 获取任务列表
const getTaskList = () => {
  loading.value = true
  try {
    const list = courseStore.getTaskQueue()
    // 应用当前的排序
    if (sortOrder.value.prop) {
      handleSortChange(sortOrder.value)
    } else {
      taskList.value = list
    }
  } catch (error) {
    ElMessage.error('获取任务列表失败：' + error.message)
  } finally {
    loading.value = false
  }
}

// 移除任务
const handleRemoveTask = async (task) => {
  try {
    courseStore.removeFromTaskQueue(task.JXBID)
    ElMessage.success('任务已移除')
    getTaskList() // 刷新列表
  } catch (error) {
    ElMessage.error('移除任务失败：' + error.message)
  }
}

// 刷新任务列表
const handleRefresh = () => {
  getTaskList()
}

// 添加一键移除处理函数
const handleClearAll = () => {
  ElMessageBox.confirm(
    '确定要清空任务队列吗？',
    '警告',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    }
  )
    .then(() => {
      courseStore.clearTaskQueue()  // 需要在courseStore中添加这个方法
      getTaskList()
      ElMessage.success('任务队列已清空')
    })
    .catch(() => {
      // 用户点击取消，不做任何操作
    })
}

// 添加间隔时间变化处理函数
const handleIntervalChange = (value) => {
  // 允许输入为空
  if (value === '') {
    requestInterval.value = ''
    return
  }

  // 移除前导零，但保留小数点前的零
  if (value.startsWith('0') && value.length > 1 && value[1] !== '.') {
    requestInterval.value = value.replace(/^0+/, '')
  }

  // 只有当输入是有效的正数时才更新后端
  if (!isNaN(Number(value)) && Number(value) > 0) {
    window.runtime.EventsEmit("only9464:interval", Number(value))
  }
  // 如果是非法值，保持输入框的值不变，不发送到后端
}

// 验证登录状态的函数
const verifyLoginStatus = async () => {
  try {
    if (!globalStore.Authorization || !globalStore.batchId) {
      throw new Error('未找到Authorization信息或批次ID')
    }
    
    const result = await window.go.only9464.App.VerifyToken(
      globalStore.Authorization,
      globalStore.batchId
    )
    
    const data = JSON.parse(result)
    
    if (data.code === 200 && data.msg === 'success') {
      return true
    } else {
      throw new Error('登录状态已失效')
    }
  } catch (error) {
    ElMessage.error('登录失效，请重新登录。即将跳转登录......')
    // 清除用户信息
    globalStore.clearUserInfo()
    // 延迟1秒后跳转到登录页面
    setTimeout(() => {
      router.push('/')
    }, 1000)
    return false
  }
}

// 修改开始选(qiang)课任务处理函数
const handleStartTask = async () => {
  if (isTaskRunning.value) {
    // 如果任务正在运行，则停止任务
    window.runtime.EventsEmit("only9464:stop")
    courseStore.setTaskRunning(false)  // 使用 store 的方法
    ElMessage.info('选(qiang)课已停止')
    // 清理事件监听
    window.runtime.EventsOff("only9464:result")
    window.runtime.EventsOff("only9464:complete")
    return
  }

  // 先验证登录状态
  const isValid = await verifyLoginStatus()
  if (!isValid) {
    return
  }

  if (taskList.value.length === 0) {
    ElMessage.warning('任务队列为空，请先添加课程')
    return
  }

  // 验证间隔时间
  const interval = Number(requestInterval.value)
  if (!interval || interval <= 0) {
    ElMessage.warning('请输入有效的间隔时间')
    requestInterval.value = '0.275'
    return
  }

  // 先清理可能存在的旧事件监听
  window.runtime.EventsOff("only9464:result")
  window.runtime.EventsOff("only9464:complete")

  try {
    loading.value = true
    // 从全局store获取必要的信息
    const Authorization = globalStore.Authorization
    const batchId = globalStore.batchId

    if (!Authorization || !batchId) {
      ElMessage.error('请先登录并选择选课批次')
      return
    }

    // 重置所有课程状态
    taskList.value.forEach(task => {
      task.status = '未开始'
      task.message = ''
    })

    // 准备课程数据
    const courses = taskList.value.map(task => ({
      clazzType: task.clazzType,
      clazzId: task.JXBID,
      secretVal: task.secretVal,
      courseName: task.KCM,
      teacher: task.SKJS
    }))

    // 设置事件监听
    window.runtime.EventsOn("only9464:result", (result) => {
      const task = taskList.value.find(t => t.JXBID === result.clazzId)
      if (task) {
        task.status = result.message
        task.message = result.message
        // 强制更新视图
        taskList.value = [...taskList.value]
      }
    })

    // 监听任务完成事件
    window.runtime.EventsOn("only9464:complete", () => {
      if (courseStore.getTaskRunning()) { // 使用 store 的方法
        ElMessage.success('选(qiang)课任务已完成')
      }
      courseStore.setTaskRunning(false)  // 使用 store 的方法
      window.runtime.EventsOff("only9464:result")
      window.runtime.EventsOff("only9464:complete")
    })

    // 调用批量选(qiang)课接口，传入初始间隔时间
    await window.go.only9464.App.BatchAddClazz(Authorization, batchId, courses, Number(requestInterval.value), selectMode.value)
    courseStore.setTaskRunning(true)  // 使用 store 的方法
    ElMessage.success('选(qiang)课任务已启动')
  } catch (error) {
    console.error('选(qiang)课失败：', error)
    ElMessage.error('选(qiang)课失败：' + error.message)
    courseStore.setTaskRunning(false)  // 使用 store 的方法
    // 发生错误时才清理事件监听
    window.runtime.EventsOff("only9464:result")
    window.runtime.EventsOff("only9464:complete")
  } finally {
    loading.value = false
  }
}

// 组件加载时获取任务列表
onMounted(() => {
  getTaskList()
})
</script>

<style scoped>
.only9464-container {
  height: 100%;
  display: flex;
  flex-direction: column;
  padding: 0 20px;
  gap: 20px;
  overflow: hidden;
}

.task-table-container {
  padding: 20px;
}

.table-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  padding: 0 8px;
}

.title-section {
  display: flex;
  align-items: center;
  gap: 8px;
}

.header-buttons {
  display: flex;
  gap: 12px;
  align-items: center;
}

.acrylic-effect {
  background: rgba(255, 255, 255, 0.3);
  backdrop-filter: blur(10px);
  border-radius: 12px;
  transition: all 0.3s ease;
}

/* 修改表格单元格样式，支持自动换行 */
.multi-line-cell {
  white-space: normal;
  word-break: break-all;
  line-height: 1.5;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .acrylic-effect {
    background: rgba(255, 255, 255, 0.15);
  }

  :deep(.el-table) {
    background-color: transparent;
    --el-table-border-color: rgba(255, 255, 255, 0.1);
    --el-table-header-bg-color: rgba(255, 255, 255, 0.05);
    --el-table-row-hover-bg-color: rgba(255, 255, 255, 0.08);
    --el-table-text-color: #e6e6e6;
    --el-table-header-text-color: #ffffff;
  }

  :deep(.el-table__empty-block) {
    background: rgba(255, 255, 255, 0.1);
    backdrop-filter: blur(10px);
    border-radius: 8px;
    margin: 8px;
  }

  :deep(.el-table__empty-text) {
    color: #909399;
  }
}

/* 单元格内容单行显示并居中 */
:deep(.el-table .cell) {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  text-align: center;
}

/* 鼠标悬停时显示完整内容 */
:deep(.el-table .el-tooltip__trigger) {
  display: inline-block;
  width: 100%;
}

/* 调整表格行高 */
:deep(.el-table__row) {
  height: auto;
}

:deep(.el-table__cell) {
  padding: 8px 0;
}

.secret-text {
  font-family: monospace;
  font-size: 0.9em;
  color: #666;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .secret-text {
    color: #999;
  }
}

.status-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
}

.status-message {
  font-size: 12px;
  color: #666;
  margin-top: 4px;
  word-break: break-all;
  max-width: 100%;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .status-message {
    color: #999;
  }
}

.status-info {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
}

.interval-text {
  font-size: 12px;
  color: #666;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .interval-text {
    color: #999;
  }
}

.interval-input {
  display: flex;
  align-items: center;
  gap: 8px;
}

.interval-label {
  white-space: nowrap;
  color: var(--el-text-color-primary);
}

/* 移除number类型输入框的上下箭头 */
:deep(.el-input__inner[type="number"]) {
  -moz-appearance: textfield;
  -webkit-appearance: textfield;
  appearance: textfield;
}
:deep(.el-input__inner[type="number"]::-webkit-outer-spin-button),
:deep(.el-input__inner[type="number"]::-webkit-inner-spin-button) {
  -webkit-appearance: none;
  appearance: none;
  margin: 0;
}

.interval-input, .mode-select {
  display: flex;
  align-items: center;
  gap: 8px;
}

.interval-label, .mode-label {
  white-space: nowrap;
  color: var(--el-text-color-primary);
}
</style>

