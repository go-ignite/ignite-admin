<template>
  <div class="container" id="ct">
      <div class="columns">
          <div class="column is-12">
              <code-table></code-table>
          </div>
      </div>
  </div>
</template>
<script>
import CodeTable from '../components/CodeTable.vue'

export default {
  components: {
    CodeTable,
  },
};
</script>
