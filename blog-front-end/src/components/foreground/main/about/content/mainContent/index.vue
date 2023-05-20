<template>
    <div class="main_content">
        <div>
            <div class="content_title">
                <h3>关于我</h3>
            </div>
            <div class="content_box">
                <p>本人学识渊博，经验丰富，代码风骚，效率恐怖。AI、UI设计、前后端无不精通，熟练掌握各种语言的框架。</p>
                <p>会DDOS一秒钟死一个站，会注入，会上传，会XSS，会破解，会嗅探，会开发，会业务，会运维，会渗透。互联网内，我无处不在。</p>
                <p>全国漏洞认识深刻，熟练掌握各种操作系统内核。认真学习过《黑客攻防技术宝典》，《逆向工程》，《游戏waigua攻防艺术》等专业书籍。深山苦练十余载，一天只睡2小时。</p>
                <p>听指挥，执行快，脾气好，不喷人，操作虎，意识强，渗透快，shell多，能过狗，能过盾。</p>
                <p>千里之外取服务器首级，瞬息之间爆管理员狗头，压安全狗如压草芥。电话通知出Bug后，秒登vpn，千里之外定位问题，瞬息之间修复上线。</p>
                <p>只有你想不到的，没有我做不到的。</p>
                <br>
                <p><b>以上都是吹的，哈哈。</b></p>
                <p>本人在职小前端一枚,熟悉Vue,React,Angular,JQuery,原生JavaScript</p>
                <p>在职期间还维护过PHP项目</p>
                <p>2021年4月实习,2021年7月正式开始打工，至今</p>
            </div>
        </div>
        <div>
            <div class="content_title">
                <h3>关于本站</h3>
            </div>
            <div class="content_box">
                <p>使用Golang重写了原作者Java的后端接口！</p>
                <p><b>建站时间：</b>2022年7月20日</p>
                <p><b>开发技术：</b><b>前端：</b>H5、css3、原生js、vue3.0、Element Plus &nbsp;&nbsp;&nbsp;&nbsp; <b>后端：</b>Golang
                    &nbsp;&nbsp;&nbsp;&nbsp;<b>数据库：</b>Mysql 8.0</p>
            </div>
        </div>
        <div>
            <div class="content_title">
                <h3>特别鸣谢</h3>
            </div>
            <div class="content_box">
                <p><b><a href="https://www.wangeditor.com/">wangEditor：</a></b>开源 Web 富文本编辑器，开箱即用，配置简单</p>
                <p><b><a href="https://element-plus.gitee.io/zh-CN/">Vue.js：</a></b>渐进式 JavaScript 框架</p>
                <p><b><a href="https://v3.cn.vuejs.org/">Element Plus：</a></b>基于 Vue 3，面向设计师和开发者的组件库</p>
                <p><b><a href="https://echarts.apache.org/zh/index.html">Apache ECharts：</a></b>一个基于 JavaScript 的开源可视化图表库</p>
                <p><b><a href="www.bpvank.com">bpvank：</a></b>感谢大佬的开源</p>
            </div>
        </div>
    </div>

    <div class="message_box">
        <div class="titleWrapper">
            <h3 class="title"><b>留言</b></h3>
        </div>

        <el-row :gutter="20">
            <el-col :span="24">
                <p>◎ 欢迎参与讨论，请在这里发表您的看法和观点。留言头像采用“腾讯QQ”的头像</p>
            </el-col>
            <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
                <el-input v-model="message.name" size="large" placeholder="昵称" />
            </el-col>
            <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
                <el-input v-model.number="message.qq" size="large" placeholder="QQ号" />
            </el-col>
            <el-col :span="24">
                <el-input v-model="message.content" :autosize="{ minRows: 4, maxRows: 5 }" type="textarea"
                    placeholder="留言内容" maxlength="50" />
            </el-col>
            <el-col :xs="24" :sm="18" :md="18" :lg="18" :xl="18">
                <el-input v-model="message.code" size="large" placeholder="验证码">
                    <template #append>
                        <VerificationCode ref="player" @getVerificationCode="getVerificationCode" title="看不清？点击换一张">
                        </VerificationCode>
                    </template>
                </el-input>
            </el-col>
            <el-col :span="24" :sm="6" :md="6" :lg="6" :xl="6">
                <el-button color="#2fa7b9" plain size="large" style="width:100%;text-align: right;" @click="submit">发布留言
                </el-button>
            </el-col>
        </el-row>
    </div>

    <div class="message_content" v-loading="loading" element-loading-text="玩命加载中...">
        <div class="titleWrapper">
            <h3 class="title"><b>最新留言</b></h3>
        </div>
        <div>
            <ul>
                <li class="item" v-for="item in messageInfo">
                    <div class="userInfo">
                        <img :src="item.messageQQ" alt="">
                        <font># {{item.messageName}}</font>
                        <l>{{item.messageDate}}</l>
                        <span @click="reply(item.messageId,item.replyContent)" class="operation">
                            <span v-if="userInfo.userType==0 && item.messageId !== replyId">
                                <el-icon>
                                    <ChatDotSquare />
                                </el-icon>
                                <span v-if="item.replyContent =='' || item.replyContent ==null">回复</span>
                                <span v-else>编辑回复</span>
                            </span>
                            <span v-if="userInfo.userType==0 && item.messageId === replyId">
                                <el-icon>
                                    <ChatDotSquare />
                                </el-icon>
                                取消回复
                            </span>
                        </span>
                        <el-popconfirm confirm-button-text="确定" cancel-button-text="取消" :icon="Delete"
                            icon-color="#F56C6C" confirm-button-type="danger" :title="'确定删除该评论吗？'"
                            @confirm="deleteMessage(item.messageId)">
                            <template #reference>
                                <span v-if="userInfo.userType==0" class="operation">
                                    <el-icon>
                                        <Delete />
                                    </el-icon>
                                    删除
                                </span>
                            </template>
                        </el-popconfirm>
                    </div>
                    <div class="commentGroup">
                        <p class="detailText">{{item.content}}</p>
                    </div>

                    <div v-show="item.messageId === replyId" class="comment_reply_box">
                        <el-row :gutter="20">
                            <el-col :span="20">
                                <el-input v-model="replyContent" maxlength="100" size="large"
                                    :placeholder="'回复@' + item.messageName+ ':'" show-word-limit />
                            </el-col>
                            <el-col :span="4">
                                <el-button plain color="#2fa7b9" size="large" style="width: 100%;"
                                    @click="submitReply(item.messageId)">提交
                                </el-button>
                            </el-col>
                        </el-row>
                    </div>

                    <div class="otherInfo" v-if="item.replyContent !== null && item.replyContent !== ''"><strong
                            style="color:#2FA7B9">TolZB</strong>：{{item.replyContent}}</div>
                </li>
            </ul>
        </div>
        <el-pagination class="showPaging" background layout="prev, pager, next" :total="total" :page-size="pageSize"
            :disabled="disabled" @current-change="changePage" v-model:currentPage="currentPage"
            @size-change="handleSizeChange" :pager-count="3" :hide-on-single-page="true" />
        <el-pagination class="hidePaging" background layout="prev, pager, next, jumper" :total="total"
            :page-size="pageSize" :disabled="disabled" @current-change="changePage" v-model:currentPage="currentPage"
            @size-change="handleSizeChange" :pager-count="3" :hide-on-single-page="true" />
    </div>
