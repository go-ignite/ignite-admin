import Vue from 'vue/dist/vue.js'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'
import './css/style.css'
import Nav from '../src/components/Nav.vue'
import CodeTable from '../src/components/CodeTable.vue'
import Footer from '../src/components/Footer.vue'


Vue.use(Buefy)

new Vue({
    el: '#nav',
    components: { 
        'nav-bar' : Nav 
    }
})

new Vue({
    el: '#ct',
    components: { 
        'code-table' : CodeTable 
    }
})

new Vue({
    el: '#footer',
    components: { 
        'foot' : Footer 
    }
})