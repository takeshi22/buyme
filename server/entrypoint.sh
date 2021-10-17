#!/bin/bash
set -e

rm -f /buyme/tmp/pids/server.pid

exec "$@"