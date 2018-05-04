<template>
  <div class="iadmin">
    <nav-bar></nav-bar>
    <div class="iadmin_views">
      <router-view />
    </div>
    <g-footer></g-footer>
  </div>
</template>

<script>
import NavBar from '@/components/NavBar.vue'
import GFooter from '@/components/Footer.vue'

export default {
  name: 'app',
  components: {
    NavBar,
    GFooter,
  },
}
</script>
