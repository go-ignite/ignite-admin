<template>
  <div class="container" id="profile">
      <div class="columns">
          <div class="column is-4 is-offset-4">
              <about-us></about-us>
          </div>
      </div>
  </div>
</template>
<script>
import AboutUs from '../components/AboutUs.vue'

export default {
  components: {
    AboutUs,
  },
};
</script>
