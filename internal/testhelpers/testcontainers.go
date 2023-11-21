package testhelpers

/*
	usage:
	testDB := testhelpers.NewTestContainerDatabase(t)
	defer testDB.Close(t)
	println(testDB.ConnectionString(t))
*/

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type TestContainerDatabase struct {
	instance testcontainers.Container
}

const (
	rootUsername = "root"
	rootPassword = "password"
	databaseName = "mytestdb"
)

func newTestContainerDatabase(t *testing.T) *TestContainerDatabase {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	req := testcontainers.ContainerRequest{
		Image:        "mysql:8.2",
		ExposedPorts: []string{"3306/tcp", "33060/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": rootPassword,
			"MYSQL_DATABASE":      databaseName,
		},
		WaitingFor: wait.ForListeningPort("3306").WithStartupTimeout(15 * time.Second),
	}
	postgres, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)
	return &TestContainerDatabase{
		instance: postgres,
	}
}

func (db *TestContainerDatabase) Port(t *testing.T) int {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	p, err := db.instance.MappedPort(ctx, "3306")
	require.NoError(t, err)
	return p.Int()
}

func (db *TestContainerDatabase) ConnectionString(t *testing.T) string {
	hostName, _ := db.instance.Host(context.Background())
	return fmt.Sprintf("%s:%s@tcp(%v:%v)/%s", rootUsername, rootPassword, hostName, db.Port(t), databaseName)
}

func (db *TestContainerDatabase) Close(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	require.NoError(t, db.instance.Terminate(ctx))
}
