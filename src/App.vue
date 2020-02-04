<template>
  <div id="app">
    <el-container id="base-container" style="position: absolute;height: 100%;width: 100%">
      <el-header height="60px" style="width: 100%; padding:0px">
        <el-row id="header-div">
            <el-col :span="3">
              <router-link to="/">
                <el-image
                  style="width: 150px; height: 60px"
                  :src="url"
                  :fit="fit"></el-image>
              </router-link>
            </el-col>
            <el-col :span="6">
            <router-link to="/"><span style="font-size: 30px;"><span style="color: #000000;">接口测试平台</span></span></router-link>
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
                    <el-dropdown-item>退出</el-dropdown-item>
                  </el-dropdown-menu>
                </el-dropdown>
            </el-col>
        </el-row>
      </el-header>
      <el-container id="aside-container">
        <el-aside width="200px">
            <el-menu style="height: 100%;"
              default-active="1-1"
              class="el-menu-vertical-demo"
              background-color="#545c64"
              text-color="#fff"
              active-text-color="#ffd04b">
              <el-submenu index="1">
                <template slot="title">
                  <i class="el-icon-location"></i>
                  <span>基础信息</span>
                </template>
                <el-menu-item index="1-1">
                    <router-link to="/pro-index"><span class="link-span"><i class="el-icon-location" style=""></i>项目管理</span></router-link>
                </el-menu-item>
                <el-menu-item index="1-2">
                    <router-link to="/env-index"><span class="link-span"><i class="el-icon-location"></i>环境管理</span></router-link>
                </el-menu-item>
                <el-menu-item index="1-3">
                    <router-link to="/api-index"><span class="link-span"><i class="el-icon-location"></i>接口管理</span></router-link>
                </el-menu-item>
                <el-menu-item index="1-4">
                    <router-link to="/case-index"><span class="link-span"><i class="el-icon-location"></i>用例管理</span></router-link>
                </el-menu-item>
                <el-menu-item index="1-5">
                    <router-link to="/suite-index"><span class="link-span"><i class="el-icon-location"></i>用例集管理</span></router-link>
                </el-menu-item>
                <el-menu-item index="1-6">
                    <router-link to="/report-index"><span class="link-span"><i class="el-icon-location"></i>测试报告管理</span></router-link>
                </el-menu-item>
                <el-menu-item index="1-7">
                    <router-link to="/plan-index"><span class="link-span"><i class="el-icon-location"></i>执行计划管理</span></router-link>
                </el-menu-item>
                <el-menu-item index="1-8">
                    <router-link to="/mock-index"><span class="link-span"><i class="el-icon-location"></i>mock服务管理</span></router-link>
                </el-menu-item>
              </el-submenu>
              <el-submenu index="2">
                <template slot="title">
                  <i class="el-icon-location"></i>
                  <span>系统设置</span>
                </template>
                <el-menu-item index="2-1">
                    <router-link to="/public-func-index"><span class="link-span"><i class="el-icon-location"></i>公共函数管理</span></router-link>
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
                }]
        }
    },
    methods: {

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