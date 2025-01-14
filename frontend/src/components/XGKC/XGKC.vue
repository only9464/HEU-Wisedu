<template>
  <div class="xgkc-container">
    <div v-if="error" class="error">
      {{ error }}
    </div>
    <div class="table-container">
      <div class="table-header">
        <div class="filter-group">
          <el-select
            v-model="xgkcStore.selectedType"
            placeholder="课程类型筛选"
            clearable
            class="type-filter"
            :popper-append-to-body="false"
          >
            <el-option
              v-for="type in availableTypes"
              :key="type"
              :label="type"
              :value="type"
              class="type-option"
            />
          </el-select>

          <el-input
            v-model="xgkcStore.searchKeyword"
            placeholder="搜索"
            clearable
            class="search-input"
          >
            <template #prefix>
              <el-icon><Search /></el-icon>
            </template>
          </el-input>

          <div class="checkbox-group">
            <el-checkbox v-model="xgkcStore.hideSelected">只看未选</el-checkbox>
            <el-checkbox v-model="xgkcStore.onlyNotFull">只看未满</el-checkbox>
            <el-checkbox v-model="xgkcStore.onlyOnline">只看网络课</el-checkbox>
          </div>
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
          show-overflow-tooltip
        />
        <el-table-column 
          prop="XGXKLB" 
          label="类型" 
          min-width="140"
          align="center"
          header-align="center"
          sortable="custom"
        >
          <template #default="scope">
            <div class="type-cell">{{ scope.row.XGXKLB }}</div>
          </template>
        </el-table-column>
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

  <!-- 添加教师对话框 -->
  <el-dialog
    v-model="addTeacherDialogVisible"
    title="添加要过滤的教师"
    width="30%"
  >
    <el-input
      v-model="newTeacher"
      placeholder="请输入教师姓名"
      clearable
    />
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="addTeacherDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="addTeacher">
          确定
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup>
import { ref, onMounted, computed, watch, onUnmounted } from 'vue'
import { useGlobalStore } from '../../stores/globalStore'
import { useCourseStore } from '../../stores/courseStore'
import { useXGKCStore } from '../../stores/xgkcStore'
import { ElMessage } from 'element-plus'
import { Refresh, Search, Plus, Delete } from '@element-plus/icons-vue'

const globalStore = useGlobalStore()
const courseStore = useCourseStore()
const xgkcStore = useXGKCStore()
const loading = ref(false)
const error = ref(null)

// 添加教师过滤相关的响应式变量
const addTeacherDialogVisible = ref(false)
const newTeacher = ref('')

// 添加教师过滤相关的方法
const showAddTeacherDialog = () => {
  addTeacherDialogVisible.value = true
  newTeacher.value = ''
}

const addTeacher = () => {
  if (newTeacher.value.trim()) {
    xgkcStore.addFilteredTeacher(newTeacher.value.trim())
    addTeacherDialogVisible.value = false
    newTeacher.value = ''  // 清空输入
  }
}

const removeTeacher = (teacher) => {
  xgkcStore.removeFilteredTeacher(teacher)
}

