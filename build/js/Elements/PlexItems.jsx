'use strict';

import PlexItem from './PlexItem.jsx';
import React from 'react';
import Masonry from 'react-masonry-component';

var MasonryContainer = Masonry(React);

export default React.createClass({
	displayName: 'PlexItems',
	render: function() {
		var nodes = this.props.items ? this.props.items.getCurrentItems().map(function(item){
			return (<PlexItem key={item.Key} item={item}/>);
		}) : [];
		return (
			<MasonryContainer>
				{nodes}
			</MasonryContainer>
		);
	}
});
