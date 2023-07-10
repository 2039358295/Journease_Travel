<template>
  <div class="详细页">
    <!-- 轮播图片 -->
    <!-- 轮播图片需要数组：spotpic[photo score ] -->
    <div class="spot-detail">
      <!-- :interval="5000" -->
      <el-carousel :interval="5000" :height="carouselHeight" style="padding: 30px">
        <el-carousel-item style="text-align: center;" v-for="(item, index) in spotpic" :key="index">
          <img :src="item.src1" alt="spot image" style="height:450px;width: 1250px;max-height: 100%; max-width: 100%" />
          <div class="carousel-caption">
            <!-- <h3>{{ item.caption }}</h3> -->
            <!-- <div class="rating">
              <el-rate v-model="item.score" disabled :show-score="true" />
            </div> -->
            <p class="detailcomment">♡{{ item.comment }}♡</p>
          </div>
        </el-carousel-item>
      </el-carousel>

    </div>
    <!-- 景点描述 -->
    <div style="margin-bottom: 50px">
      <div class="景区名称">
        <h3 style="text-align: center">{{ spotname }}</h3>
      </div>
      <div class="地址" style="text-align: center;">
        <p>{{ spotaddress }}</p>
      </div>
      <div class="正文">
        <p style="margin-right: 60px;margin-left: 60px;">{{ spotint }}</p>
      </div>
    </div>
    <!-- 门票 -->
    <div class="container"  style="text-align: center">
      <el-table :data="ticketTypes" style="width: 100%" class="table-container">
        <el-table-column prop="type" label="票型" class="table-column"></el-table-column>
        <el-table-column prop="price" label="价格（元）" class="table-column"></el-table-column>
        <el-table-column label="预订" class="table-column">
          <template #default="{ row }">
            <el-button @click="showDialog(row)"  class="button tada">预订</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <el-dialog v-model="dialogVisible" style="display:flex;justify-content: center;width: 450px;height: 250px">
      <h2 style="margin-bottom: 10px">{{ selectedTicketType.type }} - ￥{{ selectedTicketType.price }}</h2>
      <div style="margin-bottom: 10px">
        <span>游玩时间：</span>
        <el-date-picker
            v-model="selectedDate"
            type="date"
            placeholder="选择日期"
            style="width: 200px"
        ></el-date-picker>
      </div>
      <div style="margin-bottom: 10px">
        <span>购买数量：</span>
        <el-input-number v-model="selectedQuantity" :min="1" :max="10"></el-input-number>
      </div>
      <el-button @click="buyTicket">购买</el-button>
    </el-dialog>
    <!-- </div> -->
    <!-- 评论展示 -->
    <!-- 数组名称：commentList[avatar[头像]、name[昵称]、time[时间]、comment[评论内容]、likes[点赞数]] -->
    <div class="comments">
      <!-- 发表评论输入框 -->
      <div class="comment-input">
        <el-input v-model="comment" placeholder="请输入评论内容"></el-input>
        <el-rate v-model="rate" show-score text-color="#ff9900" :colors="['#99A9BF', '#F7BA2A', '#FF9900']"></el-rate>
        <el-button type="primary" @click="submitComment">提交</el-button>
      </div>
      <!-- 展示评论列表 -->
      <div class="comment-list">
        <div class="comment" v-for="(item, index) in commentList" :key="index">
          <div class="comment-info">
            <img :src="item.avatar" alt="" class="avatar">
            <div class="info">
              <div class="nickname">{{ item.name }}</div>
              <div class="time">{{ item.time }}</div>
            </div>
          </div>
          <div class="content">{{ item.comment }}</div>
          <!--        <div class="rating">-->
          <!--          <el-rate v-model="item.rating" disabled show-score text-color="#ff9900" :colors="['#99A9BF', '#F7BA2A', '#FF9900']"></el-rate>-->
          <!--          <div>-->
          <!--            <img class="likeComment"  v-bind:style="{backgroundColor: item.isliked ? '#f85b06' : 'white'}" style="z-index: 1;height: 20px;width: 23px;" size="middle" circle @click="likeComment(index)" :src="require('@/assets/likecomment.png')" >-->
          <!--            {{ item.likes }}-->
          <!--          </div>-->
          <!--        </div>-->
          <div class="rating">
            <el-rate v-model="item.rating" disabled show-score text-color="#ff9900" :colors="['#99A9BF', '#F7BA2A', '#FF9900']"></el-rate>
            <!--          <el-button type="text" icon="el-icon-thumb-up" @click="likeComment(index)">{{ item.likes }}</el-button>-->
            <!--          <img class="likeComment"  v-bind:style="{backgroundColor: item.isliked ? '#f85b06' : 'white'}" style="z-index: 1;height: 20px;width: 23px;" size="middle" circle @click="likeComment(index)" :src="require('@/assets/likecomment.png')" >-->
          </div>
        </div>
      </div>
    </div>
    <el-affix position="bottom" :offset="100">
      <el-button class="like-style" style="z-index: 1; position: fixed; bottom: 108px; right: 52px;align-items: center" size="middle" v-bind:style="{backgroundColor: this.collection ? '#fab4eb' : 'white'}" id="likeButton" circle v-on:click="toggleLike()">
        <img :src="require('@/assets/heart.png')" style="height: 20px;width: 20px; z-index: 3;margin-left: 0">
      </el-button>
    </el-affix>
  </div>
  <el-backtop :right="50" :bottom="50" />
