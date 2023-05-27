import request from "../utils/request";
import {BASE_URL} from "../constants";


/**
 * 获取归档列表
 * @param pageNum 页数
 * @param pageSize 页大小
 */
export const getArchivesList = (pageNum: number, pageSize: number): Promise<any> => request.get(`${BASE_URL}/archives/list`, {
    pageNum,
    pageSize
});