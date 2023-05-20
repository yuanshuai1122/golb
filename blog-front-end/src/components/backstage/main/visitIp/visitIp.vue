<template>
    <el-card class="box-card">
        <template #header>
            <div class="card-header">
                <h3>
                    <el-icon style="margin-right: 10px;">
                        <Refresh />
                    </el-icon>访问信息
                </h3>
                <!-- 小于768px -->
                <div class="hidden-sm-and-down">
                    <el-row :gutter="10">
                        <el-col :span="15">
                            <el-button plain style="width: 100%;" color="#2fa7b9" @click="deleteAll">批量删除</el-button>
                        </el-col>
                        <!-- <el-col :span="10"> -->
                        <!-- </el-col> -->
                        <el-col :span="3" style="display: inline-flex;justify-content: center;align-items: center;">
                            <el-icon style="font-size: 20px;color: #b3b6bc;" @click="refresh">
                                <Refresh />
                            </el-icon>
                        </el-col>
                    </el-row>
                </div>

                <!-- 大于等于768px  -->
                <div class="visible">
                    <p style="display: inline-flex;justify-content: center;align-items: center;">
                        <el-dropdown style="line-height:0px;" :hide-on-click="false" trigger="click">
                            <span class="el-dropdown-link" style="color:#2fa7b9;font-size: 16px;margin-right: 15px;"
                                title="菜单">
                                <el-icon>
                                    <Grid />
                                </el-icon>
                            </span>
                            <template #dropdown>
                                <el-dropdown-menu class="visible">
                                    <!-- <el-dropdown-item>
                                        <el-button plain style="width: 100%;" color="#2fa7b9" @click="addLink">添加链接
                                        </el-button>
                                    </el-dropdown-item> -->
                                </el-dropdown-menu>
                            </template>
                        </el-dropdown>
                        <el-icon @click="refresh">
                            <Refresh />
                        </el-icon>
                    </p>
                </div>


            </div>
        </template>
        <div>
            <el-table v-loading="loading" element-loading-text="玩命加载中..." :data="tableData" ref="multipleTableRef"
                @selection-change="handleSelectionChange" style="width: 100%;text-align: center;"
                :cell-style="{ textAlign: 'center' }"
                :header-cell-style="{ fontSize: '15px', background: 'rgb(90 96 113 / 78%)', 'color': 'white', textAlign: 'center' }">
                <!-- <el-table-column prop="urlId" label="编号" width="100" /> -->
                <el-table-column type="selection" width="55" />
                <el-table-column label="ID" width="60">
                    <template #default="scope">
                        <el-tooltip :content="scope.row.id" placement="top" effect="light">
                            <span class="highlight">{{ scope.row.id }}</span>
                        </el-tooltip>
                    </template>
                </el-table-column>
                <el-table-column label="IP">
                    <template #default="scope">
                        <el-tooltip :content="scope.row.ip" placement="top" effect="light">
                            <span>{{ scope.row.ip }}</span>
                        </el-tooltip>
                    </template>
                </el-table-column>
                <el-table-column label="City">
                    <template #default="scope">
                        <el-tooltip :content="scope.row.city" placement="top" effect="light">
                            <span>{{ scope.row.city }}</span>
                        </el-tooltip>
                    </template>
                </el-table-column>
                <el-table-column label="Time">
                    <template #default="scope">
                        <el-tooltip :content="scope.row.time" placement="top" effect="light">
                            <span>{{ scope.row.time }}</span>
                        </el-tooltip>
                    </template>
                </el-table-column>
                <el-table-column label="操作" width="150">
                    <template #default="scope">
                        <el-popconfirm confirm-button-text="确定" cancel-button-text="取消" :icon="Delete"
                            icon-color="#626AEF" :title="'确定删除吗？'" @confirm="handleDelete(scope.$index, scope.row)">
                            <template #reference>
                                <el-button size="small" type="danger" style="margin-bottom: 10px;">删除</el-button>
                            </template>
                        </el-popconfirm>
                    </template>
                </el-table-column>
            </el-table>
            <!--
			    total 总行数
			    page-size	每页显示条目个数
			    current-change 改变页码时触发
			    currentPage:当前页码
			 -->
            <el-pagination class="showPaging" background layout="prev, pager, next" :total="total" :page-size="pageSize"
                :disabled="disabled" @current-change="changePage" v-model:currentPage="currentPage"
                @size-change="handleSizeChange" :pager-count="3" :hide-on-single-page="true" />
            <el-pagination class="hidePaging" background layout="prev, pager, next, jumper" :total="total"
                :page-size="pageSize" :disabled="disabled" @current-change="changePage"
                v-model:currentPage="currentPage" @size-change="handleSizeChange" :pager-count="3"
                :hide-on-single-page="true" />
        </div>
    </el-card>
</template>

<script>
import {
    ref,
    toRefs,
    reactive,
    onMounted,
    watch
} from 'vue'
import {
    Search,
    Delete
} from '@element-plus/icons-vue'
import { ElTable } from 'element-plus'
import axios from "axios";
import {
    ElMessage,
    ElNotification,
    ElMessageBox
} from 'element-plus'
const multipleTableRef = ref(ElTable)
/**
 * 加载全部数据
 */
