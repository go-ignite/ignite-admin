import Vue from 'vue/dist/vue.js'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'
import './css/style.css'
import './css/font-awesome.min.css'
import './css/material-icons.css'
import App from './views/Code.vue';


Vue.use(Buefy)
new Vue({
  el: '#app',
  template: '<App />',
  components: { App },
});
