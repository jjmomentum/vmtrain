# Message Model
class Message
  def initialize(message)
    @message = message
  end

  def to_json
    { 'message' => @message }.to_json
  end

  def to_xml
    builder = Nokogiri::XML::Builder.new do |xml|
      xml.response do
        xml.message @message
      end
    end
    builder.to_xml
  end
end
