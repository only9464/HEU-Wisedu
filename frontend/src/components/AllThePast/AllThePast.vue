<template>
  <div class="allthepast-container">
    <div class="header">
      <div class="title-section">
        <h1>所有已修课程</h1>
        <span v-if="totalSize !== null" class="total-count">
          ({{ totalSize }}门)
        </span>
      </div>
      <div class="header-buttons">
        <el-button 
          type="info" 
          @click="showGraduationRequirements"
        >
          设置毕业要求
        </el-button>
        <el-button 
          type="danger" 
          @click="clearLocalData"
          :disabled="!hasLocalData"
        >
          清除本地查询结果
        </el-button>
        <el-button 
          type="primary" 
          @click="handleQuery"
          :loading="loading"
        >
          {{ hasLocalData ? '重新查询' : '查询' }}
        </el-button>
        <el-button 
          :type="isSettingPublic ? 'warning' : 'success'"
          @click="handleBatchSetPublic"
          :disabled="!hasLocalData"
        >
          {{ isSettingPublic ? '完成设置公选类别' : '设置公选类别' }}
        </el-button>
      </div>
    </div>

    <!-- 添加统计信息 -->
    <div class="statistics" v-if="displayData">
      <div class="stat-items">
        <el-tag 
          type="info" 
          class="stat-tag custom-blue-tag"
          :class="{ 'active-tag': activeFilter === 'required' }"
          @click="handleTagClick('required')"
        >
          必修：{{ requiredCredit }}学分
        </el-tag>
        <el-tag 
          type="info" 
          class="stat-tag custom-orange-tag"
          :class="{ 'active-tag': activeFilter === 'professional' }"
          @click="handleTagClick('professional')"
        >
          专选：{{ professionalCredit }}学分
        </el-tag>
        <el-tag 
          type="info" 
          class="stat-tag custom-purple-tag"
          :class="{ 'active-tag': activeFilter === 'public' }"
          @click="handleTagClick('public')"
        >
          公选：{{ totalPublicCredit }}学分
        </el-tag>
        <el-tag 
          v-for="(credit, category) in publicCourseCredits" 
          :key="category"
          type="info"
          class="stat-tag custom-cyan-tag"
          :class="{ 'active-tag': activeFilter === category }"
          @click="handleTagClick(category)"
        >
          {{ category }}: {{ credit }}学分
        </el-tag>
      </div>
    </div>

    <!-- 添加筛选选项 -->
    <div class="filter-options" v-if="displayData">

      <div class="filter-group">
        <span class="filter-label">成绩筛选：</span>
        <el-checkbox v-model="filters.hideExcellent">忽略优秀</el-checkbox>
        <el-checkbox v-model="filters.hideGood">忽略良好</el-checkbox>
        <el-checkbox v-model="filters.hideMedium">忽略中等</el-checkbox>
        <el-checkbox v-model="filters.hidePass">忽略及格</el-checkbox>
        <div class="checkbox-with-help">
          <el-checkbox v-model="filters.hideFail">忽略不及格</el-checkbox>
          <el-tooltip
            v-if="isSettingPublic"
            content="根据学校规定：不及格的公选课学分不算"
            placement="top"
          >
            <el-icon class="help-icon"><QuestionFilled /></el-icon>
          </el-tooltip>
        </div>
      </div>
            <!-- 非设置公选状态时才显示课程性质筛选 -->
      <div class="filter-group" v-if="!isSettingPublic">
        <span class="filter-label">课程性质：</span>
        <el-checkbox v-model="filters.hideRequired">忽略必修</el-checkbox>
        <el-checkbox v-model="filters.hidePublic">忽略公选</el-checkbox>
        <el-checkbox v-model="filters.hideProfessional">忽略专选</el-checkbox>
      </div>
    </div>

    <!-- 查询结果展示 -->
    <div v-if="displayData" class="result-container">
      <el-table 
        :data="filteredGradeData"
        style="width: 100%"
        :stripe="true"
        v-loading="loading"
        class="acrylic-effect"
        @sort-change="handleSortChange"
      >
        <el-table-column 
          type="index" 
          label="序号" 
          width="70" 
          align="center"
          header-align="center"
        />
        <el-table-column 
          prop="term" 
          label="学期" 
          width="180" 
          align="center"
          header-align="center"
          :sort-orders="['ascending', 'descending']"
          sortable
        />
        <el-table-column 
          prop="name" 
          label="课程名称" 
          width="auto"
          align="center"
          header-align="center"
          :sort-orders="['ascending', 'descending']"
          sortable
          show-overflow-tooltip
        >
          <template #default="scope">
            <span class="course-name-cell">{{ scope.row.name }}</span>
          </template>
        </el-table-column>
        <el-table-column 
          v-if="!isSettingPublic"
          prop="type" 
          label="课程性质" 
          width="160"
          align="center"
          sortable
        >
          <template #default="scope">
            <div class="type-with-category">
              <span>{{ scope.row.type }}</span>
              <el-tag 
                v-if="scope.row.category" 
                size="small" 
                type="success"
                class="category-tag"
              >
                {{ scope.row.category }}
              </el-tag>
            </div>
          </template>
        </el-table-column>
        <el-table-column 
          v-else
          prop="category" 
          label="课程类别" 
          width="120"
          align="center"
        >
          <template #default="scope">
            <el-select 
              v-model="scope.row.category" 
              placeholder="请选择"
              size="small"
              @change="(val) => handleCategoryChange(scope.row.name, val)"
            >
              <el-option label="A" value="A" />
              <el-option label="B" value="B" />
              <el-option label="C" value="C" />
              <el-option label="D" value="D" />
              <el-option label="E" value="E" />
              <el-option label="F" value="F" />
              <el-option label="A0" value="A0" />
            </el-select>
          </template>
        </el-table-column>
        <el-table-column 
          prop="credit" 
          label="学分" 
          width="80"
          align="center"
          sortable
        />
        <el-table-column 
          prop="scoreValue"
          label="成绩" 
          width="100"
          align="center"
          sortable
        >
          <template #default="scope">
            <span :class="getScoreClass(scope.row.score)">
              {{ scope.row.score }}
            </span>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <!-- 在 template 末尾添加弹窗组件 -->
    <el-dialog
      v-model="graduationDialog"
      title="设置毕业要求"
      width="30%"
      center
    >
      <span>功能正在开发</span>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="clearGraduationSettings">清空</el-button>
          <el-button type="primary" @click="graduationDialog = false">
            确定
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useGlobalStore } from '../../stores/globalStore'
import { QuestionFilled } from '@element-plus/icons-vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const loading = ref(false)
const globalStore = useGlobalStore()
const loginResult = ref(null)
const error = ref(null)

