#!/usr/bin/env ruby
require "json"
require "date"

list_filename = ARGV[0]

id2username = {}
open(list_filename).each{|line|
  m = JSON.parse(line)
  id2username[m["id"]] = m["name"]
}

STDIN.each{|line|
  item = JSON.parse(line);
  item["date"] = DateTime.strptime(item["ts"].to_s, "%s").strftime("%Y%m%dT%H%M%S%Z");
  item["username"] = id2username[item["user"]]
  puts item.to_json
}
