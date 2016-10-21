# Server APIs
class ServerApi < ApiBase
  get '/' do
    logger.info('GET /servers')
    logger.debug('Querying lab data service for servers...')
    # Fetch servers from lab data service
    servers = HTTParty.get('http://localhost:6001/api/servers')
    response = { 'servers' => servers }
    logger.debug("Successfully returning servers - RETURN: #{OK}")
    [OK, response.to_json]
  end

  post '/' do
    logger.info('POST /servers')
    logger.debug('Creating server with lab data service...')
    server = HTTParty.post('http://localhost:6001/api/servers',
      :headers => {
        'Content-type' => 'application/json'
      },
      :body => params.to_json
    )

    response = { 'server' => server }
    logger.debug("Successfully returning new server - RETURN: #{OK}")
    [OK, response.to_json]
  end

  get '/:id' do
    logger.info("POST /servers/#{params[:id]}")
    logger.debug('Querying lab data service for server...')
    # Fetch server from lab data service
    response = { 'server' => {} }
    logger.debug("Successfully returning server '#{params[:id]}' - RETURN: #{OK}")
    [OK, response.to_json]
  end

  put '/:id' do
    logger.info("PUT /servers/#{params[:id]}")
    logger.debug('Updating server with lab data service...')
    # PUT server to lab data service
    response = { 'server' => {} }
    logger.debug("Successfully returning server '#{params[:id]}' - RETURN: #{OK}")
    [OK, response.to_json]
  end

  delete '/:id' do
    logger.info("DELETE /servers/#{params[:id]}")
    logger.debug('Deleting server with lab data service...')
    # DELETE server with lab data service
    logger.debug("Successfully deleted server '#{params[:id]}' - RETURN: #{OK}")
    [OK, response.to_json]
  end
end
