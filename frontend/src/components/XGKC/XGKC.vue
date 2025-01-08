<template>
  <div class="xgkc-container">
    <h1>公选课</h1>
    <div v-if="error" class="error">
      {{ error }}
    </div>
    <div class="table-container">
      <div class="table-header">
        <div class="filter-group">
          <el-select
            v-model="selectedType"
            placeholder="课程类型筛选"
            clearable
            class="type-filter"
            :style="{ width: selectWidth }"
          >
            <el-option
              v-for="type in availableTypes"
              :key="type"
              :label="type"
              :value="type"
            />
          </el-select>
          
          <el-input
            v-model="searchKeyword"
            placeholder="搜索课程名称"
            clearable
            class="search-input"
          >
            <template #prefix>
              <el-icon><Search /></el-icon>
            </template>
          </el-input>
        </div>
        
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
            <div class="multi-line-cell">{{ scope.row.XGXKLB }}</div>
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
          prop="SFYX"
          label="操作" 
          width="120"
          align="center"
          header-align="center"
          sortable="custom"
        >
          <template #default="scope">
            <el-button 
              :type="scope.row.SFYX === '1' ? 'danger' : 'primary'"
              size="small"
              @click="scope.row.SFYX === '1' ? handleUnselect(scope.row) : handleSelect(scope.row)"
            >
              {{ scope.row.SFYX === '1' ? '退选' : '选课' }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed, watch } from 'vue'
import { useGlobalStore } from '../../stores/globalStore'
import { useCourseStore } from '../../stores/courseStore'
import { ElMessage } from 'element-plus'
import { Refresh, Search } from '@element-plus/icons-vue'

const globalStore = useGlobalStore()
const courseStore = useCourseStore()
const loading = ref(false)
const error = ref(null)
const selectedType = ref('全部')  // 选中的类型，默认选择"全部"
const searchKeyword = ref('')  // 添加搜索关键字的响应式变量

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

// 添加计算属性来计算选择框宽度
const selectWidth = computed(() => {
  if (!selectedType.value) {
    return '200px'  // 默认宽度
  }
  // 计算文本宽度（假设每个中文字符16px，英文字符8px）
  const textLength = selectedType.value.split('').reduce((acc, char) => {
    return acc + (/[\u4e00-\u9fa5]/.test(char) ? 16 : 8)
  }, 0)
  // 添加padding和其他UI元素的空间
  const width = textLength + 60
  // 设置最小和最大宽度
  return `${Math.max(200, Math.min(600, width))}px`
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
      '1'
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
      SFYX: course.SFYX
    }
    result.push(courseInfo)
  }

  // 如果选择了类型且不是"全部"，进行筛选
  if (selectedType.value && selectedType.value !== '全部') {
    result = result.filter(course => course.XGXKLB === selectedType.value)
  }

  // 添加关键字筛选
  if (searchKeyword.value) {
    result = result.filter(course => 
      course.KCM.toLowerCase().includes(searchKeyword.value.toLowerCase())
    )
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

// 获取公选课
const getXGKC = async () => {
  // 如果已经加载过数据，就不再请求
  if (courseStore.xgkcLoaded) {
    return
  }

  loading.value = true
  error.value = null
  
  try {
    if (!globalStore.Authorization || !globalStore.batchId) {
      throw new Error('缺少必要的认证信息')
    }

    const data = await window.go.XGKC.App.GetXGKC(globalStore.Authorization, globalStore.batchId)
    if (data.code !== 200) {
      throw new Error(data.msg || '获取课程列表失败')
    }
    courseStore.setXGKCData(data)
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
  courseStore.xgkcLoaded = false
  // 重新获取数据
  await getXGKC()
}

onMounted(() => {
  getXGKC()
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
  width: 400px;
  transition: width 0.3s ease;  /* 添加过渡效果 */
}

/* 确保下拉选项也能完整显示 */
:deep(.el-select-dropdown__item) {
  white-space: normal;
  height: auto;
  padding: 8px 12px;
  line-height: 1.5;
}

/* 调整下拉框的最大高度 */
:deep(.el-select-dropdown) {
  max-height: 400px;
}

/* 修改表格单元格样式，支持自动换行 */
.multi-line-cell {
  white-space: normal;
  word-break: break-all;
  line-height: 1.5;
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
  width: 200px;
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
</style>
