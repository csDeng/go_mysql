"use strict";
import $axios from '~/axios';
const Tags = {
    get(){
        return $axios.get('/tags')
    },
    add(params={}){
        return $axios.post('/tags/',params)
    },
    delete(id){
        return $axios.delete('/tags/' + id)
    }
}
export default Tags;