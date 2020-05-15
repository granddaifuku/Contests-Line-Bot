from bs4 import  BeautifulSoup
import datetime
import urllib.request
import json
import db

AT_URL = 'https://atcoder.jp/contests/'
CF_URL = 'https://codeforces.com/api/contest.list?gym=false'

AT_INFO_PATH = '/app/data/at_info.json'
CF_INFO_PATH = '/app/data/cf_info.json'

AT_TMP_PATH = '/app/data/at_tmp.json'
CF_TMP_PATH = '/app/data/cf_tmp.json'

CONTESTS_NAME_TMP = '/app/data/contests_name.json'
CONTESTS_TIME_TMP = '/app/data/contests_time.json'
CONTESTS_RANGE_TMP = '/app/data/contests_range.json'


def template_json_data(path):
    with open(path) as f:
        jsn = json.load(f)
    
    return jsn


def get_data(url, scp = False):
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


def get_upcoming_at_contests():
    r = get_data(AT_URL, True)
    soup = BeautifulSoup(r, 'html.parser')
    texts = soup.get_text()
    words = [line.strip() for line in texts.splitlines()]
    upcoming = False
    text = []
    for word in words:
        if word == '◉' or word == '':
            continue
        if word == 'Upcoming Contests':
            upcoming = True
            continue
        if word == 'Recent Contests':
            upcoming = False
        if upcoming:
            text.append(word)
    res = []
    for i in range(len(text)):
        if i < 4:
            continue
        if i % 4 == 0:
            text[i], text[i + 1] = text[i + 1], text[i]
        if i % 4 == 1:
            s = ''
            if i == 1:
                pass
            else:
                for t in text[i]:
                    if t == '+':
                        break
                    s += t
                start = datetime.datetime.strptime(s, '%Y-%m-%d %H:%M:%S')
                dur = datetime.datetime.strptime(text[i + 1], '%H:%M')
                end = start + datetime.timedelta(hours=int(dur.strftime('%H')), minutes=int(dur.strftime('%M')))
                s += ' - '
                s += end.strftime('%Y-%m-%d %H:%M:%S')
            text[i] = s
        if i % 4 != 2:
            res.append(text[i])

    return res


def get_upcoming_cf_contests():
    JST = datetime.timezone(datetime.timedelta(hours=+9), 'JST')
    contents = get_data(CF_URL)
    if contents['status'] == 'FAILED':
        print('Failed to call CF API')
        return
    res = []
    for i in range(len(contents['result'])):
        if (contents['result'][i]['phase'] == 'FINISHED'):
            break
        res.insert(0, contents['result'][i]['name'])
        start = contents['result'][i]['startTimeSeconds']
        s = ''
        start_jst = datetime.datetime.fromtimestamp(start, JST)
        start_time = datetime.datetime.strftime(start_jst, '%Y-%m-%d %H:%M:%S')
        s += start_time
        dur_sec = contents['result'][i]['durationSeconds']
        dur = datetime.timedelta(seconds=dur_sec)
        end_time = start_jst + dur
        s += ' - '
        s += end_time.strftime('%Y-%m-%d %H:%M:%S')
        res.insert(1, s)
    
    return res


def send_at_info():
    contents = template_json_data(AT_TMP_PATH)
    data = db.get_records(db.AT_TABLE)
        
    if len(data) == 0:
        for i in range(3):
            contents['body']['contents'][1]['contents'][i]['contents'][1]['text'] = '-'
    else:
        for i in range(len(data)):
            if i == 0:
                contents['body']['contents'][1]['contents'][0]['contents'][1]['text'] = data[i]['name']
                contents['body']['contents'][1]['contents'][1]['contents'][1]['text'] = data[i]['time']
                contents['body']['contents'][1]['contents'][2]['contents'][1]['text'] = data[i]['range']
            else:
                contests_name = template_json_data(CONTESTS_NAME_TMP)
                contests_time = template_json_data(CONTESTS_TIME_TMP)
                contests_range = template_json_data(CONTESTS_RANGE_TMP)
                contests_name['contents'][1]['text'] = data[i]['name']
                contests_time['contents'][1]['text'] = data[i]['time']
                contests_range['contents'][1]['text'] = data[i]['range']
                contents['body']['contents'][1]['contents'].append(contests_name)
                contents['body']['contents'][1]['contents'].append(contests_time)
                contents['body']['contents'][1]['contents'].append(contests_range)
        
    return contents


def send_cf_info():
    contents = template_json_data(CF_TMP_PATH)
    data = db.get_records(db.CF_TABLE, False)

    if len(data) == 0:
        for i in range(2):
            contents['body']['contents'][1]['contents'][i]['contents'][1]['text'] = '-'
    else:
        for i in range(len(data)):
            if i == 0:
                contents['body']['contents'][1]['contents'][0]['contents'][1]['text'] = data[i]['name']
                contents['body']['contents'][1]['contents'][1]['contents'][1]['text'] = data[i]['time']
            else:
                contests_name = template_json_data(CONTESTS_NAME_TMP)
                contests_time = template_json_data(CONTESTS_TIME_TMP)
                contests_name['contents'][1]['text'] = data[i]['name']
                contests_time['contents'][1]['text'] = data[i]['time']
                contents['body']['contents'][1]['contents'].append(contests_name)
                contents['body']['contents'][1]['contents'].append(contests_time)
    
    return contents


def format_at_info():
    res = []
    data = get_upcoming_at_contests()

    if len(data) == 0:
        info = template_json_data(AT_INFO_PATH)
        for i in info:
            info[i] = '-'
        res.append(info)
    else:
        for i in range(len(data) // 3):
            info = template_json_data(AT_INFO_PATH)
            index = 0
            for j in info:
                info[j] = data[i * 3 + index]
                index = index + 1
            res.append(info)

    return res


def format_cf_info():
    res = []
    data = get_upcoming_cf_contests()

    if len(data) == 0:
        info = template_json_data(CF_INFO_PATH)
        for i in info:
            info[i] = '-'
        res.append(info)
    else:
        for i in range(len(data) // 2):
            info = template_json_data(CF_INFO_PATH)
            index = 0
            for j in info:
                info[j] = data[i * 2 + index]
                index = index + 1
            res.append(info)

    return res