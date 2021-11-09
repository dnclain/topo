# API BDTOPOv3 

With this, you can build your own local IGN TOPO Database.
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
    * Use the viewer. 

7. Down the services

    `docker-compose down`
    OR
    `docker-compose down -v` to destroy the dataset.


## Dev mode 

Please reopen this project in a remote container from vscode.

## Credits
- Thanks to [Geosophy](https://www.geosophy.io) for its support.
- Thanks to Emmanuel S. (IGNFab) for his guidance.