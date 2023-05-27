import * as React from 'react';
import Timeline from '@mui/lab/Timeline';
import TimelineItem from '@mui/lab/TimelineItem';
import TimelineSeparator from '@mui/lab/TimelineSeparator';
import TimelineConnector from '@mui/lab/TimelineConnector';
import TimelineContent from '@mui/lab/TimelineContent';
import TimelineDot from '@mui/lab/TimelineDot';
import TimelineOppositeContent from '@mui/lab/TimelineOppositeContent';
import {Pagination, Stack} from "@mui/material";
import {useEffect, useState} from "react";
import {getArchivesList} from "../../services/archive";
import {IArchive} from "../../types/archive";
import {dateFormat} from "../../utils/dateUtils";


/**
 * 博客时间线组件
 * @param archive 时间线参数
 * @constructor
 */
const BlogTimeLine = (archive :IArchive) => {
    return (
        <>
            <Timeline position="alternate">
                <TimelineItem>
                    <TimelineOppositeContent color="text.secondary">
                        {archive.createTime}
                    </TimelineOppositeContent>
                    <TimelineSeparator>
                        <TimelineDot />
                        <TimelineConnector />
                    </TimelineSeparator>
                    <TimelineContent>{archive.title}</TimelineContent>
                </TimelineItem>
            </Timeline>
        </>
    )
}

/**
 * 归档组件
 * @constructor
 */
export default function BlogArchives() {

    const [archiveData, setArchiveData] = useState<IArchive[]>([]);

    const [archivePages, setArchivePages] = useState<number>(0);


        useEffect(()=> {
        // 获取归档列表
        getArchivesList(1, 3).then(res => {
            console.log(res)
            setArchivePages(res.data.totalPage)
            const archiveLists: IArchive[] = []
            for (let datum of res.data.list) {
                const archive: IArchive = {
                    title: datum.title,
                    createTime: datum.createTime
                }
                archiveLists.push(archive)
            }
            setArchiveData(archiveLists)
        })
    }, [])

    return (
        <>
            {
                archiveData.map((item, index) => {
                    return (
                        <BlogTimeLine
                            key={index}
                            title={item.title}
                            createTime={dateFormat(item.createTime, 'MM-DD HH:mm:ss')}
                        />
                    )
                })
            }
            <Stack sx={{alignItems: 'center'}} spacing={2}>
                <Pagination count={archivePages}/>
            </Stack>
        </>
    );
}