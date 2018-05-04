module.exports = {
  lintOnSave: false,
  devServer: {
    proxy: 'http://localhost:8000',
  },
  outputDir: '../static',
}
