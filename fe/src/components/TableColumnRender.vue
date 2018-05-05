// HACK USE EL-TABLE
<template>
  <div class="tcr">
    <el-table
      :data="tableData"
      v-bind="tableConfig"
      style="width: 100%"
    >
      <el-table-column
        v-if="showIndex"
        type="index"
        class-name="g_table_index"
        label="序号"
        width="100">
      </el-table-column>
      <el-table-column
        v-for="col in tableCols"
        v-bind="col.raw"
        :key="col.raw.prop"
      >
        <template slot-scope="scope">
          <component v-if="col.render" :is="col.render" :row="scope.row"></component>
          <slot v-else-if="col.slot" :name="col.slot" :row="scope.row" :col="col"></slot>
          <span v-else>{{displayCellData(scope, col)}}</span>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      v-if="pagination && pagination.total"
      @current-change="pageChange"
      :current-page="pagination.index"
      :page-size="pagination.size"
      :total="pagination.total"
      layout="total, prev, pager, next"
      class="t_pagination"
      background
    >
    </el-pagination>
  </div>
</template>

<script>
export default {
  props: {
    showIndex: Boolean,
    tableCols: Array,
    tableConfig: {
      type: Object,
      default: () => ({}),
    },
    tableData: {
      type: Array,
      default: () => [],
    },
    pagination: {
      type: Object,
      default: null,
    },
  },

  methods: {
    displayCellData(scope, col) {
      const value = scope.row[col.raw.prop]
      if (!col.formatter) return value
      return col.formatter(value, scope.row)
    },

    pageChange(page) {
      this.$emit('page-change', page)
    },

  },
}
</script>

<style lang="less">
.tcr {
  th {
    color: #333;
  }
  .t_pagination {
    padding: 30px 50px 80px 0;
    text-align: right;
  }
}
</style>


