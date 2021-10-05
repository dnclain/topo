package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	geojson "github.com/paulmach/go.geojson"
)

type building struct {
	id         sql.NullString
	nature     sql.NullString
	usage1     sql.NullString
	usage2     sql.NullString
	leger      sql.NullString
	etat       sql.NullString
	date_creat sql.NullString
	date_maj   sql.NullString
	date_app   sql.NullTime
	date_conf  sql.NullTime
	source     sql.NullString
	id_source  sql.NullString
	prec_plani sql.NullFloat64
	prec_alti  sql.NullFloat64
	nb_logts   sql.NullInt64
	nb_etages  sql.NullInt64
	mat_murs   sql.NullString
	mat_toits  sql.NullString
	hauteur    sql.NullFloat64
	z_min_sol  sql.NullFloat64
	z_min_toit sql.NullFloat64
	z_max_toit sql.NullFloat64
	z_max_sol  sql.NullFloat64
	origin_bat sql.NullString
	app_ff     sql.NullString
	geometry   *geojson.Geometry
}

func getGeoJSON(db *sql.DB, query string, args ...interface{}) (*geojson.FeatureCollection, error) {

	rows, err := db.Query(query, args...)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer rows.Close()

	fc := geojson.NewFeatureCollection()
	for rows.Next() {
		var a building
		if err := rows.Scan(
			&a.id,
			&a.nature,
			&a.usage1,
			&a.usage2,
			&a.leger,
			&a.etat,
			&a.date_creat,
			&a.date_maj,
			&a.date_app,
			&a.date_conf,
			&a.source,
			&a.id_source,
			&a.prec_plani,
			&a.prec_alti,
			&a.nb_logts,
			&a.nb_etages,
			&a.mat_murs,
			&a.mat_toits,
			&a.hauteur,
			&a.z_min_sol,
			&a.z_min_toit,
			&a.z_max_toit,
			&a.z_max_sol,
			&a.origin_bat,
			&a.app_ff,
			&a.geometry); err != nil {
			log.Println(err.Error())
			return nil, err
		}
		f := geojson.NewFeature(a.geometry)
		// 🥲 why no ternary operator OR 'if expression' in go, snif 🥲
		// id
		if a.id.Valid {
			_val, _ := a.id.Value()
			f.SetProperty("id", _val)
		} else {
			f.SetProperty("id", nil)
		}
		// nature
		if a.nature.Valid {
			_val, _ := a.nature.Value()
			f.SetProperty("nature", _val)
		} else {
			f.SetProperty("nature", nil)
		}
		// usage1
		if a.usage1.Valid {
			_val, _ := a.usage1.Value()
			f.SetProperty("usage1", _val)
		} else {
			f.SetProperty("usage1", nil)
		}
		// usage2
		if a.usage2.Valid {
			_val, _ := a.usage2.Value()
			f.SetProperty("usage2", _val)
		} else {
			f.SetProperty("usage2", nil)
		}
		// leger
		if a.leger.Valid {
			_val, _ := a.leger.Value()
			f.SetProperty("leger", _val)
		} else {
			f.SetProperty("leger", nil)
		}
		// etat
		if a.etat.Valid {
			_val, _ := a.etat.Value()
			f.SetProperty("etat", _val)
		} else {
			f.SetProperty("etat", nil)
		}
		// date_creat
		if a.date_creat.Valid {
			_val, _ := a.date_creat.Value()
			f.SetProperty("date_creat", _val)
		} else {
			f.SetProperty("date_creat", nil)
		}
		// date_maj
		if a.date_maj.Valid {
			_val, _ := a.date_maj.Value()
			f.SetProperty("date_maj", _val)
		} else {
			f.SetProperty("date_maj", nil)
		}
		// date_app
		if a.date_app.Valid {
			_val, _ := a.date_app.Value()
			f.SetProperty("date_app", _val)
		} else {
			f.SetProperty("date_app", nil)
		}
		// date_conf
		if a.date_conf.Valid {
			_val, _ := a.date_conf.Value()
			f.SetProperty("date_conf", _val)
		} else {
			f.SetProperty("date_conf", nil)
		}
		// source
		if a.source.Valid {
			_val, _ := a.source.Value()
			f.SetProperty("source", _val)
		} else {
			f.SetProperty("source", nil)
		}
		// id_source
		if a.id_source.Valid {
			_val, _ := a.id_source.Value()
			f.SetProperty("id_source", _val)
		} else {
			f.SetProperty("id_source", nil)
		}
		// prec_plani
		if a.prec_plani.Valid {
			_val, _ := a.prec_plani.Value()
			f.SetProperty("prec_plani", _val)
		} else {
			f.SetProperty("prec_plani", nil)
		}
		// prec_alti
		if a.prec_alti.Valid {
			_val, _ := a.prec_alti.Value()
			f.SetProperty("prec_alti", _val)
		} else {
			f.SetProperty("prec_alti", nil)
		}
		// nb_logts
		if a.nb_logts.Valid {
			_val, _ := a.nb_logts.Value()
			f.SetProperty("nb_logts", _val)
		} else {
			f.SetProperty("nb_logts", nil)
		}
		// nb_etages
		if a.nb_etages.Valid {
			_val, _ := a.nb_etages.Value()
			f.SetProperty("nb_etages", _val)
		} else {
			f.SetProperty("nb_etages", nil)
		}
		// mat_murs
		if a.mat_murs.Valid {
			_val, _ := a.mat_murs.Value()
			f.SetProperty("mat_murs", _val)
		} else {
			f.SetProperty("mat_murs", nil)
		}
		// mat_toits
		if a.mat_toits.Valid {
			_val, _ := a.mat_toits.Value()
			f.SetProperty("mat_toits", _val)
		} else {
			f.SetProperty("mat_toits", nil)
		}
		// hauteur
		if a.hauteur.Valid {
			_val, _ := a.hauteur.Value()
			f.SetProperty("hauteur", _val)
		} else {
			f.SetProperty("hauteur", nil)
		}
		// z_min_sol
		if a.z_min_sol.Valid {
			_val, _ := a.z_min_sol.Value()
			f.SetProperty("z_min_sol", _val)
		} else {
			f.SetProperty("z_min_sol", nil)
		}
		// z_min_toit
		if a.z_min_toit.Valid {
			_val, _ := a.z_min_toit.Value()
			f.SetProperty("z_min_toit", _val)
		} else {
			f.SetProperty("z_min_toit", nil)
		}
		// z_max_toit
		if a.z_max_toit.Valid {
			_val, _ := a.z_max_toit.Value()
			f.SetProperty("z_max_toit", _val)
		} else {
			f.SetProperty("z_max_toit", nil)
		}
		// z_max_sol
		if a.z_max_sol.Valid {
			_val, _ := a.z_max_sol.Value()
			f.SetProperty("z_max_sol", _val)
		} else {
			f.SetProperty("z_max_sol", nil)
		}
		// origin_bat
		if a.origin_bat.Valid {
			_val, _ := a.origin_bat.Value()
			f.SetProperty("origin_bat", _val)
		} else {
			f.SetProperty("origin_bat", nil)
		}
		// app_ff
		if a.app_ff.Valid {
			_val, _ := a.app_ff.Value()
			f.SetProperty("app_ff", _val)
		} else {
			f.SetProperty("app_ff", nil)
		}
		fc.AddFeature(f)
	}

	rerr := rows.Close()
	if rerr != nil {
		log.Println(rerr.Error())
		return nil, rerr
	}

	if err := rows.Err(); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return fc, nil
}

