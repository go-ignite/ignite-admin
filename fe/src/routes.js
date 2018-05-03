import Vue from 'vue'
import Router from 'vue-router'
import EventBus, { Event } from './utils/EventBus'

Vue.use(Router)

const router = new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      component: () => import('./views/Entry'),
      children: [
        {
          path: '',
          name: 'login',
          component: () => import('./views/Login'),
        },
        {
          path: 'about',
          name: 'about',
          component: () => import('./views/AboutUs'),
        },
        {
          path: 'code',
          name: 'code',
          component: () => import('./views/CodeTable'),
        },
        {
          path: 'status',
          name: 'status',
          component: () => import('./views/UserTable'),
        },
      ],
    },
  ],
})

router.afterEach((to, from) => {
  EventBus.$emit(Event.ROUTE_CHANGE, to.name)
})

export default router
