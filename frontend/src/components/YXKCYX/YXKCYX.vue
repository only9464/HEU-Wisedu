<template>
  <div class="yxkc-container">
    <h1>已选课程</h1>
    <div v-if="error" class="error">
      {{ error }}
    </div>
    <div class="tables-container">
      <!-- 公选课表格 -->
      <div class="table-section">
        <h2>公选课</h2>
        <el-table 
          v-loading="loading"
          :data="xgkcSelected" 
          style="width: 100%"
          :stripe="true"
          class="acrylic-effect"
        >
          <el-table-column type="index" label="序号" width="60" align="center" header-align="center"/>
          <el-table-column prop="JXBID" label="教学班ID" min-width="180" align="center" header-align="center" show-overflow-tooltip/>
          <el-table-column prop="XGXKLB" label="类型" min-width="140" align="center" header-align="center">
            <template #default="scope">
              <div class="multi-line-cell">{{ scope.row.XGXKLB }}</div>
            </template>
          </el-table-column>
          <el-table-column prop="KCM" label="课程名称" min-width="200" align="center" header-align="center" show-overflow-tooltip/>
          <el-table-column prop="SKJS" label="授课教师" min-width="120" align="center" header-align="center">
            <template #default="scope">
              <div class="multi-line-cell">{{ scope.row.SKJS }}</div>
            </template>
          </el-table-column>
          <el-table-column prop="XF" label="学分" width="80" align="center" header-align="center"/>
          <el-table-column label="操作" width="120" align="center" header-align="center">
            <template #default="scope">
              <el-button type="danger" size="small" @click="handleUnselect(scope.row, 'XGKC')">
                退选
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>

      <!-- 推荐课程表格 -->
      <div class="table-section">
        <h2>培养方案内课程</h2>
        <el-table 
          v-loading="loading"
          :data="tjkcSelected" 
          style="width: 100%"
          :stripe="true"
          class="acrylic-effect"
        >
          <el-table-column type="index" label="序号" width="60" align="center" header-align="center"/>
          <el-table-column prop="JXBID" label="教学班ID" min-width="180" align="center" header-align="center" show-overflow-tooltip/>
          <el-table-column prop="KCM" label="课程名称" min-width="200" align="center" header-align="center" show-overflow-tooltip/>
          <el-table-column prop="SKJS" label="授课教师" min-width="120" align="center" header-align="center">
            <template #default="scope">
              <div class="multi-line-cell">{{ scope.row.SKJS }}</div>
            </template>
          </el-table-column>
          <el-table-column prop="XF" label="学分" width="80" align="center" header-align="center"/>
          <el-table-column label="操作" width="120" align="center" header-align="center">
            <template #default="scope">
              <el-button type="danger" size="small" @click="handleUnselect(scope.row, 'TJKC')">
                退选
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>

      <!-- 方案外课程表格 -->
      <div class="table-section">
        <h2>跨专业选修课程</h2>
        <el-table 
          v-loading="loading"
          :data="fawkcSelected" 
          style="width: 100%"
          :stripe="true"
          class="acrylic-effect"
        >
          <el-table-column type="index" label="序号" width="60" align="center" header-align="center"/>
          <el-table-column prop="JXBID" label="教学班ID" min-width="180" align="center" header-align="center" show-overflow-tooltip/>
          <el-table-column prop="KCM" label="课程名称" min-width="200" align="center" header-align="center" show-overflow-tooltip/>
          <el-table-column prop="SKJS" label="授课教师" min-width="120" align="center" header-align="center">
            <template #default="scope">
              <div class="multi-line-cell">{{ scope.row.SKJS }}</div>
            </template>
          </el-table-column>
          <el-table-column prop="XF" label="学分" width="80" align="center" header-align="center"/>
          <el-table-column label="操作" width="120" align="center" header-align="center">
            <template #default="scope">
              <el-button type="danger" size="small" @click="handleUnselect(scope.row, 'FAWKC')">
                退选
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
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

