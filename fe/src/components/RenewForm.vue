<template>
    <div class="modal-card">
        <header class="modal-card-head">
            <p class="modal-card-title">帐号续期</p>
        </header>
        <section class="modal-card-body">
            <b-datepicker v-model="date" inline :dayNames="dayNames" :monthNames="monthNames" :min-date="minDate"></b-datepicker>
        </section>
        <footer class="modal-card-foot">
            <button class="button" type="button" @click="cancel">关闭</button>
            <button class="button is-primary" type="button" @click="onSubmit">更新</button>
        </footer>
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
      minDate: new Date(today.getFullYear(), today.getMonth(), today.getDate() + 1),
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

<style scoped>
.modal-card {
  margin: 0 auto;
  width: auto;
}
</style>
