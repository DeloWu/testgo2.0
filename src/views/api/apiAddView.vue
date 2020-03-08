<template>
    <el-container>
        <el-main>
            <el-row>
                <el-col :span="12">
                    <el-breadcrumb separator-class="el-icon-arrow-right">
                        <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
                        <el-breadcrumb-item>添加接口</el-breadcrumb-item>
                    </el-breadcrumb>
                </el-col>
                <el-col :span="12"></el-col>
            </el-row>
            <br>
            <el-row>
                <el-col :span="12">
                    <el-form ref="form" :model="form" :rules="rules" label-width="100px" label-position="top" size="mini">
                        <el-form-item label="接口名称:" prop="apiName" required>
                            <el-input v-model="form.apiName"></el-input>
                        </el-form-item>
                        <el-form-item  label="base-url:" prop="baseUrl">
                            <el-select
                                v-model="form.apiBaseUrl"
                                filterable
                                allow-create
                                default-first-option
                                placeholder="请选择或者新增url"
                                style="width:592px;">
                                <el-option
                                    v-for="item in apiBaseUrls"
                                    :key="item.envIp + ':' + item.envPort"
                                    :label="item.envIp + ':' + item.envPort"
                                    :value="item.envIp + ':' + item.envPort">
                                </el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="接口uri:" prop="apiPath" required>
                            <el-input v-model="form.apiPath"></el-input>
                        </el-form-item>
                        <el-form-item  label="请求方法:" prop="apiMethod" required>
                            <el-radio v-model="form.apiMethod" label="get">get</el-radio>
                            <el-radio v-model="form.apiMethod" label="post">post</el-radio>
                            <el-radio v-model="form.apiMethod" label="put">put</el-radio>
                            <el-radio v-model="form.apiMethod" label="patch">patch</el-radio>
                            <el-radio v-model="form.apiMethod" label="delete">delete</el-radio>
                            <el-radio v-model="form.apiMethod" label="head">head</el-radio>
                            <el-radio v-model="form.apiMethod" label="options">options</el-radio>
                        </el-form-item>
                        <el-form-item  label="Content-Type:" prop="apiContentType" required style="width:800px;">
                            <el-radio v-model="form.apiContentType" label="application/json">application/json</el-radio>
                            <el-radio v-model="form.apiContentType" label="application/x-www-form-urlencoded">application/x-www-form-urlencoded</el-radio>
                            <el-radio v-model="form.apiContentType" label="text/html; charset=UTF-8">text/html; charset=UTF-8(get方法选用)</el-radio>
                            <el-radio v-model="form.apiContentType" label="text/plain">text/plain</el-radio>
                            <el-radio v-model="form.apiContentType" label="application/xml">application/xml</el-radio>
                            <el-radio v-model="form.apiContentType" label="multipart/form-data">multipart/form-data</el-radio>
                        </el-form-item>
                        <el-form-item label="关联项目:" required>
                            <el-select v-model="form.relativePro" placeholder="请选择关联项目" multiple style="width: 300px;">
                                <el-option :label="item.proName" :value="item.id" v-for="item in relativePros" :key="item.value"></el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="接口请求头:">
                            <el-button type="primary" @click="addTableRow('apiHeaders')">添加</el-button>
                            <base-variables-table :tableData="form.apiHeaders" @removeRowByIndex="removeHeadersByIndex"></base-variables-table>
                        </el-form-item>
                        <el-form-item label="请求params:">
                            <el-button type="primary" @click="addTableRow('apiParams')">添加</el-button>
                            <base-variables-table :tableData="form.apiParams" @removeRowByIndex="removeParamsByIndex"></base-variables-table>
                        </el-form-item>
                        <el-form-item label="请求体:">
                            <el-button type="primary" @click="addTableRow('apiPostData')">添加</el-button>
                            <base-variables-table :tableData="form.apiPostData" @removeRowByIndex="removePostDataByIndex"></base-variables-table>
                        </el-form-item>
                        <el-form-item label="httpRunner-variables:">
                            <el-button type="primary" @click="addTableRow('variables')">添加</el-button>
                            <base-variables-table :tableData="form.variables" @removeRowByIndex="removeVariablesByIndex"></base-variables-table>
                        </el-form-item>
                        <el-form-item label="httpRunner-extract:">
                            <el-button type="primary" @click="addTableRow('extract')">添加</el-button>
                            <base-extract-table :tableData="form.extract" @removeRowByIndex="removeExtractByIndex"></base-extract-table>
                        </el-form-item>
                        <el-form-item label="httpRunner-validate:">
                            <el-button type="primary" @click="addTableRow('validate')">添加</el-button>
                            <base-validate-table :tableData="form.validate" @removeRowByIndex="removeValidateByIndex"></base-validate-table>
                        </el-form-item>
                        <el-form-item label="httpRunner-setup_hooks:" style="width:600px;">
                            <el-button type="primary" @click="addTableRow('setupHooks')">添加</el-button>
                            <base-hooks-table :tableData="form.setupHooks" @removeRowByIndex="removeSetupHooksByIndex"></base-hooks-table>
                        </el-form-item>
                        <el-form-item label="httpRunner-teardown_hooks:" style="width:600px;">
                            <el-button type="primary" @click="addTableRow('teardownHooks')">添加</el-button>
                            <base-hooks-table :tableData="form.teardownHooks" @removeRowByIndex="removeTeardownHooksByIndex"></base-hooks-table>
                        </el-form-item>
                        <el-form-item label="接口描述:" prop="apiDesc">
                            <el-input type="textarea" v-model="form.apiDesc"></el-input>
                        </el-form-item>
                        <el-form-item>
                            <el-button type="success" @click="save('form')">保存</el-button>
                            <el-button type="primary" @click="saveAndContinue('form')">保存并继续添加</el-button>
                            <el-button type="warning" @click="resetForm('form')">重置</el-button>
                            <el-button type="danger" @click="cancelSave()">取消</el-button>
                        </el-form-item>
                    </el-form>
                </el-col>
            </el-row>
        </el-main>
    </el-container>
