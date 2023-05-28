import React from 'react';
import NavTabs from "./components/NavTabs";
import {Container, CssBaseline} from "@mui/material";
import {Route, Routes} from "react-router-dom";
import Home from "./pages/Home";
import About from "./pages/About";
import Archives from "./pages/Archives";
import BlogDetail from "./pages/BlogDetail";

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
                            <Route path='/blog/:id' element={<BlogDetail/>}/>
                        </Routes>
                    </Container>
            </React.Fragment>
        </div>
    );
};

export default App;