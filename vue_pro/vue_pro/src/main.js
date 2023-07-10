import { createApp } from 'vue'
import App from './App.vue'
import axios from 'axios'
//引入element框架
import ElDialog from 'element-plus'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus';
import mitt from 'mitt'
import store from '@/utils/store.js'


// 引入路由
import router from './router/index'
// //定义全局css
// import './assets/css/global.css'

const app =createApp(App)
app.config.globalProperties.emitter = mitt()
for(const [key,component] of Object.entries(ElementPlusIconsVue)){
    app.component(key, component)
}
app.use(ElementPlus)
app.use(ElDialog)
app.use(store)
app.config.globalProperties.$message = ElMessage;
app.use(router)
app.mount('#app')

// Vue.config.productionTip = false;
//设置默认请求的路径
axios.defaults.baseURL="http://localhost:8080"
//Vue.prototype.$axios=axios