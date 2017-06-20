const UglifyJSPlugin = require('uglifyjs-webpack-plugin');

module.exports = {
    entry: {
        index: __dirname + "/src/index.js",
        status: __dirname + "/src/main.js",
    },
    output: {
        path: __dirname + "/static/js",
        filename: "[name].bundle.js"
    },

    module: {
        loaders: [
            { test: /\.css$/, loader: 'style-loader!css-loader' },
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
    ]
}