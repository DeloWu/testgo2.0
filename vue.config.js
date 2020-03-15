const path = require('path');

function resolve(dir) {
    return path.join(__dirname, dir)
}
module.exports = {
    lintOnSave: true,
    chainWebpack: (config) => {
        config.resolve.alias
            .set('@', resolve('src'))
            .set('@assets', resolve('src/assets'))
            .set('@components', resolve('src/components'))
            .set('@api', resolve('src/api'))
            .set('@directive', resolve('src/directive'))
            .set('@filters', resolve('src/filters'))
            .set('@router', resolve('src/router'))
            .set('@store', resolve('src/store'))
            .set('@utils', resolve('src/utils'))
            .set('@views', resolve('src/views'))
            .set('@static', resolve('src/static'))
            .set('@config', resolve('config'))
    },
    transpileDependencies: [
        'vue-echarts',
        'resize-detector'
    ],
    // module: {
    //     rules: [{
    //         test: /\.scss$/,
    //         use: [{
    //             loader: "style-loader" // 将 JS 字符串生成为 style 节点
    //         }, {
    //             loader: "css-loader" // 将 CSS 转化成 CommonJS 模块
    //         }, {
    //             loader: "sass-loader" // 将 Sass 编译成 CSS
    //         }]
    //     }]
    // }
}