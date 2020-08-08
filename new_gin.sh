# !/bin/bash

# コマンドの打たれ方
# aaa=プロジェクト名
# bbb=ディレクトリ指定（デフォルトは"../"）
# ./new_gin aaa bbb

project_name=$1

# もしプロジェクト名がなかったら終了
if [ $# = 0 ]; then
  echo "プロジェクト名を指定してください"
  exit;
fi

make_path=$2

# もしpath指定がなかったら"../"を指定
if [ $# = 1  ]; then
  make_path="../"
fi

# 1.aaaで同じディレクトリにフォルダを作成
mkdir -p $make_path/$project_name

# 全てのディレクトリをコピーしaaaに移動
cp -R ./ $make_path/$project_name

# 移動後app_nameとなっているところを全てaaaに変更
LC_ALL=C find $make_path/$project_name -type f -print0 | xargs -0 sed  -i -e "s/app_name/$project_name/g"

find $make_path/$project_name -type f -name "*-e" -exec rm {} \;

# このファイルを削除
rm ./new_gin.sh
