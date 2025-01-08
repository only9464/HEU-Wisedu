<template>
  <div class="fawkc-container">
    <h1>跨专业选修课程</h1>
    <div v-if="error" class="error">
      {{ error }}
    </div>
    <div class="table-container">
      <div class="table-header">
        <div></div>
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
        :data="processedData" 
        style="width: 100%"
        :stripe="true"
        class="acrylic-effect"
        @sort-change="handleSortChange"
        :default-sort="{ prop: 'SFYX', order: 'descending' }"
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
        />
        <el-table-column 
          prop="KCM" 
          label="课程名称" 
          min-width="200" 
          align="center"
          header-align="center"
          sortable="custom"
        />
        <el-table-column 
          prop="SKJS" 
          label="授课教师" 
          min-width="120" 
          align="center"
          header-align="center"
          sortable="custom"
        />
        <el-table-column 
          prop="XF" 
          label="学分" 
          width="80" 
          align="center"
          header-align="center"
          sortable="custom"
        />
        <el-table-column 
          prop="SFYX"
          label="操作" 
          width="200"
          align="center"
          header-align="center"
          sortable="custom"
        >
          <template #default="scope">
            <div class="operation-buttons">
              <el-button 
                :type="scope.row.SFYX === '1' ? 'danger' : 'primary'"
                size="small"
                @click="scope.row.SFYX === '1' ? handleUnselect(scope.row) : handleSelect(scope.row)"
              >
                {{ scope.row.SFYX === '1' ? '退选' : '选课' }}
              </el-button>
              
              <el-button
                v-if="scope.row.SFYX !== '1'"
                type="success"
                size="small"
                @click="handleAddToQueue(scope.row)"
                :disabled="isInQueue(scope.row.JXBID)"
                :class="{ 'is-in-queue': isInQueue(scope.row.JXBID) }"
              >
                {{ isInQueue(scope.row.JXBID) ? '已在队列' : '加入队列' }}
              </el-button>
            </div>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useGlobalStore } from '../../stores/globalStore'
import { useCourseStore } from '../../stores/courseStore'
import { ElMessage } from 'element-plus'
import { Refresh } from '@element-plus/icons-vue'

const globalStore = useGlobalStore()
const courseStore = useCourseStore()
const loading = ref(false)
const error = ref(null)

// 添加排序相关的变量和方法
const sortOrder = ref({
  prop: 'SFYX',
  order: 'descending'
})

// 选课处理函数
const handleSelect = async (course) => {
  try {
    if (!globalStore.Authorization || !globalStore.batchId) {
      throw new Error('缺少必要的认证信息')
    }

    const result = await window.go.only9464.App.AddClazz(
      globalStore.Authorization,
      globalStore.batchId,
      'FAWKC',
      course.JXBID,
      course.secretVal,
      '1'
    )

    const jsonResult = JSON.parse(result)
    if (jsonResult.code === 200) {
      ElMessage.success(jsonResult.msg || '选课成功')
      // 更新 store 中的数据
      const rows = courseStore.fawkcData.data.rows
      for (const row of rows) {
        for (const tc of row.tcList) {
          if (tc.JXBID === course.JXBID) {
            tc.SFYX = '1'
            break
          }
        }
      }
    } else {
      throw new Error(jsonResult.msg || '选课失败')
    }
  } catch (err) {
    ElMessage.error(err.message)
  }
}

// 退选处理函数
const handleUnselect = async (course) => {
  try {
    if (!globalStore.Authorization || !globalStore.batchId) {
      throw new Error('缺少必要的认证信息')
    }

    const result = await window.go.only9464.App.DelClazz(
      globalStore.Authorization,
      globalStore.batchId,
      'FAWKC',
      course.JXBID,
      course.secretVal
    )

    const jsonResult = JSON.parse(result)
    if (jsonResult.code === 200) {
      ElMessage.success(jsonResult.msg || '退选成功')
      // 更新 store 中的数据
      const rows = courseStore.fawkcData.data.rows
      for (const row of rows) {
        for (const tc of row.tcList) {
          if (tc.JXBID === course.JXBID) {
            tc.SFYX = '0'
            break
          }
        }
      }
    } else {
      throw new Error(jsonResult.msg || '退选失败')
    }
  } catch (err) {
    ElMessage.error(err.message)
  }
}

