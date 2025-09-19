const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    host: "localhost", // 允许外部访问，如果只想本地访问可以设置为 'localhost'
    port: 8022, // 端口号，可以修改为你想要的端口
    open: true, // 启动后自动打开浏览器
    // 如果需要HTTPS，可以添加以下配置
    // https: true,
  },
});
