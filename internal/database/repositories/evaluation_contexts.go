package repositories

import (
	"encoding/json"
	"github.com/placer14/gof-server/internal/database"
	"github.com/placer14/gof-server/internal/handlers/payloads"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type EvaluationContextRepository struct {
	DB *gorm.DB
}

func (ecr *EvaluationContextRepository) CreateEvaluationContext(payload payloads.ContextEvaluation) (database.EvaluationContext, error) {
	db := ecr.DB
	var evaluationContext database.EvaluationContext

	jsonAttributes, err := json.Marshal(payload.Attributes)
	if err != nil {
		return evaluationContext, err
	}
	now := time.Now().UTC()
	query := `
	INSERT INTO evaluation_contexts (kind, key, attributes, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT (kind, key) DO UPDATE SET attributes = $3, updated_at = $4`

	result := db.Raw(query, payload.Kind, payload.Key, datatypes.JSON([]byte(jsonAttributes)), now, now).Scan(&evaluationContext)
	if result.Error != nil {
		return evaluationContext, result.Error
	}
	return evaluationContext, nil
}
