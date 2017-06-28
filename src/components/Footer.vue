<template>
    <footer class="footer">
        <div class="container">
            <div class="content has-text-centered">
                <p>
                    <strong>ignite-admin</strong> by <a href="http://github.com/go-ignite">go-ignite</a>. The source code is licensed
                    <a href="http://opensource.org/licenses/mit-license.php">MIT</a>.
                </p>
                <p>
                    <a class="icon" href="https://github.com/jgthms/bulma">
                        <i class="fa fa-github"></i>
                    </a>
                </p>
            </div>
        </div>
    </footer>
</template>

<script>
export default {
    data: function () {
        return {}
    }
}
</script>

<style scoped>
.footer {
    position: fixed;
    left: 0px;
    bottom: 0px;
    height: 30px;
    width: 100%;
}
</style>