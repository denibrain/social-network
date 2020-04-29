const path = require("path")
const CopyWebpackPlugin = require('copy-webpack-plugin')
const MiniCssExtractPlugin = require('mini-css-extract-plugin')

module.exports = {
    mode: 'development',
   // context: path.resolve(__dirname, "ui"),
    entry: {
        "main": "./ui/Pages/Main/index.js",
        "signup": "./ui/Pages/SignUp",
        "signin": "./ui/Pages/SignIn",
    },
    output: {
        path: path.resolve(__dirname, "static/ui"),
        publicPath: 'ui/',
        filename: "[name].bundle.js",
        sourceMapFilename: '[name].[hash:8].map',
       // chunkFilename: '[id].[name].[hash:8].js',
    },
    module: {
        rules: [
            {
                test: /\.css$/,
                use:[MiniCssExtractPlugin.loader, 'css-loader']
            },
            {
                test: /\.scss$/,
                use:[MiniCssExtractPlugin.loader, 'css-loader', 'sass-loader']
            },
            {
                test: /\.png$/,
                use:['file-loader']
            },
            {
                test: /\.ttf$/,
                use:['file-loader']
            }
        ]
    },
    plugins: [
        new CopyWebpackPlugin(
            [
                {
                    from: path.resolve(__dirname, "ui/Images"),
                    to: path.resolve(__dirname, "static/ui")
                }
            ]
        ),
        new MiniCssExtractPlugin({
            filename: "[name].bundle.css",
        })
    ],
    resolve: {
        modules: [
          //  path.resolve(__dirname, "ui"),
            'node_modules',
        ],
        extensions: ['.js', '.jsx', '.html', '.json'],
        // alias: {
        //     jquery: "jquery/dist/jquery",
        // },
        //
        // mainFields: ["main", "module"],
    },
};