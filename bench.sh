#!/bin/bash

function run10()
{
    for t in $(seq 1 10); do
         $@ > /dev/null
    done
}

echo Go
/usr/bin/time -f '%E %M' ./aggregator-go/main ./generator/data.csv > /tmp/go.txt
echo
echo Rust
/usr/bin/time -f '%E %M' ./aggregator-rust/target/release/aggregator ./generator/data.csv > /tmp/rust.txt
echo
echo Ruby
/usr/bin/time -f '%E %M' ruby ./aggregator-ruby/main.rb ./generator/data.csv > /tmp/ruby.txt

diff -bu /tmp/go.txt /tmp/rust.txt
diff -bu /tmp/go.txt /tmp/ruby.txt
