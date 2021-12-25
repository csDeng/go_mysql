import $axios from '~/axios'

const auth = '/auths'
const users = '/api/v1/users'
const user = {
    login(params){
        return $axios.post('/login',params)
    },
    register(params){
        return $axios.post('/register',params)
    },
    get_list(pid=0){
        return $axios.get(`${users}?page=${pid}`)
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


