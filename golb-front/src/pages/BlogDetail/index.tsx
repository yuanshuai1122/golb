import React, {useEffect} from 'react';
import {useParams} from "react-router-dom";

const BlogDetail = () => {

    let params = useParams()

    useEffect(()=> {
        console.log(params.id)
    }, [])

    return (
        <div>
            博客详情页面 {params.id}
        </div>
    );
};

export default BlogDetail;