func getBuilding(db *sql.DB, key, value string) (*geojson.FeatureCollection, error) {
	return getGeoJSON(db, fmt.Sprintf("SELECT id, nature, usage1, usage2, leger, etat, date_creat, date_maj, date_app, date_conf, source, id_source, prec_plani, prec_alti, nb_logts, nb_etages, mat_murs, mat_toits, hauteur,  z_min_sol, z_min_toit, z_max_toit, z_max_sol, origin_bat, app_ff, ST_AsGeoJSON(geom) FROM %s.building WHERE %s=$1", os.Getenv("APP_DB_SCHEMA"), key), value)
}

func getBuildingIntersects(db *sql.DB, pos string) (*geojson.FeatureCollection, error) {
	return getGeoJSON(db, fmt.Sprintf("SELECT id, nature, usage1, usage2, leger, etat, date_creat, date_maj, date_app, date_conf, source, id_source, prec_plani, prec_alti, nb_logts, nb_etages, mat_murs, mat_toits, hauteur,  z_min_sol, z_min_toit, z_max_toit, z_max_sol, origin_bat, app_ff, ST_AsGeoJSON(geom) FROM %s.building WHERE ST_Intersects(geom,ST_SetSRID(ST_MakePoint(%s),4326))", os.Getenv("APP_DB_SCHEMA"), pos))
}

func getBuildingBbox(db *sql.DB, bbox string) (*geojson.FeatureCollection, error) {
	return getGeoJSON(db, fmt.Sprintf("SELECT id, nature, usage1, usage2, leger, etat, date_creat, date_maj, date_app, date_conf, source, id_source, prec_plani, prec_alti, nb_logts, nb_etages, mat_murs, mat_toits, hauteur,  z_min_sol, z_min_toit, z_max_toit, z_max_sol, origin_bat, app_ff, ST_AsGeoJSON(geom) FROM %s.building WHERE ST_Intersects(geom,ST_MakeEnvelope(%s,4326))", os.Getenv("APP_DB_SCHEMA"), bbox))
}
