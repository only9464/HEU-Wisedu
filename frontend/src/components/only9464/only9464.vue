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
          <el-button 
            type="primary"
            @click="handleStartTask"
          >
            开始抢课任务
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
          prop="JXBID" 
          label="教学班ID" 
          min-width="185" 
          align="center"
          header-align="center"
        />
        <el-table-column 
          prop="KCM" 
          label="课程名称" 
          min-width="200" 
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
import { ElMessage, ElMessageBox } from 'element-plus'
import { Refresh } from '@element-plus/icons-vue'
import { useCourseStore } from '../../stores/courseStore'
import {useGlobalStore} from "../../stores/globalStore.js";

const globalStore = useGlobalStore()
const loading = ref(false)
const taskList = ref([])
const courseStore = useCourseStore()

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

// 添加开始抢课任务处理函数
const handleStartTask = async () => {
  ElMessage.info('功能正在开发')
  // 获取任务队列内容
  const taskQueue = courseStore.getTaskQueue();

  // 将任务队列内容转换为 JSON 字符串
  const taskQueueJSON = JSON.stringify(taskQueue);

  // // 使用 ElMessageBox 弹出对话框显示任务队列内容
  // ElMessageBox.alert(
  //     JSON.stringify(taskQueue, null, 2),
  //     '任务队列内容',
  //     {
  //       confirmButtonText: '确定',
  //       dangerouslyUseHTMLString: true, // 允许使用 HTML 内容
  //       customClass: 'task-queue-dialog'
  //     }
  // );

  try {
    if (!globalStore.Authorization || !globalStore.batchId) {
      throw new Error('缺少必要的认证信息')
    }

    // 修正参数顺序：Authorization, batchID, taskListJSON
    const [results, errors] = await window.go.only9464.App.TaskListAddClazz(
        globalStore.Authorization,
        globalStore.batchId,
        taskQueueJSON
    );
    // 显示结果消息
    await ElMessageBox.alert(
        results,
        '任务执行结果',
        {
          confirmButtonText: '确定',
          dangerouslyUseHTMLString: true, // 允许使用 HTML 内容
          customClass: 'task-results-dialog'
        }
    );

    // 处理结果和错误
    let successCount = 0;
    let errorCount = 0;
    let resultsMessage = '';

    results.forEach((result, index) => {
      if (errors[index]) {
        resultsMessage += `任务 ${index + 1} 失败: ${errors[index].message}\n`;
        errorCount++;
      } else {
        resultsMessage += `任务 ${index + 1} 成功: ${result}\n`;
        successCount++;
      }
    });

    if (successCount > 0 && errorCount === 0) {
      ElMessage.success('所有任务启动成功');
    } else if (successCount > 0 && errorCount > 0) {
      ElMessage.warning(`部分任务启动成功，部分任务失败`);
    } else {
      ElMessage.error('所有任务启动失败');
    }
  } catch (error) {
    ElMessage.error('任务队列启动失败：' + error.message);
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
</style>

