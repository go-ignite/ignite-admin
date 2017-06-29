<template>
    <nav class="nav has-shadow navbar-fixed-top">
        <div class="container">
            <div class="nav-left">
                <a class="nav-item">
                    <img src="./images/favicon-96x96.png" alt="ignite">
                </a>

                <a v-if="page == 0" class="nav-item is-tab is-active" href="/status">首页</a>

                <a v-if="page == 1" class="nav-item is-tab is-active" href="/status">用户管理</a>
                <a v-if="page !=1 && page != 0" class="nav-item is-tab" href="/status">用户管理</a>

                <a v-if="page == 2" class="nav-item is-tab is-active" href="/code">邀请码管理</a>
                <a v-if="page != 2 && page != 0" class="nav-item is-tab" href="/code">邀请码管理</a>
                
                <a v-if="page == 3" class="nav-item is-tab is-active">关于</a>
                <a v-else class="nav-item is-tab">关于</a>
            </div>
            <div class="nav-right">
                <a class="nav-item is-tab" @click="onLogout">退出</a>
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
export default {
    data: function () {
        return {}
    },
    props: ['page'],
    methods: {
        onLogout(event) {
            console.log('logout clicked...');
            localStorage.setItem("token", "");
            location.href = '/';
        }
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