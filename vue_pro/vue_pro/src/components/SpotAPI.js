import axios from 'axios'

export function getNavData(data){
    return axios({
        method:'post',
        url:`user/resource`,
        data
    })
}

export function getSpotData(data){
    return axios({
        method:'post',
        url:`spot/resource`,
        data
    })
}

export function getSpotByScore(data){
    return axios({
        method:'post',
        url:`spot/OrderByScore`,
        data
    })
}

export function FilterSpotByAddress(data){
    return axios({
        method:'post',
        url:`spot/Filter1`,
        data
    })
}

export function FilterSpotByName(data){
    return axios({
        method:'post',
        url:`spot/FilterByName`,
        data
    })
}

export function SearchSpotByName(data){
    return axios({
        method:'post',
        url:`spot/SearchByName`,
        data
    })
}

export function CreateSpot1(data){
    return axios({
        method:'post',
        url:`spot1/create`,
        data
    })
}

export function DeleteSpot1(data){
    return axios({
        method:'post',
        url:`spot1/delete`,
        data
    })
}

export function Search(data){
    return axios({
        method:'post',
        url:`spot/search`,
        data
    })
}

export function Display(data){
    return axios({
        method:'post',
        url:`spot/display`,
        data
    })
}