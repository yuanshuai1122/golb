import request from "../utils/request";
import {BASE_URL} from "../constants";

/**
 * 获取文章列表
 * @param pageNum
 * @param pageSize
 * @param keywords
 */
export const getArticlesList = (pageNum: number, pageSize: number, keywords: string): Promise<any> => request.get(`${BASE_URL}/articles/list`, {
    pageNum,
    pageSize,
    keywords
});

/**
 * 根据id获取文章详情
 * @param id 文章id
 */
export const getArticlesDetail = (id: number): Promise<any> => request.get(`${BASE_URL}/articles/detail`, {
    id
});