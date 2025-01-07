<template>
  <div class="tjkc-container">
    <h1>培养方案内课程</h1>
    <div v-if="loading" class="loading">
      <el-loading :active="loading" />
    </div>
    <div v-else-if="error" class="error">
      {{ error }}
    </div>
    <div v-else class="json-container">
      <pre>{{ JSON.stringify(processedData, null, 2) }}</pre>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useGlobalStore } from '../../stores/globalStore'
import { ElMessage } from 'element-plus'

const globalStore = useGlobalStore()
const responseData = ref(null)
const loading = ref(false)
const error = ref(null)

// 处理数据的计算属性
const processedData = computed(() => {
  if (!responseData.value || !responseData.value.data || !responseData.value.data.rows) {
    return []
  }

  const result = []
  // 遍历rows中的每个课程
  for (const course of responseData.value.data.rows) {
    // 遍历每个课程的教学班列表
    for (const tc of course.tcList) {
      // 提取所需信息
      const courseInfo = {
        JXBID: tc.JXBID,
        KCM: tc.KCM,
        SKJS: tc.SKJS,
        secretVal: tc.secretVal,
        XF: tc.XF
      }
      // 将信息添加到结果列表中
      result.push(courseInfo)
    }
  }
  return result
})

// 获取培养方案内课程
const getTJKC = async () => {
  loading.value = true
  error.value = null
  
  try {
    if (!globalStore.Authorization || !globalStore.batchId) {
      throw new Error('缺少必要的认证信息')
    }

    const data = await window.go.TJKC.App.GetTJKC(globalStore.Authorization, globalStore.batchId)
    if (data.code !== 200) {
      throw new Error(data.msg || '获取课程列表失败')
    }
    responseData.value = data
  } catch (err) {
    error.value = err.message
    ElMessage.error(err.message)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  getTJKC()
})
</script>

<style scoped>
.tjkc-container {
  padding: 20px;
  max-width: 1200px;
  margin: 0 auto;
}

.loading {
  text-align: center;
  padding: 40px;
}

.error {
  color: #f56c6c;
  text-align: center;
  padding: 20px;
}

h1 {
  margin-bottom: 20px;
}

.json-container {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border-radius: 8px;
  padding: 20px;
  overflow-x: auto;
}

pre {
  margin: 0;
  white-space: pre-wrap;
  word-wrap: break-word;
  font-family: monospace;
  font-size: 14px;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .json-container {
    background: rgba(255, 255, 255, 0.05);
  }
  
  pre {
    color: #ddd;
  }
}
</style>
