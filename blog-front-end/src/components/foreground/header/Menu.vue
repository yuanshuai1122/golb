<template>
    <div class="hidden-xs-only header-menu">
        <el-menu :default-active="route.path" @select="handleSect" :router="true" :ellipsis="false" class="el-menu-demo"
            mode="horizontal" active-text-color='#2FA7B9'>
            <img src="../../../assets/logo.svg" style="cursor: pointer;width:140px;" alt="TolZB"
                @click="router.push({ path: '/' })">
            <div class="flex-grow" id="fire">
                {{ activeText }} |
            </div>
            <el-menu-item index="/">
                <span>首页</span>
            </el-menu-item>
            <el-menu-item index="/article">
                <span>精品文章</span>
            </el-menu-item>
            <el-menu-item index="/website">
                <span>学习网站</span>
            </el-menu-item>
            <!-- <el-menu-item index="/message">
                <span>留言</span>
            </el-menu-item> -->
            <el-menu-item index="/about">
                <span>关于</span>
            </el-menu-item>
            <a href="/login" target="_blank" class="loginA" v-show="!userInfo">
                登录/注册
            </a>
            <el-dropdown v-show="userInfo" style="cursor: pointer;">
                <div class="el-dropdown-link" style="margin: auto;">
                    <img v-if="userInfo.userIcon !== null && userInfo.userIcon !== ''" :src="url + userInfo.userIcon"
                        style="width: 30px;margin: 0px 10px;vertical-align:middle;border-radius: 50px;" />
                    <img v-if="userInfo.userIcon === null || userInfo.userIcon === ''"
                        src="../../../assets/default_avatar.svg"
                        style="width: 30px;margin: 0px 10px;vertical-align:middle;border-radius: 50px;">
                    <span>
                        {{ userInfo.userName }}
                    </span>

                    <el-icon class="el-icon--right">
                        <arrow-down />
                    </el-icon>
                </div>
                <template #dropdown>
                    <el-dropdown-menu v-if="userInfo.userType === 0">
                        <el-dropdown-item @click="toPath('/addArticle')">新建文章</el-dropdown-item>
                        <el-dropdown-item @click="toPath('/articleManager')">文章管理</el-dropdown-item>
                        <el-dropdown-item @click="toPath('/classifyManager')">分类管理</el-dropdown-item>
                        <el-dropdown-item @click="toPath('/commentManager')">评论管理</el-dropdown-item>
                        <el-dropdown-item @click="toPath('/userManager')">用户管理</el-dropdown-item>
                        <el-dropdown-item @click="toPath('/linkManager')">链接管理</el-dropdown-item>
                        <el-dropdown-item @click="toPath('/personSetup')">个人设置</el-dropdown-item>
                        <el-dropdown-item @click="toPath('/systemSetup')">系统设置</el-dropdown-item>
                        <el-dropdown-item @click="toPath('/carouseMap')">轮播图</el-dropdown-item>
                        <el-dropdown-item @click="toPath('/circulateText')">循环字</el-dropdown-item>
                        <el-dropdown-item @click="toPath('/visitIp')">访问信息</el-dropdown-item>
                        <el-dropdown-item divided @click="exit">退出登录</el-dropdown-item>
                    </el-dropdown-menu>
                    <el-dropdown-menu v-else>
                        <el-dropdown-item @click="toPath('/addArticle')">新建文章</el-dropdown-item>
                        <el-dropdown-item @click="toPath('/articleManager')">文章管理</el-dropdown-item>
                        <el-dropdown-item @click="toPath('/commentManager')">评论管理</el-dropdown-item>
                        <el-dropdown-item @click="toPath('/personSetup')">个人设置</el-dropdown-item>
                        <el-dropdown-item divided @click="exit">退出登录</el-dropdown-item>
                    </el-dropdown-menu>
                </template>
            </el-dropdown>
        </el-menu>
    </div>

    <div class="hidden-sm-and-up header-img">
        <img src="../../../assets/logo.svg" style="cursor: pointer;width:140px;margin: 0 auto;" alt="TolZB"
            @click="router.push({ path: '/' })">
    </div>

</template>

<script>

import {
    reactive,
    toRefs,
    onMounted,
    onBeforeUnmount
} from 'vue'

import {
    useRoute,
    useRouter
} from "vue-router"
import axios from "axios";

let textTimer = null;

