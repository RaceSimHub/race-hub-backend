package driver_test

import (
	"testing"

	"github.com/RaceSimHub/race-hub-backend/internal/service/driver"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	mockDb "github.com/RaceSimHub/race-hub-backend/internal/database/mock"
	"github.com/RaceSimHub/race-hub-backend/internal/database/sqlc"
)

type DriverSuite struct {
	suite.Suite
	driverService *driver.Driver
	mockDB        *mockDb.MockQuerier
}

func (suite *DriverSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	suite.mockDB = mockDb.NewMockQuerier(ctrl)
	suite.driverService = driver.NewDriver(suite.mockDB)
}

/*func (suite *DriverSuite) TestCreateDriver() {
	suite.mockDB.EXPECT().InsertDriver(gomock.Any(), sqlc.InsertDriverParams{
		Name:     "Michel Fiel",
		RaceName: "Bidu Fiel",
		Email:    "bidu@gmail.com",
		Phone:    "99887744",
	}).Return(int64(1), nil)

	id, err := suite.driverService.Create("Michel Fiel", "Bidu Fiel", "bidu@gmail.com", "99887744")
	suite.NoError(err)
	suite.Equal(int64(1), id)
}*/

/*func (suite *DriverSuite) TestUpdateDriver() {
	suite.mockDB.EXPECT().UpdateDriver(gomock.Any(), sqlc.UpdateDriverParams{
		ID:       1,
		Name:     "Michel Fiel",
		Email:    "bidu@gmail.com",
		Phone:    "99887744",
	}).Return(nil)

	err := suite.driverService.Update(1, "Michel Fiel", "Bidu Fiel", "bidu@gmail.com", "99887744")
	suite.NoError(err)
}*/

func (suite *DriverSuite) TestDeleteDriver() {
	suite.mockDB.EXPECT().DeleteDriver(gomock.Any(), int64(1)).Return(nil)

	err := suite.driverService.Delete(1)
	suite.NoError(err)
}

func (suite *DriverSuite) TestGetDriver() {
	suite.mockDB.EXPECT().GetDriver(gomock.Any(), int64(1)).Return(sqlc.GetDriverRow{
		ID: 1,
	}, nil)

	d, err := suite.driverService.GetByID(1)
	suite.NoError(err)
	suite.Equal(sqlc.GetDriverRow{
		ID: 1,
	}, d)
}

func (suite *DriverSuite) TestListDrivers() {
	suite.mockDB.EXPECT().SelectListDrivers(gomock.Any(), sqlc.SelectListDriversParams{
		Offset: int32(0),
		Limit:  int32(10),
	}).Return([]sqlc.SelectListDriversRow{
		{
			ID: 1,
		},
	}, nil)

	d, err := suite.driverService.GetList(0, 10)
	suite.NoError(err)
	suite.Equal([]sqlc.SelectListDriversRow{
		{
			ID: 1,
		},
	}, d)
}

func TestDriverSuite(t *testing.T) {
	suite.Run(t, new(DriverSuite))
}
