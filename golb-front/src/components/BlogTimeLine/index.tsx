import * as React from 'react';
import {useEffect, useState} from 'react';
import Timeline from '@mui/lab/Timeline';
import TimelineItem from '@mui/lab/TimelineItem';
import TimelineSeparator from '@mui/lab/TimelineSeparator';
import TimelineConnector from '@mui/lab/TimelineConnector';
import TimelineContent from '@mui/lab/TimelineContent';
import TimelineDot from '@mui/lab/TimelineDot';
import TimelineOppositeContent from '@mui/lab/TimelineOppositeContent';
import {Pagination, Stack} from "@mui/material";
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
 * 处理归档数据
 * @param data
 */
const handleArchiveData = (data: any[]) => {
    const archiveLists: IArchive[] = []
    for (let datum of data) {
        const archive: IArchive = {
            title: datum.title,
            createTime: datum.createTime
        }
        archiveLists.push(archive)
    }

    return archiveLists
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
        getArchivesList(1, 8).then(res => {
            setArchivePages(res.data.totalPage)
            const archiveLists: IArchive[] = handleArchiveData(res.data.list);
            setArchiveData(archiveLists)
        })
    }, [])

    /**
     * 处理分页数据
     * @param event 分页事件
     * @param page 页数
     */
    const handlePagination = (event: React.ChangeEvent<unknown>, page: number) => {
        // 获取归档列表
        getArchivesList(page, 8).then(res => {
            setArchivePages(res.data.totalPage)
            const archiveLists: IArchive[] = handleArchiveData(res.data.list);
            setArchiveData(archiveLists)
        })
    }

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
                <Pagination
                    count={archivePages}
                    onChange={handlePagination}
                />
            </Stack>
        </>
    );
}