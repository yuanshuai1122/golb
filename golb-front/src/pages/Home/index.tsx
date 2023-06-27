import React, {useEffect} from 'react';
import BlogPageList from "../../components/BlogPageList";
import Typography from "@mui/material/Typography";
import request from "../../utils/request";

const Home = () => {

    return (
        <div>
            <div style={{textAlign: "center", marginTop: 40, marginBottom: 30}}>
                <Typography variant="h4" gutterBottom>
                    aabb的文档
                </Typography>
                <Typography variant="subtitle2" gutterBottom>
                    这是aabb的文档，balabalabala...
                </Typography>
                <div>图标图标图标</div>
            </div>
            <BlogPageList/>
        </div>
    );
};

export default Home;