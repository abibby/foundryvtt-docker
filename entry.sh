#!/usr/bin/env bash

export XDG_DATA_HOME=/config

mkdir -p "$XDG_DATA_HOME/FoundryVTT/Config"
node /entry.js > "$XDG_DATA_HOME/FoundryVTT/Config/options.json"

$@
