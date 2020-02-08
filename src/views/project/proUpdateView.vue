<template>
    <el-container>
        <el-main>
            <el-row>
                <el-col :span="12">
                    <el-breadcrumb separator-class="el-icon-arrow-right">
                        <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
                        <el-breadcrumb-item :to="{ path: '/pro-add' }">添加项目</el-breadcrumb-item>
                    </el-breadcrumb>
                </el-col>
                <el-col :span="12"></el-col>
            </el-row>
            <br>
            <el-row>
                <el-col :span="12">
                    <el-form ref="form" :model="form" label-width="100px" label-position="top">
                      <el-form-item label="项目名称:" required>
                        <el-input v-model="form.proName"></el-input>
                      </el-form-item>
                      <el-form-item label="项目描述:">
                        <el-input type="textarea" v-model="form.proDesc"></el-input>
                      </el-form-item>
                      <el-form-item>
                        <el-button type="success" @click="save">保存</el-button>
                        <el-button type="primary" @click="saveAndContinue">保存并继续添加</el-button>
                        <el-button type="danger" @click="cancelSave">取消</el-button>
                      </el-form-item>
                    </el-form>
                </el-col>
            </el-row>
        </el-main>
    </el-container>
</template>

<script>
export default {
    name: 'proUpdate',
    data() {
        return {
            form: {
                proName: "",
                proDesc: ""
            }
        }
    },
    methods: {
        save: function(){
            window.console.log("保存并返回/pro-index")
        },
        saveAndContinue(){
            window.console.log("保存并返回/pro-add")
        },
        cancelSave(){
            window.console.log("不保存并返回/pro-index")
        }
    },
    mounted() {
        window.console.log("向后端请求数据");
        this.form = {
                proName: "项目一",
                proDesc: "项目一描述"
            }
    }
}
</script>

<style scoped>
    

</style>