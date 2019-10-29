#!/usr/bin/ruby

require 'docker'
require 'optparse'

env = {
  'MESSAGES' => 0,
  'SECONDS' => 0,
  'APPLICATIONS' => 'id1',
  'DATASYNC' => 'false',
  'DISCOVERY' => 'false'
}

OptionParser.new do |parser|
  parser.on('-m', '--messages=n', OptionParser::DecimalInteger) do |m|
    env['MESSAGES'] = m
  end

  parser.on('-s', '--seconds=n', OptionParser::DecimalInteger) do |s|
    env['SECONDS'] = s
  end

  parser.on('-g', '--discovery') do |d|
    env['DISCOVERY'] = 'true'
  end

  parser.on('-d', '--datasync') do |d|
    env['DATASYNC'] = 'true'
  end

  parser.on('-p', '--public-chat-id', String) do |p|
    env['PUBLIC_CHAT_ID'] = 'test200'
  end
  parser.on('-a', '--applications=n', OptionParser::DecimalInteger) do |app|
    applications = ''
    (1..app.to_i).each do |id|
      applications = "#{applications}id#{id}#{if id == app.to_i then '' else "," end}"
    end
    env['APPLICATIONS'] = applications
  end

end.parse!

env_string = env.map do |(k,v)|
  env_pairs = "#{k}=#{v}"
end.join(" ")

image = Docker::Image.build_from_dir('.')


container = Docker::Container.create({
  'Cmd' => './run.sh',
  'Image' => image.id})

puts "COMMAND: env #{env_string} ./run.sh"

container = container.run("env #{env_string} ./run.sh")

stats = nil
while true do
  pulled_stats = container.stats
  if pulled_stats['network'] && pulled_stats['network']['tx_bytes'] != 0
    stats = pulled_stats
    sleep 0.3
  else
    break
  end
end

if stats && stats['network']
  puts "SENT: #{stats['network']['tx_bytes']} bytes"
  puts "RECEIVED: #{stats['network']['rx_bytes']} bytes"
end
