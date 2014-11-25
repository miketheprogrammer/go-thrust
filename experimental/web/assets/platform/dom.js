/*
Simple DOM Manipulation primitives.
*/
Thrust.extend('DOM', {
  _callbacks: {},

  hasClass: function(e, className) {
    return (' '+ e.className +' ').indexOf(' '+ className +' ') >= 0;
  },

  addClass: function(e, className) {
    if (!this.hasClass(e, className)) {
      e.className = e.className + ' ' + className;
    }
  },

  setStyle: function(e, property, value) {
    e.style[property] = value;
  },

  setStyles: function(e, styles) {
    for (var style in styles) {
      this.setStyle(e, style, styles[style]);
    }
  },

  createWebView: function(opts) {
    return void(0)
  }
});
