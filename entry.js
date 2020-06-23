const config = {
  port: 30000,
  upnp: true,
  fullscreen: false,
  hostname: null,
  routePrefix: null,
  sslCert: null,
  sslKey: null,
  awsConfig: null,
  proxySSL: false,
  proxyPort: null,
  updateChannel: "beta",
  world: null,
};
const envMap = {
  port: "PORT",
  upnp: "UPNP",
  fullscreen: "FULLSCREEN",
  hostname: "HOSTNAME",
  routePrefix: "ROUTE_PREFIX",
  sslCert: "SSL_CERT",
  sslKey: "SSL_KEY",
  awsConfig: "AWS_CONFIG",
  proxySSL: "PROXY_SSL",
  proxyPort: "PROXY_PORT",
  updateChannel: "UPDATE_CHANNEL",
  world: "WORLD",
};

function bool(str) {
  return str === "true" || str === "1";
}

const configCast = {
  port: Number,
  upnp: bool,
  fullscreen: bool,
  hostname: String,
  routePrefix: String,
  sslCert: String,
  sslKey: String,
  awsConfig: String,
  proxySSL: String,
  proxyPort: String,
  updateChannel: String,
  world: String,
};

for (const key in config) {
  if (config.hasOwnProperty(key)) {
    if (process.env[envMap[key]]) {
      config[key] = configCast[key](process.env[envMap[key]]);
    }
  }
}

console.log(JSON.stringify(config, undefined, "    "));
