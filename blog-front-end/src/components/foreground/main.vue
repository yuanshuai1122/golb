<template>
    <div id="foreground">
        <div class="common-layout">
            <div id="wp" class="wp" v-if="effects04 === 1">
                <div class="xnkl">
                    <div class="deng-box2">
                        <div class="deng">
                            <div class="xian"></div>
                            <div class="deng-a">
                                <div class="deng-b">
                                    <div class="deng-t">度</div>
                                </div>
                            </div>
                            <div class="shui shui-a">
                                <div class="shui-c"></div>
                                <div class="shui-b"></div>
                            </div>
                        </div>
                    </div>
                    <div class="deng-box3">
                        <div class="deng">
                            <div class="xian"></div>
                            <div class="deng-a">
                                <div class="deng-b">
                                    <div class="deng-t">欢</div>
                                </div>
                            </div>
                            <div class="shui shui-a">
                                <div class="shui-c"></div>
                                <div class="shui-b"></div>
                            </div>
                        </div>
                    </div>
                    <div class="deng-box1">
                        <div class="deng">
                            <div class="xian"></div>
                            <div class="deng-a">
                                <div class="deng-b">
                                    <div class="deng-t">春</div>
                                </div>
                            </div>
                            <div class="shui shui-a">
                                <div class="shui-c"></div>
                                <div class="shui-b"></div>
                            </div>
                        </div>
                    </div>
                    <div class="deng-box">
                        <div class="deng">
                            <div class="xian"></div>
                            <div class="deng-a">
                                <div class="deng-b">
                                    <div class="deng-t">新</div>
                                </div>
                            </div>
                            <div class="shui shui-a">
                                <div class="shui-c"></div>
                                <div class="shui-b"></div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <el-container>
                <!-- 页眉 -->
                <el-header id="my-header">
                    <MenuTop></MenuTop>
                </el-header>
                <!-- 主要的 -->
                <el-main>
                    <router-view></router-view>
                </el-main>
                <!-- 页尾\页脚 -->
                <el-footer class="front_footer">
                    <FooterCopyright></FooterCopyright>
                </el-footer>
            </el-container>
        </div>

        <!-- Backtop 回到顶部 -->
        <BackTop></BackTop>

        <!-- 打开菜单栏固钉 -->
        <MenuRight></MenuRight>
        <Music></Music>
    </div>
</template>
<script>
// 引入顶部菜单栏自定义组件
import MenuTop from './header/Menu.vue';
// 引入页尾\页脚版权信息自定义组件
import FooterCopyright from './footer/Copyright.vue';
// 引入回到顶部固钉自定义组件
import BackTop from './affix/BackTop.vue';
// 引入打开菜单固钉自定义组件
import MenuRight from './affix/Menu.vue';
import Music from './music/Music.vue';

import {
    useRoute
} from "vue-router"
import {
    toRefs,
    onMounted,
    onBeforeUnmount,
    reactive,
    ref
} from 'vue'
import {
    ElNotification
} from 'element-plus'
import axios from "axios";
import { GetNowTime} from '@/utils/index'
function loadData(state) {
    // 查询系统设置
    let data = JSON.parse(localStorage.getItem('SystemData'))
    // 灯笼特效
    state.effects04 = data.effects04;
    axios.get("/article/showAllArticleInfo").then(res => {
        let data = JSON.stringify(res.data.data)
        localStorage.setItem("AllArticleInfo", data);
    })
}
/**
* 模块名: 弹出访问者ip+地址
* 代码描述: 弹出访问者ip+地址
* 作者:lizibin
* 创建时间:2022/12/31 20:00:49
*/
function getMachineIp() {
    fetch('http://ip-api.com/json/?lang=zh-CN')
        .then(res => res.json())
        .then(data => {
            // if (!localStorage.getItem('ipData')) {
            ElNotification({
                title: '网站提示',
                message: `欢迎IP:${data.query},来自${data.country}${data.regionName}${data.city}的朋友！`,
                position: 'bottom-right',
                type: 'success',
            })
            // }
            let ipData = JSON.stringify(data)
            localStorage.setItem('ipData', ipData)
            let time = GetNowTime();
            addVisitInfo(data.query, `${data.country}-${data.regionName}-${data.city}`,time)
        })
}
/**
* 模块名: addVisitInfo
* 代码描述:
* 作者:lizibin
* 创建时间:2023/01/02 16:00:56
*/
function addVisitInfo(ip, city,time) {
    const params = {
        ip,
        city,
        time
    }
    axios.post("/ip/insert", params).then((res) => {
        if (res.data.code == 0) {

        } else {
            ElNotification({
                title: '系统提示',
                message: '保存访问信息失败',
                type: 'error',
            });
        }
    })
}
/**
* 模块名:
* 代码描述: 高度判断
* 作者:lizibin
* 创建时间:2023/01/01 11:52:51
*/
function handleScrollx() {
    let top = document.documentElement.scrollTop || document.body.scrollTop;
    if (top > 60) {
        document.getElementById('my-header').style.display = 'none'
    } else {
        document.getElementById('my-header').style.display = 'block'
    }
}

