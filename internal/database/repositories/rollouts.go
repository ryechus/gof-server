package repositories

import (
	"encoding/json"
	"github.com/placer14/gof-server/internal/database"
	"github.com/placer14/gof-server/internal/handlers/payloads"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type RolloutRepository struct {
	DB *gorm.DB
}

func (rr *RolloutRepository) GetRollout(flagKeyUUID datatypes.UUID) (database.Rollout, error) {
	db := rr.DB
	var rollout database.Rollout

	query := `SELECT * FROM rollouts WHERE flag_key_uuid = ?`
	result := db.Raw(query, flagKeyUUID.String()).Scan(&rollout)
	if result.Error != nil {
		return rollout, result.Error
	}

	return rollout, nil
}

func (rr *RolloutRepository) SetupRollout(payload payloads.PutRolloutRule) (database.Rollout, error) {
	db := rr.DB
	var rollout database.Rollout

	jsonAttributes, err := json.Marshal(payload.ContextConfig)
	if err != nil {
		return rollout, err
	}
	query := `
	INSERT INTO rollouts (uuid, flag_key_uuid, rollout_type, config) 
	VALUES ($1, $2, $3, $4) ON CONFLICT (uuid) DO UPDATE SET config = $4`
	result := db.Raw(query, payload.UUID, payload.FlagUUID, payload.RolloutType, datatypes.JSON([]byte(jsonAttributes))).Scan(&rollout)
	if result.Error != nil {
		return rollout, result.Error
	}
	return rollout, nil
}

func (rr *RolloutRepository) DeleteRollout(rolloutUUID string) error {
	db := rr.DB
	query := `DELETE FROM rollouts WHERE uuid = ?`
	row := struct{}{}
	result := db.Raw(query, rolloutUUID).Scan(&row)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
