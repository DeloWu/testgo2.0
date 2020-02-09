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
                <el-col :span="12">
                    <el-form ref="form" :model="form" label-width="100px" label-position="top" size="mini">
                        <p>公共配置:</p>
                        <el-form-item  label="httpRunner-public-url:" required>
                            <el-select
                                v-model="form.configBaseUrl"
                                filterable
                                allow-create
                                default-first-option
                                placeholder="请选择或者新增url, e.g. http://127.0.0.1:8080"
                                style="width:592px;">
                                <el-option
                                    v-for="item in apiUrls"
                                    :key="item.value"
                                    :label="item.label"
                                    :value="item.value">
                                </el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="httpRunner-public-variables:">
                            <el-input type="textarea" v-model="form.variables" placeholder='请输入正确格式,e.g. [{"expected_status_code": 200}, {"foo":"bar"}]'></el-input>
                        </el-form-item>
                        <el-form-item label="httpRunner-public-output:">
                            <el-input type="textarea" v-model="form.output" placeholder='请输入正确格式,e.g. ["session_token"]'></el-input>
                        </el-form-item>
                        <el-form-item label="httpRunner-public-setup_hooks:">
                            <el-input type="textarea" v-model="form.setupHooks" placeholder='请输入正确格式, e.g. ["${setup_hook_prepare_kwargs($request)}","${teardown_hook_sleep_N_secs($response, n_secs)}"]'></el-input>
                        </el-form-item>
                        <el-form-item label="httpRunner-public-teardown_hooks:">
                            <el-input type="textarea" v-model="form.teardownHooks" placeholder='请输入正确格式, e.g. ["${setup_hook_prepare_kwargs($request)}","${teardown_hook_sleep_N_secs($response, n_secs)}"]'></el-input>
                        </el-form-item>
                        <el-divider></el-divider>
                        <p>私有配置:</p>
                        <el-form-item label="用例名称:" required>
                            <el-input v-model="form.caseName"></el-input>
                        </el-form-item>
                        <!-- <el-form-item label="httpRunner-variables:">
                            <el-input type="textarea" v-model="form.variables" placeholder='请输入正确格式,e.g. [{"expected_status_code": 200}, {"foo":"bar"}]'></el-input>
                        </el-form-item>
                        <span style="color:#ec1010; font-size:8px;">validate比较符包括: eq lt le gt ge len_eq len_gt len_lt contains contained_by startswith endswith</span>
                        <el-form-item label="httpRunner-validate:">
                            <el-input type="textarea" v-model="form.validate" placeholder='请输入正确格式, e.g. [["eq", "status_code", "$expected_status_code"], ["eq", "content.headers.Host", "httpbin.org"]]'></el-input>
                        </el-form-item> -->
                        <el-form-item label="拼装用例步骤:" required>
                            <el-transfer
                                filterable
                                filter-placeholder="请输入关键字搜索"
                                v-model="choiceSteps"
                                :data="allSteps"
                                :titles="['可选步骤', '当前用例步骤']">
                            </el-transfer>
                        </el-form-item>
                        <el-form-item label="指定步骤添加:extract / validate / variables">
                        </el-form-item>
                        <el-form-item label="用例描述:">
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
            choiceSteps: []
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
            this.apis.forEach((item, index) => {
                let cur_item_key = Object.keys(item)[0];
                let cur_item_value = item[Object.keys(item)[0]];
                this.allSteps.push({
                    label: "接口-" + cur_item_value,
                    key: "api"+cur_item_key
                });
            });
            this.cases.forEach((item, index) => {
                let cur_item_key = Object.keys(item)[0];
                let cur_item_value = item[Object.keys(item)[0]];
                this.allSteps.push({
                    label: "用例-" + cur_item_value,
                    key: "case-"+cur_item_key
                });
            });
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
        
    }
}
</script>

<style scoped>


</style>