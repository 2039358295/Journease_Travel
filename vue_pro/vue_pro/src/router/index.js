//用于创建整个应用的路由
import { createRouter, createWebHistory } from 'vue-router'

import Spot from "../components/Spot.vue";
import Community from "../components/Community.vue";
import Identity from "../components/Identity.vue";
import SpotDetail from "../components/SpotDetail.vue"
import Login from "../components/Login.vue"
import App from "../App.vue"



const routes = [
    { path: "/", redirect:'/spot'},
    {
      path: '/spotdetail',
      name: 'SpotDetail',
      component: SpotDetail,
    },

    { path: "/spot", 
      name:'spot',
      component: Spot ,
      },
      {
        path:"/app",
        name:"app",
        component:App,
      },
    { path: "/community", 
      name:'community',
      component: Community },
    { path: "/identity", 
      name:'identity',
      component: Identity },
    { path: "/login",
      name:'login', 
      component: Login },
    // { path: "/post",name:'post', component: Post }
  ];

  const router = createRouter({
    routes:routes,
    history:createWebHistory(process.env.BASE_URL)
})

export default router