// 处理数据的计算属性
const processedData = computed(() => {
  if (!courseStore.fawkcData || !courseStore.fawkcData.data || !courseStore.fawkcData.data.rows) {
    return []
  }

  let result = []
  for (const course of courseStore.fawkcData.data.rows) {
    for (const tc of course.tcList) {
      const courseInfo = {
        JXBID: tc.JXBID,
        KCM: tc.KCM,
        SKJS: tc.SKJS,
        secretVal: tc.secretVal,
        XF: tc.XF,
        SFYX: tc.SFYX
      }
      result.push(courseInfo)
    }
  }

  // 添加排序逻辑
  if (sortOrder.value.prop) {
    const { prop, order } = sortOrder.value
    result.sort((a, b) => {
      let compareResult = 0
      // 处理不同类型的排序
      switch (prop) {
        case 'XF':  // 学分按数字排序
          compareResult = parseFloat(a[prop]) - parseFloat(b[prop])
          break
        case 'SFYX':  // 选课状态按数字排序
          compareResult = parseInt(a[prop]) - parseInt(b[prop])
          break
        default:  // 其他字段按字符串排序
          compareResult = String(a[prop]).localeCompare(String(b[prop]), 'zh-CN')
      }
      // 根据排序方向调整结果
      return order === 'ascending' ? compareResult : -compareResult
    })
  }

  return result
})

// 添加排序处理函数
const handleSortChange = ({ prop, order }) => {
  sortOrder.value = { prop, order }
}

// 获取跨专业选修课程
const getFAWKC = async () => {
  // 如果已经加载过数据，就不再请求
  if (courseStore.fawkcLoaded) {
    return
  }

  loading.value = true
  error.value = null
  
  try {
    if (!globalStore.Authorization || !globalStore.batchId) {
      throw new Error('缺少必要的认证信息')
    }

    const data = await window.go.FAWKC.App.GetFAWKC(globalStore.Authorization, globalStore.batchId)
    if (data.code !== 200) {
      throw new Error(data.msg || '获取课程列表失败')
    }
    courseStore.setFAWKCData(data)
  } catch (err) {
    error.value = err.message
    ElMessage.error(err.message)
  } finally {
    loading.value = false
  }
}

// 手动刷新处理函数
const handleRefresh = async () => {
  // 清除已加载标记
  courseStore.fawkcLoaded = false
  // 重新获取数据
  await getFAWKC()
}

// 修改添加到任务队列的方法
const handleAddToQueue = (course) => {
  const task = {
    JXBID: course.JXBID,
    KCM: course.KCM,
    SKJS: course.SKJS,
    XF: course.XF,
    clazzType: 'FAWKC',       // 添加课程类型
    secretVal: course.secretVal  // 添加secretVal
  }
  courseStore.addToTaskQueue(task)
  ElMessage.success('已添加到任务队列')
}

// 添加判断课程是否在队列中的方法
const isInQueue = (jxbid) => {
  return courseStore.taskQueue.some(task => task.JXBID === jxbid)
}

onMounted(() => {
  getFAWKC()
})
</script>

<style scoped>
.fawkc-container {
  height: 100%;
  display: flex;
  flex-direction: column;
  padding: 0 20px;
  gap: 20px;
  overflow: hidden;
}

.error {
  color: #f56c6c;
  text-align: center;
  padding: 20px;
  margin-top: 20px;
}

h1 {
  margin-bottom: 20px;
  color: #333;
}

.table-container {
  padding: 20px;
}

.acrylic-effect {
  background: rgba(255, 255, 255, 0.3);
  backdrop-filter: blur(10px);
  border-radius: 12px;
  transition: all 0.3s ease;
}

.secret-text {
  font-family: monospace;
  font-size: 12px;
  color: #666;
  cursor: pointer;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  h1 {
    color: #fff;
  }

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

  .secret-text {
    color: #999;
  }
}

.table-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  padding: 0 8px;
}

.operation-buttons {
  display: flex;
  gap: 8px;
  justify-content: center;
}

/* 添加已在队列中按钮的样式 */
.is-in-queue {
  opacity: 0.7;
  cursor: not-allowed;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .is-in-queue {
    opacity: 0.5;
  }
}
</style>
