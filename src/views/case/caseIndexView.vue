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
                    <router-link to="/case-add"><el-button type="primary" round>添加用例</el-button></router-link>
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
                            width="350"
                            :formatter="caseStepsFormatter">
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
    import config from "@config"
    import {getProjectsByPagination} from "@api/project"
    import {getCaseTotal, getCasesByPagination, deleteCaseById} from "@api/case"
    export default {
    name: 'caseIndex',
    data() {
        return {
            dialogVisible: false,
            input: '',
            currentPage: 1,
            currentPageSize: config.pageSize,
            total: 0,
            deleteId: 0,
            tableData: [

            ],
            filterPros: [],
            options: [

            ]
        }
    },
    methods: {
        fetchData(){
            //    获取分页数据
            getCasesByPagination({pageindex: this.currentPage, pagesize: this.currentPageSize}).then(response => {
                this.tableData = response.data.data
            });
            getCaseTotal().then(response => {
                this.total = response.data.data.total
            });
            getProjectsByPagination({pageindex: 1, pagesize: 1000}).then(response => {
                this.options = response.data.data;
            });
        },
        search(){
            window.console.log('call search func');
        },
        handleSizeChange(val) {
            this.currentPageSize = val;
            getCasesByPagination({pageindex: this.currentPage, pagesize: this.currentPageSize}).then(response => {
                this.tableData = response.data.data
            });
        },
        handleCurrentChange(val) {
            this.currentPage = val;
            getCasesByPagination({pageindex: this.currentPage, pagesize: this.currentPageSize}).then(response => {
                this.tableData = response.data.data
            });
        },
        confirmDelete() {
            this.$confirm('此操作将永久删除该文件, 是否继续?', '提示', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'warning'
            }).then(() => {
              deleteCaseById(this.deleteId);
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
                this.$router.push({path:"/case-update", query:{id:row.id}});
            }else if(cell.className.indexOf("delete") >= 0){
                this.deleteId = row.id;
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
        },
        caseStepsFormatter(row){
            return "TODO: 转换步骤内容"
        }
    },
        created() {
            //获取数据,渲染页面
            this.fetchData();
        }
    }
</script>

<style scoped>

</style>