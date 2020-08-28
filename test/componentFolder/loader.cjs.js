'use strict';

Object.defineProperty(exports, '__esModule', { value: true });

const index = require('./index-d3d02e95.js');
const patch = require('./patch-a72a92bd.js');

const defineCustomElements = (win, options) => {
  if (typeof window === 'undefined') return Promise.resolve();
  return patch.patchEsm().then(() => {
  return index.bootstrapLazy([["my-component.cjs",[[1,"my-component",{"first":[1],"middle":[1],"last":[1]}]]]], options);
  });
};

exports.defineCustomElements = defineCustomElements;
