<template>
  <div>
    <a class="button is-primary batch" @click="onBatch">批量添加</a>
    <table class="table">
      <thead>
        <tr>
          <th>邀请码</th>
          <th>创建时间</th>
          <th>有效期限</th>
          <th>总流量</th>
          <th>操作</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, index) in codeList">
          <td>{{ item.InviteCode }}</td>
          <td>
            <span class="tag is-primary">{{ item.Created | dateFilter}}</span>
          </td>
          <td>{{ item.PackageLimit }} GB</td>
          <td>{{ item.AvailableLimit }} 个月</td>
          <td>
            <a @click="remove(item, index)" class="button is-danger is-small">删除</a>
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
if (localStorage.getItem("token") != "") {
  axios.defaults.headers.common['Authorization'] = "Bearer " + localStorage.getItem("token");
}

export default {
  data() {
    return {
      codeList: [],
      total: 0,
      current: 1,
      perPage: 12,
      order: '',
      size: 'is-small',
      isSimple: false
    }
  },
  filters: {
    dateFilter: (value) => {
      return value.split('T')[0];
    }
  },
  methods: {
    onBatch() {
        console.log("batch button clicked...");
    },
    remove(item, index) {
        console.log(index);
    },
    pageChanged(value) {
      let self = this;

      axios.get("/auth/code_list?pageIndex=" + value.toString() + "&pageSize=12")
        .then(function (response) {
          if (response.status == 200) {
            if (response.data.success) {
              self.codeList = response.data.data.data;
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
  },
  created() {
    let self = this;

    axios.get("/auth/code_list?pageIndex=1&pageSize=12")
      .then(function (response) {
        if (response.status == 200) {
          if (response.data.success) {
            console.log(response.data.data);
            self.codeList = response.data.data.data;
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