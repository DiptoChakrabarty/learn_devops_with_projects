const express = require("express");
const bodyParser = require("body-parser");

const app = express();
const port = process.env.PORT || 3000;

app.use(bodyParser.json());

const tasks = [];

app.post("/tasks", (req, res) => {
  const task = req.body;
  tasks.push(task);
  res.status(201).json(task);
});

app.get("/tasks", (req, res) => {
  res.json(tasks);
});

app.put("/tasks/:id", (req, res) => {
  const taskId = parseInt(req.params.id);
  const updatedTask = req.body;

  const taskIndex = tasks.findIndex((task) => task.id === taskId);

  if (taskIndex !== -1) {
    tasks[taskIndex] = { ...tasks[taskIndex], ...updatedTask };
    res.json(tasks[taskIndex]);
  } else {
    res.status(404).json({ message: "Task not found" });
  }
});

app.delete("/tasks/:id", (req, res) => {
  const taskId = parseInt(req.params.id);

  const taskIndex = tasks.findIndex((task) => task.id === taskId);

  if (taskIndex !== -1) {
    tasks.splice(taskIndex, 1);
    res.status(204).send();
  } else {
    res.status(404).json({ message: "Task not found" });
  }
});

app.listen(port, () => {
  console.log(`Todo app is running on port ${port}`);
});
