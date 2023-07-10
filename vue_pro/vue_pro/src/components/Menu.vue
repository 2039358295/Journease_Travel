<template>
  <!-- 导航栏 -->
  <div class="affix-container">
    <el-menu  mode="horizontal" background-color="rgba(70, 76, 91, 0.6)" text-color="#ccc" style="border-radius: 10px;" class="app-menu" router>
      <!-- <el-menu-item index="/home" router-link to="/home">首页</el-menu-item> -->
      <!-- <div class="logostyle">Journease</div> -->
      <img src="../assets/name.png" style="height: 60px;">
      <el-menu-item index="/spot" router-link to="/spot" @click="clickspot()">景点</el-menu-item>
      <el-menu-item index="/community" router-link to="/community" @click="clickcommunity()">社区</el-menu-item>
      <el-menu-item index="/identity" router-link to="/identity" v-show="iflogin" @click="clickidentity()">个人中心</el-menu-item>
      <el-menu-item v-show="ifshow">
        <!-- <el-input type='text' class="search-input" placeholder="搜索" prefix-icon="Search" style="width: 200px; margin-top:13px;"></el-input> -->
        <el-input type='text' v-model="this.searchContent" @change="search"  class="w-50 m-2" placeholder="搜索" style="width: 200px; margin-top:13px;">
          <!-- <el-input type="text" v-model="keyword" placeholder="搜索" prefix-icon="el-icon-search" @keyup.enter.native="search" clearable> -->
          <template #prefix>
            <el-icon class="el-input__icon" @click="onsearch()"><search /></el-icon>
          </template>
        </el-input>
      </el-menu-item>
      <el-menu-item >
        <el-button v-show="!iflogin" type="text" @click="gotoLogin()" style="margin-top:13px">登录 / 注册</el-button>
        <el-button v-show="iflogin" type="text" @click="gotoidentity()" style="margin-top:13px">你好！{{ username }} ₍˄·͈༝·͈˄*₎◞ ̑̑</el-button>
        <!-- <div v-show="iflogin">你好，{{ username }}</div> -->
        <!-- showLoginDialog = true -->
      </el-menu-item>

      <el-menu-item @click="menuidentity()">
        <div class="avatar-container" @mouseenter="ifshowPopup()" @mouseleave="hidePopup()">
          <el-avatar :src="avatarUrl" :size="40" fit="cover"></el-avatar>
        </div>
      </el-menu-item>
    </el-menu>
    <div v-show="showPopup" class="avatar-popup-container" @mouseenter="hideshowPopup = false" @mouseleave="closeshowPopup()">
      <el-menu style="border-radius: 10px;" default-active="/spot" router>
        <el-menu-item  @click="goToPersonalCenter()" router-link to="/identity"><el-icon><edit /></el-icon>个人中心</el-menu-item>
        <el-menu-item @click="goToMypost()" ><el-icon><message /></el-icon>我的贴子</el-menu-item>
        <el-menu-item @click="goToMyfav()"><el-icon><star /></el-icon>我的贴子收藏</el-menu-item>
        <el-menu-item @click="goToViewfav()"><el-icon><star /></el-icon>我的景点收藏</el-menu-item>
        <el-menu-item @click="goToTick()"><el-icon><tickets /></el-icon>我的订单</el-menu-item>
        <el-menu-item @click="logout()" class="menu-item-center">
          <svg xmlns="http://www.w3.org/2000/svg" height="1em" viewBox="0 0 512 512" style="margin-right: 9px;margin-left: 8px;">! Font Awesome Free 6.4.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc.<path d="M502.6 278.6c12.5-12.5 12.5-32.8 0-45.3l-128-128c-12.5-12.5-32.8-12.5-45.3 0s-12.5 32.8 0 45.3L402.7 224 192 224c-17.7 0-32 14.3-32 32s14.3 32 32 32l210.7 0-73.4 73.4c-12.5 12.5-12.5 32.8 0 45.3s32.8 12.5 45.3 0l128-128zM160 96c17.7 0 32-14.3 32-32s-14.3-32-32-32L96 32C43 32 0 75 0 128L0 384c0 53 43 96 96 96l64 0c17.7 0 32-14.3 32-32s-14.3-32-32-32l-64 0c-17.7 0-32-14.3-32-32l0-256c0-17.7 14.3-32 32-32l64 0z"/></svg>
          退出登录
        </el-menu-item>
      </el-menu>
    </div>
  </div>
</template>

<script>
//  import router from '@/router/index.js'
import {getToken,removeToken} from "../components/token.js"
import {mapActions} from 'vuex'
import {PostAvatar, PostEmail} from "@/api/PostAPI";
import {GetPoster} from "@/api/GetAPI";