// 修改 availableTypes 计算属性
const availableTypes = computed(() => {
  if (!courseStore.xgkcData || !courseStore.xgkcData.data || !courseStore.xgkcData.data.rows) {
    return ['全部']  // 即使没有数据也显示"全部"选项
  }

  // 使用 Set 来去重
  const typeSet = new Set()
  courseStore.xgkcData.data.rows.forEach(course => {
    typeSet.add(course.XGXKLB)
  })

  // 转换为数组，添加"全部"选项并排序
  return ['全部', ...Array.from(typeSet).sort()]
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
      'XGKC',
      course.JXBID,
      course.secretVal,
      // '1'
    )

    const jsonResult = JSON.parse(result)
    if (jsonResult.code === 200) {
      ElMessage.success(jsonResult.msg || '选课成功')
      // 更新 store 中的数据
      const rows = courseStore.xgkcData.data.rows
      for (const row of rows) {
        if (row.JXBID === course.JXBID) {
          row.SFYX = '1'
          break
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
      'XGKC',
      course.JXBID,
      course.secretVal
    )

    const jsonResult = JSON.parse(result)
    if (jsonResult.code === 200) {
      ElMessage.success(jsonResult.msg || '退选成功')
      // 更新 store 中的数据
      const rows = courseStore.xgkcData.data.rows
      for (const row of rows) {
        if (row.JXBID === course.JXBID) {
          row.SFYX = '0'
          break
        }
      }
    } else {
      throw new Error(jsonResult.msg || '退选失败')
    }
  } catch (err) {
    ElMessage.error(err.message)
  }
}

// 在 script setup 中添加
const sortOrder = ref({
  prop: 'SFYX',  // 默认按选课状态排序
  order: 'descending'  // 默认降序，使已选课程在上面
})

// 修改 processedData 计算属性
const processedData = computed(() => {
  if (!courseStore.xgkcData || !courseStore.xgkcData.data || !courseStore.xgkcData.data.rows) {
    return []
  }

  let result = []
  for (const course of courseStore.xgkcData.data.rows) {
    const courseInfo = {
      JXBID: course.JXBID,
      XGXKLB: course.XGXKLB,
      KCM: course.KCM,
      SKJS: course.SKJS,
      secretVal: course.secretVal,
      XF: course.XF,
      SFYX: course.SFYX,
      KRL: course.KRL,         // 添加课容量
      DYZYRS: course.DYZYRS,   // 添加已报人数
      SFYM: course.SFYM        // 添加是否已满
    }
    result.push(courseInfo)
  }

  // 如果选择了类型且不是"全部"，进行筛选
  if (xgkcStore.selectedType && xgkcStore.selectedType !== '全部') {
    result = result.filter(course => course.XGXKLB === xgkcStore.selectedType)
  }

  // 添加关键字筛选
  if (xgkcStore.searchKeyword) {
    result = result.filter(course => 
      course.KCM.toLowerCase().includes(xgkcStore.searchKeyword.toLowerCase())
    )
  }

  // 添加只看未满筛选
  if (xgkcStore.onlyNotFull) {
    result = result.filter(course => course.SFYM !== '1')
  }

  // 添加只看网络课筛选
  if (xgkcStore.onlyOnline) {
    result = result.filter(course => course.KCM.includes('网络'))
  }

  // 添加不看已选筛选
  if (xgkcStore.hideSelected) {
    result = result.filter(course => course.SFYX !== '1')
  }

  // 添加排序逻辑
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

// 获取公选课
const getXGKC = async () => {
  loading.value = true
  error.value = null
  
  try {
    if (!globalStore.Authorization || !globalStore.batchId) {
      throw new Error('缺少必要的认证信息，请先登录选课系统')
    }
    
    const result = await window.go.XGKC.App.GetXGKC(
      globalStore.Authorization,
      globalStore.batchId
    )
    courseStore.setXGKCData(result)
  } catch (err) {
    handleError(err)
  } finally {
    loading.value = false
  }
}

// 手动刷新处理函数
const handleRefresh = async () => {
  // 清除已加载标记
  courseStore.xgkcLoaded = false
  // 重新获取数据
  await getXGKC()
}

// 修改添加到任务队列的方法
const handleAddToQueue = (course) => {
  const task = {
    JXBID: course.JXBID,
    KCM: course.KCM,
    SKJS: course.SKJS,
    XF: course.XF,
    clazzType: 'XGKC',
    secretVal: course.secretVal,
    KCXZ: course.XGXKLB  // 使用XGXKLB作为KCXZ
  }
  courseStore.addToTaskQueue(task)
  ElMessage.success('已添加到任务队列')
}

// 添加判断课程是否在队列中的方法
const isInQueue = (jxbid) => {
  return courseStore.taskQueue.some(task => task.JXBID === jxbid)
}

// 添加选课人数列的排序方法
const handleCapacitySort = (a, b) => {
  // 计算选课比例作为排序依据
  const ratioA = parseFloat(a.DYZYRS) / parseFloat(a.KRL)
  const ratioB = parseFloat(b.DYZYRS) / parseFloat(b.KRL)
  return ratioA - ratioB
}

// 添加一键添加到任务队列的处理函数
const handleBatchAddToQueue = () => {
  // 获取当前展示的课程列表（已经过筛选）
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
      clazzType: 'XGKC',
      secretVal: course.secretVal,
      KCXZ: course.XGXKLB
    }
    courseStore.addToTaskQueue(task)
  })

  ElMessage.success(`成功添加 ${coursesToAdd.length} 门课程到任务队列`)
}

