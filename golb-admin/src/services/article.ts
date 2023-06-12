import request from "../utils/request";
import {BASE_URL} from "../constants";
import {IPostArticle} from "@/types/article";

/**
 * 创建文章
 * @param data 创建文章内容
 */
export const postArticle = (data: IPostArticle): Promise<any> => request.post(`${BASE_URL}/admin/articles/create`, data);