
window.TaskManager = (function () {
  var taskManager = Object.create(null);
  taskManager.Tasks = [];
  taskManager.render = function () {
    $('#task_list').empty()
    taskManager.Tasks.forEach(function (task) {
      $('#task_list').append(task.render())
    });
  }
  taskManager.addTask = function (task) {
    taskManager.Tasks.push(task)
  }
  taskManager.remove = function (id) {
    taskManager.Tasks = taskManager.Tasks.filter(function(task) {
      if (task.id == id) {
        return false
      } else {
        return true
      }
    })
    taskManager.render()
  }
  return taskManager
})()



window.Task = function (id, description) {
  var task = {
    id: id,
    desc: description
  }
  TaskManager.addTask(task);
  task.render = function () {
    return window.Task.Template
               .replace(/{id}/g, task.id)
               .replace(/{desc}/g, task.desc);
  }
  task.remove = function () {
    TaskManager.remove(task.id);
  }


  task.html = task.render()

  return task
}
window.Task.Template =
"<div id=\"task_item_template\">"
  +"<div id=\"{id}\"><p id=\"{id}\">{desc}</p></div>"
+"</div>"