</template>
<script>
    // 验证码组件
    import VerificationCode from '../verification-code/verificationCode.vue'

    import {
        reactive,
        ref,
        toRefs,
        unref,
        onMounted,
    } from 'vue'
    import { FormatDate} from '@/utils/index'
    import axios from "axios";

    import {
        ElNotification,
        ElMessage
    } from 'element-plus'
    import {
        Delete
    } from '@element-plus/icons-vue'
    /**
     * @desc 加载获取数据
     */
    function loadData(state) {
        state.loading = true;
        const params = {
            'currentPage': state.currentPage,
            'pageSize': state.pageSize,
        }
        axios.get("/message/pageShow", {
            params
        }).then(res => {
            // 先清空数据
            state.messageInfo = [];
            res.data.data.list.forEach(element => {
                // 时间戳格式化
                element.messageDate = FormatDate(element.messageDate);
                // 头像采用qq的
                element.messageQQ = 'https://q2.qlogo.cn/headimg_dl?dst_uin=' + element.messageQQ +
                    '&spec=100'
                // 添加
                state.messageInfo.push(element)
            });
            //JSON.parse 将从后台得到的数据转换为标准JSON格式
            //前台展示的是需要数组，JSON.parse转换后的数据，element-plus可以解析
            // state.tableData = res.data.data.list;
            state.total = res.data.data.total;
            state.pageSize = res.data.data.pageSize;
            state.currentPage = res.data.data.currentPage;
            state.loading = false;
        })
    }

    export default {
        components: {
            VerificationCode
        },
        setup(props) {
            // 获取当前登录的用户信息
            var tokenStr = sessionStorage.getItem("userInfo");
            // 子组件的方法
            const player = ref(null);
            //挂载后加载数据
            onMounted(() => {
                loadData(state);
                if (tokenStr) {
                    state.userInfo = JSON.parse(tokenStr).data[0]
                    // 表单只读状态关闭
                    state.readonly = false
                }
            });
            const state = reactive({
                // 当前登录信息
                userInfo: {},
                // 留言内容
                message: {
                    name: '',
                    qq: '',
                    content: '',
                    code: ''
                },
                // 随机生成的验证码
                verificationCode: '',
                // 留言信息
                messageInfo: [],
                // 回复内容
                replyContent: '',
                // 回复状态
                replyId: false,
                total: 0, //总条数
                pageSize: 5, //每页显示行数
                currentPage: 1, //当前页码
                loading: false, // 加载
            })

            // 接收子组件（验证码）传的值
            const getVerificationCode = (codeList) => {
                state.verificationCode = codeList;
            };

            // 切换页面的执行事件，  val 当前页码
            const changePage = (val) => {
                state.currentPage = val;
                loadData(state);
            };

            // 发布评论
            const submit = () => {
                // 判断内容是否填写完整
                if (state.message.name == '' || state.message.qq == '' || state.message.content == '' || state
                    .message.code == '') {
                    ElMessage.error('将内容填写完整，表达你的态度再评论吧！')
                } else {
                    // 判断验证码是否输入正确
                    if (state.message.code.toLowerCase() !== state.verificationCode.toLowerCase()) {
                        ElMessage.error('验证码有误，请重新输入！')
                    } else {
                        const params = {
                            messageName: state.message.name, // 留言昵称
                            messageQQ: state.message.qq, // 留言QQ
                            content: state.message.content, // 留言内容
                        }

                        axios.post("/message/insert", params).then((res) => {
                            if (res.data.code == 0) {
                                // 清空内容
                                state.message = {
                                    name: '',
                                    qq: '',
                                    content: '',
                                    code: ''
                                }
                                // 刷新留言
                                loadData(state)
                                // 调用子组件方法（更新验证码）refreshCode
                                unref(player).refreshCode();
                                ElMessage({
                                    message: '留言成功，感谢支持.',
                                    type: 'success',
                                })
                            }
                        })
                    }
                }
            }

            // 回复留言
            const reply = (id, replyContent) => {
                // 关闭回复
                if (state.replyId === id) {
                    state.replyId = false;
                }
                // 当前留言回复没有关闭又点击其他回复时
                else if (state.replyId !== id && state.replyId !== false) {
                    state.replyId = id
                    state.replyContent = replyContent
                } else { // 为false时
                    state.replyId = id
                    state.replyContent = replyContent
                }
            }

            // 提交回复留言
            const submitReply = (messageId) => {
                if (state.replyContent !== "") {
                    axios.get("/message/updateMessageReply", {
                        params: {
                            messageId: messageId, // 回复留言id
                            replyContent: state.replyContent // 回复内容
                        }
                    }).then(res => { // 回复成功
                        if (res.data.code == 0) {
                            // 清除内容
                            state.replyContent = ""
                            state.replyId = false
                            ElMessage({
                                message: '留言回复成功.',
                                type: 'success',
                            })
                            // 刷新评论数据
                            loadData(state);
                        } else {
                            ElMessage.error('回复失败')
                        }

                    })
                } else {
                    ElMessage.error('回复内容不能为空哦~')
                }
            }

            // 删除留言
            const deleteMessage = (messageId) => {
                axios.delete("/message/delete", {
                    params: {
                        messageId: messageId
                    }
                }).then(res => { // 删除成功
                    if (res.data.code == 0) {
                        ElMessage({
                            message: '留言删除成功.',
                            type: 'success',
                        })
                        // 刷新评论数据
                        loadData(state);
                    } else {
                        ElMessage.error('删除失败')
                    }

                })
            }

            return {
                ...toRefs(state),
                getVerificationCode,
                submit,
                reply,
                deleteMessage,
                Delete,
                player,
                changePage,
                submitReply
            }
        }

    }
