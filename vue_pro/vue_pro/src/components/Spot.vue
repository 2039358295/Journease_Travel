<!-- eslint-disable-next-line vue/multi-word-component-names -->
<template>
  <div class="景区缩略图页面">
    <!-- <img class="background" src="require('../assets/SpotBackground.jpg')"> -->
    <!-- <div class="header-channel-fixed-right">     -->
    <!-- 轮播图 -->
    <el-carousel :interval="4000" type="card" height="300px" style="margin-bottom: 10px;">
      <el-carousel-item v-for="(item, index) in spotpics" :key="index">
        <img :src="item.picture" alt="spot image" style="width:700px;max-height: 100%; max-width: 100%;margin-top: -14px;"/>
        <div class="carousel-caption" style="margin-block-end:15px;padding-bottom: 1px;">
          <h3 style="margin-bottom: 0px;">{{ item.name }}</h3>
          <div class="rating">
            <el-rate
                v-model="item.score"
                disabled :show-score="false"
            >
            </el-rate>
          </div>
          <p style="margin-top: 0px;">{{ item.comment }}</p>
        </div>
      </el-carousel-item>
    </el-carousel>
    <!-- 筛选地区和类型 -->
    <div>
      <el-collapse class="折叠面板" v-model="isCollapse">
        <el-collapse-item title="选择省份">
          <el-checkbox v-model="isSelectAll" @change="selectAll">{{ selectAllText }}</el-checkbox>
          <el-checkbox-group v-model="selectedProvinces" @change="handleCheckedCitiesChange">
            <el-checkbox v-for="province in provinces" :key="province" :label="province">{{ province }}</el-checkbox>
          </el-checkbox-group>
        </el-collapse-item>
        <el-collapse-item title="选择景区类型">
          <el-checkbox v-model="isSelectAllTypes" @change="selectAllTypes">{{ selectAllTypesText }}</el-checkbox>
          <el-checkbox-group v-model="selectedTypes" @change="handleSelectChange">
            <el-checkbox v-for="spottype in spottypes" :key="spottype" :label="spottype">{{ spottype }}</el-checkbox>
          </el-checkbox-group>
        </el-collapse-item>
      </el-collapse>
    </div>

    <div>
      <el-row gutter="20" style="flex-wrap: wrap;">
        <el-col :span="12" v-for="spot in spots.slice((currentPage-1)*pageSize, currentPage*pageSize)" :key="spot.id" >
          <el-card :body-style="{ padding: '0px' }" class="景点信息" style="height: 220px;width: 480px;" >
            <div style="display: flex;">
              <!-- <div class="spot" style="height: 100px;width: 400px;"> -->
              <div class="photo" style="display:inline-block;" @click="goToSpot(spot.id,spot.collection,spot.name)">
                <img :src="spot.picture" style="height: 180px;width: 180px;margin-left: 10px;margin-top: 10px"/>
              </div>
              <div class="info" style="margin-left: 10px; flex-grow: 1;">
                <div class="name" style="margin-bottom: 15px" @click="goToSpot(spot.id,spot.collection,spot.name)">{{ spot.name }}</div>
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
                  <el-button class="like-style"  size="middle" v-bind:style="{backgroundColor: spot.collection ? '#fab4eb' : 'white'}" style="z-index: 1;margin-right: 10px"  circle  v-on:click="toggleLike(spot)" >
                    <img :src="require('@/assets/heart.png')" style="height: 20px;width: 23px; z-index: 3;">
                  </el-button>
                  <span>{{spot.collections}}</span>
                </div>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>

    <!-- 翻页 -->

    <div style=" margin-top: 20px;margin-left: 535px;margin-right: 550px;">
      <el-pagination v-model:current-page="currentPage" :page-size="pageSize" layout="prev, pager, next" :total="spots.length"></el-pagination>
    </div>
    <el-backtop :right="30" :bottom="50" />
    <!-- <el-button type="primary" class="back-to-top" icon="el-Top" @click="scrollToTop"></el-button> -->
  </div>
  <!-- </div> -->
</template>

<script>
import {getSpotData, Search} from '../components/SpotAPI.js'
import { FilterSpotByAddress,CreateSpot1,DeleteSpot1,Display} from '../components/SpotAPI.js'
import {ElMessage} from "element-plus";
import { ElMessageBox } from 'element-plus'
import { getToken } from './token';
import {mapGetters} from "vuex";

