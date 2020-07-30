package system

import (
	"time"

	"github.com/seknox/trasa/models"
	logger "github.com/sirupsen/logrus"
)

func (s SystemStore) SetGlobalSetting(setting models.GlobalSettings) error {
	_, err := s.DB.Exec(`INSERT into global_settings (id, org_id, status, type, value, updated_by, updated_on) values($1, $2, $3, $4, $5, $6, $7);`,
		setting.SettingID, setting.OrgID, setting.Status, setting.SettingType, setting.SettingValue, setting.UpdatedBy, setting.UpdatedOn)
	return err
}

func (s SystemStore) GetGlobalSetting(orgID, settingType string) (models.GlobalSettings, error) {
	//logger.Trace(orgID, settingType)
	var setting models.GlobalSettings
	err := s.DB.QueryRow(`
		SELECT id, org_id,status, type, value, updated_by, updated_on FROM global_settings WHERE org_id = $1 AND type=$2;`,
		orgID, settingType).
		Scan(&setting.SettingID, &setting.OrgID, &setting.Status, &setting.SettingType, &setting.SettingValue, &setting.UpdatedBy, &setting.UpdatedOn)

	return setting, err
}

func (s SystemStore) UpdateGlobalSetting(setting models.GlobalSettings) error {

	result, err := s.DB.Exec(`UPDATE global_settings SET status = $3, value = $4, updated_by = $5, updated_on =$6  WHERE org_id = $1 AND type = $2;`,
		setting.OrgID, setting.SettingType, setting.Status, setting.SettingValue, setting.UpdatedBy, setting.UpdatedOn)

	//TODO Possible nil pointer dereference, value of result could be nil
	v, _ := result.RowsAffected()

	if err != nil || v == 0 {
		err = s.SetGlobalSetting(setting)
		return err
	}

	return err
}

func (s SystemStore) CreateSecurityRule(rule models.SecurityRule) error {

	_, err := s.DB.Exec(`INSERT into security_rules (id, org_id, name, const_name,description, scope,condition,status, source, action,created_by, created_at, last_modified )
						 values($1, $2, $3, $4, $5,$6,$7,$8, $9, $10, $11, $12, $13);`, rule.RuleID, rule.OrgID, rule.Name, rule.ConstName, rule.Description, rule.Scope, rule.Condition, rule.Status, rule.Source, rule.Action, rule.CreatedBy, rule.CreatedAt, rule.LastModified)
	return err
}

func (s SystemStore) GetSecurityRuleByName(orgID, constName string) (models.SecurityRule, error) {
	var rule models.SecurityRule
	err := s.DB.QueryRow(`SELECT id, org_id, name, const_name,description, scope,condition,status, source, action,created_by, created_at, last_modified FROM security_rules where org_id=$1 AND const_name=$2 `, orgID, constName).Scan(&rule.RuleID, &rule.OrgID, &rule.Name, &rule.ConstName, &rule.Description, &rule.Scope, &rule.Condition, &rule.Status, &rule.Source, &rule.Action, &rule.CreatedBy, &rule.CreatedAt, &rule.LastModified)
	if err != nil {
		return rule, err
	}

	return rule, nil
}

func (s SystemStore) getSecurityRules(orgID string) ([]models.SecurityRule, error) {

	var rule models.SecurityRule
	var rules []models.SecurityRule = make([]models.SecurityRule, 0)

	rows, err := s.DB.Query(`SELECT id, org_id, name, const_name,description, scope,condition,status, source, action,created_by, created_at, last_modified FROM security_rules where org_id=$1`, orgID)
	if err != nil {
		//fmt.Printf("after rows: %v\n", err)
		return rules, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&rule.RuleID, &rule.OrgID, &rule.Name, &rule.ConstName, &rule.Description, &rule.Scope, &rule.Condition, &rule.Status, &rule.Source, &rule.Action, &rule.CreatedBy, &rule.CreatedAt, &rule.LastModified)
		if err != nil {
			logger.Debug(err)
		}
		rules = append(rules, rule)
	}

	return rules, nil

}

func (s SystemStore) updateSecurityRule(orgID, statusstr, ruleID string) error {
	status := true
	if statusstr == "disabled" {
		status = false
	}
	_, err := s.DB.Exec(`UPDATE security_rules SET status = $1, last_modified = $2 WHERE org_id = $3 AND id=$4;`,
		status, time.Now().Unix(), orgID, ruleID)

	if err != nil {
		return err
	}
	return nil
}

func (s SystemStore) storeBackupMeta(backup models.Backup) error {

	_, err := s.DB.Exec(`INSERT into backups (backup_id, org_id, backup_name, backup_type,created_at )
						 values($1, $2, $3, $4, $5);`, backup.BackupID, backup.OrgID, backup.BackupName, backup.BackupType, backup.CreatedAt)

	if err != nil {
		return err
	}
	return err
}

func (s SystemStore) getBackupMeta(backup models.Backup) (models.Backup, error) {

	err := s.DB.QueryRow(`SELECT  backup_name,backup_type, created_at  FROM backups WHERE backup_id=$1 AND org_id=$2`, backup.BackupID, backup.OrgID).Scan(&backup.BackupName, &backup.BackupType, &backup.CreatedAt)

	if err != nil {
		return backup, err
	}
	return backup, err
}

func (s SystemStore) getBackupMetas(orgID string) ([]models.Backup, error) {

	var backup models.Backup
	var backups = make([]models.Backup, 0)

	rows, err := s.DB.Query(`SELECT  backup_id, org_id, backup_name,backup_type, created_at  FROM backups WHERE org_id=$1 ORDER BY created_at DESC`, orgID)
	if err != nil {
		logger.Debug(err)
		return backups, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&backup.BackupID, &backup.OrgID, &backup.BackupName, &backup.BackupType, &backup.CreatedAt)
		if err != nil {
			logger.Debug(err)
		}
		backups = append(backups, backup)
	}

	return backups, nil

}
