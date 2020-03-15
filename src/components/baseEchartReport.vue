
<template>
    <!--为echarts准备一个具备大小的容器dom-->
    <div id="main" style="width: 600px;height: 400px;"></div>
</template>
<script>
    import echarts from 'echarts'
    export default {
        name: 'base-echart-report',
        data () {
            return {
                option: {
                    tooltip: {
                        trigger: 'axis',
                        axisPointer: {            // 坐标轴指示器，坐标轴触发有效
                            type: 'shadow'        // 默认为直线，可选为：'line' | 'shadow'
                        }
                    },
                    legend: {
                        data: ['成功数量', '失败数量']
                    },
                    grid: {
                        left: '3%',
                        right: '4%',
                        bottom: '3%',
                        containLabel: true
                    },
                    xAxis: {
                        type: 'category',
                        data: ['周一', '周二', '周三', '周四', '周五', '周六', '周日']
                    },
                    yAxis: {
                        type: 'value'
                    },
                    series: [
                        {
                            name: '失败数量',
                            type: 'bar',
                            stack: '总量',
                            label: {
                                show: true,
                                position: 'insideRight'
                            },
                            data: [20, 25, 10, 40, 90, 30, 20],
                            color: ['#dd6b66'],
                        },
                        {
                            name: '成功数量',
                            type: 'bar',
                            stack: '总量',
                            label: {
                                show: true,
                                position: 'insideRight'
                            },
                            data: [220, 182, 191, 234, 290, 330, 310],
                            color: ['#50a084'],
                        },
                    ]
                }
            }
        },
        methods:{
            draw(id){
                this.charts = echarts.init(document.getElementById(id), 'light');
                this.charts.setOption(this.option)
            }
        },
        //调用
        mounted(){
            this.$nextTick(function() {
                this.draw('main')
            })
        }
    }
</script>
<style scoped>
    * {
        margin: 0;
        padding: 0;
        list-style: none;
    }
</style>