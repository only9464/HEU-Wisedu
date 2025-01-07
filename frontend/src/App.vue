<script setup>
import Sidebar from './components/nav/Sidebar.vue'
</script>

<template>
  <div class="app-container">
    <Sidebar />
    <div class="main-content">
      <div class="content">
        <div class="content-wrapper">
          <router-view></router-view>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.app-container {
  display: flex;
  width: 100%;
  height: 100vh;
  overflow: hidden;
}

.main-content {
  flex: 1;
  margin-left: 240px; /* 调整左侧间距，200px + 40px */
  min-width: 0;
  padding: 20px 20px 20px 0; /* 上、右、下、左的内边距 */
}

.content {
  height: 100%;
  overflow-y: auto;
  width: 100%;
  padding-right: 20px; /* 右侧内边距 */
}

.content-wrapper {
  max-width: 1200px;
  margin: 0 auto;
  height: 100%;
  display: flex;
  flex-direction: column;
}

/* 自定义滚动条样式 */
.content::-webkit-scrollbar {
  width: 8px;
}

.content::-webkit-scrollbar-track {
  background: transparent;
}

.content::-webkit-scrollbar-thumb {
  background: rgba(255, 255, 255, 0);
  border-radius: 4px;
}

.content::-webkit-scrollbar-thumb:hover {
  background: rgba(255, 255, 255, 0);
}
</style>