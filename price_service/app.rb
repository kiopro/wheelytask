require 'rubygems'
require 'bundler/setup'

# application
require 'sinatra/base'
require 'grpc'
require './calc'

this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(this_dir, 'proto')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'TripInformation_services_pb'

class App < Sinatra::Base
  get '/' do
    'Wheely Price Service v0.1!'
  end

  get '/healthcheck' do
    'I\'m alive!'
  end

  #####

  get '/get_price' do
    content_type :json

    request.body.rewind
    req_body = JSON.parse request.body.read

    if req_body['start_point'] and req_body['end_point']
      price = grpc_request(req_body['start_point'], req_body['end_point'])

      {price: price}.to_json
    else
      'Wrong payload'
    end
  end

	def grpc_request(start_point, end_point)
		service_address = ENV['DISTANCE_S']

    stub = Wheelypb::TripInformation::Stub.new(service_address, :this_channel_is_insecure)

    start_trip = Wheelypb::Point.new(lat: start_point["lat"], long: start_point["long"])
    end_trip = Wheelypb::Point.new(lat: end_point["lat"], long: end_point["long"])

    trip = Wheelypb::Trip.new(start: start_trip, end: end_trip)

    resp = stub.get_trip_information(trip)

    CaluclatePrice.new(resp.distance, resp.time).run
  end
end

App.run!


#docker run --rm -it -p 4567:4567 kiopro/price-service
