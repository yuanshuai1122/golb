<template>
    <!-- 菜单样式可以改成可配置 -->
    <el-menu :default-active="route.path" class="el-menu-vertical-mode" background-color="#0e111b" color="white" text-color="#c1c1c1" router
        :unique-opened="false" unique-opened :default-openeds="[route.path]">
        <div class="imgBox hidden-xs-only" >
            <img src="../../../assets/logo.svg">
        </div>
        <template v-for="v in menuData" :key="v.id" style="width:100%;">
            <el-sub-menu v-if="v.childMenu && v.childMenu.length>0" index="/articleOperation">
                <template #title >
                    <el-icon>
                        <component :is="v.icon"></component>
                    </el-icon>
                    {{v.chineseName}}
                </template>
                <el-menu-item v-for="vite in v.childMenu" :key="vite.id" :index="vite.path"  @click="byValue">
                    <template #title>
                        <el-icon>
                            <component :is="vite.icon"></component>
                        </el-icon>{{vite.chineseName}}
                    </template>
                </el-menu-item>
            </el-sub-menu>
            <el-menu-item v-else :index="v.path"  @click="byValue">
                <template #title>
                    <el-icon>
                        <component :is="v.icon"></component>
                    </el-icon>{{v.chineseName}}
                </template>
            </el-menu-item>
        </template>
    </el-menu>
</template>
<script>
    import {
        reactive,
        ref,
        onMounted
    } from 'vue'

    import {
        useRoute
    } from "vue-router"
    import axios from "axios"

    import {
        ElMessage,
        ElLoading
    } from 'element-plus'

    export default {
        components: {},
        emits: ["onDrawerVisible"],
        setup(props, {
            emit
        }) {
            const route = useRoute();
            const menuData = ref();
            onMounted(() => {
                const tokenStr = sessionStorage.getItem("userInfo");
                
                if (tokenStr) {
                    const userInfo = JSON.parse(tokenStr);
                    menuData.value = userInfo.menuInfo
                }
            });

            const byValue = () =>{
                emit("onDrawerVisible", false);
            }
            return {
                route,
                onMounted,
                menuData,
                byValue
            };
        }
    }
</script>

<style scoped>
    .imgBox {
        width: 100%;
        height: 70px;
    }

    .imgBox img {
        width: 80%;
        height: 100%;
        margin: 0 10%;
    }

    .el-menu-vertical-demo:not(.el-menu--collapse) {
        width: 200px;
        min-height: 400px;
    }

    .el-menu {
        height: 100%;
        border: 0px;
        /* color: white !important; */
        /* background-color: #0e111b!important;  */
        /* 暗夜黑 */
    }

    /* 选中 */
    :deep(.el-menu-item.is-active) {
        color: white;
        background: linear-gradient(to right, #a5aad3, #202121);
    }
</style>