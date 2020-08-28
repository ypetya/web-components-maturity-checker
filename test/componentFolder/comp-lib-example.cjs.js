'use strict';

const index = require('./index-d3d02e95.js');
const patch = require('./patch-a72a92bd.js');

patch.patchBrowser().then(options => {
  return index.bootstrapLazy([["my-component.cjs",[[1,"my-component",{"first":[1],"middle":[1],"last":[1]}]]]], options);
});
