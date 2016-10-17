# Require dependencies
require File.join(File.dirname(__FILE__), 'app', 'init')

use Rack::MethodOverrideWithParams
set :method_override, :true

use Rack::PostBodyContentTypeParser

map '/api/1/users' do
  run UserApi
end

map '/api/1/servers' do
  run ServerApi
end

map '/api/1/reservations' do
  run ReservationApi
end
