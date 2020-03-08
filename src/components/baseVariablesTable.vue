<template>
    <el-table
            name="baseVariablesTable"
            width="100%"
            :data="tableData">
        <el-table-column label="键名">
            <template slot-scope="scope">
                <el-input type="input" v-model="scope.row.keyName"></el-input>
            </template>
        </el-table-column>
        <el-table-column label="健值类型">
            <template slot-scope="scope">
                <el-select v-model="scope.row.valueType" placeholder="请选择数据类型" >
                    <el-option :label="item.label" :value="item.value" v-for="item in valueTypes" :key="item.value"></el-option>
                </el-select>
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
        name: "base-Variables-Table",
        props: {
            tableData: {
                require: true,
                type: Array,
            }
        },
        data: function () {
            return {
                valueTypes: [{label: "string", value: "string"}, {label: "int", value: "int"}, {label: "float", value: "float"}, {label: "bool", value: "bool"}, {label: "list", value: "list"}, {label: "dict", value: "dict"}, {label: "null", value: "null"},],
            }
        },

    }

</script>

<style scoped>

</style>