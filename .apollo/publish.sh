#! /bin/bash

SUBGRAPH_NAME=user-subgraph
SUPERGRAPH_NAME=cuffney-supergraph-vokiem
SUPERGRAPH_VARIANT=main

read -p "Are you sure you want to publish the schema? y/n" -n 1 -r
echo    # (optional) move to a new line
if [[ $REPLY =~ ^[Yy]$ ]]
then
    rover subgraph publish $SUPERGRAPH_NAME@$SUPERGRAPH_VARIANT --name $SUBGRAPH_NAME --schema ./graph/schema/schema.graphql
fi