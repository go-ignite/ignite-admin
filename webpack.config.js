const UglifyJSPlugin = require('uglifyjs-webpack-plugin');

var path = require('path'),
    HtmlWebpackPlugin = require('html-webpack-plugin'),
    ExtractTextPlugin = require("extract-text-webpack-plugin"),
    webpack = require('webpack');




module.exports = {

    entry: {
        index: "./src/index.js",
        status: "./src/main.js",
        code: "./src/code.js",
        about: "./src/about.js",
    },
    output: {
        path: path.resolve(__dirname, './static'),
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
                query: {
                    presets: ['es2015']
                }
            },
            {
                test: /\.(jpg|png)$/,
                use: [
                    'url-loader?limit=10000&name=images/[name].[ext]'
                ]
            },
        ]
    },

    plugins: [
        new webpack.LoaderOptionsPlugin({
            minimize: true
        }),
        new HtmlWebpackPlugin({
            template: './src/html/index.html',
            filename: 'index.html',
            chunks: ['index', 'commons'],
        }),
        new HtmlWebpackPlugin({
            template: './src/html/code.html',
            filename: 'code.html',
            chunks: ['code', 'commons'],
        }),
        new HtmlWebpackPlugin({
            template: './src/html/status.html',
            filename: 'status.html',
            chunks: ['status', 'commons'],
        }),
        new HtmlWebpackPlugin({
            template: './src/html/about.html',
            filename: 'about.html',
            chunks: ['about', 'commons'],
        }),
        new ExtractTextPlugin("css/style.css"),
        new webpack.optimize.CommonsChunkPlugin({
            name: 'commons',
            filename: 'js/commons.js',
            minChunks: 2,
        }),
    ]
}
