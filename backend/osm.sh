# TODO: Добавить обработку добавления многих городов

#!/bin/bash
source .env

osmconvert --all-to-nodes $OSM_DIR/RU-SPE.pbf -o=$OSM_DIR/RU-SPE.o5m

echo "Converted to o5m"

osmfilter $OSM_DIR/RU-SPE.o5m --keep="building AND addr*" --drop-author --drop-version -o=$OSM_DIR/ADDR-RU-SPE.o5m

echo "Filtered required values"

osmconvert $OSM_DIR/ADDR-RU-SPE.o5m -o=$OSM_DIR/ADDR-RU-SPE.csv --csv-headline --csv-separator=";" --csv="addr:street addr:housenumber @lat @lon"

echo "Converted to CSV"

rm $OSM_DIR/ADDR-RU-SPE.o5m $OSM_DIR/RU-SPE.pbf $OSM_DIR/RU-SPE.o5m

echo "Removed other files"