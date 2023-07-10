<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <div class="社区页面">
    <el-container style="height: 100%">
      <div class="侧边栏" >
        <div class="button-container">
          <div>
            <el-button class="button tada" size="middle" round @click="orderTime" style="margin-bottom: 4px;">按时间</el-button>
          </div>
          <div>
            <el-button class="button tada" size="middle" round @click="orderLike">按点赞</el-button>
          </div>

          <div style="margin-top: 20px;">
            <el-tooltip
                class="item"
                effect="light"
                content="发送帖子"
                placement="bottom"
            >

              <el-button size="large" circle @click="ifvisible() ">
                <el-icon>
                  <Plus/>
                </el-icon>
              </el-button>
            </el-tooltip>
          </div>
        </div>
      </div>

      <!-- 发帖页面 -->
      <div>
        <el-dialog v-model="visible" title="发帖">
          <el-form :model="formData" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
            <el-form-item label="标题：" prop="title">
              <el-input type='text' v-model="formData.title" placeholder="输入标题..." style="margin-top: 10px;"></el-input>
            </el-form-item>
            <el-form-item label="地区：" prop="region">
              <el-cascader v-model="formData.tag1" :options="options" :props="{label:'label',value:'label',children:'children'}" placeholder="景区所在地" :formatter="formatter"></el-cascader>
            </el-form-item>
            <!-- 下拉框 -->
            <el-form-item label="标签：" prop="tag2">
              <el-select type='text' v-model="formData.tag2" placeholder="景区类型">
                <el-option v-for="item in type" :key="item" :value="item"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="内容：" prop="content">
              <el-input type='textarea' :rows="8" v-model="formData.content" :autosize="{ minRows: 8, maxRows: 12}" placeholder="请输入正文..." style="margin-top: 10px;"></el-input>
            </el-form-item>
            <el-form-item label="图片：">
              <el-upload
                  action="http://127.0.0.1:8888/api/private/v1/upload"
                  list-type="picture-card"
                  :limit=9
                  :on-preview="handlePictureCardPreview"
                  :before-upload="beforeAvatarUpload"
                  :on-remove="handleRemove"
                  :on-success="handleSuccess"
                  :on-change="handlePhoto"
                  :headers="headerObj"
                  :show-file-list="true"
                  :auto-upload="false"
                  multiple>
                <el-icon><Plus /></el-icon>
              </el-upload>
              <el-dialog v-model="dialogVisible">
                <div class="customdialogclass">
                  <img width="100%" :src="dialogImageUrl" alt="">
                </div>
              </el-dialog>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="onSubmit" style="float:right">提交</el-button>
            </el-form-item>
          </el-form>
        </el-dialog>
      </div>
      <!-- 帖子列表 -->
      <div class="主体" >
        <div style="display: flex;">
          <div id="content" class="列表">
            <el-space wrap>
              <el-card class="板块信息" style="width: 800px"  v-for = "(item, index) in postlist" :key="index">
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
                  <div class="点赞收藏评论标签" style="position: relative;">

                    <div class="标签" >
                      <el-tag class="tag">{{ item.Tag1 }}</el-tag>
                      <el-tag class="tag">{{ item.Tag2 }}</el-tag>
                    </div>
                    <div class="点赞收藏评论" style="display: flex;" >
                      <!-- 点赞 -->
                      <div class="element">
                        <el-button class="like-style"  v-bind:style="{backgroundColor: item.isLiked ? '#d3123a' : 'white'}" style="z-index: 1;" size="middle" id="likeButton" circle  v-on:click="toggleLike(item)" >
                          <img :src="require('@/assets/heart.png')" style="height: 20px;width: 23px; z-index: 3;margin-left: 0px;">
                        </el-button>
                      </div>

                      <div class="element" style="margin-top: 4px;">
                        {{ item.Likenum }}
                      </div>

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

                          <el-icon class="comment_button"  size="30" round style="text-align: right;margin-top: 2px;cursor: pointer;" @click="showCommentSidebar(index,item.PostID)"><ChatLineRound /></el-icon>
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
                            <el-tooltip
                                class="item"
                                effect="light"
                                content="评论"
                                placement="bottom"
                            >
                              <el-icon class="comment-button" @click="SubmitComment(this.uid)" size="30" round style="text-align: center;cursor: pointer;"><Promotion /></el-icon>
                            </el-tooltip>
                          </div>
                        </el-form>

                      </div>
                      <div class="comment-list-container">
                        <el-card v-for="comment in item.commentList" :key="comment.id" class="comment-card">
                          <div class="comment-header">
                            <el-avatar :src="comment.avatar" size="medium" style="margin-right: 6px;margin-left: -3px;"></el-avatar>
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

        </div>
      </div>
      <el-backtop :right="50" :bottom="50" />
    </el-container>

  </div>
