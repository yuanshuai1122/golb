import React from 'react';
import BlogCard from "../../components/BlogCard";
import BlogPagination from "../../components/BlogPagination";

const Home = () => {
    return (
        <div>
            <br/>
            <BlogCard/>
            <BlogCard/>
            <BlogCard/>
            <BlogCard/>
            <BlogCard/>
            <BlogCard/>
            <br/>
            <BlogPagination/>
        </div>
    );
};

export default Home;