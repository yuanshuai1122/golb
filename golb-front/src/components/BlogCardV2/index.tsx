import * as React from 'react';
import { styled } from '@mui/material/styles';
import Grid from '@mui/material/Grid';
import Paper from '@mui/material/Paper';
import Typography from '@mui/material/Typography';
import ButtonBase from '@mui/material/ButtonBase';
import {IBlogList} from "../../types/blog";
import {Button} from "@mui/material";
import {useNavigate} from "react-router-dom";
import Link from '@mui/material/Link';

const Img = styled('img')({
    margin: 'auto',
    display: 'block',
    maxWidth: '100%',
    maxHeight: '100%',
});

export default function BlogCard(props: IBlogList) {

    const {id,  title, views, userName,  coverImg, createTime, abstract} = props

    return (
        <Paper
            sx={{
                p: 2,
                margin: 'auto',
                flexGrow: 1,
                backgroundColor: (theme) =>
                    theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
            }}
        >
            <Grid container spacing={2}>
                <Grid item>
                    <ButtonBase sx={{ width: 200, height: 128 }}>
                        <Img alt="complex" src = {coverImg} />
                    </ButtonBase>
                </Grid>
                <Grid item xs={12} sm container>
                    <Grid item xs container direction="column" spacing={2}>
                        <Grid item xs>
                            <Link href={"/blog/" + id} underline="hover">
                                <Typography gutterBottom variant="h6" component="div">
                                    {title}
                                </Typography>
                            </Link>
                            <Typography variant="body2" gutterBottom>
                                {abstract}
                            </Typography>
                        </Grid>
                        <Grid item>
                            <Typography sx={{ cursor: 'pointer' }} variant="body2">
                                {createTime} -- {views}
                            </Typography>
                        </Grid>
                    </Grid>
                    <Grid item>
                        <Typography variant="subtitle1" component="div">
                            {userName}
                        </Typography>
                    </Grid>
                </Grid>
            </Grid>
        </Paper>
    );
}