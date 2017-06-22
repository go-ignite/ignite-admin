<template>
  <table class="table">
    <thead>
      <tr>
        <th>Name</th>
        <th>InviteCode</th>
        <th>Created</th>
        <th>Expired</th>
        <th>Package</th>
        <th>Used</th>
        <th>Port</th>
        <th>Status</th>
        <th>Reset</th>
        <th>Destroy</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="item in statusList">
        <th>{{ item.Username }}</th>
        <td>{{ item.InviteCode }}</td>
        <td>{{ item.Created | dateFilter}}</td>
        <td>{{ item.Expired | dateFilter}}</td>
        <td>{{ item.PackageLimit }} GB</td>
        <td>{{ item.PackageUsed | bandwidth }}</td>
        <td>{{ item.ServicePort }}</td>
        <td v-if="item.Status === 0 "><font color="gray">未创建</font></td>
        <td v-else-if="item.Status === 1 "><font color="green">运行中</font></td>
        <td v-else><font color="red">已停止</font></td>
        <td @click="reset(item.Id)"><a class="button is-success is-small">重置流量</a></td>
        <td @click="destroy(item.Id)"><a class="button is-danger is-small">一键销毁</a></td>
      </tr>
    </tbody>
  </table>
</template>

<script>
import axios from 'axios';
axios.defaults.headers.common['Authorization'] = localStorage.getItem("token");

export default {
  data () {
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
    reset: (id) => {
      console.log(id);
    },
    destroy: (id) => {
      console.log(id);
    },
    confirmCustomDelete: ()=> {
            this.$dialog.confirm({
                title: 'Deleting account',
                message: 'Are you sure you want to <strong>delete</strong> your account? This action cannot be undone.',
                confirmText: 'Delete Account',
                type: 'is-danger',
                hasIcon: true,
                onConfirm: () => {
                    this.$toast.open('Account deleted!')
                }
            })
        }
  },
  created () {
    let self = this;

    axios.get("/panel/status_list")
      .then(function (response) {
        if (response.data.success) {
          console.log(response.data.data);
          self.statusList = response.data.data;
        }
      })
      .catch(function (error) {
        console.log(error);
      });
  }
}
</script>

<style scoped>
.table {
  margin: 80px auto 180px auto;
}
</style>