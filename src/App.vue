<template>
  <div>
    <g-navbar></g-navbar>
    <router-view />
    <g-footer></g-footer>
  </div>
</template>

<script>
import GNavbar from './components/Nav.vue';
import GFooter from './components/Footer.vue';

export default {
  name: 'app',
  components: {
    GNavbar,
    GFooter,
  }
};
</script>
