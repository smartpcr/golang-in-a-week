package store

import v1 "webapi/schema/v1"

type UserRepository struct {
	*OrmRepository[v1.User]
}

type ProjectRepository struct {
	*OrmRepository[v1.Project]
}

type TaskRepository struct {
	*OrmRepository[v1.Task]
}
