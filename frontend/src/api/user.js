import $axios from '~/axios'

const auth = '/auths'
const users = '/users'
const user = {
    login(params){
        return $axios.post(auth + '/login',params)
    },
    register(params){
        return $axios.post(auth + '/register',params)
    },
    get_list(){
        return $axios.get(`${users}`)
    },
    update(params){
        console.log(params)
        return $axios.put(`${users}/${params.id}`,params)
    },
    del({id}){
        return $axios.delete(`${users}/${id}`)
    },
    AdminLogin(params){
        return $axios.post(auth + '/login', params)
    }
}


export default user


