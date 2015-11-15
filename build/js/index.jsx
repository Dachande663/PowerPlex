'use strict';

import React from 'react';
import ReactDOM from 'react-dom';
import PlexLibrary from './Elements/PlexLibrary.jsx';

ReactDOM.render(<PlexLibrary data="/data.json" />, document.getElementById('library'));
