<template>
  <div id="allIden">

    <el-container>
      <el-container class="idenAll">
        <el-aside width="200px" style="margin-left: 50px;margin-top: 60px;margin-bottom: 180px;" >
          <el-menu style="border-radius: 10px;"  v-model="activeIndex" :default-active="activeIndex" class="el-menu-vertical-demo" @select="handleSelect" position:fixed>
            <el-menu-item style="border-radius: 10px;" index="1" >
              <el-icon>
                <edit />
              </el-icon>
              修改资料</el-menu-item>
            <el-menu-item index="2">
              <el-icon>
                <message />
              </el-icon>

              我的帖子</el-menu-item>
            <el-menu-item index="3">
              <el-icon>
                <star />
              </el-icon>
              我的帖子收藏</el-menu-item>
            <el-menu-item index="4">
              <el-icon>
                <star />
              </el-icon>
              我的景点收藏</el-menu-item>
            <el-menu-item style="border-radius: 10px;" index="5">
              <el-icon>
                <tickets />
              </el-icon>
              我的订单</el-menu-item>
          </el-menu>
        </el-aside>

        <el-main >
          <div v-if="activeIndex === '1'" class="custom-label">

            <el-form :model="form" label-width="80px" style="margin-top: 30px">
              <el-form-item label="头像" class="avatar-wrap">
                <el-avatar :src="form.avatar" size="large"></el-avatar>
                <div  @click="visible = true">
                  <div style="display: flex;">
                    <el-tooltip
                        class="item"
                        effect="light"
                        content="上传头像"
                        placement="top"
                    >
                      <el-icon  size="40" round style="text-align: center; margin-left: 20px; cursor: pointer;" @mouseenter="showlable = true" @mouseleave="showlable = false"><Camera /></el-icon>
                    </el-tooltip>

                    <!-- <h3 v-show="showlable" style="font-family: 宋体;margin-top: 20px;">上传头像</h3> -->
                  </div>
                </div>
              </el-form-item>

              <el-form-item label="昵称">
                <el-input style="width: 200px;" v-model="form.nickname" type="text"></el-input>
              </el-form-item>
              <el-form-item label="邮箱">
                <el-input style="width: 200px;" v-model="form.email" type="email"></el-input>
              </el-form-item>
              <el-form-item label="密码">
                <el-input style="width: 200px;" v-model="form.password" :type="showPassword ? 'text' : 'password'"></el-input>
                <el-button type="text"  @click="showPassword = !showPassword">
                  <el-icon>
                    <View />
                  </el-icon>
                </el-button>
                <!-- <el-input style="width: 200px;" v-model="form.password" type="password"></el-input> -->
              </el-form-item>
              <el-form-item>
                <el-button id="confirmButton" @click="submitForm">确认修改</el-button>
              </el-form-item>
            </el-form>
          </div>



          <!-- 上传头像 -->
          <el-dialog v-model="visible" title="上传头像" style=" width: 400px;display: block;text-align: center;">
            <div>
              <el-upload
                  class="avatar-uploader"
                  :show-file-list="showFileList"
                  limit="1"
                  action="http://127.0.0.1:8888/api/private/v1/upload"
                  list-type="picture-card"
                  :on-preview="handlePictureCardPreview"
                  :before-upload="beforeAvatarUpload"
                  :on-remove="handleRemove"
                  :on-success="handleSuccess"
                  :on-change="handlePhoto"
                  :auto-upload="false"
              >
                <img v-if="imageUrl && showFileList" :src="imageUrl" class="avatar">
                <el-icon><Plus /></el-icon>
              </el-upload>
            </div>
<!--            <div style="text-align: center;margin-top:10px;">-->

<!--              <el-button type="button" @click="confirm" >确定上传</el-button>-->
<!--            </div>-->
          </el-dialog>


          <div v-if="activeIndex === '2'" class="default-info">

            <el-space wrap>
              <el-card class="板块信息" style="width: 800px"  v-for = "(item, index) in mypostlist" :key="index">
                <div >
                  <div class="帖子头部">
                    <div class="头像" style="margin-right: 6px;">
                      <!-- <el-avatar :src="item.avatar" :size="40" fit="cover"></el-avatar> -->
                      <img :src="item.avatar" fit="cover" alt='' size="40" style="width: 100%;height: 100%;">
                    </div>
                    <div class="comment-info">
                      <div class="comment-username">{{item.UserID }}</div>
                      <div class="comment-time">{{item.RTime}}</div>
                    </div>
                  </div>
                  <h3 class="帖子名称">{{ item.Title }}</h3>
                  <!-- 根据expand的值，动态计算文章内容的行数 -->
                  <div id="ar" class="帖子内容" :style="{maxHeight: item.expand ? 'none' : '60px', overflow: 'hidden'}">{{item.Contents}}</div>
                  <el-button type="text" v-if="item.Contents.length > 60 && !item.expand" @click="expand(index)">展开全文</el-button>
                  <el-button type="text" v-if="item.expand" @click="noexpand(index)">收起</el-button>
                  <!-- 帖子图片 -->
                  <el-scrollbar style="padding: 5px;margin-top: 7px;">
                    <div class="scrollbar-flex-content">
                      <p v-for="(p, index) in item.pic" :key="index" class="scrollbar-demo-item">
                        <img :src="p" alt="spot image" style="height:450px;width: 1250px;max-height: 100%; max-width: 100%" />
                      </p>
                    </div>
                  </el-scrollbar>
                  <!-- 点赞收藏 -->
                  <div class="点赞收藏评论标签" style="position: relative;margin-top: 4px;">

                    <div class="标签" >
                      <el-tag class="tag">{{ item.Tag1 }}</el-tag>
                      <el-tag class="tag">{{ item.Tag2 }}</el-tag>
                    </div>
                    <div class="点赞收藏评论" style="display: flex;" >
                      <!-- 点赞 -->
                      <div class="element">
                        <el-button class="like-style"  v-bind:style="{backgroundColor: item.isLiked ? '#fab4eb' : 'white'}" style="z-index: 1;" size="middle" id="likeButton" circle  v-on:click="toggleLike(item)" >
                          <img :src="require('@/assets/heart.png')" style="height: 20px;width: 23px; z-index: 3;margin-left: 0px;">
                        </el-button>
                      </div>

                      <div class="element" style="margin-top: 4px;">
                        {{ item.Likenum }}
                      </div>

                      <!-- <div class="圆角框"> -->
                      <div class="element">
                        <el-button class="star-style" size="middle" circle v-bind:style="{backgroundColor: item.isClicked ? '#f5e473e2' : 'white'}" v-on:click="toggleColor(item)" >
                          <img :src="require('@/assets/star.png')" style="height: 22px;width: 24px; z-index: 3;margin-top: -3px;">
                        </el-button>
                      </div>
                      <div class="element" style="margin-top: 4px;">
                        {{ item.Collection }}
                      </div>
                      <!-- </div> -->
                      <div class="element">
                        <el-tooltip
                            class="item"
                            effect="light"
                            content="查看评论"
                            placement="bottom"
                        >

                          <el-icon class="comment_button"  size="30" round style="text-align: right;margin-top: 2px;cursor: pointer;" @click="showCommentSidebar(index,item.PostID)"><ChatLineRound /></el-icon>
                        </el-tooltip>
                      </div>

                      <div class="element">
                        <el-tooltip
                            class="item"
                            effect="light"
                            content="删除帖子"
                            placement="bottom"
                        >
                          <el-icon   size="30" round style="text-align: right;margin-top: 2px;cursor: pointer;" @click="deletePost(item.PostID)">
                            <delete />
                          </el-icon>
                        </el-tooltip>


                      </div>

                    </div>
                  </div>
                  <!-- 评论 -->
                  <div class="comment-sidebar">

                    <el-drawer class="comment-header" title="评论列表" v-model="item.commentSidebarVisible" size="25%">
                      <div class="comment-form">
                        <!-- <div class="comment-input"> -->
