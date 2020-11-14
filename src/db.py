import os
import psycopg2
import psycopg2.extras
import atcoder
import codeforces
import yukicoder


AT_TABLE = 'at_db'
CF_TABLE = 'cf_db'
YK_TABLE = 'yk_db'


def get_connection():
    db_url = os.environ.get('DATABASE_URL')

    return psycopg2.connect(db_url)


def update_at_table():
    query = ''
    query += 'DELETE FROM {};'.format(AT_TABLE)
    data = atcoder.format_at_info()
    for i in range(len(data)):
        query += 'INSERT INTO {0} (name, time, range) VALUES (\'{1}\', \'{2}\', \'{3}\');'.format(
            AT_TABLE, data[i]['name'], data[i]['time'], data[i]['range'])
    execute(query)


def update_cf_table():
    query = ''
    query += 'DELETE FROM {};'.format(CF_TABLE)
    data = codeforces.format_cf_info()
    for i in range(len(data)):
        query += 'INSERT INTO {0} (name, time) VALUES (\'{1}\', \'{2}\');'.format(
            CF_TABLE, data[i]['name'], data[i]['time'])
    execute(query)


def update_yk_table():
    query = ''
    query += 'DELETE FROM {};'.format(YK_TABLE)
    data = yukicoder.format_yk_info()
    for i in range(len(data)):
        query += 'INSERT INTO {0} (name, time) VALUES (\'{1}\', \'{2}\');'.format(
            CF_TABLE, data[i]['name'], data[i]['time'])
    execute(query)


def get_records(table_name, range=True):
    query = ''
    if range:
        query += 'SELECT name, time, range FROM {};'.format(table_name)
    else:
        query += 'SELECT name, time FROM {};'.format(table_name)
    res = execute(query, False)

    return res


def execute(query, Insert=True):
    with get_connection() as conn:
        if Insert:
            with conn.cursor() as cur:
                cur.execute(query)
                conn.commit()
        else:
            with conn.cursor(cursor_factory=psycopg2.extras.DictCursor) as cur:
                cur.execute(query)
                res = cur.fetchall()
                return res


if __name__ == '__main__':
    update_at_table()
    update_cf_table()
