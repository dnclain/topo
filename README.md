# API BDTOPOv3 

With this, you can build your own local IGN TOPO Database. Currently, it provides ONLY `building` information.
The API provides also a viewer to explore it. 

This api is eligible to a merge with [`parcellaire express`](https://github.com/esgn/api-parcellaire-express) to build only one product, probably in the future. 

> Inspired from [PARCELLAIRE EXPRESS](https://github.com/esgn/api-parcellaire-express)

## Instructions

1. Copy .env.example to .env. Change to values that suit your needs.
2. Tune you postgres with [pgtunes](https://pgtune.leopard.in.ua/#/) and update `docker-composer.yml` with the new values
3. Build images

    `docker-compose build`

4. Turn on the services

    `docker-compose up` OR `docker-compose up -d` to start in the background.

5. Import the dataset
    * Download
        `docker-compose run topo-importer python3 /tmp/download-dataset.py`
    * Import
        `docker-compose run topo-importer bash /tmp/import-data.sh`

6. Use the api
    * Use the viewer with the url defined by VIEWER_URL
    * Use the routes defined below

7. Turn off the services

    `docker-compose down`
    OR
    `docker-compose down -v` to destroy the dataset.

## Routes

* **GET** `/building/{id}` : Retrieve a building definition by its id.
  * Exemple : http://localhost:8010/parcelle/01053000BE0095
* **GET** `/building?pos={pos}` *ou* `/building?lon={lon}&lat={lat}` : Find the buildings that intersect with a geographic coordinate (WGS84)
  * Exemple : http://localhost:8010/parcelle?pos=5.2709,44.6247
* **GET** `/building?bbox={bbox}` *ou* `/building?lon_min={lon}&lat_min={lat}&lon_max={lon}&lat_max={lat}` : Find the buildings that intersect with a bounding box expressed in geographic coordinates (WGS84)
  * Exemple : http://localhost:8010/building?bbox=5.2135,44.5719,5.2709,44.6247

### Formats

#### Parameters

* `{id}` : Unique building identifier (ex: `BATIMENTXXXXX`)
* `{lon}` : Longitude (real number from -180 to 180)
* `{lat}` : Latitude (real number from -90 to 90)
* `{pos}` : Position composed from 2 coordinates (`lon,lat`)
* `{bbox}` : Bounding box composed from 4 coordonnates (`lon_min,lat_min,lon_max,lat_max`)

#### Results

Results are provided in [GeoJSON](https://geojson.org/) format.

## Dev mode 

`TODO`: add a remote container for vscode.

## Credits
- Thanks to [Geosophy](https://www.geosophy.io) for its support.
- Thanks to Emmanuel S. (IGNFab) for his guidance.