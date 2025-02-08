// easy 2432 ELO 1266
// The employee that worked

func hardestWorker(n int, logs [][]int) int {

    m:=len(logs)
    max:=0
    task:=0
    employee:=-1

    for i:=0;i<m;i++ {
        //fmt.Println(logs[i][1])
        if i == 0 {
            task = logs[0][1]
        } else {
            task = logs[i][1] - logs[i-1][1]
        }
        if task > max {
            max = task
            employee = logs[i][0]
        } else if task == max {
            if logs[i][0] < employee {
                employee = logs[i][0]
            }
        }
    }


    //fmt.Println(max)

    return employee
    
}