function loadData(state) {
    state.loading = true
    var params = {
        'currentPage': state.currentPage,
        'pageSize': state.pageSize
    }
    axios.get("/ip/getall", {
        params
    }).then(res => {
        //JSON.parse 将从后台得到的数据转换为标准JSON格式
        //前台展示的是需要数组，JSON.parse转换后的数据，element-plus可以解析
        state.tableData = res.data.data.list;
        state.total = res.data.data.total;
        state.pageSize = res.data.data.pageSize;
        state.currentPage = res.data.data.currentPage;
        state.loading = false
    })
}

export default {
    components: {
    },
    setup(props) {
        //挂载后加载数据
        onMounted(() => {
            loadData(state);
        });
        // 图片 根url
        const url = process.env.VUE_APP_URL;
        const state = reactive({
            // 表格全部信息
            tableData: [],
            total: 0, //总条数
            pageSize: 10, //每页显示行数
            currentPage: 1, //当前页码
            // 当前审核链接的信息
            loading: false,
            deleteId: []
        })

        /**
        * 模块名:
        * 代码描述: 批量删除
        * 作者:lizibin
        * 创建时间:2023/01/02 14:44:57
        */
        const deleteAll = () => {
            if (state.deleteId.length === 0) {
                ElNotification({
                    title: '失败',
                    message: '未选择数据',
                    type: 'error'
                })
            } else {
                axios.delete("/ip/deleteall", {
                    params: {
                        id: state.deleteId.toString()
                    }
                }).then(res => {
                    if (res.data.code == "0") {
                        //通知提示框
                        ElNotification({
                            title: '成功',
                            message: '删除成功',
                            type: 'success',
                        });
                        loadData(state);
                    } else {
                        ElNotification({
                            title: '失败',
                            message: '删除失败',
                            type: 'error'
                        })
                    }
                })
            }
        }
        const handleSelectionChange = (val) => {
            let id = []
            val.forEach(e => {
                id.push(e.id)
            });
            state.deleteId = id
            console.log(id);
        }

        // 切换页面的执行事件，  val 当前页码
        const changePage = (val) => {
            state.currentPage = val;
            //重新加载数据
            loadData(state);
        };

        // 分页序号不乱
        const Nindex = (index) => {
            // 当前页数 - 1 * 每页数据条数 + 1
            const page = state.currentPage // 当前页码
            const pagesize = state.pageSize // 每页条数
            return index + 1 + (page - 1) * pagesize
        }

        // 刷新按钮
        const refresh = () => {
            // 搜索表单内容
            state.searchValue = ""
            // 筛选下拉框内容
            state.auditValue = null
            // 更新数据
            loadData(state);

            ElMessage({
                message: '刷新成功',
                type: 'success',
            })
        }

        // 监听下拉框内容的改变
        watch(() => state.auditValue, (val, preVal) => {
            if (val) {
                state.searchValue = "";
                auditScreen(state)
            }
        })

        // 增加删除事件
        function useDelete(state) {
            const handleDelete = (index, row) => {
                //确定删除
                axios.delete("/ip/delete", {
                    params: {
                        id: row.id
                    }
                }).then(res => {
                    if (res.data.code == "0") {
                        //通知提示框
                        ElNotification({
                            title: '成功',
                            message: '删除成功',
                            type: 'success',
                        });
                        loadData(state);
                    } else {
                        ElNotification({
                            title: '失败',
                            message: '删除失败',
                            type: 'error'
                        })
                    }
                })
            };
            return {
                handleDelete,
            }

        };
        return {
            // 搜索、删除图标
            Search,
            Delete,
            ...toRefs(state),
            // 搜索事件
            changePage,
            ...useDelete(state),
            Nindex,
            refresh,
            handleSelectionChange,
            deleteAll
        }
    }
}
</script>

<style scoped>
.card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.card-header h3 {
    display: inline-flex;
    justify-content: center;
    align-items: center;
}

:deep(.el-card__header) {
    border-bottom: 1px solid rgb(238 238 238);
    /* background: linear-gradient(to right, #535b9a, #30bcd7); */
    color: #2fa7b9;
}

.text {
    font-size: 14px;
}

.item {
    margin-bottom: 18px;
}

.el-card {
    border-radius: 0px;
    border: none;
}

.el-card.is-always-shadow {
    box-shadow: 0px 0px 0px rgb(0 0 0 / 12%);
}

.demo-image__error .image-slot {
    font-size: 30px;
}

.demo-image__error .image-slot .el-icon {
    font-size: 30px;
}

.demo-image__error .el-image {
    width: 100%;
    height: 200px;
}

.example-showcase .el-dropdown+.el-dropdown {
    margin-left: 15px;
}

.example-showcase .el-dropdown-link {
    cursor: pointer;
    color: var(--el-color-primary);
    display: flex;
    align-items: center;
}

@media screen and (min-width: 992px) {
    .visible {
        display: none !important;
    }
}

@media screen and (max-width: 992px) {
    .visible {
        display: block !important;
    }
}


/* 分页样式 */
:deep(.el-pagination.is-background .el-pager li:not(.is-disabled).is-active) {
    background-color: #2fa7b9;
}

.el-pagination {
    margin-top: 20px;
    justify-content: center;
}

:deep(.el-table .cell) {
    -webkit-box-orient: vertical;
    display: -webkit-box;
    -webkit-line-clamp: 2;
}

@media screen and (max-width: 500px) {
    .hidePaging {
        display: none;
    }
}

@media screen and (min-width: 500px) {
    .showPaging {
        display: none;
    }
}
</style>