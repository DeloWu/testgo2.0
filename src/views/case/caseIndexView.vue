<template>
    <el-container>
        <el-main>
            <el-row>
                <el-col :span="12">
                    <el-breadcrumb separator-class="el-icon-arrow-right">
                        <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
                        <el-breadcrumb-item :to="{ path: '/case-index' }">用例管理</el-breadcrumb-item>
                    </el-breadcrumb>
                </el-col>
                <el-col :span="12"></el-col>
            </el-row>
            <br>
            <el-row>
                <el-col :span="3">
                    <el-button type="primary" round @click="addCase">添加用例</el-button>
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
                        show-overflow-tooltip
                        :height="480"
                        style="width: 100%"
                        @cell-click="operate">
                        <el-table-column
                          label="序号"
                          width="50"
                          type="index">
                        </el-table-column>
                        <el-table-column
                          show-overflow-tooltip
                          prop="caseName"
                          label="用例名称"
                          width="140">
                        </el-table-column>
                        <el-table-column
                          show-overflow-tooltip
                          prop="caseSteps"
                          label="用例步骤"
                          width="350">
                        </el-table-column>
                        <el-table-column
                          show-overflow-tooltip
                          prop="caseDesc"
                          label="接口描述">
                        </el-table-column>
                        <el-table-column
                          fixed="right"
                          label="操"
                          prop="caseIndex"
                          width="60"
                          class-name="edit">
                          <el-button type="success" icon="el-icon-edit" circle ></el-button>
                        </el-table-column>
                        <el-table-column
                          fixed="right"
                          label="作"
                          prop="caseIndex"
                          width="60"
                          class-name="delete">
                          <el-button type="danger" icon="el-icon-delete" circle></el-button>
                        </el-table-column>
                        <el-table-column
                          fixed="right"
                          label=""
                          prop="caseIndex"
                          width="60"
                          class-name="run">
                          <el-button type="primary" icon="el-icon-caret-right" circle ></el-button>
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
        <el-dialog
          title="执行用例"
          :visible.sync="dialogVisible"
          width="30%">
          <span>TODO: 设计运行用例样式</span>
          <span slot="footer" class="dialog-footer">
            <el-button @click="dialogVisible = false">取 消</el-button>
            <el-button type="primary" @click="dialogVisible = false">运 行</el-button>
          </span>
        </el-dialog>
    </el-container>
</template>

<script>
export default {
    name: 'caseIndex',
    data() {
        return {
            dialogVisible: false,
            input: '',
            currentPage: 1,
            tableData: [
                {
                    caseIndex: 11,
                    caseName: "用例一",
                    caseSteps: `[{"1":"api"}, {"4":"testcase"}]`,
                    relativePro: [11, 22],
                    createTime: 1581135254,
                    modifyTime: 1581135255,
                    caseDesc: "用例描述一",
                    configVariables: [{"expected_status_code": 200}],
                    configBaseUrl: "http://127.0.0.1:5000",
                    configOutput: ["session_token"],
                    configVerify: false,
                    extract: {"apiId": [{"extractName":"extractValue"}, {"session_token":"content.token"}],"apiId2": [{"extractName":"extractValue"}, {"session_token":"content.token"}]},
                    validate: {"apiId": [["eq", "status_code", "$expected_status_code"], ["eq", "content.headers.Host", "httpbin.org"]],"apiId2": [["eq", "status_code", "$expected_status_code"], ["eq", "content.headers.Host", "httpbin.org"]],
                    variables: {"apiId": [{"expected_status_code": 200}], "apiId2": [{"expected_status_code": 200}]}}
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
        confirmDelete() {
            this.$confirm('此操作将永久删除该文件, 是否继续?', '提示', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'warning'
            }).then(() => {
              this.$message({
                type: 'success',
                message: '删除成功!'
              });
            }).catch(() => {
              this.$message({
                type: 'info',
                message: '已取消删除'
              });
            });
          },
        addCase(){
            window.console.log("call add Case func");
        },
        operate(row, column, cell){
            if (cell.className.indexOf("edit") >= 0) {
                window.console.log("call edit func, cur row is: " + row.caseIndex);
            }else if(cell.className.indexOf("delete") >= 0){
                window.console.log("call delete func, cur row is: " + row.caseIndex);
                this.confirmDelete();
            }else if(cell.className.indexOf("run") >= 0){
                this.dialogVisible = true;
                window.console.log("call run func, cur row is: " + row.caseIndex);
            }else {
                window.console.log("no func match, cur row is: " + row.caseIndex);
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