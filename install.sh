#!/bin/bash
set -ex

go build -o bin/apps-metrics-plugin cmd/plugin/apps_metrics.go
cf uninstall-plugin AppsMetricsPlugin
cf install-plugin bin/apps-metrics-plugin -f
