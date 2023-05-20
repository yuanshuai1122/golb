/*
 Navicat Premium Data Transfer

 Source Server         : 82.157.47.194
 Source Server Type    : MySQL
 Source Server Version : 80016
 Source Host           : 82.157.47.194:3306
 Source Schema         : bpvank

 Target Server Type    : MySQL
 Target Server Version : 80016
 File Encoding         : 65001

 Date: 31/12/2022 14:33:41
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for articleinfo
-- ----------------------------
DROP TABLE IF EXISTS `articleinfo`;
CREATE TABLE `articleinfo`  (
  `articleId` int(11) NOT NULL AUTO_INCREMENT COMMENT '文章编号',
  `userId` int(11) NOT NULL COMMENT '编写用户编号',
  `userName` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '编写用户名称',
  `articleTitle` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '文章标题',
  `articleClassifyId` int(11) NOT NULL COMMENT '文章分类id',
  `articleClassifyName` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '文章分类名称',
  `articleDase` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '文章描述',
  `articleImgLitimg` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '文章缩略图',
  `articleContent` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '文章内容',
  `publishTime` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '发表时间',
  `articleState` int(11) NOT NULL COMMENT '文章状态',
  `articlePass` int(11) NULL DEFAULT 1 COMMENT '后台审核是否通过',
  `commentState` int(11) NOT NULL COMMENT '评论状态',
  `click` int(11) NOT NULL DEFAULT 0 COMMENT '阅读次数',
  `review` int(11) NOT NULL DEFAULT 0 COMMENT '评论次数',
  PRIMARY KEY (`articleId`) USING BTREE,
  INDEX `fk_articleInfo_userId`(`userId`) USING BTREE,
  INDEX `fk_articleInfo_articleClassifyId`(`articleClassifyId`) USING BTREE,
  CONSTRAINT `fk_articleInfo_userId` FOREIGN KEY (`userId`) REFERENCES `userinfo` (`userId`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 114 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '文章管理表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of articleinfo
-- ----------------------------
INSERT INTO `articleinfo` VALUES (107, 1, 'admin', '常用命令合集', 1, '技术', '平常使用次数多的命令', 'static/thumbnail/banner06.png', '<p>mysql&gt; source /home/bpvank.sql; </p><p><br></p><p>/usr/local/nginx/html/</p><p><br></p><p>/usr/local/nginx/sbin</p><p><br></p><p>/usr/local/nginx/conf/</p><p><br></p><p>/usr/local/go_server/</p><p><br></p><p>netstat -tunlp | grep 520</p><p> </p><p>kill -9 </p><p><br></p><p>nohup ./serverBlog &</p><p><br></p><p>$env:GOOS=\"linux\"</p><p><br></p><p>go build -o serverBlog .\\main.go</p><p><br></p><p>ALTER USER \'root\'@\'localhost\' IDENTIFIED BY \'\';</p><p><br></p><p><br></p>', '2022-12-22 19:35:46', 1, 2, 1, 7, 0);
INSERT INTO `articleinfo` VALUES (108, 1, 'admin', 'Centos7.6安装mysql8.0', 1, '技术', 'Centos7.6安装mysql8.0', 'static/thumbnail/4f8c4ce8a0ee4e90bc85e0d363ebd658.png', '<p><span style=\"color: rgb(77, 77, 77); background-color: rgb(255, 255, 255); font-size: 16px;\">1.在/usr/local/目录下创建mysql文件夹</span></p><pre><code >cd  /usr/local\r\nmkdir mysql\ncd /mysql</code></pre><p><span style=\"color: rgb(77, 77, 77); background-color: rgb(255, 255, 255); font-size: 16px;\">2.下载</span></p><pre><code >wget https://dev.mysql.com/get/Downloads/MySQL-8.0/mysql-8.0.16-2.el7.x86_64.rpm-bundle.tar</code></pre><p><span style=\"color: rgb(77, 77, 77); background-color: rgb(255, 255, 255); font-size: 16px;\">3.解压</span></p><pre><code >tar -xvf mysql-8.0.16-2.el7.x86_64.rpm-bundle.tar</code></pre><p><span style=\"color: rgb(77, 77, 77); background-color: rgb(255, 255, 255); font-size: 16px;\">4.安装MySQL，依次执行</span></p><pre><code >rpm -ivh mysql-community-common-8.0.16-2.el7.x86_64.rpm --nodeps --force\n\nrpm -ivh mysql-community-libs-8.0.16-2.el7.x86_64.rpm --nodeps --force\n\nrpm -ivh mysql-community-client-8.0.16-2.el7.x86_64.rpm --nodeps --force\n\nrpm -ivh mysql-community-server-8.0.16-2.el7.x86_64.rpm --nodeps --force\r</code></pre><p><span style=\"color: rgb(77, 77, 77); background-color: rgb(255, 255, 255); font-size: 16px;\"> &nbsp;查看</span></p><pre><code >rpm -qa | grep mysql</code></pre><p><span style=\"color: rgb(77, 77, 77); background-color: rgb(255, 255, 255); font-size: 16px;\">5.初始化和相关配置</span></p><pre><code >mysqld --initialize;\r\n\r\nchown mysql:mysql /var/lib/mysql -R;\r\n\r\nsystemctl start mysqld.service;\r\n\r\nsystemctl enable mysqld;</code></pre><p><span style=\"color: rgb(77, 77, 77); background-color: rgb(255, 255, 255); font-size: 16px;\">6.</span><a href=\"https://so.csdn.net/so/search?q=%E6%9F%A5%E7%9C%8B%E6%95%B0%E6%8D%AE%E5%BA%93&amp;spm=1001.2101.3001.7020\" target=\"_blank\" style=\"text-align: start;\">查看数据库</a><span style=\"color: rgb(77, 77, 77); background-color: rgb(255, 255, 255); font-size: 16px;\">初始化密码</span></p><pre><code >cat /var/log/mysqld.log | grep password</code></pre><p><span style=\"color: rgb(77, 77, 77); background-color: rgb(255, 255, 255); font-size: 16px;\">7.登录，复制初始化密码登录</span></p><pre><code >mysql -uroot -p</code></pre><p><span style=\"color: rgb(77, 77, 77); background-color: rgb(255, 255, 255); font-size: 16px;\">8.修改密码</span></p><pre><code >ALTER USER \'root\'@\'localhost\' IDENTIFIED WITH mysql_native_password BY \'xxxxxxx\';</code></pre><p><br></p>', '2022-12-22 19:45:37', 1, 2, 1, 3, 0);
INSERT INTO `articleinfo` VALUES (109, 1, 'admin', 'linux下安装nginx', 1, '技术', 'linux下安装nginx', 'static/thumbnail/4f8c4ce8a0ee4e90bc85e0d363ebd658.png', '<h1>下载安装包</h1><pre><code >cd /home   # 进入指定目录，目录请自行选择\r\nwget https://nginx.org/download/nginx-1.21.2.tar.gz   # 下载安装包，版本请自行选择</code></pre><h1 style=\"text-align: start;\">安装步骤</h1><ol style=\"text-align: start;\"><li>cd 至 nginx 安装包所在目录，进行解压；</li></ol><pre><code >tar -zxvf nginx-1.21.2.tar.gz</code></pre><p> &nbsp; &nbsp;2.<span style=\"color: rgba(0, 0, 0, 0.75); background-color: rgb(255, 255, 255); font-size: 16px;\">cd 至解压后的目录，运行./configure 进行初始化配置</span></p><pre><code >cd ./nginx-1.21.2\n./configure</code></pre><p>3.<span style=\"color: rgba(0, 0, 0, 0.75); background-color: rgb(255, 255, 255); font-size: 16px;\">执行编译操作</span></p><pre><code >make   </code></pre><p>4.<span style=\"color: rgba(0, 0, 0, 0.75); background-color: rgb(255, 255, 255); font-size: 16px;\">执行安装操作</span></p><pre><code >make install</code></pre><pre><code > whereis nginx   # 查找安装路径\n cd /usr/local/nginx   # 进入安装路径\n ./sbin/nginx   # 启动\n ./sbin/nginx -s reload   # 重载\n ./sbin/nginx -s stop   # 关闭\n ./sbin/nginx -s quit   # 优雅关闭（当请求被处理完成之后才关闭）\r\n</code></pre><p><br></p>', '2022-12-22 19:56:17', 1, 2, 1, 1, 0);
INSERT INTO `articleinfo` VALUES (110, 1, 'admin', 'Vue3', 1, '技术', '公司一直在用react，vue2忘记了，直接学vue3吧', 'http://82.157.47.194:520/static/thumbnail/4f8c4ce8a0ee4e90bc85e0d363ebd658.png', '<h1>1.项目构建</h1><pre><code >//webpackage\nnpm init vue@latest\n//vite\nnpm init vite@latest</code></pre><h1>2.VsCode Vue3开发插件</h1><p>--volar</p><p>--TypeScript Vue Plugin (Volar)</p><h1>3.指令</h1><p>v- 开头都是vue 的指令</p><p>v-text 用来显示文本</p><p>v-html 用来展示富文本</p><p>v-if 用来控制元素的显示隐藏（切换真假DOM）</p><p>v-else-if 表示 v-if 的“else if 块”。可以链式调用</p><p>v-else v-if条件收尾语句</p><p>v-show 用来控制元素的显示隐藏（display none block Css切换）</p><p>v-on 简写@ 用来给元素添加事件</p><p>v-bind 简写: &nbsp;用来绑定元素的属性Attr</p><p>v-model 双向绑定</p><p>v-for 用来遍历元素</p><p>v-on修饰符 冒泡案例 @click.stop</p><h1>4.ref</h1><p><span style=\"color: rgb(77, 77, 77); background-color: rgb(255, 255, 255); font-size: 16px;\">--ref:接受一个内部值并返回一个</span><a href=\"https://so.csdn.net/so/search?q=%E5%93%8D%E5%BA%94%E5%BC%8F&amp;spm=1001.2101.3001.7020\" target=\"_blank\" style=\"text-align: start;\">响应式</a><span style=\"color: rgb(77, 77, 77); background-color: rgb(255, 255, 255); font-size: 16px;\">且可变的 ref 对象。ref 对象仅有一个 </span>.value<span style=\"color: rgb(77, 77, 77); background-color: rgb(255, 255, 255); font-size: 16px;\"> property，指向该内部值。</span></p><p><span style=\"color: rgb(77, 77, 77); background-color: rgb(255, 255, 255); font-size: 16px;\">--isRef:判断是不是一个ref对象</span></p><p><span style=\"color: rgb(77, 77, 77); background-color: rgb(255, 255, 255); font-size: 16px;\">--shallowRef：创建一个跟踪自身.value变化的ref，但不会使其值也变成响应式的</span></p><p><span style=\"color: rgb(77, 77, 77); background-color: rgb(255, 255, 255); font-size: 16px;\">--triggerRef：强制更新页面DOM</span></p><p><span style=\"color: rgb(77, 77, 77); background-color: rgb(255, 255, 255); font-size: 16px;\">--customRef：工厂函数要求我们返回一个对象并且实现get和set适合去做防抖之类的</span></p><p><span style=\"color: rgb(77, 77, 77); background-color: rgb(255, 255, 255); font-size: 16px;\">--获取dom元素</span></p><pre><code >&lt;div ref=\"dom\"&gt;&lt;/div&gt;\n\nconst dom = ref();\n\nconsole.log(dom.value)</code></pre><h1>5.reactive</h1><p>--<span style=\"color: rgb(77, 77, 77); background-color: rgb(255, 255, 255); font-size: 16px;\">用来绑定复杂的数据类型 例如 对象 数组</span></p><p><span style=\"color: rgb(77, 77, 77); background-color: rgb(255, 255, 255); font-size: 16px;\">--</span>readonly:<span style=\"color: rgb(77, 77, 77); background-color: rgb(255, 255, 255); font-size: 16px;\">拷贝一份proxy对象将其设置为只读</span></p><p><span style=\"color: rgb(77, 77, 77); background-color: rgb(255, 255, 255); font-size: 16px;\">--</span>shallowReactive: &nbsp;<span style=\"color: rgb(77, 77, 77); background-color: rgb(255, 255, 255); font-size: 16px;\">只能对浅层的数据 如果是深层的数据只会改变值 不会改变视图</span></p><p><br></p><p><br></p>', '2022-12-25 13:08:59', 1, 2, 1, 7, 0);
INSERT INTO `articleinfo` VALUES (111, 1, 'admin', 'Nginx配置信息', 1, '技术', 'Nginx配置信息', 'static/thumbnail/banner01.png', '<pre><code >\r\n#user  nobody;\r\nworker_processes  1;\r\n\r\n\r\nevents {\r\n    worker_connections  1024;\r\n}\r\n\r\n\r\nhttp {\r\n    \r\n    include       mime.types;\r\n    default_type  application/octet-stream;\r\n    sendfile        on;\r\n    keepalive_timeout  65;\r\n    \r\n    server {\r\n        \r\n        listen       80;\r\n        server_name  localhost;\r\n        \r\n        location / {\r\n            root   html/blog;\r\n            index  index.html index.htm;\r\n            try_files $uri $uri/ /index.html;\r\n        }\r\n        # img\r\n        location /static/ {\r\n            root  /usr/local/go_server/;\r\n            autoindex on;\r\n        }  \r\n        error_page   500 502 503 504  /50x.html;\r\n        location = /50x.html {\r\n            root   html;\r\n        }\r\n    }\r\n   \r\n}\r\n</code></pre><p><br></p>', '2022-12-26 21:15:36', 1, 2, 1, 4, 0);
INSERT INTO `articleinfo` VALUES (112, 1, 'admin', '博客换库方法', 1, '技术', '博客换库方法', 'http://82.157.47.194:520/static/thumbnail/4f8c4ce8a0ee4e90bc85e0d363ebd658.png', '<h1>一.web</h1><p> &nbsp;1.打包Vue项目</p><pre><code >npm run builds</code></pre><p> &nbsp; &nbsp;产生dist文件夹</p><p> &nbsp; 2.把dist文件夹内容拷贝到服务器 /usr/local/nginx/html/blog/路径下</p><p><img src=\"http://82.157.47.194:520/static/articleImg/image.png\" alt=\"\" data-href=\"\" style=\"width: 100%;\"/></p><hr/><h1>二.go服务</h1><p> &nbsp;1.在go文件夹下shift+右键打PowerShell,依次执行</p><pre><code >$env:GOOS=\"linux\"\n\ngo build -o serverBlog .\\main.go</code></pre><p> 2.删除/usr/local/go_server/ 目录下原来的serverBlog文件,拷贝进我们新的文件</p><p> 3.执行netstat -tunlp | grep 520 找到原来占用520端口的进程</p><p> 4.kill -9 原来的进程id</p><p><img src=\"http://82.157.47.194:520/static/articleImg/zzzzzz.jpg\" alt=\"\" data-href=\"\" style=\"width: 100%;\"/></p><p> &nbsp;5.执行nohup ./serverBlog & &nbsp;启动新的go服务</p>', '2022-12-26 22:24:11', 1, 2, 1, 8, 0);
INSERT INTO `articleinfo` VALUES (113, 1, 'admin', '我什么也没忘，但有些事情只适合收藏', 5, '情感', '如何忘掉前任：\n史铁生说过，\n“我什么也没忘，\n但有些事只适合收藏”。', 'http://82.157.47.194:520/static/thumbnail/4d086e061d950a7bc8acf03edd4e97def0d3c9d3.png', '<ul><li>如果你从来没有出现，我会不会过的快乐一些</li></ul><p> &nbsp; &nbsp; &nbsp;我的微信昵称是 1014 ，为什么不是1024(程序员节)呢？ 因为我在1014结束了一段恋爱。</p><p> &nbsp; &nbsp; &nbsp;</p><p> &nbsp; &nbsp; &nbsp;在这段断断续续的感情中我学会了怎么逼对方摊牌，冷暴力，发消息不回，晾你几天，你就受不了了，哈哈哈（两次被冷暴力的经验）。</p><p> &nbsp; &nbsp; &nbsp; </p>', '2022-12-27 11:12:04', 1, 2, 1, 4, 0);
INSERT INTO `articleinfo` VALUES (114, 1, 'admin', 'Git上传命令', 1, '技术', 'Git上传命令', 'static/thumbnail/git.png', '<pre><code >git init 初始化，创建本地仓库\ngit add . 添加到本地仓库\ngit commit -m \"注释\" 添加注释\ngit remote add origin 仓库地址 连接远程仓库\ngit pull --rebase origin master 同步仓库内容\ngit push -u origin master 上传到远程仓库\r</code></pre><p><br></p>', '2022-12-31 14:29:46', 1, 2, 1, 1, 0);

-- ----------------------------
-- Table structure for backstagemenuinfo
-- ----------------------------
DROP TABLE IF EXISTS `backstagemenuinfo`;
CREATE TABLE `backstagemenuinfo`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '菜单id',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '菜单名称',
  `chineseName` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '菜单中文名称',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '菜单标题',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '菜单路径',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '菜单图标',
  `parentMenuId` int(11) NULL DEFAULT NULL COMMENT '父菜单id',
  `menuStatus` tinyint(4) NULL DEFAULT NULL COMMENT '菜单是否可见',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '后台菜单表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of backstagemenuinfo
-- ----------------------------
INSERT INTO `backstagemenuinfo` VALUES (1, 'backstageHome', '后台首页', 'Bpvank - 后台首页', '/backstageHome', 'House', 0, 1);
INSERT INTO `backstagemenuinfo` VALUES (2, 'articleOperation', '文章操作', 'Bpvank - 文章操作', '/articleOperation', 'Reading', 0, 1);
INSERT INTO `backstagemenuinfo` VALUES (3, 'addArticle', '新建文章', 'Bpvank - 新建文章', '/addArticle', 'EditPen', 2, 1);
INSERT INTO `backstagemenuinfo` VALUES (4, 'articleManager', '文章管理', 'Bpvank - 文章管理', '/articleManager', 'Edit', 2, 1);
INSERT INTO `backstagemenuinfo` VALUES (5, 'classifyManager', '分类管理', 'Bpvank - 分类管理', '/classifyManager', 'FolderOpened', 2, 0);
INSERT INTO `backstagemenuinfo` VALUES (6, 'commentManager', '评论管理', 'Bpvank - 评论管理', '/commentManager', 'ChatDotSquare', 0, 1);
INSERT INTO `backstagemenuinfo` VALUES (7, 'userManager', '用户管理', 'Bpvank - 用户管理', '/userManager', 'MagicStick', 0, 0);
INSERT INTO `backstagemenuinfo` VALUES (8, 'linkManager', '链接管理', 'Bpvank - 链接管理', '/linkManager', 'Link', 0, 0);
INSERT INTO `backstagemenuinfo` VALUES (9, 'personSetup', '个人设置', 'Bpvank - 个人设置', '/personSetup', 'User', 0, 1);
INSERT INTO `backstagemenuinfo` VALUES (10, 'systemSetup', '系统设置', 'Bpvank - 系统设置', '/systemSetup', 'Setting', 0, 0);
INSERT INTO `backstagemenuinfo` VALUES (11, 'carouseMap', '轮播图', 'Bpvank - 轮播图', '/carouseMap', 'View', 0, 0);

-- ----------------------------
-- Table structure for carousemapinfo
-- ----------------------------
DROP TABLE IF EXISTS `carousemapinfo`;
CREATE TABLE `carousemapinfo`  (
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 38 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of carousemapinfo
-- ----------------------------
INSERT INTO `carousemapinfo` VALUES ('4f8c4ce8a0ee4e90bc85e0d363ebd658.png', 'static/advertisingImg/4f8c4ce8a0ee4e90bc85e0d363ebd658.png', 35);
INSERT INTO `carousemapinfo` VALUES ('banner01.png', 'static/advertisingImg/banner01.png', 36);
INSERT INTO `carousemapinfo` VALUES ('IMG_20221101_203906.jpg', 'static/advertisingImg/IMG_20221101_203906.jpg', 37);

-- ----------------------------
-- Table structure for classifyinfo
-- ----------------------------
DROP TABLE IF EXISTS `classifyinfo`;
CREATE TABLE `classifyinfo`  (
  `classifyIntroduce` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '分类介绍',
  `classifyId` int(11) NOT NULL AUTO_INCREMENT COMMENT '分类编号',
  `classifyName` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '分类名称',
  `articleNumber` int(11) NOT NULL DEFAULT 0 COMMENT '文章数量',
  `color1` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '渐变色1',
  `color2` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '渐变色2',
  PRIMARY KEY (`classifyId`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 27 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '分类管理表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of classifyinfo
-- ----------------------------
INSERT INTO `classifyinfo` VALUES ('日常挖坑集', 1, '技术', 7, '#6D80FE', '#23D2FD');
INSERT INTO `classifyinfo` VALUES ('创意无限大', 2, '资源', 0, '#FBF40F', '#FFA9AB');
INSERT INTO `classifyinfo` VALUES ('七七八八十分精彩', 3, '杂文', 0, '#09B0E8', '#29F49A');
INSERT INTO `classifyinfo` VALUES ('记录美好生活', 4, '生活', 0, '#717CFE', '#FC83EC');
INSERT INTO `classifyinfo` VALUES ('emo患者区', 5, '情感', 1, '#535b9a', '#30bcd7');
INSERT INTO `classifyinfo` VALUES ('踩坑喜加一', 6, '其他', 0, '#FF988B', '#FF6B88');

-- ----------------------------
-- Table structure for commentinfo
-- ----------------------------
DROP TABLE IF EXISTS `commentinfo`;
CREATE TABLE `commentinfo`  (
  `commentId` int(11) NOT NULL AUTO_INCREMENT COMMENT '评论编号',
  `articleId` int(11) NOT NULL COMMENT '文章代号',
  `userId` int(11) NOT NULL COMMENT '评论用户编号',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '评论内容',
  `parentId` int(11) NULL DEFAULT 0 COMMENT '父评论',
  `commentDate` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '评论日期',
  PRIMARY KEY (`commentId`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 139 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of commentinfo
-- ----------------------------
INSERT INTO `commentinfo` VALUES (138, 112, 1, '1', 0, '2022-12-27 11:14:12');

-- ----------------------------
-- Table structure for effects
-- ----------------------------
DROP TABLE IF EXISTS `effects`;
CREATE TABLE `effects`  (
  `effects01` tinyint(1) NULL DEFAULT 0 COMMENT '特效1',
  `effects02` tinyint(1) NULL DEFAULT 0 COMMENT '特效2'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of effects
-- ----------------------------
INSERT INTO `effects` VALUES (0, 0);

-- ----------------------------
-- Table structure for logininfo
-- ----------------------------
DROP TABLE IF EXISTS `logininfo`;
CREATE TABLE `logininfo`  (
  `loginId` int(11) NOT NULL COMMENT '登录的用户编号',
  `loginTime` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '登录时间'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '登录管理表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of logininfo
-- ----------------------------
INSERT INTO `logininfo` VALUES (1, '2022-11-02 17:19:36');
INSERT INTO `logininfo` VALUES (1, '2022-11-03 09:40:57');
INSERT INTO `logininfo` VALUES (1, '2022-11-03 13:02:24');
INSERT INTO `logininfo` VALUES (1, '2022-11-05 14:42:36');
INSERT INTO `logininfo` VALUES (1, '2022-11-05 15:01:52');
INSERT INTO `logininfo` VALUES (1, '2022-11-19 16:35:25');
INSERT INTO `logininfo` VALUES (1, '2022-11-20 16:22:31');
INSERT INTO `logininfo` VALUES (1, '2022-11-21 21:24:57');
INSERT INTO `logininfo` VALUES (1, '2022-11-21 21:34:50');
INSERT INTO `logininfo` VALUES (1, '2022-11-26 13:40:55');
INSERT INTO `logininfo` VALUES (1, '2022-12-03 16:02:30');
INSERT INTO `logininfo` VALUES (1, '2022-12-04 10:12:57');
INSERT INTO `logininfo` VALUES (1, '2022-12-05 10:44:17');
INSERT INTO `logininfo` VALUES (2, '2022-12-05 11:43:07');
INSERT INTO `logininfo` VALUES (1, '2022-12-05 13:43:46');
INSERT INTO `logininfo` VALUES (1, '2022-12-05 13:44:32');
INSERT INTO `logininfo` VALUES (1, '2022-12-05 15:10:27');
INSERT INTO `logininfo` VALUES (1, '2022-12-05 15:31:22');
INSERT INTO `logininfo` VALUES (1, '2022-12-05 15:33:50');
INSERT INTO `logininfo` VALUES (1, '2022-12-10 14:01:37');
INSERT INTO `logininfo` VALUES (1, '2022-12-10 14:17:44');
INSERT INTO `logininfo` VALUES (1, '2022-12-10 14:21:31');
INSERT INTO `logininfo` VALUES (1, '2022-12-10 14:25:14');
INSERT INTO `logininfo` VALUES (1, '2022-12-10 14:25:37');
INSERT INTO `logininfo` VALUES (1, '2022-12-10 14:30:15');
INSERT INTO `logininfo` VALUES (1, '2022-12-10 14:31:39');
INSERT INTO `logininfo` VALUES (1, '2022-12-10 14:32:26');
INSERT INTO `logininfo` VALUES (2, '2022-12-10 20:21:57');
INSERT INTO `logininfo` VALUES (1, '2022-12-10 20:25:58');
INSERT INTO `logininfo` VALUES (2, '2022-12-10 20:26:11');
INSERT INTO `logininfo` VALUES (2, '2022-12-10 20:30:38');
INSERT INTO `logininfo` VALUES (1, '2022-12-10 21:22:53');
INSERT INTO `logininfo` VALUES (1, '2022-12-11 09:39:14');
INSERT INTO `logininfo` VALUES (1, '2022-12-11 11:35:55');
INSERT INTO `logininfo` VALUES (2, '2022-12-11 11:36:32');
INSERT INTO `logininfo` VALUES (1, '2022-12-11 11:36:54');
INSERT INTO `logininfo` VALUES (1, '2022-12-11 12:59:15');
INSERT INTO `logininfo` VALUES (1, '2022-12-15 23:51:32');
INSERT INTO `logininfo` VALUES (38, '2022-12-15 23:53:06');
INSERT INTO `logininfo` VALUES (1, '2022-12-16 00:05:26');
INSERT INTO `logininfo` VALUES (1, '2022-12-16 00:07:41');
INSERT INTO `logininfo` VALUES (1, '2022-12-16 21:10:11');
INSERT INTO `logininfo` VALUES (1, '2022-12-17 11:15:07');
INSERT INTO `logininfo` VALUES (1, '2022-12-17 12:44:34');
INSERT INTO `logininfo` VALUES (1, '2022-12-17 13:26:17');
INSERT INTO `logininfo` VALUES (1, '2022-12-17 20:13:40');
INSERT INTO `logininfo` VALUES (1, '2022-12-17 20:48:28');
INSERT INTO `logininfo` VALUES (1, '2022-12-17 21:01:20');
INSERT INTO `logininfo` VALUES (1, '2022-12-17 22:34:28');
INSERT INTO `logininfo` VALUES (1, '2022-12-22 19:12:51');
INSERT INTO `logininfo` VALUES (1, '2022-12-22 19:30:48');
INSERT INTO `logininfo` VALUES (1, '2022-12-24 17:37:42');
INSERT INTO `logininfo` VALUES (1, '2022-12-25 12:05:19');
INSERT INTO `logininfo` VALUES (1, '2022-12-26 20:53:29');
INSERT INTO `logininfo` VALUES (1, '2022-12-26 22:01:59');
INSERT INTO `logininfo` VALUES (1, '2022-12-26 22:14:00');
INSERT INTO `logininfo` VALUES (1, '2022-12-27 09:50:05');
INSERT INTO `logininfo` VALUES (1, '2022-12-28 16:37:57');
INSERT INTO `logininfo` VALUES (1, '2022-12-28 22:36:13');
INSERT INTO `logininfo` VALUES (1, '2022-12-30 14:14:34');
INSERT INTO `logininfo` VALUES (1, '2022-12-30 22:37:05');
INSERT INTO `logininfo` VALUES (1, '2022-12-31 14:26:42');

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message`  (
  `messageId` int(11) NOT NULL AUTO_INCREMENT COMMENT '留言编号',
  `messageName` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '留言昵称',
  `messageQQ` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '留言QQ',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '留言内容',
  `replyContent` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '回复留言内容',
  `messageDate` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '留言日期',
  PRIMARY KEY (`messageId`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 29 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of message
-- ----------------------------
INSERT INTO `message` VALUES (20, '23213', '123123', '2eeeeeeeeee', '112213123123123122', '2022-11-26 16:57:16');
INSERT INTO `message` VALUES (28, 'asdasd', 'asdasdas', 'dasdas', NULL, '2022-12-03 16:04:40');

-- ----------------------------
-- Table structure for my
-- ----------------------------
DROP TABLE IF EXISTS `my`;
CREATE TABLE `my`  (
  `articleId` int(11) NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of my
-- ----------------------------
INSERT INTO `my` VALUES (11);

-- ----------------------------
-- Table structure for systemsetup
-- ----------------------------
DROP TABLE IF EXISTS `systemsetup`;
CREATE TABLE `systemsetup`  (
  `effects03` tinyint(1) NULL DEFAULT NULL COMMENT '特效2黑白',
  `effects04` tinyint(1) NULL DEFAULT NULL,
  `effects02` tinyint(1) NULL DEFAULT 0 COMMENT '特效2',
  `effects01` tinyint(1) NULL DEFAULT 0 COMMENT '特效1',
  `stickArticle` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '置顶',
  `allArticle` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '文章展示',
  `featuredArticle` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '左侧精选文章',
  `technologyArticle` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '左侧技术专区文章',
  `resourceArticle` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '右侧资源专区文章',
  `advertising1` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '左侧广告图1',
  `advertisingLink1` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '左侧广告链接1',
  `advertising2` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '右侧广告图1',
  `advertisingLink2` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '左侧广告链接1',
  `advertising3` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '右侧广告图2',
  `advertising5` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '背景',
  `advertising4` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '右上角',
  `advertisingLink3` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '左侧广告链接1',
  `tableName` tinyint(4) NULL DEFAULT NULL COMMENT '用来改表'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of systemsetup
-- ----------------------------
INSERT INTO `systemsetup` VALUES (0, 1, 0, 0, '112', '[110,111,109,108,107]', '[113]', '[108,107,111]', '[]', '', '', '', '', '', '', 'static/advertisingImg/QQ截图20221226205311.jpg', '', 1);

-- ----------------------------
-- Table structure for urlinfo
-- ----------------------------
DROP TABLE IF EXISTS `urlinfo`;
CREATE TABLE `urlinfo`  (
  `urlId` int(11) NOT NULL AUTO_INCREMENT COMMENT '链接编号',
  `urlName` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '网站名称',
  `urlAddres` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '网站地址',
  `urlReferral` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '网站介绍',
  `urlLitimg` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '网站缩略图',
  `webmasterEmail` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '站长邮箱',
  `urlPass` int(11) NULL DEFAULT 1 COMMENT '后台审核是否通过',
  `urlType` int(11) NULL DEFAULT NULL COMMENT '链接类型',
  PRIMARY KEY (`urlId`) USING BTREE,
  UNIQUE INDEX `urlAddres`(`urlAddres`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 28 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '链接管理表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of urlinfo
-- ----------------------------
INSERT INTO `urlinfo` VALUES (26, '极简壁纸', 'https://bz.zzzmh.cn/index', '免费好用的电脑壁纸网站', 'static/upload/c5398839880411ebb6edd017c2d2eca2.jpg', 'admin@zzzmh.cn', 2, 1);
INSERT INTO `urlinfo` VALUES (27, '李文周的博客', 'https://www.liwenzhou.com/', '李文周老师的GO博客', 'static/upload/Dingtalk_20221230141934.jpg', '2457665356@qq.com', 2, 1);

-- ----------------------------
-- Table structure for userinfo
-- ----------------------------
DROP TABLE IF EXISTS `userinfo`;
CREATE TABLE `userinfo`  (
  `userId` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户编号',
  `userType` int(11) NOT NULL DEFAULT 1 COMMENT '用户类型',
  `userName` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名',
  `userPass` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码',
  `userEmail` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '邮箱',
  `userRegdate` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '注册时间',
  `userSignature` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '个性签名',
  `userIcon` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '个人头像',
  PRIMARY KEY (`userId`) USING BTREE,
  UNIQUE INDEX `userName`(`userName`) USING BTREE,
  UNIQUE INDEX `userEmail`(`userEmail`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 39 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '注册用户管理表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of userinfo
-- ----------------------------
INSERT INTO `userinfo` VALUES (1, 0, 'admin', '$2a$10$Evemu2Qj8s3U6tAA2qYhqO2/u2ezEAetR77NN2VpEJGOY5jsjdFL2', '2457665356@qq.com', '2022-07-31 10:10:56', '若没有坚持到底的勇气，还是不要开始好了', 'static/userIcon/1665192663091_avatar.jpg');

SET FOREIGN_KEY_CHECKS = 1;
