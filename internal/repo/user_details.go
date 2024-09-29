package repo

import (
	"nexbit/models"
	_ "github.com/lib/pq"
	"errors"
)

func SaveUser(dbs DBService, user models.User) error {
	result, err := dbs.NamedExec(`
	INSERT
	INTO
	users
	(name, email)
	VALUES
	(:name, :email)`, user)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows were inserted")
	}

	return nil
}

func SaveUserPreferences(dbs DBService, preferences models.UserPreferences) error {

	result, err := dbs.NamedExec(`
	INSERT
	INTO
	user_preferences
	(user_id, question_id, answer)
	VALUES
	(:user_id, :question_id, :answer)`, preferences)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows were inserted")
	}

	return nil
}
