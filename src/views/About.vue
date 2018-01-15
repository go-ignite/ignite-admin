<template>
  <div>
    <g-navbar :page="3"></g-navbar>
    <div class="container" id="profile">
        <div class="columns">
            <div class="column is-4 is-offset-4">
               <about-us></about-us>
            </div>
        </div>
    </div>
    <g-footer></g-footer>
  </div>
</template>
<script>
import GNavbar from '../components/Nav.vue';
import GFooter from '../components/Footer.vue';
import AboutUs from '../components/AboutUs.vue'

export default {
  components: {
    GNavbar,
    GFooter,
    AboutUs,
  },
};
</script>
