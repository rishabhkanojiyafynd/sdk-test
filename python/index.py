from pixelbin import PixelbinClient, PixelbinConfig
import json


url = "IMAGE_URL.png"


# create client with your API_TOKEN
config = PixelbinConfig({
    "domain": "https://api.pixelbin.io",
    "apiSecret": "API_TOKEN",
})

# Create a pixelbin instance
pixelbin: PixelbinClient = PixelbinClient(config=config)


try:
    result = pixelbin.transformation.getTransformationContext(url=url)
    json_object = json.dumps(result, indent=4)
    print(json_object)
except Exception as e:
    print(e)
