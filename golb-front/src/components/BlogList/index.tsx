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


interface MediaProps {
    loading?: boolean;
    blogList: IBlogList[]
}

function BlogCard(props: MediaProps) {
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

export default function BlogList() {

    const [blogRowsList, setBlogRowsList] = useState<IBlogRowsList[]>([]);

    useEffect(()=> {
        // 获取文章列表
        getArticlesList(1, 6, "").then(res => {
            console.log(res.data.list)
            const blogRowsList: IBlogRowsList[] = [];
            // 切片
            let arrLen = 3;  //这里一行数组最多3个
            for (let i=0; i<res.data.list.length; i+=arrLen) {
                const blogSlice = res.data.list.slice(i,i+arrLen)
                const blogList: IBlogList[] = []
                for (let blogSliceElement of blogSlice) {
                    const blog: IBlogList = {
                        coverImg: blogSliceElement.coverImg,
                        title: blogSliceElement.title,
                        channel: "aabb",
                        views: 100,
                        createTime: blogSliceElement.createTime
                    }
                    blogList.push(blog)
                }
                blogRowsList.push({
                    blogList: blogList
                });
            }
            setBlogRowsList(blogRowsList)

        })
    }, [])

    return (
        <>
            <Box sx={{ overflow: 'hidden'}}>
                {
                    blogRowsList.map((item) => {
                        return <BlogCard blogList={item.blogList} />
                    })
                }
            </Box>
            <Stack sx={{alignItems: 'center'}} spacing={2}>
                <Pagination count={10}/>
            </Stack>
        </>


    );
}
