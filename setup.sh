#!/bin/sh

# TODO(simar7): Port this over to use docker compose

gnatsd --debug &
liftbridge --raft-bootstrap-seed --level=debug &
liftbridge --data-dir /tmp/liftbridge/server-2 --port=9293 --level=debug &
liftbridge --data-dir /tmp/liftbridge/server-3 --port=9294 --level=debug