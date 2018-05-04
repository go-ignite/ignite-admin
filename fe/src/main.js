import Vue from 'vue'
import './assets/css/style.less'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'

import App from './App.vue'
import router from './routes'

Vue.use(ElementUI)

new Vue({
  router,
  render: (h) => h(App),
}).$mount('#app')
