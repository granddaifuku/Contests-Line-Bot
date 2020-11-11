import db
import utils
from bs4 import BeautifulSoup
import datetime
import re

AT_URL = 'https://atcoder.jp/contests/'
AT_INFO_PATH = '/app/data/at_info.json'
AT_TMP_PATH = '/app/data/at_tmp.json'


def get_upcoming_at_contests():
    r = utils.get_data(AT_URL, True)
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
                split_time = re.split('[:]', text[i + 1])
                dur_hour = int(split_time[0])
                day = dur_hour // 24
                hour = dur_hour % 24
                minutes = int(split_time[1])
                if day == 0:
                    dur = datetime.datetime.strptime(text[i + 1], '%H:%M')
                    end = start + \
                        datetime.timedelta(hours=int(dur.strftime(
                            '%H')), minutes=int(dur.strftime('%M')))
                else:
                    text[i + 1] = str(day) + ":" + \
                        str(hour) + ":" + str(minutes)
                    dur = datetime.datetime.strptime(text[i + 1], '%d:%H:%M')
                    end = start + \
                        datetime.timedelta(days=int(dur.strftime('%d')), hours=int(dur.strftime(
                            '%H')), minutes=int(dur.strftime('%M')))
                s += ' - '
                s += end.strftime('%Y-%m-%d %H:%M:%S')
            text[i] = s
        if i % 4 != 2:
            res.append(text[i])

    return res


def send_at_info():
    contents = utils.template_json_data(AT_TMP_PATH)
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
                contests_name = utils.template_json_data(
                    utils.CONTESTS_NAME_TMP)
                contests_time = utils.template_json_data(
                    utils.CONTESTS_TIME_TMP)
                contests_range = utils.template_json_data(
                    utils.CONTESTS_RANGE_TMP)
                contests_name['contents'][1]['text'] = data[i]['name']
                contests_time['contents'][1]['text'] = data[i]['time']
                contests_range['contents'][1]['text'] = data[i]['range']
                contents['body']['contents'][1]['contents'].append(
                    contests_name)
                contents['body']['contents'][1]['contents'].append(
                    contests_time)
                contents['body']['contents'][1]['contents'].append(
                    contests_range)

    return contents


def format_at_info():
    res = []
    data = get_upcoming_at_contests()

    if len(data) == 0:
        info = utils.template_json_data(AT_INFO_PATH)
        for i in info:
            info[i] = '-'
        res.append(info)
    else:
        for i in range(len(data) // 3):
            info = utils.template_json_data(AT_INFO_PATH)
            index = 0
            for j in info:
                info[j] = data[i * 3 + index]
                index = index + 1
            res.append(info)

    return res
