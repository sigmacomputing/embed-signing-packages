require 'uri'
require 'net/http'
require 'securerandom'
require 'openssl'
require 'jwt'

# Replace with your own values
EMBED_PATH   = 'your path here'
CLIENT_ID    = 'your clientid here'
EMBED_SECRET = 'your secret here'

def urlencode(pairs)
  URI.encode_www_form(pairs)
end

def generate_signed_embed_url()
  params = {
    ":nonce": SecureRandom.uuid,
    ":email": "xyz@xyz.com",
    ":external_user_id": "xyz@xyz.com",
    ":client_id": CLIENT_ID,
    ":time": Time.now.to_i,
    ":session_length": 3600,
    ":mode": "userbacked",
    ":external_user_team": "EmbedTeam",
    # custom controls/parameters
    # "Store-Region": "West",
  }

  url_with_params = "#{EMBED_PATH}?#{urlencode(params)}"

  signature = OpenSSL::HMAC.hexdigest('sha256', EMBED_SECRET, url_with_params)

  url_with_signature = "#{url_with_params}&:signature=#{signature}"
  return url_with_signature
end

def generate_jwt_embed_url()
  payload = {
    "iss": CLIENT_ID,
    "sub": "xyz@xyz.com",
    "jti": SecureRandom.uuid,
    "ver": "1.1",
    "aud": "sigmacomputing",
    "iat": Time.now.to_i,
    "exp": Time.now.to_i + 3600,
    "account_type": "Pro",
    "teams": ["EmbedTeam"],
  }
  token = JWT.encode(payload, EMBED_SECRET, 'HS256')
  return "https://app.sigmacomputing.com/<your-org>/<your-workbook>?:embed=true&:jwt=#{token}"
end

def main()
  url_with_signature = generate_signed_embed_url()
  puts "======= Signed Embed URL ======="
  puts url_with_signature
  puts "==============================="

  url_with_jwt = generate_jwt_embed_url()
  puts "======= JWT Embed URL ======="
  puts url_with_jwt
  puts "==============================="
end

main()
