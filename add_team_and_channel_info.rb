#!/usr/bin/env ruby
require "json"

team = ARGV[0]
id = ARGV[1]
channel = ARGV[2]

STDIN.each{|line|
  item = JSON.parse(line);
  item["team"] = team
  item["id"] = id
  item["channel"] = channel
  puts item.to_json
}
