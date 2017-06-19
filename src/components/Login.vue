<template>
    <div>
        <h1 class="title">User Login</h1>
        <b-field label="Name">
            <b-input value="John Silver"></b-input>
        </b-field>
        <b-field label="Password">
            <b-input type="password" password-reveal>
            </b-input>
        </b-field>
    </div>
</template>

<script>
module.exports = {
    data: function () {
        return {}
    }
}
</script>

<style>
#login {
  margin: 180px auto 180px auto;
}
</style>