(function(root) {

  if (!window[root]) {
    window[root] = {

      /**
       * @param target {Object} the target object to copy into
       * @param source {Object} the source object to copy from
       * @param overwrite {Boolean} indicate if we should overwrite [Optional]
       * @param transform {Function} transformation function for each item [Optional]
       *
       * @return {Object} the modified target object
       */
      _copy: function(target, source, overwrite, transform) {
        for (var key in source) {
          if (overwrite || typeof target[key] === 'undefined') {
            target[key] = transform ? transform(source[key]) : source[key];
          }
        }

        return target;
      },

      /**
       * @param name {Object} fully qualified name
       * @param value {Object} value to set [Optional]
       *
       * @return {Object} the created object
       */
      _create: function(name, value) {
        var node  = window[root], names = name ? name.split('.') : [];

        for (var i = 0, len = names.length; i < len; i++) {
          var part = names[i], nso = node[part];

          if (!nso) {
            nso = (value && i+1 == names.length) ? value : {};
            node[part] = nso;
          }
          node = nso;
        }
        return node;
      },

      /**
       * Extends the parent namespace with a source object.
       * If the namespace target doesn't exist, it will be created automatically.
       *
       * Por ejemplo:
       *
       * To extend the root namespace:
       *
       * Thrust.extend({});
       * Thrust.extend(function() { return {}; });
       * Thrust.extend("", {});
       * Thrust.extend("", function() { return {}; });
       *
       * To extend children namespaces:
       *
       * Thrust.extend("Parent", {});
       * Thrust.extend("Parent", function() { return {}; });
       * Thrust.extend("Parent.Child", {});
       * Thrust.extend("Parent.Child", function() { return {}; });
       *
       * @param target {Object|Function|String} the target object to copy into
       * @param source {Object|Function} the source object to copy from
       * @param overwrite {Boolean} indicate if we should overwrite [Optional]
       *
       * @return {Object} the modified target object
       */
      extend: function(target, source, overwrite) {
        if (Object.prototype.toString.call(target) === '[object Object]' || typeof target === 'function') {
          source = target; target = this;
        }

        return this._copy(
          typeof target === 'string' ? this._create(target) : target,
          source,
          overwrite
        );
      }
    };
  }

})("Thrust");

Thrust.extend('IPC', {

  stub: function() {
    return void(0)
  },

});

/*
Simple DOM Manipulation primitives.
*/
JIBE.extend('DOM', {
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

