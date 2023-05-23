import React from 'react';
import BlogCard from "../BlogCard";
import NavTabs from "../NavTabs";
import {Container, CssBaseline} from "@mui/material";
import Box from "@mui/material/Box";
import {Link, Route, Routes} from "react-router-dom";
import Home from "../../pages/Home";
import About from "../../pages/About";

const App = () => {
    return (
        <div>
            <React.Fragment>
                <CssBaseline />
                    <Container maxWidth="lg">
                        <NavTabs/>
                        <Link className='list-group-item' to="/about">About</Link>
                        <Link className='list-group-item' to="/">Home</Link>
                        <Routes>
                            <Route path="/" element={<Home />}></Route>
                            <Route path="/about" caseSensitive element={<About/>}></Route>
                        </Routes>
                    </Container>
            </React.Fragment>
        </div>
    );
};

export default App;