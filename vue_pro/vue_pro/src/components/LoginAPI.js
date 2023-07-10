import axios from 'axios'

export function RegisterJudge(data){
    return axios({
        method:'post',
        url:`user/register`,
        data
    })
}

export function LoginJudge(data){
    return axios({
        method:'post',
        url:`user/login`,
        data
    })
}