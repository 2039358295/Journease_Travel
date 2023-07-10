import {CreatePost} from "./PostAPI.js";

export function showPassword(){
    if(document.getElementById('Password').type === "text"){
        document.getElementById('Password').type = "password";
    }else{
        document.getElementById('Password').type = "text";
    }
}


export function onFileChange(event, post) {     //选择图片
    post.Photo = event.target.files[0].name
}

export function onSubmit() {        //提交新建帖子
    this.visible=false;
    this.formData.tag1 = this.formData.tag1[0] + "/" + this.formData.tag1[1]
    console.log(this.formData)
    CreatePost(this.formData)
        .then(response => {
            console.log(response.data);
            alert("发表成功")
            // window.location.reload()
        })
        .catch(error => {
            console.log(error);
        });
    location.reload()
}


