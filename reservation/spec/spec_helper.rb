# Test env
require 'rack/test'

ENV['RESERVATION_ENV'] = 'test'
ENV['RESERVATION_LOG_LEVEL'] = 'off'

# require the dependencies
require File.join(File.dirname(__FILE__), '..', 'app', 'init')
include HttpStatusCodes
use Rack::PostBodyContentTypeParser

# Constants
JSON_CONTENT = 'application/json;charset=utf-8'
XML_CONTENT = 'application/xml;charset=utf-8'
JAVASCRIPT_CONTENT = 'application/javascript;charset=utf-8'
HTML_CONTENT = 'text/html;charset=utf-8'

RSpec.configure do |conf|
  conf.include Rack::Test::Methods
  conf.before(:each) do
    env 'SERVER_PORT', '9292'
  end
end
