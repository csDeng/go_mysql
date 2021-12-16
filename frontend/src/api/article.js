"use strict";
import $axios from '~/axios';

const article = {
    get(id){
        return $axios.get('/blogs/'+id)
    },
    save(params){
        return $axios.post('/blogs/create', {...params} )
    },
    get_list(){
        // if(pid){
        //     return $axios.get('/blogs?pid='+pid)
        // }
        return $axios.get('/blogs')
    },
    del(id){
        return $axios.delete('/blogs/' + id)
    },
    get_index(){
        // if(pid){
        //     return $axios.get('/blogs?pid='+pid)
        // }
        return $axios.get('/blogs/index')
    },
}

export default article