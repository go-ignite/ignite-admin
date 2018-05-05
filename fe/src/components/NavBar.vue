<template>
  <div>
    <div class="nav_wrap">
      <img src="../assets/img/favicon-96x96.png" class="nav_logo" alt="ignite">
      <el-menu :default-active="activePath" class="nav_items" mode="horizontal" @select="handleSelect">
        <!-- use v-show for active. -->
        <el-menu-item v-show="!isLogin" index="/">首页</el-menu-item>
        <el-menu-item v-show="isLogin" index="/status">用户管理</el-menu-item>
        <el-menu-item v-show="isLogin" index="/code">邀请码管理</el-menu-item>
        <el-menu-item index="/about">关于</el-menu-item>
      </el-menu>
      <span v-if="isLogin" class="h_fr h_csp" @click="onLogout">退出</span>
    </div>
  </div>
</template>

<script>
import EventBus, { Event } from '../utils/EventBus'

export default {
  data() {
    return {
      isLogin: false,
      buggerActive: false,
    }
  },
  computed: {
    activePath() {
      return this.$route.path
    },
  },
  methods: {
    onLogout() {
      localStorage.removeItem('ignite_admin_token')
      location.href = '/'
    },
    toggleBurger() {
      this.buggerActive = !this.buggerActive
    },
    handleSelect(index) {
      this.$router.push({
        path: index,
      })
    },
  },
  created() {
    EventBus.$on(Event.LOGIN_SUCCESS, () => {
      this.isLogin = true
    })
    if (localStorage.getItem('ignite_admin_token')) {
      this.isLogin = true
    }
  },
}
</script>

<style lang="less">
.nav_logo {
  height: 50px;
  margin-right: 50px;
  cursor: pointer;
}
.nav_wrap {
  padding: 0 20px;
  display: flex;
  align-items: center;
  border-bottom: 1px solid #e6e6e6;
}
.nav_items {
  flex: 1;
  border-bottom: none;
}
</style>
>
.nav-l
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
