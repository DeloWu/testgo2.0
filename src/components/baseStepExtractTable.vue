<template>
    <el-table
            name="baseStepExtractTable"
            width="100%"
            :data="tableData">
        <el-table-column label="已选中步骤">
            <template slot-scope="scope">
                <el-select v-model="scope.row.id" placeholder="请选择项目">
                    <el-option
                            v-for="item in choicedSteps"
                            :key="item.id"
                            :label="item.apiName"
                            :value="item.id">
                    </el-option>
                </el-select>
            </template>
        </el-table-column>
        <el-table-column label="键名">
            <template slot-scope="scope">
                <el-input type="input" v-model="scope.row.keyName"></el-input>
            </template>
        </el-table-column>
        <el-table-column label="键值">
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
        name: "base-Step-Extract-Table",
        props: {
            tableData: {
                require: true,
                type: Array,
            },
            choicedSteps: {
                require: true,
                type: Array,
            },

        },
        data: function () {
            return {

            }
        },

    }

</script>

<style scoped>

</style>