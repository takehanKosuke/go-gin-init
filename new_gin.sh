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

# 1.aaaで同じディレクトリにフォルダを作成
mkdir -p $make_path/$project_name

# 全てのディレクトリをコピーしaaaに移動
cp -R ./ $make_path/$project_name

# 移動後app_nameとなっているところを全てaaaに変更
LC_ALL=C find $make_path/$project_name -type f -print0 | xargs -0 sed  -i -e "s/app_name/$project_name/g"

find $make_path/$project_name -type f -name "*-e" -exec rm {} \;

# 最初に使用するためのmakeファイルを削除
sed -e '1,3d' $make_path/$project_name/Makefile > $make_path/$project_name/Makefiletmp
mv $make_path/$project_name/Makefiletmp $make_path/$project_name/Makefile

# terraformの秘匿ファイルを作成
echo << EOS
aws_access_key = "awsアクセスキー"
aws_secret_key = "awsシークレットアクセスキー"
EOS >> $make_path/terraform.tfstate

# このファイルを削除
rm $make_path/$project_name/new_gin.sh
