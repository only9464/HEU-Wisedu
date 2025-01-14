<template>
  <div class="tjkc-container">
    <div v-if="error" class="error">
      {{ error }}
    </div>
    <div class="table-container">
      <div class="table-header">
        <div class="filter-group">
          <el-checkbox v-model="tjkcStore.hideSelected">只看未选</el-checkbox>
        </div>
        <div class="header-buttons">
          <el-button 
            type="primary" 
            :icon="Refresh"
            circle
            @click="handleRefresh"
            :loading="loading"
          />
          <el-button 
            type="success"
            @click="handleBatchAddToQueue"
            :loading="loading"
          >
            一键添加到任务队列({{ availableCoursesCount }})
          </el-button>
        </div>
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
          prop="KCXZ" 
          label="课程性质" 
          min-width="120" 
          align="center" 
          header-align="center"
          sortable="custom"
        >
          <template #default="scope">
            <div class="multi-line-cell">{{ scope.row.KCXZ }}</div>
          </template>
        </el-table-column>
        <el-table-column 
          prop="DYZYRS"
          label="选课人数" 
          width="120"
          align="center"
          header-align="center"
          sortable="custom"
          :sort-method="handleCapacitySort"
        >
          <template #default="scope">
            <el-tag :type="scope.row.SFYM === '1' ? 'danger' : 'success'">
              {{ scope.row.DYZYRS }} / {{ scope.row.KRL }}
            </el-tag>
          </template>
        </el-table-column>
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
import { useTJKCStore } from '../../stores/tjkcStore'
import { ElMessage } from 'element-plus'
import { Refresh } from '@element-plus/icons-vue'

const globalStore = useGlobalStore()
const courseStore = useCourseStore()
const tjkcStore = useTJKCStore()
const loading = ref(false)
const error = ref(null)

