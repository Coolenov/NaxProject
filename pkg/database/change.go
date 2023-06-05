package database

import (
	"github.com/Coolenov/FusionAPI-monorepo/initialize"
)

func ChangeLastRequestByLink(url string) error {
	_, err := initialize.DB.Exec("UPDATE scrapers SET last_request = CURRENT_TIMESTAMP WHERE link = ?", url)
	if err != nil {
		return err
	}
	return err
}
