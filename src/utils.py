from bs4 import  BeautifulSoup
import datetime
import urllib3
import json

AT_URL = 'https://atcoder.jp/contests/'
CF_URL = 'https://codeforces.com/api/contest.list?gym=false'

AT_INFO_PATH = '../data/at_info.json'
CF_INFO_PATH = '../data/cf_info.json'

AT_TMP_PATH = '../data/at_tmp.json'
CF_TMP_PATH = '../data/cf_tmp.json'

CONTESTS_NAME_TMP = '../data/contests_name.json'
CONTESTS_TIME_TMP = '../data/contests_time.json'
CONTESTS_RANGE_TMP = '../data/contests_range.json'


def template_json_data(path):
    with open(path) as f:
        jsn = json.load(f)
    
    return jsn


def get_data(url):
    http = urllib3.PoolManager()
    response = http.request('GET', url)

    return response


def get_upcoming_at_contests():
    r = get_data(AT_URL)
    soup = BeautifulSoup(r.data, 'html.parser')
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
    r = get_data(CF_URL)
    contents = json.loads(r.data.decode('utf-8'))
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
    data = get_upcoming_at_contests()
    if len(data) == 0:
        for i in range(3):
            contents['body']['contents'][1]['contents'][i]['contents'][1]['text'] = '-'
    else:
        for i in range(len(data)):
            if i <= 2:
                contents['body']['contents'][1]['contents'][i]['contents'][1]['text'] = data[i]
            else:
                if i % 3 == 0:
                    contests_info = template_json_data(CONTESTS_NAME_TMP)
                elif i % 3 == 1:
                    contests_info = template_json_data(CONTESTS_TIME_TMP)
                else:
                    contests_info = template_json_data(CONTESTS_RANGE_TMP)
                contests_info['contents'][1]['text'] = data[i]
                contents['body']['contents'][1]['contents'].append(contests_info)
        
    return contents


def send_cf_info():
    contents = template_json_data(CF_TMP_PATH)
    data = get_upcoming_cf_contests()
    if len(data) == 0:
        for i in range(2):
            contents['body']['contents'][1]['contents'][i]['contents'][1]['text'] = '-'
    else:
        for i in range(len(data)):
            if i <= 1:
                contents['body']['contents'][1]['contents'][i]['contents'][1]['text'] = data[i]
            else:
                if i % 2 == 0:
                    contests_info = template_json_data(CONTESTS_NAME_TMP)
                else:
                    contests_info = template_json_data(CONTESTS_TIME_TMP)
                contests_info['contents'][1]['text'] = data[i]
                contents['body']['contents'][1]['contents'].append(contests_info)
    
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


if __name__ == '__main__':
    print(format_at_info())
