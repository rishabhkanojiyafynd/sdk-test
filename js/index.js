const { PixelbinConfig, PixelbinClient } = require("@pixelbin/admin");

let API_MAIN_DOMAIN = "https://api.pixelbin.io";
let apiToken = "API_TOKEN";

let url = "IMAGE_URL.png";

const config = new PixelbinConfig({
  domain: API_MAIN_DOMAIN,
  apiSecret: apiToken,
});

const pixelbin = new PixelbinClient(config);
const main = async () => {
  const response = await pixelbin.transformation.getTransformationContext({
    url,
  });
  console.log(JSON.stringify(response, null, 2));
};
main();
