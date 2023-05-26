import request from "../utils/request";
import {BASE_URL} from "../constants";

/**
 * 获取文章列表
 * @param pageNum
 * @param pageSize
 * @param keywords
 */
export const getArticlesList = (pageNum: number, pageSize: number, keywords: string): Promise<any> => request.get(`${BASE_URL}/articles/list`, {});