<template>
    <el-table
            name="baseExtractTable"
            width="100%"
            :data="tableData">
        <el-table-column label="输出内容">
            <template slot-scope="scope">
                <el-input type="input" v-model="scope.row.value"></el-input>
            </template>
        </el-table-column>
        <el-table-column
                label="操作"
                width="50">
            <template  slot-scope="scope">
                <el-button type="danger" icon="el-icon-delete" circle size="mini" @click="$emit('removeRowByIndex', tableData, scope.$index)"></el-button>
            </template>
        </el-table-column>
    </el-table>
</template>

<script>
    export default {
        name: "base-Output-Table",
        props: {
            tableData: {
                require: true,
                type: Array,
            }
        },
        data: function () {
            return {

            }
        }
    }

</script>

<style scoped>

</style>