// 添加筛选条件
const filters = ref({
  hideRequired: false,
  hidePublic: false,
  hideProfessional: false,
  hideExcellent: false,
  hideGood: false,
  hideMedium: false,
  hidePass: false,
  hideFail: false,
})

// 添加设置公选状态变量
const isSettingPublic = ref(false)

// 添加筛选状态缓存
const cachedFilters = ref(null)

// 在 script setup 中添加新的响应式变量
const activeFilter = ref(null)

// 添加cookies相关的状态
const savedCookies = ref(null)

// 在其他 ref 变量后添加
const graduationDialog = ref(false)

// 显示验证码弹窗改为直接查询
const showCaptchaDialog = () => {
  handleQuery()
}

// 在组件挂载时尝试从localStorage读取cookies
onMounted(() => {
  const storedCookies = localStorage.getItem('gradeCookies')
  if (storedCookies) {
    savedCookies.value = JSON.parse(storedCookies)
  }
})

// 清除本地数据
const clearLocalData = () => {
  ElMessageBox.confirm(
    '确定要清除本地缓存的成绩数据吗？',
    '警告',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    }
  ).then(() => {
    // 清除全局存储中的数据
    globalStore.setGradeData(null)
    globalStore.setAllCourses(null)
    
    // 清除本地变量
    loginResult.value = null
    
    // 清除localStorage中的成绩数据
    localStorage.removeItem('gradeData')
    localStorage.removeItem('allCourses')
    
    ElMessage.success('本地成绩数据已清除')
  }).catch(() => {})
}

