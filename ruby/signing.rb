require 'uri'
require 'net/http'
require 'securerandom'
require 'openssl'

# Replace with your own values
EMBED_PATH   = 'your path here'
CLIENT_ID    = 'your clientid here'
EMBED_SECRET = 'your secret here'

def urlencode(pairs)
  URI.encode_www_form(pairs)
end

params = {
  ":nonce": SecureRandom.uuid,
  ":email": "test@test.com",
  ":external_user_id": "test@test.com",
  ":client_id": CLIENT_ID,
  ":time": Time.now.to_i,
  ":session_length": 3600,
  ":mode": "userbacked",
  ":external_user_team": "Embeddingtown,EmbedTeam",
}

url_with_params = "#{EMBED_PATH}?#{urlencode(params)}"

signature = OpenSSL::HMAC.hexdigest('sha256', EMBED_SECRET, url_with_params)

url_with_signature = "#{url_with_params}&:signature=#{signature}"

puts url_with_signature
