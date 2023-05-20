<template>
    <div class="music-box">
        <iframe frameborder="no" border="0" marginwidth="0" marginheight="0" width=330 height=86 src="//music.163.com/outchain/player?type=2&id=1366216050&auto=1&height=66"></iframe>
    </div>
</template>

<script>
    import {
        ref,
        reactive,
        toRefs,
        onMounted
    } from 'vue'

    export default {
        setup() {
            const state = reactive({
                
            });
            onMounted(() => {
               
            });
            
            return {
                ...toRefs(state),
            }
        }
    }
</script>

<style scoped>
    .music-box {
        position: fixed;
        bottom: 0;
        left: 0;
        display: flex;
    }
    .music-box-left,
    .music-box-right{
        width: 20px;
        height: 86px;
        background-color: red;
    }
</style>