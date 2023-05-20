<template>
    <section class="w3l-error-61-404">
        <div class="error-61-mian">
            <div class="wrapper">
                <div class="errors-16-top">
                    <div class="errors-16-info">
                        <h3> 404 </h3>
                        <p>sorry! 页面找不到了~</p>
                        <span style="color: #a7621c;line-height: 30px;cursor: pointer;" @click="goBack">返回上一页</span>
                        <form action="https://www.x6d.com/" method="post" class="d-flex error-page-form">
                            <font style="color:#ffffff;font-size:21px;font-weight:bold;line-height: 35px;">
                                你穿越到了一个未知领域! ∑(っ°Д°;)っ</font>
                            <el-button class="hidden-sm-and-down" color="#a7621c" round plain style="margin-left: 20px;"  @click="goHome">TolZB</el-button>
                        </form>

                    </div>

                </div>
                <div class="copyright-footer">
                    <div class="copy-right">
                        <p>Copyright &copy;2020 <a href="/" title="404">TolZB.</a>
                            All rights reserved.</p>
                    </div>
                </div>
            </div>
        </div>
    </section>
</template>

<script>
    import {
        useRouter
    } from 'vue-router'
    export default {
        setup() {
            const router = useRouter();
            const goHome = () => {
                router.push('/')
                sessionStorage.setItem('activeIndex1', '/')
            }
            const goBack = () => {
                router.go(-1)
            }
            // 打印立体TolZB 哈哈
            console.log(`
    ÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜ¼¼¼¼¼ÜÜÜÜÜÜÜÜÜÜÜÜ¼Ü¼¼¼ÜÜ¼¼¼
    ÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜ¼Ü¼¼ÜÜÜÜÜÜÜÜÜÜÜÜÜÜ¼Ü¼Ü¼ÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜ¼Ü¼ÜÜ¼¼¼Ü¼¼Ü¼¼¼¼¼¼
    ÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜ¼¼ÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜÜ¼¼ÜÜ¼¼ÜÜÜ¼ÜÜÜ¼¼¼¼¼ÜÜÜ¼¼Ü¼¼¼¼ÜÜ¼¼¼¼¼
    ÜÜÜÜÜÜÜÜÜÜ¼Ü¼ÜÜÜÜ¼¼¼¼¼¼Ü¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼¼ÜÜ¼¼¼¼¼¼¼ÜÜÜ¼¼
    ÜÜÜ¼ÜÜÜÜÜÜÜÜÜÜÜ¼¼¼¼¼¼ÜÜ[Ü[[[[[[[[[ÜÜÜÜ¼ÜÜÜÜ¼¼[¼[¼ÜÜ¼¼¼¼¼¼¼¼¼ÜÜÜ¼¼¼¼¼ÜÜ¼¼ÜÜ¼¼¼¼Ü¼
    ÜÜÜÜÜÜÜÜÜÜ¼Ü¼¼¼Ü¼ÜÜÜÜ[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[¼[Ü¼¼¼¼¼¼ÜÜÜ¼¼¼¼¼¼Ü¼ÜÜ¼ÜÜ¼
    ÜÜÜÜÜÜÜÜ¼ÜÜÜÜÜÜÜÜÜ[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[Ü[ÜÜ¼ÜÜÜÜ¼¼¼ÜÜÜÜÜ¼
    ÜÜÜÜÜÜÜ¼¼¼¼¼¼¼¼¼Ü[[[[[[[[[[¼[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[¼¼¼¼¼¼¼¼¼¼¼ÜÜÜÜ¼
    ÜÜÜÜ¼ÜÜ¼¼¼¼¼¼¼¼Ü[[[ÜÜÜ[¼¼¼¼Ü¼Ü[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[¼¼¼¼¼¼¼¼¼Ü¼ÜÜÜÜ
    ÜÜÜÜÜÜÜÜÜÜÜÜÜ¼¼Ü[[[¼[¼¼ü███████████▄[[░░░░░░░░░░░░░[[░░[[[[[[[[[[[[¼¼¼¼ÜÜ¼¼Ü¼ÜÜ¼
    ÜÜ¼Ü¼¼¼¼Ü¼¼¼[Ü[[[[¼Ü¼▓████████████████▄[░░░░░░░░░░░░░░░[[[[[[[[[[¼[¼¼¼¼¼¼¼¼¼¼ÜÜ¼
    ÜÜ¼Ü¼¼¼¼¼Ü[[[[[[[¼¼█████████████████████▄░░░░░░░░░░░░░░░░[[[[[[[[[[¼¼¼¼¼¼¼¼¼¼¼¼Ü
    ÜÜ¼Ü¼¼¼¼¼¼[[[Ü[[[▓████████████████████████[░░░░░░░░░░░░░░[[[[[[[[[[¼¼Ü¼¼¼¼¼¼¼¼¼¼
    █Ü¼Ü¼¼¼ÜÜ[[[[Ü[¼¼██████████████████████████░░░░░││░░░░░░[░[[[[[[[¼¼¼¼¼¼¼¼¼¼¼ÜÜÜ¼
    ██¼¼ÜÜÜ¼¼[▄▄█▄¼¼▓███████████████████████████░░░░░░│░░░░░░[[[[[[[¼[¼¼¼¼Ü¼¼¼¼¼ÜÜÜ¼
    ██████ÿÜ[▓██████████████████████████████████[░░░░░░░░░░░[[[[[[[[[¼¼¼¼¼¼¼¼¼¼¼Ü¼¼Ü
    █████████████████████████████████████████████[░░░░░░░[░[[[[[[[¼¼[¼¼¼¼¼¼¼¼¼ÜÜÜÜÜ¼
    █████████████████████████████████████████▒▓██Ü[░░░░[[[░[[[[[[[[¼¼¼Ü¼¼¼¼¼ÜÜÜÜ¼Ü¼Ü
    █████▓███████████████████████████████████████Ü¼[░[[[[[[[[[[Ü[¼¼¼¼¼¼¼Ü¼ÜÜÜÜÜÜÜ¼¼¼
    █████████████████████████████████▓███████▓████▄█[[[[[[[¼[[¼¼¼¼¼¼¼¼¼¼¼ÜÜ¼¼¼¼ÜÜ¼¼Ü
    ███████████████████████████████▓███▓▓▓▓▓▒▓▓███████[¼¼[▄[▄█▄ç¼¼¼¼Ü¼¼Ü¼ÜÜÜÜ¼¼ÜÜÜÜÜ
    ███████████████████████████████▒▓▓██████▓▒▓█████████████████████Ü¼Ü¼¼ÜÜÜÜÜÜÜ¼ÜÜÜ
    ███████████████████████████████████▓█████▓██████████████████████▒ù%ÜÜÜÜÜÜÜÜÜÜÜÜÜ
    ███████████████████████████████████▓█████▓██████████████████▓██████ÜÜ¼ÜÜÜ¼¼¼ÜÜÜ¼
    ████████████████████████████████▓▓▓▓▓██▓▓██████████████████▓▓████████¼ÉÜÜÜÜÉÉÉÜÜ
    ██████▓█████████████████████████████▓▓▓██████████████████▓███▓█▓████▓▓██████████
    █████▓▓██████████████████████████▓▓█▓████████▓████▓███████████████▒▒▓▓██████████
    ████▓▓████████████████████████████████▓▓ÜÜⁿ▓█▓▓ÜÜÜÜÜÜÜ▓██████▓█████▓▓▓██████████
    ████████▓▓████████████████████████████▓ÜÉ¼Ü¼Ü[ÜÜÜⁿ%▒▒É½▒▓▓▓██▓██▓███▓▓▓▓███▓▓█▓█
    █████▓▓▓▓▓███████████████████████████▓ÜÜ¼Ü¼ÜÉ¼%Ü¼Ü[Üⁿ▓▓▓▓▓@▓▓▓▓▓█████▓██████████
    ████▓▓▓▓▓▓▓██████████████████████████Ü¼¼Ü¼¼ÉÜ¼Ü¼Ü¼¼¼ÜÜÜ▓▓▒ÜÜ▓▓▓██████▓▓█████████
    █▓¼▓▓▓▓▓▓████████████████████▓▓▓▓███ÜÜ¼▒¼ÜÜ▓▒Ü½Ü¼¼¼Ü¼ÜÜ[▓███▓▒▒▓▓█████▓▓████████
    ¼ÉÉÖÖ▓▓▓█▓███████████████████▓▓▓███Ü¼É▒ÜÜÉÜ%▓ÉÜ¼½É¼Ü¼¼Ü¼¼[▓████▓▓▓█████▓████████
    ÉÉÉÉÉ▓▓▓▓▓██████████████████▓█████▒½ÉÉ▓¼ÉÉÉÖ▓É¼ÜÉ▒½▒ÉÉÜÜÜ¼¼Å███████████▓████████
    É¼ÜÜÜÉÖ▓▓▓████████████████████▓▓▓WÜ▒▒▓%▒▒▓▓¼▓▒Ü▒▓▒▓▒▒%ÉÜÜÜÜ¼Å▓▓█████████▓███████
    ▓████▒ÜÜ▓▓▓████████████████████▒▒▒▓▒▓▒%▓▓▓▓▓█▓█▓█▒▒▒%ÉÜÜ▄MMMÜ▓██████████████████
    █▓▓▓▓▓▓█▒▓▓████████████████████▒▒▓▓▒▓▓▒███████████▒▓██▓▓▓▒▒▒▓▓█████▓▓█▓█████████
    ▓▓▓▓▓▓▓███████████████████████▒▓▓███▓█▒▓████████████████▓▓▓▓▓▓▓██████▓▓▓████▓▓▓▒
    ▓▓▓▓▓███████████████████████████▓▒▓████▓█████████████████▓▓▓▓▓▓▓███████▓▒▒▓▒▒▓▓▓
    ▓▓▓▓▓██████████████████████████▒██▓▒▓██▓███▓██▓████████████▓▓▓▓▓▓███████▓██▓████
    ▓▓▓████████████████████████████▒▓██▒▓▓▓▓▓▓▓▓█▓██████████████▓▓▓▓▓▓▓██████▓█▓████
    ▓▓████████████████████████████▓▓██▓▓▓▓▓████▓▒▓█████████████████▓▓▓▓▓███████▒█▓▓█
    █████████████████████████████▓▓██▓▓████████▒▒▓██████████████████████▓▓███████▓██
    ████████████████████████████████▓▓█████████▒▒▓████████████████████████████████▓█
    ███████████████████████████▓████████████████▓▓███████████████████████████▓██████
    ███████████████████████████▓█████████████▓█▓▓▓█████▓███████████████████████▓████
    ████████████████████████████████████████████████████████████████████████████████
    ██████████████████████████████████████████████████████████████████████████████▓█`)
            return {
                goHome,
                goBack
            }
        }
    }
