import React from 'react';
import BlogCard from "../BlogCard";
import NavTabs from "../NavTabs";
import {Container, CssBaseline} from "@mui/material";
import Box from "@mui/material/Box";
import {Link, Route, Routes} from "react-router-dom";
import Home from "../../pages/Home";
import About from "../../pages/About";
import Archives from "../../pages/Archives";

const App = () => {
    return (
        <div>
            <React.Fragment>
                <CssBaseline />
                    <Container maxWidth="lg">
                        <NavTabs/>
                        <Routes>
                            <Route path="/" element={<Home />}></Route>
                            <Route path="/archives" element={<Archives />}></Route>
                            <Route path="/about" caseSensitive element={<About/>}></Route>
                        </Routes>
                    </Container>
            </React.Fragment>
        </div>
    );
};

export default App;