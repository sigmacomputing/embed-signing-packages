const crypto = require('crypto');
const { URLSearchParams } = require('url');

// Replace with your own values
const EMBED_PATH   = "your path here"
const CLIENT_ID    = "your clientid here"
const EMBED_SECRET = "your secret here"

function urlencode(pairs) {
  const params = new URLSearchParams();
  for (const [key, value] of Object.entries(pairs)) {
    params.append(key, value);
  }
  return params.toString();
}

const params = {
  ":nonce": crypto.randomBytes(16).toString('hex'),
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

const signature = crypto.createHmac('sha256', Buffer.from(EMBED_SECRET, 'utf8'))
  .update(Buffer.from(urlWithParams, 'utf8'))
  .digest('hex');

const urlWithSignature = `${urlWithParams}&:signature=${signature}`;

console.log(urlWithSignature);
