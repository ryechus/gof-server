package test_storage

import (
	"context"
	"fmt"
	"github.com/placer14/gof-server/internal/handlers/payloads"
	"github.com/placer14/gof-server/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/datatypes"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog/log"
)

const (
	dbName     = "gof_server_test"
	dbUser     = "postgres"
	dbPassword = "dbpw"
)

func runPostgres(ctx context.Context) testcontainers.Container {
	postgresContainer, err := postgres.Run(ctx,
		"postgres:17-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithExposedPorts("5432/tcp"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		log.Printf("failed to start container: %s", err)
		return nil
	}

	mappedPort, err := postgresContainer.MappedPort(ctx, "5432/tcp")
	if err != nil {
		log.Fatal().Msgf("failed to get mapped port: %s", err)
	}

	postgresDSN := fmt.Sprintf(
		"host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Los_Angeles",
		dbUser, dbPassword, dbName, mappedPort.Port())
	err = os.Setenv("GOF_SERVER_POSTGRES_DSN", postgresDSN)
	if err != nil {
		log.Fatal().Msgf("failed to set environment variable %s", postgresDSN)
	}
	return postgresContainer
}

func TestCreateFlagKey(t *testing.T) {
	ctx := context.Background()
	container := runPostgres(ctx)
	defer func() {
		if err := testcontainers.TerminateContainer(container); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()

	dbStorage := storage.NewDBStorage()

	key := "test-flag"
	flagType := "bool"
	variations := []payloads.FlagVariation{
		{
			Name:  "enabled",
			Value: true,
		},
		{
			Name:  "disabled",
			Value: false,
		},
	}
	err := dbStorage.CreateFlag(key, flagType, variations)
	if err != nil {
		t.Errorf("failed to create flag: %s", err)
	}

	flagWithVariations, err := dbStorage.GetFlagWithVariations(key)
	if err != nil {
		t.Errorf("failed to get flag with variations: %s", err)
	}

	assert.Equal(t, key, flagWithVariations.Key)
	assert.Equal(t, flagType, flagWithVariations.FlagType)
	assert.Equal(t, true, flagWithVariations.DefaultEnabledVariation)
	assert.Equal(t, false, flagWithVariations.DefaultDisabledVariation)
}

func TestPutRule(t *testing.T) {
	ctx := context.Background()
	container := runPostgres(ctx)
	defer func() {
		if err := testcontainers.TerminateContainer(container); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()

	dbStorage := storage.NewDBStorage()

	key := "test-flag"
	flagType := "bool"
	variations := []payloads.FlagVariation{
		{
			Name:  "enabled",
			Value: true,
		},
		{
			Name:  "disabled",
			Value: false,
		},
	}
	err := dbStorage.CreateFlag(key, flagType, variations)
	if err != nil {
		t.Errorf("failed to create flag: %s", err)
	}
	flagWithVariations, err := dbStorage.GetFlagWithVariations(key)
	flagVariations, err := dbStorage.GetFlagVariations(key)
	if err != nil {
		t.Errorf("failed to get flag with variations: %s", err)
	}

	ruleContexts := []payloads.RuleContext{
		{
			ContextKind: "user",
			Attribute:   "email",
			Values:      []any{"example@email.com"},
		},
	}
	rule := payloads.PutRule{
		FlagUUID:      flagWithVariations.FlagUUID,
		Name:          "test",
		VariationUUID: flagVariations[0].UUID,
		Priority:      1,
		RuleContexts:  ruleContexts,
	}
	err = dbStorage.PutRule(rule)
	if err != nil {
		t.Errorf("failed to get flag with variations: %s", err)
	}

	flagWithVariations, _ = dbStorage.GetFlagWithVariations(key)

	assert.Equal(t, flagVariations[0].UUID, flagWithVariations.Rules[0].VariationUUID)
}

func TestPreventPutRuleUsingUnrelatedVariation(t *testing.T) {
	ctx := context.Background()
	container := runPostgres(ctx)
	defer func() {
		if err := testcontainers.TerminateContainer(container); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()

	dbStorage := storage.NewDBStorage()

	key := "test-flag"
	flagType := "bool"
	variations := []payloads.FlagVariation{
		{
			Name:  "enabled",
			Value: true,
		},
		{
			Name:  "disabled",
			Value: false,
		},
	}
	err := dbStorage.CreateFlag(key, flagType, variations)
	if err != nil {
		t.Errorf("failed to create flag: %s", err)
	}
	flagWithVariations, _ := dbStorage.GetFlagWithVariations(key)

	ruleContexts := []payloads.RuleContext{
		{
			ContextKind: "user",
			Attribute:   "email",
			Values:      []any{"example@email.com"},
		},
	}
	rule := payloads.PutRule{
		FlagUUID:      flagWithVariations.FlagUUID,
		Name:          "test",
		VariationUUID: datatypes.NewUUIDv4().String(),
		Priority:      1,
		RuleContexts:  ruleContexts,
	}
	err = dbStorage.PutRule(rule)

	assert.Error(t, err)
}

func TestEvaluateFlag(t *testing.T) {

	ctx := context.Background()
	container := runPostgres(ctx)
	defer func() {
		if err := testcontainers.TerminateContainer(container); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()

	dbStorage := storage.NewDBStorage()

	key := "test-flag"
	flagType := "bool"
	variations := []payloads.FlagVariation{
		{
			Name:  "enabled",
			Value: true,
		},
		{
			Name:  "disabled",
			Value: false,
		},
	}
	err := dbStorage.CreateFlag(key, flagType, variations)
	if err != nil {
		t.Errorf("failed to create flag: %s", err)
	}
	flagWithVariations, err := dbStorage.GetFlagWithVariations(key)
	flagVariations, err := dbStorage.GetFlagVariations(key)
	if err != nil {
		t.Errorf("failed to get flag with variations: %s", err)
	}

	ruleContexts := []payloads.RuleContext{
		{
			ContextKind: "user",
			Attribute:   "email",
			Values:      []any{"bad@email.com"},
			Operator:    "IN",
		},
	}
	rule := payloads.PutRule{
		FlagUUID:      flagWithVariations.FlagUUID,
		Name:          "test",
		VariationUUID: flagVariations[1].UUID,
		Priority:      1,
		RuleContexts:  ruleContexts,
	}
	err = dbStorage.PutRule(rule)
	if err != nil {
		t.Errorf("failed to get flag with variations: %s", err)
	}
	dbStorage.UpdateFlag(payloads.UpdateFlag{Key: key, Enabled: true})

	testCases := map[string]struct {
		evalCtx  map[string]any
		expected bool
	}{
		"test-123": {
			evalCtx: map[string]any{
				"email": "bad@email.com",
			},
			expected: false,
		},
		"test-456": {
			evalCtx: map[string]any{
				"email": "example@email.com",
			},
			expected: true,
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			evaluationContext := payloads.EvaluateFlag{
				Context: payloads.ContextEvaluation{
					Kind:       "user",
					Key:        name,
					Attributes: test.evalCtx,
				},
			}
			value, err := dbStorage.EvaluateFlag(key, evaluationContext)
			if err != nil {
				t.Errorf("failed to evaluate flag: %s", err)
			}

			assert.Equal(t, test.expected, value.(bool))
		})
	}
}
