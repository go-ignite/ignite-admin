const UglifyJSPlugin = require('uglifyjs-webpack-plugin');

var path = require('path'), 
    HtmlWebpackPlugin = require('html-webpack-plugin'),
    ExtractTextPlugin = require("extract-text-webpack-plugin"),
   webpack = require('webpack');


module.exports = {
    entry: {
        index: __dirname + "/src/index.js",
        status: __dirname + "/src/main.js",
    },
    output: {
					   path: path.resolve(__dirname, './static'),
        filename: "js/[name].bundle.js"
    },


    module: {
         rules: [
									   {
              test: /\.css$/,
									     use: ExtractTextPlugin.extract({
                fallback: "style-loader",
                use: "css-loader"
              })
            },
												{
            test: /\.(jpg|png)$/,
            use: [
                'url-loader?limit=1&name=./images/[name].[ext]'
              ]
            },
            {
                test: /\.vue$/,
                loader: 'vue-loader'
            },
            {
                test: /.js$/,
                loader: 'babel-loader',
                exclude: /node_modules/,
                query: {
                    presets: ['es2015']
                }
            }
        ]
    },

    plugins: [
        // new UglifyJSPlugin(),
        new HtmlWebpackPlugin({
            template: './templates/index.html',
            filename: 'index.html'
									}),
        new HtmlWebpackPlugin({
            template: './templates/code.html',
            filename: 'code.html'
									}),
        new HtmlWebpackPlugin({
            template: './templates/status.html',
            filename: 'status.html'
									}),
        new HtmlWebpackPlugin({
            template: './templates/users.html',
            filename: 'users.html'
									}),
        new ExtractTextPlugin("css/style.css"),
    ]
}
