package main

import "fmt"

// T1-T2
// T2-T3
// T4-T6
// T2-T4
// T6-T7

// T1 [T2]
// T2 [T3, T4]
// T3 []
// T4 [T6]
// T6 [T7]
// T7 []

// first flatern the unqiue tasks
// and mark its dependcies
// map[Task]{
// 		Rank: 1,
// 		Dependcies: [Task,...]
// 		IsDone: bool
// }

// we need list [] of tasks ordered by execution rank/position

// check if the dependend tasks completed, then make this into the queue

type taskProp struct {
	Rank       int
	Dependcies []string
	IsDone     bool
}

func main() {
	input := [][]string{
		{"T1", "T2"},
		{"T2", "T3"},
		{"T4", "T6"},
		{"T2", "T4"},
		{"T6", "T7"},
		{"T5", "T1"},
	}

	tasksMap := make(map[string]taskProp)

	uniqueTasks := []string{}
	// first flatern the unqiue tasks
	for _, tasks := range input {
		t := tasks[0]
		dt := tasks[1]

		tprop, ok := tasksMap[t]
		if !ok {
			tasksMap[t] = taskProp{
				Rank: 0,
				Dependcies: []string{
					dt,
				},
				IsDone: false,
			}
			uniqueTasks = append(uniqueTasks, t)
		} else {
			tprop.Dependcies = append(tprop.Dependcies, dt)
			tasksMap[t] = tprop
		}

		//  add depended task into the registary
		_, ok = tasksMap[dt]
		if !ok {
			tasksMap[dt] = taskProp{
				Rank:       0,
				Dependcies: []string{},
				IsDone:     false,
			}
			uniqueTasks = append(uniqueTasks, dt)
		}
	}

	// find the rank
	totalTasks := len(tasksMap)

	fmt.Println(totalTasks, uniqueTasks)
	fmt.Printf("%v\n", tasksMap)

	// finding the order of executions
	// we need list [] of tasks ordered by execution rank/position

	// 	T3,T7, T6, T4, T2, T1
	//  [T3 T7 T6 T4 T2 T1]

	tasksOrder := []string{}
	for {
		for _, t := range uniqueTasks { // fixed len(6)
			// 1. check if has 0 dependencies
			taskProp := tasksMap[t]
			if taskProp.IsDone {
				continue
			}

			if len(taskProp.Dependcies) == 0 {
				// add it to the queue
				tasksOrder = append(tasksOrder, t)

				// mark this has done
				taskProp.IsDone = true
				tasksMap[t] = taskProp

				continue
			}

			// 2. check if the dependend tasks completed, then make this into the queue
			allSubTasksDone := true
			for _, subT := range taskProp.Dependcies {
				subTaskProp := tasksMap[subT]
				// if any one not completed then no
				if !subTaskProp.IsDone {
					allSubTasksDone = false
				}
			}

			// check if the sub tasks are done
			if allSubTasksDone {
				// add it to the queue
				tasksOrder = append(tasksOrder, t)

				// mark this has done
				taskProp.IsDone = true
				tasksMap[t] = taskProp
			}
		}

		// fmt.Println("Current Queue: ", tasksOrder, len(tasksOrder))
		if len(tasksOrder) >= totalTasks {
			break
		}
	}

	fmt.Println("Order of tasks: ", tasksOrder)
}
