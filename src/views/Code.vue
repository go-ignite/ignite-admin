<template>
  <div>
    <g-navbar :page="2"></g-navbar>

    <div class="container" id="ct">
        <div class="columns">
            <div class="column is-12">
               <code-table></code-table>
            </div>
        </div>
    </div>
    <g-footer></g-footer>
  </div>
</template>
<script>
import GNavbar from '../components/Nav.vue';
import GFooter from '../components/Footer.vue';
import CodeTable from '../components/CodeTable.vue'

export default {
  components: {
    GNavbar,
    GFooter,
    CodeTable,
  },
};
</script>
