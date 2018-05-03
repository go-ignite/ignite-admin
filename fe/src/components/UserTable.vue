<template>
  <div class="iadmin_usertable">
    <b-modal :active.sync="showModal" :height="550" :width="380" :cancel="cancelRenew">
      <renew-form @renew-success="fetchData" :showModal="showModal" :selectUser="selectUser" :cancel="cancelRenew"></renew-form>
    </b-modal>
    <t-c-r
      :tableData="statusList"
      :tableCols="tableCols"
      :pagination="pagination"
    >
      <template slot="operator" slot-scope="{ col, row }">
        <a v-if="row.Status === 1" @click="stop(item)" class="button is-warning is-small">停止服务</a>
          <a v-if="row.Status === 2" @click="start(row)" class="button is-primary is-small">启动服务</a>
          <a @click="reset(row)" class="button is-success is-small">重置流量</a>
          <a @click="renew(row)" class="button is-info is-small">服务续期</a>
          <a @click="destroy(row)" class="button is-danger is-small">一键销毁</a>
      </template>
    </t-c-r>
  </div>
</template>

<script>
import TCR from '@/components/TableColumnRender.vue'
import request from '../apis/request'
import RenewForm from './RenewForm.vue'

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
          formatter: (v) => `${v} GB`
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
            label: '操作',
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
      const index = this.statusList.findIndex(e => e.Id === iten.Id)
      this.$dialog.confirm({
        title: '停止服务',
        message:
          '是否确定 <strong>停止</strong> 用户帐号 <strong>' + item.Username + '</strong> 的服务?',
        confirmText: '停止服务',
        cancelText: '取消',
        type: 'is-warning',
        hasIcon: true,
        onConfirm: () => {
          request
            .put('/api/auth/' + item.Id.toString() + '/stop')
            .then((response) => {
              if (response.success) {
                this.statusList[index].Status = 2
                this.$toast.open('用户帐号对应服务已停止!')
              } else {
                this.$toast.open('停止服务失败!')
              }
            })
            .catch(() => {
              this.$toast.open('停止服务失败!')
            })
        },
      })
    },
    start(item) {
      const index = this.statusList.findIndex(e => e.Id === iten.Id)
      this.$dialog.confirm({
        title: '启动服务',
        message:
          '是否确定 <strong>启动</strong> 用户帐号 <strong>' + item.Username + '</strong> 的服务?',
        confirmText: '启动服务',
        cancelText: '取消',
        type: 'is-primary',
        hasIcon: true,
        onConfirm: () => {
          request
            .put('/api/auth/' + item.Id.toString() + '/start')
            .then((response) => {
              if (response.success) {
                this.statusList[index].Status = 1
                this.$toast.open('用户帐号对应服务已成功启动!')
              } else {
                this.$toast.open('停止服务失败!')
              }
            })
            .catch(() => {
              this.$toast.open('启动服务失败!')
            })
        },
      })
    },
    reset(item) {
      const index = this.statusList.findIndex(e => e.Id === iten.Id)
      this.$dialog.confirm({
        title: '重置流量',
        message:
          '是否确定 <strong>重置</strong> 用户帐号 <strong>' +
          item.Username +
          '</strong> 的本月流量?',
        confirmText: '重置',
        cancelText: '取消',
        type: 'is-info',
        hasIcon: true,
        onConfirm: () => {
          request
            .put('/api/auth/' + item.Id.toString() + '/reset')
            .then((response) => {
              if (response.success) {
                this.statusList[index].PackageUsed = 0
                this.$toast.open('用户帐号本月流量已重置!')
              } else {
                this.$toast.open('重置用户帐号本月流量失败!')
              }
            })
            .catch(() => {
              this.$toast.open('重置用户帐号本月流量失败!')
            })
        },
      })
    },
    destroy(item) {
      const index = this.statusList.findIndex(e => e.Id === iten.Id)
      this.$dialog.confirm({
        title: '销毁账户',
        message:
          '是否确定 <strong>销毁</strong> 用户帐号 <strong>' +
          item.Username +
          '</strong> ? 该操作将不可逆转',
        confirmText: '销毁帐号',
        cancelText: '取消',
        type: 'is-danger',
        hasIcon: true,
        onConfirm: () => {
          request
            .put('/api/auth/' + item.Id.toString() + '/destroy')
            .then((response) => {
              if (response.success) {
                this.statusList.splice(index, 1)
                this.$toast.open('用户帐号已销毁!')
              } else {
                this.$toast.open('销毁用户帐号失败!')
              }
            })
            .catch(() => {
              this.$toast.open('销毁用户帐号失败!')
            })
        },
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
