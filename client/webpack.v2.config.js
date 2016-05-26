var webpack = require('webpack');
var loaders = require('./webpack.loaders');
var ExtractTextPlugin = require('extract-text-webpack-plugin');
var path = require('path');

module.exports = {
  entry: [
    // Your appʼs entry point
    './v2/index.jsx'
  ],
  devtool: process.env.WEBPACK_DEVTOOL || 'source-map',
  output: {
    path: path.join(__dirname, '../public-v2'),
    filename: 'bundle.js'
  },
  resolve: {
    extensions: ['', '.js', '.jsx']
  },
  module: {
    loaders: loaders
  },
  plugins: [
    new ExtractTextPlugin('style.css', { allChunks: true }),
    new webpack.NoErrorsPlugin()
  ],
  postcss: [
    require('autoprefixer')
  ]
};
