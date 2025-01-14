import { useGlobalStore } from '../stores/globalStore'

export function updateGlobalSettings() {
  const store = useGlobalStore()
  store.setAppVersion('1.0.0')
} 