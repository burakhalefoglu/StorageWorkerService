package test

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/model"
	"StorageWorkerService/internal/service/concrete"
	"StorageWorkerService/pkg/jsonParser/gojson"
	"StorageWorkerService/test/Mocks/repository"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ChurnBlocker_SuccessIsTrue(t *testing.T) {

	//Arrange
	var testChurnBlockerDal = new(repository.MockChurnBlockerMlResultDal)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.ChurnBlockerMlResultDal = testChurnBlockerDal

	churnBlocker := concrete.ChurnBlockerMlResultManagerConstructor()

	m:= model.ChurnBlockerMlResultModel{}
	testChurnBlockerDal.On("Add", &m).Return(nil)
	message, _ := (*churnBlocker.Parser).EncodeJson(&m)


	//Act
	success, err:= churnBlocker.AddChurnBlockerMlResultData(message)


	//Assert
	assert.Equal(t, true, success)
	assert.Equal(t, "", err)
}

func Test_ChurnBlocker_SuccessIsFalse(t *testing.T) {

	//Arrange
	var testChurnBlockerDal = new(repository.MockChurnBlockerMlResultDal)
	var json = gojson.GoJsonConstructor()

	IoC.JsonParser = json
	IoC.ChurnBlockerMlResultDal = testChurnBlockerDal

	churnBlocker := concrete.ChurnBlockerMlResultManagerConstructor()

	m:= model.ChurnBlockerMlResultModel{}
	testChurnBlockerDal.On("Add", &m).Return(errors.New("FakeError"))
	message, _ := (*churnBlocker.Parser).EncodeJson(&m)


	//Act
	success, err:= churnBlocker.AddChurnBlockerMlResultData(message)


	//Assert
	assert.Equal(t, false, success)
	assert.Equal(t, "FakeError", err)
}
