<template>
    <div class="resource" v-loading="loading" element-loading-text="玩命加载中...">
        <div class="resource_title">
            <h3 class="title"><b>资源</b></h3>
        </div>
        <div class="resource_layout">
            <div v-for="(item, index) in articleInfo">
                <div class="contentBox" v-if="index === 0">
                    <a :href="'/articleDetails/' + item.articleId" :title="item.articleTitle">
                        <img v-if="item.articleImgLitimg" class="bannerImg" :src="item.articleImgLitimg"
                            :alt="item.articleTitle">
                        <h4 class="title">{{ item.articleTitle }}</h4>
                        <p>{{ item.articleClassifyName }} · {{ item.publishTime }}</p>
                    </a>
                </div>


                <div class="contentBox item" v-if="index > 0">
                    <spana :href="'/articleDetails/' + item.articleId" :title="item.articleTitle">
                        <!-- <div v-if="item.articleImgLitimg" class="banner" @click="articleDetails(item)">
                            <img :src="item.articleImgLitimg" :alt="item.articleTitle" @click="articleDetails(item)">
                        </div> -->
                        <!-- 判断是否有缩略图，没有则样式不同 -->
                        <div class="articleContent">
                            <h4 class="title">{{ item.articleTitle }}</h4>
                            <p>{{ item.articleClassifyName }} · {{ item.publishTime }}</p>
                        </div>
                    </spana>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
import axios from "axios";
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
import { FormatDate} from '@/utils/index'
function loadData(state) {
    state.loading = true;
    // 查询系统设置中设置好的资源文章的文章id
    var data = JSON.parse(localStorage.getItem('SystemData'))
    state.resourceData = JSON.parse(data.resourceArticle);
    // 查询全部文章
    var articleInfo = JSON.parse(localStorage.getItem('AllArticleInfo'))
    // axios.get("/article/showAllArticleInfo").then(res => {
    state.resourceData.forEach(val => {
        articleInfo.forEach(element => {
            // 筛选出系统设置总设置好的资源文章
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
                state.articleInfo.push(element)
            }
        });
    });
    state.loading = false;
    // })
}
export default {
    setup() {
        const state = reactive({
            // 资源文章的文章id
            resourceData: [],
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
.resource {
    background: rgb(255, 255, 255);
    border-radius: 6px;
    padding: 20px;
    margin-bottom: 20px;
}

.resource_title {
    margin-bottom: 15px;
    margin-top: -10px;
    display: flex;
    justify-content: space-between;
    border-bottom: 1px solid rgb(238, 238, 238);
    overflow: hidden;
}

.resource_title .title {
    line-height: 40px;
    color: #474749;
}

.resource_title .title b {
    display: inline-block;
    border-bottom: 4px solid #2fa7b9;
}

.resource_layout .contentBox {
    margin-bottom: 15px;
    position: relative;
    overflow: hidden;
    font-size: 12px;
    color: #999;
    line-height: 18px;
}

.resource_layout .contentBox h4:hover {
    color: #2fa7b9;
}

.resource_layout .contentBox span {
    cursor: pointer;
}

.resource_layout .bannerImg {
    display: block;
    width: 100%;
    object-fit: cover;
    margin-bottom: 10px;
    border-radius: 6px;
}

.resource_layout .title {
    font-size: 15px;
    line-height: 24px;
    margin-bottom: 10px;
    color: #727272;
    overflow: hidden;
    -webkit-box-orient: vertical;
    display: -webkit-box;
    -webkit-line-clamp: 2;
    text-align: justify;
    cursor: pointer;
}

.resource_layout p {
    font-size: 12px;
    color: #999;
}

.resource_layout .contentBox.item {
    padding-top: 15px;
    border-top: 1px solid #eee;
}

.resource_layout .contentBox.item .banner {
    width: 30%;
    height: 60px;
    overflow: hidden;
    float: left;
    border-radius: 4px;
    margin-right: 15px;
    cursor: pointer;
}

.resource_layout .contentBox.item .banner img {
    width: 100%;
    height: 100%;
    object-fit: cover;
}

/* .technology_layout .contentBox.item .title {
        font-size: 15px;
        overflow: hidden;
        -webkit-box-orient: vertical;
        display: -webkit-box;
        -webkit-line-clamp: 2;
        text-align: justify;
        cursor: pointer;
        line-height: 20px;
        color: #727272;
    } */

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