"use strict";
import $axios from '~/axios';

const prefix = "/api/v1/articles"
const article = {
    get(id){
        return $axios.get(prefix+id)
    },
    save(params){
        return $axios.post(prefix, {...params} )
    },
    get_list(pid=0){
        return $axios.get(prefix + '?page='+pid)
    },
    del(id){
        return $axios.delete(prefix + "/" + id)
    },
    get_index(){
        return $axios.get('index')
    },
}

export default article