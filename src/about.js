import Vue from 'vue/dist/vue.js'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'
import './css/style.css'
import './css/font-awesome.min.css'
import Nav from '../src/components/Nav.vue'
import Profile from '../src/components/Profile.vue'
import Footer from '../src/components/Footer.vue'


Vue.use(Buefy)

new Vue({
    el: '#nav',
    components: { 
        'nav-bar' : Nav 
    }
})

new Vue({
    el: '#profile',
    components: { 
        'profile' : Profile 
    }
})

new Vue({
    el: '#footer',
    components: { 
        'foot' : Footer 
    }
})
