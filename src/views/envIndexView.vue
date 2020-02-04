<template>
    <el-container>
        <el-main>
            <el-row>
                <el-col :span="12">
                    <el-breadcrumb separator-class="el-icon-arrow-right">
                        <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
                        <el-breadcrumb-item :to="{ path: '/env-index' }">环境管理</el-breadcrumb-item>
                    </el-breadcrumb>
                </el-col>
                <el-col :span="12"></el-col>
            </el-row>
            <br>
            <el-row>
                <el-col :span="3">
                    <el-button type="primary" round @click="addEnv">添加环境</el-button>
                </el-col>
                <!-- <el-col :span="5"></el-col> -->
                <el-col :span="4" :offset="11">
                    <el-select v-model="filterPros" multiple placeholder="请选择项目">
                        <el-option
                          v-for="item in options"
                          :key="item.value"
                          :label="item.label"
                          :value="item.value">
                        </el-option>
                    </el-select>
                </el-col>
                <el-col :span="1">
                    <el-button icon="el-icon-search" circle @click="filter"></el-button>
                </el-col>
                <el-col :span="4" :offset="0">
                    <el-input
                      placeholder="请输入内容搜索"
                      v-model="input"
                      clearable>
                    </el-input>
                </el-col>
                <el-col :span="1">
                    <el-button icon="el-icon-search" circle @click="search"></el-button>
                </el-col>
            </el-row>
            <el-row>
                <el-col>
                    <el-table
                        :data="tableData"
                        stripe
                        :height="480"
                        style="width: 100%"
                        @cell-click="oprEnv">
                        <el-table-column
                          label="序号"
                          width="80"
                          type="index">
                        </el-table-column>
                        <el-table-column
                          prop="envName"
                          label="环境名称"
                          width="150">
                        </el-table-column>
                        <el-table-column
                          prop="envIp"
                          label="域名/IP"
                          width="240">
                        </el-table-column>
                        <el-table-column
                          prop="envPort"
                          label="端口"
                          width="80">
                        </el-table-column>
                        <el-table-column
                          prop="envDesc"
                          label="环境描述">
                        </el-table-column>
                        <el-table-column
                          prop="relativePro"
                          label="所属项目">
                        </el-table-column>
                        <el-table-column
                          fixed="right"
                          label="操作"
                          prop="envIndex"
                          width="60">
                          <el-button type="success" icon="el-icon-edit" circle ></el-button>
                        </el-table-column>
                        <el-table-column
                          fixed="right"
                          label=""
                          prop="envIndex"
                          width="60">
                          <el-button type="danger" icon="el-icon-delete" circle></el-button>
                        </el-table-column>
                        <el-table-column label="" width="100" fixed="right"></el-table-column>

                    </el-table>
                </el-col>
            </el-row>
        </el-main>
          <el-footer height="80px">
            <el-row>
                <el-pagination
                  @size-change="handleSizeChange"
                  @current-change="handleCurrentChange"
                  :current-page="currentPage"
                  :page-sizes="[10, 20, 50, 100]"
                  :page-size="10"
                  layout="total, sizes, prev, pager, next, jumper"
                  :total="200">
                </el-pagination>
            </el-row>
          </el-footer>
    </el-container>
</template>

<script>
export default {
    name: 'proIndex',
    data() {
        return {
            input: '',
            currentPage: 1,
            tableData: [
                {
                    envIndex: 11,
                    envName: "环境一",
                    envIp: "1.1.1.1",
                    envPort: 8080,
                    envDesc: "环境描述一",
                    relativePro: 1,
                },
                {
                    envIndex: 22,
                    envName: "环境二",
                    envIp: "https://www.baidu.com",
                    envPort: 443,
                    envDesc: "环境描述二",
                    relativePro: 2,
                },
                {
                    envIndex: 33,
                    envName: "环境二",
                    envIp: "http://www.bejson.com",
                    envPort: 80,
                    envDesc: "环境描述三",
                    relativePro: 3,
                },
            ],
            filterPros: [],
            options: [{
                value: 1,
                label: "项目一"
            },{
                value: 2,
                label: "项目二"
            },
            {
                value: 3,
                label: "项目三"
            },
            ]
        }
    },
    methods: {
        search(){
            window.console.log('call search func');
        },
        handleSizeChange(val) {
            window.console.log(`每页 ${val} 条`);
        },
        handleCurrentChange(val) {
            window.console.log(`当前页: ${val}`);
        },
        addEnv(){
            window.console.log("call add env func");
        },
        oprEnv(row, column){
            if (column.id.indexOf("column_7") >= 0) {
                window.console.log("call edit func, cur row is: " + row.envIndex);
            }else if(column.id.indexOf("column_8") >= 0){
                window.console.log("call delete func, cur row is: " + row.envIndex);
            }else {
                window.console.log("no func match, cur row is: " + row.envIndex);
            }
        },
        filter(){
            window.console.log("call filter func");
        }
    }
}
</script>

<style scoped>
    

</style>