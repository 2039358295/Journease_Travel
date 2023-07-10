import axios from 'axios'

export function updateUser(data){
    return axios.postForm(`user/update`,data)
}

export function GetTicketlist(userid){
    return axios.post('scenery/getorders',JSON.stringify({
        userid:   userid,
    }))
}

export function getInfo(data){
    return axios({
        method:'post',
        url:`user/getInfo`,
        data
    })
}

export function getCollectSpot(data){
    return axios({
        method:'post',
        url:`user/getCollection`,
        data
    })
}