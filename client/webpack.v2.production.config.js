var webpack = require('webpack');var ExtractTextPlugin = require('extract-text-webpack-plugin');
var path = require('path');
var loaders = require('./webpack.loaders');

module.exports = {
  entry: [
    // Your appʼs entry point
    './v2/index.jsx'
  ],
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
    new ExtractTextPlugin('style.css', { allChunks: true })
  ],
  postcss: [
    require('autoprefixer')
  ]
};
