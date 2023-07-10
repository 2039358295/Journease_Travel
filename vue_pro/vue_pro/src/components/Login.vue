<style scoped>
.form-input .el-input__inner {
  border: none; /* 移除默认边框 */
  border-radius: 0; /* 移除圆角 */
  border-bottom: 1px solid #ccc; /* 添加底部横线 */
}

* {
  box-sizing: border-box;
}
.headstyle {
  display: flex;
  margin-top: -60px;
  align-items: center;
  /* justify-content: center;  */
  margin-left: 30px;
  margin-bottom: 20px;
}

body {
  font-family: 'Montserrat',sans-serif;
  /* background-image: linear-gradient(120deg,#a1c4fd,#cc9ce0); */
  background-image: linear-gradient(120deg,#429dda,#a1c4fd,#cc9ce0);
  /* background-image: linear-gradient(120deg, #a1c4fd 0%, #c2e9fb 100%); */
  display: flex;
  flex-direction: column;
  justify-content: center;
  /* align-items: center; */
  height: 100vh;
  margin: -20px 0 50px;
}

h1 {
  font-weight: bold;
  margin-bottom: 20px;
  margin-left: 5px;
  font-size: 38px;
  font-weight:80;
  font-family: Forte;
}

.登录注册{
  font-family: 华文新魏;
  font-weight: bold;
  font-size: 35px;
}

p {
  font-size: 16px;
  line-height: 30px;
  /* letter-spacing: .5px; */
  margin: 20px 0 30px;
  font-family: 华文新魏;
}
span {
  font-size: 12px;
}

a {
  color: #333;
  font-size: 14px;
  text-decoration: none;
  margin: 15px 0;
}
.container {
  background: #fbfbfb;
   border-radius: 10px;
   box-shadow: 0 14px 28px rgba(0, 0, 0, .25), 0 10px 10px rgba(0, 0, 0, .22);
  position: relative;
    overflow: hidden;
   /* margin: auto; */
   width: 768px;
  max-width: 100%;
  min-height: 480px;
  margin: auto;
   margin-top: 170px;
}
.form-container form {
  background: #fff;
  display: flex;
  flex-direction: column;
  padding: 0 50px;
  height: 100%;
  justify-content: center;
  text-align: center;
}
.social-container {
  margin: 20px 0;
}

.social-container a {
  border: 1px solid #ddd;
  border-radius: 50%;
  display: inline-flex;
  justify-content: center;
  align-items: center;
  margin: 0 5px;
  height: 40px;
  width: 40px;
}

.social-container a:hover {
  background-color: #eee;

}

.txtb {
  /* border-bottom: 2px solid #adadad; */
  position: relative;
  margin: 10px 0;
}

.txtb input {
  font-size: 15px;
  color: #333;
  border: none;
  width: 100%;
  outline: none;
  background: none;
  padding: 0 3px;
  height: 35px;
}

.txtb span::before {
  content: attr(data-placeholder);
  position: absolute;
  top: 50%;
  left: 5px;
  color: #adadad;
  transform: translateY(-50%);
  transition: .5s;
}
.txtb span::after {
  content: '';
  position: relative;
  left: 0%;
  top: 100%;
  width: 0%;
  height: 2px;
  background-image: linear-gradient(120deg,#429dda,#a1c4fd,#cc9ce0);
  transition: .5s;
}

.input_text{
  border-radius: 20px;
  height: 25px;
  width: 215px;
}

.focus + span::before {
  top: -5px;
}

.focus + span::after {
  width: 100%;
}

button {
  border-radius: 20px;
  border: 1px solid #ff4b2b;
  background: #ff4b2b;
  color: #fff;
  font-size: 12px;
  font-weight: bold;
  padding: 12px 45px;
  letter-spacing: 1px;
  text-transform: uppercase;
  transition: transform 80ms ease-in;
  cursor: pointer;
}

button:active {
  transform: scale(.95);
}

button:focus {
  outline: none;
}

button.ghost {
  background: transparent;
  border-color: #fff;
}

.form-container {
  position: absolute;
  top: 0;
  height: 100%;
  transition: all .6s ease-in-out;
}

.form-container button {
  /* background: linear-gradient(120deg,#4B7FA2,#d4e5ef); */
  background-image: linear-gradient(120deg,#429dda,#a1c4fd,#cc9ce0);
  border: none;
  background-size: 200%;
  color: #fff;
  transition: .5s;
}

.form-container button:hover {
  background-position: right;
}



.sign-in-container {
  left: 0;
  width: 50%;
  z-index: 2;
}
.sign-in-container form a {
  position: relative;
  left: 100px;
}
.sign-up-container {
  left: 0;
  width: 50%;
  z-index: 1;
  opacity: 0;
}
.sign-up-container button {
  margin-top: 20px;
}
.overlay-container {
  position:absolute;
  top: 0;
  left: 50%;
  width: 50%;
  height: 100%;
  overflow: hidden;
  transition: transform .6s ease-in-out;
  z-index: 100;
}

.overlay {
  /* background-image: linear-gradient(120deg,#4B7FA2,#d4e5ef); */
  /* background-image: linear-gradient(120deg,#429dda,#cc9ce0); */
  background-image: linear-gradient(120deg,#429dda,#a1c4fd,#cc9ce0);
  /* background-image: linear-gradient(120deg, #a1c4fd 0%, #c2e9fb 100%); */
  color: #fff;
  position: relative;
  left: -100%;
  height: 100%;
  width: 200%;
  transform: translateY(0);
  transition: transform .6s ease-in-out;
}
.overlay-panel {
  position: absolute;
  top: 0;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 0 40px;
  height: 100%;
  width: 50%;
  text-align: center;
  transform: translateY(0);
  transition: transform .6s ease-in-out;
}
.overlay-right {
  right: 0;
  transform: translateY(0);

}

.overlay-left {
  transform: translateY(-20%);
}

.container.right-panel-active .sign-in-container {
  transform: translateY(100%);
}

.container.container.right-panel-active .overlay-container {
  transform: translateX(-100%);
}
.container.right-panel-active .sign-up-container {
  transform: translateX(100%);
  opacity: 1;
  z-index: 5;
}
.container.container.right-panel-active .overlay {
  transform: translateX(50%);
}
.container.container.right-panel-active .overlay-left {
  transform: translateY(0);
}
.container.container.right-panel-active .overlay-right {
  transform: translateY(20%);
}
#logoName{
  font-family: 'Lobster', cursive;
}
.登录注册页面{
  background-image: url('../assets/Iden.jpg');
  position: relative;
  overflow: hidden;
  height: 100vh;
}

</style>
<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <div class="登录注册页面" >
    <div class="container" :class="{ 'right-panel-active': isPanelActive }">
      <div class="form-container sign-up-container">
        <!-- <form> -->
        <el-form ref="registerForm" :model="registerForm" :rules="registerRules" label-width="80px" style="margin-top: 10px;">
          <div class="headstyle" >
            <img id="logo" src='../assets/logoname.jpg' style="display: block;margin-left: -10px;width: 130px;height: 110px;">
            <div class="登录注册">注册</div>
          </div>
          <div >
            <el-form-item label="昵称" icon="el-icon-user" prop="nickname">
              <el-input class="input_text" v-model="registerForm.nickname"></el-input>
            </el-form-item>
            <el-form-item label="邮箱" prop="email">
              <el-input class="input_text" v-model="registerForm.email"></el-input>
            </el-form-item>
            <el-form-item label="密码" prop="password">
              <el-input class="input_text" type="password" v-model="registerForm.password"></el-input>
            </el-form-item>
            <el-form-item label="确认密码" prop="confirmPassword">
              <el-input class="input_text" type="password" v-model="registerForm.confirmPassword"></el-input>
            </el-form-item>
          </div>
          <button type="button" style="width: 280px;margin-left: 20px;" @click="handleRegister">注册</button>
        </el-form>
      </div>

      <div class="form-container sign-in-container">
        <!-- <form action="#"> -->
        <el-form  ref="loginForm" :model="loginForm" :rules="loginRules" label-width="80px;margin-left: -10px;" >
          <div class="headstyle" >
            <img id="logo" src='../assets/logoname.jpg' style="display: block;margin-left: -15px;width: 130px;height: 110px;">
            <div class="登录注册">登录</div>
          </div>

          <div >
            <el-form-item label="邮箱" prop="email" style="margin-top: 20px;">
              <el-input class="input_text" type="email" v-model="loginForm.email" ></el-input>
            </el-form-item>
          </div>
          <div >
            <el-form-item label="密码" prop="password">
              <el-input class="input_text" type="password" v-model="loginForm.password"></el-input>
            </el-form-item>
          </div>
          <!-- <a href="#">忘记密码？</a> -->
          <button type="button" @click="handleLogin()" style="margin-top: 25px;">登录</button>
          <!-- </form> -->
        </el-form>
      </div>

      <div class="overlay-container">
        <div class="overlay">
          <div class="overlay-panel overlay-left">
            <h1>Hello  Traveler!</h1>

            <p>立即登录Jouernease，开启您的专属旅程吧</p>
            <button class="ghost" @click="isPanelActive = false">登录</button>
          </div>
          <div class="overlay-panel overlay-right">
            <h1>Welcome  to Journease !</h1>
            <!-- <p>Don't have an account yet? Click the 'Register' button to sign up and join our amazing journey!</p> -->
            <p>立即注册加入我们，一起开启奇妙的旅程吧</p>
            <button class="ghost" @click="isPanelActive = true">注册</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {LoginJudge} from '../components/LoginAPI.js'
import {RegisterJudge} from '../components/LoginAPI.js'
import { setToken } from './token'
import { ElMessage } from 'element-plus'

export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name:'Login',
  data() {
    return {
      uid: '202',
      isPanelActive: false,
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
  methods: {
    // 点击登录
    handleLogin() {
      //向后端发送登录
      var params1=new URLSearchParams();
      params1.append('email',this.loginForm.email);
      params1.append('password',this.loginForm.password);
      LoginJudge(params1).then(response=>{
        if(response.data.code===200){
          ElMessage.success("登录成功！");
          this.uid=response.data.uid;
          this.ava=response.data.ava;
          this.userName=response.data.userName;
          setToken(response.data.token);
          sessionStorage.setItem(
              'ID',
              JSON.stringify(
                  {
                    id:response.data.uid,
                    name:response.data.userName,
                  }
              )
          )


          setTimeout(() => {
            this.$router.go(0)
            this.$forceUpdate()
          }, 0.00001)
          console.log("点了登陆了！")
          this.$router.push('/spot');
        }
        else{
          let a=response.data.msg
          ElMessage.error(a);
        }
        // setTimeout(() => {
        //   this.$router.go(0)
        // }, 0.00001)
      })
          .catch(error=>{
            console.log(error);
          });
    },
    handleRegister(){
      var params=new URLSearchParams()
      params.append('nickname',this.registerForm.nickname)
      params.append('password',this.registerForm.password)
      params.append('email',this.registerForm.email)
      params.append('password1',this.registerForm.confirmPassword)

      RegisterJudge(params).then(response=>{
        if(response.data.code==200){
          ElMessage.success("注册成功！");
        }
        else{
          let a=response.data.msg
          ElMessage.error(a);
        }
        // setTimeout(() => {
        //   this.$router.go(0)
        // }, 0.00001)
      })
          .catch(error=>{
            console.log(error);
          });
    }
  }
}
</script>