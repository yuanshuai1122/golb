import * as React from 'react';
import {styled} from '@mui/material/styles';
import Grid from '@mui/material/Grid';
import Paper from '@mui/material/Paper';
import Typography from '@mui/material/Typography';
import ButtonBase from '@mui/material/ButtonBase';
import {IBlogList} from "../../types/blog";
import Link from '@mui/material/Link';
import {dateFormat} from "../../utils/dateUtils";

const Img = styled('img')({
    margin: 'auto',
    display: 'block',
    maxWidth: '100%',
    maxHeight: '100%',
});

export default function BlogCard(props: IBlogList) {

    const {id,  title, views, userName,  coverImg, createTime, abstract, categoryName} = props

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
                                {dateFormat(createTime, 'YYYY-MM-DD HH:mm')} | 分类：{categoryName}｜浏览量：{views}
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