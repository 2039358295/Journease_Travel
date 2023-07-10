import axios from "axios";

export function DeleteUserFP(userid,postid){
    return axios.delete(`/community/DeleteFP`,{
        data:JSON.stringify({
            userID:userid,
            postID:postid,
        })
    })
}


export function deletePost(postID){         //删除某个帖子
    axios.delete(`/community/delete/${postID}`)
        .then(response => {
            console.log(response)
            if (response.status===200) {
                alert('删除成功！');
                window.location.reload()
            } else {
                alert('删除失败！');
            }
        })
        .catch(error => {
            console.log(error);
        });
}