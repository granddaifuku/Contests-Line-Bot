import urllib.request
import json

CONTESTS_NAME_TMP = '/app/data/contests_name.json'
CONTESTS_TIME_TMP = '/app/data/contests_time.json'
CONTESTS_RANGE_TMP = '/app/data/contests_range.json'


def template_json_data(path):
    with open(path) as f:
        jsn = json.load(f)

    return jsn


def get_data(url, scp=False):
    req = urllib.request.Request(url)
    try:
        with urllib.request.urlopen(req) as res:
            if scp:
                return res.read()
            else:
                return json.load(res)
    except urllib.error.HTTPError as err:
        print(err.code)
    except urllib.error.URLError as err:
        print(err.reason)
