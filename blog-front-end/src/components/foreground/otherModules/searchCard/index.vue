<template>
    <!-- 卡片/搜索功能 -->
    <el-card class="search_card">
        <div class="search_img_box">
            <div><img :src="uploadURL + url" /></div>
            <div class="text">
                <div>镇</div>
                <div>站</div>
                <div>神</div>
                <div>兽</div>
            </div>
        </div>
        <div class="search_input_box">
            <!-- 搜索框 -->
            <el-input v-model="input" class="w-50 m-2 search_input" size="large" placeholder="请输入关键字进行搜索"
                clearable="true" :prefix-icon="Search" @keyup.enter.native="searchContent('')" />
        </div>
        <div class="search_link_box">
            <p>
                <el-link type="info" :underline="false" @click="searchContent('Golang')">Golang</el-link>
                <el-link type="info" :underline="false" @click="searchContent('资源')">资源</el-link>
                <el-link type="info" :underline="false" @click="searchContent('VUE')">VUE</el-link>
                <el-link type="info" :underline="false" @click="searchContent('技术')">技术</el-link>
            </p>

        </div>
    </el-card>
</template>

<script>
import {
    onBeforeMount,
    onMounted,
    reactive,
    toRefs
} from "vue"
import {
    Search
} from '@element-plus/icons-vue'
import {
    ElMessage
} from 'element-plus'
import {
    useRouter
} from 'vue-router'


export default {
    setup() {
        onMounted(() => {
            var url = JSON.parse(localStorage.getItem('SystemData')).advertising4;
            state.url = url
            // alert(state.nonce)
        })
        const state = reactive({
            input: '',
            nonce: null,
            url: null,
            uploadURL: process.env.VUE_APP_URL
        });

        const router = new useRouter();

        const searchContent = (val) => {
            if (val == "") {
                if (state.input !== "") {
                    router.push({
                        path: "/search",
                        query: {
                            keyword: state.input
                        }
                    })
                } else {
                    ElMessage.error('请输入关键字进行搜索')
                }
            } else {
                router.push({
                    path: "/search",
                    query: {
                        keyword: val
                    }
                })
            }


        }

        return {
            // ...toRefs 对响应式数据解构赋值
            ...toRefs(state),
            Search,
            searchContent
        };
    }
};
</script>

<style scoped>
.el-card {
    border-radius: 10px;
    margin-bottom: 20px;
}

:deep(.el-card__body) {
    padding: 0;
}

.search_img_box {
    display: flex;
    width: 100%;
    overflow: hidden;
    margin-bottom: 20px;
}

.search_img_box img {
    width: 190px;
    height: 190px;
    -o-object-fit: cover;
    object-fit: cover;
}

.search_img_box .text {
    font-family: "微软雅黑";
    -webkit-text-stroke: 1px #1d1d20;
    /*文字描边*/
    -webkit-text-fill-color: transparent;
    /*设置文字的填充颜色*/
    font-weight: 600;
    width: 103px;
    height: 190px;
    display: flex;
    flex-direction: column;
    justify-content: space-evenly;
    align-items: center;
}

.search_input_box {
    text-align: center;
    font-size: 12px;
    margin-bottom: 8px;
}

.search_input {
    width: 80%;
    border-radius: 50px;
}

.search_link_box {
    height: 18%;
    padding: 0px 15%;
    margin-bottom: 30px;
}

.search_link_box p {
    white-space: nowrap;
    text-overflow: ellipsis;
    overflow: hidden;
}

:deep(.el-input--large .el-input__wrapper) {
    margin-top: 0px;
    border-radius: 50px;
}

.el-link {
    margin-right: 10px;
}

.el-link .el-icon--right.el-icon {
    vertical-align: text-bottom;
}
</style>