// 添加排序相关的变量和方法
const sortOrder = ref({
  prop: 'SFYX',  // 默认按选课状态排序
  order: 'descending'  // 默认降序，使已选课程在上面
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
      'TJKC',
      course.JXBID,
      course.secretVal,
      // '1'
    )

    const jsonResult = JSON.parse(result)
    if (jsonResult.code === 200) {
      ElMessage.success(jsonResult.msg || '选课成功')
      // 更新 store 中的数据
      const rows = courseStore.tjkcData.data.rows
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
      'TJKC',
      course.JXBID,
      course.secretVal
    )

    const jsonResult = JSON.parse(result)
    if (jsonResult.code === 200) {
      ElMessage.success(jsonResult.msg || '退选成功')
      // 更新 store 中的数据
      const rows = courseStore.tjkcData.data.rows
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

// 添加到任务队列
const handleAddToQueue = (course) => {
  const task = {
    JXBID: course.JXBID,
    KCM: course.KCM,
    SKJS: course.SKJS,
    XF: course.XF,
    clazzType: 'TJKC',
    secretVal: course.secretVal,
    KCXZ: course.KCXZ
  }
  courseStore.addToTaskQueue(task)
  ElMessage.success('已添加到任务队列')
}

// 处理数据的计算属性
const processedData = computed(() => {
  if (!courseStore.tjkcData || !courseStore.tjkcData.data || !courseStore.tjkcData.data.rows) {
    return []
  }

  let result = []
  for (const course of courseStore.tjkcData.data.rows) {
    for (const tc of course.tcList) {
      const courseInfo = {
        JXBID: tc.JXBID,
        KCM: tc.KCM,
        SKJS: tc.SKJS,
        secretVal: tc.secretVal,
        XF: tc.XF,
        SFYX: tc.SFYX,
        KCXZ: tc.KCXZ,
        KRL: tc.KRL,         // 添加课容量
        DYZYRS: tc.DYZYRS,   // 添加已报人数
        SFYM: tc.SFYM        // 添加是否已满
      }
      result.push(courseInfo)
    }
  }

  // 添加只看未选筛选
  if (tjkcStore.hideSelected) {
    result = result.filter(course => course.SFYX !== '1')
  }

  // 修改排序逻辑
  if (sortOrder.value.prop) {
    const { prop, order } = sortOrder.value
    result.sort((a, b) => {
      let compareResult = 0
      // 处理不同类型的排序
      switch (prop) {
        case 'XF':  // 学分按数字排序
        case 'KRL':  // 课容量按数字排序
        case 'DYZYRS':  // 已报人数按数字排序
          compareResult = parseFloat(a[prop]) - parseFloat(b[prop])
          break
        case 'SFYM':  // 是否已满按数字排序
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

// 添加选课人数列的排序方法
const handleCapacitySort = (a, b) => {
  // 计算选课比例作为排序依据
  const ratioA = parseFloat(a.DYZYRS) / parseFloat(a.KRL)
  const ratioB = parseFloat(b.DYZYRS) / parseFloat(b.KRL)
  return ratioA - ratioB
}

// 添加排序处理函数
const handleSortChange = ({ prop, order }) => {
  sortOrder.value = { prop, order }
}

// 修改错误处理逻辑
const handleError = (err) => {
  if (err.message && err.message.includes('缺少必要的认证信息')) {
    error.value = '缺少必要的认证信息，请先登录选课系统'
  } else {
    error.value = err.message
  }
  ElMessage.error(error.value)
}

// 获取培养方案内课程
const getTJKC = async () => {
  loading.value = true
  error.value = null
  
  try {
    if (!globalStore.Authorization || !globalStore.batchId) {
      throw new Error('缺少必要的认证信息，请先登录选课系统')
    }
    
    const result = await window.go.TJKC.App.GetTJKC(
      globalStore.Authorization,
      globalStore.batchId
    )
    courseStore.setTJKCData(result)
  } catch (err) {
    handleError(err)
  } finally {
    loading.value = false
  }
}

// 手动刷新处理函数
const handleRefresh = async () => {
  // 清除已加载标记
  courseStore.tjkcLoaded = false
  // 重新获取数据
  await getTJKC()
}

// 添加判断课程是否在队列中的方法
const isInQueue = (jxbid) => {
  return courseStore.taskQueue.some(task => task.JXBID === jxbid)
}

// 添加一键添加到任务队列的处理函数
const handleBatchAddToQueue = () => {
  // 获取当前展示的课程列表
  const currentCourses = processedData.value
  
  // 过滤掉已选课程和已在队列中的课程
  const coursesToAdd = currentCourses.filter(course => 
    course.SFYX !== '1' && !isInQueue(course.JXBID)
  )

  if (coursesToAdd.length === 0) {
    ElMessage.warning('没有可添加的课程')
    return
  }

  // 批量添加到任务队列
  coursesToAdd.forEach(course => {
    const task = {
      JXBID: course.JXBID,
      KCM: course.KCM,
      SKJS: course.SKJS,
      XF: course.XF,
      clazzType: 'TJKC',
      secretVal: course.secretVal,
      KCXZ: course.KCXZ
    }
    courseStore.addToTaskQueue(task)
  })

  ElMessage.success(`成功添加 ${coursesToAdd.length} 门课程到任务队列`)
}

// 添加计算属性：计算可添加到队列的课程数量
const availableCoursesCount = computed(() => {
  return processedData.value.filter(course => 
    course.SFYX !== '1' && !isInQueue(course.JXBID)
  ).length
})

onMounted(() => {
  getTJKC()
})
</script>

<style scoped>
.tjkc-container {
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

/* 确保表头文字和排序图标正确对齐 */
:deep(.el-table__header) th {
  padding: 8px 0;
  height: 40px;
}

:deep(.el-table__header) .cell {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 4px;
}

/* 添加按钮组样式 */
.header-buttons {
  display: flex;
  gap: 12px;
  align-items: center;
}

.filter-group {
  display: flex;
  gap: 16px;
  align-items: center;
}

.el-checkbox {
  margin: 0;
  --el-checkbox-font-size: 14px;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .el-checkbox {
    --el-checkbox-text-color: var(--el-text-color-primary);
    --el-checkbox-checked-text-color: var(--el-color-primary);
  }
}
</style>