<!--                        <div style="margin-top: 5px;">-->
<!--                          <el-avatar :src="currentUserAvatar" size="medium" style="margin-left: -18px;"></el-avatar>-->
<!--                        </div>-->
                        <el-form v-model="commentForm" label-width="0" >
                          <div class="comment-wrapper">
                            <el-input
                                type="textarea"
                                style="width: 320px;height: 100px;display: flex;margin-left: 10px; resize: none;"
                                v-model="commentForm.comment"
                                placeholder="请输入您的评论..."
                            ></el-input>
                            <el-icon class="comment-button" @click="SubmitComment(this.uid)" size="30" round style="text-align: center;cursor: pointer;"><Promotion /></el-icon>
                          </div>
                        </el-form>

                        <!-- </div> -->

                      </div>
                      <div class="comment-list-container">
                        <el-card v-for="comment in item.commentList" :key="comment.id" class="comment-card">
                          <div class="comment-header">
                            <el-avatar :src="comment.avatar" size="medium"></el-avatar>
                            <div class="comment-info">
                              <div class="comment-username">{{ comment.name }}</div>
                              <div class="comment-time">{{ comment.time }}</div>
                            </div>
                          </div>
                          <div class="comment-content">{{ comment.content }}</div>
                        </el-card>
                      </div>
                    </el-drawer>
                  </div>
                </div>
              </el-card>
            </el-space>
          </div>

          <div v-if="activeIndex === '3'" class="default-info">

            <el-space wrap>
              <el-card class="板块信息" style="width: 800px"  v-for = "(item, index) in FPpostlist" :key="index">
                <div >
                  <div class="帖子头部">
                    <div class="头像" style="margin-right: 6px;">
                      <!-- <el-avatar :src="item.avatar" :size="40" fit="cover"></el-avatar> -->
                      <img :src="item.avatar" fit="cover" alt='' size="40" style="width: 100%;height: 100%;">
                    </div>
                    <div class="comment-info">
                      <div class="comment-username">{{item.UserID }}</div>
                      <div class="comment-time">{{item.RTime}}</div>
                    </div>
                  </div>
                  <h3 class="帖子名称">{{ item.Title }}</h3>
                  <!-- 根据expand的值，动态计算文章内容的行数 -->
                  <div id="ar" class="帖子内容" :style="{maxHeight: item.expand ? 'none' : '60px', overflow: 'hidden'}">{{item.Contents}}</div>
                  <el-button type="text" v-if="item.Contents.length > 60 && !item.expand" @click="expand(index)">展开全文</el-button>
                  <el-button type="text" v-if="item.expand" @click="noexpand(index)">收起</el-button>
                  <!-- 帖子图片 -->
                  <el-scrollbar style="padding: 5px;margin-top: 7px;">
                    <div class="scrollbar-flex-content">
                      <p v-for="(p, index) in item.pic" :key="index" class="scrollbar-demo-item">
                        <img :src="p" alt="spot image" style="height:450px;width: 1250px;max-height: 100%; max-width: 100%" />
                      </p>
                    </div>
                  </el-scrollbar>
                  <!-- 点赞收藏 -->
                  <div class="点赞收藏评论标签" style="position: relative;margin-top: 4px;">

                    <div class="标签" >
                      <el-tag class="tag">{{ item.Tag1 }}</el-tag>
                      <el-tag class="tag">{{ item.Tag2 }}</el-tag>
                    </div>
                    <div class="点赞收藏评论" style="display: flex;" >
                      <!-- 点赞 -->
                      <div class="element">
                        <el-button class="like-style"  v-bind:style="{backgroundColor: item.isLiked ? '#fab4eb' : 'white'}" style="z-index: 1;" size="middle" id="likeButton" circle  v-on:click="toggleLike(item)" >
                          <img :src="require('@/assets/heart.png')" style="height: 20px;width: 23px; z-index: 3;margin-left: 0px;">
                        </el-button>
                      </div>

                      <div class="element" style="margin-top: 4px;">
                        {{ item.Likenum }}
                      </div>


                      <!-- <div class="圆角框"> -->
                      <div class="element">
                        <el-button class="star-style" size="middle" circle v-bind:style="{backgroundColor: item.isFP ? '#f5e473e2' : 'white'}" v-on:click="toggleColor(item)" >
                          <img :src="require('@/assets/star.png')" style="height: 22px;width: 24px; z-index: 3;margin-top: -3px;">
                        </el-button>
                      </div>
                      <div class="element" style="margin-top: 4px;">
                        {{ item.Collection }}
                      </div>
                      <div class="element">
                        <el-tooltip
                            class="item"
                            effect="light"
                            content="查看评论"
                            placement="bottom"
                        >

                          <el-icon class="comment_button"  size="30" round style="text-align: right;margin-top: 2px;cursor: pointer;" @click="showCommentSidebar1(index,item.PostID)"><ChatLineRound /></el-icon>
                        </el-tooltip>
                      </div>

                    </div>
                  </div>
                  <!-- 评论 -->
                  <div class="comment-sidebar">

                    <el-drawer class="comment-header" title="评论列表" v-model="item.commentSidebarVisible" size="25%">
                      <div class="comment-form">
                        <!-- <div class="comment-input"> -->
<!--                        <div style="margin-top: 5px;">-->
<!--                          <el-avatar :src="currentUserAvatar" size="medium" style="margin-left: -18px;"></el-avatar>-->
<!--                        </div>-->
                        <el-form v-model="commentForm" label-width="0" >
                          <div class="comment-wrapper">
                            <el-input
                                type="textarea"
                                style="width: 320px;height: 100px;display: flex;margin-left: 10px; resize: none;"
                                v-model="commentForm.comment"
                                placeholder="请输入您的评论..."
                            ></el-input>
                            <el-icon class="comment-button" @click="SubmitComment(this.uid)" size="30" round style="text-align: center;cursor: pointer;"><Promotion /></el-icon>
                          </div>
                        </el-form>


                      </div>
                      <div class="comment-list-container">
                        <el-card v-for="comment in item.commentList" :key="comment.id" class="comment-card">
                          <div class="comment-header">
                            <el-avatar :src="comment.avatar" size="medium"></el-avatar>
                            <div class="comment-info">
                              <div class="comment-username">{{ comment.name }}</div>
                              <div class="comment-time">{{ comment.time }}</div>
                            </div>
                          </div>
                          <div class="comment-content">{{ comment.content }}</div>
                        </el-card>
                      </div>
                    </el-drawer>
                  </div>
                </div>
              </el-card>
            </el-space>
          </div>

          <div v-if="activeIndex === '4'" class="default-info">
            <div class="spots" >
              <el-card :body-style="{ padding: '0px' }"  style="opacity: 0.85; height: 230px;width: 450px;margin: 10px;" v-for="(spot, index) in spots" :key="index" >
                <div style="display: flex;">
                  <div class="photo" style="display:inline-block;" @click="goToSpot(spot.id,spot.collection,spot.name)">
                    <img :src="spot.picture" style="height: 180px;width: 180px;margin-left: 10px;margin-top: 10px"/>
                  </div>
                  <div class="info" style="margin-left: 10px; flex-grow: 1;">
                    <div class="name" style="margin-bottom: 5px" @click="goToSpot(spot.id,spot.collection,spot.name)">{{ spot.name }}</div>
                    <div style="margin-bottom: 10px" @click="goToSpot(spot.id,spot.collection,spot.name)">
                      <span class="address">{{ spot.address }}</span>
                      <span class="address1">{{spot.address1}}</span>
                      <el-divider direction="vertical"></el-divider>
                      <span> <el-tag>{{ spot.type }}</el-tag></span>
                    </div>
                    <div v-truncate-text="40" class="text" style="margin-bottom: 20px" @click="goToSpot(spot.id,spot.collection,spot.name)">
                      {{spot.text}}
                    </div>
                    <div class="comment" style="margin-bottom: 30px" @click="goToSpot(spot.id,spot.collection,spot.name)">
                      “{{spot.comment}}”
                    </div>

                    <div class="rating" style="margin-bottom: 0px" @click="goToSpot(spot.id,spot.collection,spot.name)">
                      <el-rate v-model="spot.score" disabled :show-score="true" />
                    </div>
                    <div style="display: flex; justify-content:end; margin-right: 10px; margin-bottom: 0px;align-items: center;">
                      <el-button class="like-style"  size="middle" v-bind:style="{backgroundColor: spot.collection ? '#fab4eb' : 'white'}" style="z-index: 1;margin-right: 10px"  circle  v-on:click="toggleSpot(spot)" >
                        <img :src="require('@/assets/heart.png')" style="height: 20px;width: 23px; z-index: 3;">
                      </el-button>
                      {{spot.collections}}
                    </div>

                  </div>
                </div>
              </el-card>
            </div>
          </div>


          <div v-if="activeIndex === '5'" class="default-info">

            <div class="ticket-container" v-for="ticket in ticketlist" :key="ticket.id">
              <div class="image-container">
                <img :src="getImage(ticket.type)" alt="Example" class="ticket-image">
                <div class="text-overlay">
                  <h class="ticket-name">
                    {{ ticket.sceneryname }}
                  </h>
                  <p class="ticket-date">日期：{{ticket.date}}  人数：{{ ticket.num }}
                  </p>
                </div>
                <div class="ticket-price">
                  {{ ticket.price }}
                </div>
              </div>
            </div>
          </div>

        </el-main>
      </el-container>
    </el-container>
    <el-backtop :right="50" :bottom="50" />
  </div>
