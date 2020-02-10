<template>
    <el-container>
        <el-main>
            <el-row>
                <el-col :span="12">
                    <el-breadcrumb separator-class="el-icon-arrow-right">
                        <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
                        <el-breadcrumb-item>添加用例集</el-breadcrumb-item>
                    </el-breadcrumb>
                </el-col>
                <el-col :span="12"></el-col>
            </el-row>
            <br>
            <el-row>
                <el-col :span="24">
                    <el-form ref="form" :model="form" label-width="100px" label-position="top" size="mini">
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
                                    v-for="item in apiUrls"
                                    :key="item.value"
                                    :label="item.label"
                                    :value="item.value">
                                </el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="httpRunner-public-variables:" style="width:600px;">
                            <el-input type="textarea" v-model="form.configVariables" placeholder='请输入正确格式,e.g. [{"expected_status_code": 200}, {"foo":"bar"}]'></el-input>
                        </el-form-item>
                        <el-divider></el-divider>
                        <p>私有配置:</p>
                        <el-form-item label="用例集名称:" required style="width:600px;">
                            <el-input v-model="form.suiteName"></el-input>
                        </el-form-item>
                        <!-- <el-form-item label="httpRunner-variables:">
                            <el-input type="textarea" v-model="form.variables" placeholder='请输入正确格式,e.g. [{"expected_status_code": 200}, {"foo":"bar"}]'></el-input>
                        </el-form-item>
                        <span style="color:#ec1010; font-size:8px;">validate比较符包括: eq lt le gt ge len_eq len_gt len_lt contains contained_by startswith endswith</span>
                        <el-form-item label="httpRunner-validate:">
                            <el-input type="textarea" v-model="form.validate" placeholder='请输入正确格式, e.g. [["eq", "status_code", "$expected_status_code"], ["eq", "content.headers.Host", "httpbin.org"]]'></el-input>
                        </el-form-item> -->
                        <el-form-item label="拼装用例集:" required>
                            <el-transfer
                                filterable
                                filter-placeholder="请输入关键字搜索"
                                v-model="choicedSteps"
                                :data="allSteps"
                                :titles="['可选用例', '当前用例集合']">
                            </el-transfer>
                        </el-form-item>
                        <el-form-item label="指定步骤添加:variables / parameters">
                            <el-button type="primary" @click="addExtract">添加</el-button>
                            <div id="stepZone">
                                <el-row>
                                    <div name="stepExtract">
                                        <el-table
                                        name="stepExtractTable"
                                        width="100%"
                                        :data="stepExtractTableData">
                                            <el-table-column
                                                label="已选中用例"
                                                width="180">
                                                <template slot-scope="scope">
                                                    <el-select v-model="stepExtractTableData[scope.$index]['stepId']" placeholder="请选择">
                                                        <el-option
                                                            v-for="item in choicedSteps"
                                                            :key="item"
                                                            :label="allStepsDict[item]"
                                                            :value="item">
                                                        </el-option>
                                                    </el-select>
                                                </template>
                                            </el-table-column>
                                            <el-table-column
                                                label="variables内容"
                                                width="250">
                                                <template slot-scope="scope">
                                                    <el-input type="textarea" v-model="scope.row.variables"></el-input>
                                                </template>
                                            </el-table-column>
                                            <el-table-column
                                                label="parameters内容"
                                                width="250">
                                                <template slot-scope="scope">
                                                    <el-input type="textarea" v-model="scope.row.parameters"></el-input>
                                                </template>
                                            </el-table-column>
                                            <el-table-column
                                                label="操作"
                                                width="50">
                                                <template  slot-scope="scope">
                                                    <el-button type="danger" icon="el-icon-delete" circle size="mini" @click="removeByIndex(stepExtractTableData,scope.$index)"></el-button>
                                                </template>
                                            </el-table-column>
                                        </el-table>
                                    </div>
                                </el-row>
                            </div>
                        </el-form-item>
                        <el-form-item label="用例描述:" style="width:600px;">
                            <el-input type="textarea" v-model="form.caseDesc"></el-input>
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
import {removeInArrayByIndex} from '@utils/common.js'
export default {
    name: 'proAdd',
    data() {
        return {
            form: {
            
            },
            apiUrls: [],
            relativePros: [],
            // 可选api类型步骤
            apis: [],
            // 可选case类型步骤
            cases: [],
            // 当前所有可选用例步骤
            allSteps: [],
            // 当前选中用例步骤, TODO:提交表单记得转换格式["case-1", "api-3"] =>[{"1":"case"}, {"3":"api"}]
            choicedSteps: [],
            // 用于做已选中步骤的map映射 e.g. {"api-1":"apiName1", "case-2":"caseName2"}
            allStepsDict: {},
            // 格式为: {"apiId": [{"extractName":"extractValue"}, {"session_token":"content.token"}],"apiId2": [{"extractName":"extractValue"}, {"session_token":"content.token"}]}
            stepExtractTableData: []
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
        },
        // 根据this.apis / this.steps 生成transfer的可选步骤 data
        generateAllStep(){
            this.cases.forEach((item, index) => {
                let cur_item_key = "case-" + Object.keys(item)[0];
                let cur_item_value = "用例-" + item[Object.keys(item)[0]];
                this.allSteps.push({
                    label: cur_item_value,
                    key: cur_item_key
                });
                this.allStepsDict[cur_item_key]=cur_item_value;
            });
        },
        splitSeparator(str, separator){
            return str.split(separator)
        },
        addExtract(){
            this.stepExtractTableData.push({});
        },
        // 移除指定步骤的httpRunner参数
        removeByIndex(array, removeIndex){
            this.stepExtractTableData = removeInArrayByIndex(array, removeIndex);
        }

    },
    created(){
        window.console.log("请求后端获取所有项目");
        this.relativePros = [{label:"项目一", value: 11}, {label:"项目二", value: 22}];
        this.apiUrls = [{label: "https:127.0.0.1:443", value: "https:127.0.0.1:443"}, {label: "https:127.0.0.1:8080", value: "https:127.0.0.1:8080"}]
        window.console.log("请求后端获取所有接口步骤");
        this.apis = [{"1": "apiName1"}, {"3": "apiName2"}, {"5": "apiName3"}];
        this.cases = [{"1": "caseName1"}, {"2": "caseName2"}, {"4": "caseName3"}];
        this.generateAllStep();
        // this.choicedSteps = [];
    }
}
</script>

<style scoped>
   div.el-divider.el-divider--horizontal {
        width: 603px;
   }
</style>