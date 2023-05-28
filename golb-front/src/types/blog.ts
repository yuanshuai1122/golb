
// 博客列表types
interface IBlogList {
    id: number
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