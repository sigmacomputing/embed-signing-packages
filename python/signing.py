import hashlib
import hmac
import time
import jwt

try:
    from urllib.parse import quote
except:
    from urllib import quote
import uuid

# Replace with your own values
EMBED_PATH = "your path here"
CLIENT_ID = "your clientid here"
EMBED_SECRET = "your secret here"


def urlencode(pairs):
    # Sigma doesn't support percent-encoded colons and commas in date range params
    # Percent-encoding those characters isn't required by RFC, so explicitly don't
    # encode them!
    def myquote(val):
        return quote(
            str(val),
            safe=",:",
        )

    return "&".join(myquote(k) + "=" + myquote(v) for k, v in pairs.items())


def generateSignedEmbedUrl():

    params = {
        ":nonce": uuid.uuid4(),
        ":email": "xyz@xyz.com",
        ":external_user_id": "1",
        ":client_id": CLIENT_ID,
        ":time": time.time(),
        ":session_length": 3600,
        ":mode": "view",
        ":mode": "userbacked",
        ":external_user_team": "Embedded Users,EmbeddingTown",
        ":account_type": "embedUser",
        # custom controls/parameters
        # "Store-Region": "West",
    }
    url_with_params = EMBED_PATH + "?" + urlencode(params)

    signature = hmac.new(
        EMBED_SECRET.encode("utf-8"), url_with_params.encode("utf-8"), hashlib.sha256
    ).hexdigest()

    url_with_signature = url_with_params + "&" + urlencode({":signature": signature})

    return url_with_signature


def generateJWTEmbedUrl():
    payload = {
        "iss": CLIENT_ID,
        "sub": "xyz@xyz.com",
        "jti": str(uuid.uuid4()),
        "ver": "1.1",
        "aud": "sigmacomputing",
        "iat": int(time.time()),
        "exp": int(time.time()) + 3600,
        "account_type": "Pro",
        "teams": ["Embedded Users", "EmbeddingTown"],
    }
    token = jwt.encode(
        payload=payload,
        key=EMBED_SECRET,
        algorithm="HS256",
        headers={ "kid": CLIENT_ID },
    )
    url_with_token = (
        "https://app.sigmacomputing.com/<your org>/<your workbook>?:embed=true&:jwt="
        + token
    )
    return url_with_token


def main():

    print("========== JWT Embed URL ==========")
    print(generateJWTEmbedUrl())
    print("====================================")

    print("========== Signed Embed URL ==========")
    print(generateSignedEmbedUrl())
    print("====================================")


if __name__ == "__main__":
    main()
