#!/bin/bash

export PGUSER=$POSTGRES_USER
export PGDATABASE=$POSTGRES_DB
export PGPASSWORD=$POSTGRES_PASSWORD

psql -c "DROP SCHEMA IF EXISTS $POSTGRES_SCHEMA CASCADE;"
psql -c "CREATE SCHEMA $POSTGRES_SCHEMA;"
psql -c "CREATE EXTENSION IF NOT EXISTS postgis;"

cd /tmp/topo-express/
append=''
src_epsg=2154
dst_epsg=4326

for f in *.7z; do
    if [[ $f =~ "RGAF09UTM20" ]]
    then
        src_epsg=5490
    elif [[ $f =~ "RGM04UTM38S" ]]
    then
        src_epsg=4471
    elif [[ $f =~ "RGR92UTM40S" ]]
    then
        src_epsg=2975
    elif [[ $f =~ "UTM22RGFG95" ]]
    then
        src_epsg=2972
    fi
    7z x $f
    xdir=`basename $f .7z`
    cd "$(find $xdir -iname "BATIMENT.shp" -printf '%h' -quit)"
    shp2pgsql -s $src_epsg:$dst_epsg -D $append BATIMENT.shp $POSTGRES_SCHEMA.building | psql
    cd /tmp/topo-express/
    rm -rf $xdir
    append='-a'
done

# To search by coordinates or square
psql -c "CREATE INDEX parcelle_geom_idx ON $POSTGRES_SCHEMA.building USING GIST (geom)"
# To search by id
psql -c "CREATE INDEX parcelle_idu_idx ON $POSTGRES_SCHEMA.building (id)"
psql -c "VACUUM ANALYZE"