// 计算属性：处理后的课程数据（用于表格显示）
const processedGradeData = computed(() => {
  if (!displayData.value) return []
  
  return displayData.value.datas.xscjcx.rows.map(item => {
    const score = item.XSZCJMC || item.ZCJ?.toString() || '--'
    return {
      term: item.XNXQDM_DISPLAY,
      name: item.KCM,
      type: item.KCXZDM_DISPLAY,
      credit: item.XF,
      score: score,
      scoreValue: getScoreValue(score),
      category: globalStore.getCourseCategory(item.KCM),
      marked: false,
      notes: '',
    }
  }).sort((a, b) => {
    return b.term.localeCompare(a.term)
  })
})

// 显示数据（优先使用编辑后的数据）
const displayData = computed(() => {
  if (loginResult.value) return loginResult.value
  return globalStore.allCourses ? JSON.parse(globalStore.allCourses) : null
})

// 计算属性：是否有本地数据
const hasLocalData = computed(() => {
  return !!globalStore.gradeData
})

// 修改为过滤后的数据
const filteredGradeData = computed(() => {
  let data = processedGradeData.value

  // 首先应用标签过滤
  if (activeFilter.value) {
    data = data.filter(item => {
      switch (activeFilter.value) {
        case 'required':
          return item.type === '必修'
        case 'professional':
          return item.type === '专选'
        case 'public':
          return item.type === '公选'
        default:
          // 处理 A-F 和 A0 类别
          return item.type === '公选' && item.category === activeFilter.value
      }
    })
  }

  // 然后应用其他筛选条件
  // 如果是设置公选状态，只显示公选课
  if (isSettingPublic.value) {
    data = data.filter(item => item.type === '公选')
  } else {
    // 正常的课程性质筛选
    data = data.filter(item => {
      if (filters.value.hideRequired && item.type === '必修') return false
      if (filters.value.hidePublic && item.type === '公选') return false
      if (filters.value.hideProfessional && item.type === '专选') return false
      return true
    })
  }

  // 成绩筛选
  data = data.filter(item => {
    const score = item.score
    if (filters.value.hideExcellent && (score === '优秀' || parseFloat(score) >= 90)) return false
    if (filters.value.hideGood && (score === '良好' || (parseFloat(score) >= 80 && parseFloat(score) < 90))) return false
    if (filters.value.hideMedium && (score === '中等' || (parseFloat(score) >= 70 && parseFloat(score) < 80))) return false
    if (filters.value.hidePass && (score === '及格' || (parseFloat(score) >= 60 && parseFloat(score) < 70))) return false
    if (filters.value.hideFail && (score === '不及格' || parseFloat(score) < 60)) return false
    return true
  })

  return data
})

// 修改查询函数
const handleQuery = async () => {
  loading.value = true
  error.value = null

  try {
    // 检查是否有保存的cookies
    const storedCookies = localStorage.getItem('gradeCookies')
    if (!storedCookies) {
      // 如果没有cookies，提示用户去首页登录
      ElMessageBox.confirm(
        '需要先登录教务系统才能查询成绩。是否立即跳转到首页进行登录？',
        '提示',
        {
          confirmButtonText: '立即跳转',
          cancelButtonText: '取消',
          type: 'warning'
        }
      ).then(() => {
        // 用户确认后跳转到首页
        router.push('/')
      })
      return
    }

    // 使用保存的cookies进行查询
    ElMessage.info('正在查询成绩，请稍候...')
    const result = await window.go.login.App.QueryAllGrade(
      JSON.parse(storedCookies)
    )

    if (result.data) {
      globalStore.setGradeData(JSON.stringify(result.data))
      ElMessage.success('查询成功')
      
      // 更新cookies（如果有更新）
      if (result.cookies) {
        localStorage.setItem('gradeCookies', JSON.stringify(result.cookies))
      }
    } else {
      throw new Error('查询失败：返回结果为空')
    }
  } catch (err) {
    console.error('查询失败：', err)
    error.value = err.message
    
    // 如果是cookies失效导致的错误，提示用户重新登录
    if (err.message.includes('cookies') || err.message.includes('登录')) {
      // 清除失效的cookies
      localStorage.removeItem('gradeCookies')
      
      ElMessageBox.confirm(
        '登录状态已失效，需要重新登录教务系统。是否立即跳转到首页进行登录？',
        '提示',
        {
          confirmButtonText: '立即跳转',
          cancelButtonText: '取消',
          type: 'warning'
        }
      ).then(() => {
        // 用户确认后跳转到首页
        router.push('/')
      })
    } else {
      ElMessage.error(err.message)
    }
  } finally {
    loading.value = false
  }
}

