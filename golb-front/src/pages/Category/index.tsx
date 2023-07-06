import React, {useEffect} from 'react';
import {makeStyles} from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import FolderIcon from '@material-ui/icons/Folder';
import {JSX} from 'react/jsx-runtime';
import {getCategoryList} from "../../services/category";
import {ICategory} from "../../types/category";
import {ListItemButton} from "@mui/material";

const useStyles = makeStyles((theme) => ({
    root: {
        flexGrow: 1,
        maxWidth: 752,
    },
    demo: {
        backgroundColor: theme.palette.background.paper,
    },
    title: {
        margin: theme.spacing(4, 0, 2),
    },
}));

function generate(element: JSX.Element) {
    return [0, 1, 2].map((value) =>
        React.cloneElement(element, {
            key: value,
        }),
    );
}

export default function Category() {
    const classes = useStyles();
    const [dense, setDense] = React.useState(false);
    const [secondary, setSecondary] = React.useState(false);
    const [categoryList, setCategoryList] = React.useState<ICategory[]>([]);

    useEffect(()=> {
        // 设置紧凑模式
        setDense(false)
        // 设置显示第二行
        setSecondary(true)
        // 查询分类列表
        getCategoryList().then(r => {
            console.log(r.data)
            setCategoryList(r.data)
        }).catch(e => {
            console.log(e)
        })
    }, [])


    return (
        <div className={classes.root}>
            <Grid container spacing={2}>
                <Grid item xs={12} md={6}>
                    <Typography variant="h6" className={classes.title}>
                        分类
                    </Typography>
                    <div className={classes.demo}>
                        <List dense={dense}>
                            {
                                categoryList.map((item, index) => {
                                    return (
                                        <div key={index}>
                                            <ListItem>
                                                <ListItemButton
                                                    dense={true}
                                                    onClick={()=> {
                                                        console.log(item.id)
                                                    }
                                                }
                                                >
                                                    <ListItemIcon>
                                                        <FolderIcon />
                                                    </ListItemIcon>
                                                    <ListItemText
                                                        primary = {item.categoryName}
                                                        secondary={secondary ? item.categoryDesc : null}
                                                    />
                                                </ListItemButton>
                                            </ListItem>
                                        </div>
                                    )
                                })
                            }
                        </List>
                    </div>
                </Grid>
            </Grid>
        </div>
    );
}
