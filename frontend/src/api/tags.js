"use strict";
import $axios from '~/axios';
const prefix = "/api/v1/tags"
const Tags = {
    get(){
        return $axios.get('/')
    },
    add(params={}){
        return $axios.post(prefix,params)
    },
    delete(id){
        return $axios.delete(prefix + id)
    }
}
export default Tags;