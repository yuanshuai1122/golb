import React, {useEffect, useState} from 'react';
import {useParams} from "react-router-dom";
import {getArticlesDetail} from "../../services/article";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import {BlogInfo} from "../../types/blog";
import ReactMarkdown from "react-markdown";
import gfm from 'remark-gfm'

const BlogDetail = () => {

    const [blogInfo, setBlogInfo] = useState<BlogInfo>();

    let params = useParams()

    useEffect(()=> {
        console.log(params.id)
        if (params.id != undefined) {
            getArticlesDetail(Number.parseInt(params.id)).then(res =>{
                setBlogInfo(res.data)
            })
        }
    }, [])

    return (
        <>
            <br/>
            <Box sx={{ width: '100%'}}>
                <Typography variant="h5" gutterBottom>
                    {blogInfo?.title}
                </Typography>
                <hr/>
                <Typography variant="subtitle1" gutterBottom>
                    <ReactMarkdown
                        remarkPlugins={[gfm]}>
                        {blogInfo?.content === undefined ? "" : blogInfo.content}
                    </ReactMarkdown>
                </Typography>
            </Box>
        </>
    );
};

export default BlogDetail;