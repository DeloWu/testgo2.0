<template>
    <el-container>
        <el-main>
            <el-row>
                <el-col :span="12">
                    <el-breadcrumb separator-class="el-icon-arrow-right">
                        <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
                        <el-breadcrumb-item>编辑用例集</el-breadcrumb-item>
                    </el-breadcrumb>
                </el-col>
                <el-col :span="12"></el-col>
            </el-row>
            <br>
            <el-row>
                <el-col :span="24">
                    <el-form ref="form" :model="form" :rules="rules" label-width="100px" label-position="top" size="mini">
                        <p>公共配置:</p>
                        <el-form-item  label="httpRunner-public-url:">
                            <el-select
                                    v-model="form.configBaseUrl"
                                    filterable
                                    allow-create
                                    default-first-option
                                    placeholder="请选择或者新增url, e.g. http://127.0.0.1:8080"
                                    style="width:600px;">
                                <el-option
                                        v-for="item in apiBaseUrls"
                                        :key="item.envIp + ':' + item.envPort"
                                        :label="item.envIp + ':' + item.envPort"
                                        :value="item.envIp + ':' + item.envPort">
                                </el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="关联项目:" required>
                            <el-select v-model="form.relativePro" placeholder="请选择关联项目" multiple style="width: 300px;">
                                <el-option :label="item.proName" :value="item.id" v-for="item in relativePros" :key="item.value"></el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="httpRunner-public-variables:" style="width:600px;">
                            <el-button type="primary" @click="addTableRow('configVariables')">添加</el-button>
                            <base-variables-table :tableData="form.configVariables" @removeRowByIndex="removeVariablesByIndex"></base-variables-table>
                        </el-form-item>
                        <el-divider></el-divider>
                        <p>私有配置:</p>
                        <el-form-item label="用例集名称:" prop="suiteName" required style="width:600px;">
                            <el-input v-model="form.suiteName"></el-input>
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
                                    v-model="form.suiteSteps"
                                    :data="suiteStepOptions"
                                    :titles="['可选步骤', '当前用例步骤']">
                            </el-transfer>
                        </el-form-item>
                        <el-form-item label="指定步骤添加:variables / parameters">
                            <el-button type="primary" @click="addStepVariables">添加variables</el-button>
                            <el-row>
                                <el-col :span="18">
                                    <base-step-variables-table :tableData="form.stepVariables" :choicedSteps="choicedCaseSteps" @removeRowByIndex="removeStepVariablesByIndex"></base-step-variables-table>
                                </el-col>
                            </el-row>
                            <el-button type="primary" @click="addStepParameters">添加parameters</el-button>
                            <el-row>
                                <el-col :span="18">
                                    <base-step-variables-table :tableData="form.stepParameters" :choicedSteps="choicedCaseSteps" @removeRowByIndex="removeStepParametersByIndex"></base-step-variables-table>
                                </el-col>
                            </el-row>
                        </el-form-item>
                        <el-form-item label="用例描述:" style="width:600px;">
                            <el-input type="textarea" v-model="form.suiteDesc"></el-input>
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
    import baseVariablesTable from '@components/baseVariablesTable'
    import baseStepVariablesTable from '@components/baseStepVariablesTable'
    import {removeInArrayByIndex} from '@utils/common.js'
    import {getProjectsByPagination} from "@api/project"
    import {getEnvironmentsByPagination} from "@api/environment"
    import {getCasesByPagination} from "@api/case"
    import {getSuiteById, editSuite} from "@api/suite"
    export default {
    name: 'suiteUpdate',
        components: {baseVariablesTable, baseStepVariablesTable},
        data() {
        return {
            form: {
                configVariables: [],
                stepVariables: [],
                stepParameters: [],
                suiteSteps: [],
            },
            rules: {
                suiteName: [
                    { required: true, message: '请输入用例集名称', trigger: 'blur' },
                ],
            },
            apiUrls: [],
            relativePros: [],
            // 可选api类型步骤
            apis: [],
            // 可选case类型步骤
            cases: [],
            // 当前所有可选用例步骤
            apiBaseUrls: [],
            //当前可选步骤搜索的项目
            filterPros: [],
            //所有项目
            projectOptions: [

            ],
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
        save(formData) {
            this.$refs[formData].validate((valid) => {
                if (valid) {
                    //需要格式转换
                    let postForm = this.formTransform();
                    editSuite(postForm).then(response => {
                        const code = response.data.code;
                        if(code == 200){
                            this.successMessage();
                            this.$router.push('/suite-index');
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
                    editSuite(postForm).then(response => {
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
            this.$router.push('/suite-index')
        },
        addTableRow(name){
            if(name === "configHeaders"){
                this.form.configHeaders.push({});
            }
            else if (name === "apiParams"){
                this.form.apiParams.push({});
            }
            else if (name === "configVariables"){
                this.form.configVariables.push({});
            }
            else if (name === "StepVariables"){
                this.form.stepVariables.push({});
            }
            else {
                window.console.log("无法识别: " + name);
            }
        },
        addStepVariables(){
            this.form.stepVariables.push({});
        },
        addStepParameters(){
            this.form.stepParameters.push({});
        },
        // 移除variables单行
        removeVariablesByIndex(array, removeIndex){
            this.form.configVariables = removeInArrayByIndex(array, removeIndex);
        },
        // 移除stepVariables单行
        removeStepVariablesByIndex(array, removeIndex){
            this.form.stepVariables = removeInArrayByIndex(array, removeIndex);
        },
        // 移除stepParameters单行
        removeStepParametersByIndex(array, removeIndex){
            this.form.stepParameters = removeInArrayByIndex(array, removeIndex);
        },
        filter(){
            window.console.log("根据this.filterPros 搜索所有api和case");
            // this.suiteStepOptions = [...]
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
        suiteStepOptions: function () {
            let options = [];
            if(this.cases != null){
                for (let case_ of this.cases){
                    let labelName = "用例-" + case_.caseName;
                    options.push({key: case_.id, label: labelName})
                }
            }

            return options
        },
        choicedCaseSteps: function () {
            let options = [];
            if(this.form.suiteSteps != []){
                let idLabelMap = {};
                for (let item of this.suiteStepOptions){
                    idLabelMap[item.key] = item.label;
                }
                for (let id_ of this.form.suiteSteps){
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
        getCasesByPagination({pageindex: 1, pagesize: 1000}).then(response => {
            if(response.data.data != null){
                this.cases = response.data.data;
            }
        });
        const updateId = this.$router.currentRoute.query.id;
        getSuiteById(updateId).then(response => {
            this.form = response.data.data;
        });
    }
}
</script>

<style scoped>
   div.el-divider.el-divider--horizontal {
        width: 603px;
   }
</style>