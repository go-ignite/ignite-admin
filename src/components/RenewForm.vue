<template>
    <div class="modal-card">
        <header class="modal-card-head">
            <p class="modal-card-title">帐号续期</p>
        </header>
        <section class="modal-card-body">
            <b-datepicker v-model="date" inline></b-datepicker>
        </section>
        <footer class="modal-card-foot">
            <button class="button" type="button" @click="$emit('close')">关闭</button>
            <button class="button is-primary" type="button" @click="onSubmit">更新</button>
        </footer>
    </div>
</template>

<script>
import request from '../apis/request';

export default {
    data() {
        return {
            date: new Date()
        }
    },
    methods: {
        onSubmit() {
            let self = this;
            request.post("/api/auth/code_generate", {
                "amount": self.amount,
                "limit": self.limit,
                "available": self.available,
            }).then((response) => {
                self.$toast.open('生成邀请码成功!');
                self.$emit('close');
            }).catch((error) => {
                self.$toast.open({
                    message: '生成邀请码失败!',
                    type: 'is-danger'
                })
            });
        }
    }
}
</script>

<style scoped>
.modal-card {
    margin: 0 auto;
    width: auto;
}
</style>
