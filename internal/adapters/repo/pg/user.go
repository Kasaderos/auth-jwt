package pg

import (
	"birthday-bot/internal/domain/entities"
	"context"
)

func (d *St) UserGet(ctx context.Context, id int64) (*entities.UserSt, error) {
	var result entities.UserSt

	err := d.db.QueryRow(ctx, `
		select
			id,
  			name,
  			email,
  			password,
			birthdate
		from users
		where id = $1
	`, id).Scan(
		&result.ID,
		&result.Name,
		&result.Email,
		&result.Password,
		&result.Birthdate,
	)

	return &result, err
}

func (d *St) UserUpdate(ctx context.Context, id int64, obj *entities.UserCUSt) error {
	fields := d.getUserFields(obj)
	cols := d.tPrepareFieldsToUpdate(fields)

	fields["cond_id"] = id

	return d.db.ExecM(ctx, `
		update users 
		set `+cols+`
		where id = ${cond_id}
	`, fields)
}

func (d *St) getUserFields(obj *entities.UserCUSt) map[string]any {
	result := map[string]any{}

	if obj.Name != nil {
		result["name"] = *obj.Name
	}

	if obj.Email != nil {
		result["email"] = *obj.Email
	}

	if obj.Password != nil {
		result["password"] = *obj.Password
	}

	if obj.Birthdate != nil {
		result["birthdate"] = *obj.Birthdate
	}

	return result
}

func (d *St) UserCreate(ctx context.Context, obj *entities.UserCUSt) (int64, error) {
	fields := d.getUserFields(obj)
	cols, values := d.tPrepareFieldsToCreate(fields)

	var newId int64

	err := d.db.QueryRowM(ctx, `
		insert into users (`+cols+`)
		values (`+values+`)
		returning id
	`, fields).Scan(&newId)

	return newId, err
}

func (d *St) UserDelete(ctx context.Context, id int64) error {
	return d.db.Exec(ctx, `
		delete from users 
		where id = $1
	`, id)
}
