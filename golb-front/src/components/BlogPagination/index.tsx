import React from 'react';
import {Pagination, Stack} from "@mui/material";

const BlogPagination = () => {
    return (
        <Stack spacing={2}>
            <Pagination count={10}/>
        </Stack>
    );
};

export default BlogPagination;