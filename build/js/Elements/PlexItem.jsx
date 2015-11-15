'use strict';

import React from 'react';

export default React.createClass({
	displayName: 'PlexItem',
	render: function() {
		return (
			<div className="col-md-4">
				<div className="panel panel-primary">
					<div className="panel-body">
						<h2>{this.props.item.Title}</h2>
						<p className="text-muted"><strong>Type:</strong> {this.props.item.Type}</p>
						<p className="text-muted"><strong>Year:</strong> {this.props.item.Year}</p>
						<p className="text-muted"><strong>Rating:</strong> {this.props.item.Rating}/10</p>
					</div>
				</div>
			</div>
		);
	}
});
