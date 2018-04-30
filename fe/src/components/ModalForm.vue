<template>
    <div class="modal-card">
        <section class="modal-card-body">
            <b-field label="数量">
                <b-input v-model.number="amount" placeholder="输入邀请码数量" type="number" min="1" max="100" required>
                </b-input>
            </b-field>
    
            <b-field label="流量 (GB)">
                <b-input v-model.number="limit" placeholder="输入月流量" type="number" min="1" max="999999" required>
                </b-input>
            </b-field>
    
            <b-field label="有效期 (月)">
                <b-input v-model.number="available" placeholder="输入有效月数" type="number" min="1" max="999999" required>
                </b-input>
            </b-field>
        </section>
        <footer class="modal-card-foot">
            <button class="button" type="button" @click="$emit('closeModal')">关闭</button>
            <button class="button is-primary" type="button" @click="onSubmit">批量生成</button>
        </footer>
    </div>
</template>

<script>
import request from '../apis/request'

export default {
  data() {
    return {
      amount: 1,
      limit: 1,
      available: 1,
    }
  },
  methods: {
    onSubmit() {
      request
        .post('/api/auth/code_generate', {
          amount: this.amount,
          limit: this.limit,
          available: this.available,
        })
        .then((response) => {
          if (response.success) {
            this.$toast.open('生成邀请码成功!')
          } else {
            this.$toast.open('生成邀请码失败!')
          }
          this.$emit('closeModal')
        })
        .catch((error) => {
          this.$toast.open({
            message: '生成邀请码失败!',
            type: 'is-danger',
          })
        })
    },
  },
}
</script>

<style scoped>
.modal-card {
  margin: 0 auto;
  width: auto;
}
</style>
