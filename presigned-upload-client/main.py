import ast
import sys
import requests

def get_presigned_url(key: str):
    res = requests.get("http://localhost:8395/s3/PresignUrl/upload?key="+key)
    # res = requests.get("http://google.com")
    body = res.content.decode('utf-8')
    res_dic = ast.literal_eval(body)
    for k,v in res_dic.items():
        print(k, v)
    return res_dic

def upload_with_presigned_url(res_dict):
    res = requests.put(res_dict['URL'], headers=res_dict['SignedHeader'])
    print(res.status_code)

if __name__ == "__main__":
    res_dict = get_presigned_url(sys.argv[1])
    upload_with_presigned_url(res_dict)