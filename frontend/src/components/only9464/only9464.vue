<template>
  <div class="only9464-container">
    <h1>抢课助手</h1>
    
    <!-- 任务队列表格 -->
    <div class="task-table-container">
      <div class="table-header">
        <h2>任务队列</h2>
        <el-button 
          type="primary" 
          :icon="Refresh"
          circle
          @click="handleRefresh"
          :loading="loading"
        />
      </div>
      
      <el-table 
        v-loading="loading"
        :data="taskList" 
        style="width: 100%"
        :stripe="true"
        class="acrylic-effect"
      >
        <el-table-column 
          type="index" 
          label="序号" 
          width="60" 
          align="center"
          header-align="center"
        />
        <el-table-column 
          prop="JXBID" 
          label="教学班ID" 
          min-width="180" 
          align="center"
          header-align="center"
          show-overflow-tooltip
        />
        <el-table-column 
          prop="KCM" 
          label="课程名称" 
          min-width="200" 
          align="center"
          header-align="center"
          show-overflow-tooltip
        />
        <el-table-column 
          prop="SKJS" 
          label="授课教师" 
          min-width="120" 
          align="center"
          header-align="center"
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
        />
        <el-table-column 
          prop="clazzType" 
          label="课程类型" 
          width="120" 
          align="center"
          header-align="center"
        >
          <template #default="scope">
            <el-tag :type="getTagType(scope.row.clazzType)">
              {{ getClazzTypeLabel(scope.row.clazzType) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column 
          prop="secretVal" 
          label="密钥值" 
          min-width="180" 
          align="center"
          header-align="center"
          show-overflow-tooltip
        >
          <template #default="scope">
            <span class="secret-text">{{ scope.row.secretVal }}</span>
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
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { Refresh } from '@element-plus/icons-vue'
import { useCourseStore } from '../../stores/courseStore'

const loading = ref(false)
const taskList = ref([])
const courseStore = useCourseStore()

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

// 获取任务列表
const getTaskList = () => {
  loading.value = true
  try {
    taskList.value = courseStore.getTaskQueue()
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
</style>

