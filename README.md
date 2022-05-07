# rakuten-card-statement-bot

楽天カード明細情報をLINEに通知するBOT

## Configuration
### LINE Messaging API のアクセス情報を設定

`.env.example`をコピーし、`.env`を作成。  
作成した`.env`に、アクセス情報を記述してください。  
_Heroku などのサービスで、定義されている環境変数があれば、サービスの値が優先されて使用されます。_

#### RAKUTEN_SELECT_CARD_NO

取得したいカードのセレクトボックスの順番を記述する。  
1枚目なら**1**、2枚目なら**2**を記述する。  

![カード選択](https://raw.githubusercontent.com/suhrr/readme-images/master/rakuten-card-statement-bot/select_card.png)

## Screenshots
![LINEスクリーンショット](https://raw.githubusercontent.com/suhrr/readme-images/master/rakuten-card-statement-bot/line_1.jpg)