</script>

<style>
    html {
        scroll-behavior: smooth;
    }

    body,
    html {
        margin: 0;
        padding: 0;
        font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto,
            Oxygen-Sans, Ubuntu, Cantarell, "Helvetica Neue", sans-serif;
    }

    * {
        box-sizing: border-box;
    }

    .d-grid {
        display: grid;
    }

    .d-flex {
        display: flex;
        display: -webkit-flex;
    }

    .text-center {
        text-align: center;
    }

    .text-left {
        text-align: left;
    }

    .text-right {
        text-align: right;
    }

    button,
    input,
    select {
        -webkit-appearance: none;
        outline: none;
    }

    button,
    .btn,
    select {
        cursor: pointer;
    }

    a {
        text-decoration: none;
    }

    iframe {
        border: none;
    }

    ul {
        margin: 0;
        padding: 0
    }

    h1,
    h2,
    h3,
    h4,
    h5,
    h6,
    p {
        margin: 0;
        padding: 0
    }

    p {
        color: #485460;
    }

    .p-relative {
        position: relative;
    }

    .tlinks {
        text-indent: -9999px;
        height: 0;
        line-height: 0;
        font-size: 0;
        overflow: hidden;
    }

    .p-absolute {
        position: absolute;
    }

    .p-fixed {
        position: fixed;
    }

    .p-sticky {
        position: sticky;
    }

    .btn,
    button,
    .actionbg {
        border-radius: 2px;
        -webkit-border-radius: 2px;
        -moz-border-radius: 2px;
        -o-border-radius: 2px;
        -ms-border-radius: 2px;
    }

    .btn:hover,
    button:hover {
        transition: 0.5s ease;
        -webkit-transition: 0.5s ease;
        -o-transition: 0.5s ease;
        -ms-transition: 0.5s ease;
        -moz-transition: 0.5s ease;
    }

    /*--/wrapper--*/
    .wrapper {
        width: 100%;
        padding-right: 15px;
        padding-left: 15px;
        margin-right: auto;
        margin-left: auto;
    }

    @media (min-width: 576px) {
        .wrapper {
            max-width: 540px;
        }
    }

    @media (min-width: 768px) {
        .wrapper {
            max-width: 720px;
        }
    }

    @media (min-width: 992px) {
        .wrapper {
            max-width: 960px;
        }
    }

    @media (min-width: 1200px) {
        .wrapper {
            max-width: 1140px;
        }
    }

    .wrapper-full {
        width: 100%;
        padding-right: 15px;
        padding-left: 15px;
        margin-right: auto;
        margin-left: auto;
    }

    /*--//wrapper--*/
    .w3l-error-61-404 .error-61-mian {
        background: url(../../../assets/banner/banner01.png) no-repeat bottom;
        background-size: cover;
        -webkit-background-size: cover;
        -moz-background-size: cover;
        -o-background-size: cover;
        -moz-background-size: cover;
        height: 100vh;
        z-index: 0;
        position: relative;
        display: grid;
        align-items: center;
    }

    .w3l-error-61-404 .error-61-mian:before {
        content: "";
        background: rgba(66, 21, 2, 0.45);
        position: absolute;
        top: 0;
        min-height: 100%;
        left: 0;
        right: 0;
        z-index: -1;
    }

    .w3l-error-61-404 .errors-16-top {
        max-width: 600px;
    }

    .w3l-error-61-404 .error-61-mian h3 {
        font-size: 223px;
        line-height: 200px;
        color: #fff;
        font-weight: bold;
        margin-bottom: 20px;
        text-shadow: 0 1px 2px rgba(0, 0, 0, .6);
    }

    .w3l-error-61-404 .error-61-mian p {
        font-size: 20px;
        line-height: 30px;
        color: #fff;
        letter-spacing: 1px;
    }

    .w3l-error-61-404 .error-page-form {
        border: 2px solid #fff;
        border-radius: 25px;
        padding: 5px 20px;
        margin: 20px 0;
    }

    .w3l-error-61-404 .error-page-form input {
        color: #fff;
        font-size: 16px;
        border: none;
        outline: none;
        display: block;
        padding: 0.4em 1em;
        background: rgba(114, 51, 23, 0.75);
        letter-spacing: 1px;
        width: 100%;
        border-radius: 25px;
        margin-right: 1%;
    }

    ::-webkit-input-placeholder {
        /* Chrome/Opera/Safari */
        color: #fff;
        background: transparent;
    }

    ::-moz-placeholder {
        /* Firefox 19+ */
        color: #fff;
        background: transparent;
    }

    :-ms-input-placeholder {
        /* IE 10+ */
        color: #fff;
        background: transparent;
    }

    :-moz-placeholder {
        /* Firefox 18- */
        color: #fff;
        background: transparent;
    }

    .w3l-error-61-404 .copyright-footer {

        margin: 100px auto 0;

    }

    .w3l-error-61-404 .copy-right p {
        font-size: 14px;
        line-height: 29px;
        color: #f9f9f9;
    }

    .w3l-error-61-404 .copy-right a {
        color: #f9f9f9;
    }

    @media (max-width: 1280px) {
        .w3l-error-61-404 .error-61-mian h3 {
            font-size: 180px;
            line-height: 170px;
        }
    }

    @media (max-width: 1280px) {
        .w3l-error-61-404 .error-61-mian h3 {
            font-size: 125px;
            line-height: 120px;
            color: #fff;
            font-weight: 700;
            margin-bottom: 20px;
        }
    }

    @media (max-width: 990px) {
        .w3l-error-61-404 .error-61-mian h3 {
            font-size: 125px;
            line-height: 110px;
        }
    }

    @media (max-width: 667px) {
        .w3l-error-61-404 .error-61-mian h3 {
            font-size: 115px;
            line-height: 110px;
        }
    }

    @media (max-width:600px) {
        .w3l-error-61-404 .error-61-mian {
            background: url(../../../assets/banner/banner01.png) no-repeat right;
            background-size: cover;
            -webkit-background-size: cover;
            -moz-background-size: cover;
            -o-background-size: cover;
            -moz-background-size: cover;
        }
    }

    @media (max-width: 440px) {
        .w3l-error-61-404 .error-61-mian h3 {
            font-size: 105px;
            line-height: 100px;
        }
    }

    @media (max-width: 384px) {
        .w3l-error-61-404 .error-61-mian h3 {
            font-size: 95px;
            line-height: 90px;
        }
    }
</style>