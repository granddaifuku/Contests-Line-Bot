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
    data = utils.get_upcoming_at_contests()
    execute(AT_TABLE)


def update_cf_table():
    sql_string = ''
    sql_string += 'DELETE FROM {};'.format(CF_TABLE)
    data = utils.get_upcoming_cf_contests()
    execute(CF_TABLE)


def execute(table_name, query):
    conn = get_connection()
    cur = conn.cursor()
    sql_string = 'INSERT INTO {}'.format(table_name)
    cur.execute(sql_string)
    cur.execute(query)