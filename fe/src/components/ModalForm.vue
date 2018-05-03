<template>
    <div class="modal-card">
        <el-form>
          <el-form-item label="数量">
            <el-input v-model.number="amount" placeholder="输入邀请码数量" type="number" min="1" max="100" required>
            </el-input>
          </el-form-item>
          <el-form-item label="流量（GB）">
            <el-input v-model.number="limit" placeholder="输入月流量" type="number" min="1" max="999999" required>
            </el-input>
          </el-form-item>
          <el-form-item label="有效期（月）">
            <el-input v-model.number="available" placeholder="输入有效月数" type="number" min="1" max="999999" required></el-input>
          </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
          <el-button @click="$emit('closeModal')">关闭</el-button>
          <el-button type="primary" @click="onSubmit">批量生成</el-button>
        </span>
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
            this.$message.success('生成邀请码成功!')
          } else {
            this.$message.error('生成邀请码失败!')
          }
          this.$emit('closeModal')
        })
        .catch((error) => {
          this.$message.error('生成邀请码失败!')
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
