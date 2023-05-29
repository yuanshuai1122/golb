import React, {useEffect} from 'react';
import {useParams} from "react-router-dom";
import {getArticlesDetail} from "../../services/article";

const BlogDetail = () => {

    let params = useParams()

    useEffect(()=> {
        console.log(params.id)
        if (params.id != undefined) {
            getArticlesDetail(Number.parseInt(params.id)).then(res =>{
                console.log(res)
            })
        }
    }, [])

    return (
        <div>
            博客详情页面 {params.id}
        </div>
    );
};

export default BlogDetail;