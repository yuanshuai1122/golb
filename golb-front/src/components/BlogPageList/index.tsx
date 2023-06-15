import * as React from 'react';
import {useEffect, useState} from 'react';
import Box from '@mui/material/Box';
import {Pagination, Stack} from "@mui/material";
import {getArticlesList} from "../../services/article";
import {IBlogList, IBlogRowsList} from "../../types/blog";
import BlogCard from "../BlogCard";


/**
 * 处理博客页数据
 * @param data 数据源
 */
const handlePageData = (data: any[]) => {
    const blogCardList: IBlogList[] = []
    data.map((item, index) => {
        const blog: IBlogList = {
            id: item.id,
            coverImg: item.coverImg,
            title: item.title,
            abstract: item.abstract,
            userName: "aabb",
            views: item.views,
            categoryName: item.categoryName,
            createTime: item.createTime,
        }
        blogCardList.push(blog)
    })

    return blogCardList
}

/**
 * 博客页面列表组件
 * @constructor
 */
export default function BlogPageList() {

    const [blogRowsList, setBlogRowsList] = useState<IBlogRowsList[]>([]);
    const [blogList, setBlogList] = useState<IBlogList[]>([]);
    const [totalPages, setTotalPages] = useState<number>(0);

    useEffect(()=> {
        // 获取文章列表
        getArticlesList(1, 8, "").then(res => {
            // 总页数
            setTotalPages(res.data.totalPage)
            // 处理博客数据
            // 处理博客数据
            const blogs: IBlogList[] = handlePageData(res.data.list);
            setBlogList(blogs)

        }).catch(e => {
            console.log(e)
        })
    }, [])

    // 分页handle
    const handlePagination = (event: React.ChangeEvent<unknown>, page: number) => {
        // 获取文章列表
        getArticlesList(page, 8, "").then(res => {
            // 总页数
            setTotalPages(res.data.totalPage)
            // 处理博客数据
            const blogs: IBlogList[] = handlePageData(res.data.list);
            setBlogList(blogs)
        }).catch(e => {
            console.log(e)
        })
    }

    return (
        <>
            <Box sx={{ overflow: 'hidden'}}>
                {
                    blogList.map((item, index) => {
                        return (
                            <>
                                <BlogCard
                                    key={index}
                                    id={item.id}
                                    title={item.title}
                                    abstract={item.abstract}
                                    userName={item.userName}
                                    views={item.views}
                                    coverImg={item.coverImg}
                                    categoryName={item.categoryName}
                                    createTime={item.createTime}
                                />
                            </>
                        )
                    })
                }
            </Box>
            <Stack sx={{alignItems: 'center'}} spacing={2}>
                <Pagination
                    count={totalPages}
                    onChange={handlePagination}
                />
            </Stack>
        </>


    );
}
