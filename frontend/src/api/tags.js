"use strict";
import $axios from '~/axios';
const prefix = "/api/v1/tags"
const Tags = {
    get_list(pid=0){
        return $axios.get(prefix+"?page="+pid)
    },
    add(params={}){
        return $axios.post(prefix,params)
    },
    delete(id){
        return $axios.delete(prefix + id)
    }
}
export default Tags;