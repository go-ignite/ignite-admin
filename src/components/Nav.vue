<template>
    <nav class="nav has-shadow navbar-fixed-top">
        <div class="container">
            <div class="nav-left">
                <a class="nav-item">
        <img src="/static/images/favicon-96x96.png" alt="ignite">
      </a>
                <a class="nav-item is-tab is-hidden-mobile is-active">首页</a>
                <a class="nav-item is-tab is-hidden-mobile">关于</a>
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
  }
}
</script>

<style>
.nav {
    position:fixed;
    left:0px;
    top:0px;
    width:100%;
}
</style>