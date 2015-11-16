/**
 * @file task_test.go
 * @author Mikhail Klementyev jollheef<AT>riseup.net
 * @license GNU GPLv3
 * @date November, 2015
 * @brief test work with task table
 */

package db

import (
	"errors"
	"fmt"
	"testing"
)

func TestCreateTaskTable(*testing.T) {

	db, err := InitDatabase(dbPath)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = createTaskTable(db)
	if err != nil {
		panic(err)
	}
}

func TestAddTask(*testing.T) {

	db, err := InitDatabase(dbPath)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	task := Task{255, "n", "d", 10, 100, 50, true, "f", 1, 10, true}

	err = AddTask(db, &task)
	if err != nil {
		panic(err)
	}

	if task.ID != 1 {
		panic(errors.New("Task id not correct"))
	}
}

func TestGetTasks(*testing.T) {

	db, err := InitDatabase(dbPath)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	ntasks := 150

	for i := 0; i < ntasks; i++ {

		task := Task{ID: 255, Name: fmt.Sprintf("%d", i)}

		err = AddTask(db, &task)
		if err != nil {
			panic(err)
		}
	}

	tasks, err := GetTasks(db)
	if err != nil {
		panic(err)
	}

	if len(tasks) != ntasks {
		panic(errors.New("Mismatch get tasks length"))
	}

	for i := 0; i < ntasks; i++ {

		if tasks[i].Name != fmt.Sprintf("%d", i) && tasks[i].ID != i {
			panic(errors.New("Get invalid task"))
		}
	}
}
