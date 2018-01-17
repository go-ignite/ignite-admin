import Vue from 'vue';
import Router from 'vue-router';
import EventBus, {Event} from './utils/EventBus'
import About from './views/About';
import Code from './views/Code';
import Index from './views/Index';
import Status from './views/Status';
import Container from './components/Container.vue';

Vue.use(Router);

const router = new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      component: Container,
      children: [
        {
          path: '',
          name: 'index',
          component: Index,
        },
        {
          path: 'about',
          name: 'about',
          component: About,
        },
        {
          path: 'code',
          name: 'code',
          component: Code,
        },
        {
          path: 'status',
          name: 'status',
          component: Status,
        },
      ],
    },
  ],
});

router.afterEach((to, from) => {
  EventBus.$emit(Event.ROUTE_CHANGE, to.name);
})

export default router;
