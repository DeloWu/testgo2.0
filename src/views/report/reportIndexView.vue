<template>
    <el-container>
        <el-main>
            <el-row>
                <el-col :span="12">
                    <el-breadcrumb separator-class="el-icon-arrow-right">
                        <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
                        <el-breadcrumb-item :to="{ path: '/report-index' }">测试报告管理</el-breadcrumb-item>
                    </el-breadcrumb>
                </el-col>
                <el-col :span="12"></el-col>
            </el-row>
            <br>
            <el-row>
                <el-col :span="3">
                    <!-- <el-button type="primary" round @click="addPro">添加报告</el-button> -->
                </el-col>
                <el-col :span="14"></el-col>
                <el-col :span="4" :offset="18">
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
                        @cell-click="operate"
                        :default-sort = "{prop: 'reportCreateTime', order: 'descending'}">
                        <el-table-column
                          label="序号"
                          width="80"
                          type="index">
                        </el-table-column>
                        <el-table-column
                          show-overflow-tooltip
                          prop="reportCreateTime"
                          label="创建时间"
                          width="250"
                          sortable
                          :formatter="dateTimeFormater">
                        </el-table-column>
                        <el-table-column
                          show-overflow-tooltip
                          prop="reportRunTime"
                          label="执行时长"
                          width="250"
                          sortable
                          :formatter="runTimeFormat">
                        </el-table-column>
                        <el-table-column
                          show-overflow-tooltip
                          prop="reportUser"
                          label="执行者"
                          width="250"
                          sortable>
                        </el-table-column>
                        <el-table-column
                          fixed="right"
                          label="操作"
                          prop="reportIndex"
                          width="60"
                          class-name="open">
                          <el-button type="success" icon="el-icon-search" circle ></el-button>
                        </el-table-column>
                        <el-table-column
                          fixed="right"
                          label=""
                          prop="reportIndex"
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
                  :page-size="10"
                  layout="total, sizes, prev, pager, next, jumper"
                  :total="200">
                </el-pagination>
            </el-row>
        </el-footer>
    </el-container>
</template>

<script>
import { formatDate, RunTimeFormater } from "@utils/timeOperator"
export default {
    name: 'reportIndex',
    data() {
        return {
            input: '',
            currentPage: 1,
            tableData: [
                {
                    reportId: 11,
                    reportRunTime: 3600,
                    reportUser: "admin",
                    reportPath: "",
                    createTime: 1581135254,
                    modifyTime: 1581135255,
                },
                {
                    reportId: 22,
                    reportRunTime: 3601,
                    reportUser: "admin",
                    reportPath: "",
                    createTime: 1581135254,
                    modifyTime: 1581135255,
                },
                {
                    reportId: 33,
                    reportRunTime: 3602,
                    reportUser: "admin",
                    reportPath: "",
                    createTime: 1581135251,
                    modifyTime: 1581135251,
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
        addPro(){
            window.console.log("call add pro func");
        },
        operate(row, column, cell){
            if (cell.className.indexOf("open") >= 0) {
                window.console.log("call open func, cur row is: " + row.reportId);
            }else if(cell.className.indexOf("delete") >= 0){
                window.console.log("call delete func, cur row is: " + row.reportId);
                this.confirmDelete();
            }else {
                window.console.log("no func match, cur row is: " + row.reportId);
            }
        },
        // 时间戳转化为datetime格式
        dateTimeFormater(row) {
            try{
               return formatDate(row.createTime);
            } catch(e){
                window.console.log(e);
                return row.createTime
            }
        },
        // 3601s => 1h0m1s
        runTimeFormat(row) {
            try{
                return RunTimeFormater(row.reportRunTime)
            } catch(e){
                window.console.log(e);
                return row.reportRunTime
            }
        }
    }
}
</script>

<style scoped>
    

</style>