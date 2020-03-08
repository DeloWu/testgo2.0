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
                    <router-link to="/env-add"><el-button type="primary" round>添加环境</el-button></router-link>
                </el-col>
                <!-- <el-col :span="5"></el-col> -->
                <el-col :span="4" :offset="11">
                    <el-select v-model="filterPros" multiple placeholder="请选择项目">
                        <el-option
                          v-for="item in options"
                          :key="item.id"
                          :label="item.proName"
                          :value="item.id">
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
                          width="80"
                          type="index">
                        </el-table-column>
                        <el-table-column
                          prop="envName"
                          label="环境名称"
                          width="150">
                        </el-table-column>
                        <el-table-column
                          show-overflow-tooltip
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
                          show-overflow-tooltip
                          prop="envDesc"
                          label="环境描述">
                        </el-table-column>
                        <el-table-column
                          prop="relativePro"
                          label="所属项目"
                          :formatter="relativeProFormatter">
                        </el-table-column>
                        <el-table-column
                          fixed="right"
                          label="操作"
                          prop="envIndex"
                          width="60"
                          class-name="edit">
                          <el-button type="success" icon="el-icon-edit" circle ></el-button>
                        </el-table-column>
                        <el-table-column
                          fixed="right"
                          label=""
                          prop="envIndex"
                          width="60"
                          class-name=”delete“>
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
                  :page-size="currentPageSize"
                  layout="total, sizes, prev, pager, next, jumper"
                  :total="total">
                </el-pagination>
            </el-row>
          </el-footer>
    </el-container>
</template>

<script>
    import config from "@config"
    import {getProjectsByPagination} from "@api/project"
    import {getEnvironmentTotal, getEnvironmentsByPagination, deleteEnvironmentById} from "@api/environment"
    export default {
    name: 'envIndex',
    data() {
        return {
            input: '',
            currentPage: 1,
            currentPageSize: config.pageSize,
            total: 0,
            deleteId: 0,
            tableData: [

            ],
            filterPros: [],
            options: [

            ],
        }
    },
    methods: {
        fetchData(){
            getProjectsByPagination({pageindex: 1, pagesize: 1000}).then(response => {
                this.options = response.data.data;
            });
            //    获取分页数据
            getEnvironmentsByPagination({pageindex: this.currentPage, pagesize: this.currentPageSize}).then(response => {
                this.tableData = response.data.data
            });
            getEnvironmentTotal().then(response => {
                this.total = response.data.data.total
            });
        },
        search(){
            window.console.log('call search func');
        },
        handleSizeChange(val) {
            this.currentPageSize = val;
            getEnvironmentsByPagination({pageindex: this.currentPage, pagesize: this.currentPageSize}).then(response => {
                this.tableData = response.data.data
            });
        },
        handleCurrentChange(val) {
            this.currentPage = val;
            getEnvironmentsByPagination({pageindex: this.currentPage, pagesize: this.currentPageSize}).then(response => {
                this.tableData = response.data.data
            });
        },
        confirmDelete() {
            this.$confirm('此操作将永久删除该文件, 是否继续?', '提示', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'warning'
            }).then(() => {
                deleteEnvironmentById(this.deleteId);
                this.$message({
                    type: 'success',
                    message: '删除成功!'
                });
                this.deleteId = 0;
                // 刷新数据
                this.fetchData()
            }).catch(() => {
              this.$message({
                type: 'info',
                message: '已取消删除'
              });
            });
          },
        operate(row, column, cell){
            if (cell.className.indexOf("edit") >= 0) {
                this.$router.push({path:"/env-update", query:{id:row.id}});
            }else if(cell.className.indexOf("delete") >= 0){
                this.deleteId = row.id;
                this.confirmDelete();
            }else {
                window.console.log("no func match, cur row is: " + row.envIndex);
            }
        },
        filter(){
            window.console.log("call filter func");
        },
        relativeProFormatter(row){
            return this.id_proName_map[row.relativePro[0]]
        }
    },
        computed: {
            id_proName_map: function () {
                // 用于所属项目的id和名称映射
                let map = {};
                for (const item of this.options){
                    map[item.id] = item.proName;
                }
                return map
            }
        },
    created() {
        //获取数据,渲染页面
        this.fetchData();
    },
}
</script>

<style scoped>
    

</style>