// 根据成绩返回对应的样式类名
const getScoreClass = (score) => {
  if (!score || score === '--') return ''
  const numScore = parseFloat(score)
  if (isNaN(numScore)) {
    // 处理优秀、良好等级制成绩
    switch (score) {
      case '优秀': return 'score-excellent'
      case '良好': return 'score-good'
      case '中等': return 'score-medium'
      case '及格': return 'score-pass'
      case '不及格': return 'score-fail'
      default: return ''
    }
  } else {
    // 处理百分制成绩
    if (numScore >= 90) return 'score-excellent'
    if (numScore >= 80) return 'score-good'
    if (numScore >= 70) return 'score-medium'
    if (numScore >= 60) return 'score-pass'
    return 'score-fail'
  }
}

// 获取成绩对应的数值（用于排序）
const getScoreValue = (score) => {
  if (!score || score === '--') return -1
  const numScore = parseFloat(score)
  if (!isNaN(numScore)) return numScore
  
  // 处理等级制成绩
  switch (score) {
    case '优秀': return 95
    case '良好': return 85
    case '中等': return 75
    case '及格': return 65
    case '不及格': return 0
    default: return -1
  }
}

// 修改批量设置公选类别的处理函数
const handleBatchSetPublic = () => {
  if (isSettingPublic.value) {
    // 完成设置 - 恢复之前的筛选状态
    isSettingPublic.value = false
    if (cachedFilters.value) {
      Object.assign(filters.value, cachedFilters.value)
      cachedFilters.value = null
    }
    ElMessage.success('已完成公选类别设置')
  } else {
    // 开始设置 - 缓存当前筛选状态
    cachedFilters.value = { ...filters.value }
    // 清空所有筛选条件
    Object.keys(filters.value).forEach(key => {
      filters.value[key] = false
    })
    // 默认勾选"忽略不及格"
    filters.value.hideFail = true
    isSettingPublic.value = true
    ElMessage.success('已进入公选类别设置模式')
  }
}

// 添加处理类别变化的方法
const handleCategoryChange = (courseName, category) => {
  globalStore.setCourseCategory(courseName, category)
}

// 判断成绩是否及格的辅助函数
const isPassingGrade = (score) => {
  if (!score || score === '--') return false
  const numScore = parseFloat(score)
  if (!isNaN(numScore)) return numScore >= 60
  // 处理文字形式的成绩
  return ['优秀', '良好', '中等', '及格'].includes(score)
}

// 修改学分统计的计算属性
const publicCourseCredits = computed(() => {
  if (!processedGradeData.value) return {}
  
  const credits = {}
  processedGradeData.value
    .filter(course => {
      // 只统计公选课且及格的课程
      return course.type === '公选' && isPassingGrade(course.score)
    })
    .forEach(course => {
      const category = course.category || '未分类的公选课'
      credits[category] = (credits[category] || 0) + parseFloat(course.credit)
    })
  
  // 按类别A-F和A0排序
  const orderedCredits = {}
  const order = ['A', 'B', 'C', 'D', 'E', 'F', 'A0', '未分类的公选课']
  order.forEach(category => {
    if (credits[category]) {
      orderedCredits[category] = Number(credits[category].toFixed(1))  // 保留一位小数
    }
  })
  
  return orderedCredits
})

