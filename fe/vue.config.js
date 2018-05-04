module.exports = {
  lintOnSave: false,
  devServer: {
    proxy: 'http://localhost:8000'
  },
  outputDir: '../static',
  configureWebpack: config => {
    if (process.env.NODE_ENV === 'production') {
      config.output.publicPath = '/static/'
    }
  }
}
