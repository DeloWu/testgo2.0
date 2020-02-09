<template>
    <el-container>
        <el-main>
            <el-row>
                <el-col :span="24">
                    <el-table
                    :data="tableData"
                    border
                    style="width: 80%">
                        <el-table-column
                          prop="proCounts"
                          label="项目数量"
                          width="100">
                        </el-table-column>
                        <el-table-column
                          prop="envCounts"
                          label="环境数量"
                          width="100">
                        </el-table-column>
                        <el-table-column
                          prop="apiCounts"
                          label="接口数量"
                          width="100">
                        </el-table-column>
                        <el-table-column
                          prop="caseCounts"
                          label="用例数量"
                          width="100">
                        </el-table-column>
                        <el-table-column
                          prop="suiteCounts"
                          label="用例集数量"
                          width="100">
                        </el-table-column>
                          <el-table-column
                          prop="reportCounts"
                          label="报告数量"
                          width="100">
                          </el-table-column>
                          <el-table-column
                          prop="planCounts"
                          label="执行计划数量"
                          width="120">
                          </el-table-column>
                        <el-table-column
                          prop="mockCounts"
                          label="mock服务数量">
                        </el-table-column>
                    </el-table>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="24">
                    <h3>近期用例执行情况</h3><br>
                    <div>
                        <H1>TODO: 数据可视化表格</H1>
                    </div>
                </el-col>
            </el-row>
        </el-main>
          <el-footer height="80px"></el-footer>
    </el-container>
</template>

<script>
export default {
    name: 'index',
    data() {
        return {
                tableData: [{
                    proCounts: 0,
                    envCounts: 0,
                    apiCounts: 0,
                    caseCounts: 0,
                    suiteCounts: 0,
                    reportCounts: 0,
                    planCounts: 0,
                    mockCounts: 0
                }]
        }
    },
    methods: {

    }
}
</script>

<style scoped>
    

</style>