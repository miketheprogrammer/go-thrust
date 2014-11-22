var Api = window.Api = {}

Api.Get = function (key) {
  $.get("/db/get/"+key, function (err, res) {
    
  })
}