</template>
<script>

import { ElMessage, ElMessageBox } from 'element-plus'
//import {getToken} from '../components/token.js'
import {
  GetComments, GetNickname,
  GetScenery,
  GetSpotpic,
  PostComment,
  PostLikes,
  PostOrders
} from "../api/SpotDetailAPI";
import {PostAvatar, PostEmail} from "@/api/PostAPI";
import {getToken} from "@/components/token";
import {GetPoster} from "@/api/GetAPI";


export default{
  name:'SpotDetail',
  // props:{
  //   spotid,
  // },
  data(){
    return{
      uid:0,
      spotID:2,
      avatar:null,
      collection:'',
      nickname:'',
      spotname:null,
      spotint:null,
      spotaddress:null,
      ismaked:false,
      carouselHeight:'450px',
      // spotname:'黄山',
      // spotarticle:'五岳归来不看山，黄山归来不看岳',
      // spotaddress:'安徽省黄山市',
      spotpic:[
        //   {
        //     src1:require('../assets/logoname.jpg'),
        //     score:3.5,
        //     comment:'五岳归来不看山，黄山归来不看岳',
        //   },
        //   {
        //     src1:require('../components/2.jpg'),
        //     // src2:require('../components/3.png'),
        //     // caption:'黄山',
        //     score:5,
        //     comment:'五岳归来不看山，黄山归来不看岳',
        //   },
        //   {
        //     src1:require('../assets/logoname.jpg'),
        //     // src2:require('../components/3.png'),
        //     // caption:'黄山',
        //     score:5,
        //     comment:'五岳归来不看山，黄山归来不看岳',
        //   },
        //   {
        //     src1:require('../components/2.jpg'),
        //     // src2:require('../components/3.png'),
        //     // caption:'黄山',
        //     score:5,
        //     comment:'五岳归来不看山，黄山归来不看岳',
        //   },
        //   {
        //     src1:require('../components/2.jpg'),
        //     // src2:require('../components/3.png'),
        //     // caption:'黄山',
        //     score:5,
        //     comment:'五岳归来不看山，黄山归来不看岳',
        //   },
        //   {
        //     src1:require('../components/2.jpg'),
        //     // src2:require('../components/3.png'),
        //     // caption:'黄山',
        //     score:'4',
        //     comment:'五岳归来不看山，黄山归来不看岳',
        //   },
      ],
      // props:{

      // },
      ticketTypes: [
        { type: '成人票', price: null },
        { type: '儿童票', price: null },
        { type: '学生票', price: null },
      ],
      dialogVisible: false,
      selectedTicketType: null,
      selectedDate: null,
      selectedQuantity: 1,
      selectedSname: 'nidaya',
      comment: '',
      rate: 0,
      commentList: [
        {
          id: 0,
          userid:0,
          // avatar: 'https://picsum.photos/50/50',
          name: '小明',
          time: '2023-05-28 14:30:00',
          comment: '这个景区真的很不错，值得一去！',
          rating: 4,
          likes: 10,
        },
      ],
    }
  },
  methods: {
    showDialog(row) {
      if(this.uid){
        this.selectedTicketType = row
        this.dialogVisible = true
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
    },
    buyTicket() {
      if (!this.selectedDate || !this.selectedQuantity) {
        this.$message.error('请选择游玩时间和购买数量')
        return
      }
      PostOrders(this.selectedTicketType.type,this.selectedDate,this.selectedQuantity,this.selectedSname,this.selectedTicketType.price)
      const totalPrice = this.selectedTicketType.price * this.selectedQuantity
      ElMessage.success(`您需要支付${totalPrice}元`)
      this.dialogVisible = false
    },
    submitComment() {
      if(this.uid){
        if (!this.comment) {
          ElMessage.error('请输入评论内容');
          return;
        }
        const comment = {
          //ldy:this.uid,
          // avatar: this.avatar,
          nickname: this.nickname,
          time: new Date().toLocaleString(),
          comment: this.comment,
          uid: this.uid,
          rating: this.rate,
          likes: 0,
        };

        PostComment(comment)
            .then(data => {
              console.log(data)
              // this.commentList.unshift(comment);
              this.comment = '';
              this.rate = 0;
              ElMessage.success('评论成功');
              window.location.reload();
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
    },
    likeComment(index) {
      PostLikes(this.commentList[index]);
      this.commentList[index].isliked = !this.commentList[index].isliked;
      if (this.commentList[index].isliked) {
        this.commentList[index].likes++;
      } else {
        this.commentList[index].likes--;
      }
    },
  },
  created() {
    PostEmail(getToken())
        .then(response=>{
          this.uid = response.data.uid
        })
    let Spot = JSON.parse(sessionStorage.getItem('spotID'))
    if (Spot) {
      this.spotID = Spot.spot
      this.collection=Spot.collection
      this.selectedSname=Spot.sname
    }
    console.log(this.spotID)
    // this.uid=getToken();

    // if(this.userID){
    //   GetAllPosts()
    //   .then(response => {
    //     console.log(response.data)
    //     this.commentList = response.data
    //   })
    //   .catch(error => {
    //     console.log(error);
    //   })
    // }
    PostEmail(getToken())
        .then(response=>{
          this.uid = response.data.uid
          GetComments(this.uid,this.spotID)
              .then(response => {
                console.log(response)
                this.commentList = response.data.commentList
                this.commentList.forEach(comment=>{
                    GetPoster(comment.userid)
                        .then(response=>{
                          comment.name = response.data.name
                        })
                    PostAvatar(comment.userid)
                        .then(response=>{
                          console.log(response)
                          comment.avatar = URL.createObjectURL(new Blob([response.data],{type:"image/jpeg"}))
                        })
                })
              })
              .catch(error => {
                console.log(error)
              })
        })
    PostEmail(getToken())
        .then(response=>{
          this.uid = response.data.uid
          GetScenery(this.uid,this.spotID)
              .then(response => {
                console.log(response)
                this.ticketTypes[0].price = response.data.price1
                this.ticketTypes[1].price = response.data.price2
                this.ticketTypes[2].price = response.data.price3
                this.spotname = response.data.name
                this.spotint = response.data.introduction
                this.spotaddress = response.data.address
              })
        })
    PostEmail(getToken())
        .then(response=>{
          this.uid = response.data.uid
          GetSpotpic(this.uid,this.spotID)
              .then(response => {
                console.log(response)
                this.spotpic = response.data.spotpic
              })
        })

    PostEmail(getToken())
        .then(response=>{
          this.uid = response.data.uid
          GetNickname(this.uid,this.spotID)
              .then(response => {
                this.nickname = response.data.nickname
              })
          PostAvatar(this.uid)
              .then(response=>{
                console.log(response)
                this.avatar = URL.createObjectURL(new Blob([response.data],{type:"image/jpeg"}))
              })
        })

  }
}
</script>
<style scoped>
.详细页{
  margin-top: 58px;
  background:linear-gradient(to top, #f8f5f4 0%, #ace0f9 100%);
  /*background-image: linear-gradient(-225deg, #FFFEFF 0%, #D7FFFE 100%);*/
  background-size: cover;
  position: relative;
  background-attachment: fixed;
}
/* 评论渲染 */
.comments {
  max-width: 600px;
  margin: 0 auto;
  padding: 50px;
  box-sizing: border-box;
}
.comment-input {
  display: flex;
  align-items: center;
  margin-top: 35px;
  margin-bottom: 20px;
}
.comment-input .el-input {
  flex: 1;
  margin-right: 20px;
}
.comment-input .el-rate {
  margin-right: 20px;
}
.likeComment{
  cursor:pointer;
}
.comment-list .comment {
  display: flex;
  flex-direction: column;
  margin-bottom: 20px;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 4px;
  background: white;
  opacity: 1;
}
.comment-list .comment .comment-info {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}
.comment-list .comment .comment-info .avatar {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  margin-right: 10px;
}
.comment-list .comment .comment-info .info {
  flex: 1;
}
.comment-list .comment .comment-info .nickname {
  font-size: 16px;
  font-weight: bold;
  margin-bottom: 5px;
}
.comment-list .comment .comment-info .time {
  font-size: 14px;
  color: #999;
}
.comment-list .comment .content {
  margin-bottom: 10px;
  font-size: 14px;
  line-height: 1.5;
}
.comment-list .comment .rating {
  display: flex;
  align-items: center;
}
.comment-list .comment .rating .el-rate {
  margin-right: 10px;
}
.comment-list .comment .rating .el-button {
  font-size: 14px;
  color: #999;
  padding: 0;
  margin-left: 10px;
}
/* 门票渲染 */
.container {
  /* max-width: 800px; */
  width: 500px;
  margin: 0px auto;
  margin-top: 20px;
  text-align: center;
  position: relative;
  opacity: 0.8;
}

.el-table {
  margin-top: 20px;
}

.el-table__header th {
  font-weight: bold;
  background-color: #f5f5f5;
}

.el-table__body td {
  text-align: center;
}

.el-dialog {
  text-align: center;
}

.el-dialog__body {
  padding: 20px;
}

.el-date-picker {
  width: 100%;
  margin-top: 10px;
}

.el-input-number {
  margin-top: 10px;
}

/* .el-button {
  margin-top: 20px;
  background-color: #409eff;
  color: #fff;
  border-color: #409eff;
}

.el-button:hover {
  background-color: #66b1ff;
  color: #fff;
  border-color: #66b1ff;
} */

/* 轮播渲染 */
/* .spot-detail {
position: relative;
background-size: cover;
} */
/* 设置轮播图片容器的最大高度和宽度 */
/* .spot-detail {
  max-height: 100%;
  max-width: 100%;
  width: 100%;
  height: 100%;
  overflow: hidden;
  position: relative;
} */
/* 设置轮播图片的最大高度和宽度 */
/* .el-carousel-item img {
  max-height: 100%;
  max-width: 100%;
  display: block;
  margin: 0 auto;
  width: 100%;
  height: 100%;
  object-fit: cover;
} */


.carousel-caption {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  height: 80px;
  background-color: rgba(43, 41, 41, 0.7);
  color: #fff;
  /* padding: 5px; */
}
.detailcomment{
  //font-family: 云峰飞云体;
  font-family: 云峰飞云体,serif;
  font-size: 20px;
}


.景区名称 h3 {
  font-family: 华文新魏, sans-serif;
  font-size: 36px; /* 设置标题字体大小 */
  font-weight: bold; /* 设置标题字体粗细 */
  color: #333; /* 设置标题字体颜色 */
  letter-spacing: 2px; /* 设置字母之间的间距 */
}

.地址 p {
  font-family: 华文行楷;
  font-size: 23px; /* 设置地址字体大小 */
  color: #666; /* 设置地址字体颜色 */
  line-height: 1.2; /* 设置地址行高 */
}

.正文 p {
  font-family: 楷体;
  font-size: 20px; /* 设置正文字体大小 */
  line-height: 1.6; /* 设置正文行高 */
  color: #1c1c1c; /* 设置正文字体颜色 */
}
.like-style {
  display: inline-block;
  margin-left: auto;
  /*padding: 1em 2em;*/
  background-color: rgba(252, 3, 2, 0.72);
  color: #fff;
  /*border-radius: 4px;*/
  border: none;
  cursor: pointer;
  position: relative;
  box-shadow: 0 2px 25px rgba(233, 30, 99, 0.5);
  /*outline: 0;*/
  transition: transform ease-in 0.1s, background-color ease-in 0.1s, box-shadow ease-in 0.25s;
}

.like-style::before {
  position: absolute;
  content: '';
  left: -2em;
  right: -2em;
  top: -2em;
  bottom: -2em;
  pointer-events: none;
  background-repeat: no-repeat;
  background-image: radial-gradient(circle, #ff002f 20%, transparent 20%),
  radial-gradient(circle, #ff0081 20%, transparent 20%),
  radial-gradient(circle, #ff0081 20%, transparent 20%),
  radial-gradient(circle, #ff0081 20%, transparent 20%),
  radial-gradient(circle, #ff0081 20%, transparent 20%),
  radial-gradient(circle, #ff0081 20%, transparent 20%),
  radial-gradient(circle, #ff0081 20%, transparent 20%),
  radial-gradient(circle, #ff0081 20%, transparent 20%),
  radial-gradient(circle, #ff0081 20%, transparent 20%),
  radial-gradient(circle, #ff0081 20%, transparent 20%),
  radial-gradient(circle, #ff0081 20%, transparent 20%),
  radial-gradient(circle, #ff0081 20%, transparent 20%),
  radial-gradient(circle, #ff0081 20%, transparent 20%),
  radial-gradient(circle, #ff0081 20%, transparent 20%),
  radial-gradient(circle, #ff0081 20%, transparent 20%),
  radial-gradient(circle, #ff0081 20%, transparent 20%);
  background-position: 5% 44%, -5% 20%, 7% 5%, 23% 0%, 37% 0, 58% -2%, 80% 0%, 100% -2%, -5% 80%,
  100% 55%, 2% 100%, 23% 100%, 42% 100%, 60% 95%, 70% 96%, 100% 100%;
  background-size: 0% 0%;
  transition: background-position .5s ease-in-out, background-size .75s ease-in-out;
}

.like-style:active::before {
  transition:0s;
  background-size: 10% 10%, 20% 20%, 15% 15%, 20% 20%, 18% 18%, 10% 10%, 15% 15%, 10% 10%, 18% 18%,
  15% 15%, 20% 20%, 18% 18%, 20% 20%, 15% 15%, 10% 10%, 20% 20%;
  background-position: 18% 40%, 20% 31%, 30% 30%, 40% 30%, 50% 30%, 57% 30%, 65% 30%, 80% 32%, 15% 60%,
  83% 60%, 18% 70%, 25% 70%, 41% 70%, 50% 70%, 64% 70%, 80% 71%;
}
.table-container {
  display: flex;
  justify-content: center;
  border: 1px solid #ccc;
  border-radius: 10px;
  padding: 10px;
  background-color: #a1c4fd;
}
.table-column{
  background-color: #a1c4fd;
}

html,body{
  margin: 0;
  height: 100%;
  display: grid;
  place-content: center;
}
:root{
  --primary-color: royalblue;
}

.button{
  padding: 5px 16px;
  color: #000000d9;
  border: 1px solid #d9d9d9;
  line-height: 1.4;
  box-shadow: 0 2px #00000004;
  cursor: pointer;
  transform: scale(1);
  animation: jump 0s;
  transition: .3s;
  background-color: #abbcee; /* 设置按钮的背景颜色 */
  border-radius: 5px; /* 设置按钮的圆角 */
}
.button:focus-visible{
  outline: 0;
}
.button[type="primary"]{
  border-color: transparent;
  background-color: var(--primary-color);
  color: #fff;
  text-shadow: 0 -1px 0 rgb(0 0 0 / 12%);
  box-shadow: 0 2px #0000000b;
}
.button[type="dashed"]{
  border-style: dashed;
}
.button[type="text"]{
  border-color: transparent;
  box-shadow: none;
}
.button:not([type]):hover,
.button[type="dashed"]:hover,
.button[type="dashed"]:focus{
  color: var(--primary-color);
  border-color: currentColor;
}
.button[type="primary"]:hover{
  filter: brightness(1.1);
}
.button[type="text"]:hover,
.button[type="text"]:focus{
  background: rgba(0,0,0,.018);
}
.tada{
  animation: tada 0s;
}
.rubberBand{
  animation: rubberBand 0s;
}
.jello{
  animation: jello 0s;
}
body:hover .button{
  animation-duration: 1s;
}
.button.button:active{
  filter: brightness(.9);
  animation: none;
}
@keyframes tada {
  from {
    transform: scale3d(1, 1, 1)
  }

  10%, 20% {
    transform: scale3d(.9, .9, .9) rotate3d(0, 0, 1, -3deg)
  }

  30%, 50%, 70%, 90% {
    transform: scale3d(1.1, 1.1, 1.1) rotate3d(0, 0, 1, 3deg)
  }

  40%, 60%, 80% {
    transform: scale3d(1.1, 1.1, 1.1) rotate3d(0, 0, 1, -3deg)
  }

  to {
    transform: scale3d(1, 1, 1)
  }
}

@keyframes rubberBand {
  0% {
    -webkit-transform: scaleX(1);
    transform: scaleX(1)
  }

  30% {
    -webkit-transform: scale3d(1.25,.75,1);
    transform: scale3d(1.25,.75,1)
  }

  40% {
    -webkit-transform: scale3d(.75,1.25,1);
    transform: scale3d(.75,1.25,1)
  }

  50% {
    -webkit-transform: scale3d(1.15,.85,1);
    transform: scale3d(1.15,.85,1)
  }

  65% {
    -webkit-transform: scale3d(.95,1.05,1);
    transform: scale3d(.95,1.05,1)
  }

  75% {
    -webkit-transform: scale3d(1.05,.95,1);
    transform: scale3d(1.05,.95,1)
  }

  to {
    -webkit-transform: scaleX(1);
    transform: scaleX(1)
  }
}
@keyframes jello {
  0%,11.1%,to {
    -webkit-transform: translateZ(0);
    transform: translateZ(0)
  }

  22.2% {
    -webkit-transform: skewX(-12.5deg) skewY(-12.5deg);
    transform: skewX(-12.5deg) skewY(-12.5deg)
  }

  33.3% {
    -webkit-transform: skewX(6.25deg) skewY(6.25deg);
    transform: skewX(6.25deg) skewY(6.25deg)
  }

  44.4% {
    -webkit-transform: skewX(-3.125deg) skewY(-3.125deg);
    transform: skewX(-3.125deg) skewY(-3.125deg)
  }

  55.5% {
    -webkit-transform: skewX(1.5625deg) skewY(1.5625deg);
    transform: skewX(1.5625deg) skewY(1.5625deg)
  }

  66.6% {
    -webkit-transform: skewX(-.78125deg) skewY(-.78125deg);
    transform: skewX(-.78125deg) skewY(-.78125deg)
  }

  77.7% {
    -webkit-transform: skewX(.390625deg) skewY(.390625deg);
    transform: skewX(.390625deg) skewY(.390625deg)
  }

  88.8% {
    -webkit-transform: skewX(-.1953125deg) skewY(-.1953125deg);
    transform: skewX(-.1953125deg) skewY(-.1953125deg)
  }
}

</style>
