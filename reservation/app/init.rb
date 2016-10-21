require 'rubygems'
require 'bundler/setup'
require 'logging'

#
# Set RESERVATION_ENV first and then load bundle
ENV['RACK_ENV'] = RESERVATION_ENV = ENV['RESERVATION_ENV'] || 'production'
Bundler.require(:default, RESERVATION_ENV)

APP_LOG_DIR = File.join(File.dirname(__FILE__), '..', 'log')

#
# Disable verification of self-signed certs in production
if Sinatra::Base.environment != :production
  HTTParty::Basement.default_options.update(verify: false)
end

#
# Build application namespace
# ==============================================================================
#
# Require all config files first
lib_files = {
  config: Dir[File.join(File.dirname(__FILE__), '..', 'config', '*.rb')],
  app: Dir[File.join(File.dirname(__FILE__), '..', 'app', '*.rb')],
  models: Dir[File.join(File.dirname(__FILE__), '..', 'app', 'lib', 'models', '*.rb')],
}

#
# Order of execution dictates config files load first
lib_files[:config].each{|lib| require lib}

#
# Require individual files where order of execution matters
require File.join(File.dirname(__FILE__), '..', 'app', 'lib', 'util', 'http_status_codes')
require File.join(File.dirname(__FILE__), '..', 'app', 'api_base')

#
# Bring in all modules, apps, validators, models and representers
[
  lib_files[:app],
  lib_files[:models]
].each do |lib_list|
  lib_list.each{|lib| require lib}
end
