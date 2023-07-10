<template>
    <div class="personal-center">
      <div class="nav-bar">
        <el-menu :default-active="activeIndex" class="el-menu-vertical-demo" @select="handleSelect">
          <el-menu-item index="1"><i class="el-icon-edit"></i>修改资料</el-menu-item>
          <el-menu-item index="2"><i class="el-icon-message"></i>我的帖子</el-menu-item>
          <el-menu-item index="3"><i class="el-icon-star-on"></i>我的景点收藏</el-menu-item>
          <el-menu-item index="4"><i class="el-icon-star-off"></i>我的帖子收藏</el-menu-item>
          <el-menu-item index="5"><i class="el-icon-s-order"></i>我的订单</el-menu-item>
        </el-menu>
      </div>
      <div class="content">
        <div v-if="activeIndex === '1'" class="personal-info">
          <h2>个人资料</h2>
          <el-form :model="form" label-width="80px">
            <el-form-item label="头像">
              <el-avatar :src="form.avatar" size="large"></el-avatar>
            </el-form-item>
            <el-form-item label="昵称">
              <el-input v-model="form.nickname"></el-input>
            </el-form-item>
            <el-form-item label="密码">
              <el-input v-model="form.password" type="password"></el-input>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="submitForm">确定</el-button>
            </el-form-item>
          </el-form>
        </div>
        <div v-else class="default-info">
          <h2>个人资料</h2>
          <div class="info-item">
            <span class="label">头像：</span>
            <el-avatar :src="userInfo.avatar" size="large"></el-avatar>
          </div>
          <div class="info-item">
            <span class="label">昵称：</span>
            <span class="value">{{ userInfo.nickname }}</span>
          </div>
          <div class="info-item">
            <span class="label">密码：</span>
            <span class="value">******</span>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: 'PersonalCenter',
    data() {
      return {
        activeIndex: '1',
        form: {
          avatar: 'https://randomuser.me/api/portraits/men/1.jpg',
          nickname: 'John Doe',
          password: '123456'
        },
        userInfo: {
          avatar: 'https://randomuser.me/api/portraits/men/1.jpg',
          nickname: 'John Doe'
        }
      }
    },
    methods: {
      handleSelect(index) {
        this.activeIndex = index
      },
      submitForm() {
        // 提交表单数据
        console.log(this.form)
      }
    }
  }
  </script>
  
  <style>
  .personal-center {
    display: flex;
    flex-direction: row;
    height: 100%;
  }
  .nav-bar {
    width: 200px;
    background-color: #f0f0f0;
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
  </style>
  <!-- eslint-disable-next-line vue/multi-word-component-names -->
<template> 
    <div>
  
  <div class="header-channel-fixed-right">     <!-- 筛选地区 -->
  <el-checkbox
    :indeterminate="isIndeterminate"
    v-model="checkAll"
    @change="handleCheckAllChange"
    >全选</el-checkbox
  >
  <el-checkbox-group style="max-height: 50px"
    v-model="checkedCities"
    @change="handleCheckedCitiesChange"
  >
    <el-checkbox v-for="city in cities" :label="city" :key="city"
      >{{city}}</el-checkbox
    >
  </el-checkbox-group>
  
  </div> 
  
  
      <div class="search-wrapper"> 
        <input type="text" placeholder="请输入景区名称" v-model="searchText" />
        <input type="text" placeholder="请输入景区名称" />
        <select style="margin-right: 100px;" v-model="selectedType"> 
          <option value="">全部</option>
          <option value="名胜古迹">名胜古迹</option>
          <option value="博物馆">博物馆</option>
          <option value="山水圣地">自然风光</option>
        </select>
      </div>
  
      <!-- <div class="spots"> -->
        <el-card v-for="(spot, index) in visibleSpots"  :key="index" class="spot-wrapper">
          <div class="spot" style="height: 80px;width: 200px;">
            <el-link href="https://element.eleme.io" target="_blank">{{ spot.name }}</el-link>
            <!-- <h3>{{ spot.name }}</h3> -->
            <p>{{ spot.address }}</p>
            
            <div class="info">
              <span>分数{{ spot.score }}</span>
              <span>收藏数：{{ spot.collection}}</span>
            </div>
          </div>
        </el-card>
      <!-- </div> -->
  
      <div class="pagination">
        <button :disabled="currentPage === 1" @click="toPage(currentPage - 1)">上一页</button>
        <span>{{ currentPage }}/{{ totalPages }}</span>
        <button :disabled="currentPage === totalPages" @click="toPage(currentPage + 1)">下一页</button>
      </div>
      <el-backtop :right="50" :bottom="50" />
      <!-- <el-button type="primary" class="back-to-top" icon="el-Top" @click="scrollToTop"></el-button> -->
    </div>
    
  </template>
  
  <script>
  const cityOptions= ['黑龙江省','辽宁省','吉林省','河北省','河南省','湖北省','湖南省','山东省','山西省','陕西省','安徽省','浙江省','江苏省','福建省','广东省','海南省','四川省','云南省','贵州省','青海省','甘肃省','江西省','台湾省','内蒙古自治区','宁夏回族自治区','新疆维吾尔自治区','西藏自治区','广西壮族自治区','香港','澳门']
  // const cityOptions = ['上海', '北京', '广州', '深圳','青岛']
  export default {
    // eslint-disable-next-line vue/multi-word-component-names
    name:'Spot',
    data() {
      return {
        spots: [
          { address:"山东", name:'',comment:'',collection:'',id:11,picture:'',score:'',text:'',type:'名胜古迹'},
          { address:"山东", name:'',comment:'',collection:'',id:11,picture:'',score:'',text:'',type:'名胜古迹'},
          { address:"山东", name:'山东博物馆',comment:'',collection:'',id:11,picture:'',score:'',text:'',type:'博物馆'},
          { address:"山东", name:'',comment:'',collection:'',id:11,picture:'',score:'',text:'',type:'名胜古迹'},
          { address:"山东", name:'',comment:'',collection:'',id:11,picture:'',score:'',text:'',type:'名胜古迹'},
          { address:"山东", name:'',comment:'',collection:'',id:11,picture:'',score:5,text:'',type:'名胜古迹'},
          { address:"山东", name:'',comment:'',collection:'',id:11,picture:'',score:3.9,text:'',type:'名胜古迹'},
          { address:"山东", name:'',comment:'',collection:'',id:11,picture:'',score:4.2,text:'',type:'名胜古迹'},
          { address:"山东", name:'',comment:'',collection:100,id:11,picture:'',score:4.7,text:'',type:'名胜古迹'},
          { address:"广东", name:'鼎湖山',comment:'天然大氧吧',collection:99,id:11,picture:'',score:4.8,text:'',type:'自然风光'},
          { address:"山东", name:'',comment:'',collection:5,id:11,picture:'',score:5.0,text:'',type:'自然风光'},
        ],
        searchText: '',
        selectedType: '',
        pageSize: 10,
        currentPage: 1,
        checkAll: false,
        checkedCities: ['上海', '北京'],
        cities: cityOptions,
        isIndeterminate: true,
      }
    },
    computed: {
      visibleSpots() {
        let spots = this.spots.filter((spot) => {
          return (
            // spot.name.includes(this.searchText) &&
            // (this.selectedType === '' || spot.type === this.selectedType)
            this.selectedType === '' || spot.type === this.selectedType
          )
        })
        // eslint-disable-next-line vue/no-side-effects-in-computed-properties
        this.totalPages = Math.ceil(spots.length / this.pageSize)
        return spots.slice(
          (this.currentPage - 1) * this.pageSize,
          this.currentPage * this.pageSize
        )
      },
    },
    methods: {
      toPage(page) {
        if (page < 1 || page > this.totalPages) return
        this.currentPage = page
      },
      // 回到顶部函数
      scrollToTop() {
        window.scroll({ top: 0, left: 0, behavior: 'smooth' })
      },
  
      handleCheckAllChange(val) {
        this.checkedCities = val ? cityOptions : []
        this.isIndeterminate = false
      },
      handleCheckedCitiesChange(value) {
        let checkedCount = value.length
        this.checkAll = checkedCount === this.cities.length
        this.isIndeterminate =
          checkedCount > 0 && checkedCount < this.cities.length
      },
  
    },
    // created() {
    //   GetAllPosts()
    //       .then(response => {
    //         console.log(response.data)
    //         this.spots = response.data
    //       })
    //       .catch(error => {
    //         console.log(error);
    //       })
    // },
  }
  </script>
  
  <style>
  .el-card:nth-child(2n+1){
  margin-left: 50px;
  }
  .el-card:nth-child(2n){
  margin-left: 100px;
  }
  .spots {
    display: flex;
    flex-wrap: wrap;
    /* height: 50px;
    width: 150px; */
    justify-content: space-between;
  }
  
  .spot-wrapper {
    padding: 20px;
          background: rgba(255, 255, 255, 0.194);
          /* align-items: center; */
          display: flex;
          /* flex-direction: column; */
          justify-content: space-between;
    height: 80px;
    width: 200px;
    width: calc(50% - 10px);
    margin-bottom: 20px;
  }
  
  .spot {
    margin-top: 80px;
    padding: 10px;
    border: 1px solid #ccc;
  }
  
  .info {
    display: flex;
    justify-content: space-between;
    margin-top: 10px;
  }
  
  .pagination {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-top: 20px;
  }
  
  .back-to-top {
    position: fixed;
    bottom: 50px;
    right: 50px;
    width: 50px;
    height: 50px;
    border-radius: 50%;
    background-color: #ccc;
    background-image: url('https://cdn.jsdelivr.net/npm/vue-cli-service-global-style-loader/lib/style/back-to-top.png');
    background-repeat: no-repeat;
    background-size: 24px 24px;
    background-position: center;
    cursor: pointer;
  }
  
  .search-wrapper {
    display: flex;
    margin-top: 250px;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
  }
  
  .header-channel-fixed-right {
    /* display: inline-grid; */
    max-height: 70px;
    flex: 1;
    align-items: center;
    height: 95%;
    position: relative;
  }
  
  /* @media (max-width: 1099.9px)
  .header-channel-fixed-right .left-bottom, .header-channel-fixed-right-right, .header-channel-fixed-right-left .left-top, .header-channel-fixed-right {
    font-size: 13px;
    grid-gap: 14px 8px;
  } */
  
  </style>
  