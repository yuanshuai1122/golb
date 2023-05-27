
// 博客列表types
import React from "react";

interface IBlogList {
    coverImg: string
    title: string
    userName: string
    views: number
    createTime: string
}

interface IBlogRowsList {
    blogList: IBlogList[]
}

export type {
    IBlogList,
    IBlogRowsList
}