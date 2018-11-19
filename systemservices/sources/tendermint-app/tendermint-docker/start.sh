#!/bin/sh

tendermint init
tendermint node --proxy_app tcp://service:26658
