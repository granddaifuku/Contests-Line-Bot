import db
import datetime
import utils

CF_URL = 'https://codeforces.com/api/contest.list?gym=false'
CF_INFO_PATH = '/app/data/cf_info.json'
CF_TMP_PATH = '/app/data/cf_tmp.json'


def get_upcoming_cf_contests():
    JST = datetime.timezone(datetime.timedelta(hours=+9), 'JST')
    contents = utils.get_data(CF_URL)
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


def send_cf_info():
    contents = utils.template_json_data(CF_TMP_PATH)
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
                contests_name = utils.template_json_data(
                    utils.CONTESTS_NAME_TMP)
                contests_time = utils.template_json_data(
                    utils.CONTESTS_TIME_TMP)
                contests_name['contents'][1]['text'] = data[i]['name']
                contests_time['contents'][1]['text'] = data[i]['time']
                contents['body']['contents'][1]['contents'].append(
                    contests_name)
                contents['body']['contents'][1]['contents'].append(
                    contests_time)

    return contents


def format_cf_info():
    res = []
    data = get_upcoming_cf_contests()

    if len(data) == 0:
        info = utils.template_json_data(CF_INFO_PATH)
        for i in info:
            info[i] = '-'
        res.append(info)
    else:
        for i in range(len(data) // 2):
            info = utils.template_json_data(CF_INFO_PATH)
            index = 0
            for j in info:
                info[j] = data[i * 2 + index]
                index = index + 1
            res.append(info)

    return res
