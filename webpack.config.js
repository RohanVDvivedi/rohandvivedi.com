
const MiniCssExtractPlugin = require('mini-css-extract-plugin');

module.exports = {
    entry: './rjs/main.js',
    output: {
        path: __dirname + '/public/static/',
        filename: 'js/bundle.js'
    },
    module: {
        rules: [
            {
                test: /\.(js|jsx)$/,
                exclude: /node_modules/,
                use: ['babel-loader']    
            },
            {
                test: /\.css$/,
                use: [MiniCssExtractPlugin.loader, "css-loader"]
            }
        ]  
    },
    plugins: [
        new MiniCssExtractPlugin({
            filename: 'css/bundle.css'
        }),
    ]
};