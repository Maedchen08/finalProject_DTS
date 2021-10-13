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
	db,mock,err=sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("Error intialize DATA-DOG : %v",err)
	}
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{Conn: db, SkipInitializeWithVersion: true}), &gorm.Config{})
	repo = NewTransactionRepository(gormDB)
	exitVal := m.Run()
	defer db.Close()
	os.Exit(exitVal)
	

}

func TestTransactionRepository_CreateTransaction(t *testing.T) {
	t.Run("should return success", func(t *testing.T) {
		inputTransaction := models.Transactions{
			TypeTransactionId: 1,
			CustomerId:        1,
			AgentId:           1,
			Address:           "Jl. H. Dahlan",
			ProvinceId:        32,
			RegencyId:         3204,
			DistrictId:        3204190,
			Amount:            1000000,
			StatusTransaction: 3,
			Rating:            5,
			RatingComment:     "Perfect Service ",
		}
		outputTransactionID := 1
		queryExpected:= "INSERT INTO `transaction`(`type_transaction_id`,`customers_id`,`agents_id`,`address`,`transaction_prov_id`,`transaction_regency_id`,`transaction_district_id`,`amount`,`status_transaction`,`rating`,`rating_comment`) VALUES (?,?,?,?,?,?,?,?,?,?,?)"

		mock.ExpectBegin()
		mock.ExpectExec(queryExpected).
			WillReturnResult(sqlmock.NewResult(1,1)).
			WillReturnError(nil)
		mock.ExpectCommit()

		actualTransactionID,err := repo.Save(&inputTransaction)
		assert.NoError(t,err)
		assert.Equal(t,actualTransactionID,outputTransactionID)

	})
}
