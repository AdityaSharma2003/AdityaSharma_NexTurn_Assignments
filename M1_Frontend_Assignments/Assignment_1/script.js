const taskInput = document.getElementById('taskInput');
const addTaskBtn = document.getElementById('addTaskBtn');
const taskList = document.getElementById('taskList');
const noTasksMessage = document.getElementById('noTasksMessage');

const taskCounter = document.createElement('p');
taskCounter.id = 'taskCounter';
taskCounter.textContent = 'Pending tasks: 0';
taskList.parentElement.insertBefore(taskCounter, taskList);

addTaskBtn.addEventListener('click', function () {
  const taskText = taskInput.value;

  if (taskText === '') {
    alert('Please enter a task.');
    return;
  }

  noTasksMessage.style.display = 'none';

  const taskItem = createTaskItem(taskText);
  taskList.appendChild(taskItem);

  taskInput.value = '';

  updateTaskCounter();
  saveData();
});

taskList.addEventListener('click', function (e) {
  if (e.target.tagName === 'LI') {
    e.target.classList.toggle('completed');
    updateTaskCounter();
    saveData();
  }
});

function createTaskItem(taskText) {
  const taskItem = document.createElement('li');
  taskItem.textContent = taskText;
  taskItem.draggable = true;

  const editBtn = document.createElement('button');
  editBtn.textContent = 'Edit';
  editBtn.addEventListener('click', function () {
    enableEditing(taskItem);
  });

  const deleteBtn = document.createElement('button');
  deleteBtn.textContent = 'Delete';
  deleteBtn.addEventListener('click', function () {
    taskList.removeChild(taskItem);

    if (taskList.childElementCount === 0) {
      noTasksMessage.style.display = 'block';
    }

    updateTaskCounter();
    saveData();
  });

  taskItem.appendChild(editBtn);
  taskItem.appendChild(deleteBtn);

  taskItem.addEventListener('dragstart', handleDragStart);
  taskItem.addEventListener('dragover', handleDragOver);
  taskItem.addEventListener('drop', handleDrop);
  taskItem.addEventListener('dragend', handleDragEnd);

  return taskItem;
}

function enableEditing(taskItem) {
  const originalText = taskItem.firstChild.textContent;
  taskItem.innerHTML = '';

  const inputField = document.createElement('input');
  inputField.type = 'text';
  inputField.value = originalText;
  inputField.style.flexGrow = '1';

  const saveBtn = document.createElement('button');
  saveBtn.textContent = 'Save';
  saveBtn.addEventListener('click', function () {
    const updatedText = inputField.value.trim();
    if (updatedText) {
      taskItem.innerHTML = '';
      taskItem.textContent = updatedText;

      taskItem.appendChild(createEditButton(taskItem));
      taskItem.appendChild(createDeleteButton(taskItem));

      taskItem.draggable = true;
      taskItem.addEventListener('dragstart', handleDragStart);
      taskItem.addEventListener('dragover', handleDragOver);
      taskItem.addEventListener('drop', handleDrop);
      taskItem.addEventListener('dragend', handleDragEnd);

      saveData();
    } else {
      alert('Task cannot be empty.');
    }
  });

  const cancelBtn = document.createElement('button');
  cancelBtn.textContent = 'Cancel';
  cancelBtn.addEventListener('click', function () {
    taskItem.innerHTML = originalText;
    taskItem.appendChild(createEditButton(taskItem));
    taskItem.appendChild(createDeleteButton(taskItem));
  });

  taskItem.appendChild(inputField);
  taskItem.appendChild(saveBtn);
  taskItem.appendChild(cancelBtn);
}

function createEditButton(taskItem) {
  const editBtn = document.createElement('button');
  editBtn.textContent = 'Edit';
  editBtn.addEventListener('click', function () {
    enableEditing(taskItem);
  });
  return editBtn;
}

function createDeleteButton(taskItem) {
  const deleteBtn = document.createElement('button');
  deleteBtn.textContent = 'Delete';
  deleteBtn.addEventListener('click', function () {
    taskList.removeChild(taskItem);

    if (taskList.childElementCount === 0) {
      noTasksMessage.style.display = 'block';
    }

    updateTaskCounter();
    saveData();
  });
  return deleteBtn;
}

function handleDragStart(e) {
  e.dataTransfer.setData('text/plain', e.target.id);
  e.target.classList.add('dragging');
}

function handleDragOver(e) {
  e.preventDefault();
  const draggingItem = document.querySelector('.dragging');
  const currentHoveredItem = e.target;

  if (currentHoveredItem.tagName === 'LI' && currentHoveredItem !== draggingItem) {
    const bounding = currentHoveredItem.getBoundingClientRect();
    const offset = e.clientY - bounding.top;

    if (offset > bounding.height / 2) {
      currentHoveredItem.after(draggingItem);
    } else {
      currentHoveredItem.before(draggingItem);
    }
  }
}

function handleDrop(e) {
  e.preventDefault();
  e.target.classList.remove('dragging');
  saveData();
}

function handleDragEnd(e) {
  e.target.classList.remove('dragging');
}

function saveData() {
  localStorage.setItem('data', taskList.innerHTML);
}

function showData() {
  const savedData = localStorage.getItem('data');

  if (savedData) {
    taskList.innerHTML = savedData;

    const taskItems = taskList.querySelectorAll('li');

    taskItems.forEach((taskItem) => {
      const deleteBtn = taskItem.querySelector('button:nth-child(2)');
      const editBtn = taskItem.querySelector('button:nth-child(1)');

      deleteBtn.addEventListener('click', function () {
        taskList.removeChild(taskItem);

        if (taskList.childElementCount === 0) {
          noTasksMessage.style.display = 'block';
        }

        updateTaskCounter();
        saveData();
      });

      editBtn.addEventListener('click', function () {
        enableEditing(taskItem);
      });

      taskItem.draggable = true;
      taskItem.addEventListener('dragstart', handleDragStart);
      taskItem.addEventListener('dragover', handleDragOver);
      taskItem.addEventListener('drop', handleDrop);
      taskItem.addEventListener('dragend', handleDragEnd);
    });
  }

  updateTaskCounter();
}

function updateTaskCounter() {
  const pendingTasks = taskList.querySelectorAll('li:not(.completed)').length;
  taskCounter.textContent = `Pending tasks: ${pendingTasks}`;
}

showData();
