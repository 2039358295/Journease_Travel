import axios from "axios";


export function AddLike(PostID, CommentID){         //评论点赞
    return axios({
        method:'put',
        url:'/community/Like',
        data:JSON.stringify({
            PostID:PostID,
            CommentID: CommentID,
        })
    })
}

export function CancelLike(PostID, CommentID){          //取消评论点赞
    return axios({
        method:'put',
        url:'/community/Dislike',
        data:JSON.stringify({
            PostID:PostID,
            CommentID: CommentID,
        })
    })
}

export function DislikePost(postid){
    return axios.put("/community/DislikePost",JSON.stringify({
            PostID: parseInt(postid)
        })
    )
}

export function LikePost(postid){
    return axios.put("/community/LikePost",JSON.stringify({
            PostID: parseInt(postid)
        })
    )
}