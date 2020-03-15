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
}