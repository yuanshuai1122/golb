import request from "../utils/request";
import {BASE_URL} from "../constants";


/**
 * 获取分类列表
 */
export const getCategoryList= (): Promise<any> => request.get(`${BASE_URL}/category/list`);