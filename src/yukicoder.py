import db
import datetime
import utils

YUKICODER_URL = 'https://yukicoder.me/api/v1/contest/future'
YK_INFO_PATH = '/app/data/yk_info.json'
YK_TMP_PATH = '/app/data/yk_tmp.json'


def get_upcoming_yukicoder_contests():
    contests = utils.get_data(YUKICODER_URL)
    if len(contests) == 0:
        print("No contents to display")
        return
    res = []
    for contest in contests:
        name = contest['Name']
        start = datetime.datetime.fromisoformat(
            contest['Date']).replace(tzinfo=None)
        end = datetime.datetime.fromisoformat(
            contest['EndDate']).replace(tzinfo=None)
        s = ''
        s += datetime.datetime.strftime(start, '%Y-%m-%d %H:%M:%S')
        s += ' - '
        s += datetime.datetime.strftime(end, '%Y-%m-%d %H:%M:%S')
        res.insert(0, name)
        res.insert(1, s)

    return res


def send_yk_info():
    contents = utils.template_json_data(YK_TMP_PATH)
    data = db.get_records(db.YK_TABLE, False)

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


def format_yk_info():
    res = []
    data = get_upcoming_yukicoder_contests()

    if len(data == 0):
        info = utils.template_json_data(YK_INFO_PATH)
        for i in info:
            info[i] = '-'
        res.append(info)
    else:
        for i in range(len(data) // 2):
            info = utils.template_json_data(YK_INFO_PATH)
            index = 0
            for j in info:
                info[j] = data[i * 2 + index]
                index = index + 1
            res.append(info)

    return res


if __name__ == "__main__":
    get_upcoming_yukicoder_contests()
