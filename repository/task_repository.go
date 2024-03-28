package database

import (
	"time"
	"todo-list/domain"

	"github.com/upper/db/v4"
)

const TasksTableName = "tasks"

type task struct {
	Id          uint64            `db:"id,omitempty"`
	User_id     uint64            `db:"user_id"`
	Title       string            `db:"title"`
	Description *string           `db:"description"`
	Status      domain.TaskStatus `db:"status"`
	Date        *time.Time        `db:"date"`
}

type TaskRepository struct {
	coll db.Collection
	sess db.Session
}

func NewTaskRepository(sess db.Session) TaskRepository {
	return TaskRepository{
		coll: sess.Collection(TasksTableName),
		sess: sess,
	}
}

func (r TaskRepository) Save(u domain.Task) (domain.Task, error) {
	usr := r.mapDomainToModel(u)
	err := r.coll.InsertReturning(&usr)
	if err != nil {
		return domain.Task{}, err
	}
	newUser := r.MapModelToDomain(usr)
	return newUser, err
}

func (r TaskRepository) mapDomainToModel(u domain.Task) task {
	return task{
		Id:          u.Id,
		User_id:     u.User_id,
		Title:       u.Title,
		Description: u.Description,
		Status:      u.Status,
		Date:        u.Date,
	}
}

func (r TaskRepository) MapModelToDomain(u task) domain.Task {
	return domain.Task{
		Id:          u.Id,
		User_id:     u.User_id,
		Title:       u.Title,
		Description: u.Description,
		Status:      u.Status,
		Date:        u.Date,
	}
}

func (r TaskRepository) FinById(id uint64) (domain.Task, error) {
	var u task
	err := r.coll.Find("id = ?", id).One(&u)
	if err != nil {
		return domain.Task{}, err
	}
	usr := r.MapModelToDomain(u)
	return usr, nil
}

// ф-ія оновлення
func (r TaskRepository) Update(u domain.Task) (domain.Task, error) {
	usr := r.mapDomainToModel(u)
	err := r.coll.Find("id = ?", u.Id).Update(&usr)
	if err != nil {
		return domain.Task{}, err
	}
	updatedUser := r.MapModelToDomain(usr)
	return updatedUser, nil
}

func (r TaskRepository) Delete(id uint64) error {
	err := r.coll.Find("id = ?", id).Delete()
	return err
}
