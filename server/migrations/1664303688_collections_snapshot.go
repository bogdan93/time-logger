package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

// Auto generated migration with the most recent collections configuration.
func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `[
			{
				"id": "systemprofiles0",
				"created": "2022-09-23 17:24:07.457",
				"updated": "2022-09-23 17:25:08.668",
				"name": "profiles",
				"system": true,
				"schema": [
					{
						"system": true,
						"id": "pbfielduser",
						"name": "userId",
						"type": "user",
						"required": true,
						"unique": true,
						"options": {
							"maxSelect": 1,
							"cascadeDelete": true
						}
					},
					{
						"system": false,
						"id": "pbfieldname",
						"name": "name",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "pbfieldavatar",
						"name": "avatar",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [
								"image/jpg",
								"image/jpeg",
								"image/png",
								"image/svg+xml",
								"image/gif"
							],
							"thumbs": null
						}
					}
				],
				"listRule": "userId = @request.user.id",
				"viewRule": "userId = @request.user.id",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null
			},
			{
				"id": "dcmes2w11ekjy6z",
				"created": "2022-09-23 17:25:08.668",
				"updated": "2022-09-23 17:25:08.668",
				"name": "projects",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "s7dpuotl",
						"name": "name",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": 1,
							"max": 64,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "xzd2knyx",
						"name": "odooid",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": 1,
							"max": 4,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "alao4azd",
						"name": "userId",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "t5na7s5x",
						"name": "isReadonly",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					}
				],
				"listRule": "userId = '' || userId = @request.user.id",
				"viewRule": "userId = '' || userId = @request.user.id",
				"createRule": "@request.user.id = userId \u0026\u0026 isReadonly = false",
				"updateRule": "@request.user.id = userId \u0026\u0026 isReadonly = false",
				"deleteRule": "@request.user.id = userId \u0026\u0026 isReadonly = false"
			},
			{
				"id": "fmke95gci1s7air",
				"created": "2022-09-23 17:25:08.668",
				"updated": "2022-09-23 17:25:08.668",
				"name": "workhour_details",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "2z1478r9",
						"name": "details",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": 5,
							"max": 124,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "jy25rdnj",
						"name": "short",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": 1,
							"max": 124,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "xhrytbae",
						"name": "isThisWork",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "qvguofkh",
						"name": "isReadonly",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "ychblliu",
						"name": "userId",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"listRule": "userId = '' || userId = @request.user.id",
				"viewRule": "userId = '' || userId = @request.user.id",
				"createRule": "@request.user.id = userId \u0026\u0026 isReadonly = false",
				"updateRule": "@request.user.id = userId \u0026\u0026 isReadonly = false",
				"deleteRule": "@request.user.id = userId \u0026\u0026 isReadonly = false"
			},
			{
				"id": "nw5axwb5dveuoj1",
				"created": "2022-09-23 17:25:08.668",
				"updated": "2022-09-23 17:25:08.668",
				"name": "workday_logs",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "ef1tyxvg",
						"name": "day",
						"type": "number",
						"required": true,
						"unique": false,
						"options": {
							"min": 1,
							"max": 32
						}
					},
					{
						"system": false,
						"id": "xxh4wgm4",
						"name": "month",
						"type": "number",
						"required": true,
						"unique": false,
						"options": {
							"min": 1,
							"max": 13
						}
					},
					{
						"system": false,
						"id": "jm4xnvsz",
						"name": "year",
						"type": "number",
						"required": true,
						"unique": false,
						"options": {
							"min": 1,
							"max": 3000
						}
					},
					{
						"system": false,
						"id": "eck9z21x",
						"name": "workedHours",
						"type": "json",
						"required": true,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "jq4xqdrz",
						"name": "userId",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"listRule": "@request.user.id = userId",
				"viewRule": "@request.user.id = userId",
				"createRule": "@request.user.id = userId",
				"updateRule": "@request.user.id = userId",
				"deleteRule": "@request.user.id = userId"
			},
			{
				"id": "kzn01p58622qfv7",
				"created": "2022-09-23 17:25:08.668",
				"updated": "2022-09-23 17:25:08.668",
				"name": "signatures",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "x4zcmmz7",
						"name": "name",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": 1,
							"max": 40,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "lbdcmiwb",
						"name": "field",
						"type": "file",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 15242880,
							"mimeTypes": [
								"image/jpg",
								"image/jpeg",
								"image/png"
							],
							"thumbs": [
								"100x100f"
							]
						}
					},
					{
						"system": false,
						"id": "oimfnqko",
						"name": "userId",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"listRule": "@request.user.id = userId",
				"viewRule": "@request.user.id = userId",
				"createRule": "@request.user.id = userId",
				"updateRule": "@request.user.id = userId",
				"deleteRule": "@request.user.id = userId"
			},
			{
				"id": "zrzweq0ios2p2bz",
				"created": "2022-09-23 17:25:08.669",
				"updated": "2022-09-23 17:25:08.669",
				"name": "companies",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "wjcijxda",
						"name": "name",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": 1,
							"max": 124,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "8judklkp",
						"name": "nrRegCom",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": 1,
							"max": 124,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "c2rihcuv",
						"name": "cui",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": 1,
							"max": 20,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "th7fcjl0",
						"name": "sediu",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": 1,
							"max": 512,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "2rajxaaw",
						"name": "iban",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": 1,
							"max": 124,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "hggf1w8a",
						"name": "bankName",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": 1,
							"max": 124,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "mbgh2ubx",
						"name": "capitalSocial",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": 1,
							"max": 124,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "zv10laof",
						"name": "delegat_nume",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": 1,
							"max": 124,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "xi8itzff",
						"name": "delegat_cnp",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": 1,
							"max": 124,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "ad5eprj7",
						"name": "delegat_ciSerie",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": 1,
							"max": 124,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "5hhfowgu",
						"name": "delegat_ciNr",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": 1,
							"max": 124,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "sqy6wgey",
						"name": "delegat_eliberatDe",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": 1,
							"max": 124,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "wbdoah6b",
						"name": "judet",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": 1,
							"max": 50,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "6mbu72lz",
						"name": "isReadonly",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "c5fm4bed",
						"name": "userId",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"listRule": "userId = '' || userId = @request.user.id",
				"viewRule": "userId = '' || userId = @request.user.id",
				"createRule": "@request.user.id = userId \u0026\u0026 isReadonly = false",
				"updateRule": "@request.user.id = userId \u0026\u0026 isReadonly = false",
				"deleteRule": "@request.user.id = userId \u0026\u0026 isReadonly = false"
			},
			{
				"id": "8jc9ewp67eh6r2t",
				"created": "2022-09-23 17:25:08.669",
				"updated": "2022-09-23 17:25:08.669",
				"name": "companies_contract_details",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "d9dkcpyz",
						"name": "description",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": 1,
							"max": 512,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "cpc5eto8",
						"name": "userId",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"listRule": "@request.user.id = userId",
				"viewRule": "@request.user.id = userId",
				"createRule": "@request.user.id = userId",
				"updateRule": "@request.user.id = userId",
				"deleteRule": "@request.user.id = userId"
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		// no revert since the configuration on the environment, on which
		// the migration was executed, could have changed via the UI/API
		return nil
	})
}
