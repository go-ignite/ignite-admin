import Vue from 'vue';
import Buefy from 'buefy';
import 'buefy/lib/buefy.css';
import './css/style.css';
import './css/font-awesome.min.css';
import './css/material-icons.css';

import App from './App.vue';
import router from './routes'

Vue.use(Buefy);

new Vue({
  el: '#app',
  render: h => h(App),
  router,
});
