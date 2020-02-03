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
                          prop="pro_counts"
                          label="项目数量"
                          width="100">
                        </el-table-column>
                        <el-table-column
                          prop="env_counts"
                          label="环境数量"
                          width="100">
                        </el-table-column>
                        <el-table-column
                          prop="api_counts"
                          label="接口数量"
                          width="100">
                        </el-table-column>
                        <el-table-column
                          prop="case_counts"
                          label="用例数量"
                          width="100">
                        </el-table-column>
                        <el-table-column
                          prop="suite_counts"
                          label="用例集数量"
                          width="100">
                        </el-table-column>
                          <el-table-column
                          prop="report_counts"
                          label="报告数量"
                          width="100">
                          </el-table-column>
                          <el-table-column
                          prop="plan_counts"
                          label="执行计划数量"
                          width="120">
                          </el-table-column>
                        <el-table-column
                          prop="mock_counts"
                          label="mock服务数量">
                        </el-table-column>
                    </el-table>
                </el-col>
            </el-row>
            <el-rol>
                <el-col :span="24">
                    <h3>近期用例执行情况</h3><br>
                    <div>
                        <H1>TODO: 数据可视化表格</H1>
                    </div>
                </el-col>
            </el-rol>
        </el-main>
          <el-footer height="80px"></el-footer>
    </el-container>
</template>

<script>
export default {
    name: 'App',
    data() {
        return {
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
    

</style>