</template>


<script>
import {mapGetters} from "vuex";
// eslint-disable-next-line no-unused-vars
import { getSpotData } from '../components/SpotAPI.js'
import { ElMessage } from "element-plus";
import { CreateSpot1,DeleteSpot1} from '../components/SpotAPI.js'
import { ElMessageBox } from 'element-plus'
import {deletePost, DeleteUserFP} from "../api/DeleteAPI.js";
import { ref } from 'vue'
import {updateUser, getInfo, getCollectSpot, GetTicketlist} from '../components/IdentityAPI.js'
import {getToken} from "@/components/token";
import {AddUserFP, PostAvatar, PostEmail} from "@/api/PostAPI";
import {GetComments, GetFPpost, GetMyPostList, GetPoster, GetUserFP} from "@/api/GetAPI";
import {
  PostImg,
  PostImg1,
  PostImg2,
  PostImg3,
  PostImg4,
  PostImg5,
  PostImg6,
  PostImg7,
  PostImg8, SubmitComment
} from "@/api/PostAPI";
import Cookies from "js-cookie";
import {DislikePost, LikePost} from "@/api/PutAPI";
// import {GetAllPosts, PostPosts} from '@/api/GetAPI'
export default {
  name: 'PersonalCenter',
  computed:{
    ...mapGetters(['getSelectedTabIndex']),
  },
  data() {
    return {
      FPpostlist:[],
      commentForm:{
        comment:'',
      },
      uid:'',
      id:'',

      showPassword: false,
      likesshow:true,
      visible:false,
      // 上传头像的文件
      fileList:[],
      imageUrl: '',
      dialogVisible: false,
      dialogImageUrl: '',
      showFileList: true,
      autoUpload: true,
      activeIndex: '1',
      // 个人信息
      form: {
        avatar: 'https://randomuser.me/api/portraits/men/1.jpg',
        nickname: 'John Doe',
        password: '123456',
        add:'广东省肇庆市',
      },
      userInfo: {
        avatar: 'https://randomuser.me/api/portraits/men/1.jpg',
        nickname: 'John Doe'
      },
      ticketlist:[
        {
          ticketid:'',
          spotname:'黄山',
          selectedTicketType:'成人票',
          selectedDate:'2023年6月1日',
          price:'80',
          selectedQuantity:'2',
        },
        {spotname:'山东博物馆',
          spotPicture:'',
          selectedTicketType:'学生票',
          selectedDate:'2023年6月1日',
          price:'80',
          selectedQuantity:'2',
        },
        {
          ticketid:'',
          spotname:'珠海长隆乐园度假区',
          selectedTicketType:'儿童票',
          selectedDate:'2023年6月1日',
          price:'20',
          selectedQuantity:'2',
        },
      ],
      mypostlist:[{
        commentList:[{
          id: 1,
          user: {
            username: '小明',
            avatar: 'https://randomuser.me/api/portraits/men/1.jpg',
          },
          time: '2023-06-06 20:00',
          content: '这是一条评论',
        },
          {
            id: 2,
            user: {
              username: '小红',
              avatar: 'https://randomuser.me/api/portraits/women/1.jpg',
            },
            time: '2023-06-06 20:01',
            content: '这是另一条评论',
          },
          {
            id: 3,
            user: {
              username: '小丽',
              avatar: 'https://randomuser.me/api/portraits/men/1.jpg',
            },
            time: '2023-06-06 20:00',
            content: '这是一条评论',
          },],
        commentSidebarVisible:false,
        expand: false,
        avatar:require('../components/zyx.jpg'),
        title:'帖子标题1',
        UserID:'zyxwcy',
        Contents:'顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"',
        RTime:'2023.06.06',
        Likenum:10,
        collection:10,
        Tag1:'安徽合肥',
        Tag2:'名胜古迹',
        isClicked:false,
        isLiked:false,
        pic:[require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg'),require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg'),require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg')]
      },
        {
          commentList:[{
            id: 1,
            user: {
              username: '小明',
              avatar: 'https://randomuser.me/api/portraits/men/1.jpg',
            },
            time: '2023-06-06 20:00',
            content: '这是一条评论',
          },
            {
              id: 2,
              user: {
                username: '小红',
                avatar: 'https://randomuser.me/api/portraits/women/1.jpg',
              },
              time: '2023-06-06 20:01',
              content: '这是另一条评论',
            },
            {
              id: 3,
              user: {
                username: '小丽',
                avatar: 'https://randomuser.me/api/portraits/men/1.jpg',
              },
              time: '2023-06-06 20:00',
              content: '这是一条评论',
            },],
          commentSidebarVisible:false,
          expand: false,
          avatar:require('../components/zyx.jpg'),
          title:'帖子标题1',
          UserID:'zyxwcy',
          Contents:'顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"',
          RTime:'2023.06.06',
          Likenum:10,
          collection:10,
          Tag1:'安徽合肥',
          Tag2:'名胜古迹',
          isClicked:false,
          isLiked:false,
          pic:[require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg'),require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg')]
        },
        {
          commentList:[{
            id: 1,
            user: {
              username: '小明',
              avatar: 'https://randomuser.me/api/portraits/men/1.jpg',
            },
            time: '2023-06-06 20:00',
            content: '这是一条评论',
          },
            {
              id: 2,
              user: {
                username: '小红',
                avatar: 'https://randomuser.me/api/portraits/women/1.jpg',
              },
              time: '2023-06-06 20:01',
              content: '这是另一条评论',
            },
            {
              id: 3,
              user: {
                username: '小丽',
                avatar: 'https://randomuser.me/api/portraits/men/1.jpg',
              },
              time: '2023-06-06 20:00',
              content: '这是一条评论',
            },],
          commentSidebarVisible:false,
          expand: false,
          avatar:require('../components/zyx.jpg'),
          title:'帖子标题1',
          UserID:'zyxwcy',
          Contents:'顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"',
          RTime:'2023.06.06',
          Likenum:10,
          collection:10,
          Tag1:'安徽合肥',
          Tag2:'名胜古迹',
          isClicked:false,
          isLiked:false,
          pic:[require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg'),require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg')]
        },
        {
          commentList:[{
            id: 1,
            user: {
              username: '小明',
              avatar: 'https://randomuser.me/api/portraits/men/1.jpg',
            },
            time: '2023-06-06 20:00',
            content: '这是一条评论',
          },
            {
              id: 2,
              user: {
                username: '小红',
                avatar: 'https://randomuser.me/api/portraits/women/1.jpg',
              },
              time: '2023-06-06 20:01',
              content: '这是另一条评论',
            },
            {
              id: 3,
              user: {
                username: '小丽',
                avatar: 'https://randomuser.me/api/portraits/men/1.jpg',
              },
              time: '2023-06-06 20:00',
              content: '这是一条评论',
            },],
          commentSidebarVisible:false,
          expand: false,
          avatar:require('../components/zyx.jpg'),
          title:'帖子标题1',
          UserID:'zyxwcy',
          Contents:'顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"',
          RTime:'2023.06.06',
          Likenum:10,
          collection:10,
          Tag1:'安徽合肥',
          Tag2:'名胜古迹',
          isClicked:false,
          isLiked:false,
          pic:[require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg'),require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg'),require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg')]
        },
        {
          commentList:[{
            id: 1,
            user: {
              username: '小明',
              avatar: 'https://randomuser.me/api/portraits/men/1.jpg',
            },
            time: '2023-06-06 20:00',
            content: '这是一条评论',
          },
            {
              id: 2,
              user: {
                username: '小红',
                avatar: 'https://randomuser.me/api/portraits/women/1.jpg',
              },
              time: '2023-06-06 20:01',
              content: '这是另一条评论',
            },
            {
              id: 3,
              user: {
                username: '小丽',
                avatar: 'https://randomuser.me/api/portraits/men/1.jpg',
              },
              time: '2023-06-06 20:00',
              content: '这是一条评论',
            },],
          // commentSidebarVisible:false,
          expand: false,
          avatar:require('../components/zyx.jpg'),
          title:'帖子标题1',
          UserID:'zyxwcy',
          Contents:'顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"',
          RTime:'2023.06.06',
          Likenum:10,
          collection:10,
          Tag1:'安徽合肥',
          Tag2:'名胜古迹',
          isClicked:false,
          isLiked:false,
          commentSidebarVisible:false,
          pic:[require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg'),]
        },
      ],
      postlist:[{
        commentList:[{
          id: 1,
          user: {
            username: '小明',
            avatar: 'https://randomuser.me/api/portraits/men/1.jpg',
          },
          time: '2023-06-06 20:00',
          content: '这是一条评论',
        },
          {
            id: 2,
            user: {
              username: '小红',
              avatar: 'https://randomuser.me/api/portraits/women/1.jpg',
            },
            time: '2023-06-06 20:01',
            content: '这是另一条评论',
          },
          {
            id: 3,
            user: {
              username: '小丽',
              avatar: 'https://randomuser.me/api/portraits/men/1.jpg',
            },
            time: '2023-06-06 20:00',
            content: '这是一条评论',
          },],
        commentSidebarVisible:false,
        expand: false,
        avatar:require('../components/zyx.jpg'),
        Title:'帖子标题1',
        UserID:'zyxwcy',
        Contents:'顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"',
        RTime:'2023.06.06',
        Likenum:10,
        collection:10,
        Tag1:'安徽合肥',
        Tag2:'名胜古迹',
        isClicked:false,
        isLiked:false,
        pic:[require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg'),require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg'),require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg')]
      },
        {
          commentList:[{
            id: 1,
            user: {
              username: '小明',
              avatar: 'https://randomuser.me/api/portraits/men/1.jpg',
            },
            time: '2023-06-06 20:00',
            content: '这是一条评论',
          },
            {
              id: 2,
              user: {
                username: '小红',
                avatar: 'https://randomuser.me/api/portraits/women/1.jpg',
              },
              time: '2023-06-06 20:01',
              content: '这是另一条评论',
            },
            {
              id: 3,
              user: {
                username: '小丽',
                avatar: 'https://randomuser.me/api/portraits/men/1.jpg',
              },
              time: '2023-06-06 20:00',
              content: '这是一条评论',
            },],
          commentSidebarVisible:false,
          expand: false,
          avatar:require('../components/zyx.jpg'),
          title:'帖子标题1',
          UserID:'zyxwcy',
          Contents:'顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"',
          RTime:'2023.06.06',
          Likenum:10,
          collection:10,
          Tag1:'安徽合肥',
          Tag2:'名胜古迹',
          isClicked:false,
          isLiked:false,
          pic:[require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg'),require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg')]
        },

        {
          commentList:[{
            id: 1,
            user: {
              username: '小明',
              avatar: 'https://randomuser.me/api/portraits/men/1.jpg',
            },
            time: '2023-06-06 20:00',
            content: '这是一条评论',
          },
            {
              id: 2,
              user: {
                username: '小红',
                avatar: 'https://randomuser.me/api/portraits/women/1.jpg',
              },
              time: '2023-06-06 20:01',
              content: '这是另一条评论',
            },
            {
              id: 3,
              user: {
                username: '小丽',
                avatar: 'https://randomuser.me/api/portraits/men/1.jpg',
              },
              time: '2023-06-06 20:00',
              content: '这是一条评论',
            },],
          commentSidebarVisible:false,
          expand: false,
          avatar:require('../components/zyx.jpg'),
          title:'帖子标题1',
          UserID:'zyxwcy',
          Contents:'顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"',
          RTime:'2023.06.06',
          Likenum:10,
          collection:10,
          Tag1:'安徽合肥',
          Tag2:'名胜古迹',
          isClicked:false,
          isLiked:false,
          pic:[require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg'),require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg')]
        },
        {
          commentList:[{
            id: 1,
            user: {
              username: '小明',
              avatar: 'https://randomuser.me/api/portraits/men/1.jpg',
            },
            time: '2023-06-06 20:00',
            content: '这是一条评论',
          },
            {
              id: 2,
              user: {
                username: '小红',
                avatar: 'https://randomuser.me/api/portraits/women/1.jpg',
              },
              time: '2023-06-06 20:01',
              content: '这是另一条评论',
            },
            {
              id: 3,
              user: {
                username: '小丽',
                avatar: 'https://randomuser.me/api/portraits/men/1.jpg',
              },
              time: '2023-06-06 20:00',
              content: '这是一条评论',
            },],
          commentSidebarVisible:false,
          expand: false,
          avatar:require('../components/zyx.jpg'),
          title:'帖子标题1',
          UserID:'zyxwcy',
          Contents:'顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"',
          RTime:'2023.06.06',
          Likenum:10,
          collection:10,
          Tag1:'安徽合肥',
          Tag2:'名胜古迹',
          isClicked:false,
          isLiked:false,
          pic:[require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg'),require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg'),require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg')]
        },
        {
          commentList:[{
            id: 1,
            user: {
              username: '小明',
              avatar: 'https://randomuser.me/api/portraits/men/1.jpg',
            },
            time: '2023-06-06 20:00',
            content: '这是一条评论',
          },
            {
              id: 2,
              user: {
                username: '小红',
                avatar: 'https://randomuser.me/api/portraits/women/1.jpg',
              },
              time: '2023-06-06 20:01',
              content: '这是另一条评论',
            },
            {
              id: 3,
              user: {
                username: '小丽',
                avatar: 'https://randomuser.me/api/portraits/men/1.jpg',
              },
              time: '2023-06-06 20:00',
              content: '这是一条评论',
            },],
          // commentSidebarVisible:false,
          expand: false,
          avatar:require('../components/zyx.jpg'),
          title:'帖子标题1',
          UserID:'zyxwcy',
          Contents:'顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"',
          RTime:'2023.06.06',
          Likenum:10,
          collection:10,
          Tag1:'安徽合肥',
          Tag2:'名胜古迹',
          isClicked:false,
          isLiked:false,
          commentSidebarVisible:false,
          pic:[require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg'),]
        },
      ],
      spots: [
        { address:"山东省", name:'',comment:'',collection:'',id:11,picture:'',score:'',text:'',type:'名胜古迹'},
        { address:"北京市", name:'故宫博物院',comment:'紫禁城',collection:'',id:11,picture:'',score:3.8,text:'',type:'名胜古迹'},
        { address:"山东省", name:'山东博物馆',comment:'宏伟壮阔，值得一看',collection:'',id:11,picture:'',score:'',text:'',type:'博物馆'},
        { address:"山东省", name:'',comment:'',collection:'',id:11,picture:'',score:'',text:'',type:'名胜古迹'},
        { address:"广西省", name:'',comment:'',collection:'',id:11,picture:'',score:'',text:'',type:'名胜古迹'},
        { address:"广东省", name:'',comment:'',collection:'',id:11,picture:'',score:5,text:'',type:'名胜古迹'},
        { address:"山东省", name:'',comment:'',collection:'',id:11,picture:'',score:3.9,text:'',type:'名胜古迹'},
        { address:"山东省", name:'',comment:'',collection:'',id:11,picture:'',score:4.2,text:'',type:'名胜古迹'},
        { address:"山东省", name:'',comment:'',collection:100,id:11,picture:'',score:4.7,text:'',type:'名胜古迹'},
        { address:"广东省", name:'鼎湖山',comment:'天然大氧吧',collection:99,id:11,picture:'',score:4.8,text:'',type:'自然风光'},
        { address:"山东省", name:'',comment:'',collection:5,id:11,picture:'',score:5.0,text:'',type:'自然风光'},
      ],
      LikeCache:[],
    }
  },
  setup() {
    const isCollect = ref(false)
    const isLike = ref(false)
    const handleCollect = () => {
      isCollect.value = !isCollect.value
      if (isCollect.value) {
        ElMessage.success('收藏成功')
      } else {
        ElMessage.success('取消收藏成功')
      }
    }

    const handleLike= () => {
      isLike.value = !isLike.value
      if (isLike.value) {
        ElMessage.success('点赞成功')
      } else {
        ElMessage.success('取消点赞成功')
      }
    }

    return {
      isCollect,
      isLike,
      handleCollect,
      handleLike,
    }
  },

  created() {
    this.$watch(()=>this.getSelectedTabIndex,this.gotoIndex);
    this.uid= getToken();

    PostEmail(this.uid).then(response=>{
      this.id = response.data.uid
      console.log(this.id)
      var params = new URLSearchParams()

      GetTicketlist(this.id).then(response=>{
        console.log(response)
        this.ticketlist = response.data.orderList
        console.log(response.data.orderList)
      })

      params.append('uid', this.id)
      // 个人信息

      PostAvatar(response.data.uid)
          .then(response => {
            this.form.avatar = URL.createObjectURL(new Blob([response.data],{type:"image/jpeg"}))
          })

      getInfo(params).then(response => {
        this.form.email = response.data.email;
        this.form.nickname = response.data.nickname;
        this.form.password = '';
      })
          .catch(error => {
            console.log(error);
          })
      this.loadLikeCache()
      GetMyPostList(this.id)
          .then(response => {
            console.log(response)
            this.mypostlist = response.data.posts
            this.mypostlist.forEach(item =>{
              item.isFP = false;
              item.isLiked = false;
              var uid
              PostEmail(this.uid).then(response=>{
                uid = response.data.uid
                GetUserFP(uid)
                    .then(response => {
                      this.mypostlist.forEach(item => {
                        response.data.data.forEach(fp => {
                          if(item.PostID === fp.postID){
                            item.isFP = true
                          }
                        })
                        let liked = this.LikeCache.find(temp => temp.userID === uid && temp.PostID === item.PostID)
                        if (liked){
                          item.isLiked = true
                        }
                      })
                    })
              })
              GetPoster(item.UserID)
                  .then(response=>{
                    PostAvatar(item.UserID)
                        .then(p=>{
                          item.avatar = URL.createObjectURL(new Blob([p.data],{type:"image/jpeg"}))
                        })
                    item.UserID = response.data.name
                  })
                  .catch(err => {
                    console.log(err)
                  })
              const piclist = []
              PostImg(item.PostID)
                  .then(response => {
                    item.isphoto = false
                    if (response.data.size !== 0){
                      piclist.push(URL.createObjectURL(new Blob([response.data],{type:'image/jpeg'})))
                      PostImg1(item.PostID).then(response=>{if (response.data.size !== 0){piclist.push(URL.createObjectURL(new Blob([response.data],{type:'image/jpeg'})))}}).catch(err=>{console.log(err)})
                      PostImg2(item.PostID).then(response=>{if (response.data.size !== 0){piclist.push(URL.createObjectURL(new Blob([response.data],{type:'image/jpeg'})))}}).catch(err=>{console.log(err)})
                      PostImg3(item.PostID).then(response=>{if (response.data.size !== 0){piclist.push(URL.createObjectURL(new Blob([response.data],{type:'image/jpeg'})))}}).catch(err=>{console.log(err)})
                      PostImg4(item.PostID).then(response=>{if (response.data.size !== 0){piclist.push(URL.createObjectURL(new Blob([response.data],{type:'image/jpeg'})))}}).catch(err=>{console.log(err)})
                      PostImg5(item.PostID).then(response=>{if (response.data.size !== 0){piclist.push(URL.createObjectURL(new Blob([response.data],{type:'image/jpeg'})))}}).catch(err=>{console.log(err)})
                      PostImg6(item.PostID).then(response=>{if (response.data.size !== 0){piclist.push(URL.createObjectURL(new Blob([response.data],{type:'image/jpeg'})))}}).catch(err=>{console.log(err)})
                      PostImg7(item.PostID).then(response=>{if (response.data.size !== 0){piclist.push(URL.createObjectURL(new Blob([response.data],{type:'image/jpeg'})))}}).catch(err=>{console.log(err)})
                      PostImg8(item.PostID).then(response=>{if (response.data.size !== 0){piclist.push(URL.createObjectURL(new Blob([response.data],{type:'image/jpeg'})))}}).catch(err=>{console.log(err)})
                      item.isphoto = true
                    }
                    console.log(piclist)
                    console.log(piclist.length)
                    item.pic = piclist
                    console.log(item.pic)

                  })
                  .catch(err =>{
                    console.log(err)
                  })
              GetComments(item.PostID)
                  .then(response => {
                    item.commentList = response.data.data
                    item.commentList.forEach(comment=>{
                      GetPoster(comment.userID)
                          .then(response => {
                            comment.name = response.data.name
                            PostAvatar(comment.userID)
                                .then(p=>{
                                  comment.avatar = URL.createObjectURL(new Blob([p.data],{type:"image/jpeg"}))
                                })
                          })
                    })
                  })
            })
          })
          .catch(error => {
            console.log(error);
          })

      // 收藏景区
      getCollectSpot(params).then(response => {
        console.log(response)
        this.spots = response.data.data
      })
          .catch(error => {
            console.log(error);
          })
      this.FPpostlist=[],
          GetFPpost(this.id)
              .then(response => {
                console.log(response)
                this.FPpostlist = response.data.posts
                this.FPpostlist.forEach(item =>{
                  item.isFP = false;
                  item.isLiked = false;
                  var uid
                  PostEmail(this.uid).then(response=>{
                    uid = response.data.uid
                    GetUserFP(uid)
                        .then(response => {
                          this.FPpostlist.forEach(item => {
                            response.data.data.forEach(fp => {
                              if(item.PostID === fp.postID){
                                item.isFP = true
                              }
                            })
                            let liked = this.LikeCache.find(temp => temp.userID === uid && temp.PostID === item.PostID)
                            if (liked){
                              item.isLiked = true
                            }
                          })
                        })
                  })
                  GetPoster(item.UserID)
                      .then(response=>{
                        PostAvatar(item.UserID)
                            .then(p=>{
                              item.avatar = URL.createObjectURL(new Blob([p.data],{type:"image/jpeg"}))
                            })
                        item.UserID = response.data.name
                      })
                      .catch(err => {
                        console.log(err)
                      })
                  const piclist = []
                  PostImg(item.PostID)
                      .then(response => {
                        item.isphoto = false
                        if (response.data.size !== 0){
                          piclist.push(URL.createObjectURL(new Blob([response.data],{type:'image/jpeg'})))
                          PostImg1(item.PostID).then(response=>{if (response.data.size !== 0){piclist.push(URL.createObjectURL(new Blob([response.data],{type:'image/jpeg'})))}}).catch(err=>{console.log(err)})
                          PostImg2(item.PostID).then(response=>{if (response.data.size !== 0){piclist.push(URL.createObjectURL(new Blob([response.data],{type:'image/jpeg'})))}}).catch(err=>{console.log(err)})
                          PostImg3(item.PostID).then(response=>{if (response.data.size !== 0){piclist.push(URL.createObjectURL(new Blob([response.data],{type:'image/jpeg'})))}}).catch(err=>{console.log(err)})
                          PostImg4(item.PostID).then(response=>{if (response.data.size !== 0){piclist.push(URL.createObjectURL(new Blob([response.data],{type:'image/jpeg'})))}}).catch(err=>{console.log(err)})
                          PostImg5(item.PostID).then(response=>{if (response.data.size !== 0){piclist.push(URL.createObjectURL(new Blob([response.data],{type:'image/jpeg'})))}}).catch(err=>{console.log(err)})
                          PostImg6(item.PostID).then(response=>{if (response.data.size !== 0){piclist.push(URL.createObjectURL(new Blob([response.data],{type:'image/jpeg'})))}}).catch(err=>{console.log(err)})
                          PostImg7(item.PostID).then(response=>{if (response.data.size !== 0){piclist.push(URL.createObjectURL(new Blob([response.data],{type:'image/jpeg'})))}}).catch(err=>{console.log(err)})
                          PostImg8(item.PostID).then(response=>{if (response.data.size !== 0){piclist.push(URL.createObjectURL(new Blob([response.data],{type:'image/jpeg'})))}}).catch(err=>{console.log(err)})
                          item.isphoto = true
                        }
                        item.pic = piclist

                      })
                      .catch(err =>{
                        console.log(err)
                      })
                  GetComments(item.PostID)
                      .then(response => {
                        item.commentList = response.data.data
                        item.commentList.forEach(comment=>{
                          GetPoster(comment.userID)
                              .then(response => {
                                comment.name = response.data.name
                                PostAvatar(comment.userID)
                                    .then(p=>{
                                      comment.avatar = URL.createObjectURL(new Blob([p.data],{type:"image/jpeg"}))
                                    })
                              })
                        })
                      })
                })
              })
              .catch(error => {
                console.log(error);
              })
    })

  },


  methods: {
    deletePost,
    showCommentSidebar(index,postid){
      this.mypostlist[index].commentSidebarVisible = true;
      Cookies.set('postID', postid)
    },
    showCommentSidebar1(index,postid){
      this.FPpostlist[index].commentSidebarVisible = true;
      Cookies.set('postID', postid)
    },
    SubmitComment,
    saveLikeCache(){
      let likeCacheJson = JSON.stringify(this.LikeCache)
      localStorage.setItem('likeCache', likeCacheJson)
    },
    loadLikeCache(){
      let likeCacheJson = localStorage.getItem('likeCache')
      if(likeCacheJson){
        this.LikeCache = JSON.parse(likeCacheJson)
      }
    },

    // 上传头像
    beforeAvatarUpload(file) {
      const isJPG = file.type === 'image/jpeg';
      const isPNG = file.type === 'image/png';
      const isLt2M = file.size / 1024 / 1024 < 2;

      if (!isJPG && !isPNG) {
        this.$message.error('上传头像图片只能是 JPG 或 PNG 格式');
      }
      if (!isLt2M) {
        this.$message.error('上传头像图片大小不能超过 2MB');
      }
      return (isJPG || isPNG) && isLt2M;
    },
    handleRemove(file, fileList) {
      console.log(file, fileList);
    },
    handlePictureCardPreview(file) {
      this.dialogImageUrl = file.url;
      this.dialogVisible = true;
    },


    // handleSuccess(response, file, fileList) {
    //   this.imageUrl = URL.createObjectURL(file.raw);
    //   this.showFileList = false;
    //   this.autoUpload = false;
    // },
    handlePhoto(file){
      console.log(file)
      this.form.avatar = file
    },
    gotoIndex(){
      this.activeIndex=this.getSelectedTabIndex;
      console.log(this.getSelectedTabIndex);
      console.log(this.activeIndex);
    },

    // 票型选择
    getImage(ticketType) {
      if (ticketType === '成人票') {
        return require('@/assets/t1.png');
      } else if (ticketType === '学生票') {
        return require('@/assets/t2.png');
      } else if (ticketType === '儿童票') {
        return require('@/assets/t3.png');
      } else {
        return '';
      }
    },

    goToSpot(spotid,c,s){
      sessionStorage.setItem(
          'spotID',
          JSON.stringify(
              {
                spot:spotid,
                collection:c,
                sname: s
              }
          )
      )
      this.$router.push('/spotdetail')
    },

    expand(index) {
      this.postlist[index].expand = true;
    },
    noexpand(index) {
      this.postlist[index].expand = false;
    },

    // 收藏景点函数
    toggleSpot(item) {
      // if(this.uid){
        item.collection = !item.collection;
        var params = new URLSearchParams();
        params.append('token', getToken())
        params.append("sid", item.id);

        if(item.collection) {
          item.collections += 1;
          CreateSpot1(params).then(response => {
            let a=response.data.msg
            ElMessage.success({
              message: a,
              type: 'success',
            })
          })
              .catch(error => {
                console.log(error);
              });
        }
        else {
          item.collections -= 1;

          DeleteSpot1(params).then(response => {
            let a=response.data.msg
            ElMessage.success({
              message: a,
              type: 'success',
            })
          })
              .catch(error => {
                console.log(error);
              });
        }
      // }

      // else{
      //   ElMessageBox.confirm('游客无法收藏，请先登录', '提示', {
      //     confirmButtonText: '去登录',
      //     cancelButtonText: '取消',
      //     type: 'warning'
      //   }).then(() => {
      //     // 用户点击了确定按钮，跳转至登录页面
      //     this.$router.push('/login');
      //   }).catch(() => {
      //     // 用户点击了取消按钮，不做操作
      //   });
      // }

    },



    // 点击点赞按钮的响应
    toggleLike(temp) {
      if(this.uid){

        PostEmail(this.uid).then(response=>{
          let uid = response.data.uid
          let liked = this.LikeCache.find(item => item.userID === uid && item.PostID === temp.PostID)
          if (liked){
            let index = this.LikeCache.findIndex(item => item.userID === liked.userID && item.PostID === liked.PostID)
            if(index !== -1){
              this.LikeCache.splice(index, 1)
              temp.Likenum -= 1
            }
            console.log("pop")
            this.saveLikeCache()
            //删除帖子点赞
            ElMessage.success({
              message: '取消点赞成功',
              type: 'success',
            })

            DislikePost(temp.PostID)
          }else {
            this.LikeCache.push({userID: uid, PostID: temp.PostID})

            ElMessage.success({
              message: '点赞成功',
              type: 'success',
            })

            this.saveLikeCache()
            //添加帖子点赞
            temp.Likenum += 1
            LikePost(temp.PostID)
          }
          console.log(this.LikeCache)
          temp.isLiked = !temp.isLiked;
        })

      }

      else{
        ElMessageBox.confirm('游客无法点赞，请先登录', '提示', {
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

    // 帖子收藏
    toggleColor(item) {
      if(this.uid){
        item.isFP = !item.isFP;
        PostEmail(this.uid).then(response=>{
          this.id = response.data.uid
          if(item.isFP){
            item.collection += 1;
            AddUserFP(this.id, item.PostID)
            ElMessage.success({
              message: '收藏成功',
              type: 'success',
            })

          } else {
            item.collection -= 1;
            DeleteUserFP(this.id, item.PostID)
            ElMessage.success({
              message: '取消收藏成功',
              type: 'success',
            })

          }
        })
      }

      else{
        ElMessageBox.confirm('游客无法收藏，请先登录', '提示', {
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


    // clicklikes() {
    //   this.likesshow = !this.likesshow;
    //   this.form.number = !this.likesshow
    //       ? this.form.number + 1
    //       : this.form.number - 1;
    // },

    handleSelect(index) {
      this.activeIndex = index
    },

    // 修改个人资料
    submitForm() {
      // 提交表单数据
      console.log(this.form)
      PostEmail(getToken())
          .then(response=>{
            console.log(response)
            this.form.uid = response.data.uid
          })

      updateUser(this.form).then(response => {
        console.log(response);
        if(response.data.code===200){
          ElMessage.success("更新信息成功！");
        }
        else{
          // let a=response.data.msg
          // ElMessage.error(a);
        }

      })
          .catch(error => {
            console.log(error);
          });
    }

  }
}
</script>

<style>
#confirmButton{
  background-color: rgb(255, 255, 251);
  border-radius: 8px;
  padding: 2px;
  width: 80px;
}

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}
.custom-label .el-form-item__label {
  font-family: 宋体;
  color: rgba(255, 255, 255, 0.83);
}


.comment-button {
  position: absolute;
  bottom: 0;
  right: 0;
  margin: 10px;
}

.comment_button :hover{
  transform: scale(1.1);
}

.idenAll{
  background-image: url('../assets/Iden.jpg');
  position: relative;
  overflow: hidden;
  height: 100vh;

}
.personal-center {
  display: flex;
  flex-direction: row;
  height: 100%;
}
.nav-bar {
  /* margin-left: 80px; */
  margin-top: 150px;
  position: fixed;
  width: 200px;
}
.default-info .info-item {
  margin-bottom: 10px;
}
.default-info .label {
  display: inline-block;
  width: 80px;
  font-weight: bold;
}
.default-info .value {
  display: inline-block;
}

.image-container {
  position: relative;
  height: 300px;
}

.ticket-image{
  width:720px;
  height: 250px;
  border-radius: 10px;
  box-shadow: 0 14px 28px rgba(0, 0, 0, .25), 0 10px 10px rgba(.28, 0, 0, .22);
}



.ticket-container {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.ticket-name{
  font-family:华文中宋;
  /* font-family:幼圆; */
  font-weight: bold;
  font-size: 35px;
  color: rgba(240, 248, 255, 0.945);
}

.ticket-date{
  font-family: 仿宋;
  color:  rgba(247, 248, 249, 0.909);
  font-size: 10px;
}

.image-container {
  position: relative;
}

.上传头像图标{
  width: 50px;
  margin-top: 10px;
  height: 50px;
}

.ticket-price{
  position: absolute;
  right: 68px;
  top:65px;
  z-index: 1;
  font-weight: bold;
  /* margin-bottom: 20px;
  margin-left: 5px; */
  font-size: 50px;
  font-weight:80;
  font-family: Forte;
  /* font-family: "Montserrat", sans-serif; */
}
.text-overlay {
  /* border: 1px solid white; */
  position: absolute;
  align-items: center;
  top: 25px;
  left: 12px;
  z-index: 1;
}


.address{
  font-family: 华文行楷;
}
.address1{
  font-family: 华文行楷;
}

.like-style {
  display: inline-block;
  background-color: #fab4eb;
  color: #ffffffdd;
  border: none;
  cursor: pointer;
  position: relative;
  box-shadow: 0 2px 20px #fab4eb;
  outline: 0;
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
  background-image: radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
    /*  */
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%);
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

.text {
  font-family: 楷体;
  font-size: 15px;
  /*overflow: hidden;*/
  /*text-overflow: ellipsis;*/
}
.comment{
  font-family: 华文行楷;
  font-size: 15px;
}
.name{
  font-family: 华文新魏;
  font-weight: bold;
  font-size: 35px;
}
.carousel-caption {
  position: absolute;
  text-align: center;
  bottom: 0px;
  left: 0;
  right: 0;
  background-color: rgba(43, 41, 41, 0.7);
  color: #fff;

  /* padding: 5px; */
}
.景区缩略图页面{
  background-image: url('../assets/back.jpg');
  position: relative;
  background-size: cover;
  background-attachment: fixed;
  margin-top: 58px;

}

.景点收藏{
  text-align: right;
}
.景点信息 {
  margin: 5px auto;
  padding: 10px;
  /* background-image: linear-gradient(to top, #fff1eb 0%, #ace0f9 100%); */
  display: flex;
  opacity: 0.85;
  flex-direction: column;
  justify-content: space-between;
}
.spots {
  /* width: calc(100% - 220px - 5px); */
  margin-left: 200px;
  margin-right: 50px;
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
}


.info {
  /* display: flex; */
  justify-content: space-between;
  margin-top: 0px;
  border: 0px;
}


.板块信息 {
  /* margin: 5px auto; */
  /* margin-left: 200px; */
  /* margin-right: 200px; */
  border-left: 10px;
  border-right: 10px;
  padding: 20px;
  background: rgba(255, 255, 255, 0.194);
  /* align-items: center; */
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  vertical-align:top;
}



.导航栏 {
  width: 100%;
  height: 50px;
  background-color: rgba(51, 51, 51, 0.4);
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
}
#allIden{
  background-position: center;
  background-repeat: no-repeat;
  height: 100vh; /* 设置容器高度为视窗高度 */
  background-image: url('../assets/IdenBack.jpg');
  background-position: 100lvh;
  background-size: cover;
  position: relative;
  background-attachment: fixed;
  margin-top: 58px;
}
.喜欢按钮{
  background: none;
  border: none;
  padding: 0;
}

/* 帖子图片展示 */
.scrollbar-flex-content {
  display: flex;
}
.scrollbar-demo-item {
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 180px;
  height: 170px;
  margin: 10px;
  text-align: center;
  border-radius: 4px;
  background: var(--el-color-danger-light-9);
  color: var(--el-color-danger);
}
.comment-sidebar {
  text-align: left;
  width: 100px;
}
.comment_button :hover{
  transform: scale(1.2);
}

.bubble_button :hover{
  transform: scale(1.2);
}

.comment-sidebar-container {
  position: fixed;
  /* text-align: left; */
  top: 0;
  right: -40%;
  bottom: 0;
  width: 40%;
  background-color: #fff;
  z-index: 999;
  transition: transform 0.3s ease-out;
  height: 100vh;
}

.comment-sidebar-visible {
  transform: translateX(-50%);
}

.comment-header {
  display: flex;
  /* align-items: center; */
  /* justify-content: space-between; */
  padding: 10px;
  font-size: 18px;
  font-weight: bold;
  border-bottom: 1px solid #eee;
}

.comment-header i {
  cursor: pointer;
}

.comment-form {
  margin-top: 2px;
  margin-bottom: 30px;
  /* justify-content: flex-start;  */
  display: flex;
}

.comment-card {
  margin-bottom: 20px;
  width: 320px;
  margin-left:10px;
}

.comment-header {
  display: flex;
  align-items: center;
}

.comment-avatar {
  margin-right: 10px;
}

.comment-info {

  margin-left: 9px;
  overflow: auto;
  /* justify-content: flex-start; */
  margin-left: 0;
}

.comment-username {
  font-size: 16px;
  font-weight: bold;
  margin-bottom: 5px;
}

.comment-time {
  font-size: 12px;
  color: #999;
}
.comment-list-container {
  max-height: 75%;
  overflow-y: auto;
}


.comment-wrapper {
  position: relative;
  margin-top: -20px;
}

.comment-button {
  position: absolute;
  bottom: 0;
  right: 0;
  margin: 10px;
}
.customdialogclass{
  height: 300px;
  overflow: auto;
}

.logo {
  width: 34px;
  height: 34px;
  border-radius: 50%;
  overflow: hidden;
}

.right-function a {
  color: white;
  text-decoration: none;
}
.主体 {
  height: 100%;
  width: 1000px;
  margin-left: 250px;
  display: flex;
  justify-content: space-between;
}

.star-style{
  display: inline-block;

  background-color: #f5e473e2;
  color: #ffffffdd;

  border: none;
  cursor: pointer;
  position: relative;
  box-shadow: 0 2px 25px #f5e473b8;
  /* #fce239 */
  outline: 0;
  transition: transform ease-in 0.1s, background-color ease-in 0.1s, box-shadow ease-in 0.25s;
}

.star-style::before {
  position: absolute;
  content: '';
  left: -2em;
  right: -2em;
  top: -2em;
  bottom: -2em;
  pointer-events: none;
  background-repeat: no-repeat;
  background-image: radial-gradient(circle, #f5e473e2 20%, transparent 20%),
  radial-gradient(circle, #f5e473e2 20%, transparent 20%),
  radial-gradient(circle, #f5e473e2 20%, transparent 20%),
  radial-gradient(circle, #f5e473e2 20%, transparent 20%),
  radial-gradient(circle, #f5e473e2 20%, transparent 20%),
  radial-gradient(circle, #f5e473e2 20%, transparent 20%),
  radial-gradient(circle, #f5e473e2 20%, transparent 20%),
  radial-gradient(circle, #f5e473e2 20%, transparent 20%),
  radial-gradient(circle, #f5e473e2 20%, transparent 20%),
  radial-gradient(circle, #f5e473e2 20%, transparent 20%),
  radial-gradient(circle, #f5e473e2 20%, transparent 20%),
  radial-gradient(circle, #f5e473e2 20%, transparent 20%),
  radial-gradient(circle, #f5e473e2 20%, transparent 20%),
  radial-gradient(circle, #f5e473e2 20%, transparent 20%),
  radial-gradient(circle, #f5e473e2 20%, transparent 20%);
  background-position: 5% 44%, -5% 20%, 7% 5%, 23% 0%, 37% 0, 58% -2%, 80% 0%, 100% -2%, -5% 80%,
  100% 55%, 2% 100%, 23% 100%, 42% 100%, 60% 95%, 70% 96%, 100% 100%;
  background-size: 0% 0%;
  transition: background-position .5s ease-in-out, background-size .75s ease-in-out;
}

.star-style:active::before {
  transition:0s;
  background-size: 10% 10%, 20% 20%, 15% 15%, 20% 20%, 18% 18%, 10% 10%, 15% 15%, 10% 10%, 18% 18%,
  15% 15%, 20% 20%, 18% 18%, 20% 20%, 15% 15%, 10% 10%, 20% 20%;
  background-position: 18% 40%, 20% 31%, 30% 30%, 40% 30%, 50% 30%, 57% 30%, 65% 30%, 80% 32%, 15% 60%,
  83% 60%, 18% 70%, 25% 70%, 41% 70%, 50% 70%, 64% 70%, 80% 71%;
}

.like-style {
  display: inline-block;

  background-color: #fab4eb;
  color: #ffffffdd;

  border: none;
  cursor: pointer;
  position: relative;
  box-shadow: 0 2px 20px #fab4eb;
  outline: 0;
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
  background-image: radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%),
  radial-gradient(circle, #fab4eb 20%, transparent 20%);
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
.comment-content{
  margin-top: 3px;
  font-family: 等线;
  font-size: 15px;
}
.个人中心 {
  width: 190px;
  height: 300px;
  padding: 15px;
  background: rgba(255,255,255,0.6);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-between;
}

.头像 {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  overflow: hidden;
}

.侧边栏 {
  margin-left: 30px;
  margin-top: 40px;
  width: 200px;
  position: fixed;
  text-align: right;

}
.板块列表{
  height: 100%;
  overflow: hidden;
  width: calc(100% - 190px - 5px);

  padding: 15px;
  background: rgba(255,255,255,0.6);
}
.tag{
  margin-right: 8px;
  size: 15px;
}
.点赞收藏标签{
  display: flex;
  width:400px;
}

.标签{
  float: left;
  width: 30%;
  size: large;
}
.element{
  margin-right: 10px;
}

.点赞收藏评论{
  text-align: center;
  display: flex;
  float: right;
  width: 30%;
}


.圆角框{
  /* align-items: left; */
  padding: 2px;
  height: 30px;
  width: 65px;
  margin-right: 10px;
  background: rgba(255, 255, 255, 0.906);
  border-radius: 8px;
}

ol {
  list-style: none;
}

.板块信息 {
  margin: 6px auto;
  padding: 10px;
  /* opacity: 50%; */
  background-color: #eeeeeecb;;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.button-container {

  display: flex;
  flex-direction: column;
}

.评论 {
  margin: 0;
  padding:0;
  height: 50px;
  width: 500px;
  font-family:楷体,sans-serif;
  font-size:9pt;
  background: rgba(165, 220, 235, 0.71);
  text-align:left;
  font-variant:small-caps;
  flex-direction: column;
  justify-content: space-between;
}
.帖子名称 {
  font-size: 23px;
  text-align: center;
  font-family: 方正粗黑宋简体;
  font-weight: bold;

}

.主贴时间 {

  font-size: 10px;
  text-align: right;
}
.帖主 {
  font-size: 10px;
  text-align: left;
}
.帖子内容 {
  padding: 5px;
  text-indent: 2em;
  font-size: 14px;
  /* font-family: 楷体;
  font-weight: 50; */
}

.帖子头部{
  display: flex;
}
.点击 {
  display: flex;
  justify-content: right;
  text-align: center;
  font-size: 10px;
}
.数量 {
  margin: 0px 10px 0 0;
}

#title {
  height: 30px;
  width: 520px;
  text-indent: 10px;
  border-radius: 10px;
  outline: none;
  border: none;
  background-color:rgba(166, 213, 222, 0.39);
}

#submit {
  height: 40px;
  width: 80px;

  color: white;
  background: orange;
  outline: none;
  border: none;
  border-radius: 10px;

}

.标签{
  order: 1;
}

</style>
  