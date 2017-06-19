<template>
    <div>
        <h1 class="title">User Login</h1>
        <b-field label="Name">
            <b-input v-model="username"></b-input>
        </b-field>
        <b-field label="Password">
            <b-input type="password" password-reveal v-model="password">
            </b-input>
        </b-field>
        <a class="button is-primary" @click="onLogin">Login</a>
    </div>
</template>

<script>
module.exports = {
    data: function () {
        return {
            username: '',
            password: ''
        }
    },
    methods: {
        onLogin: function (event) {
            console.log('Username:' + this.username + ' Password:' + this.password);
        }
    }
}
</script>

<style>
#login {
    margin: 170px auto 180px auto;
}
</style>