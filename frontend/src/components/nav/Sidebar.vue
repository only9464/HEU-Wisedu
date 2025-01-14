<template>
  <div class="sidebar-container">
    <div class="sidebar acrylic-effect">
      <div class="menu-container">
        <ul>

          <li>
            <router-link to="/" class="nav-link">
              <span>首页-登录</span>
            </router-link>
          </li>

          <li>
            <router-link to="/TJKC" class="nav-link">
              <span>培养方案内课程</span>
            </router-link>
          </li>
          <li>
            <router-link to="/FAWKC" class="nav-link">
              <span>跨专业选修课</span>
            </router-link>
          </li>
          <li>
            <router-link to="/XGKC" class="nav-link">
              <span>公选课</span>
            </router-link>
          </li>
          <li>
            <router-link to="/YXKCYX" class="nav-link">
              <span>已选课程</span>
            </router-link>
          </li>
          <li>
            <router-link to="/AllThePast" class="nav-link">
              <span>所有学期已修课程</span>
            </router-link>
          </li>
          <li>
            <router-link to="/only9464" class="nav-link">
              <span>选(qiang)课助手</span>
            </router-link>
          </li>

          <!-- <li>
            <router-link to="/jsfinder" class="nav-link">
              <span>JSFinder</span>
            </router-link>
          </li> -->
        </ul>
      </div>
      <div class="bottom-menu">
        <ul>

          <li>
            <router-link to="/setting" class="nav-link" :class="{ 'has-update': hasUpdate }">
              <span>
                设置
                <el-tag v-if="hasUpdate" size="small" type="danger" effect="dark" class="update-tag">
                  有更新
                </el-tag>
              </span>
            </router-link>
          </li>

        </ul>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useGlobalStore } from '../../stores/globalStore'
import { storeToRefs } from 'pinia'

const globalStore = useGlobalStore()
const hasUpdate = computed(() => !globalStore.isLatestVersion)

onMounted(() => {
  globalStore.checkUpdate()
})
</script>

<style scoped>
.sidebar-container {
  position: fixed;
  top: 20px;
  left: 20px;
  bottom: 20px;
  width: 200px;
  z-index: 100;
}

.sidebar {
  height: 100%;
  border-radius: 16px;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

.menu-container {
  padding-top: 20px;
  flex-grow: 1;
}

.bottom-menu {
  padding-bottom: 20px;
}

.acrylic-effect {
  background: linear-gradient(
    135deg,
    rgba(255, 255, 255, 0.15),
    rgba(255, 255, 255, 0.08)
  );
  backdrop-filter: blur(20px) saturate(125%);
  -webkit-backdrop-filter: blur(20px) saturate(125%);
  border: 1px solid rgba(255, 255, 255, 0.1);
  box-shadow: 0 8px 32px 0 rgba(0, 0, 0, 0.15);
}

.sidebar ul {
  list-style-type: none;
  padding: 0;
  margin: 0;
}

.sidebar li {
  margin: 8px 0;
}

.nav-link {
  color: #000000;
  text-decoration: none;
  padding: 12px 24px;
  display: block;
  position: relative;
  transition: all 0.3s ease;
  margin: 0 8px;
  border-radius: 8px;
  font-weight: 600;
  letter-spacing: 0.5px;
  text-shadow: 0 1px 1px rgba(255, 255, 255, 0.3);
  background: linear-gradient(
    90deg,
    rgba(255, 255, 255, 0.3),
    rgba(255, 255, 255, 0.2)
  );
}

.nav-link:hover {
  background: linear-gradient(
    90deg,
    rgba(255, 255, 255, 0.4),
    rgba(255, 255, 255, 0.3)
  );
  color: #1e8fa3;
  text-shadow: 0 1px 2px rgba(255, 255, 255, 0.1);
}

.nav-link:active {
  background: rgba(255, 255, 255, 0.3);
}

.router-link-active {
  background: linear-gradient(
    90deg,
    rgba(255, 255, 255, 0.5),
    rgba(255, 255, 255, 0.4)
  );
  color: #1e8fa3;
  font-weight: 700;
  text-shadow: 0 1px 2px rgba(255, 255, 255, 0.01);
}

.nav-link::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 8px;
  opacity: 0;
  transition: opacity 0.3s ease;
}

.nav-link:hover::before,
.router-link-active::before {
  opacity: 1;
}

.nav-link span {
  position: relative;
  z-index: 1;
}

.has-update {
  position: relative;
  animation: pulse 2s infinite;
  background: linear-gradient(
    90deg,
    rgba(255, 68, 68, 0.2),
    rgba(255, 68, 68, 0.1)
  ) !important;
}

.update-tag {
  margin-left: 8px;
  font-size: 12px;
}

@keyframes pulse {
  0% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.02);
  }
  100% {
    transform: scale(1);
  }
}

@media (prefers-color-scheme: dark) {
  .acrylic-effect {
    background: linear-gradient(
      135deg,
      rgba(255, 255, 255, 0.1),
      rgba(255, 255, 255, 0.05)
    );
  }
  
  .nav-link {
    color: rgba(255, 255, 255, 0.9);
    text-shadow: 0 2px 4px rgba(0, 0, 0, 0.5);
    background: linear-gradient(
      90deg,
      rgba(255, 255, 255, 0.15),
      rgba(255, 255, 255, 0.1)
    );
  }

  .nav-link:hover,
  .router-link-active {
    color: #ffffff;
  }

  .has-update {
    background: linear-gradient(
      90deg,
      rgba(255, 68, 68, 0.3),
      rgba(255, 68, 68, 0.2)
    ) !important;
  }
}
</style>