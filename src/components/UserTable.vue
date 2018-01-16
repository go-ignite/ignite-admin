<template>
  <div>
    <table class="table">
      <thead>
        <tr>
          <th>用户名</th>
          <th>服务类型</th>
          <th>创建时间</th>
          <th>过期时间</th>
          <th>总流量</th>
          <th>已使用</th>
          <th>端口号</th>
          <th>状态</th>
          <th>操作</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, index) in statusList" :key="item.Id">
          <th>{{ item.Username }}</th>
          <td>{{ item.ServiceType }}</td>
          <td>
            <span class="tag is-primary">{{ item.Created | dateFilter}}</span>
          </td>
          <td>
            <span class="tag is-primary">{{ item.Expired | dateFilter}}</span>
          </td>
          <td>{{ item.PackageLimit }} GB</td>
          <td>{{ statusList[index].PackageUsed | bandwidth }}</td>
          <td>{{ item.ServicePort }}</td>
          <td v-if="item.Status === 0 ">
            <font color="gray">未创建</font>
          </td>
          <td v-else-if="item.Status === 1 ">
            <font color="green">运行中</font>
          </td>
          <td v-else>
            <font color="red">已停止</font>
          </td>
          <td>
            <a v-if="item.Status === 1" @click="stop(item, index)" class="button is-warning is-small">停止服务</a>
            <a v-if="item.Status === 2" @click="start(item, index)" class="button is-primary is-small">启动服务</a>
            <a @click="reset(item, index)" class="button is-success is-small">重置流量</a>
            <a @click="destroy(item, index)" class="button is-danger is-small">一键销毁</a>
          </td>
        </tr>
      </tbody>
    </table>
  
    <b-pagination v-on:change="pageChanged" :total="total" :current.sync="current" :order="order" :size="size" :simple="isSimple" :per-page="perPage">
    </b-pagination>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      statusList: [],
      total: 0,
      current: 1,
      perPage: 12,
      order: '',
      size: 'is-small',
      isSimple: false
    }
  },
  filters: {
    bandwidth: (value) => {
      return value.toFixed(2).toString() + ' GB';
    },
    dateFilter: (value) => {
      return value.split('T')[0];
    }
  },
  methods: {
    pageChanged(value) {
      let self = this;

      axios.get("/auth/status_list?pageIndex=" + value.toString() + "&pageSize=12")
        .then(function (response) {
          if (response.status == 200) {
            if (response.data.success) {
              self.statusList = response.data.data.data;
              self.total = response.data.data.total;
            }
          }
        })
        .catch(function (error) {
          if (error.response.status == 401) {
            localStorage.setItem("token", "");
            location.href = '/';
          }
        });
    },
    stop(item, index) {
      self = this;
      this.$dialog.confirm({
        title: '停止服务',
        message: '是否确定 <strong>停止</strong> 用户帐号 <strong>' + item.Username + '</strong> 的服务?',
        confirmText: '停止服务',
        cancelText: '取消',
        type: 'is-warning',
        hasIcon: true,
        onConfirm: () => {
          axios.put("/auth/" + item.Id.toString() + "/stop")
            .then((response) => {
              if (response.status == 200) {
                if (response.data.success) {
                  self.statusList[index].Status = 2;
                  self.$toast.open('用户帐号对应服务已停止!');
                } else {
                  self.$toast.open('停止服务失败!');
                }
              } else {
                self.$toast.open('停止服务失败!');
              }
            })
            .catch((error) => {
              console.log(error);
              if (error.response.status == 401) {
                localStorage.setItem("token", "");
                location.href = '/';
              }
            });
        }
      })
    },
    start(item, index) {
      self = this;
      this.$dialog.confirm({
        title: '启动服务',
        message: '是否确定 <strong>启动</strong> 用户帐号 <strong>' + item.Username + '</strong> 的服务?',
        confirmText: '启动服务',
        cancelText: '取消',
        type: 'is-primary',
        hasIcon: true,
        onConfirm: () => {
          axios.put("/auth/" + item.Id.toString() + "/start")
            .then((response) => {
              if (response.status == 200) {
                if (response.data.success) {
                  self.statusList[index].Status = 1;
                  self.$toast.open('用户帐号对应服务已成功启动!');
                } else {
                  self.$toast.open('停止服务失败!');
                }
              } else {
                self.$toast.open('启动服务失败!');
              }
            })
            .catch((error) => {
              console.log(error);
              if (error.response.status == 401) {
                localStorage.setItem("token", "");
                location.href = '/';
              }
            });
        }
      })
    },
    reset(item, index) {
      self = this;
      this.$dialog.confirm({
        title: '重置流量',
        message: '是否确定 <strong>重置</strong> 用户帐号 <strong>' + item.Username + '</strong> 的本月流量?',
        confirmText: '重置',
        cancelText: '取消',
        type: 'is-info',
        hasIcon: true,
        onConfirm: () => {
          axios.put("/auth/" + item.Id.toString() + "/reset")
            .then((response) => {
              if (response.status == 200) {
                if (response.data.success) {
                  self.statusList[index].PackageUsed = 0;
                  self.$toast.open('用户帐号本月流量已重置!');
                } else {
                  self.$toast.open('重置用户帐号本月流量失败!');
                }
              } else {
                self.$toast.open('重置用户帐号本月流量失败!');
              }
            })
            .catch((error) => {
              console.log(error);
              if (error.response.status == 401) {
                localStorage.setItem("token", "");
                location.href = '/';
              }
            });
        }
      })
    },
    destroy(item, index) {
      self = this;
      this.$dialog.confirm({
        title: '销毁账户',
        message: '是否确定 <strong>销毁</strong> 用户帐号 <strong>' + item.Username + '</strong> ? 该操作将不可逆转',
        confirmText: '销毁帐号',
        cancelText: '取消',
        type: 'is-danger',
        hasIcon: true,
        onConfirm: () => {
          axios.put("/auth/" + item.Id.toString() + "/destroy")
            .then((response) => {
              if (response.status == 200) {
                if (response.data.success) {
                  self.statusList.splice(index, 1);
                  self.$toast.open('用户帐号已销毁!');
                } else {
                  self.$toast.open('销毁用户帐号失败!');
                }
              } else {
                self.$toast.open('销毁用户帐号失败!');
              }
            })
            .catch((error) => {
              console.log(error);
              if (error.response.status == 401) {
                localStorage.setItem("token", "");
                location.href = '/';
              }
            });
        }
      })
    }
  },
  created() {
    if (localStorage.getItem("token") != "") {
      axios.defaults.headers.common['Authorization'] = "Bearer " + localStorage.getItem("token");
    }
    let self = this;

    axios.get("/auth/status_list?pageIndex=1&pageSize=12")
      .then(function (response) {
        if (response.status == 200) {
          if (response.data.success) {
            console.log(response.data.data);
            self.statusList = response.data.data.data;
            self.total = response.data.data.total;
          }
        }
      })
      .catch(function (error) {
        if (error.response.status == 401) {
          localStorage.setItem("token", "");
          location.href = '/';
        }
      });
  }
}
</script>

<style scoped>
.table {
  margin: 80px auto 20px auto;
}

.table th,
.table td {
  text-align: center;
}

.pagination {
  margin: 20px auto 80px auto;
}
</style>
