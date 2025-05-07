const crypto = require("crypto");
const { URLSearchParams } = require("url");
const jwt = require("jsonwebtoken");

// Replace with your own values
const EMBED_PATH = "your path here";
const CLIENT_ID = "your clientid here";
const EMBED_SECRET = "your secret here";

function urlencode(pairs) {
  const params = new URLSearchParams();
  for (const [key, value] of Object.entries(pairs)) {
    params.append(key, value);
  }
  return params.toString();
}

function generateSignedEmbedUrl() {
  const params = {
    ":nonce": crypto.randomBytes(16).toString("hex"),
    ":email": "xyz@xyz.com",
    ":external_user_id": "xyz@xyz.com",
    ":client_id": CLIENT_ID,
    ":time": Math.floor(new Date().getTime() / 1000),
    ":session_length": 3600,
    ":mode": "view",
    ":external_user_team": "EmbedTeam",
    // custom controls/parameters
    // "Store-Region": "West",
  };

  const urlWithParams = `${EMBED_PATH}?${urlencode(params)}`;

  const signature = crypto
    .createHmac("sha256", Buffer.from(EMBED_SECRET, "utf8"))
    .update(Buffer.from(urlWithParams, "utf8"))
    .digest("hex");

  const urlWithSignature = `${urlWithParams}&:signature=${signature}`;
  return urlWithSignature;
}

function generateJwtEmbedUrl() {
  const payload = {
    iss: CLIENT_ID,
    sub: "xyz@xyz.com",
    aud: "sigmacomputing",
    jti: crypto.randomBytes(16).toString("hex"),
    ver: "1.1",
    iat: Math.floor(new Date().getTime() / 1000),
    exp: Math.floor(new Date().getTime() / 1000) + 3600,
    teams: ["EmbedTeam"],
    account_type: "Pro",
  };

  const token = jwt.sign(payload, EMBED_SECRET, {
    algorithm: "HS256",
    keyid: CLIENT_ID,
  });
  const urlWithToken = `https://app.sigmacomputing.com/<your org>/<your workbook>?:embed=true&:jwt=${token}`;
  return urlWithToken;
}

function main() {
  const jwtUrl = generateJwtEmbedUrl();
  console.log(
    `===========JWT Embed URL============
    ${jwtUrl}
    ==========================================`
  );

  const urlWithSignature = generateSignedEmbedUrl();
  console.log(
    `===========Signed Embed URL============
    ${urlWithSignature}
    ==========================================`
  );
}

main();
