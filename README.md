# ms-protobuf

このリポジトリは、MSプロジェクトにおける **Protobufのスキーマ定義を一元管理** するためのリポジトリです。  
全サービス間で一貫したスキーマ設計・共有を可能にします。

---

## 📁 ディレクトリ構成

```
.
├── ms/ # サービスごとのスキーマ管理ディレクトリ
│ ├── apifront/v1/
│ │ └── api.proto
│ └── user/v1/
│ ├── resource.proto
│ └── user.proto
├── buf.yaml
├── buf.lock
└── README.md
```

---

## 🛠 運用方法

現時点ではCIを導入しておらず、以下の手順で **ローカルから手動で操作** します：

```sh
# Lint（構文チェック）
buf lint

# RegistryへPush（BSR）
buf push
```
