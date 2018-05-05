<template>
  <div class="iadmin_usertable g_wrap">
    <el-dialog
      @close="cancelRenew"
      :visible.sync="showModal"
      title="服务续期"
      width="380">
      <renew-form @renew-success="fetchData" :showModal="showModal" :selectUser="selectUser" :cancel="cancelRenew"></renew-form>
    </el-dialog>
    <t-c-r
      :tableData="statusList"
      :tableCols="tableCols"
      :pagination="pagination"
    >
      <template slot="operator" slot-scope="{ col, row }">
        <el-dropdown>
           <el-button type="primary" size="mini">
            操作<i class="el-icon-arrow-down el-icon--right"></i>
          </el-button>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item v-if="row.Status === 0" @click.native="start(row)">启动服务</el-dropdown-item>
            <el-dropdown-item v-if="row.Status === 1" @click.native="stop(row)">停止服务</el-dropdown-item>
            <el-dropdown-item @click.native="reset(row)" divided>重置流量</el-dropdown-item>
            <el-dropdown-item @click.native="renew(row)">服务续期</el-dropdown-item>
            <el-dropdown-item @click.native="destroy(row)">一键销毁</el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
    </template>
    </t-c-r>
  </div>
</template>

<script>
import TCR from '@/components/TableColumnRender.vue'
import RenewForm from '@/components/RenewForm.vue'
import request from '@/apis/request'

export default {
  computed: {
    tableCols() {
      return [
        {
          raw: {
            label: '用户名',
            prop: 'Username',
          },
        },
        {
          raw: {
            label: '服务类型',
            prop: 'ServiceType',
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
            label: '过期时间',
            prop: 'Expired',
          },
          formatter: (v) => this.dateFilter(v)
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
            label: '已使用',
            prop: 'PackageUsed',
          },
          formatter: (v) => this.bandwidth(v)
        },
        {
          raw: {
            label: '端口号',
            prop: 'ServicePort',
          },
        },
        {
          raw: {
            label: '状态',
            prop: 'Status',
          },
          formatter: (v) => this.statusFormat(v)
        },
        {
          raw: {
            label: '',
            width: '150px',
          },
          slot: 'operator',
        },
      ]
    },

  },
  data() {
    return {
      pagination: {
        total: 0,
        size: 12,
      },
      statusList: [],
      isSimple: false,
      showModal: false,
      selectUser: {},
    }
  },
  filters: {
    bandwidth: (value) => {
      return value.toFixed(2).toString() + ' GB'
    },
    dateFilter: (value) => {
      return value.split('T')[0]
    },
  },
  methods: {
    pageChanged(value) {
      request
        .get(`/api/auth/status_list?pageIndex=${value.toString()}&pageSize=${this.pagination.size}`)
        .then((response) => {
          if (response.success) {
            this.statusList = response.data.data
            this.pagination.total = response.data.total
          }
        })
    },
    stop(item) {
      const index = this.statusList.findIndex(e => e.Id === item.Id)
      this.$confirm(`是否确定停止用户${item.Username}的服务?`, '停止服务', {
        confirmButtonText: '停止服务',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
          request
            .put('/api/auth/' + item.Id.toString() + '/stop')
            .then((response) => {
              if (response.success) {
                this.statusList[index].Status = 2
                this.$message('用户帐号对应服务已停止!')
              } else {
                this.$message('停止服务失败!')
              }
            })
            .catch(() => {
              this.$message('停止服务失败!')
            })
        })
    },
    start(item) {
      const index = this.statusList.findIndex(e => e.Id === item.Id)
      this.$confirm(`是否确定启动用户${item.Username}的服务?`, '启动服务', {
        confirmButtonText: '启动服务',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
          request
            .put('/api/auth/' + item.Id.toString() + '/start')
            .then((response) => {
              if (response.success) {
                this.statusList[index].Status = 1
                this.$message('用户帐号对应服务已成功启动!')
              } else {
                this.$message.error(response.message)
              }
            })
            .catch(() => {
              this.$message.error('启动服务失败')
            })
      })
    },
    reset(item) {
      const index = this.statusList.findIndex(e => e.Id === item.Id)
      this.$confirm(`是否确定重置用户账号${item.Username}的本月流量?`, '重置流量', {
        confirmButtonText: '重置',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
          request
            .put('/api/auth/' + item.Id.toString() + '/reset')
            .then((response) => {
              if (response.success) {
                this.statusList[index].PackageUsed = 0
                this.$message('用户帐号本月流量已重置!')
              } else {
                this.$message('重置用户帐号本月流量失败!')
              }
            })
            .catch(() => {
              this.$message('重置用户帐号本月流量失败!')
            })
      })
    },
    destroy(item) {
      const index = this.statusList.findIndex(e => e.Id === item.Id)
      this.$confirm(`是否确定销毁用户账号${item.Username}?该操作将不可逆转`, '销毁账号', {
        confirmButtonText: '销毁账号',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
          request
            .put('/api/auth/' + item.Id.toString() + '/destroy')
            .then((response) => {
              if (response.success) {
                this.statusList.splice(index, 1)
                this.$message('用户帐号已销毁!')
              } else {
                this.$message('销毁用户帐号失败!')
              }
            })
            .catch(() => {
              this.$message('销毁用户帐号失败!')
            })
      })
    },
    renew(item) {
      this.selectUser = item
      this.showModal = true
    },
    cancelRenew() {
      this.selectUser = {}
      this.showModal = false
    },
    fetchData() {
      request.get(`/api/auth/status_list?pageIndex=1&pageSize=${this.pagination.size}`).then((response) => {
        if (response.success) {
          this.statusList = response.data.data
          this.pagination.total = response.data.total
        }
      })
    },
    dateFilter: (value) => {
      return value.split('T')[0]
    },
    bandwidth: (value) => {
      return value.toFixed(2).toString() + ' GB'
    },
    statusFormat(v) {
      if (v === 0) return '未创建'
      if (v === 1) return '运行中'
      return '已停止'
    }
  },
  created() {
    this.fetchData()
  },
  components: {
    RenewForm,
    TCR,
  },
}
</script>

<style scoped>
.table {
  margin: 80px auto 20px auto;
  width: 100%;
}

.table th,
.table td {
  text-align: center;
}
</style>
