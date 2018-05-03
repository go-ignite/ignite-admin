<template>
    <div class="iadmin_codetable g_wrap">
        <el-button @click="showModal = true" type="primary">批量生成</el-button>
        <t-c-r
          :tableData="codeList"
          :tableCols="tableCols"
          :pagination="pagination"
          @page-change="pageChanged"
        >
        </t-c-r>
        <el-dialog
          @close="closed"
          :visible.sync="showModal"
          title="批量生成邀请码"
          width="360">
          <modal-form @closeModal="showModal = false"></modal-form>
        </el-dialog>
    </div>
</template>

<script>
import TCR from '@/components/TableColumnRender.vue'
import ModalForm from '@/components/ModalForm.vue'
import request from '../apis/request'

export default {
  computed: {
    tableCols() {
      const removeFn = this.remove
      return [
        {
          raw: {
            label: '邀请码',
            prop: 'InviteCode',
          },
        },
        {
          raw: {
            label: '创建时间',
            prop: 'Created',
          },
          formatter: (v) => this.dateFilter(v)
        },
        {
          raw: {
            label: '有效期限',
            prop: 'AvailableLimit',
          },
          formatter: (v) => `${v}个月`
        },
        {
          raw: {
            label: '总流量',
            prop: 'PackageLimit',
          },
          formatter: (v) => `${v} GB`
        },
        {
          raw: {
            label: '操作',
          },
          render: {
            props: {
              row: {
                type: Object,
                default: () => ({}),
              },
            },
            render() {
              return (
                <el-button onClick={() => removeFn(this.row)} type="danger" icon="el-icon-delete" circle></el-button>
              )
            },
          },
        },
      ]
    },
  },
  data() {
    return {
      codeList: [],
      pagination: {
        total: 0,
        size: 12,
      },
      order: '',
      size: 'is-small',
      isSimple: false,
      showModal: false,
    }
  },
  filters: {
    dateFilter: (value) => {
      return value.split('T')[0]
    },
  },
  methods: {
    closed() {
      //Refresh all the valid invite codes.
      request.get(`/api/auth/code_list?pageIndex=1&pageSize=${this.pagination.size}`).then((response) => {
        if (response.success) {
          this.codeList = response.data.data
          this.pagination.total = response.data.total
        }
      })
    },
    remove(item, index) {
      // TODO: FIX remove effect page
      this.$confirm('是否确定删除邀请码?', '提示', {
        title: '删除邀请码',
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        const removeId = item.Id
        request
          .put(`/api/auth/${removeId}/remove`)
          .then((response) => {
            if (response.success) {
              const index = this.codeList.findIndex(e => e.Id === removeId)
              if (index > -1) {
                this.codeList.splice(index, 1)
                this.$message.success('邀请码已删除')
              }
            } else {
              this.$message.error('删除邀请码失败!')
            }
          })
          .catch(() => {
            this.$message.error('删除邀请码失败!')
          })
      }).catch(() => {})
    },
    pageChanged(value) {
      request
        .get(`/api/auth/code_list?pageIndex=${value}&pageSize=${this.pagination.size}`)
        .then((response) => {
          if (response.success) {
            this.codeList = response.data.data
            this.pagination.total = response.data.total
          }
        })
    },
    dateFilter: (value) => {
      return value.split('T')[0]
    },
  },
  created() {
    request.get(`/api/auth/code_list?pageIndex=1&pageSize=${this.pagination.size}`).then((response) => {
      if (response.success) {
        this.codeList = response.data.data
        this.pagination.total = response.data.total
      }
    })
  },
  components: {
    TCR,
    ModalForm,
  },
}
</script>

<style scoped>
.batch {
  margin: 80px auto 0px auto;
}

.table {
  margin: 20px auto 20px auto;
}

.table th,
.table td {
  text-align: center;
}

.pagination {
  margin: 20px auto 80px auto;
}
</style>
