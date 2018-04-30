<template>
  <div class="iadmin">
    <g-navbar></g-navbar>
    <div class="iadmin_views">
      <router-view />
    </div>
    <g-footer></g-footer>
  </div>
</template>

<script>
import GNavbar from '@/components/Nav.vue'
import GFooter from '@/components/Footer.vue'

export default {
  name: 'app',
  components: {
    GNavbar,
    GFooter,
  },
}
</script>
