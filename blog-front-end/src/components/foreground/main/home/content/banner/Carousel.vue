<template>
    <!-- 走马灯/轮播图 -->
    <el-carousel trigger="hover" :autoplay="true">
        <el-carousel-item v-for="(img,index) in imgs" :key="index">
            <img :src="img.url">
            <h3 class="small justify-center" text="2xl">{{ img.title }}</h3>
        </el-carousel-item>
    </el-carousel>
    <!-- <el-carousel type="card" :autoplay="false" arrow="never" ref="slideCarousel" @change="setActiveItem">
        <el-carousel-item v-for="item in imgs" :key="item.id">
            <img class="element-img" alt="暂无" :src="item.url" @click="uploadDialogOpen">
        </el-carousel-item>
    </el-carousel> -->
</template>
<script>
 import axios from "axios";
    import {
        onMounted,
        reactive,
        toRefs,
    } from "vue"

    function loadData(state) {
        // 查询全部文章
        axios.get("/carouseMap/searchCarouseMap").then(res => {
            let data = res.data.data;
            const url = process.env.VUE_APP_URL;
            data.forEach(e => {
                e.url = url+e.url
            });
            state.imgs = data;
        })
    }

    export default {
        setup() {
           //挂载后加载数据
           onMounted(() => {
                loadData(state);
            });
            const state = reactive({
                // baseUrl: process.env.VUE_APP_URL,
                imgs: [],
            })
            // const imgs = [{
            //         name: "banner01",
            //         url: require('../../../../../../assets/banner/banner02.png'),
            //         // title: "人生能有几回博，此时不博何时博;"
            //     }
            //     , {
            //         name: "banner02",
            //         url: require('../../../../../../assets/banner/banner03.png'),
            //         // title: "青，取之于蓝而青于蓝;"
            //     }, {
            //         name: "banner03",
            //         url: require('../../../../../../assets/banner/banner04.png'),
            //         // title: "冰，水为之而寒于水；"
            //     }, {
            //         name: "banner04",
            //         url: require('../../../../../../assets/banner/banner05.png'),
            //         // title: "业精于勤而荒于嬉，行成于思而毁于随；"
            //     }, {
            //         name: "banner05",
            //         url: require('../../../../../../assets/banner/banner06.png'),
            //         // title: "人生能有几回博，此时不博何时博;"
            //     }
            // ]
            
            // 滑动切换
            const slideBanner = () => {
                //选中的轮播图
                var box = document.querySelector('.el-carousel__container');
                var startPoint = 0;
                var stopPoint = 0;
                //重置坐标
                var resetPoint = function () {
                    startPoint = 0;
                    stopPoint = 0;
                }
                //手指按下
                box.addEventListener("touchstart", function (e) {
                    //手指点击位置的X坐标
                    startPoint = e.changedTouches[0].pageX;
                });
                //手指滑动
                box.addEventListener("touchmove", function (e) {
                    //手指滑动后终点位置X的坐标
                    stopPoint = e.changedTouches[0].pageX;
                });
                //当手指抬起的时候，判断图片滚动离左右的距离
                let that = this
                box.addEventListener("touchend", function (e) {
                    if (stopPoint == 0 || startPoint - stopPoint == 0) {
                        resetPoint();
                        return;
                    }
                    if (startPoint - stopPoint > 0) {
                        resetPoint();
                        that.$refs.slideCarousel.next();
                        return;
                    }
                    if (startPoint - stopPoint < 0) {
                        resetPoint();
                        that.$refs.slideCarousel.prev();
                        return;
                    }
                });
            }

            return {
                ...toRefs(state),
                slideBanner,
            }
        }
    };
</script>

<style scoped>
    /* 走马灯、轮播图样式 */
    .el-carousel {
        background: white;
        border-radius: 10px;
        overflow: hidden;
    }

    .el-carousel__item img {
        display: block;
        width: 100%;
        height: 100%;
        overflow: hidden;
        object-fit: cover;
        background-size: 100% 100%;
        background-repeat:no-repeat;
    }

    /* css注释：设置了浏览器宽度不大于768px时 */
    @media screen and (max-width: 768px) {
        :deep(.el-carousel__container) {
            height: 200px;
        }

        :deep(.el-carousel__button) {
            width: 20px;
        }
    }

    @media screen and (max-width: 500px) {
        :deep(.el-carousel__container) {
            height: 150px;
        }

        :deep(.el-carousel__button) {
            width: 10px;
        }
    }


    .el-carousel__item h3 {
        color: white;
        opacity: 0.75;
        position: absolute;
        top: 0;
        bottom: 0;
        left: 0;
        right: 0;
        margin: auto 0;
        display: inline-flex;
        justify-content: center;
        align-items: center;
        text-align: center;
    }
</style>