// 修改必修课学分统计
const requiredCredit = computed(() => {
  if (!processedGradeData.value) return 0
  const total = processedGradeData.value
    .filter(course => {
      return course.type === '必修' && isPassingGrade(course.score)
    })
    .reduce((sum, course) => sum + parseFloat(course.credit), 0)
  return Number(total.toFixed(1))  // 保留一位小数
})

// 修改专选课学分统计
const professionalCredit = computed(() => {
  if (!processedGradeData.value) return 0
  const total = processedGradeData.value
    .filter(course => course.type === '专选' && isPassingGrade(course.score))
    .reduce((sum, course) => sum + parseFloat(course.credit), 0)
  return Number(total.toFixed(1))  // 保留一位小数
})

// 添加公选课总学分统计
const totalPublicCredit = computed(() => {
  if (!processedGradeData.value) return 0
  const total = processedGradeData.value
    .filter(course => course.type === '公选' && isPassingGrade(course.score))
    .reduce((sum, course) => sum + parseFloat(course.credit), 0)
  return Number(total.toFixed(1))  // 保留一位小数
})

// 修改总门数的计算属性
const totalSize = computed(() => {
  if (!processedGradeData.value) return null
  
  // 使用 Set 来去除重复的课程名称
  const uniqueCourses = new Set(
    processedGradeData.value.map(course => course.name)
  )
  
  return uniqueCourses.size
})

// 添加处理标签点击的方法
const handleTagClick = (filter) => {
  if (activeFilter.value === filter) {
    // 如果点击的是当前激活的标签，则取消过滤
    activeFilter.value = null
  } else {
    // 否则设置新的过滤条件
    activeFilter.value = filter
  }
}

// 添加新的方法
const showGraduationRequirements = () => {
  graduationDialog.value = true
}

const clearGraduationSettings = () => {
  // 暂时为空函数
}
</script>