</template>

<script>
    import baseVariablesTable from '@components/baseVariablesTable'
    import baseValidateTable from '@components/baseValidateTable'
    import baseExtractTable from '@components/baseExtractTable'
    import baseHooksTable from '@components/baseHooksTable'
    import {removeInArrayByIndex} from '@utils/common.js'
    import {getProjectsByPagination} from "@api/project"
    import {getEnvironmentsByPagination} from "@api/environment"
    import {addApi} from "@api/api";
    export default {
        name: 'apiAdd',
        components: {baseVariablesTable, baseValidateTable, baseExtractTable, baseHooksTable},
        data() {
            return {
                form: {
                    // 格式为: [{"keyName":"", "valueType":"", "value":""}, {, , }...]
                    apiHeaders: [],
                    // 格式为: [{"keyName":"", "valueType":"", "value":""}, {, , }...]
                    apiParams: [],
                    // 格式为: [{"keyName":"", "valueType":"", "value":""}, {, , }...]
                    apiPostData: [],
                    // 格式为: [{"keyName":"", "valueType":"", "value":""}, {, , }...]
                    variables: [],
                    // 格式为: [{"keyName":"", "value":""}, {, , }...]
                    extract: [],
                    // 格式为: [{"keyName":"", "comparator":"", "valueType":"", "value":""}], {, , }...
                    validate: [],
                    // 格式为: ["${setup_hook_prepare_kwargs($request)}", ...]
                    setupHooks: [],
                    // 格式为: ["${setup_hook_prepare_kwargs($request)}", ...]
                    teardownHooks: []
                },
                rules: {
                    apiName: [
                        { required: true, message: '请输入接口名称', trigger: 'blur' },
                    ],
                    apiPath: [
                        { required: true, message: '请输入apiPath', trigger: 'blur' },
                    ],
                    apiMethod: [
                        { required: true, message: '请输入请求方法', trigger: 'blur' },
                    ],
                    apiContentType: [
                        { required: true, message: '请输入ContentType', trigger: 'blur' },
                    ],
                    relativePro: [
                        { required: true, message: '请选择关联项目', trigger: 'blur' },
                    ]
                },
                apiBaseUrls: [],
                relativePros: [],
            }
    },
        methods: {
            resetForm(formName) {
                this.$refs[formName].resetFields();
            },
            successMessage(){
                this.$message({
                    message: '数据添加成功',
                    type: 'success'
                });
            },
            failMessage(){
                this.$message({
                    message: '数据添加失败',
                    type: 'error'
                });
            },
            save: function(formData){
                this.$refs[formData].validate((valid) => {
                    if (valid) {
                        //需要格式转换
                        let postForm = this.formTransform();
                        addApi(postForm).then(response => {
                            const code = response.data.code;
                            if(code == 200){
                                this.successMessage();
                                this.$router.push('/api-index');
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
                        addApi(postForm).then(response => {
                            const code = response.data.code;
                            if(code == 200){
                                this.successMessage();
                                this.$router.go(0);
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
                this.$router.push('/api-index')
            },
            addTableRow(name){
                if(name === "apiHeaders"){
                    this.form.apiHeaders.push({});
                }
                else if (name === "apiParams"){
                    this.form.apiParams.push({});
                }
                else if (name === "apiPostData"){
                    this.form.apiPostData.push({});
                }
                else if (name === "variables"){
                    this.form.variables.push({});
                }
                else if (name === "extract"){
                    this.form.extract.push({});
                }
                else if (name === "validate"){
                    this.form.validate.push({});
                }
                else if (name === "setupHooks"){
                    this.form.setupHooks.push({});
                }
                else if (name === "teardownHooks"){
                    this.form.teardownHooks.push({});
                }
                else {
                    window.console.log("无法识别: " + name);
                }
            },
            // 移除headers单行
            removeHeadersByIndex(array, removeIndex){
                this.form.apiHeaders = removeInArrayByIndex(array, removeIndex);
            },
            // 移除params单行
            removeParamsByIndex(array, removeIndex){
                this.form.apiParams = removeInArrayByIndex(array, removeIndex);
            },
            // 移除postData单行
            removePostDataByIndex(array, removeIndex){
                this.form.apiPostData = removeInArrayByIndex(array, removeIndex);
            },
            // 移除variables单行
            removeVariablesByIndex(array, removeIndex){
                this.form.variables = removeInArrayByIndex(array, removeIndex);
            },
            // 移除extract单行
            removeExtractByIndex(array, removeIndex){
                this.form.extract = removeInArrayByIndex(array, removeIndex);
            },
            // 移除validate单行
            removeValidateByIndex(array, removeIndex){
                this.form.validate = removeInArrayByIndex(array, removeIndex);
            },
            // 移除setupHooks单行
            removeSetupHooksByIndex(array, removeIndex){
                this.form.setupHooks = removeInArrayByIndex(array, removeIndex);
            },
            // 移除teardownHooks单行
            removeTeardownHooksByIndex(array, removeIndex){
                this.form.teardownHooks = removeInArrayByIndex(array, removeIndex);
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
        created(){
        getProjectsByPagination({pageindex: 1, pagesize: 1000}).then(response => {
            this.relativePros = response.data.data
        });
        getEnvironmentsByPagination({pageindex: 1, pagesize: 1000}).then(response => {
            this.apiBaseUrls = response.data.data
        });

    }
}
</script>

<style scoped>


</style>