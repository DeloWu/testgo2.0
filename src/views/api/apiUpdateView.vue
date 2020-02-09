<template>
    <el-container>
        <el-main>
            <el-row>
                <el-col :span="12">
                    <el-breadcrumb separator-class="el-icon-arrow-right">
                        <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
                        <el-breadcrumb-item>编辑接口</el-breadcrumb-item>
                    </el-breadcrumb>
                </el-col>
                <el-col :span="12"></el-col>
            </el-row>
            <br>
            <el-row>
                <el-col :span="12">
                    <el-form ref="form" :model="form" label-width="100px" label-position="top" size="mini">
                        <el-form-item label="接口名称:" required>
                            <el-input v-model="form.apiName"></el-input>
                        </el-form-item>
                        <el-form-item  label="接口url:" required>
                            <el-select
                                v-model="form.apiUrl"
                                filterable
                                allow-create
                                default-first-option
                                placeholder="请选择或者新增url"
                                style="width:592px;">
                                <el-option
                                    v-for="item in apiUrls"
                                    :key="item.value"
                                    :label="item.label"
                                    :value="item.value">
                                </el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item  label="请求方法:" required>
                            <el-radio v-model="form.apiMethod" label="get">get</el-radio>
                            <el-radio v-model="form.apiMethod" label="post">post</el-radio>
                            <el-radio v-model="form.apiMethod" label="put">put</el-radio>
                            <el-radio v-model="form.apiMethod" label="patch">patch</el-radio>
                            <el-radio v-model="form.apiMethod" label="delete">delete</el-radio>
                            <el-radio v-model="form.apiMethod" label="head">head</el-radio>
                            <el-radio v-model="form.apiMethod" label="options">options</el-radio>
                        </el-form-item>
                        <el-form-item  label="Content-Type:" required style="width:800px;">
                            <el-radio v-model="form.apiContentType" label="application/json">application/json</el-radio>
                            <el-radio v-model="form.apiContentType" label="application/x-www-form-urlencoded">application/x-www-form-urlencoded</el-radio>
                            <el-radio v-model="form.apiContentType" label="text/html; charset=UTF-8">text/html; charset=UTF-8(get方法选用)</el-radio>
                            <el-radio v-model="form.apiContentType" label="text/plain">text/plain</el-radio>
                            <el-radio v-model="form.apiContentType" label="application/xml">application/xml</el-radio>
                            <el-radio v-model="form.apiContentType" label="multipart/form-data">multipart/form-data</el-radio>
                        </el-form-item>
                        <el-form-item label="接口请求头:">
                            <el-input type="textarea" v-model="form.apiHeader" placeholder="请输入正确json格式"></el-input>
                        </el-form-item>
                        <el-form-item label="请求params:">
                            <el-input type="textarea" v-model="form.apiParams" placeholder="请输入正确json格式"></el-input>
                        </el-form-item>
                        <el-form-item label="请求体:">
                            <el-input type="textarea" v-model="form.apiPostData" placeholder="请输入正确格式"></el-input>
                        </el-form-item>
                        <el-form-item label="关联项目:" required>
                            <el-select v-model="form.relativePro" placeholder="请选择关联项目" style="width: 300px;">
                                <el-option :label="item.label" :value="item.value" v-for="item in relativePros" :key="item.value"></el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="httpRunner-variables:">
                            <el-input type="textarea" v-model="form.variables" placeholder='请输入正确格式,e.g. [{"expected_status_code": 200}, {"foo":"bar"}]'></el-input>
                        </el-form-item>
                        <span style="color:#ec1010; font-size:8px;">validate比较符包括: eq lt le gt ge len_eq len_gt len_lt contains contained_by startswith endswith</span>
                        <el-form-item label="httpRunner-validate:">
                            <el-input type="textarea" v-model="form.validate" placeholder='请输入正确格式, e.g. [["eq", "status_code", "$expected_status_code"], ["eq", "content.headers.Host", "httpbin.org"]]'></el-input>
                        </el-form-item>
                        <el-form-item label="接口描述:">
                            <el-input type="textarea" v-model="form.apiDesc"></el-input>
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
            relativePros: []
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
    created(){
        this.form = {
                    apiIndex: 11,
                    apiName: "接口一",
                    apiUrl: "http://127.0.0.1:8080/",
                    apiMethod: "get",
                    apiContentType: "application/json",
                    apiHeader: JSON.stringify({}),
                    apiParams: JSON.stringify({"foo": "bar"}),
                    apiPostData: JSON.stringify({"foo1": "bar1"}),
                    apiDesc: "接口描述一",
                    relativePro: 11,
                    createTime: 1581135254,
                    modifyTime: 1581135255,
                    variables: JSON.stringify([{"expected_status_code": 200}]),
                    validate: JSON.stringify([["eq", "status_code", "$expected_status_code"], ["eq", "content.headers.Host", "httpbin.org"]])
                };
        window.console.log("请求后端获取所有项目");
        this.relativePros = [{label:"项目一", value: 11}, {label:"项目二", value: 22}];
        this.apiUrls = [{label: "https:127.0.0.1:443", value: "https:127.0.0.1:443"}, {label: "https:127.0.0.1:8080", value: "https:127.0.0.1:8080"}]
    }
}
</script>

<style scoped>
    

</style>