# API BDTOPOv3 

With this, you can build your own local IGN TOPO Database. Currently, it provides ONLY `building` information.
The API provides also a viewer to explore it. 

This api is eligible to a merge with [`parcellaire express`](https://github.com/esgn/api-parcellaire-express) to build only one product, probably in the future. 

> Inspired from [PARCELLAIRE EXPRESS](https://github.com/esgn/api-parcellaire-express)

## Instructions

### Dev environement

1. Copy .env.example to .env. Change to values that suit your needs.
2. Generate `docker-compose.yml` file

    `docker-compose -f docker-compose.common.yml -f docker-compose.dev.yml config > docker-compose.yml`

3. Tune you postgres with [pgtunes](https://pgtune.leopard.in.ua/#/) and update `docker-composer.yml` with the new values

4. Build images

    `docker-compose build`

5. Eventually push the images to a registry for the stack environment.

    ```bash
    docker login <registry>
    docker-compose push
    ```

6. Turn on the services

    `docker-compose up` OR `docker-compose up -d` to start in the background.

7. Import the dataset
    * Download
        `docker-compose run -ti topo-importer python3 /tmp/download-dataset.py`
    * Import
        `docker-compose run -ti topo-importer bash /tmp/import-data.sh`

8. Use the api
    * Use the viewer with the url defined by VIEWER_URL
    * Use the routes defined below

9.  Turn off the services

    `docker-compose down`
    OR
    `docker-compose down -v` to destroy the dataset (only if the data are in volumes)

### Stack/traefik (=production) environment

1. Copy `.env.example` to `.env`. Change to values that suit the production environment.
2. Ensure `docker-compose` is installed : [Installation guide](https://docs.docker.com/compose/install/)
3. Generate `docker-stack.yml` file

    `docker-compose -f docker-compose.common.yml -f docker-compose.stack.yml config > docker-stack.yml`

4. Tune you postgres with [pgtunes](https://pgtune.leopard.in.ua/#/) and update `docker-stack.yml` with the new values

5. Turn on the services

    ```bash
    # Deploy
    docker stack deploy topo -c docker-stack.yml --with-registry-auth
    # Check 
    docker stack ps --no-trunc
    # Service reference 
    docker service ls
    ```    

6. Import the dataset
    * Download
  
      `docker exec -ti topo_topo-importer.XXXXX python3 /tmp/download-dataset.py`

    * Import

      `docker exec -ti topo_topo-importer.XXXXX /bin/bash /tmp/import-data.sh`


7. Use the api
    * Use the viewer with the url defined by VIEWER_URL
    * Use the routes defined below

8.  Turn off the services

    `docker stack rm topo`


## Routes

* **GET** `/building/{id}` : Retrieve a building definition by its id.
  * Exemple : http://localhost:8010/building/01053000BE0095
* **GET** `/building?pos={pos}` *ou* `/building?lon={lon}&lat={lat}` : Find the buildings that intersect with a geographic coordinate (WGS84)
  * Exemple : http://localhost:8010/building?pos=5.2709,44.6247
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