</script>

<style scoped>
    .main_content,
    .message_box,
    .message_content {
        width: 100%;
        height: auto;
        overflow: hidden;
        background-color: white;
        padding: 25px;
        line-height: 30px;
        color: #474749;
        border-radius: 10px;
        margin-bottom: 20px;
        box-sizing: border-box;
    }

    .message_box .titleWrapper,
    .message_content .titleWrapper {
        margin-bottom: 15px;
        margin-top: -10px;
        border-bottom: 1px solid #eee;
        overflow: hidden;
        display: flex;
        justify-content: space-between;
    }

    .message_box b,
    .message_content b {
        display: inline-block;
        border-bottom: 4px solid #2FA7B9;
    }

    .content_box {
        width: 100%;
    }

    .content_box>h4 {
        margin-left: 10px;
    }

    .content_box>p {
        text-indent: 2em;
        line-height: 35px;
    }

    .content_title>h3 {
        font-size: 20px;
        font-weight: 700;
        padding-left: 10px;
        margin: 15px 0;
        border-left: 4px solid #2FA7B9;
    }

    .message_box .el-col {
        margin-bottom: 20px;
    }

    :deep(.el-input-group__append, .el-input-group__prepend) {
        background: white;
    }

    .message_box p {
        font-size: 14px;
        color: #666;
    }

    .message_content .userInfo {
        display: inline-flex;
        justify-content: center;
        align-items: center;
    }

    .message_content .userInfo img {
        width: 40px;
        border-radius: 50px;
        margin-right: 10px;
    }

    .message_content .userInfo>font {
        font-size: 15px;
        color: #333;
    }

    .message_content .userInfo>l {
        font-size: 12px;
        color: #999;
        margin: 0 10px;
    }

    .message_content .commentGroup {
        padding: 15px;
        border-radius: 4px;
        background: #f6f6f6;
        overflow: hidden;
    }

    .message_content .detailText {
        line-height: 20px;
        font-size: 13px;
        color: #666;
    }

    .message_content .otherInfo {
        font-size: 13px;
        color: #666;
        line-height: 18px;
        margin-top: 10px;
        word-break: break-all;
    }

    .message_content .item {
        margin-bottom: 20px;
        padding-bottom: 20px;
        border-bottom: 1px dashed #eee;
    }

    .comment_reply_box {
        width: 100%;
        margin: 15px 0;
    }

    .message_content .operation {
        cursor: pointer;
        color: #777777;
        font-size: 12px;
        margin-right: 10px;
        display: inline-flex;
        justify-content: center;
        align-items: center;
    }

    .message_content .operation span{
        display: inline-flex;
        justify-content: center;
        align-items: center;
    }

    .message_content .operation .el-icon {
        margin-right: 5px;
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