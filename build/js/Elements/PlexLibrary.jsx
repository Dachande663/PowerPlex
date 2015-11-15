'use strict';

import PlexItems from './PlexItems.jsx';
import PourOver from 'pourover';
import React from 'react';

export default React.createClass({
	displayName: 'PlexLibrary',
	loadItems: function() {
		var self = this;
		jQuery.getJSON("/assets/test/data.json", function(resp){


			var items = [];
			$.each(resp, function(el, item){
				items.push(item);
			});
			var collection = new PourOver.Collection(items);



			var type_filter   = PourOver.makeExactFilter('Type', ['movie', 'show', 'season', 'episode']);
			collection.addFilters([type_filter]);


			var NameSort = PourOver.Sort.extend({
				fn: function(a, b) {
					if(a.Title < b.Title) return -1;
					if(a.Title > b.Title) return 1;
					return 0;
				}
			});

			var RatingSort = PourOver.Sort.extend({
				fn: function(a, b) {
					if(a.Rating < b.Rating) return 1;
					if(a.Rating > b.Rating) return -1;
					return 0;
				}
			});

			var YearSort = PourOver.Sort.extend({
				fn: function(a, b) {
					if(a.Year < b.Year) return -1;
					if(a.Year > b.Year) return 1;
					return 0;
				}
			});

			collection.addSorts([
				new NameSort('Title'),
				new RatingSort('Rating'),
				new YearSort('Year')
			]);




			var view = new PourOver.View('default_view', collection, { page_size: 500 });
			view.collection.filters.Type.query('show');
			view.setSort('Title');

			self.setState({ items: view });

		});
	},
	getInitialState: function() {
		return { items: null };
	},
	componentDidMount: function() {
		this.loadItems();
	},
	onClickSort: function(sort, e) {
		e.preventDefault();
		var view = this.state.items;
		view.setSort(sort);
		this.setState({ items: view });
	},
	onClickType: function(type, e) {
		e.preventDefault();
		var view = this.state.items;
		view.collection.filters.Type.clearQuery()
		view.collection.filters.Type.query(type);
		this.setState({ items: view });
	},
	render: function() {
		return (
			<div className="row">

				<div className="col-md-3">
					<div className="panel panel-default">
						<div className="panel-heading">Sort</div>
						<div className="panel-body">
							<ul className="list-unstyled">
								<li><a href="#" onClick={this.onClickSort.bind(this, 'Title')}>Name</a></li>
								<li><a href="#" onClick={this.onClickSort.bind(this, 'Rating')}>Rating</a></li>
								<li><a href="#" onClick={this.onClickSort.bind(this, 'Year')}>Year</a></li>
							</ul>
						</div>
					</div>
					<div className="panel panel-default">
						<div className="panel-heading">Type</div>
						<div className="panel-body">
							<ul className="list-unstyled">
								<li><a href="#" onClick={this.onClickType.bind(this, 'movie')}>Movie</a></li>
								<li><a href="#" onClick={this.onClickType.bind(this, 'show')}>Show</a></li>
								<li><a href="#" onClick={this.onClickType.bind(this, 'season')}>Season</a></li>
								<li><a href="#" onClick={this.onClickType.bind(this, 'episode')}>Episode</a></li>
							</ul>
						</div>
					</div>
				</div>

				<div className="col-md-9">
					<PlexItems items={this.state.items} />
				</div>

			</div>
		);
	}
});
