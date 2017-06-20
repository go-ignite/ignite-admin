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
import axios from 'axios';

export default {
    data: function () {
        return {
            username: '',
            password: ''
        }
    },
    methods: {
        onLogin: function (event) {
            console.log('Username:' + this.username + ' Password:' + this.password);

            axios.post("/login", {
                "username": this.username,
                "password": this.password
            })
            .then(function (response) {
                if(response.data.success) {
                    console.log(response.message);
                    location.href = '/panel/status';
                }
            })
            .catch(function(error){
                console.log(error);
            });
        }
    }
}
</script>

<style>
#login {
    margin: 170px auto 180px auto;
}
</style>