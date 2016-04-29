// search layout view
var Backbone = require('backbone');
var Marionette = require('backbone.marionette');
var BaskView = require('./storelist.js');
var template = require('../templates/store.hbs');

module.exports = Backbone.Marionette.LayoutView.extend({

  className: 'b-search',

  template : function(model) {
    model.title = 'Search results';
    return template(model);
  },

  onBeforeShow: function() {
    this.showChildView('resultsRegion', new BaskView({
      collection: this.collection
    }));
  },

  regions: {
    resultsRegion: '.region-snaplist'
  }
});
