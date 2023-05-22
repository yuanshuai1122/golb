import React from 'react';
import BlogCard from "../BlogCard";
import NavTabs from "../NavTabs";
import {Container, CssBaseline} from "@mui/material";
import Box from "@mui/material/Box";

const App = () => {
    return (
        <div>
            <React.Fragment>
                <CssBaseline />
                    <Container maxWidth="lg">
                        <NavTabs/>
                        <BlogCard/>
                        <Box sx={{ bgcolor: '#cfe8fc', height: '100vh' }} />
                    </Container>
            </React.Fragment>
        </div>
    );
};

export default App;