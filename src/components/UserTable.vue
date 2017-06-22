<template>
  <table class="table">
    <thead>
      <tr>
        <th>Name</th>
        <th>InviteCode</th>
        <th>Package</th>
        <th>Used</th>
        <th>Port</th>
        <th>Status</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="item in statusList">
        <th>{{ item.Username }}</th>
        <td>{{ item.InviteCode }}</td>
        <td>{{ item.PackageLimit }} GB</td>
        <td>{{ item.PackageUsed | bandwidth }}</td>
        <td>{{ item.ServicePort }}</td>
        <td>{{ item.ServiceStatus }}</td>
      </tr>
    </tbody>
  </table>
</template>

<script>
import axios from 'axios';

export default {
  data () {
    return {
      statusList: []
    }
  },
  filters: {
    bandwidth: (value) => {
      return value.toFixed(2).toString() + ' GB';
    }
  },
  created () {
    console.log('Token is:' + localStorage.getItem("token"));
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