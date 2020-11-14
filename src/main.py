from flask import Flask, request, abort
import os
import atcoder
import codeforces
import yukicoder

from linebot import (
    LineBotApi, WebhookHandler
)
from linebot.exceptions import (
    InvalidSignatureError, LineBotApiError
)
from linebot.models import (
    MessageEvent, TextMessage, FlexSendMessage
)

app = Flask(__name__)

CHANNEL_ACCESS_TOKEN = os.environ.get('LINE_CHANNEL_ACCESS_TOKEN')
CHANNEL_SECRET = os.environ.get('LINE_CHANNEL_SECRET')

line_bot_api = LineBotApi(CHANNEL_ACCESS_TOKEN)
handler = WebhookHandler(CHANNEL_SECRET)


@app.route("/callback", methods=['POST'])
def callback():
    signature = request.headers['X-Line-Signature']
    body = request.get_data(as_text=True)
    app.logger.info('Request body: ' + body)

    try:
        handler.handle(body, signature)
    except InvalidSignatureError:
        print('Invalid signature. Please check your channel access token/channel secret.')
        abort(400)
    return 'OK'


@handler.add(MessageEvent, message=TextMessage)
def handle_message(event):
    user_id = event.source.user_id
    to = user_id
    if hasattr(event.source, "group_id"):
        to = event.source.group_id
    TARGET = 'コンテスト'
    if not TARGET in event.message.text:
        return
    cf_data = codeforces.send_cf_info()
    cf_message = FlexSendMessage(
        alt_text='hello',
        contents=cf_data
    )
    at_data = atcoder.send_at_info()
    at_message = FlexSendMessage(
        alt_text='hello',
        contents=at_data
    )
    yk_data = yukicoder.send_yk_info()
    yk_message = FlexSendMessage(
        alt_text='hello',
        contents=yk_data
    )

    try:
        line_bot_api.push_message(
            to,
            messages=cf_message)
        line_bot_api.push_message(
            to,
            messages=at_message)
        line_bot_api.push_message(
            to,
            messages=yk_message)
    except LineBotApiError:
        print('Failed to Send Contests Information')


if __name__ == '__main__':
    port = int(os.getenv('PORT', 5000))
    app.run(host='0.0.0.0', port=port)
