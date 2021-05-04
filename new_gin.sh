# !/bin/bash

read -p "プロジェクト名を入力してください: " project_name

# もしプロジェクト名がなかったら終了
if [[ $project_name = ""  ]]; then
  echo "プロジェクト名を指定してください"
  exit;
fi

read -p "作成するディレクトリを入力してください（指定がない場合は./配下に作成されます）: " make_path

# もしpath指定がなかったら"../"を指定
if [[ $make_path = ""  ]]; then
  make_path=".."
fi

read -p "通信方式を選択してください。指定がない場合ははapiになります（api/grpc）: " connection
# 通信方式を選択されなかったら"api"を指定
if [[ $connection = ""  ]]; then
  connection="api"
fi

# grpcの場合swaggerディレクトリを削除
if [[ $connection = "grpc" ]]; then
  rm -rf $make_path/$project_name/docker/swagger
else
  rm -rf $make_path/$project_name/proto
fi



# 1.aaaで同じディレクトリにフォルダを作成
mkdir -p $make_path/$project_name

# 全てのディレクトリをコピーしaaaに移動
cp -R ./ $make_path/$project_name

# 移動後app_nameとなっているところを全てaaaに変更
grep -lr 'app_name' $make_path/$project_name/ | xargs sed -i -e "s/app_name/${project_name}/g"
find $make_path/$project_name/* -name "*-e" -exec rm {} \;

# makefileを適切なものに変更
rm $make_path/$project_name/Makefile
mv $make_path/$project_name/Makefile_new $make_path/$project_name/Makefile

# 不要なmigration fileを削除
rm $make_path/$project_name/app/db/migrations/*.sql

# このファイルを削除
rm $make_path/$project_name/new_gin.sh

echo '🎉作成が完了しました🎉'
echo 'cloud runを使うときは「variable.tf」のproject変数の内容を変更してください'
echo 'またcloud run apiを有効化してください'
