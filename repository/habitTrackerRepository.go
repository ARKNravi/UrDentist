package repository

import (
	"errors"
	"time"

	"github.com/ARKNravi/HACKFEST-BE/database"
	"github.com/ARKNravi/HACKFEST-BE/model"
	"gorm.io/gorm"
)

type TaskRepository struct{}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

func (r *TaskRepository) CompleteTask(profileID uint, taskID uint, completionTime *time.Time) error {
    db, err := database.Connect()
    if err != nil {
        return err
    }

    var existingCompletion model.TaskCompletion
    if err := db.Where("profile_id = ? AND task_id = ? AND DATE(date) = ?", profileID, taskID, completionTime.Format("2006-01-02")).First(&existingCompletion).Error; err != nil {
        if !errors.Is(err, gorm.ErrRecordNotFound) {
            return err
        }
    } else {
        if existingCompletion.Completed {
            return errors.New("this task has already been completed today")
        } else {
            existingCompletion.Completed = true
            db.Save(&existingCompletion)
            return nil
        }
    }

    completion := model.TaskCompletion{
        ProfileID: profileID,
        TaskID:    taskID,
        Completed: true,
        Date:      completionTime,
    }
    result := db.Create(&completion)
    return result.Error
}

func (r *TaskRepository) GetTasksByDate(profileID uint, date *time.Time) ([]model.TaskCompletion, error) {
    db, err := database.Connect()
    if err != nil {
        return nil, err
    }
    var tasks []model.TaskCompletion
    result := db.Preload("Profile").Preload("Profile.User").Preload("Task").Where("profile_id = ? AND DATE(date) = ?", profileID, date.Format("2006-01-02")).Find(&tasks)
    return tasks, result.Error
}

func (r *TaskRepository) UndoTask(profileID uint, taskID uint, completionTime *time.Time) error {
    db, err := database.Connect()
    if err != nil {
        return err
    }

    var completion model.TaskCompletion
    if err := db.Where("profile_id = ? AND task_id = ? AND DATE(date) = ?", profileID, taskID, completionTime.Format("2006-01-02")).First(&completion).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return errors.New("this task has not been completed today")
        }
        return err
    }

    completion.Completed = false
    result := db.Save(&completion)
    return result.Error
}

func (r *TaskRepository) GetAllTasks() ([]model.Task, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	var tasks []model.Task
	result := db.Find(&tasks)
	return tasks, result.Error
}

func (r *TaskRepository) GetProfileCreationDate(profileID uint) (time.Time, error) {
    db, err := database.Connect()
    if err != nil {
        return time.Time{}, err
    }

    var profileCreationDate time.Time
    err = db.Model(&model.Profile{}).
        Select("created_at").
        Where("id = ?", profileID).
        Scan(&struct {
            CreatedAt time.Time
        }{}).
        Error

    if err != nil {
        return time.Time{}, err
    }

    return profileCreationDate, nil
}

func (r *TaskRepository) GetCompletedTasks(profileID uint, startDate *time.Time, endDate *time.Time) (map[string]interface{}, error) {
    db, err := database.Connect()
    if err != nil {
        return nil, err
    }

    var completions []model.TaskCompletion
    err = db.Where("profile_id = ? AND completed = true AND date BETWEEN ? AND ?", profileID, startDate.Format("2006-01-02"), endDate.Format("2006-01-02")).Find(&completions).Error
    if err != nil {
        return nil, err
    }

    totalCompletedDays := 0
    totalCompletedPoints := 0

    uniqueDays := make(map[string]map[uint]bool)
    for _, completion := range completions {
        dateStr := completion.Date.Format("2006-01-02")

        if _, ok := uniqueDays[dateStr]; !ok {
            uniqueDays[dateStr] = make(map[uint]bool)
            totalCompletedDays++
        }

        uniqueDays[dateStr][completion.TaskID] = true

        if len(uniqueDays[dateStr]) == 5 {
            for taskID := range uniqueDays[dateStr] {
                task := model.Task{}
                db.First(&task, taskID)
                totalCompletedPoints += task.Points
            }
            uniqueDays[dateStr] = make(map[uint]bool)
        }
    }

    result := map[string]interface{}{
        "completedTasks": len(completions),
        "DayCompleted":   totalCompletedDays,
        "TotalPoints":    totalCompletedPoints,
    }

    return result, nil
}


func (r *TaskRepository) GetProfile(profileID uint) (model.Profile, error) {
    db, err := database.Connect()
    if err != nil {
        return model.Profile{}, err
    }

    var profile model.Profile
    err = db.First(&profile, profileID).Error
    return profile, err
}