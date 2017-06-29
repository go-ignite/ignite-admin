<template>
    <div class="modal-card">
            <header class="modal-card-head">
                <p class="modal-card-title">批量生成邀请码</p>
            </header>
            <section class="modal-card-body">
                <b-field label="数量">
                    <b-input
                        v-model="amount"
                        placeholder="输入邀请码数量"
                        required>
                    </b-input>
                </b-field>

                <b-field label="流量 (GB)">
                    <b-input
                        v-model="limit"
                        placeholder="输入月流量"
                        required>
                    </b-input>
                </b-field>

                <b-field label="有效期 (月)">
                    <b-input
                        v-model="available"
                        placeholder="输入有效月数"
                        required>
                    </b-input>
                </b-field>
            </section>
            <footer class="modal-card-foot">
                <button class="button" type="button" @click="$emit('close')">关闭</button>
                <button class="button is-primary" type="button" @click="onSubmit">批量生成</button>
            </footer>
    </div>
</template>

<script>
    export default {

    data() {
        return {
            amount: 1,
            limit: 1,
            available: 1,
        }
    },
        props: ['refresh'],
        methods: {
            onSubmit() {
                console.log("submit button clicked...");
                this.$emit('close');
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