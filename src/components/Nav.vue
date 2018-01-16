<template>
    <nav class="nav has-shadow navbar-fixed-top">
        <div class="container">
            <div class="nav-left">
                <a class="nav-item">
                    <img src="./images/favicon-96x96.png" alt="ignite">
                </a>
    
                <template v-if="isLogin">
                    <router-link class="nav-item is-tab" :to="{name: 'status'}"
                        :class="{'is-active': page === 1}">用户管理</router-link>
                    <router-link class="nav-item is-tab" :to="{name: 'code'}"
                        :class="{'is-active': page === 2}">邀请码管理</router-link>
                </template>
                <template v-else>
                    <router-link class="nav-item is-tab" :class="{'is-active': page === 0}"
                      :to="{name: 'index'}">首页</router-link>
                </template>
    
                <router-link class="nav-item is-tab" :to="{name: 'about'}"
                    :class="{'is-active': page === 3}">关于</router-link>
            </div>
            <div class="nav-right">
                <span v-if="isLogin" class="nav-item is-tab" @click="onLogout">退出</span>
            </div>
            <span class="nav-toggle">
                <span></span>
                <span></span>
                <span></span>
            </span>
        </div>
    </nav>
</template>

<script>
import EventBus from '../utils/EventBus';

export default {
    data: function () {
        return {
            isLogin: false
        }
    },
    props: ['page'],
    methods: {
        onLogout(event) {
            console.log('logout clicked...');
            localStorage.setItem("token", "");
            location.href = '/';
        }
    },
    created() {
        if (localStorage.getItem("token") != "") {
            this.isLogin = true;
        }
    },
    mounted() {
        EventBus.$on('login-success', () => {
            this.isLogin = true;
        })
    }
}
</script>

<style scoped>
.nav {
    position: fixed;
    left: 0px;
    top: 0px;
    width: 100%;
}
</style>
