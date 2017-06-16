module.exports = {
  entry:  __dirname + "/src/main.js",
  output: {
    path: __dirname + "/static/js",
    filename: "bundle.js"
  },

  module: {
    loaders: [
      { test: /\.css$/, loader: 'style-loader!css-loader' },
    ]
  }
}