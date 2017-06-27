<template>
  <table class="table">
    <thead>
      <tr>
        <th>用户名</th>
        <th>邀请码</th>
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
      <tr v-for="(item, index) in statusList">
        <th>{{ item.Username }}</th>
        <td>{{ item.InviteCode }}</td>
        <td><span class="tag is-primary">{{ item.Created | dateFilter}}</span></td>
        <td><span class="tag is-primary">{{ item.Expired | dateFilter}}</span></td>
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
          <a @click="reset(item, index)" class="button is-success is-small">重置流量</a>
          <a @click="confirmUserDelete(item.Id, item.Username)" class="button is-danger is-small">一键销毁</a>
        </td>
      </tr>
    </tbody>
  </table>
</template>

<script>
import axios from 'axios';
if (localStorage.getItem("token") != "") {
  axios.defaults.headers.common['Authorization'] = "Bearer " + localStorage.getItem("token");
}

export default {
  data() {
    return {
      statusList: []
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
    reset(item, index) {
      self = this;
      this.$dialog.confirm({
        title: '重置流量',
        message: '是否确定 <strong>重置</strong> 用户帐号 <strong>' + item.Username + '</strong> 的本月流量?',
        confirmText: '确定重置',
        cancelText: '取消',
        type: 'is-info',
        hasIcon: true,
        onConfirm: () => {
          axios.put("/auth/" + item.Id.toString() + "/reset")
            .then( (response) => {
              if (response.status == 200) {
                if (response.data.success) {
                  self.statusList[index].PackageUsed = 0;
                  self.$toast.open('用户帐号本月流量已重置!');
                }
              } else {
                self.$toast.open('重置用户帐号本月流量失败!');
              }
            })
            .catch( (error) => {
              console.log(error);
              if (error.response.status == 401) {
                localStorage.setItem("token", "");
                location.href = '/';
              }
            });
        }
      })
    },
    confirmUserDelete(id, name) {
      this.$dialog.confirm({
        title: '销毁账户',
        message: '是否确定 <strong>销毁</strong> 用户帐号 <strong>' + name + '</strong> ? 该操作将不可逆转',
        confirmText: '销毁帐号',
        cancelText: '取消',
        type: 'is-danger',
        hasIcon: true,
        onConfirm: () => {
          this.$toast.open('用户帐号已销毁!')
        }
      })
    }
  },
  created() {
    let self = this;

    axios.get("/auth/status_list")
      .then(function (response) {
        if (response.status == 200) {
          if (response.data.success) {
            console.log(response.data.data);
            self.statusList = response.data.data;
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
  margin: 80px auto 180px auto;
}

.table th,
.table td {
  text-align: center;
}
</style>