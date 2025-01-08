<template>
  <div class="xgkc-container">
    <h1>公选课</h1>
    <div v-if="error" class="error">
      {{ error }}
    </div>
    <div class="table-container">
      <el-table 
        v-loading="loading"
        :data="processedData" 
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
        />
        <el-table-column 
          prop="KCM" 
          label="课程名称" 
          min-width="200" 
          align="center"
          header-align="center"
        />
        <el-table-column 
          prop="SKJS" 
          label="授课教师" 
          min-width="120" 
          align="center"
          header-align="center"
        />
        <el-table-column 
          prop="XF" 
          label="学分" 
          width="80" 
          align="center"
          header-align="center"
        />
        <el-table-column 
          prop="secretVal" 
          label="密钥" 
          min-width="150" 
          align="left"
          header-align="center"
          show-overflow-tooltip
        >
          <template #default="scope">
            <el-tooltip 
              :content="scope.row.secretVal" 
              placement="top" 
              :show-after="500"
            >
              <span class="secret-text">{{ scope.row.secretVal.slice(0, 20) }}...</span>
            </el-tooltip>
          </template>
        </el-table-column>
        <el-table-column 
          label="操作" 
          width="120" 
          align="center"
          header-align="center"
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
import { ref, onMounted, computed } from 'vue'
import { useGlobalStore } from '../../stores/globalStore'
import { useCourseStore } from '../../stores/courseStore'
import { ElMessage } from 'element-plus'

const globalStore = useGlobalStore()
const courseStore = useCourseStore()
const loading = ref(false)
const error = ref(null)

// 选课处理函数
const handleSelect = async (course) => {
  try {
    if (!globalStore.Authorization || !globalStore.batchId) {
      throw new Error('缺少必要的认证信息')
    }

    const result = await window.go.only9464.App.AddClazz(
      globalStore.Authorization,
      globalStore.batchId,
      'XGKC',  // 这里改为 XGKC
      course.JXBID,
      course.secretVal,
      '1'
    )

    const jsonResult = JSON.parse(result)
    if (jsonResult.code === 200) {
      ElMessage.success(jsonResult.msg || '选课成功')
      // 更新课程状态
      course.SFYX = '1'
    } else {
      throw new Error(jsonResult.msg || '选课失败')
    }
  } catch (err) {
    ElMessage.error(err.message)
  }
}

// 退选处理函数
const handleUnselect = (course) => {
  ElMessage.info('退选功能开发中...')
}

// 处理数据的计算属性
const processedData = computed(() => {
  if (!courseStore.xgkcData || !courseStore.xgkcData.data || !courseStore.xgkcData.data.rows) {
    return []
  }

  const result = []
  // 遍历rows中的每个课程
  for (const course of courseStore.xgkcData.data.rows) {
    // 提取所需信息
    const courseInfo = {
      JXBID: course.JXBID,
      KCM: course.KCM,
      SKJS: course.SKJS,
      secretVal: course.secretVal,
      XF: course.XF,
      SFYX: course.SFYX
    }
    // 将信息添加到结果列表中
    result.push(courseInfo)
  }
  return result
})

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

onMounted(() => {
  getXGKC()
})
</script>

<style scoped>
/* 样式部分与其他组件相同 */
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
}
</style>
