<template>
    <el-container>
        <el-main>
            <el-row>
                <el-col :span="12">
                    <el-breadcrumb separator-class="el-icon-arrow-right">
                        <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
                        <el-breadcrumb-item>编辑项目</el-breadcrumb-item>
                    </el-breadcrumb>
                </el-col>
                <el-col :span="12"></el-col>
            </el-row>
            <br>
            <el-row>
                <el-col :span="12">
                    <el-form ref="form" :model="form" :rules="rules" label-width="100px" label-position="left" size="medium">
                        <el-form-item label="项目名称:" prop="proName" required>
                            <el-input v-model="form.proName"></el-input>
                        </el-form-item>
                        <el-form-item label="项目描述:" prop="proDesc">
                            <el-input type="textarea" v-model="form.proDesc"></el-input>
                        </el-form-item>
                        <el-form-item>
                            <el-button type="success" @click="save('form')">保存</el-button>
                            <el-button type="primary" @click="saveAndContinue('form')">保存并继续添加</el-button>
                            <el-button type="danger" @click="cancelSave">取消</el-button>
                        </el-form-item>
                    </el-form>
                </el-col>
            </el-row>
        </el-main>
    </el-container>
</template>

<script>
    import {getProjectById, editProject} from "@api/project"

    export default {
    name: 'proUpdate',
    data() {
        return {
            form: {
                // proName: "",
                // proDesc: ""
            },
            rules: {
                proName: [
                    { required: true, message: '请输入项目名称', trigger: 'blur' },
                ]
            },
        }
    },
    methods: {
        resetForm(formName) {
            this.$refs[formName].resetFields();
        },
        successMessage(){
            this.$message({
                message: '数据更新成功',
                type: 'success'
            });
        },
        failMessage(){
            this.$message({
                message: '数据更新失败',
                type: 'error'
            });
        },
        save: function(formData){
            this.$refs[formData].validate((valid) => {
                if (valid) {
                    window.console.log(formData);
                    window.console.log(this.form);
                    editProject(this.form).then(response => {
                        const code = response.data.code;
                        if(code == 200){
                            this.successMessage();
                            this.$router.push('/pro-index');
                        }else{
                            this.failMessage();
                        }
                    });
                } else {
                    window.console.log('表单格式校验失败!');
                    return false;
                }
            });
        },
        saveAndContinue(formData){
            this.$refs[formData].validate((valid) => {
                if (valid) {
                    editProject(this.form).then(response => {
                        const code = response.data.code;
                        if(code == 200){
                            this.successMessage();
                            this.form = {};
                        }else{
                            this.failMessage();
                        }
                    });
                } else {
                    window.console.log('表单格式校验失败!');
                    return false;
                }
            });
        },
        cancelSave(){
            this.$router.push('/pro-index')
        }
    },
    created() {
        const updateId = this.$router.currentRoute.query.id;
        getProjectById(updateId).then(response => {
            this.form = response.data.data
        });
    }
}
</script>

<style scoped>
    

</style>