import * as React from 'react';
import Box from '@mui/material/Box';
import Tabs from '@mui/material/Tabs';
import Tab from '@mui/material/Tab';
import {Link, matchPath, useLocation} from "react-router-dom";




function useRouteMatch(patterns: readonly string[]) {
    const { pathname } = useLocation();

    for (let i = 0; i < patterns.length; i += 1) {
        const pattern = patterns[i];
        const possibleMatch = matchPath(pattern, pathname);
        if (possibleMatch !== null) {
            return possibleMatch;
        }
    }

    return null;
}

export default function NavTabs() {

    const routeMatch = useRouteMatch(['/', '/archives', '/about']);
    const currentTab = routeMatch?.pattern?.path;

    return (
        <Box sx={{ width: '100%' }}>
            <Tabs value={currentTab} centered>
                <Tab label="首页" value="/" to="/" component={Link} />
                <Tab label="归档" value="/archives" to="/archives" component={Link} />
                <Tab label="关于我" value="/about" to="/about" component={Link} />
            </Tabs>
        </Box>
    );
}