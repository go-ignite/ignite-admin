<template>
    <nav class="navbar is-fixed-top">
        <div class="container">
            <div class="navbar-brand">
                <a class="navbar-item">
                    <img src="../assets/images/favicon-96x96.png" alt="ignite">
                </a>
                <div class="navbar-burger" @click="toggleBurger">
                    <span></span>
                    <span></span>
                    <span></span>
                </div>
            </div>
            <div class="navbar-menu container" :class="{'is-active': buggerActive}">
                <div class="navbar-start">
                    <template v-if="isLogin">
                        <router-link class="navbar-item is-tab" :to="{name: 'status'}"
                            :class="{'is-active': currentRoute === 'status'}">用户管理</router-link>
                        <router-link class="navbar-item is-tab" :to="{name: 'code'}"
                            :class="{'is-active': currentRoute === 'code'}">邀请码管理</router-link>
                    </template>
                    <template v-else>
                        <router-link class="navbar-item is-tab" :class="{'is-active': currentRoute === 'index'}"
                        :to="{name: 'index'}">首页</router-link>
                    </template>
        
                    <router-link class="navbar-item is-tab" :to="{name: 'about'}"
                        :class="{'is-active': currentRoute === 'about'}">关于</router-link>
                </div>
                <div class="navbar-end">
                    <span v-if="isLogin" class="navbar-item is-tab" @click="onLogout">退出</span>
                </div>
            </div>
        </div>
    </nav>
</template>

<script>
import EventBus, { Event } from '../utils/EventBus'

export default {
  data() {
    return {
      isLogin: false,
      currentRoute: '',
      buggerActive: false,
    }
  },
  methods: {
    onLogout() {
      localStorage.setItem('token', '')
      location.href = '/'
    },
    changeNavActive(routeName) {
      this.currentRoute = routeName
    },
    toggleBurger() {
      this.buggerActive = !this.buggerActive
    },
  },
  created() {
    if (localStorage.getItem('token') != '') {
      this.isLogin = true
    }
  },
  mounted() {
    this.changeNavActive(this.$router.currentRoute.name)
    EventBus.$on(Event.LOGIN_SUCCESS, () => {
      this.isLogin = true
    })
    EventBus.$on(Event.ROUTE_CHANGE, this.changeNavActive)
  },
}
</script>

<style scoped>
.navbar {
  box-shadow: 0 2px 3px hsla(0, 0%, 4%, 0.1);
}

.nav-item {
  display: flex;
  align-items: center;
  cursor: pointer;
}
.nav-item img {
  max-height: 1.75rem;
}
@media screen and (max-width: 1023px) {
  .navbar-menu .navbar-item.is-tab.is-active {
    border-bottom: none;
  }
}
</style>
