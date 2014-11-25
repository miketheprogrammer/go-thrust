Thrust.extend("bus", {
  _dispatcher: {
    registry: [],
    dispatch: function(commandresponse) {
      if ('string' === typeof commandresponse) {
        commandresponse = JSON.parse(commandresponse);
      }
      registry.forEach(function(handler) {
        if (handler.type == '*') {
           handler.handle(commandresponse);
           return;
        }
        if (handler.type == commandresponse._type) {
           handler.handle(commandresponse);
        }
      })
    }
  },
  registerByType: function (type, fn) {
    this._dispatcher.registry.push({
      type: type,
      handle: fn
    })
  },
  registerByTargetID: function (id, fn) {
    var matchFn = function (commandresponse) {
      if (commandresponse._target_id === id) {
        fn(commandresponse)
      }
    }
    this._dispatcher.registry.push({
      type: "*",
      handle: matchFn
    })
  }

})


Thrust.bus.registerByType("event", function (commandresponse) {

})

Thrust.bus.registerByType("reply", function (commandresponse) {

})

Thrust.bus.registerByType("invoke", function (commandresponse) {

})

Thrust.bus.registerByTargetID(1, function (commandresponse) {

})