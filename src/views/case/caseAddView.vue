<template>
    <el-container>
        <el-main>
            <el-row>
                <el-col :span="12">
                    <el-breadcrumb separator-class="el-icon-arrow-right">
                        <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
                        <el-breadcrumb-item>添加用例</el-breadcrumb-item>
                    </el-breadcrumb>
                </el-col>
                <el-col :span="12"></el-col>
            </el-row>
            <br>
            <el-row>
                <el-col :span="24">
                    <el-form ref="form" :model="form" :rules="rules" label-width="100px" label-position="top" size="mini">
                        <p>公共配置:</p>
                        <el-form-item  label="httpRunner-public-url:" prop="configBaseUrl" required>
                            <el-select
                                    v-model="form.configBaseUrl"
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
                        <el-form-item label="httpRunner-public-headers:" style="width:600px;">
                            <el-button type="primary" @click="addTableRow('configHeader')">添加</el-button>
                            <base-variables-table :tableData="form.configHeaders" @removeRowByIndex="removeHeadersByIndex"></base-variables-table>
                        </el-form-item>
                        <el-form-item label="httpRunner-public-variables:" style="width:600px;">
                            <el-button type="primary" @click="addTableRow('configVariables')">添加</el-button>
                            <base-variables-table :tableData="form.configVariables" @removeRowByIndex="removeVariablesByIndex"></base-variables-table>
                        </el-form-item>
                        <el-form-item label="httpRunner-public-output:" style="width:600px;">
                            <el-button type="primary" @click="addTableRow('configOutput')">添加</el-button>
                            <base-output-table :tableData="form.configOutput" @removeRowByIndex="removeOutputByIndex"></base-output-table>
                        </el-form-item>
                        <el-form-item label="httpRunner-public-setup_hooks:" style="width:600px;">
                            <el-button type="primary" @click="addTableRow('setupHooks')">添加</el-button>
                            <base-hooks-table :tableData="form.setupHooks" @removeRowByIndex="removeSetupHooksByIndex"></base-hooks-table>
                        </el-form-item>
                        <el-form-item label="httpRunner-public-teardown_hooks:" style="width:600px;">
                            <el-button type="primary" @click="addTableRow('teardownHooks')">添加</el-button>
                            <base-hooks-table :tableData="form.teardownHooks" @removeRowByIndex="removeTeardownHooksByIndex"></base-hooks-table>
                        </el-form-item>
                        <el-divider></el-divider>
                        <p>私有配置:</p>
                        <el-form-item label="用例名称:" prop="caseName" required style="width:600px;">
                            <el-input v-model="form.caseName"></el-input>
                        </el-form-item>
                        <el-form-item label="关联项目:" required>
                            <el-select v-model="form.relativePro" placeholder="请选择关联项目" multiple style="width: 300px;">
                                <el-option :label="item.proName" :value="item.id" v-for="item in relativePros" :key="item.value"></el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="拼装用例步骤:" required>
                            <el-select v-model="filterPros" multiple placeholder="请选择项目">
                                <el-option
                                        v-for="item in projectOptions"
                                        :key="item.id"
                                        :label="item.proName"
                                        :value="item.id">
                                </el-option>
                            </el-select>
                            <el-button icon="el-icon-search" circle @click="filter"></el-button>
                            <el-transfer
                                filterable
                                filter-placeholder="请输入关键字搜索"
                                v-model="form.caseSteps"
                                :data="allSteps"
                                :titles="['可选步骤', '当前用例步骤']">
                            </el-transfer>
                        </el-form-item>
                        <el-form-item label="指定接口步骤添加:extract / validate / variables">
                            <el-button type="primary" @click="addTableRow('stepExtract')">添加extract</el-button>
                            <el-row>
                                <el-col :span="15">
                                    <base-step-extract-table :tableData="form.stepExtracts" :choicedSteps="choicedApiSteps" @removeRowByIndex="removeStepExtractByIndex"></base-step-extract-table>
                                </el-col>
                            </el-row>
                            <el-button type="primary" @click="addTableRow('StepValidate')">添加validate</el-button>
                            <el-row>
                                <el-col :span="20">
                                    <base-step-validate-table :tableData="form.stepValidates" :choicedSteps="choicedApiSteps" @removeRowByIndex="removeStepValidateByIndex"></base-step-validate-table>
                                </el-col>
                            </el-row>
                            <el-button type="primary" @click="addTableRow('StepVariables')">添加variables</el-button>
                            <el-row>
                                <el-col :span="18">
                                    <base-step-variables-table :tableData="form.stepVariables" :choicedSteps="choicedApiSteps" @removeRowByIndex="removeStepVariablesByIndex"></base-step-variables-table>
                                </el-col>
                            </el-row>
                        </el-form-item>
                        <el-form-item label="用例描述:" style="width:600px;">
                            <el-input type="textarea" v-model="form.caseDesc"></el-input>
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
    import baseStepExtractTable from '@components/baseStepExtractTable'
    import baseStepValidateTable from '@components/baseStepValidateTable'
    import baseStepVariablesTable from '@components/baseStepVariablesTable'
    import baseHooksTable from '@components/baseHooksTable'
    import baseOutputTable from '@components/baseOutputTable'
    import {removeInArrayByIndex} from '@utils/common.js'
    import {getProjectsByPagination} from "@api/project"
    import {getEnvironmentsByPagination} from "@api/environment"
    import {getApisByPagination} from "@api/api"
    import {getCasesByPagination, addCase} from "@api/case"


    export default {
        name: 'caseAdd',
        components: {baseVariablesTable, baseHooksTable, baseOutputTable, baseStepExtractTable, baseStepValidateTable, baseStepVariablesTable},
        data() {
            return {
                form: {
                    configHeaders: [],
                    configVariables: [],
                    configOutput: [],
                    setupHooks: [],
                    teardownHooks: [],
                    stepExtracts: [],
                    stepValidates: [],
                    stepVariables: [],
                    // 当前选中用例步骤
                    caseSteps: [],
                },
                rules: {
                    configBaseUrl: [
                        { required: true, message: '请输入httpRunner-public-url', trigger: 'blur' },
                    ],
                    caseName: [
                        { required: true, message: '请输入用例名称', trigger: 'blur' },
                    ],
                },
                relativePros: [],
                // 可选api类型步骤
                apis: [],
                // 可选case类型步骤
                cases: [],
                apiBaseUrls: [],
                //当前可选步骤搜索的项目
                filterPros: [],
                //所有项目
                projectOptions: [

                ]
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
                        addCase(postForm).then(response => {
                            const code = response.data.code;
                            if(code == 200){
                                this.successMessage();
                                this.$router.push('/case-index');
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
                        addCase(postForm).then(response => {
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
                this.$router.push('/case-index')
            },
            addTableRow(name){
                if(name === "configHeaders"){
                    this.form.configHeaders.push({});
                }
                else if (name === "apiParams"){
                    this.form.apiParams.push({});
                }
                else if (name === "apiPostData"){
                    this.form.apiPostData.push({});
                }
                else if (name === "configOutput"){
                    this.form.configOutput.push({});
                }
                else if (name === "configHeader"){
                    this.form.configHeaders.push({});
                }
                else if (name === "configVariables"){
                    this.form.configVariables.push({});
                }
                else if (name === "extract"){
                    this.form.extract.push({});
                }
                else if (name === "stepExtract"){
                    this.form.stepExtract.push({});
                }
                else if (name === "validate"){
                    this.form.validate.push({});
                }
                else if (name === "StepValidate"){
                    this.form.stepValidates.push({});
                }
                else if (name === "StepVariables"){
                    this.form.stepVariables.push({});
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
                this.form.configHeaders = removeInArrayByIndex(array, removeIndex);
            },
            // 移除variables单行
            removeVariablesByIndex(array, removeIndex){
                this.form.configVariables = removeInArrayByIndex(array, removeIndex);
            },
            // 移除output单行
            removeOutputByIndex(array, removeIndex){
                this.form.configOutput = removeInArrayByIndex(array, removeIndex);
            },
            // 移除setupHooks单行
            removeSetupHooksByIndex(array, removeIndex){
                this.form.setupHooks = removeInArrayByIndex(array, removeIndex);
            },
            // 移除teardownHooks单行
            removeTeardownHooksByIndex(array, removeIndex){
                this.form.teardownHooks = removeInArrayByIndex(array, removeIndex);
            },
            // 移除stepExtract单行
            removeStepExtractByIndex(array, removeIndex){
                this.form.stepExtracts = removeInArrayByIndex(array, removeIndex);
            },
            // 移除stepValidate单行
            removeStepValidateByIndex(array, removeIndex){
                this.form.stepValidates = removeInArrayByIndex(array, removeIndex);
            },
            // 移除stepVariables单行
            removeStepVariablesByIndex(array, removeIndex){
                this.form.stepVariables = removeInArrayByIndex(array, removeIndex);
            },
            filter(){
                window.console.log("根据this.filterPros 搜索所有api和case");
                // this.allSteps = [...]
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
        computed: {
            //组装成[{"key": "", "label": ""}, ...]
            apiStepOptions: function () {
                let options = [];
                if(this.apis != null){
                    for (let api of this.apis){
                        let labelName = "接口-" + api.apiName;
                        options.push({key: api.id, label: labelName})
                    }
                }
                return options
            },
            //组装成[{"key": "", "label": ""}, ...]
            caseStepOptions: function () {
                let options = [];
                if(this.cases != null){
                    for (let case_ of this.cases){
                        let labelName = "用例-" + case_.caseName;
                        options.push({key: case_.id, label: labelName})
                    }
                }

                return options
            },
            allSteps: function () {
                let options = this.apiStepOptions.concat(this.caseStepOptions);
                return options
            },
            choicedApiSteps: function () {
                let options = [];
                if(this.form.caseSteps != []){
                    let idLabelMap = {};
                    for (let item of this.apiStepOptions){
                        idLabelMap[item.key] = item.label;
                    }
                    for (let id_ of this.form.caseSteps){
                        if(id_ in idLabelMap){
                            options.push({id: id_, apiName: idLabelMap[id_]})
                        }
                    }
                }
                return options
            }
        },
        created(){
            getEnvironmentsByPagination({pageindex: 1, pagesize: 1000}).then(response => {
                this.apiBaseUrls = response.data.data
            });
            getProjectsByPagination({pageindex: 1, pagesize: 1000}).then(response => {
                this.projectOptions = response.data.data;
                this.relativePros = response.data.data;
            });
            getApisByPagination({pageindex: 1, pagesize: 1000}).then(response => {
                if(response.data.data != null){
                    this.apis = response.data.data;
                }
            });
            getCasesByPagination({pageindex: 1, pagesize: 1000}).then(response => {
                if(response.data.data != null){
                    this.cases = response.data.data;
                }
            });

    }
}
</script>

<style scoped>

</style>