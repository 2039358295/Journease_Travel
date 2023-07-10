import axios from "axios";
import Cookies from "js-cookie"
import {ElMessageBox} from "element-plus";
import {getToken} from "@/components/token";

export function PostImg(postID){
    return axios.get(`/community/img/`+postID,{
        responseType:'blob'
    })
}

export function PostImg1(postID){
    return axios.get(`/community/img1/`+postID,{
        responseType:'blob'
    })
}

export function PostImg2(postID){
    return axios.get(`/community/img2/`+postID,{
        responseType:'blob'
    })
}

export function PostImg3(postID){
    return axios.get(`/community/img3/`+postID,{
        responseType:'blob'
    })
}

export function PostImg4(postID){
    return axios.get(`/community/img4/`+postID,{
        responseType:'blob'
    })
}

export function PostImg5(postID){
    return axios.get(`/community/img5/`+postID,{
        responseType:'blob'
    })
}

export function PostImg6(postID){
    return axios.get(`/community/img6/`+postID,{
        responseType:'blob'
    })
}

export function PostImg7(postID){
    return axios.get(`/community/img7/`+postID,{
        responseType:'blob'
    })
}

export function PostImg8(postID){
    return axios.get(`/community/img8/`+postID,{
        responseType:'blob'
    })
}

export function PostEmail(data){
    return axios({
        method:'post',
        url:`user/uid`,
        data:JSON.stringify(data)
    })
}

export function AddUserFP(userid,postid){
    return axios.post('community/AddFP',JSON.stringify({
        userID:userid,
        postID:postid,
    }))
}
export function CreatePost(data){
    console.log(data)
    return axios.postForm('community/post', data);
}//发表帖子

export function SubmitComment(uid){
    if(getToken()){
        let userID
        PostEmail(uid)
            .then(response=>{
                userID = response.data.uid
                return axios.post(`/community/create`,JSON.stringify({
                    userID:userID,
                    commentID:0,
                    postID:JSON.parse(Cookies.get("postID")),
                    content:this.commentForm.comment,
                    time:""
                }))
                    .then((response) => {
                        console.log(response)
                        if (response.status===200){
                            alert("发表成功！");
                            window.location.reload()
                        }else{
                            alert("发表失败！");
                            window.location.reload()
                        }
                    })
                    .then((data) => console.log(data))
                    .catch((error) => console.error(error))
            })
    }
    else{
        ElMessageBox.confirm('请先登录', '提示', {
            confirmButtonText: '去登录',
            cancelButtonText: '取消',
            type: 'warning'
        }).then(() => {
            // 用户点击了确定按钮，跳转至登录页面
            this.$router.push('/login');
        }).catch(() => {
            // 用户点击了取消按钮，不做操作
        });
    }
}

export function PostAvatar(userID){
    return axios.get(`/user/avatar/`+userID,{
        responseType:'blob'
    })
}

