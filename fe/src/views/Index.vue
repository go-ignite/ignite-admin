<template>
  <div class="container" id="main-container">
      <div class="columns">
          <div class="column is-4 is-offset-4">
              <login-view></login-view>
          </div>
      </div>
  </div>
</template>
<script>
import LoginView from '../components/Login.vue'

export default {
  components: {
    LoginView,
  },
};
</script>
