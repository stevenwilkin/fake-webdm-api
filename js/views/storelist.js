// storelist view
var Snaplist = require('./snaplist.js');
var StorelistItemView = require('./storelist-item.js');
var template = require('../templates/storelist.hbs');

module.exports = Snaplist.extend({
  childView: StorelistItemView,

  template : function(model) {
    return template(model);
  }
});