export default {
    setup() {
        const route = useRoute();
        const router = useRouter();
        const state = reactive({
            // activeIndex: sessionStorage.getItem('activeIndex1'), // 默认为空
            // screenWidth: document.body.clientWidth, //屏幕可视区宽度
            // 当前登录的用户信息
            userInfo: '',
            // 服务器路径
            url: process.env.VUE_APP_URL,
            top_text: [],
            activeText: '',
        });
        onMounted(() => {
            loadData(state);
            // 获取用户信息
            const tokenStr = sessionStorage.getItem("userInfo");
            if (tokenStr) {
                const userInfo = JSON.parse(tokenStr);
                state.userInfo = userInfo.data[0];
                // console.log(state.userInfo.userType);
            }
        });
        onBeforeUnmount(() => {
            clearInterval(textTimer)
        });
        const loadData = (state) => {
            // 查询循环字
            const params = {
                'currentPage': 1,
                'pageSize': 9999999,
            }
            axios.get("/text/getall", {
                params
            }).then(res => {
                let arr = []
                res.data.data.list.forEach(e => {
                    arr.push(e.text)
                });
                state.top_text = arr;
                circulateText()
            })
        }
        // 退出登录
        const exit = () => {
            // 清除用户登录信息
            window.sessionStorage.removeItem("userInfo");
            // 刷新界面
            window.location.reload(true);
        }
        // 前往某个界面
        const toPath = (path) => {
            window.location.href = path
        }
        /**
        * 模块名:
        * 代码描述: 文字循环
        * 作者:lizibin
        * 创建时间:2023/01/02 09:56:46
        */
        const circulateText = () => {
            // 随机选择textArr的一个值
            let random = Math.floor(Math.random() * (state.top_text.length));
            let nowText = state.top_text[random];
            let i = 0;
            textTimer = setInterval(function () {
                if (state.activeText.length < nowText.length) {
                    state.activeText += nowText[i++];
                } else {
                    state.activeText = ''
                    clearInterval(textTimer)
                    circulateText()
                }
            }, 300)
        }
        // // 菜单激活回调  index:path 路由 当前选择菜单的路由 index
        // const handleSect = ((index) => {
        //     sessionStorage.setItem('activeIndex1', index)
        // })
        // // 捕获 听器  获取窗口宽高
        // window.addEventListener('resize', () => {
        //     // 窗口宽度 screenWidth
        //     state.screenWidth = document.body.clientWidth;
        // })
        // // 监听器1：当前监听当前窗口是否改变，val为修改后的值,preVal为修改前的值
        // watch(() => state.screenWidth, (val, preVal) => {
        //     // 判断窗口改变后是否大于768px，大于则将sessionStorage中存储的值= activeIndex(当前选中菜单的index)
        //     // sessionStorage用于存储两个菜单其中一个更改了选项的路由(就是菜单的index)
        //     if (val > 768) {
        //         // sessionStorage存储的值传给activeIndex
        //         state.activeIndex = sessionStorage.getItem('activeIndex1');
        //         // 监听器2：当前监听sessionStorage第一次进入界面的时候是为空的
        //         // 那么就应该给sessionStorage赋初始值('/')
        //         watch(() => sessionStorage.getItem('activeIndex1'), (val,
        //             preVal) => { // val为修改后的值,preVal为修改前的值
        //             // 为什么是val不是preVal，因为我们赋初始值时赋的时sessionStorage.getItem('activeIndex1')
        //             // 一开始sessionStorage.getItem('activeIndex1')是没有数据的，所有为空，则需要在这进行非空验证
        //             // 为空则赋值'/'(首页)
        //             if (val == null) {
        //                 // 为空则菜单的index就为'/'(首页)
        //                 state.activeIndex = "/"
        //             }
        //             console.log(val);
        //             console.log(val == null);
        //             console.log(val == "");
        //         }, {
        //             immediate: true, // 立刻监听
        //             deep: true, // 深度监听
        //         })
        //     }
        // }, {
        //     immediate: true, // 立刻监听
        //     deep: true, // 深度监听
        // })
        return {
            ...toRefs(state),
            // handleSect,
            route,
            router,
            exit,
            toPath,
        }
    }
}
</script>

<style scoped>
/* 顶部菜单栏 */
.flex-grow {
    flex-grow: 1;
}

.el-menu--horizontal {
    border-bottom: none;
}

.el-menu--horizontal .el-menu-item:not(.is-disabled):focus,
.el-menu--horizontal .el-menu-item:not(.is-disabled):hover {
    color: #2FA7B9;
    background-color: rgba(47, 167, 185, 0.1);
}

.header-img {
    display: flex;
    align-items: center;
    text-align: center;
    line-height: 90px;
}

.el-menu--horizontal>.el-menu-item a {
    color: #303133;
}

.el-menu--horizontal>.el-menu-item a:hover {
    color: #2FA7B9;
}

.loginA {
    line-height: 56px;
    font-size: 14px;
    margin-left: 10px;
}

.example-showcase .el-dropdown-link {
    cursor: pointer;
    color: var(--el-color-primary);
    display: flex;
    align-items: center;
}

:deep(.el-dropdown-menu__item:not(.is-disabled):focus) {
    background-color: rgba(47, 167, 185, 0.1);
    color: #2fa7b9;
}

#fire {
    text-align: center;
    letter-spacing: 1px;
    line-height: 58px;
    font-family: '方正启体简体';
}
</style>