// 分别计算三种类型的已选课程
const xgkcSelected = computed(() => {
  if (!courseStore.xgkcData?.data?.rows) return []
  return courseStore.xgkcData.data.rows.filter(course => course.SFYX === '1')
})

const tjkcSelected = computed(() => {
  if (!courseStore.tjkcData?.data?.rows) return []
  const result = []
  for (const course of courseStore.tjkcData.data.rows) {
    for (const tc of course.tcList) {
      if (tc.SFYX === '1') {
        result.push(tc)
      }
    }
  }
  return result
})

const fawkcSelected = computed(() => {
  if (!courseStore.fawkcData?.data?.rows) return []
  const result = []
  for (const course of courseStore.fawkcData.data.rows) {
    for (const tc of course.tcList) {
      if (tc.SFYX === '1') {
        result.push(tc)
      }
    }
  }
  return result
})

// 修改退选处理函数
const handleUnselect = async (course, type) => {
  try {
    if (!globalStore.Authorization || !globalStore.batchId) {
      throw new Error('缺少必要的认证信息')
    }

    const result = await window.go.only9464.App.DelClazz(
      globalStore.Authorization,
      globalStore.batchId,
      type,  // 使用传入的课程类型
      course.JXBID,
      course.secretVal
    )

    const jsonResult = JSON.parse(result)
    if (jsonResult.code === 200) {
      ElMessage.success(jsonResult.msg || '退选成功')
      // 根据类型更新对应的数据
      switch (type) {
        case 'XGKC':
          if (courseStore.xgkcData?.data?.rows) {
            const courseIndex = courseStore.xgkcData.data.rows.findIndex(c => c.JXBID === course.JXBID)
            if (courseIndex !== -1) {
              courseStore.xgkcData.data.rows[courseIndex].SFYX = '0'
            }
          }
          break
        case 'TJKC':
          if (courseStore.tjkcData?.data?.rows) {
            for (const row of courseStore.tjkcData.data.rows) {
              for (const tc of row.tcList) {
                if (tc.JXBID === course.JXBID) {
                  tc.SFYX = '0'
                  break
                }
              }
            }
          }
          break
        case 'FAWKC':
          if (courseStore.fawkcData?.data?.rows) {
            for (const row of courseStore.fawkcData.data.rows) {
              for (const tc of row.tcList) {
                if (tc.JXBID === course.JXBID) {
                  tc.SFYX = '0'
                  break
                }
              }
            }
          }
          break
      }
    } else {
      throw new Error(jsonResult.msg || '退选失败')
    }
  } catch (err) {
    ElMessage.error(err.message)
  }
}

// 刷新处理函数
const handleRefresh = async () => {
  // 清除所有已加载标记
  courseStore.xgkcLoaded = false
  courseStore.tjkcLoaded = false
  courseStore.fawkcLoaded = false
  
  // 重新获取所有数据
  await Promise.all([
    window.go.XGKC.App.GetXGKC(globalStore.Authorization, globalStore.batchId),
    window.go.TJKC.App.GetTJKC(globalStore.Authorization, globalStore.batchId),
    window.go.FAWKC.App.GetFAWKC(globalStore.Authorization, globalStore.batchId)
  ]).catch(err => {
    error.value = err.message
    ElMessage.error(err.message)
  })
}

onMounted(() => {
  handleRefresh()
})
</script>

<style scoped>
.yxkc-container {
  height: 100%;
  display: flex;
  flex-direction: column;
  padding: 0 20px;
  gap: 20px;
  overflow: auto;  /* 改为 auto 允许滚动 */
}

.tables-container {
  display: flex;
  flex-direction: column;
  gap: 30px;
  padding: 20px;
}

.table-section {
  background: var(--el-bg-color);
  border-radius: 8px;
  padding: 20px;
}

h2 {
  margin-bottom: 20px;
  font-size: 1.2em;
  color: var(--el-text-color-primary);
}

/* ... 其他样式保持不变 ... */
</style>
