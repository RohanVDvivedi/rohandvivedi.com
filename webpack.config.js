module.exports = {
    entry: './rjs/main.js',
    output: {
        path: __dirname + '/public/static/js',
        publicPath: '/',
        filename: 'bundle.js'
    },
    module: {
        rules: [
            {
                test: /\.(js|jsx)$/,
                exclude: /node_modules/,
                use: ['babel-loader']    
            }
        ]  
    }
};