const cityOptions= ['黑龙江省','辽宁省','吉林省','河北省','河南省','湖北省','湖南省','山东省','山西省','陕西省','安徽省','浙江省','江苏省','福建省','广东省','海南省','四川省','云南省','贵州省','青海省','甘肃省','江西省','台湾省','内蒙古自治区','宁夏回族自治区','新疆维吾尔自治区','西藏自治区','广西壮族自治区','香港','澳门']
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name:'Spot',
  components:{
    // SpotDetail,
  },
  data() {
    return {
      selectedProvinces: [],
      isCollapse: true,
      isSelectAll: false,
      isSelectAllTypes: false,
      selectedTypes:[],
      // checkedCities:[],
      // checkAll:false,
      // isIndeterminate:false,
      // provinces : ['黑龙江省','辽宁省','吉林省','河北省','河南省','湖北省','湖南省','山东省','山西省','陕西省','安徽省','浙江省','江苏省','福建省','广东省','海南省','四川省','云南省','贵州省','青海省','甘肃省','江西省','台湾省','内蒙古自治区','宁夏回族自治区','新疆维吾尔自治区','西藏自治区','广西壮族自治区','香港','澳门'],
      spottypes:['名胜古迹', '博物馆', '自然风光', '地标景观'],
      spots: [],
      spotpics:[],
      searchText: '',
      selectedType: '',
      pageSize: 10,
      currentPage: 1,
      provinces:cityOptions,
      uid:808,
    }
  },
  computed: {
    ...mapGetters(["getSearchContent"]),
    selectAllText() {
      if (this.isSelectAll) {
        return '取消全选'
      } else {
        return '全选'
      }
    },
    selectAllTypesText() {
      return this.isSelectAllTypes ? '取消全选' : '全选';
    },
  },
  // 显示景区介绍内容
  directives: {
    'truncate-text': {
      mounted(el, binding) {
        const maxLength = binding.value;
        const content = el.textContent;
        if (content.length > maxLength) {
          el.textContent = content.slice(0, maxLength) + '...';
        }
      }
    }
  },

  methods: {
    searchHandler() {
      const searchContent = this.getSearchContent;
      var params1 = new URLSearchParams();
      params1.append('token', getToken())
      params1.append("name", searchContent);

      Search(params1).then(response => {
        this.spots = response.data.data;
      })
          .catch(error => {
            console.log(error);
          });
    },
    //翻页函数
    toPage(page) {
      if (page < 1 || page > this.totalPages) return
      this.currentPage = page
    },

    // 地区选择处理函数
    handleCheckedCitiesChange() {
      const checkedCount = this.selectedProvinces.length
      this.checkAll = checkedCount === this.provinces.length
      this.isIndeterminate = checkedCount > 0 && checkedCount < this.provinces.length

      // 多选框改变，向后端发送请求
      var params = new URLSearchParams();
      params.append('token', getToken())
      params.append("address", this.selectedProvinces);
      params.append('type',this.selectedTypes)

      FilterSpotByAddress(params).then(response => {
        this.spots = response.data.data;
      })
          .catch(error => {
            console.log(error);
          });

    },


    toggleCollapse() {
      this.isCollapse = !this.isCollapse
    },

    // 全选的处理函数
    selectAll() {
      if (this.isSelectAll) {
        // 向后端发送多选框内容
        this.selectedProvinces = this.provinces
        var params = new URLSearchParams();
        params.append('token', getToken())
        params.append("address", this.selectedProvinces);
        params.append('type',this.selectedTypes)

        FilterSpotByAddress(params).then(response => {
          this.spots = response.data.data;
        })
            .catch(error => {
              console.log(error);
            });
      } else {
        this.selectedProvinces = []
      }
    },

    selectAllTypes() {
      if (this.isSelectAllTypes) {
        this.selectedTypes = this.spottypes

        var params = new URLSearchParams();
        params.append('token', getToken())
        params.append("address", this.selectedProvinces);
        params.append('type',this.selectedTypes)

        FilterSpotByAddress(params).then(response => {
          this.spots = response.data.data;
        })
            .catch(error => {
              console.log(error);
            });
      } else {
        this.selectedTypes = []
      }
    },

    // 景区类型选择改变，向后端发送请求
    handleSelectChange() {
      // 当选择的值发生变化时，向后端发送选中的值并请求数据
      var params = new URLSearchParams();
      params.append('token', getToken())
      params.append("address", this.selectedProvinces);
      params.append('type',this.selectedTypes)

      FilterSpotByAddress(params).then(response => {
        this.spots = response.data.data;
      })
          .catch(error => {
            console.log(error);
          });
    },

    // 进入景区详情页的函数
    goToSpot(spotid,c,s) {
      if (getToken()==null){
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
      else{
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
      }

    },

    // 收藏景点函数
    toggleLike(item) {
      if( getToken()!=null){
        item.collection = !item.collection;
        var params = new URLSearchParams();
        params.append('token', getToken())
        params.append("sid", item.id);

        if(item.collection) {
          item.collections=item.collections*10/10 + 1;

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

  },


  created() {
    this.$watch(()=>this.getSearchContent,this.searchHandler)
    // 将uid传到Spot页面
    let User = JSON.parse(sessionStorage.getItem('userID'))
    if (User) {
      this.uid = User.user
    }
    var params2 = new URLSearchParams()
    params2.append('token', getToken())

    // 得到景区数据
    getSpotData(params2).then(response => {
      console.log(response.data)
      this.spots = response.data.data
      console.log(response.data);
    })
        .catch(error => {
          console.log(error);
        })

    // 得到展示景区数据
    Display(params2).then(response => {
      this.spotpics = response.data.data
    })
        .catch(error => {
          console.log(error);
        })
  },
}
</script>

<style scoped>

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
  flex-direction: column;
  justify-content: space-between;
  opacity: 0.85;
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

.景区缩略图页面{
  padding: 20px;
  /* text-align: center; */
  /* align-items: center; */
}


.折叠面板{
  color:lightskyblue;
  opacity: 0.8;
}

</style>