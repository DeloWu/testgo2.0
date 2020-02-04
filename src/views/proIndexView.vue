<template>
    <el-container>
        <el-main>
            <el-row>
                <el-col :span="12">
                    <el-breadcrumb separator-class="el-icon-arrow-right">
                        <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
                        <el-breadcrumb-item :to="{ path: '/pro-index' }">项目管理</el-breadcrumb-item>
                    </el-breadcrumb>
                </el-col>
                <el-col :span="12"></el-col>
            </el-row>
            <br>
            <el-row>
                <el-col :span="3">
                    <el-button type="primary" round @click="addPro">添加项目</el-button>
                </el-col>
                <el-col :span="14"></el-col>
                <el-col :span="6" :offset="13">
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
                        @cell-click="oprPro">
                        <el-table-column
                          label="序号"
                          width="80"
                          type="index">
                        </el-table-column>
                        <el-table-column
                          prop="proName"
                          label="项目名称"
                          width="250">
                        </el-table-column>
                        <el-table-column
                          prop="proDesc"
                          label="项目描述">
                        </el-table-column>
                        <el-table-column
                          fixed="right"
                          label="操作"
                          prop="proIndex"
                          width="60">
                          <el-button type="success" icon="el-icon-edit" circle ></el-button>
                        </el-table-column>
                        <el-table-column
                          fixed="right"
                          label=""
                          prop="proIndex"
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
                    proIndex: 11,
                    proName: "项目一",
                    proDesc: "项目描述一",
                },
                {
                    proIndex: 22,
                    proName: "项目二",
                    proDesc: "项目描述二",
                },
                {
                    proIndex: 33,
                    proName: "项目三",
                    proDesc: "项目描述三",
                },
                {
                    proIndex: 44,
                    proName: "项目四",
                    proDesc: "项目描述四",
                },
                {
                    proIndex: 22,
                    proName: "项目二",
                    proDesc: "项目描述二",
                },
                {
                    proIndex: 33,
                    proName: "项目三",
                    proDesc: "项目描述三",
                },
                {
                    proIndex: 44,
                    proName: "项目四",
                    proDesc: "项目描述四",
                },
                {
                    proIndex: 22,
                    proName: "项目二",
                    proDesc: "项目描述二",
                },
                {
                    proIndex: 33,
                    proName: "项目三",
                    proDesc: "项目描述三",
                },
                {
                    proIndex: 44,
                    proName: "项目四",
                    proDesc: "项目描述四",
                }
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
        addPro(){
            window.console.log("call add pro func");
        },
        oprPro(row, column){
            if (column.id.indexOf("column_4") >= 0) {
                window.console.log("call edit func, cur row is: " + row.proIndex);
            }else if(column.id.indexOf("column_5") >= 0){
                window.console.log("call delete func, cur row is: " + row.proIndex);
            }else {
                window.console.log("no func match, cur row is: " + row.proIndex);
            }
        },
    }
}
</script>

<style scoped>
    

</style>