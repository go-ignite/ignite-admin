import Vue from 'vue/dist/vue.js'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'
import './css/style.css'
import Nav from './components/Nav.vue'
import Login from './components/Login.vue'
import Footer from './components/Footer.vue'


Vue.use(Buefy)

new Vue({
    el: '#nav',
    components: { 
        'nav-bar' : Nav 
    }
})

new Vue({
    el: '#login',
    components: { 
        'login' : Login 
    }
})

new Vue({
    el: '#footer',
    components: { 
        'foot' : Footer 
    }
})