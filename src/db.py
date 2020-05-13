import utils
import psycopg2
import os 

AT_TABLE = 'at_db'
CF_TABLE = 'cf_db'

def get_connection():
    dst = os.environ.get('DATABASE_URL')

    return psycopg2.connect(dst)


def update_at_table():
    sql_string = ''
    sql_string += 'DELETE FROM {};'.format(AT_TABLE)
    data = utils.format_at_info()
    for i in range(len(data)):
        sql_string += 'INSERT INTO {0} (contest) VALUES (\'{1}\');'.format(AT_TABLE, data[i])
    print(sql_string)
    execute(AT_TABLE)


def update_cf_table():
    sql_string = ''
    sql_string += 'DELETE FROM {};'.format(CF_TABLE)
    data = utils.format_cf_info()
    for i in range(len(data)):
        sql_string += 'INSERT INTO {0} (contest) VALUES (\'{1}\');'.format(CF_TABLE, data[i])
    print(sql_string)
    execute(CF_TABLE)


def execute(query):
    conn = get_connection()
    cur = conn.cursor()
    cur.execute(query)

if __name__ == '__main__':
    update_cf_table()