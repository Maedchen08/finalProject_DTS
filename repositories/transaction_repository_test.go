package repositories

import (
	"AntarJemput-Be-C/models"
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/tj/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	repo *TransactionRepo
	mock sqlmock.Sqlmock
	db   *sql.DB
	err  error
)

func TestMain(m *testing.M) {
	db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("Error intialize DATA-DOG : %v", err)
	}
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{Conn: db, SkipInitializeWithVersion: true}), &gorm.Config{})
	repo = NewTransactionRepository(gormDB)
	exitVal := m.Run()
	defer db.Close()
	os.Exit(exitVal)

}

func TestTransactionRepo_Save(t *testing.T) {
	t.Run("Should return success", func(t *testing.T) {
		query := "INSERT INTO `transaction` (`created_at`,`updated_at`,`deleted_at`,`type_transaction_id`,`customers_id`,`agents_id`,`address`,`transaction_prov_id`,`transaction_regency_id`,`transaction_district_id`,`amount`,`status_transaction`,`rating`,`rating_comment`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
		payload := &models.Transactions{
			TypeTransactionId: 1,
			CustomerId:        1,
			AgentId:           1,
			Address:           "Jalan H. Dahlan",
			ProvinceId:        32,
			RegencyId:         3204,
			DistrictId:        320490,
			Amount:            1000000,
			StatusTransaction: 3,
			Rating:            5,
			RatingComment:     "Good Service",
		}
		expectedId := 1

		mock.ExpectBegin()
		mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(nil)
		mock.ExpectCommit()

		id, err := repo.Save(payload)
		assert.NoError(t, err)
		assert.Equal(t, id, expectedId)

	})
}

//failed, tolong dicoba lagi
func TestTransactionRepo_GetById(t *testing.T) {
	t.Run("Should return success", func(t *testing.T) {
		outputTransactionDetail := models.Transactions{
			TypeTransactionId: 1,
			CustomerId:        1,
			AgentId:           1,
			Address:           "Jalan H. Dahlan",
			ProvinceId:        32,
			RegencyId:         3204,
			DistrictId:        3204190,
			Amount:            10000000,
			StatusTransaction: 3,
			Rating:            5,
			RatingComment:     "Good Service",
		}
		inputId := 1
		query := "SELECT `type_transaction_id`,`customers_id`,`agents_id`,`address`,`transaction_prov_id`,`transaction_regency_id`,`transaction_district_id`,`amount`,`status_transaction`,`rating`,`rating_comment` FROM transaction WHERE id=?"

		mock.ExpectQuery(query).
			WithArgs(inputId).
			WillReturnRows(sqlmock.NewRows([]string{"type_transaction_id", "customers_id", "agents_id", "address", "transaction_prov_id", "transaction_regency_id", "transaction_district_id", "amount", "status_transaction", "rating", "rating_comment"}).
				AddRow(outputTransactionDetail.TypeTransactionId, outputTransactionDetail.CustomerId, outputTransactionDetail.AgentId, outputTransactionDetail.Address, outputTransactionDetail.ProvinceId, outputTransactionDetail.RegencyId, outputTransactionDetail.DistrictId, outputTransactionDetail.Amount, outputTransactionDetail.Rating, outputTransactionDetail.RatingComment)).
			WillReturnError(nil)
		actual, err := repo.GetById(inputId)
		assert.Nil(t, err)
		assert.Equal(t, actual, outputTransactionDetail)
	})
}

//failed, tolong dicobaa lagi
func TestTransactionRepo_GetAll(t *testing.T) {
	t.Run("Should return success", func(t *testing.T) {
		outputTransactionDetail := []models.Transactions{
			{
				TypeTransactionId: 1,
				CustomerId:        1,
				AgentId:           1,
				Address:           "Jalan H. Dahlan",
				ProvinceId:        32,
				RegencyId:         3204,
				DistrictId:        320490,
				Amount:            1000000,
				StatusTransaction: 3,
				Rating:            5,
				RatingComment:     "Good Service",
			},
			{
				TypeTransactionId: 1,
				CustomerId:        1,
				AgentId:           1,
				Address:           "Jalan H. Dahlan",
				ProvinceId:        32,
				RegencyId:         3204,
				DistrictId:        320490,
				Amount:            1000000,
				StatusTransaction: 3,
				Rating:            5,
				RatingComment:     "Good Service",
			},
			{
				TypeTransactionId: 1,
				CustomerId:        1,
				AgentId:           1,
				Address:           "Jalan H. Dahlan",
				ProvinceId:        32,
				RegencyId:         3204,
				DistrictId:        320490,
				Amount:            1000000,
				StatusTransaction: 3,
				Rating:            5,
				RatingComment:     "Good Service",
			},
		}
		queryExpected := "SELECT * FROM `transaction` WHERE `transaction`.`deleted_at` IS NULL"

		mock.ExpectQuery(queryExpected).
			WillReturnRows(sqlmock.NewRows([]string{"created_at", "updated_at", "deleted_at", "type_transaction_id", "customers_id", "agents_id", "address", "transaction_prov_id", "transaction_regency_id", "transaction_district_id", "amount", "status_transaction", "rating", "rating_comment"}).
				AddRow(outputTransactionDetail)).WillReturnError(nil)

		actual, err := repo.GetAll()
		assert.NoError(t, err)

		assert.Equal(t, actual, outputTransactionDetail)
	})
}

