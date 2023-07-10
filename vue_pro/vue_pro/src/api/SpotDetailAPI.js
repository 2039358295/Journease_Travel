import axios from 'axios'

export function GetComments(userid,sceneryid){
    return axios.post('scenery/getcomments',JSON.stringify({
        userid:     userid,
        sceneryid:  sceneryid,
    }))
}

export function GetScenery(userid,sceneryid){
    return axios.post('scenery/getscenery', JSON.stringify({
        userid:     userid,
        sceneryid:  sceneryid,
    }))
}

export function GetSpotpic(userid,sceneryid){
    console.log(userid.type)
    return axios.post('scenery/getspotpic',JSON.stringify({
            userid:     userid,
            sceneryid:  sceneryid,
        })
    )
}

export function PostComment(comment){
    return axios.post('scenery/postcomment', JSON.stringify({
            avatar:   comment.avatar,
            name:     comment.nickname,
            userid:     comment.uid,
            time:     comment.time,
            comment:  comment.comment,
            rating:   comment.rating,
            likes:    comment.likes,
        })
    )
}

export function GetNickname(userid,sceneryid){
    return axios.post('scenery/getnickname', JSON.stringify({
            userid:     userid,
            sceneryid:  sceneryid,
        })
    )
}

export function PostLikes(comment){
    return axios.post('scenery/postlikes',JSON.stringify({
        id:       comment.id,
        likes:    comment.likes,
    }))
}

export function PostOrders(type,date,num,sname,price){
    return axios.post('scenery/postorders',JSON.stringify({
        type:     type,
        date:     date,
        num:      num,
        sceneryname: sname,
        price: price,
    }))
}