# User APIs
class UserApi < ApiBase
  get '/' do
    logger.info('GET /users')
    logger.debug('Querying lab data service for users...')
    users = HTTParty.get('http://localhost:6001/api/users')
    response = { 'users' => users }
    logger.debug("Successfully returning users - RETURN: #{OK}")
    [OK, response.to_json]
  end

  post '/' do
    logger.info('POST /users')
    logger.debug('Creating user with lab data service...')
    user = HTTParty.post('http://localhost:6001/api/users',
      :headers => {
        'Content-type' => 'application/json'
      },
      :body => params.to_json
    )
    response = { 'user' => user }
    logger.debug("Successfully returning new user - RETURN: #{OK}")
    [OK, response.to_json]
  end

  get '/:id' do
    logger.info("POST /users/#{params[:id]}")
    logger.debug('Querying lab data service for user...')
    # Fetch user from lab data service
    response = { 'user' => {} }
    logger.debug("Successfully returning user '#{params[:id]}' - RETURN: #{OK}")
    [OK, response.to_json]
  end

  put '/:id' do
    logger.info("PUT /users/#{params[:id]}")
    logger.debug('Updating user with lab data service...')
    # PUT user to lab data service
    response = { 'user' => {} }
    logger.debug("Successfully returning user '#{params[:id]}' - RETURN: #{OK}")
    [OK, response.to_json]
  end

  delete '/:id' do
    logger.info("DELETE /users/#{params[:id]}")
    logger.debug('Deleting user with lab data service...')

    delete_message = HTTParty.delete("http://localhost:6001/api/users/#{params[:id]}")
    response = { 'message' => delete_message }

    logger.debug("Successfully deleted user '#{params[:id]}' - RETURN: #{OK}")
    [OK, response.to_json]
  end
end