</template>


<script>
import {ElAvatar, ElButton, ElCard, ElForm, ElFormItem, ElInput, ElMessage, ElMessageBox} from 'element-plus'
import {options} from '../components/data'
import {getToken} from './token';
import Cookies from "js-cookie"
import {GetAllPosts, GetComments, GetPoster, GetPosts, GetUserFP, SearchPost} from "@/api/GetAPI";
import {
  AddUserFP, PostAvatar,
  PostEmail,
  PostImg,
  PostImg1,
  PostImg2,
  PostImg3,
  PostImg4,
  PostImg5,
  PostImg6,
  PostImg7,
  PostImg8,
  SubmitComment
} from "@/api/PostAPI";
import {DislikePost, LikePost} from "@/api/PutAPI";
import {DeleteUserFP} from "@/api/DeleteAPI";
import {onSubmit} from "@/api/FunctionAPI";
import { mapGetters } from "vuex";

export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name:'Community',
  components: {
    ElButton,
    // ElDrawer,
    ElAvatar,
    ElForm,
    ElFormItem,
    ElInput,
    ElCard,
  },
  computed:{
    ...mapGetters(['getSearchContent']),
  },
  data(){
    return{
      rules:{
        title:[{
          required:true,message:'请输入标题',trigger:'blur'
        }],
        region:[{
          type:'array',required:true,message:'请选择地区',trigger:'blur'
        }],
        tag2:[{
          required:true,message:'请选择类型',trigger:'blur'
        }],
        content:[{
          required:true,message:'请输入内容',trigger:'blur'
        }],
      },
      headerObj:{
        Authorization:'Bearer ' + localStorage.getItem('token')
      },
      commentForm:{
        comment:'',
      },
      LikeCache:[{
        userID:0,
        PostID:0,
      }],
      LikeComment:[{
        userID:0,
        PostID:0,
        CommentID:0,
      },],

      currentUserAvatar:'https://randomuser.me/api/portraits/men/2.jpg',
      submittingComment:false,
      options,
      showcommentText:false,
      uid:'',
      dialogImageUrl: [],
      dialogVisible: false,
      value:'',
      hideshowPopup:true,
      ruleFormtitle:'',
      region:[],
      content:'',
      fileList:[],
      picture:'',
      visible:false,
      props: {
        value: 'value',
        label: 'label',
        children: 'children'
      },
      ruleForm: {
        ruleFormtitle:'',
        region: '',
        value:'',
        content: '',
        picture: '',
        images: []
      },

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
        title:'帖子标题1',
        UserID:'zyxwcy',
        Contents:'顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"顾名思义就是用来灌水的跟贴。何谓灌水呢？灌水"，故言而之就是"发表没有实际阅读意"',
        RTime:'2023.06.06',
        Likenum:19,
        collection:18,
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
          pic:[require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg'),]
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
          pic:[require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg'),]
        },
        {
          commentList:[],
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
          pic:[require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg'),]
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
          ],
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
          pic:[require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg'),]
        },{
          commentList:[
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
          pic:[require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg'),]
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
          pic:[require('../components/黄山1.jpg'),require('../components/黄山2.jpg'),require('../components/黄山1.jpg'),]
        },
      ],

      id:0,
      formData:{
        userID:'',
        postID:'',
        title:'',
        tag1:[],
        tag2:[],
        content:'',
        photo:null,
        fileList:[],
        likenum:''
      },
      type:['自然风光','博物馆','名胜古迹','地标观景'],


    }
  },

  created(){
    this.$watch(() => this.getSearchContent,this.searchHandler)
    console.log(this.getSearchContent)
    // 获取当前用户id
    this.uid = getToken();

    GetPosts()
        .then(response => {
          this.postlist = response.data.data
          console.log(this.postlist)
          this.postlist.forEach(item =>{
            item.isFP = false;
            item.isLiked = false;
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
            const requests = []
            requests.push(PostImg(item.PostID))
            requests.push(PostImg1(item.PostID))
            requests.push(PostImg2(item.PostID))
            requests.push(PostImg3(item.PostID))
            requests.push(PostImg4(item.PostID))
            requests.push(PostImg5(item.PostID))
            requests.push(PostImg6(item.PostID))
            requests.push(PostImg7(item.PostID))
            requests.push(PostImg8(item.PostID))
            Promise.all(requests)
                .then(responses => {
                  item.isphoto = false;
                  responses.forEach(response => {
                    if (response.data.size !== 0) {
                      piclist.push(URL.createObjectURL(new Blob([response.data], {type: 'image/jpeg'})));
                    }
                  });
                  item.isphoto = true;
                  item.pic = piclist
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
                  console.log(item.commentList)
                })
          })
          var uid
          PostEmail(this.uid).then(response=>{
            uid = response.data.uid
            GetUserFP(uid)
                .then(response => {
                  this.postlist.forEach(item => {
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
        })
        .catch(error => {
          console.log(error);
        })
    this.loadLikeCache()

  },


  methods:{
    formatter(value, selectedOptions) {
      return selectedOptions.map(option => option.label).join(' / ');
    },

    searchHandler(){
      const searchContent =this.getSearchContent;
      console.log(searchContent)
      SearchPost(searchContent)
          .then(response => {
            console.log(response)
            this.postlist = response.data.posts
            this.postlist.forEach(item =>{
              item.isFP = false;
              item.isLiked = false;
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
              const requests = []
              requests.push(PostImg(item.PostID))
              requests.push(PostImg1(item.PostID))
              requests.push(PostImg2(item.PostID))
              requests.push(PostImg3(item.PostID))
              requests.push(PostImg4(item.PostID))
              requests.push(PostImg5(item.PostID))
              requests.push(PostImg6(item.PostID))
              requests.push(PostImg7(item.PostID))
              requests.push(PostImg8(item.PostID))
              Promise.all(requests)
                  .then(responses => {
                    item.isphoto = false;
                    responses.forEach(response => {
                      if (response.data.size !== 0) {
                        piclist.push(URL.createObjectURL(new Blob([response.data], {type: 'image/jpeg'})));
                      }
                    });
                    item.isphoto = true;
                    item.pic = piclist
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
                    console.log(item.commentList)
                  })
            })
            var uid
            PostEmail(this.uid).then(response=>{
              uid = response.data.uid
              GetUserFP(uid)
                  .then(response => {
                    this.postlist.forEach(item => {
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
          })
          .catch(error => {
            console.log(error);
          })
    },
    orderTime(){
      GetAllPosts()
          .then(response => {
            this.postlist = response.data.data
            this.postlist.forEach(item =>{
              item.isFP = false;
              item.isLiked = false;
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
              const requests = []
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
                    Promise.all(requests)
                        .then(responses => {
                          item.isphoto = false;
                          responses.forEach(response => {
                            if (response.data.size !== 0) {
                              piclist.push(URL.createObjectURL(new Blob([response.data], {type: 'image/jpeg'})));
                            }
                          });
                          item.isphoto = true;
                        })
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
                    console.log(item.commentList)
                  })
            })
            var uid
            PostEmail(this.uid).then(response=>{
              uid = response.data.uid
              GetUserFP(uid)
                  .then(response => {
                    this.postlist.forEach(item => {
                      response.data.data.forEach(fp => {
                        if(item.PostID === fp.postID){
                          item.isFP = 'true'
                        }
                      })
                      let liked = this.LikeCache.find(temp => temp.userID === uid && temp.PostID === item.PostID)
                      if (liked){
                        item.isLiked = true
                      }
                    })
                  })
            })
          })
          .catch(error => {
            console.log(error);
          })
      this.loadLikeCache()
    },
    orderLike(){
      GetPosts()
          .then(response => {
            this.postlist = response.data.data
            this.postlist.forEach(item =>{
              item.isFP = false;
              item.isLiked = false;
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
              const requests = []
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
                    Promise.all(requests)
                        .then(responses => {
                          item.isphoto = false;
                          responses.forEach(response => {
                            if (response.data.size !== 0) {
                              piclist.push(URL.createObjectURL(new Blob([response.data], {type: 'image/jpeg'})));
                            }
                          });
                          item.isphoto = true;
                        })
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
                    console.log(item.commentList)
                  })
            })
            var uid
            PostEmail(this.uid).then(response=>{
              uid = response.data.uid
              GetUserFP(uid)
                  .then(response => {
                    this.postlist.forEach(item => {
                      response.data.data.forEach(fp => {
                        if(item.PostID === fp.postID){
                          item.isFP = 'true'
                        }
                      })
                      let liked = this.LikeCache.find(temp => temp.userID === uid && temp.PostID === item.PostID)
                      if (liked){
                        item.isLiked = true
                      }
                    })
                  })
            })
          })
          .catch(error => {
            console.log(error);
          })
      this.loadLikeCache()
    },
    ifshowPopup(){
      if(this.uid){
        this.showcommentText=true
      }
    },
    hidePopup() {
      setTimeout(() => {
        if (this.hideshowPopup) {
          this.showcommentText = false;
        }
      }, 200);
    },

    closeshowPopup(){
      this.showcommentText = false;
      this.hideshowPopup = true;
    },

    // 判断是否能发帖
    ifvisible(){
      if(this.uid){
        this.visible = true;
        PostEmail(this.uid).then(response=>{
          console.log(response)
          this.formData.userID = response.data.uid
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

    mypost(){
      if(this.userID){
        this.$router.push('/identity')
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


    mycollection(){
      if(this.userID){
        this.$router.push('/identity')
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


    showCommentSidebar(index,postid){
      this.postlist[index].commentSidebarVisible = true;
      Cookies.set('postID', postid)
    },

    SubmitComment,

    submitComment() {
      // 帖子评论提交
      if(this.uid){
        this.submittingComment = true;
        setTimeout(() => {
          this.commentList.push({
            id: this.commentList.length + 1,
            user: {
              username: '匿名用户',
              avatar: 'https://randomuser.me/api/portraits/lego/1.jpg',
            },
            time: new Date().toLocaleString(),
            content: this.commentForm.comment,
          });
          this.commentForm.comment = '';
          this.submittingComment = false;
        }, 1000);
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

    saveLikeCache(){
      let likeCacheJson = JSON.stringify(this.LikeCache)
      localStorage.setItem('likeCache', likeCacheJson)
    },

    // 点赞缓存池
    loadLikeCache(){
      let likeCacheJson = localStorage.getItem('likeCache')
      if(likeCacheJson){
        this.LikeCache = JSON.parse(likeCacheJson)
      }
    },

// 帖子收藏
    toggleColor(item) {
      if(this.uid){
        item.isFP = !item.isFP;
        PostEmail(this.uid).then(response=>{
          this.id = response.data.uid
          if(item.isFP){
            item.Collection += 1;
            AddUserFP(this.id, item.PostID)
            ElMessage.success({
              message: '收藏成功',
              type: 'success',
            })

          } else {
            item.Collection -= 1;
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

    beforeAvatarUpload:function(file){
      const isJPG = file.type === 'image/jpeg';
      const isPNG = file.type === 'image/png';
      const isPG = (isJPG || isPNG)                                       //限制图片格式为jpg / png
      const isLt100M = file.size / 1024 / 1024 < 100;                         //限制图片大小

      if (!isPG) {
        this.$message.error('上传头像图片只能是 JPG 或 PNG 格式!');
        return false;
      }
      if (!isLt100M) {
        this.$message.error('上传头像图片大小不能超过 100MB!');
        return false;
      }
      return true;
    },
    onSubmit,

    handlePictureCardPreview(file){
      this.dialogImageUrl = file.url
      this.dialogVisible = true
    },
    handleSuccess(response, file, fileList) {
      // response表示上传结果
      // file表示当前上传的文件
      // fileList表示当前成功上传的图片数组
      // 将图片保存
      // eslint-disable-next-line no-unused-vars
      const imageUrl = response.data.url // 获取上传成功的图片 URL
      // eslint-disable-next-line no-unused-vars
      const id = this.$route.params.id // 获取其他需要的数据，如当前用户 ID
      this.fileList = fileList
      this.form.images.push(response.data.url);
    },
    ValueChange(){
      this.formData.tag1 = this.$refs['Area'].getCheckedNodes()[0].pathLabels.join(',')
    },
    handleRemove(file,fileList){
      // file代表将要删除的文件
      // fileList为删除此图片之后剩下的图片列表
      // 更新
      console.log(file, fileList);
    },
    handlePhoto(file){
      console.log(file)
      this.formData.fileList.push(file)
    },


    expand(index) {
      this.postlist[index].expand = true;
    },
    noexpand(index) {
      this.postlist[index].expand = false;
    }
  },

}
</script>


<style >

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

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.社区页面 {
  height: 100%;
  /* background-image: linear-gradient(120deg, #e0c3fc 0%, #8ec5fc 100%); */
  background-image: url('../assets/communityBack.jpg');
  background-size: cover;
  position: relative;
  background-attachment: fixed;
  margin-top: 58px;
  /* background-repeat: no-repeat; */
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
  margin-left: 600px;
  display: flex;
  justify-content: space-between;
}

.点赞收藏评论标签{
  margin-top: 10px;
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

:root{
  --primary-color: royalblue;
}

.button{
  padding: 5px 16px;
  color: #000000d9;
  border: 1px solid #d9d9d9;
  /* background-color: transparent; */
  border-radius: 2px;
  line-height: 1.4;
  box-shadow: 0 2px #00000004;
  cursor: pointer;
  transform: scale(1);
  animation: jump 0s;
  transition: .3s;
}
.button:focus-visible{
  outline: 0;
}

.button[type="primary"]{
  border-color: #fff ;
  background-color: var(--primary-color);
  color: #fff;
  text-shadow: 0 -1px 0 rgb(0 0 0 / 12%);
  box-shadow: 0 2px #0000000b;
}
.button[type="dashed"]{
  border-style: dashed;
}
.button[type="text"]{
  /* border-color: transparent; */
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