export default {
    components: {
        MenuTop,
        FooterCopyright,
        BackTop,
        MenuRight,
        Music
    },
    setup() {
        onMounted(() => {
            loadData(state);
            getMachineIp();
            window.addEventListener("scroll", handleScrollx, true);
        });
        onBeforeUnmount(() => {
            // localStorage.removeItem("ipData"); 
        });
        const state = reactive({
            baseURL: process.env.VUE_APP_URL,
            effects04: null,
        })
        const route = useRoute()
        return {
            ...toRefs(state),
        }
    }
}

</script>

<style scoped>
/* 首页底部版权样式 */
.front_footer {
    border-top: 1px solid #ebebeb;
    background: #fff;
}

.ct2 .mn {
    width: 770px;
}

.ct2 .sd {
    width: 218px;
}

@media screen and (max-width:768px) {
    .xnkl {
        display: none;
    }
}

.deng-box {
    position: fixed;
    top: -40px;
    right: 150px;
    z-index: 9999;
    pointer-events: none;
}

.deng-box1 {
    position: fixed;
    top: -30px;
    right: 10px;
    z-index: 9999;
    pointer-events: none
}

.deng-box2 {
    position: fixed;
    top: -40px;
    left: 150px;
    z-index: 9999;
    pointer-events: none
}

.deng-box3 {
    position: fixed;
    top: -30px;
    left: 10px;
    z-index: 9999;
    pointer-events: none
}

.deng-box1 .deng,
.deng-box3 .deng {
    position: relative;
    width: 120px;
    height: 90px;
    margin: 50px;
    background: #d8000f;
    background: rgba(216, 0, 15, .8);
    border-radius: 50% 50%;
    -webkit-transform-origin: 50% -100px;
    -webkit-animation: swing 5s infinite ease-in-out;
    box-shadow: -5px 5px 30px 4px #fc903d
}

.deng {
    position: relative;
    width: 120px;
    height: 90px;
    margin: 50px;
    background: #d8000f;
    background: rgba(216, 0, 15, .8);
    border-radius: 50% 50%;
    -webkit-transform-origin: 50% -100px;
    -webkit-animation: swing 3s infinite ease-in-out;
    box-shadow: -5px 5px 50px 4px #fa6c00
}

.deng-a {
    width: 100px;
    height: 90px;
    background: #d8000f;
    background: rgba(216, 0, 15, .1);
    margin: 12px 8px 8px 8px;
    border-radius: 50% 50%;
    border: 2px solid #dc8f03
}

.deng-b {
    width: 45px;
    height: 90px;
    background: #d8000f;
    background: rgba(216, 0, 15, .1);
    margin: -4px 8px 8px 26px;
    border-radius: 50% 50%;
    border: 2px solid #dc8f03
}

.xian {
    position: absolute;
    top: -20px;
    left: 60px;
    width: 2px;
    height: 20px;
    background: #dc8f03
}

.shui-a {
    position: relative;
    width: 5px;
    height: 20px;
    margin: -5px 0 0 59px;
    -webkit-animation: swing 4s infinite ease-in-out;
    -webkit-transform-origin: 50% -45px;
    background: orange;
    border-radius: 0 0 5px 5px
}

.shui-b {
    position: absolute;
    top: 14px;
    left: -2px;
    width: 10px;
    height: 10px;
    background: #dc8f03;
    border-radius: 50%
}

.shui-c {
    position: absolute;
    top: 18px;
    left: -2px;
    width: 10px;
    height: 35px;
    background: orange;
    border-radius: 0 0 0 5px
}

.deng:before {
    position: absolute;
    top: -7px;
    left: 29px;
    height: 12px;
    width: 60px;
    content: " ";
    display: block;
    z-index: 999;
    border-radius: 5px 5px 0 0;
    border: solid 1px #dc8f03;
    background: orange;
    background: linear-gradient(to right, #dc8f03, orange, #dc8f03, orange, #dc8f03)
}

.deng:after {
    position: absolute;
    bottom: -7px;
    left: 10px;
    height: 12px;
    width: 60px;
    content: " ";
    display: block;
    margin-left: 20px;
    border-radius: 0 0 5px 5px;
    border: solid 1px #dc8f03;
    background: orange;
    background: linear-gradient(to right, #dc8f03, orange, #dc8f03, orange, #dc8f03)
}

.deng-t {
    font-family: 黑体, Arial, Lucida Grande, Tahoma, sans-serif;
    font-size: 3.2rem;
    color: #dc8f03;
    font-weight: 700;
    line-height: 85px;
    text-align: center
}

.night .deng-box,
.night .deng-box1,
.night .deng-t {
    background: 0 0 !important
}

@-moz-keyframes swing {
    0% {
        -moz-transform: rotate(-10deg)
    }

    50% {
        -moz-transform: rotate(10deg)
    }

    100% {
        -moz-transform: rotate(-10deg)
    }
}

@-webkit-keyframes swing {
    0% {
        -webkit-transform: rotate(-10deg)
    }

    50% {
        -webkit-transform: rotate(10deg)
    }

    100% {
        -webkit-transform: rotate(-10deg)
    }
}
</style>