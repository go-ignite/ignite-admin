<template>
    <div>
        <h1 class="title">管理登录</h1>
        <b-field label="用户名">
            <b-input v-model="username" @keyup.enter.native="onLogin"></b-input>
        </b-field>
        <b-field label="密码">
            <b-input type="password" password-reveal v-model="password" @keyup.enter.native="onLogin">
            </b-input>
        </b-field>
        <a class="button is-primary" @click="onLogin">登入</a>
    </div>
</template>

<script>
import request from '../apis/request'
import EventBus, { Event } from '../utils/EventBus'

export default {
  data() {
    return {
      username: '',
      password: '',
    }
  },
  methods: {
    onLogin(event) {
      request
        .post('/api/login', {
          username: this.username,
          password: this.password,
        })
        .then((response) => {
          if (response.success) {
            localStorage.setItem('token', response.data)
            EventBus.$emit(Event.LOGIN_SUCCESS)
            this.$router.push({ name: 'status' })
          } else {
            this.$toast.open({
              message: '用户名或者密码不正确!',
              type: 'is-danger',
            })
          }
        })
    },
  },
}
</script>

<style scoped>
</style>
