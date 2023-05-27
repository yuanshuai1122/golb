import * as React from 'react';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Skeleton from '@mui/material/Skeleton';
import {Pagination, Stack} from "@mui/material";
import {useEffect, useState} from "react";
import {getArticlesList} from "../../services/article";
import {IBlogList, IBlogRowsList} from "../../types/blog";
import {dateFormat} from "../../utils/dateUtils";


/**
 * 处理博客页数据
 * @param data 数据源
 */
const handlePageData = (data: any) => {
    const blogRowsList: IBlogRowsList[] = [];
    // 切片
    let arrLen = 3;  //这里一行数组最多3个
    for (let i=0; i<data.length; i+=arrLen) {
        const blogSlice = data.slice(i,i+arrLen)
        const blogList: IBlogList[] = []
        for (let blogSliceElement of blogSlice) {
            const blog: IBlogList = {
                coverImg: blogSliceElement.coverImg,
                title: blogSliceElement.title,
                userName: "aabb",
                views: 100,
                createTime: blogSliceElement.createTime
            }
            blogList.push(blog)
        }
        blogRowsList.push({
            blogList: blogList
        });
    }

    return blogRowsList;
}

/**
 * 博客卡片组件
 * @param props 参数
 * @constructor
 */
interface BlogCardProps {
    loading?: boolean;
    blogList: IBlogList[]
}
function BlogCard(props: BlogCardProps) {
    const { loading = false, blogList } = props;

    return (
        <Grid container wrap="nowrap" sx={{justifyContent: 'center'}} item spacing={3}>
            {(loading ? Array.from(new Array(3)) : blogList).map((item, index) => (
                <Box key={index} sx={{ width: 210, marginRight: 1, marginLeft: 1, my: 5 }}>
                    {item ? (
                        <img
                            style={{ width: 210, height: 140, borderRadius: 3}}
                            alt={item.title}
                            src={item.coverImg}
                        />
                    ) : (
                        <Skeleton variant="rectangular" width={210} height={100} />
                    )}
                    {item ? (
                        <Box sx={{ pr: 2 }}>
                            <Typography gutterBottom variant="body2">
                                {item.title}
                            </Typography>
                            <Typography display="block" variant="caption" color="text.secondary">
                                {item.channel}
                            </Typography>
                            <Typography variant="caption" color="text.secondary">
                                {`${item.views} 次浏览 • ${dateFormat(item.createTime, 'YYYY年MM月DD日')}`}
                            </Typography>
                        </Box>
                    ) : (
                        <Box sx={{ pt: 0.5 }}>
                            <Skeleton />
                            <Skeleton width="60%" />
                        </Box>
                    )}
                </Box>
            ))}
        </Grid>
    );
}

/**
 * 博客页面列表组件
 * @constructor
 */
export default function BlogPageList() {

    const [blogRowsList, setBlogRowsList] = useState<IBlogRowsList[]>([]);
    const [totalPages, setTotalPages] = useState<number>(0);

    useEffect(()=> {
        // 获取文章列表
        getArticlesList(1, 6, "").then(res => {
            console.log(res.data)
            // 总页数
            setTotalPages(res.data.totalPage)
            // 处理博客数据
            const blogRows: IBlogRowsList[] = handlePageData(res.data.list)
            setBlogRowsList(blogRows)

        })
    }, [])

    // 分页handle
    const handlePagination = (event: React.ChangeEvent<unknown>, page: number) => {
        // 获取文章列表
        getArticlesList(page, 6, "").then(res => {
            console.log("翻页获取文章列表")
            console.log(res.data)
            // 总页数
            setTotalPages(res.data.totalPage)
            // 处理博客数据
            const blogRows: IBlogRowsList[] = handlePageData(res.data.list)
            setBlogRowsList(blogRows)

        })
    }

    return (
        <>
            <Box sx={{ overflow: 'hidden'}}>
                {
                    blogRowsList.map((item, index) => {
                        return <BlogCard key={index} blogList={item.blogList} />
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
