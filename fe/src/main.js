import Vue from 'vue'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'
import 'mdi/css/materialdesignicons.min.css'
import './assets/css/style.less'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'

import App from './App.vue'
import router from './routes'

Vue.use(Buefy)
Vue.use(ElementUI)

new Vue({
  router,
  render: (h) => h(App),
}).$mount('#app')