func TestTransactionRepo_DeleteTransaction(t *testing.T) {
	t.Run("Should delete a transaction successfully", func(t *testing.T) {
		queryExpected := "UPDATE `transaction` SET `deleted_at`=? WHERE id =? AND `transaction`.`deleted_at` IS NULL"
		inputId := 1
		mock.ExpectBegin()
		mock.ExpectExec(queryExpected).
			WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(nil)
		mock.ExpectCommit()

		err := repo.DeleteTransaction(inputId)
		assert.NoError(t, err)
	})
}

func TestTransactionRepo_ChangeDone(t *testing.T) {
	t.Run("Should return success change status done", func(t *testing.T) {
		inputStatus := models.Transactions{
			Id: 1,
		}
		queryExpected := "UPDATE `transaction` SET `status_transaction`=?,`updated_at`=? WHERE `id` = ? AND `transaction`.`deleted_at` IS NULL"

		mock.ExpectBegin()
		mock.ExpectExec(queryExpected).
			WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(nil)
		mock.ExpectCommit()

		err := repo.ChangeDone(&inputStatus)
		assert.NoError(t, err)
	})
}

func TestTransactionRepo_ChangeReject(t *testing.T) {
	t.Run("Should return success change status reject", func(t *testing.T) {
		inputStatus := models.Transactions{
			Id: 2,
			// StatusTransaction: 2,
		}
		queryExpected := "UPDATE `transaction` SET `status_transaction`=?,`updated_at`=? WHERE `id` = ? AND `transaction`.`deleted_at` IS NULL"

		mock.ExpectBegin()
		mock.ExpectExec(queryExpected).
			WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(nil)
		mock.ExpectCommit()

		err := repo.ChangeReject(&inputStatus)
		assert.NoError(t, err)
	})
}

func TestTransactionRepo_ChangeConfirmed(t *testing.T) {
	t.Run("Should return success change status confirm", func(t *testing.T) {
		inputStatus := models.Transactions{
			Id: 2,
			// StatusTransaction: 2,
		}
		queryExpected := "UPDATE `transaction` SET `status_transaction`=?,`updated_at`=? WHERE `id` = ? AND `transaction`.`deleted_at` IS NULL"

		mock.ExpectBegin()
		mock.ExpectExec(queryExpected).
			WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(nil)
		mock.ExpectCommit()

		err := repo.ChangeConfirmed(&inputStatus)
		assert.NoError(t, err)
	})
}

func TestTransactionRepo_RatingTransaction(t *testing.T) {
	t.Run("should return success", func(t *testing.T) {
		inputRating := models.Transactions{
			Rating:        5,
			RatingComment: "Good Service",
		}
		queryExpected := "UPDATE `transaction` SET `rating`=?,`rating_comment`=?,`updated_at`=? WHERE `id` = ? AND `transaction`.`deleted_at` IS NULL"
		mock.ExpectBegin()
		mock.ExpectExec(queryExpected).
			WillReturnResult(sqlmock.NewResult(1, 1)).
			WillReturnError(nil)
		mock.ExpectCommit()

		err := repo.RatingTransaction(&inputRating)
		assert.NoError(t, err)
	})
}

func TestTransactionRepo_GetByAgentId(t *testing.T) {
	t.Run("Should return success", func(t *testing.T) {
		outputTransactionDetail := models.Transactions{
			TypeTransactionId: 1,
			CustomerId:        1,
			AgentId:           1,
			Address:           "Jalan H. Dahlan",
			ProvinceId:        32,
			RegencyId:         3204,
			DistrictId:        320490,
			Amount:            1000000,
			StatusTransaction: 3,
			Rating:            5,
			RatingComment:     "Good Service",
		}

		inputId := 1
		query := "SELECT `type_transaction_id`,`customers_id`,`agents_id`,`address`,`transaction_prov_id`,`transaction_regency_id`,`transaction_district_id`,`amount`,`status_transaction`,`rating`,`rating_comment` FROM transaction WHERE agents_id=?"

		mock.ExpectQuery(query).
			WithArgs(inputId).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "type_transaction_id", "customers_id", "agents_id", "address", "transaction_prov_id", "transaction_regency_id", "transaction_district_id", "amount", "status_transaction", "rating", "rating_comment"}).
				AddRow(outputTransactionDetail.Id, outputTransactionDetail.CreatedAt, outputTransactionDetail.UpdatedAt, outputTransactionDetail.DeletedAt, outputTransactionDetail.TypeTransactionId, outputTransactionDetail.CustomerId, outputTransactionDetail.AgentId, outputTransactionDetail.Address, outputTransactionDetail.ProvinceId, outputTransactionDetail.RegencyId, outputTransactionDetail.DistrictId, outputTransactionDetail.Amount, outputTransactionDetail.Rating, outputTransactionDetail.RatingComment)).
			WillReturnError(nil)

		actual, err := repo.GetByAgentId(inputId)
		assert.Nil(t, err)
		assert.Equal(t, actual, outputTransactionDetail)
	})
}
