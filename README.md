
### vue.js + Go + firebase

下記の素晴らしいハンズオンを実施。

* https://qiita.com/po3rin/items/d3e016d01162e9d9de80


### 利用手順

#### フロントエンド(Vue.js)

* main.jsの変更とfirebase用のAPIキー等の設定

```sh
$ mv src/main.example.js src/main.js
```

* main.jsを開いて、APIキー等の値をセット

```js
import Vue from 'vue'
import App from './App.vue'

import router from './router'
import firebase from 'firebase'

Vue.config.productionTip = false

const config = {
  apiKey: 'YOUR_API_KEY',
  authDomain: 'YOUR_DOMAIN',
  databaseURL: 'YOUR_DATABSE_URL',
  projectId: 'YOUR_PROJECT_ID',
  storageBucket: '',
  messagingSenderId: 'YOUR_SENDER_ID'
}
firebase.initializeApp(config)

firebase.auth().onAuthStateChanged(user => {
  new Vue({
    router,
    render: h => h(App)
  }).$mount('#app')
})

```


* Docker Toolbox以外を利用している場合は、./src/consts.jsのAPI_SERVERの接続先を192.168.99.100からlocalhostに変更

./src/consts.js

```js
export const API_SERVER = 'http://192.168.99.100'
export const PORT_NUMBER = 8000
```


* フロントエンド側の起動

```sh
$ npm install -g yarn
$ yarn install
$ yarn serve
```

#### バックエンド

* backend/app/credentialsに、firebaseから取得したjsonファイルを配置
* .env.exampleを.envに変更

```sh
$ mv .env.example ./.env
```

* backend/app/credentialsに配置したjsonファイル名をFIREBASE_JSON_FILE_NAME変数に設定

```sh
# firebase ... [Project Overview]->[プロジェクトの設定]->[サービスアカウント]->Firebase Admin SDK->新しい秘密鍵の生成
FIREBASE_JSON_FILE_NAME=your_firebase_json_file.json
```

* Dockerコンテナ起動

```sh
$ docker-compose up --build
```

#### 接続URL

* フロントエンド

http://localhost:8080/signin

