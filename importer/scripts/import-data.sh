#!/bin/bash

KEEP_FILES=NO

POSITIONAL=()
while [[ $# -gt 0 ]]; do
  key="$1"

  case $key in
    -k|--keep-files)
      KEEP_FILES=YES
      shift # past argument
      shift # past value
      ;;
    *)    # unknown option
      POSITIONAL+=("$1") # save it in an array for later
      shift # past argument
      ;;
  esac
done

set -- "${POSITIONAL[@]}" # restore positional parameters

echo "KEEPING FILES  = ${KEEP_FILES}"

export PGUSER=$POSTGRES_USER
export PGDATABASE=$POSTGRES_DB
export PGPASSWORD=$POSTGRES_PASSWORD
export PGHOST=$POSTGRES_HOST

echo "USER:$PGUSER"
echo "DB:$PGDATABASE "
echo -n "PASSWORD:"
echo "$PGPASSWORD" | wc -m
echo "HOST:$PGHOST"

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
    echo "🗜 Unzipping archive $f"
    7z x $f
    echo "🟢 Done"
    xdir=`basename $f .7z`
    cd "$(find $xdir -iname "BATIMENT.shp" -printf '%h' -quit)"
    echo "💾 Importing to database with conversion $src_epsg => $dst_epsg"
    shp2pgsql -s $src_epsg:$dst_epsg -D $append BATIMENT.shp $POSTGRES_SCHEMA.building | psql
    echo "🟢 Done"

    echo "🧽 Cleaning $xdir and $f"
    cd /tmp/topo-express/
    # delete the extracted folder
    echo "- Deleting temporary folder"
    rm -rf $xdir
    # delete the archive
    if [[$KEEP_FILES = "NO"]]
    then
        echo "- Deleting the archive"
        rm -rf $f
    else
        echo "- Keeping the archive"
    fi
    echo "🟢 Done"
    append='-a'
done

echo "🗂 Indexing geometries..."
# To search by coordinates or square
psql -c "CREATE INDEX parcelle_geom_idx ON $POSTGRES_SCHEMA.building USING GIST (geom)"
echo "🟢 Done"
echo "🗂 Indexing ids..."
# To search by id
psql -c "CREATE INDEX parcelle_idu_idx ON $POSTGRES_SCHEMA.building (id)"
echo "🟢 Done"
echo "Vacuum..."
psql -c "VACUUM ANALYZE"
echo "🟢 Done"
