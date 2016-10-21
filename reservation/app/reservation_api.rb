require 'base64'

# Reservation APIs
class ReservationApi < ApiBase
  get '/' do
    logger.info('GET /reservations')
    logger.debug('Querying lab data service for reservations...')
    reservations = HTTParty.get('http://localhost:6001/api/reservations')
    response = { 'reservations' => reservations }
    logger.debug("Successfully returning reservations - RETURN: #{OK}")
    [OK, response.to_json]
  end

  post '/' do
    logger.info('POST /reservations')
    logger.debug('Creating reservation with lab data reservation...')
    reservation = HTTParty.post('http://localhost:8080/api/topic/reservation_test',
      :headers => {
        'Content-type' => 'application/json'
      },
      :body => {
        'message': Base64.encode64(params.to_s)
      }.to_json
    )
    response = { 'reservation' => reservation }
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
