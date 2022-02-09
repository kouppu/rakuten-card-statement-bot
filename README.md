# rakuten-card-statement-bot

楽天カード明細情報を LINE に通知する BOT

## Configuration

`.env.example`を`.env`に変更し、内容編集する。  
Heroku などのサービスで定義されている環境変数があれば、`.env`ではなくサービスの値が優先されて使用される。

### RAKUTEN_SELECT_CARD_NO

取得したいカードのセレクトボックスの順番を記述する。  
1 枚目なら`1`、2 枚目なら`2`を記述する。  

![カード選択](https://raw.githubusercontent.com/suhrr/readme-images/master/rakuten-card-statement-bot/select_card.png)

## Screenshots
![LINEスクリーンショット](https://raw.githubusercontent.com/suhrr/readme-images/master/rakuten-card-statement-bot/line_1.jpg)
