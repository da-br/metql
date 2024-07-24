package database_test

import (
	"testing"

	"codeburg.com/da-br/metql/internal/database"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type DataBaseSuite struct {
	suite.Suite
}

func Test_BadMasterBaseFile_ReturnsError(t *testing.T) {
	db := database.NewDatabase("example/not_a.db")
	err := db.Start()

	require.Error(t, err)
}

func Test_FetchAll_ReturnCorrectCount(t *testing.T) {
	db := database.NewDatabase("../../example/masterbase.db")
	db.Start()

	result, err := db.Fetch(database.FetchQuery{})

	require.Nil(t, err, "err was not nil")

	res := *result
	require.Len(t, res, 3, "should have returned 3 rows")

	expectedModelNames := []string{"lower", "middle", "upper"}
	require.Contains(t, expectedModelNames, res[0].ModelName)
}
