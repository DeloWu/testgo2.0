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
                    <el-form ref="form" :model="form" label-width="100px" label-position="left" size="medium">
                        <el-form-item label="环境名称:" required>
                            <el-input v-model="form.envName"></el-input>
                        </el-form-item>
                        <el-form-item label="环境ip:" required>
                            <el-input v-model="form.envIp"></el-input>
                        </el-form-item>
                        <el-form-item label="环境端口:" required>
                            <el-input v-model="form.envPort"></el-input>
                        </el-form-item>
                        <el-form-item label="关联项目:" required>
                            <el-select v-model="form.relativePro.value" placeholder="请选择关联项目">
                                <el-option :label="item.label" :value="item.value" v-for="item in relativePros" :key="item.value"></el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="项目描述:">
                            <el-input type="textarea" v-model="form.envDesc"></el-input>
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
    name: 'envUpdate',
    data() {
        return {
            form: {
            
            },
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
    mounted(){
        window.console.log("请求后端获取所有项目");
        this.relativePros = [{label:"项目一", value: 11}, {label:"项目二", value: 22}];
        window.console.log("created_hook:向后端请求数据");
        this.form = {
                envName: "环境一",
                envIp: "https:127.0.0.1",
                envPort: 443,
                relativePro: {label:"项目一", value: 11},
                envDesc: "环境一描述"
            };
    }
}
</script>

<style scoped>


</style>