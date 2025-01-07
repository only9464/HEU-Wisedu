<template>
  <div class="First-container">
    <h1>Acrylic第1个界面</h1>
    <!-- 输入框，用于接受用户输入的两个整数 -->
    <div>
      <input type="number" v-model.number="numberA" placeholder="Enter number A" />
      <input type="number" v-model.number="numberB" placeholder="Enter number B" />
    </div>
    <!-- 按钮，点击后执行计算 -->
    <button @click="calculateSum">计算结果 乘法</button>
    <!-- 展示计算结果 -->
    <div v-if="result !== null">
      <p>Result: {{ result }}</p>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

// 使用 ref 创建响应式变量，用于存储输入数据和计算结果
const numberA = ref(0)
const numberB = ref(0)
const result = ref(null)

// 定义方法，调用 Go 后端的 First 方法并保存结果
async function calculateSum() {
  try {
    result.value = await window.go.First.App.First(numberA.value, numberB.value)
    console.log("Result from First:", result.value)
  } catch (error) {
    console.error("Failed to call First:", error)
  }
}
</script>

<style scoped>
.First-container {
  max-width: 400px;
  margin: 0 auto;
  text-align: center;
  padding: 20px;
}

input {
  width: 80px;
  margin: 5px;
}

button {
  margin-top: 10px;
}

p {
  margin-top: 10px;
  font-size: 18px;
}
</style>
