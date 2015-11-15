module.exports = {

	cache: true,
	entry: './js/index.jsx',
	output: {
		filename: 'explorer.js',
		path: __dirname + '/../src/public/explorer'
	},

	module: {
		loaders: [
			{
				test: /\.jsx$/,
				exclude: /node_modules/,
				loader: 'babel-loader?cacheDirectory&presets[]=es2015&presets[]=react'
				// query: {
				// 	presets:['react']
				// }
				// 'loaders': ['jsx?harmony']
			}
		],
		postLoaders: [
			{ loader: 'transform/cacheable?brfs' }
		]
	},

	// externals: {
	// 	'react-dom': 'ReactDOM'
	// },

	resolve: {
		extensions: ['', '.js', '.jsx']
	}

};
