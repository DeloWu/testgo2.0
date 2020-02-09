<template >
    <div id="app">
        <template v-if="templateSeen">
            <div id="app-container">
                <el-container id="base-container" style="position: absolute;height: 100%;width: 100%">
                  <el-header height="50px" style="width: 100%; padding:0px">
                    <el-row id="header-div">
                        <el-col :span="3">
                          <router-link to="/">
                            <el-image
                              style="width: 150px; height: 52px"
                              :src="url"
                              :fit="fit"></el-image>
                          </router-link>
                        </el-col>
                        <el-col :span="6">
                        <router-link to="/"><span style="font-size: 30px;"><span style="color: rgb(51, 122, 175);">接口测试平台</span></span></router-link>
                        <!-- <router-link to="/"><span class="link-span" style="font-size: 30px;">接口测试平台</span></router-link> -->
                        </el-col>
                        <el-col :span="12"></el-col>
                        <el-col :span="3" :offset="12">
                            <el-dropdown>
                              <span class="el-dropdown-link">
                                <el-avatar shape="circle" size="large" src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"></el-avatar>
                                <span>欢迎，admin</span>
                              </span>
                              <el-dropdown-menu slot="dropdown">
                                <el-dropdown-item>
                                    <router-link to="/login"><span style="color: #606266;">退出</span></router-link>
                                </el-dropdown-item>
                              </el-dropdown-menu>
                            </el-dropdown>
                        </el-col>
                    </el-row>
                  </el-header>
                  <el-container id="aside-container">
                    <el-aside width="200px">
                        <el-menu style="height: 100%;"
                          :default-openeds="['baseInfo', 'systemSetting']"
                          class="el-menu-vertical-demo"
                          background-color="#545c64"
                          text-color="#fff"
                          active-text-color="#ffd04b"
                          router
                          >
                          <el-submenu index="baseInfo">
                            <template slot="title">
                              <i class="el-icon-location"></i>
                              <span>基础信息</span>
                            </template>
                            <el-menu-item index="/pro-index">
                                <i class="el-icon-location" style=""></i>项目管理
                            </el-menu-item>
                            <el-menu-item index="/env-index">
                                <i class="el-icon-location"></i>环境管理
                            </el-menu-item>
                            <el-menu-item index="/api-index">
                                <i class="el-icon-location"></i>接口管理
                            </el-menu-item>
                            <el-menu-item index="/case-index">
                                <i class="el-icon-location"></i>用例管理
                            </el-menu-item>
                            <el-menu-item index="/suite-index">
                                <i class="el-icon-location"></i>用例集管理
                            </el-menu-item>
                            <el-menu-item index="/report-index">
                                <i class="el-icon-location"></i>测试报告管理
                            </el-menu-item>
                            <el-menu-item index="/plan-index">
                                <i class="el-icon-location"></i>执行计划管理
                            </el-menu-item>
                            <el-menu-item index="/mock-server-index">
                                <i class="el-icon-location"></i>mock服务管理
                            </el-menu-item>
                          </el-submenu>
                          <el-submenu index="systemSetting">
                            <template slot="title">
                              <i class="el-icon-location"></i>
                              <span>系统设置</span>
                            </template>
                            <el-menu-item index="/public-func-index">
                                <i class="el-icon-location"></i>公共函数管理
                            </el-menu-item>
                          </el-submenu>
                        </el-menu>
                    </el-aside>
                    <!-- 其他页面注入 -->
                    <router-view></router-view>
                    <!-- <el-container>
                        <router-view></router-view>
                      <el-main>
                      </el-main>
                      <el-footer height="80px"></el-footer>
                    </el-container> -->
                  </el-container>
                </el-container>
              </div>
        </template>
        <template v-else>
            <div id="app-container">
                <router-view></router-view>
            </div>
        </template>
    </div>
</template>

<script>
export default {
    name: 'App',
    data() {
        return {
                fit: 'fill',
                url: require('../static/img/testGo-logo5.png'),
                tableData: [{
                    pro_counts: 0,
                    env_counts: 0,
                    api_counts: 0,
                    case_counts: 0,
                    suite_counts: 0,
                    report_counts: 0,
                    plan_counts: 0,
                    mock_counts: 0
                }],
                // 不需要公共页头页尾的页面路径
                noHeaderRouterPaths: ["/login"]
        }
    },
    methods: {

    },
    computed: {
        templateSeen: function(){
            return this.noHeaderRouterPaths.indexOf(this.$route.path) < 0 
        }
    }
}
</script>

<style scoped>
    .link-span {
        /* color:rgb(255, 208, 75) */
        color:rgb(255, 255, 255)
    }
    a {
      text-decoration: none;
    }
     
    .router-link-active {
      text-decoration: none;
    }
</style>