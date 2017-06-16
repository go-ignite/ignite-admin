import Vue from 'vue/dist/vue.js'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'
import '../src/style.css'
import UserTable from '../src/components/UserTable.vue'
import Footer from '../src/components/Footer.vue'


Vue.use(Buefy)

var ut = new Vue({
    el: '#ut',
    components: { 
        'user-table' : UserTable 
    }
})

var footer = new Vue({
    el: '#footer',
    components: { 
        'ft' : Footer 
    }
})