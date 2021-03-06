// snaplist view
var _ = require('lodash');
var $ = require('jquery');
var Backbone = require('backbone');
Backbone.$ = $;
var Marionette = require('backbone.marionette');
var SnaplistItemView = require('../views/snaplist-item.js');
var EmptySnaplistView = require('./empty-snaplist.js');
var template = require('../templates/snaplist.hbs');
var CONF = require('../config.js');

module.exports = Marionette.CompositeView.extend({

  childViewContainer: '.b-snaplist',

  template : function(model) {
    return template(model);
  },

  childView: SnaplistItemView,

  emptyView: EmptySnaplistView

});