// 在组件卸载时重置筛选状态
onUnmounted(() => {
  // 如果需要在组件卸载时重置状态，取消注释下面这行
  // xgkcStore.resetFilters()
})

onMounted(() => {
  getXGKC()
})

// 添加计算属性：计算可添加到队列的课程数量
const availableCoursesCount = computed(() => {
  return processedData.value.filter(course => 
    course.SFYX !== '1' && !isInQueue(course.JXBID)
  ).length
})
</script>

<style scoped>
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

.xgkc-container {
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
    color: #ffffff;
  }

  .acrylic-effect {
    background: rgba(255, 255, 255, 0.2);
  }

  :deep(.el-table) {
    background-color: transparent;
    --el-table-border-color: rgba(255, 255, 255, 0.1);
    --el-table-header-bg-color: rgba(255, 255, 255, 0.05);
    --el-table-row-hover-bg-color: rgba(255, 255, 255, 0.08);
    --el-table-text-color: #ffffff;
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

  :deep(.el-select) {
    --el-select-input-focus-border-color: var(--el-color-primary);
  }
}

.table-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  padding: 0 8px;
  gap: 16px;
}

.type-filter {
  width: 240px;  /* 增加宽度以适应长文本 */
  transition: width 0.3s ease;
}

/* 确保下拉选项也能完整显示 */
:deep(.el-select-dropdown__item) {
  white-space: normal;  /* 允许文字换行 */
  height: auto;
  padding: 8px 12px;
  line-height: 1.5;
}

/* 调整下拉框的最大高度 */
:deep(.el-select-dropdown) {
  max-height: 400px;
  min-width: 240px !important;  /* 确保下拉框和选择框宽度一致 */
}

/* 确保选择框内的文字不被截断 */
:deep(.el-select .el-input__inner) {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

/* 调整表格行高 */
:deep(.el-table__row) {
  height: auto;
}

:deep(.el-table__cell) {
  padding: 8px 0;
}

.filter-group {
  display: flex;
  gap: 16px;
  align-items: center;
  flex: 1;
}

.search-input {
  width: 120px;  /* 调整宽度以适应"搜索"两个字 */
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .search-input :deep(.el-input__wrapper) {
    background-color: rgba(255, 255, 255, 0.05);
  }
  
  .search-input :deep(.el-input__inner) {
    color: var(--el-text-color-primary);
  }
  
  .search-input :deep(.el-input__prefix) {
    color: var(--el-text-color-secondary);
  }
}

/* 添加操作列的样式 */
.operation-cell {
  display: flex;
  align-items: center;
  justify-content: center;
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

/* 添加复选框样式 */
.el-checkbox {
  margin-left: 8px;
  --el-checkbox-font-size: 14px;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .el-checkbox {
    --el-checkbox-text-color: var(--el-text-color-primary);
    --el-checkbox-checked-text-color: var(--el-color-primary);
  }
}

/* 修改复选框组样式 */
.checkbox-group {
  display: flex;
  gap: 16px;
  align-items: center;
}

.el-checkbox {
  margin: 0;  /* 移除原有的margin-left */
  --el-checkbox-font-size: 14px;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .el-checkbox {
    --el-checkbox-text-color: var(--el-text-color-primary);
    --el-checkbox-checked-text-color: var(--el-color-primary);
  }
}

/* 添加按钮组样式 */
.header-buttons {
  display: flex;
  gap: 12px;
  align-items: center;
}

/* 添加类型列的特殊样式 */
.type-cell {
  white-space: normal;  /* 允许文字换行 */
  word-break: break-all;  /* 在任意字符间断行 */
  line-height: 1.5;
  padding: 4px 0;
}
</style>
