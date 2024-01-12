package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ARKNravi/HACKFEST-BE/repository"
	"github.com/gin-gonic/gin"
)

func CompleteTask(c *gin.Context) {
    profileIDStr := c.Param("profileID")
    profileID, err := strconv.Atoi(profileIDStr)
    if err != nil {
        c.JSON(400, gin.H{"error": "Invalid profile ID"})
        return
    }
    taskIDStr := c.Param("taskID")
    taskID, err := strconv.ParseUint(taskIDStr, 10, 32)
    if err != nil {
        c.JSON(400, gin.H{"error": "Invalid task ID"})
        return
    }
    repo := repository.NewTaskRepository()
    location, err := time.LoadLocation("Asia/Jakarta")
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    currentTime := time.Now().In(location)
    err = repo.CompleteTask(uint(profileID), uint(taskID), &currentTime)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"message": "Task completed successfully"})
}

func GetTasksByDate(c *gin.Context) {
    profileIDStr := c.Param("profileID")
    profileID, err := strconv.Atoi(profileIDStr)
    if err != nil {
        c.JSON(400, gin.H{"error": "Invalid profile ID"})
        return
    }
    dateStr := c.Query("date")
    date, err := time.Parse("2006-01-02", dateStr)
    if err != nil {
        c.JSON(400, gin.H{"error": "Invalid date format, use YYYY-MM-DD"})
        return
    }
    repo := repository.NewTaskRepository()
    tasks, err := repo.GetTasksByDate(uint(profileID), &date)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    // Prepare the response
    var response []map[string]interface{}
    for _, task := range tasks {
        item := map[string]interface{}{
            "Task ID":    task.TaskID,
            "Profile ID": task.ProfileID,
            "Bulan":      task.Date.Month(),
            "Tanggal":    task.Date.Day(),
            "Tahun":      task.Date.Year(),
            "Completed":  task.Completed,
            "Nama":  task.Task.Name,
        }
        response = append(response, item)
    }

    c.JSON(200, response)
}

func UndoTask(c *gin.Context) {
    profileIDStr := c.Param("profileID")
    profileID, err := strconv.Atoi(profileIDStr)
    if err != nil {
        c.JSON(400, gin.H{"error": "Invalid profile ID"})
        return
    }
    taskIDStr := c.Param("taskID")
    taskID, err := strconv.ParseUint(taskIDStr, 10, 32)
    if err != nil {
        c.JSON(400, gin.H{"error": "Invalid task ID"})
        return
    }
    repo := repository.NewTaskRepository()
    location, err := time.LoadLocation("Asia/Jakarta")
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    currentTime := time.Now().In(location)
    err = repo.UndoTask(uint(profileID), uint(taskID), &currentTime)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"message": "Task has been marked as not completed"})
}

func GetAllTasks(c *gin.Context) {
	repo := repository.NewTaskRepository()
	tasks, err := repo.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func GetCompletedTasks(c *gin.Context) {
    profileIDStr := c.Param("profileID")
    profileID, err := strconv.Atoi(profileIDStr)
    if err != nil {
        c.JSON(400, gin.H{"error": "Invalid profile ID"})
        return
    }

    repo := repository.NewTaskRepository()
    profile, err := repo.GetProfile(uint(profileID))
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    currentTime := time.Now()
    startYear := min(currentTime.Year(), profile.CreatedAt.Year())

    taskData := make(map[string]map[string]interface{})

    yearStart := time.Date(startYear, time.January, 1, 0, 0, 0, 0, time.UTC)
    midYear := time.Date(startYear, time.July, 1, 0, 0, 0, 0, time.UTC)
    endOfJune := midYear.AddDate(0, 0, -1)

    for t := yearStart; t.Year() <= currentTime.Year(); t = t.AddDate(1, 0, 0) {
        if t.Year() == currentTime.Year() || t.Year() == currentTime.Year()-1 {
            endDate := endOfJune
            if currentTime.Before(endOfJune) {
                endDate = currentTime
            }

            result, err := repo.GetCompletedTasks(profile.ID, &yearStart, &endDate)
            if err != nil {
                c.JSON(500, gin.H{"error": err.Error()})
                return
            }

            period := fmt.Sprintf("%s %d", getPeriodLabel(t), t.Year())
            taskData[period] = result

            if currentTime.After(midYear) && currentTime.Month() <= time.December {
  
                endOfYear := time.Date(t.Year(), time.December, 31, 23, 59, 59, 999999999, time.UTC)

                result, err = repo.GetCompletedTasks(profile.ID, &midYear, &endOfYear)
                if err != nil {
                    c.JSON(500, gin.H{"error": err.Error()})
                    return
                }

                period := fmt.Sprintf("%s %d", getPeriodLabel(midYear), midYear.Year())
                taskData[period] = result
            }
        }

        yearStart = yearStart.AddDate(1, 0, 0)
        midYear = midYear.AddDate(1, 0, 0)
        endOfJune = endOfJune.AddDate(1, 0, 0)
    }

    // Convert the map to a slice for consistent ordering
    var resultData []map[string]interface{}
    for period, data := range taskData {
        entry := make(map[string]interface{})
        entry[period] = data
        resultData = append(resultData, entry)
    }

    if len(resultData) == 0 {
        c.JSON(200, gin.H{"message": "No completed tasks found for the given periods"})
    } else {
        c.JSON(200, gin.H{"data": resultData})
    }
}

func getPeriodLabel(t time.Time) string {
    if t.Month() >= time.January && t.Month() <= time.June {
        return "Jan - Jun"
    } else if t.Month() >= time.July && t.Month() <= time.December {
        return "Jul - Dec"
    }
    return ""
}
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}






