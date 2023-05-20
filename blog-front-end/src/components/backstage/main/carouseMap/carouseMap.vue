<template>
    <div class="systemSetup">
        <el-row :gutter="20">
            <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24">
                <div class="left_box box-card">
                    <h3 class="title">
                        <el-icon style="margin-right: 10px;">
                            <Setting />
                        </el-icon>
                        首页展示设置
                    </h3>
                    <div class="content">
                        <h4 class="text-color">轮播图设置</h4>
                        <p class="text-color">提示：1、最多上次5张图片，单张大小不超过5M。</p>
                        <el-upload
                            v-model:file-list="fileList"
                            :on-success="handleAvatarSuccess" :before-upload="beforeUploadFunction"
                            :action="uploadURL"
                            :limit="5"
                            list-type="picture-card"
                            :on-preview="handlePictureCardPreview"
                            :on-remove="handleRemove"
                        >
                            <el-icon><Plus /></el-icon>
                        </el-upload>

                        <el-dialog v-model="dialogVisible">
                            <img w-full :src="dialogImageUrl" style="width:100%" alt="Preview Image" />
                        </el-dialog>
                    </div>
                    <div class="content" style="text-align: center;">
                        <el-button plain color="#2fa7b9" @click="leftSaveSettings">保存设置</el-button>
                    </div>
                </div>
            </el-col>
        </el-row>
    </div>

</template>
<script>
    import {
        ref,
        toRefs,
        reactive,
        onMounted,
        watch
    } from 'vue'
    import axios from "axios";
    import {
        ElMessage,
        ElNotification,
        ElMessageBox
    } from 'element-plus'
// import { fa } from 'element-plus/es/locale';
    /**
     * @desc 加载获取数据
     */
    function loadData(state) {
        // 查询全部文章
        axios.get("/carouseMap/searchCarouseMap").then(res => {
            let data = res.data.data;
            data.forEach(e => {
                e.response = {
                    data: e.url
                }
                e.url = state.baseUrl + e.url
            });
            state.fileList = data;
        })
    }
    export default ({
        setup() {
            //挂载后加载数据
            onMounted(() => {
                loadData(state);
            });
            const state = reactive({
                // 广告上传到服务器的路径
                uploadURL: process.env.VUE_APP_URL + "upload/advertisingImgUpload",
                baseUrl: process.env.VUE_APP_URL,
                fileList:[],
                dialogImageUrl: "",
                dialogVisible: false,
                // disabledFlag: false,
            })
            // 图片上传格式、大小要求
            const beforeUploadFunction = (rawFile) => {
                if (rawFile.type !== 'image/jpeg' &&
                    rawFile.type !== 'image/jpg' &&
                    rawFile.type !== 'image/png' &&
                    rawFile.type !== 'image/gif') {
                    ElMessage.error('仅支持图片格式.png | .jpg | .jpeg | .gif ')
                    return false
                } else if (rawFile.size / 1024 / 1024 > 5) {
                    ElMessage.error('仅支持大小不超过5MB的图片!')
                    return false
                }
                return true
            }

            // 图片上传成功后执行的函数
            const handleAvatarSuccess = (response) => {
                let data = JSON.parse(JSON.stringify(state.fileList));
                // if(data.length >= 5) {
                //     state.disabledFlag = true
                // }
            }

            const leftSaveSettings = () => {
                let data = JSON.parse(JSON.stringify(state.fileList));
                let sendArr = [];
                data.forEach(e => {
                    let obj = {
                        name: e.name,
                        url: e.response.data
                    }
                    sendArr.push(obj)
                });
                const params = sendArr
                
                axios.post("/carouseMap/updateCarouseMap", params).then((res) => {
                    if (res.data.code == 0) {
                        loadData(state);
                        ElMessage({
                            message: '保存成功',
                            type: 'success',
                        })
                        // 刷新界面
                        // location.reload()
                    }
                })
            }

            const handlePictureCardPreview = (uploadFile) => {
                 state.dialogImageUrl = uploadFile.url
                 state.dialogVisible = true
            }

            return {
                ...toRefs(state),
                beforeUploadFunction,
                handleAvatarSuccess,
                leftSaveSettings,
                handlePictureCardPreview
            }
        }
    })
</script>

<style>
    .systemSetup .el-col {
        margin-bottom: 20px;
    }

    .systemSetup .title {
        color: #2fa7b9;
        margin-bottom: 10px;
        padding: 10px 20px;
        display: inline-flex;
        justify-content: center;
        align-items: center;
    }


    .left_box {
        width: 100%;
        height: auto;
        background: white;
        padding: 20px;
        /* 添加此属性 padding间距不扩大div */
        box-sizing: border-box;
    }

    .right_box {
        width: 100%;
        height: auto;
        background: white;
        padding: 20px;
        /* 添加此属性 padding间距不扩大div */
        box-sizing: border-box;
    }

    .systemSetup .content {
        width: 100%;
        padding: 10px 20px;
        /* 添加此属性 padding间距不扩大div */
        box-sizing: border-box;
    }

    .systemSetup .content h4 {
        line-height: 45px;
        color: #8f8f8f;
    }

    .systemSetup .content h5 {
        line-height: 35px;
        color: #8f8f8f;
        font-size: 15px;
        font-weight: lighter;
    }

    .systemSetup .content p {
        line-height: 25px;
        color: #8f8f8f;
        font-size: 12px;
    }
</style>