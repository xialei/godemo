import requests
import json

def get_file_content(filePath):
    with open(filePath, 'rb') as fp:
        return fp.read()

class CommonOcr(object):
    def __init__(self, img_path):
        self._app_key = '5c6b*************************4d85'  # your app_key
        self._app_secret = '5c6b*************************4d85'  # your app_secret
        self._img_path = img_path

    def recognize(self):
        url = 'https://ocr-api.ccint.com/cci_ai/service/v1/business_card?lang=1,2,3&img_info=1'
        head = {}
        try:
            image = get_file_content(self._img_path)
            head['app-key'] = self._app_key
            head['app-secret'] = self._app_secret
            result = requests.post(url, data=image, headers=head)
            return result.text
        except Exception as e:
            return e

if __name__ == "__main__":
    response = CommonOcr(r'example.jpg')
    print(response.recognize())