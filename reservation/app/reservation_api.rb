# Reservation APIs
class ReservationApi < ApiBase
  get '/' do
    logger.info('GET /reservations')
    logger.debug('Querying lab data service for reservations...')
    # Fetch reservations from lab data service
    response = { 'reservations' => [] }
    logger.debug("Successfully returning reservations - RETURN: #{OK}")
    [OK, response.to_json]
  end

  post '/' do
    logger.info('POST /reservations')
    logger.debug('Creating user with lab data reservation...')
    # POST reservation to lab data service
    response = { 'reservation' => {} }
    logger.debug("Successfully returning new reservation - RETURN: #{OK}")
    [OK, response.to_json]
  end

  get '/:id' do
    logger.info("POST /reservations/#{params[:id]}")
    logger.debug('Querying lab data service for reservation...')
    # Fetch reservation from lab data service
    response = { 'reservation' => {} }
    logger.debug("Successfully returning reservation '#{params[:id]}' - RETURN: #{OK}")
    [OK, response.to_json]
  end

  put '/:id' do
    logger.info("PUT /reservations/#{params[:id]}")
    logger.debug('Updating reservation with lab data service...')
    # PUT reservation to lab data service
    response = { 'reservation' => {} }
    logger.debug("Successfully returning reservation '#{params[:id]}' - RETURN: #{OK}")
    [OK, response.to_json]
  end

  delete '/:id' do
    logger.info("DELETE /reservations/#{params[:id]}")
    logger.debug('Deleting reservation with lab data service...')
    # DELETE reservation with lab data service
    logger.debug("Successfully deleted reservation '#{params[:id]}' - RETURN: #{OK}")
    [OK, response.to_json]
  end
end