export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name:'Menu',
  data(){
    return{
      ifshow:true,
      searchContent:'',
      // username:'zlt',
      currentroute:'',
      username:'',
      iflogin:false,
      uid:'',
      hideshowPopup:true,
      showPopup: false,
      avatarUrl:"",
      input:'',
      showLoginDialog: false,
      showRegisterDialog: false,
      activeName: 'login',
      loginForm: {
        email: '',
        password: ''
      },
      loginRules: {
        email: [
          { required: true, message: '请输入邮箱', trigger: 'blur' },
          { type: 'email', message: '请输入正确的邮箱地址', trigger: ['blur', 'change'] }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' }
        ]
      },
      registerForm: {
        nickname: '',
        password: '',
        confirmPassword: '',
        email: '',
        activationCode: ''
      },
      registerRules: {
        nickname: [
          { required: true, message: '请输入昵称', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' }
        ],
        confirmPassword: [
          { required: true, message: '请再次输入密码', trigger: 'blur' },
          {
            validator: (rule, value, callback) => {
              if (value !== this.registerForm.password) {
                callback(new Error('两次输入的密码不一致'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        email: [
          { required: true, message: '请输入邮箱', trigger: 'blur' },
          { type: 'email', message: '请输入正确的邮箱地址', trigger: ['blur', 'change'] }
        ],
        activationCode: [
          { required: true, message: '请输入邮箱激活码', trigger: 'blur' }
        ],
      }
    }
  },
  methods: {//   ...mapActions(['setSearchContent']),
    ...mapActions(['updateSelectedTabIndex','setSearchContent']),
    onsearch(){
      const searchContent = this.searchContent;
      this.setSearchContent(searchContent);
    },
    goToPersonalCenter(){
      this.ifshow=false;
      this.$store.dispatch('updateSelectedTabIndex','1');
      this.$router.push('/identity');
    },
    goToMypost(){
      this.ifshow=false;
      this.$store.dispatch('updateSelectedTabIndex','2');
      this.$router.push('/identity');
    },
    goToMyfav(){
      this.ifshow=false;
      this.$store.dispatch('updateSelectedTabIndex','3');
      this.$router.push('/identity');
    },
    goToViewfav(){
      this.ifshow=false;
      this.$store.dispatch('updateSelectedTabIndex','4');
      this.$router.push('/identity');
    },
    goToTick(){
      this.ifshow=false;
      this.$store.dispatch('updateSelectedTabIndex','5');
      this.$router.push('/identity');
    },
    menuidentity(){
      if(this.uid){
        this.ifshow=false;
        this.$router.push('/identity')
      }
      else{
        this.$message.error('登录后查看个人中心')
      }
    },
    clickidentity(){
      this.ifshow=false;
    },
    clickspot(){
      this.ifshow=true;
    },
    clickcommunity(){
      this.ifshow=true;
    },
    closeshowPopup(){
      this.showPopup = false;
      this.hideshowPopup = true;
    },
    hidePopup() {
      setTimeout(() => {
        if (this.hideshowPopup) {
          this.showPopup = false;
        }
      }, 200);
    },
    gotoLogin() {
      this.$router.push({ name: 'login' })
    },
    gotoidentity(){
      this.$router.push({ name: 'identity' })
      this.currentindex = '/identity'
    },
    logout(){
      this.userID=removeToken()
      this.avatarUrl = require("../assets/logo.jpg")
      sessionStorage.clear()
      setTimeout(() => {
        this.$router.go(0)
      }, 0.00001)
      this.$router.push('/spot');
    },
    ifshowPopup(){
      if(this.uid){
        this.showPopup=true
      }
    },
  },
  // 和个人中心的请求一样
  created() {
    this.uid = getToken();
    if(getToken()){
      this.iflogin = true
    }
    console.log(this.userID);
    PostEmail(getToken())
        .then(response=>{
          GetPoster(response.data.uid)
              .then(response=>{
                console.log(response)
                this.username = response.data.name
              })
          PostAvatar(response.data.uid)
              .then(response=>{
                console.log(response)
                this.avatarUrl = URL.createObjectURL(new Blob([response.data],{type:"image/jpeg"}))
              })
        })
        .catch(err=>{
          console.log(err)
          let Spot = JSON.parse(sessionStorage.getItem('ID'))
          if(Spot !== null){
            GetPoster(Spot.id)
                .then(response=>{
                  this.username = response.data.name
                })
            PostAvatar(Spot.id)
                .then(response=>{
                  this.avatarUrl = URL.createObjectURL(new Blob([response.data],{type:"image/jpeg"}))
                })
          }
          else {
            this.username = "游客"
            this.avatarUrl = require("../assets/logo.jpg")
          }
          //this.avatarUrl=URL.Spot.ava;
        })
    // if(!this.uid){
    //   this.avatarUrl=require('../assets/logoname.jpg');
    // }
    // else{
    //   this.avatarUrl=''
    // }
    },
}
</script>

<style scoped>
.logostyle{
  font-weight: bold;
  margin-bottom: 20px;
  margin-top: 10px;
  margin-left: 5px;
  font-size: 30px;
  font-weight:80;
  /* background-image: linear-gradient(120deg,#429dda,#a1c4fd,#cc9ce0); */
  font-family: Forte;
  /* font-family: Avenir, Helvetica, Arial, sans-serif; */
}
.avatar-popup-container {
  position: absolute;
  top: 60px;
  right: 0;
  display: inline-block;
  background-color: #ffffff;
  border: 1px solid #ebeef5;
  /* box-shadow: 0 2px 12px 0 rgba(254, 247, 247, 0.1); */
  z-index: 2;
}
/* 导航栏始终置于页面顶部 */
.affix-container {
  text-align: center;
  border-radius: 4px;
  /*滚动广告幅*/
  /* padding-top: 100px;
  margin: 0 auto;
  width: 1920px;
  height: 711px;
  position: relative; */
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 999;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 60px;
  background-color: rgb(63, 63, 64);
}

.avatar-container {
  display: flex;
  align-items: center;
  margin-top: 10px;
}

.avatar-container:hover{
  transform: scale(1.6);
}

#app {
  /* font-family: Avenir, Helvetica, Arial, sans-serif; */
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  /* color: #2c3e50; */
  margin-top: 0px;
}

.app-menu {
  /* background-color: rgba(240, 237, 237, 0.4); */
  /* position: fixed; */
  top:0px;
  width: 100%;
  display: flex;
  justify-content: space-between;
}

</style>