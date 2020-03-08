<template>
    <el-container>
        <el-main>
            <el-row>
                <el-col :span="12">
                    <el-breadcrumb separator-class="el-icon-arrow-right">
                        <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
                        <el-breadcrumb-item>编辑环境</el-breadcrumb-item>
                    </el-breadcrumb>
                </el-col>
                <el-col :span="12"></el-col>
            </el-row>
            <br>
            <el-row>
                <el-col :span="12">
                    <el-form ref="form" :model="form" :rules="rules" label-width="100px" label-position="left" size="medium">
                        <el-form-item label="环境名称:" prop="envName" required>
                            <el-input v-model="form.envName"></el-input>
                        </el-form-item>
                        <el-form-item label="环境ip:" prop="envIp" required>
                            <el-input v-model="form.envIp"></el-input>
                        </el-form-item>
                        <el-form-item label="环境端口:" prop="envPort" required>
                            <el-input v-model="form.envPort"></el-input>
                        </el-form-item>
                        <el-form-item label="关联项目:" required>
                            <el-select v-model="form.relativePro" placeholder="请选择关联项目" multiple>
                                <el-option :label="item.proName" :value="item.id" v-for="item in relativePros" :key="item.value"></el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="项目描述:" prop="envDesc">
                            <el-input type="textarea" v-model="form.envDesc"></el-input>
                        </el-form-item>
                        <el-form-item>
                            <el-button type="success" @click="save('form')">保存</el-button>
                            <el-button type="primary" @click="saveAndContinue('form')">保存并继续添加</el-button>
                            <el-button type="danger" @click="cancelSave()">取消</el-button>
                        </el-form-item>
                    </el-form>
                </el-col>
            </el-row>
        </el-main>
    </el-container>
</template>

<script>
    import {getProjectsByPagination} from "@api/project"
    import {getEnvironmentById, editEnvironment} from "@api/environment"
    export default {
        name: 'envUpdate',
        data() {
            return {
                form: {
                    // envName: "",
                    // envIp: "",
                    // envPort: "0",
                    // envDesc: "",
                    // relativePro: []
                },
                rules: {
                    envName: [
                        { required: true, message: '请输入环境名称', trigger: 'blur' },
                    ],
                    envIp: [
                        { required: true, message: '请输入环境ip', trigger: 'blur' },
                    ],
                    envPort: [
                        { required: true, message: '请输入环境端口', trigger: 'blur' },
                    ],

                },
                relativePros: []
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
                        //需要格式转换
                        let postForm = this.formTransform();
                        editEnvironment(postForm).then(response => {
                            const code = response.data.code;
                            if(code == 200){
                                this.successMessage();
                                this.$router.push('/env-index');
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
                        //需要格式转换
                        let postForm = this.formTransform();
                        editEnvironment(postForm).then(response => {
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
                this.$router.push('/env-index')
            },
            // 发送请求前的格式转换
            formTransform: function () {
                //深拷贝
                let postForm = JSON.parse(JSON.stringify(this.form));
                if(typeof postForm.relativePro == "string"){
                    postForm.relativePro = [postForm.relativePro]
                }
                return postForm
            }
        },
        mounted(){
            getProjectsByPagination({pageindex: 1, pagesize: 1000}).then(response => {
                this.relativePros = response.data.data
            });
            const updateId = this.$router.currentRoute.query.id;
            getEnvironmentById(updateId).then(response => {
                this.form = response.data.data;
                this.form.relativePro = this.form.relativePro[0];
            });
    }
}
</script>

<style scoped>


</style>