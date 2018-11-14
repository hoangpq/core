#!/bin/sh

tendermint init
tendermint node --log_level debug --proxy_app tcp://service:26658
