require 'sinatra'
require 'httparty'
require 'logging'

# Base API Class
class ApiBase < Sinatra::Base
  include HttpStatusCodes

  # make the default content type JSON.
  # Actions that require something else can simply change it using the same call
  before { content_type 'application/json', :charset => 'utf-8' }
  before { headers 'Access-Control-Allow-Origin' => '*' }
  before { headers 'Access-Control-Allow-Methods' => 'GET,POST,PUT,DELETE,OPTIONS' }
  before { headers 'Access-Control-Allow-Headers' => 'Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With' }

  set :raise_errors, false
  set :show_exceptions, false
  #
  # Setup global logging logic
  include Logging.globally
  layout = Logging.layouts.pattern(:pattern => "%d %-5l [%c]: %m\n")
  #
  # Define which appenders to send log data
  Logging.logger.root.add_appenders(Logging.appenders.file("#{::APP_LOG_DIR}/output.log", :layout => layout))
  Logging.logger.root.add_appenders(Logging.appenders.stdout(:layout => layout)) if STDOUT.tty?

  Logging.logger.root.level = ENV['RESERVATION_LOG_LEVEL'] ||= 'info'
  #
  # Generic before filter for all API requests
  before do
    halt 200 if request.request_method == 'OPTIONS'
  end

  def body_to_json(request)
    if !request.content_length.nil? && request.content_length != '0'
      return MultiJson.decode(request.body.read)
    else
      logger.error("Unable to read request body due to content length - RETURN: #{BAD_REQUEST}")
      halt [BAD_REQUEST, Message.new('Invalid Request').to_json]
    end
  end

  def stringify_params
    string = '?'
    params.each do |key, value|
      string += "#{key}=#{value}&"
    end
    string.chop!
  end

  error do
    handle_error(env['sinatra.error'])
  end

  def handle_error(error)
    logger.error("#{error} - RETURN: #{ERROR}")
    [ERROR, Message.new(error.to_s).to_json]
  end
end
