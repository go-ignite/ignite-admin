import Vue from 'vue';
import Buefy from 'buefy';
import 'buefy/lib/buefy.css';
import 'mdi/css/materialdesignicons.min.css'
import './css/style.css';

import App from './App.vue';
import router from './routes'

Vue.use(Buefy);

new Vue({
  el: '#app',
  render: h => h(App),
  router,
});
