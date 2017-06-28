<template>
    <nav class="nav has-shadow navbar-fixed-top">
        <div class="container">
            <div class="nav-left">
                <a class="nav-item">
                    <img src="./images/favicon-96x96.png" alt="ignite">
                </a>
                <a class="nav-item is-tab is-active">首页</a>
                <a class="nav-item is-tab">关于</a>
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