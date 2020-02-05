<template>
    <el-container>
        <el-main>
            <el-row>
                <el-col :span="12">
                    <el-breadcrumb separator-class="el-icon-arrow-right">
                        <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
                        <el-breadcrumb-item :to="{ path: '/api-index' }">接口管理</el-breadcrumb-item>
                    </el-breadcrumb>
                </el-col>
                <el-col :span="12"></el-col>
            </el-row>
            <br>
            <el-row>
                <el-col :span="3">
                    <el-button type="primary" round @click="addApi">添加接口</el-button>
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
                        @cell-click="operate">
                        <el-table-column
                          label="序号"
                          width="50"
                          type="index">
                        </el-table-column>
                        <el-table-column
                          :show-overflow-tooltip="true"
                          prop="apiName"
                          label="接口名称"
                          width="140">
                        </el-table-column>
                        <el-table-column
                          :show-overflow-tooltip="true"
                          prop="apiUrl"
                          label="接口路径"
                          width="240">
                        </el-table-column>
                        <el-table-column
                          prop="apiMethod"
                          label="请求方式"
                          width="80">
                        </el-table-column>
                        <el-table-column
                          :show-overflow-tooltip="true"
                          prop="apiRequestData"
                          label="请求参数">
                        </el-table-column>
                        <el-table-column
                          :show-overflow-tooltip="true"
                          prop="apiDesc"
                          label="接口描述">
                        </el-table-column>
                        <el-table-column
                          prop="relativePro"
                          label="所属项目">
                        </el-table-column>
                        <el-table-column
                          fixed="right"
                          label="操"
                          prop="apiIndex"
                          width="60"
                          class-name="edit">
                          <el-button type="success" icon="el-icon-edit" circle ></el-button>
                        </el-table-column>
                        <el-table-column
                          fixed="right"
                          label="作"
                          prop="apiIndex"
                          width="60"
                          class-name=”delete“>
                          <el-button type="danger" icon="el-icon-delete" circle></el-button>
                        </el-table-column>
                        <el-table-column
                          fixed="right"
                          label=""
                          prop="apiIndex"
                          width="60"
                          class-name=”run“
                          @click="dialogVisible = true">
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
          title="发送请求"
          :visible.sync="dialogVisible"
          width="30%">
          <span>TODO: 设计运行接口样式</span>
          <span slot="footer" class="dialog-footer">
            <el-button @click="dialogVisible = false">取 消</el-button>
            <el-button type="primary" @click="dialogVisible = false">运 行</el-button>
          </span>
        </el-dialog>
    </el-container>
</template>

<script>
export default {
    name: 'apiIndex',
    data() {
        return {
            dialogVisible: false,
            input: '',
            currentPage: 1,
            tableData: [
                {
                    apiIndex: 11,
                    apiName: "接口一",
                    apiUrl: "http://127.0.0.1:8080/",
                    apiMethod: "get",
                    apiRequestData: `{"foo": "bar"}`,
                    apiDesc: "接口描述一",
                    relativePro: 1,
                },
                {
                    apiIndex: 22,
                    apiName: "接口二",
                    apiUrl: "https://127.0.0.1:443/v1/foo/bar",
                    apiMethod: "post",
                    apiRequestData: `{"foo": "bar"}`,
                    apiDesc: "接口描述二",
                    relativePro: 1,
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
        addApi(){
            window.console.log("call add api func");
        },
        operate(row, column, cell){
            if (cell.className.indexOf("edit") >= 0) {
                window.console.log("call edit func, cur row is: " + row.apiIndex);
            }else if(cell.className.indexOf("delete") >= 0){
                window.console.log("call delete func, cur row is: " + row.apiIndex);
                this.confirmDelete();
            }else if(cell.className.indexOf("run") >= 0){
                this.dialogVisible = true;
                window.console.log("call run func, cur row is: " + row.apiIndex);
            }else {
                window.console.log("no func match, cur row is: " + row.apiIndex);
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