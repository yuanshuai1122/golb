<template v-if="isRouterAlive">
    <!-- 弹窗添加链接 -->
    <el-dialog v-model="addLinkDialogFormVisible" title="添加链接" width="400px" center @close="addLinkCloseDialog(false)">
        <el-form ref="ruleFormRef" :model="form" :rules="rules">
            <el-form-item prop="text" label="文字：">
                <el-input v-model="form.text" placeholder="请输入" />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="addLinkCloseDialog(false)">取消</el-button>
                <el-button plain color="#2fa7b9" @click="submitAdd">添加</el-button>
            </span>
        </template>
    </el-dialog>
</template>

<script>
import {
    reactive,
    ref,
    toRefs,
    nextTick
} from 'vue'

import {
    ElNotification,
} from 'element-plus'

import axios from "axios";

export default {
    props: {
        addLinkDialogFormVisible: Boolean
    },
    emits: ["onAddLinkCloseDialog"],
    setup(props, {
        emit
    }) {

        // 获取页面元素
        const ruleFormRef = ref('');

        const state = reactive({
            // 弹窗状态
            addLinkDialogFormVisible: props.addLinkDialogFormVisible,
            form: {
                text: "",

            },
        })

        // 网站名称非空验证
        const validateUrlName = (rule, value, callback) => {
            if (!value) {
                return callback(new Error('请填写此字段~'))
            } else {
                callback()
            }
        };

        // 约束规则 正则
        const rules = reactive({
            text: [{
                validator: validateUrlName,
                trigger: 'blur'
            }],
        });


        // 局部组件刷新
        const isRouterAlive = ref(true);
        // 删除组件
        isRouterAlive.value = false;
        nextTick(() => {
            // 重新渲染组件
            isRouterAlive.value = true;
        });
        // 取消
        const addLinkCloseDialog = (visible) => {
            // 关闭窗口
            emit("onAddLinkCloseDialog", visible);
            // 重置表单项，将其值重置为初始值，并移除校验结果
            ruleFormRef.value.resetFields()
        };

        // 链接添加
        function addLink(ruleFormRef) {
            const submitAdd = () => {
                // 1：校验链接信息
                ruleFormRef.value.validate((valid) => {
                    // 如果表单校验成功 提交数据到后台
                    if (valid) {
                        const params = {
                            text: state.form.text,
                        }
                        axios.post("/text/insert", params).then((res) => {
                            if (res.data.code == 0) {
                                // 关闭弹窗
                                emit("onAddLinkCloseDialog", false, res.data);
                                ElNotification({
                                    title: '成功',
                                    message: '添加成功',
                                    type: 'success',
                                });
                                // 重置表单项，将其值重置为初始值，并移除校验结果
                                ruleFormRef.value.resetFields()
                            }
                        })

                    }
                })

            }
            return {
                submitAdd
            }
        }

        const dialogVisible = ref(false)

        return {
            ...toRefs(state),
            rules,
            addLinkCloseDialog,
            ...addLink(ruleFormRef),
            dialogVisible,
            ruleFormRef,
            isRouterAlive
        }
    }


}
</script>
<style scoped>
.el-button--text {
    margin-right: 15px;
}

.dialog-footer button:first-child {
    margin-right: 10px;
}

.el-select {
    width: 100%;
}

:global(.el-dialog__title) {
    color: #2fa7b9 !important;
    font-size: 20px !important;
    font-weight: 700 !important;
}

:deep(.el-upload-dragger) {
    padding: 0px !important;
    width: 100%;
    height: 130px;
}

:global(.el-dialog__body) {
    padding: 25px calc(var(--el-dialog-padding-primary) + 5px) 0px !important;
}
</style>