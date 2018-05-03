<template>
    <div class="renewform">
        <section class="rf_body">
            <el-date-picker
              v-model="date"
              type="date"
              :picker-options="pickerOptions"
              placeholder="选择日期">
            </el-date-picker>
            <b-datepicker v-model="date" inline :dayNames="dayNames" :monthNames="monthNames" :min-date="minDate"></b-datepicker>
        </section>
        <span slot="footer" class="dialog-footer">
          <el-button @click="cancel">关闭</el-button>
          <el-button type="primary" @click="onSubmit">更新</el-button>
        </span>
    </div>
</template>

<script>
import request from '../apis/request'

export default {
  props: {
    cancel: Function,
    selectUser: Object,
  },

  data() {
    const today = new Date()
    return {
      date: today,
      dayNames: ['日', '一', '二', '三', '四', '五', '六'],
      monthNames: [
        '一月',
        '二月',
        '三月',
        '四月',
        '五月',
        '六月',
        '七月',
        '八月',
        '九月',
        '十月',
        '十一月',
        '十二月',
      ],
      pickerOptions: {
        disabledDate(time) {
          return time.getTime() < Date.now();
        },
      }
    }
  },

  methods: {
    onSubmit() {
      request
        .put(`/api/auth/${this.selectUser.Id}/renew`, {
          expired: this.date.valueOf() / 1000,
        })
        .then((response) => {
          if (response.success) {
            this.$message.success('账号续期成功!')
            this.$emit('renew-success')
          } else {
            this.$message.error(response.message)
          }
          this.cancel()
        })
    },
  },

  mounted() {
    this.date = new Date(this.selectUser.Expired)
  },
}
</script>
<style lang="less">
.renewform {
  .rf_body {
    padding: 20px 0;
  }
}
</style>
