import React from 'react';
import BlogCard from "../../components/BlogCard";
import BlogPagination from "../../components/BlogPagination";
import BlogCarousel from "../../components/BlogCarousel";
import BlogList from "../../components/BlogList";

const Home = () => {
    return (
        <div>
            <br/>
            <BlogList/>
            <br/>
            <BlogPagination/>
        </div>
    );
};

export default Home;