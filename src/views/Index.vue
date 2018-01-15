<template>
  <div>
    <g-navbar v-bind:page="0"></g-navbar>

    <div class="container" id="main-container">
        <div class="columns">
            <div class="column is-4 is-offset-4">
                <login-view></login-view>
            </div>
        </div>
    </div>

    <g-footer></g-footer>
  </div>

</template>
<script>
import GNavbar from '../components/Nav.vue';
import GFooter from '../components/Footer.vue';
import LoginView from '../components/Login.vue'

export default {
  components: {
    GNavbar,
    GFooter,
    LoginView,
  },
};
</script>
