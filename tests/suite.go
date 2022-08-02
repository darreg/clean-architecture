//go:build integration

package tests

import (
	"fmt"
	"strings"
	"time"

	"github.com/alrund/yp-1-project/internal/application/app"
	"github.com/alrund/yp-1-project/internal/domain/port"
	"github.com/alrund/yp-1-project/internal/infrastructure/builder"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
)

type IntegrationTestSuite struct {
	suite.Suite
	logger            port.Logger
	dockerComposeFile string
	identifier        string
	env               map[string]string
	app               *app.App
}

func NewIntegrationTestSuite(
	logger port.Logger,
	dockerComposeFile string,
	env map[string]string,
) suite.TestingSuite {
	ts := new(IntegrationTestSuite)
	ts.logger = logger
	ts.dockerComposeFile = dockerComposeFile
	ts.env = env
	return ts
}

func (s *IntegrationTestSuite) SetupSuite() {
	composeFilePaths := []string{s.dockerComposeFile}
	s.identifier = strings.ToLower(uuid.New().String())
	compose := testcontainers.NewLocalDockerCompose(composeFilePaths, s.identifier)
	err := compose.WithCommand([]string{"up", "-d"}).WithEnv(s.env).Invoke().Error
	if err != nil {
		s.logger.Fatal(fmt.Errorf("could not run compose file: %v - %v", composeFilePaths, err))
	}

	time.Sleep(time.Second * 10)

	s.app, err = builder.Builder(&app.Config{
		Debug:                 false,
		RunAddress:            "localhost:3000",
		DatabaseURI:           "postgres://dev:dev@localhost:" + s.env["POSTGRES_PORT"] + "/dev?sslmode=disable",
		AccrualSystemAddress:  "http://localhost:8080",
		AccrualSystemMethod:   "/api/orders/%s",
		CipherPass:            "J53RPX6",
		SessionCookieDuration: "24h",
		SessionCookieName:     "sessionID",
		MigrationDir:          "../migrations",
	}, s.logger)
	if err != nil {
		s.TearDownSuite()
		s.logger.Fatal(err)
	}
}

func (s *IntegrationTestSuite) TearDownSuite() {
	composeFilePaths := []string{s.dockerComposeFile}
	compose := testcontainers.NewLocalDockerCompose(composeFilePaths, s.identifier)
	err := compose.WithEnv(s.env).Down().Error
	if err != nil {
		s.logger.Fatal(fmt.Errorf("could not run compose file: %v - %v", composeFilePaths, err))
	}
}

func (s *IntegrationTestSuite) SetupTest()                            {}
func (s *IntegrationTestSuite) BeforeTest(suiteName, testName string) {}
func (s *IntegrationTestSuite) AfterTest(suiteName, testName string)  {}
func (s *IntegrationTestSuite) TearDownTest()                         {}
