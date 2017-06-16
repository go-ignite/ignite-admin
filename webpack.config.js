const UglifyJSPlugin = require('uglifyjs-webpack-plugin');

module.exports = {
    entry: __dirname + "/src/main.js",
    output: {
        path: __dirname + "/static/js",
        filename: "bundle.js"
    },

    module: {
        loaders: [
            { test: /\.css$/, loader: 'style-loader!css-loader' },
            {
                test: /\.vue$/,
                loader: 'vue-loader'
            }
        ]
    },

    plugins: [
        // new UglifyJSPlugin(),
    ]
}