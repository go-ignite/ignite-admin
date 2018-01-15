<template>
  <div>
    <g-navbar :page="1"></g-navbar>

    <div class="container" id="ut">
        <div class="columns">
            <div class="column is-12">
                <user-table></user-table>
            </div>
        </div>
    </div>
    <g-footer></g-footer>
  </div>
</template>
<script>
import GNavbar from '../components/Nav.vue';
import GFooter from '../components/Footer.vue';
import UserTable from '../components/UserTable.vue'

export default {
  components: {
    GNavbar,
    GFooter,
    UserTable,
  },
};
</script>
