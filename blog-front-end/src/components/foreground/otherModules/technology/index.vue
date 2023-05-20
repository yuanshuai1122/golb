<template>
    <div class="technology" v-loading="loading" element-loading-text="玩命加载中...">
        <div class="technology_title">
            <h3 class="title"><b>技术</b></h3>
        </div>
        <div class="technology_layout">
            <div v-for="(item, index) in articleInfo">
                <div class="contentBox item">
                    <a :href="'/articleDetails/' + item.articleId" :title="item.articleTitle">
                        <div class="articleContent">
                            <h4 class="title">{{ item.articleTitle }}</h4>
                            <p>{{ item.articleClassifyName }} · {{ item.publishTime }}</p>
                        </div>
                    </a>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
import {
    val
} from "dom7";
import {
    forEach
} from "lodash";
import {
    useRouter
} from 'vue-router'
import {
    onMounted,
    reactive,
    toRefs,
} from "vue";
import { FormatDate } from '@/utils/index'
function loadData(state) {
    state.loading = true;
    // 查询系统设置中设置好的精选文章的文章id
    var data = JSON.parse(localStorage.getItem('SystemData'))
    // axios.get("/systemSetup/showAllSystemSetup").then(res => {
    state.technologyData = JSON.parse(data.technologyArticle);
    // })

    // 查询全部文章
    var articleInfo = JSON.parse(localStorage.getItem('AllArticleInfo'))
    // axios.get("/article/showAllArticleInfo").then(res => {
    articleInfo.forEach(element => {
        // 筛选出系统设置总设置好的精选文章
        state.technologyData.forEach(val => {
            if (element.articleId === val) {
                // 时间戳格式化
                element.publishTime = FormatDate(element.publishTime);
                // 图片 根url
                const url = process.env.VUE_APP_URL;
                // 缩略图 判断是点击上传的还是，网络图片
                if (element.articleImgLitimg != "" && !element.articleImgLitimg
                    .includes('http') && !element.articleImgLitimg
                        .includes('https')) {
                    element.articleImgLitimg = url + element.articleImgLitimg
                }

                // 添加
                state.articleInfo.unshift(element)
            }
        });
    });
    state.loading = false;
    // })
}
export default {
    setup() {
        const state = reactive({
            // 精选文章的文章id
            technologyData: [],
            // 筛选出来的文章内容
            articleInfo: [],
            loading: false
        });
        //挂载后加载数据
        onMounted(() => {
            loadData(state);
        });
        return {
            //...toRefs 对响应式数据 解构赋值   
            ...toRefs(state),
        };
    }
}
</script>

<style scoped>
.technology {
    background: rgb(255, 255, 255);
    border-radius: 6px;
    padding: 20px;
    margin-bottom: 20px;
}

.technology_title {
    margin-bottom: 15px;
    margin-top: -10px;
    display: flex;
    justify-content: space-between;
    border-bottom: 1px solid rgb(238, 238, 238);
    overflow: hidden;
}

.technology_title .title {
    line-height: 40px;
    color: #474749;
}

.technology_title .title b {
    display: inline-block;
    border-bottom: 4px solid #2fa7b9;
}

.technology_layout .contentBox {
    margin-bottom: 15px;
    position: relative;
    overflow: hidden;
    font-size: 12px;
    color: #999;
    line-height: 18px;
}

.technology_layout .contentBox h4:hover {
    color: #2fa7b9;
}

.technology_layout .contentBox span {
    cursor: pointer;
}

.technology_layout .title {
    font-size: 14px;
    line-height: 24px;
    margin-bottom: 10px;
    color: #454545;
}

.technology_layout p {
    font-size: 12px;
    color: #999;
}

.technology_layout .contentBox.item {
    padding: 10px 0;
    border-bottom: 1px solid #eee;
}

.technology_layout .contentBox.item .title {
    font-size: 15px;
    overflow: hidden;
    -webkit-box-orient: vertical;
    display: -webkit-box;
    -webkit-line-clamp: 2;
    text-align: justify;
    cursor: pointer;
    line-height: 20px;
    color: #727272;
}

.technology_layout .contentBox.item .title:hover {
    color: #2fa7b9;
}

.articleContent {
    width: 100%;
    float: left;
    word-break: break-all;
    text-align: justify;
}
</style>