
// 博客列表types
interface IBlogList {
    id: number
    coverImg: string
    title: string
    userName: string
    views: number
    createTime: string
}

// 博客行数列表
interface IBlogRowsList {
    blogList: IBlogList[]
}

// 博客详情
interface BlogInfo {
    abstract: string
    content: string
    coverImg: string
    createTime: string
    title: string
    updateTime: string
    views: number

}

export type {
    IBlogList,
    IBlogRowsList,
    BlogInfo
}