<style scoped>
.allthepast-container {
  padding: 20px;
  max-width: 800px;
  margin: 0 auto;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.result-container {
  margin-top: 20px;
  padding: 20px;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border-radius: 12px;
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.result-content {
  background: rgba(0, 0, 0, 0.05);
  padding: 16px;
  border-radius: 8px;
  overflow-x: auto;
}

pre {
  margin: 0;
  white-space: pre-wrap;
  word-wrap: break-word;
}

.header-buttons {
  display: flex;
  gap: 12px;
  align-items: center;
}

.title-section {
  display: flex;
  align-items: center;
  gap: 8px;
}

.total-count {
  font-size: 1.2em;
  color: var(--el-text-color-primary);
  font-weight: normal;
}

/* 成绩颜色样式 */
.score-excellent {
  color: #67c23a;
  font-weight: bold;
}

.score-good {
  color: #409eff;
  font-weight: bold;
}

.score-medium {
  color: #e6a23c;
  font-weight: bold;
}

.score-pass {
  color: #909399;
}

.score-fail {
  color: #f56c6c;
  font-weight: bold;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .result-container {
    background: rgba(255, 255, 255, 0.05);
  }

  .result-content {
    background: rgba(255, 255, 255, 0.05);
  }

  .total-count {
    color: var(--el-text-color-primary);
  }

  .score-excellent {
    color: #95d475;
  }

  .score-good {
    color: #79bbff;
  }

  .score-medium {
    color: #eebe77;
  }

  .score-pass {
    color: #a6a9ad;
  }

  .score-fail {
    color: #f89898;
  }
}

/* 只保留基本的表格样式 */
:deep(.el-table__header) .cell {
  display: flex;
  align-items: center;
  justify-content: center;
}

.acrylic-effect {
  background: rgba(255, 255, 255, 0.3);
  backdrop-filter: blur(10px);
  border-radius: 12px;
  transition: all 0.3s ease;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .acrylic-effect {
    background: rgba(255, 255, 255, 0.15);
  }
}

.filter-options {
  margin-bottom: 20px;
  padding: 16px;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border-radius: 12px;
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.filter-group {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 12px;
}

.filter-group:last-child {
  margin-bottom: 0;
}

.filter-label {
  font-weight: bold;
  color: var(--el-text-color-primary);
  min-width: 80px;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .filter-options {
    background: rgba(255, 255, 255, 0.05);
  }
}

.checkbox-with-help {
  display: flex;
  align-items: center;
  gap: 4px;
}

.help-icon {
  color: var(--el-text-color-secondary);
  font-size: 16px;
  cursor: help;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .help-icon {
    color: var(--el-text-color-secondary);
  }
}

.course-name-cell {
  white-space: normal;
  word-break: break-all;
  line-height: 1.5;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-align: center;
}

:deep(.el-table .cell) {
  white-space: normal !important;
  line-height: 1.5;
}

/* 确保表格行高足够 */
:deep(.el-table__row) {
  min-height: 50px;
}

.type-with-category {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
}

.category-tag {
  font-size: 12px;
  padding: 0 4px;
  height: 20px;
  line-height: 18px;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .category-tag {
    border-color: var(--el-color-success-dark);
    background-color: var(--el-color-success-dark);
  }
}

.statistics {
  margin-bottom: 20px;
  padding: 12px 16px;  /* 减小上下内边距 */
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border-radius: 12px;
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.stat-items {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  gap: 8px;
}

.stat-tag {
  font-size: 12px;
  padding: 0 8px;
  height: 24px;
  line-height: 22px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.stat-tag:hover {
  opacity: 0.8;
}

.active-tag {
  transform: scale(1.05);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .active-tag {
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.4);
  }
}

/* 分隔线样式 */
:deep(.el-divider--vertical) {
  height: 20px;
  margin: 0 4px;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .statistics {
    background: rgba(255, 255, 255, 0.05);
  }
}

/* 自定义蓝色标签 */
.custom-blue-tag {
  background-color: #3498db !important;
  border-color: #3498db !important;
  color: white !important;
}

/* 自定义橙色标签 - 使用更深的橙色 */
.custom-orange-tag {
  background-color: #d35400 !important;
  border-color: #d35400 !important;
  color: white !important;
}

/* 自定义紫色标签 - 调得更暗一些 */
.custom-purple-tag {
  background-color: #5b2c6f !important;
  border-color: #5b2c6f !important;
  color: white !important;
}

/* 自定义青色标签 */
.custom-cyan-tag {
  background-color: #16a085 !important;
  border-color: #16a085 !important;
  color: white !important;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .custom-blue-tag {
    background-color: #2980b9 !important;
    border-color: #2980b9 !important;
  }

  .custom-orange-tag {
    background-color: #e67e22 !important;
    border-color: #e67e22 !important;
  }

  .custom-purple-tag {
    background-color: #6c3483 !important;
    border-color: #6c3483 !important;
  }

  .custom-cyan-tag {
    background-color: #1abc9c !important;
    border-color: #1abc9c !important;
  }
}

.captcha-section {
  display: flex;
  align-items: center;
  gap: 16px;
}

.captcha-group {
  display: flex;
  gap: 16px;
}

.captcha-container {
  display: flex;
  align-items: center;
  gap: 8px;
}

.captcha-container img {
  height: 32px;
  cursor: pointer;
  border-radius: 4px;
}

.captcha-input {
  width: 80px;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .captcha-container img {
    border: 1px solid rgba(255, 255, 255, 0.1);
  }
}

.captcha-dialog-content {
  display: flex;
  flex-direction: column;
  gap: 16px;
  padding: 0 20px;
}

.captcha-row {
  display: flex;
  align-items: center;
  gap: 12px;
  width: 100%;
}

.captcha-row img {
  height: 32px;
  cursor: pointer;
  border-radius: 4px;
}

.captcha-input {
  flex-grow: 1;
}

.dialog-footer {
  display: flex;
  justify-content: space-between;
  padding: 0 20px;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .captcha-row img {
    border: 1px solid rgba(255, 255, 255, 0.1);
  }
}
</style>