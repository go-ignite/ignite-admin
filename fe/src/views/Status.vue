<template>
  <div class="container" id="ut">
      <div class="columns">
          <div class="column is-12">
              <user-table></user-table>
          </div>
      </div>
  </div>
</template>
<script>
import UserTable from '../components/UserTable.vue'

export default {
  components: {
    UserTable,
  },
};
</script>
