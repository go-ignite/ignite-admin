import Vue from 'vue/dist/vue.js'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'
import '../src/style.css'
import Nav from '../src/components/Nav.vue'
import Login from '../src/components/Login.vue'
import Footer from '../src/components/Footer.vue'


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