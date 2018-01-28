<template>
    <div>
        <a class="button is-primary batch" @click="showModal = true">批量生成</a>
    
        <b-modal v-on:close="closed" :active.sync="showModal" :component="ModalForm" :width="360">
        </b-modal>
    
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
                    <td>{{ item.AvailableLimit }} 个月</td>
                    <td>{{ item.PackageLimit }} GB</td>
                    <td>
                        <a @click="remove(item, index)" class="button is-warning is-small">删除</a>
                    </td>
                </tr>
            </tbody>
        </table>
    
        <b-pagination v-on:change="pageChanged" :total="total" :current.sync="current" :order="order" :size="size" :simple="isSimple" :per-page="perPage">
        </b-pagination>
    </div>
</template>

<script>
import request from '../apis/request';
import ModalForm from './ModalForm.vue';

export default {
    data() {
        return {
            codeList: [],
            total: 0,
            current: 1,
            perPage: 12,
            order: '',
            size: 'is-small',
            isSimple: false,
            showModal: false,
            ModalForm
        }
    },
    filters: {
        dateFilter: (value) => {
            return value.split('T')[0];
        }
    },
    methods: {
        closed() {
            console.log("modal dialog closed...");
            let self = this;

            //Refresh all the valid invite codes.
            request.get("/api/auth/code_list?pageIndex=1&pageSize=12")
                .then(function (response) {
                    if (response.success) {
                        self.codeList = response.data.data;
                        self.total = response.data.total;
                    }
                })
        },
        remove(item, index) {
            self = this;
            this.$dialog.confirm({
                title: '删除邀请码',
                message: '是否确定 <strong>删除</strong> 邀请码 <strong>' + item.InviteCode + '</strong> ?',
                confirmText: '确认删除',
                cancelText: '取消',
                type: 'is-warning',
                hasIcon: true,
                onConfirm: () => {
                    request.put("/api/auth/" + item.Id.toString() + "/remove")
                        .then((response) => {
                                if (response.success) {
                                    self.codeList.splice(index, 1);
                                    self.$toast.open('邀请码已删除!');
                                } else {
                                    self.$toast.open('删除邀请码失败!');
                                }
                            })
                        .catch(() => {
                            self.$toast.open('删除邀请码失败!');

                        })
                }
            })

        },
        pageChanged(value) {
            let self = this;

            request.get("/api/auth/code_list?pageIndex=" + value.toString() + "&pageSize=12")
                .then(function (response) {
                    if (response.success) {
                        self.codeList = response.data.data;
                        self.total = response.data.total;
                    }
                })
        }
    },
    created() {
        let self = this;

        request.get("/api/auth/code_list?pageIndex=1&pageSize=12")
            .then(function (response) {
                if (response.success) {
                    self.codeList = response.data.data;
                    self.total = response.data.total;
                }
            })
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
