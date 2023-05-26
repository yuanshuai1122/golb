
// 博客列表types
interface IBlogList {
    coverImg: string
    title: string
    channel: string
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