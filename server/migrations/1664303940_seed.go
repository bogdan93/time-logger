package migrations

import (
	"github.com/pocketbase/dbx"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		arniaProject := dbx.Params{
			"id":         "1",
			"name":       "Arnia",
			"odooId":     "40",
			"isReadonly": true,
		}

		dswissProject := dbx.Params{
			"id":         "2",
			"name":       "Dswiss",
			"odooId":     "192",
			"isReadonly": true,
		}
		campointProject := dbx.Params{
			"id":         "3",
			"name":       "Campoint",
			"odooId":     "102",
			"isReadonly": true,
		}
		db.Insert("projects", campointProject).Execute()
		db.Insert("projects", dswissProject).Execute()
		db.Insert("projects", arniaProject).Execute()

		// workhour_details
		developmentWKDetails := dbx.Params{
			"id":         "1",
			"details":    "Development",
			"short":      "🔧",
			"isThisWork": true,
			"isReadonly": true,
		}
		db.Insert("workhour_details", developmentWKDetails).Execute()

		leaveWKDetails := dbx.Params{
			"id":         "2",
			"details":    "Leave",
			"short":      "🏝️",
			"isThisWork": false,
			"isReadonly": true,
		}
		db.Insert("workhour_details", leaveWKDetails).Execute()

		nationalDayWKDetails := dbx.Params{
			"id":         "3",
			"details":    "National Day",
			"short":      "🇷🇴",
			"isThisWork": false,
			"isReadonly": true,
		}
		db.Insert("workhour_details", nationalDayWKDetails).Execute()

		// companies
		arniaCompany := dbx.Params{
			"id":         "1",
			"name":       "ARNIA SOFTWARE SRL",
			"nrRegCom":   "J40/8840/2006",
			"cui":        "RO18719824",
			"sediu":      "BUCUREȘTI, bd. Lascar Catargiu 47-53, Europe House, etajul 6, sector 1, 010665",
			"iban":       "RO 22 MIRO 0000 0701 9602 0101",
			"bankName":   "Procredit Bank – Sucursala Baneasa",
			"isReadonly": true,
		}
		db.Insert("companies", arniaCompany).Execute()

		return nil
	}, func(db dbx.Builder) error {
		// add down queries...

		return nil
	})
}
