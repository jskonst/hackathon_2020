const path = require("path");
const HtmlWebpackPlugin = require("html-webpack-plugin");

const  SERVER_HOST = 'localhost';
const  SERVER_PORT = 3000;
const SERVER_WEB_HOST = 'localhost';
const SERVER_WEB_PORT = '4000';


const proxy = `http://${SERVER_HOST}:${SERVER_PORT}`;

module.exports = {
  // mode: NODE_ENV || 'production',
  mode: 'production',
  entry: "./src/index.tsx",
  devtool: 'inline-source-map',
  output: {
    publicPath: "/",
    filename: "bundle.js",
    path: path.resolve(__dirname, "dist")
  },

  resolve: {
    extensions: [".ts", ".tsx", ".js", ".json"]
  },

  module: {
    rules: [
      { test: /\.tsx?$/, loader: "awesome-typescript-loader", exclude: /node_modules/ },
      {
        test: /\.(css|s[ac]ss)$/i,
        use: [
          { loader: 'style-loader' },
          {
            loader: "css-loader", options: {
              importLoaders: 1,
              // modules: true
            }
          },
          "sass-loader"
        ]
      },
      {
        test: /\.(png|jpe?g|gif)$/i,
        use: [
          {
            loader: 'file-loader',
          },
        ],
      },
    ]
  },

  plugins: [
    new HtmlWebpackPlugin({
      template: "./public/index.html",
      favicon: "./public/favicon.ico",
      filename: "index.html",
      inject: "body"
    })
  ],

  devServer: {
    allowedHosts: [
      SERVER_HOST,
    ],
    historyApiFallback: true,
    stats: 'errors-only',
    open: true,
    hot: true,
    host: SERVER_WEB_HOST,
    port: SERVER_WEB_PORT,
    proxy: {
      '/positions': proxy,
      '/api': proxy,
      '/acc': proxy,
      '/auth': proxy,
      '/socket.io': {
        target: proxy,
        ws: true
      },
      '/images': proxy
    }
  }
};
