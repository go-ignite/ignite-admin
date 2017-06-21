<template>
  <table class="table">
    <thead>
      <tr>
        <th>Name</th>
        <th>
          <abbr title="Played">InviteCode</abbr>
        </th>
        <th>
          <abbr title="Won">Package</abbr>
        </th>
        <th>
          <abbr title="Drawn">Used</abbr>
        </th>
        <th>
          <abbr title="Lost">Port</abbr>
        </th>
        <th>
          <abbr title="Goals for">Status</abbr>
        </th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="item in statusList">
        <th>{{ item.Username }}</th>
        <td>
        </td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
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
  created () {
    console.log('UserTab created...');
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