<template>
  <div>
    <g-navbar></g-navbar>
    <router-view />
    <g-footer></g-footer>
  </div>
</template>

<script>
import GNavbar from './Nav.vue';
import GFooter from './Footer.vue';

export default {
  name: 'app',
  components: {
    GNavbar,
    GFooter,
  }
};
</script>
