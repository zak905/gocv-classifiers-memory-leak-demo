#!/bin/bash

for i in {1..100}
do
   curl -s localhost:8080/detect?url=https://statics.sportskeeda.com/editor/2018/09/f0e48-1536333568-800.jpg
done

echo "---------------- done"