const UglifyJSPlugin = require('uglifyjs-webpack-plugin');

var path = require('path');
var HtmlWebpackPlugin = require('html-webpack-plugin');
var ExtractTextPlugin = require("extract-text-webpack-plugin");
var webpack = require('webpack');

function resolve (dir) {
    return path.join(__dirname, dir)
}

module.exports = {
    entry: {
        index: "./src/main.js",
    },
    output: {
        path: path.resolve(__dirname, '../static'),
        filename: "js/[name].bundle.js",
        publicPath: '/static/'
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
                test: /\.vue$/,
                loader: 'vue-loader'
            },
            {
                test: /\.js$/,
                loader: 'babel-loader',
                exclude: /node_modules/,
            },
            {
                test: /\.(jpg|png)$/,
                use: [
                    'url-loader?limit=10000&name=images/[name].[ext]'
                ]
            },
            // package fonts
            {
                test: /\.(eot|svg|ttf|woff|woff2)\w*/,
                loader: 'file-loader?outputPath=fonts/&name=[name].[ext]'
            },
        ]
    },
    resolve: {
        extensions: ['.js', '.vue', '.json'],
        alias: {
          '@': resolve('src'),
          'vue$': 'vue/dist/vue.js',
        }
    },
    plugins: [
        new webpack.LoaderOptionsPlugin({
            minimize: true
        }),
        new HtmlWebpackPlugin({
            template: './src/index.html',
            filename: 'index.html',
            chunks: ['index', 'commons'],
        }),
        new ExtractTextPlugin("css/style.css"),
        new webpack.optimize.CommonsChunkPlugin({
            name: 'commons',
            filename: 'js/commons.js',
            minChunks: 2,
        }),
    ]
}
