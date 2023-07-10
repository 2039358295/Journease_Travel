<template>
<Menu></Menu>



    <router-view></router-view>

  
</template>

<script>
// import { ref, getCurrentInstance } from 'vue';
import Menu from './components/Menu.vue'
// import { ElButton, ElDialog } from 'element-plus';
// import 'element-plus/lib/theme-chalk/index.css';
// import Login from './components/Login.vue'
// import Community from './components/Community.vue'
// import Spot from './components/Spot.vue'
// import Home from './components/Home.vue'
// import Login from './components/Login.vue'
// import Register from './components/Register.vue'
// import Post from './components/Post.vue'
// import SignIn from './components/SignIn.vue'
// import { ref } from 'vue'
export default {
  name: 'App',
  components: {
   
    Menu,
},
}
</script>
  