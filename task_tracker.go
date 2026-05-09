package main

type Task struct {
	ID          int
	title       string
	description string
	progress    string
}

type TaskTrackerInterface interface {
	addTask(newTask Task)
	updateTask(id int, newTask Task)
	deleteTask(id int)
	markTask(id int, newProgress string)
	listTask(currentProgress string) []Task
}

type TaskTracker struct {
	tasks []Task
}

func (t *TaskTracker) addTask(newTask Task) {
	t.tasks = append(t.tasks, newTask)
}

func (t *TaskTracker) updateTask(id int, newTask Task) {
	for i := range t.tasks {
		if t.tasks[i].ID == id {
			newTask.ID = id
			t.tasks[i] = newTask
			return
		}
	}
}

func (t *TaskTracker) deleteTask(id int) {
	for i := range t.tasks {
		if t.tasks[i].ID == id {
			t.tasks = append(t.tasks[:i], t.tasks[i+1:]...)
			return
		}
	}
}

func (t *TaskTracker) markTask(id int, newProgress string) {
	for i := range t.tasks {
		if t.tasks[i].ID == id {
			t.tasks[i].progress = newProgress
			return
		}
	}
}

func (t *TaskTracker) listTask(currentProgress string) []Task {
	var filteredTasks []Task
	for i := range t.tasks {
		if t.tasks[i].progress == currentProgress {
			filteredTasks = append(filteredTasks, t.tasks[i])
		}
	}
	return filteredTasks
}

func main() {
	var choice string
	for choice != "Exit" {
		switch choice {
		case "Add Task":

		case "Update Task":

		case "Delete Task":

		case "Mark Task":

		case "List Task":

		default:
			choice = "Exit"
		}
	}

	// scan for user input
	// reader := bufio.NewReader(os.Stdin)
	// line, err := reader.ReadString('\n')
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("read line: %s-\n", line)

}
