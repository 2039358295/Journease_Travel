import axios from "axios";


export function GetPosts(){
        return axios({
            method:'get',
            url:`/community/posts`,
        })
}//获取社区所有帖子(点赞数)

export function GetAllPosts(){
    return axios({
        method:'get',
        url:`/community/posts/time`,
    })
}//获取社区所有帖子(点赞数)

export function GetPoster(UserID){
    return axios({
        method:'get',
        url:`/user/username/`+UserID,
    })
}

export function GetComments(postID){
    return axios({
        method:'get',
        url:`/community/`+postID
    })
}

export function GetUserFP(userid){
    return axios({
        method:'get',
        url:'community/UserFP/'+userid,
    })
}//获取用户的帖子收藏

export function GetMyPostList(UserID){
    return axios({
        method:'get',
        url:`/community/MyPost/`+UserID,
    })
}//查找我的贴子

export function GetFPpost(postID){
    return axios({
        method:'get',
        url:'/community/FPpost/'+postID,
    })
}

export function SearchPost(tag){
    return axios({
        method:'get',
        url:'/community/SearchPost',
        params: {
            tag: tag
        }
    })
}
